package class02_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"bgw7-aulas/go-bases/day02/class01"
)

func TestExercicio1_SalarioMenos50K(t *testing.T) {
	//arrange
	salario := 50000
	impostoEsperado := 0.0
	porcentagemEsperada := 0.0

	//act
	imposto, porcentagem := class01.CalculaImposto(float64(salario))

	//assert
	require.Equal(t, impostoEsperado, imposto, "Imposto esperado %.2f Imposto obtido %.2f", impostoEsperado, imposto)
	require.Equal(t, porcentagemEsperada, porcentagem, "Porcentagem esperada %.2f Porcentagem obtida %.2f", porcentagemEsperada, porcentagem)
}

func TestExercicio1_SalarioMais50K(t *testing.T) {
	//arrange
	salario := 100000
	impostoEsperado := 17000.0
	porcentagemEsperada := 0.17

	//act
	imposto, porcentagem := class01.CalculaImposto(float64(salario))

	//assert
	require.Equal(t, impostoEsperado, imposto, "Imposto esperado %.2f Imposto obtido %.2f", impostoEsperado, imposto)
	require.Equal(t, porcentagemEsperada, porcentagem, "Porcentagem esperada %.2f Porcentagem obtida %.2f", porcentagemEsperada, porcentagem)
}

func TestExercicio1_SalarioMais150K(t *testing.T) {
	//arrange
	salario := 200000
	impostoEsperado := 54000.0
	porcentagemEsperada := 0.27

	//act
	imposto, porcentagem := class01.CalculaImposto(float64(salario))

	//assert
	require.Equal(t, impostoEsperado, imposto, "Imposto esperado %.2f Imposto obtido %.2f", impostoEsperado, imposto)
	require.Equal(t, porcentagemEsperada, porcentagem, "Porcentagem esperada %.2f Porcentagem obtida %.2f", porcentagemEsperada, porcentagem)
}
