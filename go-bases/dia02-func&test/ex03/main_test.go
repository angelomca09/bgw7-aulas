/* Exercício 3 - Teste de salário

A empresa marítima não concorda com os resultados obtidos nos cálculos salariais e, por isso, pede que realizemos
uma série de testes em nosso programa. Precisaremos realizar os seguintes testes em nosso código:

- Calcule o salário da categoria "A".
- Calcule o salário da categoria "B".
- Calcule o salário da categoria "C".
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExercicio3_CalculaSalarioA(t *testing.T) {
	// Arrange
	categoria := "A"
	minutos := 120
	salarioEsperado := 9000.00

	// Act
	salario := calculaSalario(categoria, minutos)

	// Assert
	require.Equal(t, salarioEsperado, salario, "O salário esperado é %.2f, mas o salário obtido é %.2f", salarioEsperado, salario)
}
func TestExercicio3_CalculaSalarioB(t *testing.T) {
	// Arrange
	categoria := "B"
	minutos := 120
	salarioEsperado := 3600.00

	// Act
	salario := calculaSalario(categoria, minutos)

	// Assert
	require.Equal(t, salarioEsperado, salario, "O salário esperado é %.2f, mas o salário obtido é %.2f", salarioEsperado, salario)
}
func TestExercicio3_CalculaSalarioC(t *testing.T) {
	// Arrange
	categoria := "C"
	minutos := 120
	salarioEsperado := 2000.00

	// Act
	salario := calculaSalario(categoria, minutos)

	// Assert
	require.Equal(t, salarioEsperado, salario, "O salário esperado é %.2f, mas o salário obtido é %.2f", salarioEsperado, salario)
}
