package objetos

type IElementoMovible interface {
	Mover(limX int, limY int)
	ObtenerDireccion() [2]int
	CambiarDireccion(xdir int, ydir int)
	ConfigurarDireccion(xdir int, ydir int)
	ConfigurarVelocidad(v int)
}
