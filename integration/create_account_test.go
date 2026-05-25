package integration_test

import (
	"fmt"
	"laughing-goggles/gen/api"
	"laughing-goggles/testutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount_201Created(t *testing.T) {
	t.Parallel()

	// Arrange
	srv := newTestServer(t)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/accounts", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateAccountJSONRequestBody{AccountId: 123, InitialBalance: "100.23344"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestCreateAccount_400BadRequest_NegativeBalance(t *testing.T) {
	t.Parallel()

	// Arrange
	srv := newTestServer(t)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/accounts", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateAccountJSONRequestBody{AccountId: 123, InitialBalance: "-50.00"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateAccount_400BadRequest_ZeroBalance(t *testing.T) {
	t.Parallel()

	// Arrange
	srv := newTestServer(t)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/accounts", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateAccountJSONRequestBody{AccountId: 123, InitialBalance: "0"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateAccount_400BadRequest_InvalidBalance(t *testing.T) {
	t.Parallel()

	// Arrange
	srv := newTestServer(t)

	// Act
	resp, err := http.Post(
		fmt.Sprintf("%s/accounts", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, api.CreateAccountJSONRequestBody{AccountId: 123, InitialBalance: "not-a-number"}),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateAccount_409Conflict_DuplicateAccountID(t *testing.T) {
	t.Parallel()

	// Arrange
	srv := newTestServer(t)
	reqBody := api.CreateAccountJSONRequestBody{AccountId: 123, InitialBalance: "100.23344"}

	first, err := http.Post(
		fmt.Sprintf("%s/accounts", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, reqBody),
	)
	require.NoError(t, err)
	first.Body.Close()
	require.Equal(t, http.StatusCreated, first.StatusCode)

	// Act
	second, err := http.Post(
		fmt.Sprintf("%s/accounts", srv.URL),
		testutil.ContentTypeJSON,
		testutil.MustJSON(t, reqBody),
	)
	require.NoError(t, err)
	defer second.Body.Close()

	// Assert
	require.Equal(t, http.StatusConflict, second.StatusCode)
}
