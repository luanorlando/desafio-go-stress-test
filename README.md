# Objetivo
Criar um sistema de CLI (Command Line Interface) em Go para realizar testes de carga em serviços web. O usuário deverá fornecer a URL do serviço, o número total de requisições e a quantidade de chamadas simultâneas. Ao final, o sistema deve gerar um relatório detalhado da execução.

# Como funciona
O projeto usa CobraCLI e o objetivo é realizar teste de carga em serviços web. O projeto foi configurado para rodar e passar três parâmetros obrigatórios

# Entrada de Parâmetros
A aplicação deve aceitar os seguintes parâmetros via linha de comando:
- [x] --url: URL do serviço a ser testado.
- [x] --requests: Número total de requisições a serem realizadas.
- [x] --concurrency: Número de chamadas simultâneas.
 
# Rodando o projeto sem Docker 

```go
go run main.go --url=http://google.com --requests=10 --concurrency=2
```

|url|Número de requests|concorrência|
|---|------------------|------------|
|http://google.com|10|2|

Isso acontece porque o arquivo de iniciação do cobraCLI em root.go foi configurado para funcionar exatamente assim

#### Atribuindo os valores de entrada as variaveis
```Go
rootCmd.Flags().StringVar(&url, "url", "", "URL do serviço a ser testado (obrigatório)")
rootCmd.Flags().IntVar(&requests, "requests", 0, "Número total de requisições (obrigatório)")
rootCmd.Flags().IntVar(&concurrency, "concurrency", 0, "Número de chamadas simultâneas (obrigatório)")
```

E para informar que esses valores são obrigatórios usamos o método `MarkFlagRequired`
```Go
rootCmd.MarkFlagRequired("url")
rootCmd.MarkFlagRequired("requests")
rootCmd.MarkFlagRequired("concurrency")
```

> Se rodar o projeto sem passar 1 desses valores o programa irá dar erro.

#### Arquivo de entrada

O arquivo de entrada da aplicação contém apenas a chamada da função `Execute` do package `cmd`

```go
func main() {
	cmd.Execute()
}
```

#### Buildando o projeto via Docker

Para buildar o projeto usando o docker devemos respeitar as mesma exigências do cobraCLI onde as variaveis são obrigátórias para isso pasta usar o comando

```go
docker run load-tester --url=http://google.com --requests=100 --concurrency=5
```

Com isso será obtido o relatório final

> Iniciando teste de carga para: http://google.com
> Requisições: 100 | Concorrência: 5
>==========================================
>           RELATÓRIO DE TESTE             
>==========================================
>Tempo total gasto:       9.339263837s
>Total de requisições:    100
>Requisições com 200 OK:  100
>Distribuição de Status HTTP:
>  - Status 200:            100
>==========================================
 
## Execução do Teste:
- [x] Realizar requisições HTTP para a URL especificada.
- [x] Distribuir as requisições de acordo com o nível de concorrência definido.
- [x] Garantir que o número total de requisições (--requests) seja cumprido exato.

​## 2.Geração de Relatório: Ao final da execução, o sistema deve apresentar no console as seguintes informações:
- [x] Tempo total gasto na execução.
- [x] Quantidade total de requests realizados.
- [x] Quantidade de requests com status HTTP 200.
- [x] Distribuição de outros códigos de status HTTP (ex: quantidade de 404, 500, etc...).
 
# Execução via Docker (Obrigatório)
- [x] A aplicação deve ser empacotada em uma imagem Docker para facilitar a execução. O comando de teste deve seguir o formato abaixo:
docker run <sua-imagem-docker> --url=http://google.com --requests=1000 --concurrency=10
 
# Entregável
- [x] Código Fonte: Repositório contendo a implementação do CLI.
- [x] Dockerfile: Arquivo de configuração para construção da imagem.
- [x] README: Instruções detalhadas de como buildar a imagem e executar o comando de teste.
 
# Regras de Entrega
Repositório Exclusivo: O repositório deve conter apenas o projeto em questão.
Branch Principal: Todo o código deve estar na branch main.
