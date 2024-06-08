package galaxybus

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

// TestNewGalaxies is a helper method for testing.
func TestNewGalaxies(n int) []NewGalaxy {
	newGals := make([]NewGalaxy, n)

	idx := rand.Intn(1000)
	for i := 0; i < n; i++ {
		idx++

		nu := NewGalaxy{
			Name:        Names.MustParse(fmt.Sprintf("Name%d", idx)),
			OwnerUserID: uuid.New(),
		}

		newGals[i] = nu
	}

	return newGals
}

// TestSeedGalaxies is a helper method for testing.
func TestSeedGalaxies(ctx context.Context, n int, api *Business) ([]Galaxy, error) {
	newGals := TestNewGalaxies(n)

	gals := make([]Galaxy, len(newGals))
	for i, nu := range newGals {
		gal, err := api.Create(ctx, nu)
		if err != nil {
			return nil, fmt.Errorf("seeding galaxy: idx: %d : %w", i, err)
		}

		gals[i] = gal
	}

	return gals, nil
}
