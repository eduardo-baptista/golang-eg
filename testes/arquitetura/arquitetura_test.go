package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	// roda de forma paralela com os outros testes
	t.Parallel()

	//  nao executa o teste com base na arquitetura do sistema
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	// ..
	t.Fail()
}
