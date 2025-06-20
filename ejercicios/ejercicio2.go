package ejercicios

type Tarea struct {
	Nombre string
	Tiempo int
}

func Ejercicio2(tareas []Tarea) {
	for i := 0; i < len(tareas); i++ {
		interCambiar(tareas, i)
	}
}

func interCambiar(t []Tarea, indice int) {
	indiceMin := indice
	for k := indiceMin + 1; k < len(t); k++ {
		if t[k].Tiempo < t[indiceMin].Tiempo {
			indiceMin = k
		}
	}

	if indiceMin != indice {
		aux := t[indice]

		t[indice] = t[indiceMin]
		t[indiceMin] = aux
	}
}
