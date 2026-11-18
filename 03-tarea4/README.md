# 03-tarea4: Conversor de unidades organizado en paquetes

## Descripción y estructura
Proyecto dividido en el punto de entrada `main.go` y el paquete interno `internal/almacenamiento`, encargado de las constantes de bytes, validación de unidades y conversión con redondeo.
## Cómo ejecutar
Desde `03-tarea4`, compilar y ejecutar con `go build -o convertir.exe .` y `.\convertir.exe <valor> <origen> <destino>`.
- [03-tarea4](./03-tarea4) - Conversor de unidades modularizado con paquete `internal/almacenamiento` (Paso 3: lógica movida al paquete).