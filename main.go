package goarea

import "math"

// Pi ... (o comentário é obrigatório)
const Pi = 3.141592653979

// Circ área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect área do retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não visível
func triagEq(base, altura float64) float64 {
	return (base * altura) / 2
}
