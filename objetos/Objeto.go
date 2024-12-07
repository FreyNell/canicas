package objetos

type Objeto struct {
	posX int
	posY int
}

func (o *Objeto) Ubicar(x int, y int) {
	o.posX = x
	o.posY = y
}

func (o *Objeto) ObtenerPosicion() [2]int {
	return [2]int{o.posX, o.posY}
}
