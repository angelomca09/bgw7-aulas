package hunter_test

import (
	"testdoubles/hunter"
	"testdoubles/positioner"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWhiteShark_Hunt(t *testing.T) {
	t.Run("case 1: hunter catches the prey", func(t *testing.T) {
		// Arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() *positioner.Position {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() float64 {
			return 50
		}

		sm := simulator.NewSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) bool {
			return true
		}

		impl := hunter.NewWhiteShark(100, &positioner.Position{X: 0, Y: 0, Z: 0}, sm)

		// Act
		err := impl.Hunt(pr)

		// Assert
		expErr := error(nil)
		expMockCalls := 1
		require.ErrorIs(t, err, expErr)
		require.Equal(t, expMockCalls, sm.Calls.CanCatch)
	})

	t.Run("case 2: hunter is slower than prey", func(t *testing.T) {
		// Arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() *positioner.Position {
			return &positioner.Position{X: 0, Y: 0, Z: 0}
		}
		pr.GetSpeedFunc = func() float64 {
			return 50
		}

		sm := simulator.NewSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) bool {
			return false
		}

		impl := hunter.NewWhiteShark(100, &positioner.Position{X: 0, Y: 0, Z: 0}, sm)

		// Act
		err := impl.Hunt(pr)

		// Assert
		expErr := hunter.ErrCanNotHunt
		expMockCalls := 1
		require.ErrorIs(t, err, expErr)
		require.Equal(t, expMockCalls, sm.Calls.CanCatch)
	})

	t.Run("case 3: hunter is to distant to the prey", func(t *testing.T) {
		// Arrange
		pr := prey.NewPreyStub()
		pr.GetPositionFunc = func() *positioner.Position {
			return &positioner.Position{X: 500, Y: 500, Z: 500}
		}
		pr.GetSpeedFunc = func() float64 {
			return 50
		}

		sm := simulator.NewSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) bool {
			return false
		}

		impl := hunter.NewWhiteShark(100, &positioner.Position{X: 0, Y: 0, Z: 0}, sm)

		// Act
		err := impl.Hunt(pr)

		// Assert
		expErr := hunter.ErrCanNotHunt
		expMockCalls := 1
		require.ErrorIs(t, err, expErr)
		require.Equal(t, expMockCalls, sm.Calls.CanCatch)
	})
}

func TestWhiteShark_CreateWhiteShark(t *testing.T) {
	t.Run("case 1: create white shark with random func", func(t *testing.T) {
		// Arrange
		sm := simulator.NewSimulatorMock()
		sm.CanCatchFunc = func(hunter, prey *simulator.Subject) bool {
			return true
		}

		// Act
		shark := hunter.CreateWhiteShark(sm)

		// Assert
		require.NotNil(t, shark)
	})
}
