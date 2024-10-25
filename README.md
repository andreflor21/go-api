# API em go

## [Padrão de projeto](http://github.com/golang-standards/project-layout)

### Diretórios

-   **internal**

    Todas as coisas referente a esse projeto, não será exportada para outros projetos

-   **pkg**

    Coisas comuns em projetos ou outras libs externas, será exportado podendo ser usado em outros projetos

-   **cmd**

    Onde será gerado o executavel, onde o arquivo `main.go` estará dentro da pasta com o nome do projeto ex: `go-api/main.go`

-   **configs**

    Onde será guardados os templates de configurações, variáveis de ambiente, etc

-   **test**

    Guardar arquivos de testes e2e, postman collections para testar, arquivos de testes adicionais

-   **api**

    Guardar documentação da API
