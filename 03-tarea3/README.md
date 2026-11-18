# 03-tarea3: Conversor de unidades de almacenamiento

## Descripción
Programa de consola escrito en Go que convierte cantidades de almacenamiento digital entre bytes (`B`), kilobytes (`KB`), megabytes (`MB`) y gigabytes (`GB`) usando múltiplos binarios de 1024. Valida los argumentos recibidos por línea de comandos y termina con códigos de salida específicos según el tipo de error.

## Cómo ejecutar
Desde la carpeta `03-tarea3`, compilar y ejecutar pasando el valor, la unidad de origen y la unidad de destino:

    go build -o convertir.exe .
    .\convertir.exe 2048 KB MB

Significado de los códigos de salida:
- `0`: Conversión realizada correctamente.
- `1` (`ErrUso`): No se proporcionaron exactamente 3 argumentos.
- `2` (`ErrValor`): El primer argumento no es un número válido o es negativo.
- `3` (`ErrUnidad`): La unidad de origen o destino no es `B`, `KB`, `MB` o `GB`.

## Pruebas realizadas

| Comando | Salida obtenida | Código |
| :--- | :--- | :---: |
| `.\convertir.exe 2048 KB MB` | `2048 KB = 2.00 MB` | `0` |
| `.\convertir.exe 1 gb mb` | `1 GB = 1024.00 MB` | `0` |
| `.\convertir.exe 1500 B KB` | `1500 B = 1.46 KB` | `0` |
| `.\convertir.exe 0.5 MB KB` | `0.5 MB = 512.00 KB` | `0` |
| `.\convertir.exe abc KB MB` | `Error: el valor debe ser un numero valido mayor o igual a 0` | `2` |
| `.\convertir.exe -5 KB MB` | `Error: el valor debe ser un numero valido mayor o igual a 0` | `2` |
| `.\convertir.exe 5 TB MB` | `Error: unidad no reconocida (TB)` | `3` |
| `.\convertir.exe 5 KB` | `Uso: convertir <valor> <unidad_origen> <unidad_destino>` | `1` |

## Preguntas

**a) ¿Por qué el valor está en `os.Args[1]` y no en `os.Args[0]`?**
Porque en Go el slice `os.Args` guarda en la posición `0` el nombre o ruta del propio programa ejecutable (por ejemplo `.\convertir.exe`). Los argumentos que uno escribe en la terminal después del comando comienzan en la posición `1`, por lo que `os.Args[1]` contiene el valor numérico, `os.Args[2]` la unidad de origen y `os.Args[3]` la unidad de destino.

**b) ¿Por qué los mensajes de error se escriben en `stderr` y no con `fmt.Println`?**
Porque `fmt.Println` escribe en la salida estándar (`stdout`), reservada únicamente para el resultado válido del programa. Si otro script captura la salida de este comando, un texto de error en `stdout` contaminaría los datos. Al usar `fmt.Fprintln(os.Stderr, ...)`, los mensajes de error salen por un canal separado sin mezclarse con la salida normal.

**c) ¿Qué pasaría en tu programa si escribieras `valor, _ := strconv.ParseFloat(...)` e ignoraras el error al ejecutar `convertir abc KB MB`?**
Si se ignora el error con `_`, `strconv.ParseFloat` falla al leer `"abc"` y devuelve `0` como valor por defecto. Como `0` no es menor que cero (`valor < 0` resulta falso), el programa no detectaría el error, seguiría ejecutándose e imprimiría `0 KB = 0.00 MB` con código de salida `0` en vez de terminar con código `2`.

**d) ¿Qué ventaja tiene declarar los códigos de salida con `const` en lugar de escribir `1`, `2` y `3` dentro de `os.Exit`?**
Escribir `os.Exit(ErrValor)` o `os.Exit(ErrUnidad)` hace que el código sea más claro y fácil de leer, ya que el nombre de la constante explica el motivo del fallo sin tener que adivinar qué significa cada número. Además, si después se necesita cambiar algún código de salida, solo se modifica en un único lugar dentro del bloque `const`.