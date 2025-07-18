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

func TestTuna_CreateTuna(t *testing.T) {
	t.Run("case 1: a tuna is created", func(t *testing.T) {
		// Arrange
		// ...

		// Act
		tuna := prey.CreateTuna()

		// Assert
		require.NotNil(t, tuna)
	})
}

func TestTuna_GetSpeed_WithStub(t *testing.T) {
	t.Run("case 1: get implemented speed", func(t *testing.T) {
		// Arrange
		p := prey.NewPreyStub()
		p.GetSpeedFunc = func() (speed float64) {
			return 100
		}

		// Act
		result := p.GetSpeed()

		// Assert
		require.EqualValues(t, 100, result)
	})

	t.Run("case 2: get not implemented speed", func(t *testing.T) {
		// Arrange
		p := prey.NewPreyStub()

		// Act
		result := p.GetSpeed()

		// Assert
		require.EqualValues(t, 0, result)
	})

	t.Run("case 3: get implemented position", func(t *testing.T) {
		// Arrange
		p := prey.NewPreyStub()
		p.GetPositionFunc = func() (position *positioner.Position) {
			return &positioner.Position{X: 100, Y: 100, Z: 100}
		}

		// Act
		result := p.GetPosition()

		// Assert
		require.EqualValues(t, &positioner.Position{X: 100, Y: 100, Z: 100}, result)
	})

	t.Run("case 4: get not implemented position", func(t *testing.T) {
		// Arrange
		p := prey.NewPreyStub()

		// Act
		result := p.GetPosition()

		// Assert
		require.Nil(t, result)
	})
}
