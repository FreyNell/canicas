package objetos

import (
	"fmt"
)

type Mirror struct {
	inclinadoIzquierda bool
	Objeto
}

func (m *Mirror) ConfigurarOrientacion(izq bool) {
	m.inclinadoIzquierda = izq
}

func (m *Mirror) Colision(o IElementoMoviblePosicionable) {
	var direccion [2]int = o.ObtenerDireccion()

	switch direccion[0] {

	case 0: // solo se está moviendo vertical
		if direccion[1] == 1 {
			// bajando
			direccion[1] = 0
			if m.inclinadoIzquierda {
				direccion[0] = 1
			} else {
				direccion[0] = -1
			}
		} else {
			// direccion[1] == -1 subiendo
			direccion[1] = 0
			if m.inclinadoIzquierda {
				direccion[0] = -1
			} else {
				direccion[0] = 1
			}
		}

	case 1: // viene por la izquierda
		direccion[0] = 0
		if m.inclinadoIzquierda {
			direccion[1] = 1
		} else {
			direccion[1] = -1
		}

	case -1: // viene por la derecha
		direccion[0] = 0
		if m.inclinadoIzquierda {
			direccion[1] = -1
		} else {
			direccion[1] = 1
		}

	default:
		fmt.Println("Direccion extraña (?)")

	}

	o.CambiarDireccion(direccion[0], direccion[1])

}

func (m *Mirror) String() string {
	return fmt.Sprintf("Mirror-> X: %d, Y: %d, Orientacion: %s", m.posX, m.posY, m.inclinadoIzquierda)
}

func (m *Mirror) Representar() string {
	return "M"
}
