package resourcebus

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

// TestNewResources is a helper method for testing.
func TestNewResources(n int) []NewResource {
	newRes := make([]NewResource, n)

	idx := rand.Intn(1000)
	for i := 0; i < n; i++ {
		idx++

		nu := NewResource{
			Name:         Names.MustParse(fmt.Sprintf("Sogeimaic%d", idx)),
			GalaxyID:     uuid.New(),
			AddedUserID:  uuid.New(),
			ResourceType: "Kammris Iron",
			CR:           int16(rand.Intn(1000)),
			CD:           int16(rand.Intn(1000)),
			DR:           int16(rand.Intn(1000)),
			FL:           int16(rand.Intn(1000)),
			HR:           int16(rand.Intn(1000)),
			MA:           int16(rand.Intn(1000)),
			PE:           int16(rand.Intn(1000)),
			OQ:           int16(rand.Intn(1000)),
			SR:           int16(rand.Intn(1000)),
			UT:           int16(rand.Intn(1000)),
			ER:           int16(rand.Intn(1000)),
		}

		newRes[i] = nu
	}

	return newRes
}

// TestSeedResources is a helper method for testing.
func TestSeedResources(ctx context.Context, n int, api *Business) ([]Resource, error) {
	newRes := TestNewResources(n)

	resSlice := make([]Resource, len(newRes))
	for i, nu := range newRes {
		res, err := api.Create(ctx, nu)
		if err != nil {
			return nil, fmt.Errorf("seeding resource: idx: %d : %w", i, err)
		}

		resSlice[i] = res
	}

	return resSlice, nil
}
