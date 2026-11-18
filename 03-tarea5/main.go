package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	LimiteSubtitulo = 42
	ErrUso          = 1
	ErrVacio        = 2
	ErrLimite       = 3
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Uso: analizar \"<texto>\"")
		os.Exit(ErrUso)
	}

	texto := os.Args[1]
	if strings.TrimSpace(texto) == "" {
		fmt.Fprintln(os.Stderr, "Error: el texto no puede estar vacio ni contener solo espacios")
		os.Exit(ErrVacio)
	}

	var palabras, letras, digitos, mayusculas, multibyte int
	enPalabra := false

	for _, r := range texto {
		if unicode.IsLetter(r) {
			letras++
			if unicode.IsUpper(r) {
				mayusculas++
			}
		}
		if unicode.IsDigit(r) {
			digitos++
		}
		if utf8.RuneLen(r) > 1 {
			multibyte++
		}
		if unicode.IsSpace(r) {
			enPalabra = false
		} else if !enPalabra {
			palabras++
			enPalabra = true
		}
	}

	bytes := len(texto)
	runas := utf8.RuneCountInString(texto)

	fmt.Printf("texto: %q\n", texto)
	fmt.Printf("bytes: %d\n", bytes)
	fmt.Printf("runas: %d\n", runas)
	fmt.Printf("palabras: %d\n", palabras)
	fmt.Printf("letras: %d\n", letras)
	fmt.Printf("dígitos: %d\n", digitos)
	fmt.Printf("mayúsculas: %d\n", mayusculas)
	fmt.Printf("runas de más de 1 byte: %d\n", multibyte)

	if runas > LimiteSubtitulo {
		fmt.Fprintf(os.Stderr, "Error: el texto tiene %d runas y supera el limite de %d\n", runas, LimiteSubtitulo)
		os.Exit(ErrLimite)
	}

	fmt.Println("límite de 42 runas: dentro")
}