// TDD - Test Driven Development
// Propoe o teste antes mesmo de ter a função pronta já imaginando o que a função vai fazer
package formas

import (
	//"math"
	"testing"
)

type cenarioDeTeste struct {
	areaEsperada float64
	forma        Forma
}

func TestArea(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{10, Retangulo{5, 2}},
		{0, Retangulo{0, 0}},
		{78.539816, Circulo{5}},
		{0, Circulo{0}},
	}

	for _, cenario := range cenariosDeTeste {
		areaCalculada := cenario.forma.Area()
		if areaCalculada != cenario.areaEsperada {
			t.Errorf("A área recebida %f é diferente da esperada %f",
				areaCalculada,
				cenario.areaEsperada)
		}
	}

	// t.Run("Retângulo", func(t *testing.T) {
	// 	ret := Retangulo{10, 12}
	// 	areaEsperada := float64(120)
	// 	areaRecebida := ret.Area()

	// 	if areaEsperada != areaRecebida {
	// 		t.Fatalf("A área recebida %f é diferente da esperada %f",
	// 			areaRecebida,
	// 			areaEsperada)
	// 	}
	// })

	// t.Run("Círculo", func(t *testing.T) {
	// 	circ := Circulo{10}
	// 	areaEsperada := float64(math.Pi * 100)
	// 	areaRecebida := circ.Area()

	// 	if areaEsperada != areaRecebida {
	// 		t.Fatalf("A área recebida %f é diferente da esperada %f",
	// 			areaRecebida,
	// 			areaEsperada)
	// 	}

	// })
}
