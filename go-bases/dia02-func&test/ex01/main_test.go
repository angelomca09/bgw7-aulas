/*
	Exercício 1 - Teste de tributação salarial

A empresa de chocolates que anteriormente precisava calcular o imposto de seus funcionários ao depositar seus
salários agora nos pediu para validar se os cálculos desses impostos estão corretos.
Para isso, eles nos pediram que realizássemos os testes correspondentes para:

- Calcule o imposto caso o funcionário ganhe menos de US$ 50.000.
- Calcule o imposto caso o funcionário ganhe mais de US$ 50.000.
- Calcule o imposto caso o funcionário ganhe mais de US$ 150.000.
*/
package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExercicio1_SalarioMenos50K(t *testing.T) {
	//arrange
	salario := 50000
	impostoEsperado := 0.0
	porcentagemEsperada := 0.0

	//act
	imposto, porcentagem := calculaImposto(float64(salario))

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
	imposto, porcentagem := calculaImposto(float64(salario))

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
	imposto, porcentagem := calculaImposto(float64(salario))

	//assert
	require.Equal(t, impostoEsperado, imposto, "Imposto esperado %.2f Imposto obtido %.2f", impostoEsperado, imposto)
	require.Equal(t, porcentagemEsperada, porcentagem, "Porcentagem esperada %.2f Porcentagem obtida %.2f", porcentagemEsperada, porcentagem)
}
