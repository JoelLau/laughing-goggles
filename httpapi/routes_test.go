package httpapi_test

import (
	"fmt"
	"laughing-goggles/account"
	"laughing-goggles/httpapi"
	"laughing-goggles/testutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, svc httpapi.AccountsService) *httptest.Server {
	t.Helper()

	srv := httptest.NewServer(httpapi.NewHandler(testutil.DiscardLogger, svc))
	t.Cleanup(srv.Close)

	return srv
}

func TestLivez_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/livez", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestReadyz_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/readyz", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestReadyz_500InternalServerError(t *testing.T) {
	t.Parallel()

	// Arrange
	svc := account.NewAccountService(nil)
	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/readyz", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
