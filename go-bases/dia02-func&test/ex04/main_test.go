/* Exercício 4 - Testando o cálculo de estatísticas

Os professores da universidade colombiana participaram do programa de análise de dados do Google, que premia os
melhores estatísticos da região. Portanto, os professores nos pediram para verificar o funcionamento correto
de nossos cálculos estatísticos. A tarefa a seguir é solicitada:

- Realize um teste para calcular o número mínimo de notas.
- Realize um teste para calcular o número máximo de notas.
- Realize um teste para calcular o número médio de notas.
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExercicio4_CalculaMinimo(t *testing.T) {
	// Arrange
	notas := []int{2, 3, 3, 4, 10, 2, 4, 5}
	minimoEsperado := 2.0

	// Act
	minFunc, err := operation(minimum)
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
	averageFunc, err := operation(average)
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
	maxFunc, err := operation(maximum)
	require.NoError(t, err, "Erro ao obter função de máximo")

	// Assert
	máximo := maxFunc(notas...)
	require.Equal(t, maximoEsperado, máximo, "O valor máximo esperado é %.2f, mas o obtido é %.2f", maximoEsperado, máximo)
}
