package integration_test

import (
	"fmt"
	"laughing-goggles/account"
	"net/http"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestGetAccountByID_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	svc := account.NewAccountService()

	_, err := svc.CreateAccount(account.CreateAccountParams{AccountID: 123, InitialBalance: decimal.RequireFromString("100.23344")})
	require.NoError(t, err) // sanity check

	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/accounts/123", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetAccountByID_404NotFound(t *testing.T) {
	t.Parallel()

	// Arrange
	svc := account.NewAccountService()
	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/accounts/999", srv.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
