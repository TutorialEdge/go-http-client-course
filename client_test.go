package client_test

import (
	"context"
	"testing"

	pokeClient "github.com/TutorialEdge/go-http-client-course"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}

var _ = Describe("Client", func() {
	Describe("Fetching pokemon with a client", func() {
		Context("with a regular client with no options", func() {

			client := pokeClient.NewClient()
			var pokemon pokeClient.Pokemon
			var err error
			pokemon, err = client.GetPokemonByName(context.Background(), "pikachu")

			It("should not have errored", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should have the correct id", func() {
				Expect(pokemon.ID).To(Equal(25))
			})
		})
	})
})
