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

func TestCreateAccount_201Created(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/accounts", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateAccountJSONRequestBody{AccountId: 123, InitialBalance: "100.23344"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestCreateAccount_400BadRequest_NegativeBalance(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/accounts", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateAccountJSONRequestBody{AccountId: 123, InitialBalance: "-50.00"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateAccount_400BadRequest_ZeroBalance(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/accounts", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateAccountJSONRequestBody{AccountId: 123, InitialBalance: "0"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateAccount_400BadRequest_InvalidBalance(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)
	srv := newTestServer(t, svc)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/accounts", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateAccountJSONRequestBody{AccountId: 123, InitialBalance: "not-a-number"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateAccount_409Conflict_DuplicateAccountID(t *testing.T) {
	t.Parallel()

	// Arrange
	pool := testutil.NewTestPgxPool(t)
	svc := account.NewAccountService(pool)

	_, err := svc.CreateAccount(t.Context(), account.CreateAccountParams{AccountID: 123, InitialBalance: decimal.RequireFromString("100.23344")})
	require.NoError(t, err) // sanity check

	srv := newTestServer(t, svc)

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
	assert.Equal(t, http.StatusConflict, resp.StatusCode)
}
