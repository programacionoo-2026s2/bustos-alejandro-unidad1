// Paquete almacenamiento provee constantes y funciones para convertir unidades digitales.
package almacenamiento

import "math"

// B, KB, MB y GB definen la equivalencia en bytes de cada unidad de almacenamiento.
const (
	B  = 1.0
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// Bytes devuelve cuántos bytes vale la unidad indicada en mayúsculas y si es válida.
func Bytes(unidad string) (float64, bool) {
	switch unidad {
	case "B":
		return B, true
	case "KB":
		return KB, true
	case "MB":
		return MB, true
	case "GB":
		return GB, true
	default:
		return 0, false
	}
}

// Convertir calcula el valor convertido entre unidades y lo redondea a dos decimales.
func Convertir(valor, bytesOrigen, bytesDestino float64) float64 {
	resultado := valor * bytesOrigen / bytesDestino
	return redondear(resultado)
}

func redondear(v float64) float64 {
	return math.Round(v*100) / 100
}