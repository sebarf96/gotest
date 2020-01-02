package calculator_test

import (
	
	"testing"
)

func TestAdd(t *testing.T) {
	sum := calculator.Add(1, 2)
	if sum != 3 {
		t.Error("La suma no es correcta ")
		t.Fail()
	} else {
		t.Log("Test finalizado correctamente")
	}
}
