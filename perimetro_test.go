package perimetro

import (
	"math"
	"testing"
)


// CAMINHOS FELIZES
func TestCalcularPerimetroTriangulo(t *testing.T) {
	lado := float32(4)
	perimetroEsperado := float32(12)
	perimetro, err := CalcularPerimetroTriangulo(lado)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if perimetro != perimetroEsperado {
		t.Errorf("Perímetro esperado: %f, obtido: %f", perimetroEsperado, perimetro)
	}
}

func TestCalcularPerimetroRetangulo(t *testing.T) {
	base := float32(5)
	altura := float32(3)
	perimetroEsperado := float32(16)
	perimetro, err := CalcularPerimetroRetangulo(base, altura)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if perimetro != perimetroEsperado {
		t.Errorf("Perímetro esperado: %f, obtido: %f", perimetroEsperado, perimetro)
	}
}

func TestCalcularPerimetroQuadrado(t *testing.T) {
	lado := float32(3)
	perimetroEsperado := float32(12)
	perimetro, err := CalcularPerimetroQuadrado(lado)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if perimetro != perimetroEsperado {
		t.Errorf("Perímetro esperado: %f, obtido: %f", perimetroEsperado, perimetro)
	}
}

func TestCalcularPerimetroCirculo(t *testing.T) {
	raio := float32(2)
	perimetroEsperado := float32(2 * math.Pi * raio)
	perimetro, err := CalcularPerimetroCirculo(raio)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if perimetro != perimetroEsperado {
		t.Errorf("Perímetro esperado: %f, obtido: %f", perimetroEsperado, perimetro)
	}
}
// CAMINHOS INFELIZES

func TestErroCalcularPerimetroTriangulo(t *testing.T) {
	lado := float32(0)
	_, err := CalcularPerimetroTriangulo(lado)
	if err == nil {
		t.Errorf("Esperava um erro, mas nenhum erro foi retornado")
	}
	if err.Error() != "Valor inválido. O lado do triângulo deve ser maior que zero." {
		t.Errorf("Mensagem de erro esperada: 'Valor inválido. O lado do triângulo deve ser maior que zero.', obtida: '%s'", err.Error())
	}
}

func TestErroCalcularPerimetroRetangulo(t *testing.T) {
	base := float32(-2)
	altura := float32(5)
	_, err := CalcularPerimetroRetangulo(base, altura)
	if err == nil {
		t.Errorf("Esperava um erro, mas nenhum erro foi retornado")
	}
	if err.Error() != "Valores inválidos. A base e a altura do retângulo devem ser maiores que zero." {
		t.Errorf("Mensagem de erro esperada: 'Valores inválidos. A base e a altura do retângulo devem ser maiores que zero.', obtida: '%s'", err.Error())
	}
}

func TestErroCalcularPerimetroQuadrado(t *testing.T) {
	lado := float32(0)
	_, err := CalcularPerimetroQuadrado(lado)
	if err == nil {
		t.Errorf("Esperava um erro, mas nenhum erro foi retornado")
	}
	if err.Error() != "Valor inválido. O lado do quadrado deve ser maior que zero." {
		t.Errorf("Mensagem de erro esperada: 'Valor inválido. O lado do quadrado deve ser maior que zero.', obtida: '%s'", err.Error())
	}
}

func TestErroCalcularPerimetroCirculo(t *testing.T) {
	raio := float32(-3)
	_, err := CalcularPerimetroCirculo(raio)
	if err == nil {
		t.Errorf("Esperava um erro, mas nenhum erro foi retornado")
	}
	if err.Error() != "Valor inválido. O raio do círculo deve ser maior que zero." {
		t.Errorf("Mensagem de erro esperada: 'Valor inválido. O raio do círculo deve ser maior que zero.', obtida: '%s'", err.Error())
	}
}
