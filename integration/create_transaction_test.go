package integration_test

import (
	"fmt"
	"laughing-goggles/account"
	"laughing-goggles/gen/api"
	"laughing-goggles/testutil"
	"net/http"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTransaction_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	_, err := svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 123, InitialBalance: decimal.RequireFromString("123")})
	require.NoError(t, err)
	_, err = svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 456, InitialBalance: decimal.RequireFromString("456")})
	require.NoError(t, err)

	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/transactions", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateTransactionJSONRequestBody{SourceAccountId: 123, DestinationAccountId: 456, Amount: "0.05"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	src, err := svc.GetAccountByID(t.Context(), 123)
	require.NoError(t, err)
	assert.True(t, src.Balance.Equal(decimal.RequireFromString("122.95")))

	dst, err := svc.GetAccountByID(t.Context(), 456)
	require.NoError(t, err)
	assert.True(t, dst.Balance.Equal(decimal.RequireFromString("456.05")))
}

func TestCreateTransaction_400BadRequest_NegativeAmount(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	_, err := svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 123, InitialBalance: decimal.RequireFromString("100")})
	require.NoError(t, err)
	_, err = svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 456, InitialBalance: decimal.RequireFromString("100")})
	require.NoError(t, err)

	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/transactions", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateTransactionJSONRequestBody{SourceAccountId: 123, DestinationAccountId: 456, Amount: "-50.00"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateTransaction_400BadRequest_ZeroAmount(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	_, err := svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 123, InitialBalance: decimal.RequireFromString("100")})
	require.NoError(t, err)
	_, err = svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 456, InitialBalance: decimal.RequireFromString("100")})
	require.NoError(t, err)

	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/transactions", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateTransactionJSONRequestBody{SourceAccountId: 123, DestinationAccountId: 456, Amount: "0"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateTransaction_400BadRequest_InvalidAmount(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	_, err := svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 123, InitialBalance: decimal.RequireFromString("100")})
	require.NoError(t, err)
	_, err = svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 456, InitialBalance: decimal.RequireFromString("100")})
	require.NoError(t, err)

	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/transactions", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateTransactionJSONRequestBody{SourceAccountId: 123, DestinationAccountId: 456, Amount: "not-a-number"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateTransaction_404NotFound_NonExistentSource(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	_, err := svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 123, InitialBalance: decimal.RequireFromString("100")})
	require.NoError(t, err)
	_, err = svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 456, InitialBalance: decimal.RequireFromString("100")})
	require.NoError(t, err)

	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/transactions", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateTransactionJSONRequestBody{SourceAccountId: 999, DestinationAccountId: 456, Amount: "50.00"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCreateTransaction_404NotFound_NonExistentDestination(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	_, err := svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 123, InitialBalance: decimal.RequireFromString("100")})
	require.NoError(t, err)
	_, err = svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 456, InitialBalance: decimal.RequireFromString("100")})
	require.NoError(t, err)

	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/transactions", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateTransactionJSONRequestBody{SourceAccountId: 123, DestinationAccountId: 999, Amount: "50.00"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCreateTransaction_400BadRequest_InsufficientBalance(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	_, err := svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 123, InitialBalance: decimal.RequireFromString("10")})
	require.NoError(t, err)
	_, err = svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 456, InitialBalance: decimal.RequireFromString("100")})
	require.NoError(t, err)

	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/transactions", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateTransactionJSONRequestBody{SourceAccountId: 123, DestinationAccountId: 456, Amount: "50.00"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
