# Aulas IT Bootcamp - Go
Repositório para salvar o conteúdo das aulas do IT Bootcamp - Go

## Configurações
Para criar o módulo `bgw7-aulas`, foi usado:
```bash
go mod init bgw7-aulas
```
Esse comando criou o arquivo `go.mod`.

Para adicionar a dependencia de `testify/require`, foi usado:
```bash
go get github.com/stretchr/testify/require
```
Por sua vez, este comando gerou os arquivos `go.sum`.


## CR - Caio

### Configuração do Ambiente de Teste

Atualmente, a configuração do banco de dados está hardcoded no código. Talvez seja interessante utilizar variáveis de ambiente para tornar a configuração mais flexível e segura.

Além disso, vale a pena pesquisar sobre **thread safety** (proteção contra múltiplas execuções simultâneas). Por exemplo, pode-se utilizar:

```go
var registerOnce sync.Once
```
para garantir que certas operações (como o registro do driver do banco) ocorram apenas uma vez.


### Testes de Repositório

Achei interessante a abordagem de separar cada teste em uma função única. Isso ajuda a manter os testes bem organizados e isolados, facilitando a manutenção e a leitura do código.

### Tratamento de Erros

Consideraria criar erros específicos em vez de retornar apenas `nil` em todos os cenários de falha. Isso facilita o diagnóstico dos problemas e melhora a robustez da aplicação.


