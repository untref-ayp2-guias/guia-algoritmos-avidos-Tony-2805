package ejercicios

var billetes = []int{10000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}

func Cambiar(cantidad int) map[int]int {
	billetera := make(map[int]int)

	CambiarRecur(cantidad, 0, &billetera)
	return billetera
}

func CambiarRecur(cantidad, indCambio int, b *map[int]int) {
	if cantidad <= 0 {
		return
	}
	if cantidad < billetes[indCambio] {
		for i := indCambio; i < len(billetes); i++ {
			if cantidad >= billetes[i] {
				indCambio = i
				break
			}
		}
	}
	agregarBillete(billetes[indCambio], *b)
	cantidad -= billetes[indCambio]

	CambiarRecur(cantidad, indCambio, b)
}

func CambiarNormal(cantidad int) map[int]int {
	indCambio := 0
	billetera := make(map[int]int)

	for cantidad > 0 {
		if cantidad >= billetes[indCambio] {
			agregarBillete(billetes[indCambio], billetera)
			cantidad -= billetes[indCambio]
		} else {
			indCambio++
		}
	}

	return billetera
}

func agregarBillete(billete int, b map[int]int) {
	cantidad, ok := b[billete]

	if !ok {
		b[billete] = 1
		return
	}
	nuevaCantidad := cantidad + 1
	b[billete] = nuevaCantidad

}
