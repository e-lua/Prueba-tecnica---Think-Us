// main_test.go

package main

import "testing"

func TestEncrypt(t *testing.T) {

	// Caso 1: Mensaje encriptado
	caso_1 := Encrypt("dcj", "I love prOgrAmming!")
	if caso_1 != "dcjI ldcjovdcje prdcjOgrdcjAmmdcjing!" {
		t.Errorf("TestEncrypt | Caso 1: Mensaje encriptado incorrecto")
	}

	// Caso 2: Clave nula o vacia
	caso_2 := Encrypt("", "I love prOgrAmming!")
	if caso_2 != "DCJI lDCJovDCJe prDCJOgrDCJAmmDCJing!" {
		t.Errorf("TestEncrypt | Caso 2: Mensaje encriptado incorrecto")
	}

	// Caso 3: Mensaje vacio
	caso_3 := Encrypt("dcj", "")
	if caso_3 != "" {
		t.Errorf("TestEncrypt | Caso 3: Mensaje encriptado incorrecto")
	}

}

func TestRemoveConsecutives(t *testing.T) {

	// Caso 1: Arreglo inicial
	caso_1 := RemoveConsecutives([]int{3, 4, -7, 5, -6, 2, 5, -1, 8})
	for _, value := range caso_1 {
		if value != 5 && value != 8 {
			t.Errorf("RemoveConsecutives | Caso 1: Valores obtenidos incorrectos")
		}
	}

	// Caso 2: Arreglo vacio
	caso_2 := RemoveConsecutives([]int{})
	if len(caso_2) != 0 {
		t.Errorf("RemoveConsecutives | Caso 2: Valores obtenidos incorrectos")
	}

}
