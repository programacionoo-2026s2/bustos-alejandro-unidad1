# 03-tarea5: Analizador de subtítulos accesibles

## Descripción
Programa de consola escrito en Go que recibe un texto por línea de comandos y calcula su tamaño real en bytes (`len`), la cantidad de caracteres Unicode o runas (`utf8.RuneCountInString`), el número de palabras, letras, dígitos, mayúsculas y runas multibyte. Además, comprueba si el texto respeta el límite de 42 runas recomendado para subtítulos accesibles.

## Cómo ejecutar
Desde la carpeta `03-tarea5`, compilar y ejecutar encerrando el texto entre comillas dobles para que la terminal lo pase como un único argumento en `os.Args[1]`:

    go build -o analizar .
    ./analizar "Canción de cuna 2026"

Si el texto tiene espacios y no se usan comillas, la consola divide cada palabra en un argumento separado (`len(os.Args) > 2`) y el programa termina con error de uso.

Significado de los códigos de salida:
- `0`: Ejecución correcta y texto dentro del límite de 42 runas.
- `1` (`ErrUso`): No se envió exactamente 1 argumento entre comillas.
- `2` (`ErrVacio`): El texto está vacío o solo contiene espacios en blanco.
- `3` (`ErrLimite`): El texto supera las 42 runas permitidas.

## Pruebas realizadas

| Comando | bytes | runas | palabras | letras | dígitos | mayúsc. | runas >1 byte | Código |
| :--- | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| `analizar "Hola mundo"` | 10 | 10 | 2 | 9 | 0 | 1 | 0 | `0` |
| `analizar "Canción de cuna 2026"` | 21 | 20 | 4 | 13 | 4 | 1 | 1 | `0` |
| `analizar "ÁÉÍ ñandú"` | 14 | 9 | 2 | 8 | 0 | 3 | 5 | `0` |
| `analizar "  espacios   múltiples  "` | 25 | 24 | 2 | 17 | 0 | 0 | 1 | `0` |
| `analizar "   "` | - | - | - | - | - | - | - | `2` |
| `analizar "Este subtítulo es demasiado largo para leerse bien"` | 51 | 50 | 8 | 43 | 0 | 1 | 1 | `3` |
| `analizar` | - | - | - | - | - | - | - | `1` |
| `analizar hola mundo` | - | - | - | - | - | - | - | `1` |

## Preguntas

**a) En `Canción de cuna 2026` tu programa muestra 21 bytes y 20 runas. ¿Por qué no coinciden y cuál de los dos números usarías para el límite de subtítulos?**
No coinciden porque Go guarda los `string` en UTF-8: los 19 caracteres ASCII básicos (letras sin tilde, espacios y números) ocupan 1 byte cada uno, mientras que la `ó` ocupa 2 bytes, dando 21 bytes en total para 20 caracteres reales. Para el límite de subtítulos se debe usar el número de **runas** (`20`), porque la accesibilidad mide cuántos caracteres visibles lee el espectador en pantalla y no el peso en memoria del texto.

**b) En `for i, r := range texto`, ¿qué valen `i` y `r` en cada vuelta? Si `texto` es `ñu`, ¿qué valores toma `i`?**
En cada vuelta de `range` sobre una cadena, `r` es la runa (carácter Unicode) actual e `i` es el índice del byte donde empieza esa runa. Si `texto` es `"ñu"`, en la primera vuelta `r` vale `'ñ'` e `i` vale `0`; como `'ñ'` ocupa 2 bytes en UTF-8, en la segunda vuelta `r` vale `'u'` e `i` toma el valor `2`.

**c) ¿Por qué las palabras se cuentan con la variable `enPalabra` y no contando los espacios? Usa el caso 4 como ejemplo.**
Porque contar espacios falla cuando existen espacios repetidos o espacios al inicio y al final. En el caso 4 (`"  espacios   múltiples  "`), la cadena tiene 7 espacios (2 al comienzo, 3 en el medio y 2 al final), pero solo contiene 2 palabras. Con la bandera `enPalabra`, el programa únicamente suma `1` cuando pasa de un espacio a una runa que no es espacio, ignorando los espacios consecutivos.