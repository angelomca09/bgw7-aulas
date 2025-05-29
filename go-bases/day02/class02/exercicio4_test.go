package class02_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"bgw7-aulas/go-bases/day02/class01"
)

func TestExercicio4_CalculaMinimo(t *testing.T) {
	// Arrange
	notas := []int{2, 3, 3, 4, 10, 2, 4, 5}
	minimoEsperado := 2.0

	// Act
	minFunc, err := class01.Operation(class01.Minimum)
	require.NoError(t, err, "Erro ao obter função de mínimo")

	// Assert
	minimo := minFunc(notas...)
	require.Equal(t, minimoEsperado, minimo, "O valor mínimo esperado é %.2f, mas o obtido foi %.2f", minimoEsperado, minimo)
}

func TestExercicio4_CalculaMedio(t *testing.T) {
	// Arrange
	notas := []int{2, 3, 3, 4, 1, 2, 4, 5}
	medioEsperado := 3.0

	// Act
	averageFunc, err := class01.Operation(class01.Average)
	require.NoError(t, err, "Erro ao obter função de médio")

	// Assert
	medio := averageFunc(notas...)
	require.Equal(t, medioEsperado, medio, "O valor médio esperado é %.2f, mas o obtido foi %.2f", medioEsperado, medio)
}

func TestExercicio4_CalculaMaximo(t *testing.T) {
	// Arrange
	notas := []int{2, 3, 3, 4, 10, 2, 4, 5}
	maximoEsperado := 10.0

	// Act
	maxFunc, err := class01.Operation(class01.Maximum)
	require.NoError(t, err, "Erro ao obter função de máximo")

	// Assert
	máximo := maxFunc(notas...)
	require.Equal(t, maximoEsperado, máximo, "O valor máximo esperado é %.2f, mas o obtido é %.2f", maximoEsperado, máximo)
}
