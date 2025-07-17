package hunt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		//given
		shark := NewWhiteShark(true, false, 100)
		tuna := NewTuna("Tuna", 50)

		//when
		err := shark.Hunt(tuna)

		//then

		require.NoError(t, err)
		require.False(t, shark.Hungry)
		require.True(t, shark.Tired)

	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		//given
		shark := NewWhiteShark(false, false, 100)
		tuna := NewTuna("Tuna", 50)

		// when
		err := shark.Hunt(tuna)

		//then
		require.Error(t, err)
		require.Equal(t, ErrSharkIsNotHungry, err)
		require.False(t, shark.Hungry)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		//given
		shark := NewWhiteShark(true, true, 100)
		tuna := NewTuna("Tuna", 50)

		// when
		err := shark.Hunt(tuna)

		//then
		require.Error(t, err)
		require.Equal(t, ErrSharkIsTired, err)
		require.True(t, shark.Tired)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		//given
		shark := NewWhiteShark(true, false, 49)
		tuna := NewTuna("Tuna", 50)

		// when
		err := shark.Hunt(tuna)

		//then
		require.Error(t, err)
		require.Equal(t, ErrSharkIsSlower, err)
		require.Less(t, shark.Speed, tuna.Speed)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		//given
		shark := NewWhiteShark(false, false, 100)

		// when
		err := shark.Hunt(nil)

		//then
		require.Error(t, err)
		require.Equal(t, ErrTunaIsNil, err)
	})
}
