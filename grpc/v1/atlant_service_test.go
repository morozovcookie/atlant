package v1

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/morozovcookie/atlant"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestAtlantHandler_Fetch(t *testing.T) {
	tt := []struct {
		name string

		config *MockAtlantServiceConfig

		productFetcher *atlant.MockProductFetcher
		fetchInput     []interface{}
		fetchOutput    []interface{}

		productStorer *atlant.MockProductStorer
		storeInput    []interface{}
		storeOutput   []interface{}

		clock     *MockClock
		nowOutput []interface{}

		req *FetchRequest
		url string

		expected *empty.Empty

		wantErr bool
	}{
		{
			name: "pass",

			config: &MockAtlantServiceConfig{},

			productFetcher: &atlant.MockProductFetcher{},
			fetchInput: []interface{}{
				context.Background(),
				time.Unix(10, 0),
			},
			fetchOutput: []interface{}{
				[]atlant.Product{
					{
						Name:      "sample",
						Price:     1.01,
						CreatedAt: time.Unix(10, 0),
					},
				},
				error(nil),
			},

			productStorer: &atlant.MockProductStorer{},
			storeInput: []interface{}{
				context.Background(),
				[]atlant.Product{
					{
						Name:      "sample",
						Price:     1.01,
						CreatedAt: time.Unix(10, 0),
					},
				},
			},
			storeOutput: []interface{}{
				error(nil),
			},

			clock: &MockClock{},
			nowOutput: []interface{}{
				time.Unix(10, 0),
			},

			req: &FetchRequest{
				URL: "http://127.0.0.1:8080/sample.csv",
			},
			url: "http://127.0.0.1:8080/sample.csv",

			expected: &empty.Empty{},
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			test.clock.On("NowInUTC").
				Return(test.nowOutput...)

			u, err := url.Parse(test.url)
			if err != nil {
				t.Fatal(err)
			}

			test.productFetcher.On("Fetch",
				append([]interface{}{}, test.fetchInput[0], u, test.fetchInput[1])...).
				Return(test.fetchOutput...)

			test.productStorer.On("Store", test.storeInput...).
				Return(test.storeOutput...)

			test.config.On("ProductFetcherInstance").
				Return([]interface{}{test.productFetcher}...)
			test.config.On("ProductStorerInstance").
				Return([]interface{}{test.productStorer}...)
			test.config.On("ClockInstance").
				Return([]interface{}{test.clock}...)

			handler := NewAtlantService(test.config, zap.NewNop())
			actual, err := handler.Fetch(context.Background(), test.req)
			if (err != nil) != test.wantErr {
				t.Error(err)
				t.FailNow()
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}
