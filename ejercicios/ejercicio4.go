package ejercicios

import "errors"

type Farolas struct {
	Posicion  int
	Radio     int
	Encendida bool
}

func EncenderFarolas(farolas []Farolas, x []int) ([]Farolas, error) {
	farolasOn := []Farolas{}

	for _, punto := range x {
		if estaCubierto(punto, farolasOn) {
			continue
		}

		var posibleFarola Farolas
		encontro := false
		for _, farAct := range farolas {
			if farAct.Cubre(punto) {
				posibleFarola = farAct
				encontro = true
			}
		}

		if !encontro {
			return nil, errors.New("no se pueden cubrir todos los puntos")
		} else {
			farolasOn = append(farolasOn, posibleFarola)
		}
	}
	return farolasOn, nil
}

func (f Farolas) Cubre(x int) bool {
	comienzo := f.Posicion - f.Radio
	final := f.Posicion + f.Radio

	return x >= comienzo && x <= final
}

func estaCubierto(punto int, farolas []Farolas) bool {
	if len(farolas) <= 0 {
		return false
	}

	for _, farAct := range farolas {
		if farAct.Cubre(punto) {
			return true
		}
	}

	return false
}
