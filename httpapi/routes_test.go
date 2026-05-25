package httpapi_test

import (
	"fmt"
	"laughing-goggles/gen/server"
	"laughing-goggles/httpapi"
	"laughing-goggles/testutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLivez_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	srv := httptest.NewServer(httpapi.NewHandler(testutil.DiscardLogger))
	defer srv.Close()

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
	srv := httptest.NewServer(httpapi.NewHandler(testutil.DiscardLogger))
	defer srv.Close()

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/readyz", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateAccount_201Created(t *testing.T) {
	t.Parallel()

	// Arrange
	srv := httptest.NewServer(httpapi.NewHandler(testutil.DiscardLogger))
	defer srv.Close()

	// Act
	reqBody := server.CreateAccountJSONRequestBody{AccountId: 123, InitialBalance: "100.00"}
	resp, err := http.Post(fmt.Sprintf("%s/accounts", srv.URL), "application/json", testutil.MustJSON(t, reqBody))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusCreated, resp.StatusCode)
}
