package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Sempre começa com Test
// go test só roda os testes
// go test .\ ... roda tdos os tests
// go test -v -> mostra o teste q está rodando e o tempo gasto
// go test --cover -> ver cobertura dos testes
// go test --coverprofile -> cobertura.txt gerar um relatorio com o teste
// go tool cover --func=cobertura.txt -> lê o arquivo de cobertura é mostra de forma legivel
// go tool cover --html=cobertura.txt -> lê o arquivo de cobertura é mostra em um relatorio html
func TestTipoDeEndereco(t *testing.T) {
	// permite rodar os teste em paralelo roda todos os q tem essa chamada a função
	t.Parallel()
	enderecoParaTeste := "Avenida paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo de endereço recebido é diferente do esperado! Esperava %s Recebeu %s",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido,
		)
	}

}

func TestTipoDeEnderecoComCenarios(t *testing.T) {
	// permite rodar os teste em paralelo roda todos os q tem essa chamada a função
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Estrada ABC", "Estrada"},
		{"Praça ABC", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de endereço recebido é diferente do esperado! Esperava %s Recebeu %s",
				cenario.retornoEsperado,
				tipoDeEnderecoRecebido,
			)
		}
	}

}
