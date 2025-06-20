package ejercicios

type Objeto struct {
	Nombre string
	Peso   int
	Valor  int
}

func Ejercicio3(objetos []Objeto, capacidad int) map[Objeto]float32 {
	mochila := make(map[Objeto]float32)
	//ordenar objetos
	sortPorValorNeto(objetos)

	for _, objeto := range objetos {
		if capacidad == 0 {
			break
		}

		if objeto.Peso <= capacidad {
			capacidad -= objeto.Peso
			mochila[objeto] = 1
		} else {
			canTParcial := float32(capacidad) / float32(objeto.Peso)
			mochila[objeto] = canTParcial
			capacidad = 0
		}
	}
	return mochila
}

func (o Objeto) ValorNeto() float32 {
	return float32(o.Valor) / float32(o.Peso)
}

func sortPorValorNeto(objetos []Objeto) {
	n := len(objetos)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if objetos[j].ValorNeto() > objetos[maxIdx].ValorNeto() {
				maxIdx = j
			}
		}
		// Intercambio: el objeto con mayor valor/peso al frente
		objetos[i], objetos[maxIdx] = objetos[maxIdx], objetos[i]
	}
}
