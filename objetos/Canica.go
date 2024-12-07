package objetos

import (
	"fmt"
)

type Canica struct {
	xdir      int
	ydir      int
	velocidad int
	Objeto
}

func (c *Canica) ConfigurarVelocidad(v int) {
	c.velocidad = v
}

func (c *Canica) ObtenerDireccion() [2]int {
	return [2]int{c.xdir, c.ydir}
}

func (c *Canica) ConfigurarDireccion(xdir int, ydir int) {
	c.xdir = xdir
	c.ydir = ydir
}

func (c *Canica) Mover(limX int, limY int) {
	var posX int = c.posX + c.velocidad*c.xdir
	var posY int = c.posY + c.velocidad*c.ydir

	if posX < limX {
		c.posX = posX
	} else {
		c.posX = 0
	}

	if posY < limY {
		c.posY = posY
	} else {
		c.posY = 0
	}

	if posY < 0 {
		c.posY = limY - 1
	}

	if posX < 0 {
		c.posX = limX - 1
	}

}

func (c *Canica) CambiarDireccion(xdir int, ydir int) {
	c.xdir = xdir
	c.ydir = ydir
}

func (c *Canica) String() string {
	return fmt.Sprintf("Canica-> X: %d, Y: %d, Vel: %d, DirX: %d, DirY: %d", c.posX, c.posY, c.velocidad, c.xdir, c.ydir)
}

func (c *Canica) Representar() string {
	return "C"
}
