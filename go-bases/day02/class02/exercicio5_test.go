package class02_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"bgw7-aulas/go-bases/day02/class01"
)

func TestExercicio5_VeificaAnimalCaes(t *testing.T) {
	// Arrange
	quantidadeAnimais := 3
	alimentoEsperado := 30000

	// Act
	animalDog, msg := class01.Animal(class01.Dog)
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
	animalCat, msg := class01.Animal(class01.Cat)
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
	animalHamster, msg := class01.Animal(class01.Hamster)
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
	animalTarantula, msg := class01.Animal(class01.Tarantula)
	require.Empty(t, msg, "Mensagem de erro inesperada: %s", msg)

	// Assert
	alimento := animalTarantula(quantidadeAnimais)
	require.Equal(t, alimentoEsperado, alimento, "Alimento esperado %.2f Alimento obtido %.2f", alimentoEsperado, alimento)
}
