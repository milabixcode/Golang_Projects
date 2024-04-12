// TESTE DE UNIDADE

// Um pacote por pasta
// exceção: TESTES

package enderecos_test

import (
	// . = alias
	. "introducao-testes/enderecos"
	"testing"
)	

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//TestNomeDaFunção
func TestTipoDeEndereco(t *testing.T) {
	//teste rodado em paralelo com as outras funções q tbm rodam em paralelo
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Estrada ABC", "Estrada"},
		{"Rodovia ABC", "Rodovia"},
		{"Praça ABC", "Tipo Inválido"},
		{"RUA DOS BOBOS", "Rua"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Teste quebrou!")
	}
}

// enderecoParaTeste := "Rua Paulista"

// tipoDeEnderecoEsperado := "Avenida"
// tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

// if tipoDeEnderecoEsperado != tipoDeEnderecoRecebido {
// 	t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
// 		tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
// }
