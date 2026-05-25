package handlers_test

import (
	"fmt"
	"laughing-goggles/handlers"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLivez_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	server := httptest.NewServer(handlers.NewHandler(slog.New(slog.DiscardHandler)))
	defer server.Close()

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/livez", server.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestReadyz_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	server := httptest.NewServer(handlers.NewHandler(slog.New(slog.DiscardHandler)))
	defer server.Close()

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/readyz", server.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateAccount_201Created(t *testing.T) {
	t.Parallel()

	// Arrange
	server := httptest.NewServer(handlers.NewHandler(slog.New(slog.DiscardHandler)))
	defer server.Close()

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/accounts", server.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusCreated, resp.StatusCode)
}
