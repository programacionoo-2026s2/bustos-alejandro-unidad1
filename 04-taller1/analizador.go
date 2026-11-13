package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	texto := "Programación en Go: rápido, simple y ¡divertido! Año 2026."

	// 1. Contar bytes y runas
	bytes := len(texto)
	runas := utf8.RuneCountInString(texto)

	// 2. Contar palabras
	palabras := len(strings.Fields(texto))

	// 3. Contar mayúsculas y dígitos
	mayusculas := 0
	digitos := 0
	for _, r := range texto {
		if unicode.IsUpper(r) {
			mayusculas++
		}
		if unicode.IsDigit(r) {
			digitos++
		}
	}

	// 4. Contar letras con tilde o ñ
	conTilde := 0
	letrasEspeciales := "áéíóúÁÉÍÓÚñÑ"
	for _, r := range texto {
		if strings.ContainsRune(letrasEspeciales, r) {
			conTilde++
		}
	}

	// Mostrar resultados
	fmt.Printf("Bytes: %d\n", bytes)
	fmt.Printf("Runas: %d\n", runas)
	fmt.Printf("Palabras: %d\n", palabras)
	fmt.Printf("Mayúsculas: %d\n", mayusculas)
	fmt.Printf("Dígitos: %d\n", digitos)
	fmt.Printf("Con tilde/ñ: %d\n", conTilde)

	// 5. Explicación de la diferencia entre bytes y runas
	fmt.Println("\nExplicación: Los bytes y runas no coinciden porque Go usa UTF-8. Los caracteres normales pesan 1 byte, pero los especiales (tildes, ñ, ¡) ocupan 2 bytes. Por eso hay más bytes que letras visibles.")
}
