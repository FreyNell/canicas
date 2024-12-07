package objetos

type IPosicionable interface {
	Ubicar(x int, y int)
	ObtenerPosicion() [2]int
	Representar() string
}
