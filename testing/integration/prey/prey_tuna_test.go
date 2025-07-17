package prey_test

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTuna_GetSpeed(t *testing.T) {
	t.Run("case 1: positive speed", func(t *testing.T) {
		// Arrange
		speed := 100.0
		position := &positioner.Position{X: 0, Y: 0, Z: 0}
		tuna := prey.NewTuna(speed, position)

		// Act
		result := tuna.GetSpeed()

		// Assert
		require.Equal(t, speed, result)
	})
}

func TestTuna_GetPosition(t *testing.T) {
	t.Run("case 1: positive position", func(t *testing.T) {
		// Arrange
		speed := 100.0
		position := &positioner.Position{X: 0, Y: 0, Z: 0}
		tuna := prey.NewTuna(speed, position)

		// Act
		result := tuna.GetPosition()

		// Assert
		require.Equal(t, position, result)
	})
}
