/* Exercício 5 - Calcular a quantidade de alimentos

O abrigo de animais enviou uma reclamação porque o cálculo total de alimentos a serem comprados não estava correto e
eles compraram menos alimentos do que precisavam. Para manter nosso cliente satisfeito, teremos de realizar os testes
necessários para verificar se tudo está funcionando corretamente:

- Verifique o cálculo da quantidade de alimento para cães.
- Verifique o cálculo da quantidade de alimento para gatos.
- Verifique o cálculo da quantidade de alimento para hamsters.
- Verifique o cálculo da quantidade de alimento para tarântulas.

Exemplo:

func TestDog(t *testing.T)
func TestCat(t *testing.T)
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExercicio5_VeificaAnimalCaes(t *testing.T) {
	// Arrange
	quantidadeAnimais := 3
	alimentoEsperado := 30000

	// Act
	animalDog, msg := animal(dog)
	require.Empty(t, msg, "Mensagem de erro inesperada: %s", msg)

	// Assert
	alimento := animalDog(quantidadeAnimais)
	require.Equal(t, alimentoEsperado, alimento, "Alimento esperado %.2f Alimento obtido %.2f", alimentoEsperado, alimento)
}
func TestExercicio5_VeificaAnimalGatos(t *testing.T) {
	// Arrange
	quantidadeAnimais := 3
	alimentoEsperado := 15000

	// Act
	animalCat, msg := animal(cat)
	require.Empty(t, msg, "Mensagem de erro inesperada: %s", msg)

	// Assert
	alimento := animalCat(quantidadeAnimais)
	require.Equal(t, alimentoEsperado, alimento, "Alimento esperado %.2f Alimento obtido %.2f", alimentoEsperado, alimento)
}
func TestExercicio5_VeificaAnimalHamsters(t *testing.T) {
	// Arrange
	quantidadeAnimais := 3
	alimentoEsperado := 450

	// Act
	animalHamster, msg := animal(hamster)
	require.Empty(t, msg, "Mensagem de erro inesperada: %s", msg)

	// Assert
	alimento := animalHamster(quantidadeAnimais)
	require.Equal(t, alimentoEsperado, alimento, "Alimento esperado %.2f Alimento obtido %.2f", alimentoEsperado, alimento)
}
func TestExercicio5_VeificaAnimalTarantulas(t *testing.T) {
	// Arrange
	quantidadeAnimais := 3
	alimentoEsperado := 750

	// Act
	animalTarantula, msg := animal(tarantula)
	require.Empty(t, msg, "Mensagem de erro inesperada: %s", msg)

	// Assert
	alimento := animalTarantula(quantidadeAnimais)
	require.Equal(t, alimentoEsperado, alimento, "Alimento esperado %.2f Alimento obtido %.2f", alimentoEsperado, alimento)
}
