package objetos

import (
	"fmt"
)

type Portal struct {
	Objeto
	conexion *Portal
}

func (p *Portal) Colision(o IElementoMoviblePosicionable) {
	o.Ubicar(p.conexion.posX, p.conexion.posY)
}

func (p *Portal) Conectar(o *Portal) {
	p.conexion = o
}

func (c *Portal) String() string {
	return fmt.Sprintf("Portal-> X: %d, Y: %d", c.posX, c.posY)
}

func (b *Portal) Representar() string {
	return "P"
}
