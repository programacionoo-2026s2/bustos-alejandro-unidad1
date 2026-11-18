# 03-tarea4: Conversor de unidades organizado en paquetes

## Descripción y estructura
Programa de línea de comandos en Go que convierte valores de almacenamiento entre `B`, `KB`, `MB` y `GB` (en múltiplos de 1024), separando la lógica de negocio en un paquete interno reutilizable y dejando la lectura de argumentos y códigos de salida en el paquete principal.

Árbol de carpetas del proyecto:

    03-tarea4/
    ├── go.mod
    ├── main.go
    ├── README.md
    └── internal/
        └── almacenamiento/
            └── almacenamiento.go

- `go.mod`: Define el módulo `tarea4`, base para la ruta de importación `tarea4/internal/almacenamiento`.
- `main.go`: Valida la cantidad de argumentos (`os.Args`), convierte el número inicial (`strconv.ParseFloat`), invoca al paquete `almacenamiento` e imprime resultados o errores en `stderr`.
- `internal/almacenamiento/almacenamiento.go`: Define las constantes exportadas (`B`, `KB`, `MB`, `GB`), la validación de unidades (`Bytes`), el cálculo (`Convertir`) y la función privada `redondear`.

## Cómo ejecutar
Desde la carpeta `03-tarea4`, compilar y ejecutar pasando el valor, la unidad de origen y la unidad de destino:

    go build -o convertir.exe .
    .\convertir.exe 2048 KB MB

Significado de los códigos de salida:
- `0`: Conversión realizada con éxito.
- `1` (`ErrUso`): No se recibieron exactamente 3 argumentos.
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

## Experimento del compilador
Al reemplazar temporalmente `almacenamiento.Convertir(...)` por `almacenamiento.redondear(valor)` en `main.go` y ejecutar `go build ./...`, el compilador detuvo la compilación con el siguiente mensaje:

    # tarea4
    .\main.go:40:16: undefined: almacenamiento.redondear

## Preguntas

**a) ¿Por qué `Bytes` y `Convertir` empiezan con mayúscula y `redondear` con minúscula? ¿Qué ganas al dejar `redondear` oculta?**
En Go la visibilidad se controla con la inicial del identificador: `Bytes` y `Convertir` empiezan con mayúscula para ser exportadas y poder usarse desde `main.go` mediante `almacenamiento.Bytes` y `almacenamiento.Convertir`. En cambio, `redondear` empieza con minúscula porque es un detalle interno de cálculo que solo necesita `Convertir`. Al mantenerla oculta se protege la encapsulación del paquete: ningún otro archivo puede llamarla por error ni depender de ella si en el futuro cambia la forma en que se redondean los decimales.

**b) ¿Qué mensaje dio el compilador al usar `almacenamiento.redondear` desde `main.go` y qué regla de Go lo explica?**
El compilador mostró `undefined: almacenamiento.redondear`. Esto ocurre porque en Go todo identificador cuyo nombre inicia con letra minúscula tiene alcance privado a su propio paquete (`package almacenamiento`), por lo que resulta invisible e inaccesible para cualquier otro paquete externo como `main`.

**c) ¿Para qué sirve que el paquete esté dentro de `internal/` y qué pasaría si otro módulo intentara importarlo?**
El directorio `internal/` es una regla especial de Go que restringe quién puede importar los paquetes alojados dentro de él: solo el código perteneciente al mismo árbol de módulo (`tarea4`) tiene permiso para importar `"tarea4/internal/almacenamiento"`. Si otro módulo externo intentara importarlo, el compilador de Go rechazaría la compilación indicando que el uso de un paquete `internal` no está permitido desde fuera del módulo.