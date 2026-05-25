package integration_test

import (
	"laughing-goggles/account"
	"laughing-goggles/httpapi"
	"laughing-goggles/testutil"
	"net/http/httptest"
	"testing"
)

func newTestServer(t *testing.T, svc *account.AccountService) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(
		httpapi.NewHandler(
			testutil.DiscardLogger,
			svc,
		),
	)
	t.Cleanup(srv.Close)
	return srv
}
