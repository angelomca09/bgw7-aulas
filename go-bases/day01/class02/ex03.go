package class02

import "fmt"

func Exercicio3() {
	fmt.Println("\nExercicio 3\n-----------------")

	meses := map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}

	var mes int
	fmt.Println("Digite o número do mês (1-12): ")
	fmt.Scanln(&mes)

	if mes >= 1 && mes <= 12 {
		fmt.Printf("O mês %d é %s.\n", mes, meses[mes])
	} else {
		fmt.Println("Número do mês inválido. Por favor, digite um número entre 1 e 12.")
	}

}
