package v1

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/morozovcookie/atlant"
	"go.uber.org/zap"
)

func TestProductStorer_Store(t *testing.T) {
	tt := []struct {
		name string

		producer *MockProducer

		initTransactionsInput  []interface{}
		initTransactionsOutput []interface{}

		beginTransactionInput  []interface{}
		beginTransactionOutput []interface{}

		abortTransactionInput  []interface{}
		abortTransactionOutput []interface{}

		produceParams []struct {
			input  []interface{}
			output []interface{}
		}

		commitTransactionInput  []interface{}
		commitTransactionOutput []interface{}

		pp []atlant.Product

		expected error

		wantErr bool
	}{
		{
			name: "pass",

			producer: &MockProducer{},

			initTransactionsInput: []interface{}{
				context.Background(),
			},
			initTransactionsOutput: []interface{}{
				error(nil),
			},

			beginTransactionInput: []interface{}{
				context.Background(),
			},
			beginTransactionOutput: []interface{}{
				error(nil),
			},

			abortTransactionInput:  nil,
			abortTransactionOutput: nil,

			produceParams: []struct {
				input  []interface{}
				output []interface{}
			}{
				{
					input: []interface{}{
						context.Background(),
						bytes.NewBuffer(append([]byte(`{"name":"sample","price":1.01,"created_at":10000000000}`), '\n')),
					},
					output: []interface{}{
						error(nil),
					},
				},
			},

			commitTransactionInput: []interface{}{
				context.Background(),
			},
			commitTransactionOutput: []interface{}{
				error(nil),
			},

			pp: []atlant.Product{
				{
					Name:      "sample",
					Price:     1.01,
					CreatedAt: time.Unix(10, 0),
				},
			},
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			test.producer.On("InitTransactions", test.initTransactionsInput...).
				Return(test.initTransactionsOutput...)
			test.producer.On("BeginTransaction", test.beginTransactionInput...).
				Return(test.beginTransactionOutput...)
			test.producer.On("AbortTransaction", test.abortTransactionInput...).
				Return(test.abortTransactionOutput...)

			for _, produce := range test.produceParams {
				test.producer.On("Produce", produce.input...).
					Return(produce.output...)
			}

			test.producer.On("CommitTransaction", test.commitTransactionInput...).
				Return(test.commitTransactionOutput...)

			storer := NewProductStorer(test.producer, zap.NewNop())
			actual := storer.Store(context.Background(), test.pp...)
			if (actual != nil) != test.wantErr {
				t.Error(actual)
				t.Fail()
			}
		})
	}
}
