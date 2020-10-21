package http

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/morozovcookie/atlant"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestProductFetcher_Fetch(t *testing.T) {
	tt := []struct {
		name string

		client    *MockClient
		getInput  []interface{}
		getOutput []interface{}

		url  string
		mark time.Time

		expected []atlant.Product

		wantErr bool
	}{
		{
			name: "pass",

			client: &MockClient{},
			getInput: []interface{}{
				context.Background(),
				"http://127.0.0.1:8081/sample.csv",
			},
			getOutput: []interface{}{
				&http.Response{
					Status:        "OK",
					StatusCode:    200,
					Proto:         "HTTP/1.1",
					ProtoMajor:    1,
					ProtoMinor:    1,
					Body:          ioutil.NopCloser(bytes.NewBufferString(`sample;1.01`)),
					ContentLength: int64(len(`sample;1.01`)),
				},
				error(nil),
			},

			url:  "http://127.0.0.1:8081/sample.csv",
			mark: time.Unix(10, 0),

			expected: []atlant.Product{
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
			test.client.On("Get", test.getInput...).
				Return(test.getOutput...)

			u, err := url.Parse("http://127.0.0.1:8081/sample.csv")
			if err != nil {
				t.Fatal(err)
			}

			fetcher := NewProductFetcher(test.client, zap.NewNop())
			actual, err := fetcher.Fetch(context.Background(), u, test.mark)
			if (err != nil) != test.wantErr {
				t.Error(err)
				t.FailNow()
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}
