package controllers_test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/numary/ledger/pkg/api"
	"github.com/numary/ledger/pkg/health"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthController(t *testing.T) {

	type testCase struct {
		name                 string
		healthChecksProvider []interface{}
		expectedStatus       int
		expectedResult       map[string]string
	}

	var tests = []testCase{
		{
			name: "all-ok",
			healthChecksProvider: []interface{}{
				func() health.NamedCheck {
					return health.NewNamedCheck("test1", health.CheckFn(func(ctx context.Context) error {
						return nil
					}))
				},
				func() health.NamedCheck {
					return health.NewNamedCheck("test2", health.CheckFn(func(ctx context.Context) error {
						return nil
					}))
				},
			},
			expectedStatus: http.StatusOK,
			expectedResult: map[string]string{
				"test1": "OK",
				"test2": "OK",
			},
		},
		{
			name: "one-failing",
			healthChecksProvider: []interface{}{
				func() health.NamedCheck {
					return health.NewNamedCheck("test1", health.CheckFn(func(ctx context.Context) error {
						return nil
					}))
				},
				func() health.NamedCheck {
					return health.NewNamedCheck("test2", health.CheckFn(func(ctx context.Context) error {
						return errors.New("failure")
					}))
				},
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResult: map[string]string{
				"test1": "OK",
				"test2": "failure",
			},
		},
	}

	for _, tc := range tests {

		options := make([]fx.Option, 0)
		for _, p := range tc.healthChecksProvider {
			options = append(options, health.ProvideHealthCheck(p))
		}
		options = append(options, fx.Invoke(func(lc fx.Lifecycle, h *api.API) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					rec := httptest.NewRecorder()
					req := httptest.NewRequest("GET", "/_health", nil)

					h.ServeHTTP(rec, req)

					if !assert.Equal(t, tc.expectedStatus, rec.Result().StatusCode) {
						return nil
					}
					ret := make(map[string]string)
					err := json.NewDecoder(rec.Result().Body).Decode(&ret)
					if !assert.NoError(t, err) {
						return nil
					}
					if !assert.Equal(t, tc.expectedResult, ret) {
						return nil
					}
					return nil
				},
			})
		}))
	}

}
