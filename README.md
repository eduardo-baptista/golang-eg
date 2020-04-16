### Comandos:

- go env - exibi variaveis de ambiente entre elas o GOROOT que exibe o local de instalação do Go Lang e o GOPATH é o local onde se encontra o projeto
- got get - instala pacotes externos
- go version - exibe a versão do go
- godoc -http:6060 - habilita a documentação offline
- go vet <SCRIPT> - Retorna diagnostico do código
- go test -cover - Mostra a porcentagem de cobertura do teste em relação ao código
- go test -coverprofile=resultado.out - gera um relatório de testes
- go tool cover -func=resultado.out - exibe o relatório gerado em go test
- go tool cover -html=resultado.out - exibe o relatório gerado em go test
