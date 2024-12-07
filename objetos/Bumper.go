package objetos

import (
	"fmt"
)

type Bumper struct {
	Objeto
}

func (b *Bumper) Colision(o IElementoMoviblePosicionable) {
	var direccion [2]int = o.ObtenerDireccion()
	o.CambiarDireccion(-1*direccion[0], -1*direccion[1])
}

func (b *Bumper) String() string {
	return fmt.Sprintf("Bumper-> X: %d, Y: %d", b.posX, b.posY)
}

func (b *Bumper) Representar() string {
	return "B"
}
