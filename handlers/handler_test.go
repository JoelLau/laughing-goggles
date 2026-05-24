package handlers_test

import (
	"fmt"
	"laughing-goggles/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLivez_200OK(t *testing.T) {
	t.Parallel()

	// Arrange
	server := httptest.NewServer(handlers.NewHandler())
	defer server.Close()

	// Act
	resp, err := http.Get(fmt.Sprintf("%s/api/livez", server.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
