package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	client "github.com/TutorialEdge/go-http-client-course"
	"github.com/stretchr/testify/assert"
)

func TestGetPokemonByName(t *testing.T) {
	t.Run("can get and parse a pokemon by its name", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, `{"name":"pikachu"}`)
		}))
		defer ts.Close()

		myClient := client.NewClient()
		resp, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, "pikachu", resp.Name)

	})

}
