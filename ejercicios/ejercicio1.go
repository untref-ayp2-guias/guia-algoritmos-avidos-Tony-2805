package ejercicios

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

func SelectorRecursivo(actividades []Actividad) []Actividad {
	//var actSelect []Actividad
	if len(actividades) == 0 {
		return nil
	}
	//return interSelectorRecur(actividades, actSelect, i)
	return seleccionarRecur(actividades, 0)
}

func interSelectorRecur(acts, actsSelect []Actividad, indice int) []Actividad {

	if len(acts) <= indice {
		return actsSelect //no hace falta devolver un cero, se puede devolver las actividades ya elegidas
	}

	if indice == 0 {
		actsSelect = append(actsSelect, acts[0])
		indice++
	} else {
		ultAct := actsSelect[len(actsSelect)-1]
		posActi := ultAct
		for i := indice; i < len(acts); i++ {

			actual := acts[i]
			if actual.Inicio > ultAct.Fin {
				posActi = actual
				indice = i
				actsSelect = append(actsSelect, actual)
				break
			}
		}
		if posActi == ultAct {
			return actsSelect
		}
	}

	selecFinal := interSelectorRecur(acts, actsSelect, indice)
	return selecFinal
}

func seleccionarRecur(actividades []Actividad, i int) []Actividad {

	if i >= len(actividades) {
		return nil
	}

	// Seleccionamos la actividad actual
	seleccionada := actividades[i]

	// Buscar la próxima actividad compatible
	var siguiente int
	for siguiente = i + 1; siguiente < len(actividades); siguiente++ {
		if actividades[siguiente].Inicio >= seleccionada.Fin {
			break
		}
	}

	// Llamada recursiva con la próxima actividad compatible
	return append(seleccionarRecur(actividades, siguiente), seleccionada)
}
