package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/morozovcookie/atlant"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

//
type ProductChanging struct {
	//
	CreatedAt int64 `bson:"created_at"`

	//
	OldPrice float64 `bson:"old_price"`

	//
	NewPrice float64 `bson:"new_price"`

	//
	RequestID string `bson:"request_id"`

	//
	ChangeID string `bson:"change_id"`
}

//
type Product struct {
	//
	UpdateCount *int `bson:"update_count"`

	//
	CreatedAt int64 `bson:"created_at"`

	//
	UpdatedAt *int64 `bson:"updated_at"`

	//
	Price float64 `bson:"price"`

	//
	ID string `bson:"product_id"`

	//
	Name string `bson:"name"`

	//
	ChangeHistory []ProductChanging `bson:"change_history"`
}

//
type MongoCollector interface {
	Collection() (mc *mongo.Collection)
}

//
type ProductStorage struct {
	//
	products *mongo.Collection

	//
	logger *zap.Logger
}

//
func NewProductStorage(mc MongoCollector, logger *zap.Logger) *ProductStorage {
	return &ProductStorage{
		products: mc.Collection(),

		logger: logger,
	}
}

//
func (ps *ProductStorage) StoreProduct(ctx context.Context, p *atlant.Product) (err error) {
	var (
		filter = bson.D{{Key: "product_id", Value: p.ID()}}
		mp     = &Product{
			CreatedAt:     p.CreatedAt().UnixNano(),
			Price:         p.Price(),
			ID:            p.ID(),
			Name:          p.Name(),
			ChangeHistory: make([]ProductChanging, 0, len(p.ChangeHistory())),
		}
	)

	for _, c := range p.ChangeHistory() {
		mp.ChangeHistory = append(mp.ChangeHistory, ProductChanging{
			OldPrice:  c.OldPrice,
			NewPrice:  c.NewPrice,
			RequestID: c.RequestID,
			ChangeID:  c.ChangeID,
			CreatedAt: c.CreatedAt.UnixNano(),
		})
	}

	if p.UpdateCount() != 0 {
		mp.UpdateCount = new(int)
		*(mp.UpdateCount) = p.UpdateCount()
	}

	if p.UpdatedAt().UnixNano() != 0 {
		mp.UpdatedAt = new(int64)
		*(mp.UpdatedAt) = p.UpdatedAt().UnixNano()
	}

	_, err = ps.products.UpdateOne(ctx, filter, bson.M{"$set": mp}, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	return nil
}

//
func (ps *ProductStorage) GetByProductID(ctx context.Context, productID string) (p *atlant.Product, err error) {
	var (
		mp     = &Product{}
		filter = bson.D{{Key: "product_id", Value: productID}}
	)

	if err = ps.products.FindOne(ctx, filter).Decode(mp); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil, nil
		}

		return nil, err
	}

	p = atlant.NewProduct(mp.Name, mp.Price, time.Unix(0, mp.CreatedAt))

	for _, c := range mp.ChangeHistory {
		p.ApplyChange(&atlant.ProductChanging{
			OldPrice:  c.OldPrice,
			NewPrice:  c.NewPrice,
			RequestID: c.RequestID,
			ChangeID:  c.ChangeID,
			CreatedAt: time.Unix(0, c.CreatedAt),
		})
	}

	return p, nil
}

//
func (ps *ProductStorage) List(
	ctx context.Context,
	start atlant.StartParameter,
	limit atlant.LimitParameter,
	opts atlant.ProductSortingOptions,
) (
	pp []atlant.Product,
	err error,
) {
	rows, err := ps.products.Find(ctx, bson.D{}, initListOptions(start, limit, opts))
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, logger *zap.Logger) {
		if closeErr := rows.Close(ctx); closeErr != nil {
			logger.Error("close cursor error", zap.Error(err))
			err = closeErr
		}
	}(ctx, ps.logger)

	pp = make([]atlant.Product, 0, limit.Int64())

	var (
		p  *atlant.Product
		mp Product
	)

	for rows.Next(ctx) {
		if err = rows.Decode(&mp); err != nil {
			return nil, err
		}

		p = atlant.NewProduct(mp.Name, mp.Price, time.Unix(0, mp.CreatedAt))

		for _, c := range mp.ChangeHistory {
			p.ApplyChange(&atlant.ProductChanging{
				OldPrice:  c.OldPrice,
				NewPrice:  c.NewPrice,
				RequestID: c.RequestID,
				ChangeID:  c.ChangeID,
				CreatedAt: time.Unix(0, c.CreatedAt),
			})
		}

		pp = append(pp, *p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pp, nil
}

func initListOptions(
	start atlant.StartParameter,
	limit atlant.LimitParameter,
	opts atlant.ProductSortingOptions,
) (
	mdbOpts *options.FindOptions,
) {
	mdbOpts = options.
		Find().
		SetSkip(start.Int64()).
		SetLimit(limit.Int64())

	var (
		sort = bson.D{}

		fromDomainFieldToMongoFieldMap = map[atlant.SortingField]string{
			"name":         "name",
			"price":        "price",
			"created_at":   "created_at",
			"updated_at":   "updated_at",
			"update_count": "update_count",
		}

		fromDomainSortingDirectionToMongoSortingDirection = map[atlant.SortingDirection]int{
			atlant.SortingDirectionAsc:  1,
			atlant.SortingDirectionDesc: -1,
		}
	)

	for _, opt := range opts {
		sort = append(sort, bson.E{
			Key:   fromDomainFieldToMongoFieldMap[opt.Field],
			Value: fromDomainSortingDirectionToMongoSortingDirection[opt.Direction],
		})
	}

	return mdbOpts.SetSort(sort)
}
