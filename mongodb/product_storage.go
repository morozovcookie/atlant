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
func (ps *ProductStorage) Store(ctx context.Context, pp ...atlant.Product) (err error) {
	if len(pp) != 1 {
		return ErrTooMuchObjectsForStore
	}

	var (
		p      = pp[0]
		filter = bson.D{{Key: "product_id", Value: p.ID()}}
		mp     = &Product{
			CreatedAt: p.CreatedAt.UnixNano(),
			Price:     p.Price,
			ID:        p.ID(),
			Name:      p.Name,
		}
	)

	if p.UpdateCount != 0 {
		mp.UpdateCount = new(int)
		*(mp.UpdateCount) = p.UpdateCount
	}

	if p.UpdatedAt.UnixNano() != 0 {
		mp.UpdatedAt = new(int64)
		*(mp.UpdatedAt) = p.UpdatedAt.UnixNano()
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

	if err = ps.products.FindOne(ctx, filter).Decode(&mp); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil, nil
		}

		return nil, err
	}

	p = &atlant.Product{
		Price:     mp.Price,
		Name:      mp.Name,
		CreatedAt: time.Unix(0, mp.CreatedAt),
	}

	if mp.UpdateCount != nil {
		p.UpdateCount = *(mp.UpdateCount)
	}

	if mp.UpdatedAt != nil {
		p.UpdatedAt = time.Unix(0, *(mp.UpdatedAt))
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
	var (
		mdbOpts = options.
			Find().
			SetSkip(start.Int64()).
			SetLimit(limit.Int64())
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

	rows, err := ps.products.Find(ctx, bson.D{}, mdbOpts.SetSort(sort))
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
		i = 0

		p  atlant.Product
		mp Product
	)

	for rows.Next(ctx) {
		if err = rows.Decode(&mp); err != nil {
			return nil, err
		}

		p = atlant.Product{
			Price:     mp.Price,
			Name:      mp.Name,
			CreatedAt: time.Unix(0, mp.CreatedAt),
		}

		if mp.UpdateCount != nil {
			p.UpdateCount = *(mp.UpdateCount)
		}

		if mp.UpdatedAt != nil {
			p.UpdatedAt = time.Unix(0, *(mp.UpdatedAt))
		}

		pp = append(pp, p)

		i++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pp[:i], nil
}
