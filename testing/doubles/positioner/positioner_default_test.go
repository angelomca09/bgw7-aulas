package positioner_test

import (
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPositionerDefault_GetLinearDistance(t *testing.T) {
	t.Run("case 1: positive positions", func(t *testing.T) {
		// Arrange
		from := &positioner.Position{X: 1, Y: 2, Z: 3}
		to := &positioner.Position{X: 4, Y: 5, Z: 6}
		expected := 5.196
		positioner := positioner.NewPositionerDefault()

		// Act
		result := positioner.GetLinearDistance(from, to)

		// Assert
		require.InDelta(t, expected, result, 0.001)

	})

	t.Run("case 2: negative positions", func(t *testing.T) {
		// Arrange
		from := &positioner.Position{X: -1, Y: -2, Z: -3}
		to := &positioner.Position{X: -4, Y: -5, Z: -6}
		expected := 5.196
		positioner := positioner.NewPositionerDefault()

		// Act
		result := positioner.GetLinearDistance(from, to)

		// Assert
		require.InDelta(t, expected, result, 0.001)

	})

	t.Run("case 3: no decimal return", func(t *testing.T) {
		// Arrange
		from := &positioner.Position{X: 0, Y: 0, Z: 0}
		to := &positioner.Position{X: 4, Y: 3, Z: 0}
		expected := 5.0
		positioner := positioner.NewPositionerDefault()

		// Act
		result := positioner.GetLinearDistance(from, to)

		// Assert
		require.Equal(t, expected, result)

	})
}
