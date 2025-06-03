/* Exercício 2 - Clima

Uma empresa de meteorologia deseja um aplicativo em que possa ter a temperatura, a umidade e a pressão atmosférica de
 diferentes locais.

- Declare 3 variáveis especificando o tipo de dados; como valor, elas devem ter a temperatura, a umidade e a pressão
	do local onde você está.
- Imprima os valores das variáveis no console.
- Que tipo de dados você atribuiria às variáveis?
*/

package main

import "fmt"

func main() {
	fmt.Println("Exercício 2\n-----------------")

	var temperatura float64 = 36.5
	var umidade float64 = 75.0
	var pressao float64 = 1013.25

	fmt.Println("Temperatura:", temperatura, "°C")
	fmt.Println("Umidade:", umidade, "%")
	fmt.Println("Pressão:", pressao, "hPa")

}
