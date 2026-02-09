// Package bulk provides support for bulk operations.
package bulk

import (
	"errors"
	"fmt"
)

// MaxBatchSize defines the maximum number of items allowed in a bulk operation.
const MaxBatchSize = 100

// ErrBatchSizeExceeded is returned when a bulk operation exceeds the maximum batch size.
var ErrBatchSizeExceeded = errors.New("batch size exceeds maximum allowed")

// ValidateBatchSize checks if the provided count is within the allowed batch size limit.
// Returns an error if the count exceeds MaxBatchSize or is zero.
func ValidateBatchSize(count int) error {
	if count == 0 {
		return errors.New("batch cannot be empty")
	}
	if count > MaxBatchSize {
		return fmt.Errorf("%w: got %d, max %d", ErrBatchSizeExceeded, count, MaxBatchSize)
	}
	return nil
}
