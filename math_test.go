package main

import "testing"

func TestSoma(t *testing.T) {
	result := Soma(15,15)
	expected := 30

	if result != expected {
		t.Errorf("Resultado da soma inválido: Resultado %d. Esperado %d", result,  expected)
	}
}