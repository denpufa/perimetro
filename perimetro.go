package perimetro

import (
	"math"
)

// PerimetroError representa um erro relacionado ao cálculo do perímetro
type PerimetroError struct {
	Mensagem string
}

func (e *PerimetroError) Error() string {
	return e.Mensagem
}

// CalcularPerimetroTriangulo calcula o perímetro de um triângulo equilátero
func CalcularPerimetroTriangulo(ladoTriangulo float32) (float32, error) {
	if ladoTriangulo <= 0 {
		return 0, &PerimetroError{"Valor inválido. O lado do triângulo deve ser maior que zero."}
	}

	perimetro := ladoTriangulo * 3

	return perimetro, nil
}

// CalcularPerimetroRetangulo calcula o perímetro de um retângulo
func CalcularPerimetroRetangulo(BaseRetangulo, alturaRetangulo float32) (float32, error) {
	if BaseRetangulo <= 0 || alturaRetangulo <= 0 {
		return 0, &PerimetroError{"Valores inválidos. A base e a altura do retângulo devem ser maiores que zero."}
	}

	perimetro := BaseRetangulo*2 + alturaRetangulo*2

	return perimetro, nil
}

// CalcularPerimetroQuadrado calcula o perímetro de um quadrado
func CalcularPerimetroQuadrado(ladoQuadrado float32) (float32, error) {
	if ladoQuadrado <= 0 {
		return 0, &PerimetroError{"Valor inválido. O lado do quadrado deve ser maior que zero."}
	}

	perimetro := 4 * ladoQuadrado

	return perimetro, nil
}

// CalcularPerimetroCirculo calcula o perímetro de um círculo
func CalcularPerimetroCirculo(raio float32) (float32, error) {
	if raio <= 0 {
		return 0, &PerimetroError{"Valor inválido. O raio do círculo deve ser maior que zero."}
	}

	perimetro := 2 * math.Pi * raio

	return perimetro, nil
}
