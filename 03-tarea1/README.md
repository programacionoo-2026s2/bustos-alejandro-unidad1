Tarea de clase 1: Programa suma.go

1. ¿Qué imprime?
Imprime: `total: 0`

2. Explicación del problema (Ámbito)
El problema ocurre debido al **ámbito** (*scope*) de las variables y el uso del operador de asignación corta `:=`. Dentro de la función `suma()`, se declara una variable `total := 0`. Sin embargo, al entrar al bucle `for`, se utiliza `total := total + i`. Al usar `:=`, Go no está actualizando la variable `total` de la función, sino que está declarando una **nueva variable local** llamada `total` cuyo **ámbito** se limita exclusivamente al bloque del bucle. Por lo tanto, la variable `total` original de la función nunca cambia su valor de `0` y es la que se retorna al final.

3. Corrección cambiando un solo carácter
Para corregirlo cambiando un solo carácter, se debe reemplazar los dos puntos `:` por un signo igual `=`, cambiando `:=` por `=`. Esto realiza una asignación sobre la variable existente en lugar de declarar una nueva:
```go
total = total + i