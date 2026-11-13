package main

import "fmt"

type Permiso uint8

const (
	Ejecutar Permiso = 1 << iota
	Escribir
	Leer
)

func describir(p Permiso) string {
	r := '-'
	w := '-'
	x := '-'

	if p&Leer != 0 {
		r = 'r'
	}
	if p&Escribir != 0 {
		w = 'w'
	}
	if p&Ejecutar != 0 {
		x = 'x'
	}

	return string([]byte{byte(r), byte(w), byte(x)})
}

func main() {
	usuario := Leer | Escribir | Ejecutar
	grupo := Leer | Ejecutar
	otros := Leer

	fmt.Printf("%s%s%s\n", describir(usuario), describir(grupo), describir(otros))
	fmt.Printf("numérico: %d%d%d\n", usuario, grupo, otros)

	otros |= Escribir
	grupo &^= Ejecutar

	fmt.Printf("%s%s%s\n", describir(usuario), describir(grupo), describir(otros))
}
