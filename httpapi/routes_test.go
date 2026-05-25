package httpapi_test

import (
	"fmt"
	"laughing-goggles/account"
	"laughing-goggles/httpapi"
	"laughing-goggles/testutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(httpapi.NewHandler(testutil.DiscardLogger, account.NewAccountService()))
	t.Cleanup(srv.Close)
	return srv
}

func TestLivez_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	srv := newTestServer(t)

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/livez", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestReadyz_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	srv := newTestServer(t)

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/readyz", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
