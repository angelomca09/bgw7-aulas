/*
	Exercício 2 - Calcular a média

A escola relatou que as operações para calcular a média não estão sendo realizadas corretamente, portanto, agora
somos obrigados a realizar os testes correspondentes:

- Calcular a média das notas dos alunos.
*/
package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExercicio2_CalculaMedia(t *testing.T) {
	// Arrange
	notas := []int{10, 20, 30, 40, 50}
	mediaEsperada := 30.0

	// Act
	media := calculaMedia(notas)

	// Assert
	require.Equal(t, mediaEsperada, media, "A média esperada é %.2f, mas a média obtida é %.2f", mediaEsperada, media)
}
