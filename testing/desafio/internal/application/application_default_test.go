package application

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplicationDefault_NewApplicationDefault(t *testing.T) {
	t.Run("Create a new ApplicationDefault without parameters", func(t *testing.T) {
		application := NewApplicationDefault(nil)
		require.NotNil(t, application)
	})

	t.Run("Create a new ApplicationDefault with parameters", func(t *testing.T) {
		application := NewApplicationDefault(&ConfigApplicationDefault{
			Addr: ":8080",
		})
		require.NotNil(t, application)
	})
}

func TestApplicationDefault_SetUp(t *testing.T) {
	t.Run("Set up the application", func(t *testing.T) {
		application := NewApplicationDefault(nil)
		err := application.SetUp()
		require.NoError(t, err)
	})
}

func TestApplicationDefault_TearDown(t *testing.T) {
	t.Run("Tear down the application", func(t *testing.T) {
		application := NewApplicationDefault(nil)
		err := application.TearDown()
		require.NoError(t, err)
	})
}
