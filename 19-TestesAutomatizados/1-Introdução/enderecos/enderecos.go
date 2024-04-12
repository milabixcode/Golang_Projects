package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)

	// RUA DOS BOBOS
	// {"RUA", "DOS", "BOBOS"}
	primeiraPalavraEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTipoValido = true
		}
	}

	if enderecoTipoValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo Inválido"
}