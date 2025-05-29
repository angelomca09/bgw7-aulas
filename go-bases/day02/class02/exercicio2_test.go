package class02_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"bgw7-aulas/go-bases/day02/class01"
)

func TestExercicio2_CalculaMedia(t *testing.T) {
	// Arrange
	notas := []int{10, 20, 30, 40, 50}
	mediaEsperada := 30.0

	// Act
	media := class01.CalculaMedia(notas)

	// Assert
	require.Equal(t, mediaEsperada, media, "A média esperada é %.2f, mas a média obtida é %.2f", mediaEsperada, media)
}
