# Bitácora de Desarrollo - Unidad 1 (IFT 365)

Registro cronológico de avances, dificultades técnicas y soluciones aplicadas en el repositorio `bustos-alejandro-unidad1`.

---

## 01-taller1 · Configuración inicial y variables (`estudiante.go`)
- **Actividad:** Inicialización del primer módulo en Go (`go mod init`), declaración de variables con datos personales e impresión formateada en consola.
- **Resultado:** Ejecución correcta del programa base y sincronización inicial con el repositorio remoto en GitHub.

## 02-taller1 · Función `init()` y alcance de variables (`ejercicio1.go`)
- **Actividad:** Implementación de bloques `init()` previos a `main()` para inicializar estado y verificar el orden de ejecución en Go.
- **Error encontrado y solución:** La primera versión no respetaba la estructura esperada para la inicialización de variables. Se corrigió el orden de declaración y asignación dentro de `init()` antes de realizar el commit de actualización.

## 03-taller1 y 04-taller1 · Tipos de datos, cadenas y control de flujo
- **Actividad:** Resolución de ejercicios prácticos en Go sobre manejo de cadenas, conteo de runas con `unicode/utf8`, división de palabras con `strings.Fields` y estructuras condicionales.
- **Resultado:** Programas probados en consola y documentados en sus respectivos archivos `README.md`.

## 03-tarea1 · Ámbito de variables y sombreado (`suma.go`)
- **Actividad:** Análisis de variable shadowing y verificación estática con herramientas de diagnóstico de Go.
- **Error encontrado y solución:** Al intentar revisar el sombreado de variables en PowerShell, se escribió `shadow` como un comando independiente tras `go vet`, generando un error de cmdlet no reconocido[cite: 1]. Además, al copiar los comandos de Git se incluyeron comillas invertidas extra que bloquearon la terminal[cite: 1]. Se solucionó cancelando la entrada con `Ctrl + C` y ejecutando `git add .`, `git commit` y `git push` sin caracteres adicionales.

## 03-tarea2 · Permisos estilo UNIX con `iota` y operadores de bits (`permisos.go`)
- **Actividad:** Definición del tipo `Permiso uint8` y constantes desplazadas con `1 << iota` ordenadas según `chmod` (`Ejecutar = 1`, `Escribir = 2`, `Leer = 4`), implementando la función `describir` con `&` y modificando permisos mediante `|=` y `&^=`.
- **Error encontrado y solución:** Al pegar comandos desde bloques de código en la terminal de VS Code se copió el encabezado `bash` junto con comillas invertidas, provocando errores de sintaxis en PowerShell[cite: 3]. Se resolvió limpiando la línea de comandos y ejecutando únicamente los comandos puros de Git.

## 03-tarea3 · Conversor de unidades de almacenamiento (`main.go`)
- **Actividad:** Construcción incremental en 4 commits de una herramienta CLI que convierte entre `B`, `KB`, `MB` y `GB` (múltiplos de 1024), valida argumentos con `os.Args` y `strconv.ParseFloat`, normaliza unidades con `strings.ToUpper` y reporta errores por `os.Stderr` usando constantes de salida (`ErrUso = 1`, `ErrValor = 2`, `ErrUnidad = 3`).
- **Errores encontrados y soluciones:**
  1. **Código pegado en consola:** Durante la creación inicial de `main.go`, se pegó el código fuente de Go directamente en PowerShell en lugar del editor de archivos, generando un `ParserError`[cite: 5]. Se corrigió creando y editando `main.go` desde el Explorador de VS Code.
  2. **Ejecutable en Windows:** Al compilar con `go build -o convertir .`, Windows no reconocía el binario sin extensión al invocarlo desde PowerShell. Se solucionó compilando con `go build -o convertir.exe .` para ejecutar los 8 casos de prueba, verificar `$LASTEXITCODE` y eliminar posteriormente el `.exe` antes de subir los cambios al repositorio.
  ## 03-tarea4 · Modularización con paquetes e `internal/`
- **Avance 1 (Monolito):** Creación de `03-tarea4`, inicialización del módulo `tarea4` y prueba de referencia en `main.go`. Resultados obtenidos en los 8 casos: `2.00 MB` (0), `1024.00 MB` (0), `1.46 KB` (0), `512.00 KB` (0), error de valor (2), error negativo (2), unidad `TB` inválida (3) y argumentos incompletos (1).
- **Avance 2 (Paquete interno):** Creación de `internal/almacenamiento/almacenamiento.go` con las constantes exportadas (`B`, `KB`, `MB`, `GB`), funciones exportadas (`Bytes`, `Convertir`) y la función privada `redondear`. Verificado con `go doc ./internal/almacenamiento`.