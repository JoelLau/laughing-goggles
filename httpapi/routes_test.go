package httpapi_test

import (
	"fmt"
	"laughing-goggles/account"
	"laughing-goggles/api"
	"laughing-goggles/httpapi"
	"laughing-goggles/testutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestLivez_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	srv := httptest.NewServer(httpapi.NewHandler(testutil.DiscardLogger, nil))
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
	srv := httptest.NewServer(httpapi.NewHandler(testutil.DiscardLogger, nil))
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
	store := account.NewAccountRepository()
	srv := httptest.NewServer(httpapi.NewHandler(testutil.DiscardLogger, store))
	defer srv.Close()

	// Act
	reqBody := api.CreateAccountJSONRequestBody{AccountId: 123, InitialBalance: "100.23344"}
	resp, err := http.Post(
		fmt.Sprintf("%s/accounts", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, reqBody),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	got, err := store.GetAccountByID(123)
	require.NoError(t, err)

	want := account.Account{ID: 123, Balance: decimal.RequireFromString("100.23344")}
	require.Equal(t, want, got)
}
