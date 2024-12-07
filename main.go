package main

import (
	"os"
    "os/exec"
    "runtime"
	"time"

	"canicas/entorno"
	"canicas/objetos"
)

func clearTerminal() {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func main() {
	tablero := entorno.ObtenerInstancia()
	tablero.ConfigurarAlto(20)
	tablero.ConfigurarAncho(20)

	c1 := &objetos.Canica{}
	c2 := &objetos.Canica{}
	c3 := &objetos.Canica{}

	b1 := &objetos.Bumper{}
	b2 := &objetos.Bumper{}
	b3 := &objetos.Bumper{}
	b4 := &objetos.Bumper{}

	m1 := &objetos.Mirror{}
	m2 := &objetos.Mirror{}

	p1 := &objetos.Portal{}
	p2 := &objetos.Portal{}

	tablero.AnadirObjeto(c1, 0, 0)
	tablero.AnadirObjeto(c2, 18, 0)
	tablero.AnadirObjeto(c3, 19, 0)

	tablero.AnadirObjeto(b1, 2, 2)
	tablero.AnadirObjeto(b2, 2, 18)
	tablero.AnadirObjeto(b3, 18, 2)
	tablero.AnadirObjeto(b4, 18, 18)

	tablero.AnadirObjeto(m1, 10, 0)
	tablero.AnadirObjeto(m2, 19, 10)

	tablero.AnadirObjeto(p1, 10, 10)
	tablero.AnadirObjeto(p2, 10, 19)

	c1.ConfigurarDireccion(1, 0)
	c2.ConfigurarDireccion(0, 1)
	c3.ConfigurarDireccion(0, 1)

	c1.ConfigurarVelocidad(1)
	c2.ConfigurarVelocidad(1)
	c3.ConfigurarVelocidad(1)

	m1.ConfigurarOrientacion(true)
	m2.ConfigurarOrientacion(false)

	p1.Conectar(p2)
	p2.Conectar(p1)

	tablero.MostrarTablero()
	for i := 0; i < 100; i++ {
		time.Sleep(500 * time.Millisecond)
		clearTerminal()
		tablero.AvanzarPaso(c1)
		tablero.AvanzarPaso(c2)
		tablero.AvanzarPaso(c3)
		tablero.MostrarTablero()
	}

}
