package entorno

import (
	"fmt"
	"sync"

	"canicas/objetos"
)

type tablero struct {
	ancho int
	alto  int
	mux   sync.Mutex
	objs  []objetos.IPosicionable
}

var (
	t    *tablero
	once sync.Once
)

func ObtenerInstancia() *tablero {
	once.Do(func() {
		t = &tablero{}
	})

	return t
}

func (t *tablero) ObtenerAncho() int {
	return t.ancho
}

func (t *tablero) ObtenerAlto() int {
	return t.alto
}

func (t *tablero) ConfigurarAncho(ancho int) {
	t.ancho = ancho
}

func (t *tablero) ConfigurarAlto(alto int) {
	t.alto = alto
}

func (t *tablero) AnadirObjeto(o objetos.IPosicionable, x int, y int) {
	var posicion [2]int = o.ObtenerPosicion()
	superpos, _ := t.VerificarSuperposicion(posicion[0], posicion[1])
	if superpos {
		fmt.Sprintf("Err! - Ya existe un objeto en esta posición. (%d,%d)", x, y)
	} else {
		o.Ubicar(x, y)
		t.objs = append(t.objs, o)
	}

}

func (t *tablero) AvanzarPaso(o objetos.IElementoMoviblePosicionable) {
	o.Mover(t.ancho, t.alto)
	var posicion [2]int = o.ObtenerPosicion()
	superpos, elem := t.VerificarSuperposicion(posicion[0], posicion[1])
	if superpos {
		elem.Colision(o)
	} else {

	}

}

func (t *tablero) VerificarSuperposicion(x int, y int) (bool, objetos.IElementoEstatico) {
	for i := len(t.objs); i > 0; i-- {
		var posicion [2]int
		posicion = t.objs[i-1].ObtenerPosicion()
		if posicion[0] == x && posicion[1] == y {
			// Verificar si también implementa IElementoEstatico
			if estaticoObj, esEstatico := t.objs[i-1].(objetos.IElementoEstatico); esEstatico {
				return true, estaticoObj // Devolver solo si es IElementoEstatico
			}
		}
	}

	return false, nil
}

func (t *tablero) MostrarTablero() {
	var matriz [][]string
	var posicion [2]int

	for i := 0; i < t.alto; i++ {
		vector := make([]string, t.ancho)
		for j := 0; j < t.ancho; j++ {
			vector[j] = " "
		}
		matriz = append(matriz, vector)
	}

	// Asigna los objetos a sus posiciones en la matriz
	for i := 0; i < len(t.objs); i++ {
		posicion = t.objs[i].ObtenerPosicion()
		matriz[posicion[1]][posicion[0]] = t.objs[i].Representar()
	}

	for i := 0; i < t.alto; i++ {
		for j := 0; j < t.ancho; j++ {
			fmt.Print(matriz[i][j])
		}
		fmt.Print("\n")
	}
}

func (t *tablero) String() string {
	return fmt.Sprintf("Tablero-> Alto: %d, Ancho: %d", t.alto, t.ancho)
}
