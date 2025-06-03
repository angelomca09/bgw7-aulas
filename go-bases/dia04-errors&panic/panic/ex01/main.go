/* Exercício 1 - Dados do cliente

Uma empresa de contabilidade precisa ter acesso aos dados de seus funcionários para poder fazer várias liquidações.
Para isso, eles têm todos os detalhes necessários em um arquivo .txt .

Você terá que desenvolver a funcionalidade para poder ler o arquivo .txt indicado pelo cliente, mas ele não passou o arquivo para
ser lido pelo nosso programa.
Desenvolva o código necessário para ler os dados do arquivo chamado "customers.txt" (lembre-se do que vimos sobre o pkg "os").
Como não temos o arquivo necessário, receberemos um erro e, nesse caso, o programa deve gerar um pânico ao tentar ler um arquivo
que não existe, exibindo a seguinte mensagem “The indicated file was not found or is damaged”.

Apesar disso, a mensagem "execução concluída" sempre será impressa no console.
*/

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("execução concluída")
	}()

	readFile("customers.txt")
}

func readFile(filePath string) (content []byte, err error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	if filePath == "" {
		err = errors.New("file path cannot be empty")
		return
	}
	content, err = os.ReadFile(filePath)
	if err != nil {
		err = errors.New("The indicated file was not found or is damaged")
		panic(err)
	}
	return
}
