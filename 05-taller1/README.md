# 05-taller1: Estadísticas de notas con función variádica y error centinela

## Descripción
Programa en Go que implementa la función variádica `estadisticas(nums ...float64)` con resultados nombrados (`menor`, `mayor`, `promedio`, `err`). Calcula la calificación mínima, la máxima y el promedio de cualquier cantidad de notas usando las funciones integradas `min` y `max`, y devuelve el error centinela `ErrSinDatos` cuando se invoca sin argumentos.

## Cómo ejecutar
Desde la carpeta `05-taller1`, ejecutar:

    go run .

Qué imprime cada línea de salida:
- **Línea 1 (`4 9.75 7.0625 <nil>`):** Resultado directo de `estadisticas(8.5, 6, 9.75, 4)` pasando valores sueltos; muestra el menor, el mayor, el promedio sin redondear y `<nil>` porque no hubo error.
- **Línea 2 (`min=7.00 max=10.00 prom=8.17`):** Resultado formateado a 2 decimales (`%.2f`) tras expandir el slice `notas := []float64{10, 7, 7.5}` con `estadisticas(notas...)`.
- **Línea 3 (`error: no se recibieron datos`):** Mensaje impreso tras llamar a `estadisticas()` sin argumentos y verificar el error centinela mediante `errors.Is(err, ErrSinDatos)`.

## Pruebas realizadas

Salida obtenida al ejecutar `go run .`:

    4 9.75 7.0625 <nil>
    min=7.00 max=10.00 prom=8.17
    error: no se recibieron datos

## Preguntas

**a) ¿Por qué los resultados se llaman `menor` y `mayor` y no `min` y `max`? Explica qué pasaría con la llamada `min(menor, n)` si el resultado se llamara `min`.**
Porque desde Go 1.21 `min` y `max` son funciones integradas del lenguaje. Los resultados con nombre (`menor, mayor, promedio float64, err error`) actúan como variables locales declaradas dentro de `estadisticas`; si un resultado se llamara `min`, esa variable de tipo `float64` ocultaría (haría *shadowing*) a la función integrada `min`, y al intentar ejecutar `min(menor, n)` el compilador daría un error indicando que no se puede llamar a una variable que no es una función (`cannot call non-function min (variable of type float64)`).

**b) ¿Qué hacen los tres puntos en `estadisticas(notas...)`? ¿Qué mensaje te da Go si los quitas y por qué?**
Los tres puntos después del slice `notas` desempaquetan sus elementos para pasarlos como argumentos individuales de tipo `float64` a la función variádica. Si se quitan los tres puntos y se escribe `estadisticas(notas)`, Go detiene la compilación con el mensaje `cannot use notas (variable of type []float64) as float64 value in argument to estadisticas`, porque la firma espera valores sueltos `float64` y no una lista completa de tipo `[]float64`.
