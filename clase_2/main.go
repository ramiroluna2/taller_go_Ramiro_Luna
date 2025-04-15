package main

import (
	"errors"
)

type dispositivo struct {
	nombre string
	estado bool
}

type Controlable interface {
	Encender() error
	Apagar() error
	EstadoActual() string
}

func (d *dispositivo) Encender() error {
	if d.estado {
		return errors.New("El dispositivo %s ya está encendido")
	}
	d.estado = true
	return nil
}

func (d *dispositivo) Apagar() error {
	if !d.estado {
		return errors.New("El dispositivo %s ya está apagado")
	}
	d.estado = false
	return nil
}

func (d *dispositivo) EstadoActual() string {
	if d.estado {
		return "El dispositivo " + d.nombre + " está encendido"
	}
	return "El dispositivo " + d.nombre + " está apagado"
}

func main() {

}
