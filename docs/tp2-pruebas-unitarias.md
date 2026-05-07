# TP2 — Pruebas Unitarias con TDD

## Información general

| Campo | Detalle |
|---|---|
| Proyecto | Sistema de Gestión de Ferretería |
| Lenguaje | Go 1.22 |
| Archivo de pruebas | `tests/unit/ferreteria_test.go` |
| Total de casos | 21 |
| Funciones cubiertas | 5 (`calcularNuevoStock`, `calcularUrgencia`, `validarProducto`, `estadoDesdeNombre`, `calcularFaltante`) + Patrón State (`EstadoPedido`) |

---

## Estrategia de pruebas

Las pruebas se organizan con **TDD** (Test-Driven Development):
la lógica de negocio pura fue extraída del `main.go` original en funciones
independientes (sin dependencias de base de datos), permitiendo verificarlas
de forma aislada.

Se aplican dos técnicas:

- **Partición de equivalencia**: se agrupa el dominio de entrada en clases
  cuyos valores se comportan de igual forma. Se elige un representante de
  cada clase.
- **Análisis de valores límite**: se prueban los extremos de cada clase
  (mínimo, máximo, justo debajo/arriba del umbral).

---

## Casos de prueba

### Función: `calcularNuevoStock(stockActual, cantidad int, tipo string) (int, error)`

Reproduce la lógica del `handleMovimientoStock` del sistema original.

---

#### TC-01 — ENTRADA válida

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | `stockActual=10`, `cantidad=5`, `tipo="ENTRADA"` |
| **Resultado esperado** | `nuevoStock=15`, `error=nil` |
| **Nombre del test** | `TestCalcularNuevoStock_EntradaValida` |

Clase de equivalencia válida: `tipo = "ENTRADA"` con cantidad positiva suma
correctamente al stock.

---

#### TC-02 — SALIDA exactamente igual al stock disponible

| Campo | Detalle |
|---|---|
| **Técnica** | Valor límite |
| **Datos de entrada** | `stockActual=5`, `cantidad=5`, `tipo="SALIDA"` |
| **Resultado esperado** | `nuevoStock=0`, `error=nil` |
| **Nombre del test** | `TestCalcularNuevoStock_SalidaExactaAlStock` |

Límite inferior permitido: la cantidad retirada es exactamente igual al stock.
El resultado debe ser 0 sin error.

---

#### TC-03 — SALIDA con cantidad = stock + 1 (supera el límite)

| Campo | Detalle |
|---|---|
| **Técnica** | Valor límite |
| **Datos de entrada** | `stockActual=5`, `cantidad=6`, `tipo="SALIDA"` |
| **Resultado esperado** | Error `"stock insuficiente"` |
| **Nombre del test** | `TestCalcularNuevoStock_SalidaExcedeStock` |

Un valor justo por encima del límite permitido debe producir error.

---

#### TC-04 — AJUSTE reemplaza el stock

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | `stockActual=100`, `cantidad=30`, `tipo="AJUSTE"` |
| **Resultado esperado** | `nuevoStock=30`, `error=nil` |
| **Nombre del test** | `TestCalcularNuevoStock_AjusteValido` |

Clase válida: cualquier valor de ajuste reemplaza directamente el stock
independientemente del valor anterior.

---

#### TC-05 — Tipo de movimiento inválido

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | `stockActual=10`, `cantidad=5`, `tipo="DEVOLUCION"` |
| **Resultado esperado** | Error `"tipo inválido"` |
| **Nombre del test** | `TestCalcularNuevoStock_TipoInvalido` |

Clase inválida: cualquier tipo distinto de `ENTRADA`, `SALIDA` o `AJUSTE` es rechazado.

---

### Función: `calcularUrgencia(stockActual, stockMinimo int) string`

Reproduce la clasificación de urgencia del `handleReposicion`.

---

#### TC-06 — Stock igual a 0 → urgencia CRITICO

| Campo | Detalle |
|---|---|
| **Técnica** | Valor límite |
| **Datos de entrada** | `stockActual=0`, `stockMinimo=10` |
| **Resultado esperado** | `"CRITICO"` |
| **Nombre del test** | `TestCalcularUrgencia_StockCero` |

El valor 0 es el límite absoluto inferior: representa ausencia total de stock,
caso más crítico posible.

---

#### TC-07 — Stock = stockMinimo/2 (límite entre ALTA y MEDIA)

| Campo | Detalle |
|---|---|
| **Técnica** | Valor límite |
| **Datos de entrada** | `stockActual=5`, `stockMinimo=10` |
| **Resultado esperado** | `"ALTA"` |
| **Nombre del test** | `TestCalcularUrgencia_StockEnLimiteAlta` |

El umbral exacto `stock <= stockMinimo/2` clasifica como `ALTA`.
Se prueba el valor en el límite.

---

#### TC-08 — Stock = stockMinimo/2 + 1 (pasa a MEDIA)

| Campo | Detalle |
|---|---|
| **Técnica** | Valor límite |
| **Datos de entrada** | `stockActual=6`, `stockMinimo=10` |
| **Resultado esperado** | `"MEDIA"` |
| **Nombre del test** | `TestCalcularUrgencia_StockSobreLimiteAlta` |

Un valor justo por encima del umbral abandona la clase `ALTA` y cae en `MEDIA`.

---

### Función: `validarProducto(nombre string, precioVenta float64) error`

Reproduce la validación de campos obligatorios del `handleCrearProducto`.

---

#### TC-09 — Producto con datos válidos

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | `nombre="Martillo"`, `precioVenta=150.0` |
| **Resultado esperado** | `error=nil` |
| **Nombre del test** | `TestValidarProducto_DatosValidos` |

Clase válida: nombre no vacío y precio positivo deben pasar sin error.

---

#### TC-10 — Nombre vacío

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | `nombre=""`, `precioVenta=150.0` |
| **Resultado esperado** | Error |
| **Nombre del test** | `TestValidarProducto_NombreVacio` |

Clase inválida: nombre vacío es campo obligatorio faltante.

---

#### TC-11 — Precio exactamente 0 (límite inferior inválido)

| Campo | Detalle |
|---|---|
| **Técnica** | Valor límite |
| **Datos de entrada** | `nombre="Llave"`, `precioVenta=0` |
| **Resultado esperado** | Error |
| **Nombre del test** | `TestValidarProducto_PrecioCero` |

El precio 0 es el límite que divide válido (>0) de inválido (≤0).

---

#### TC-12 — Precio negativo

| Campo | Detalle |
|---|---|
| **Técnica** | Valor límite |
| **Datos de entrada** | `nombre="Llave"`, `precioVenta=-1.0` |
| **Resultado esperado** | Error |
| **Nombre del test** | `TestValidarProducto_PrecioNegativo` |

Valor por debajo del límite 0: confirma que la clase inválida incluye negativos.

---

### Patrón State — `EstadoPedido`

Las transiciones del patrón State se prueban de forma pura, sin BD.

---

#### TC-13 — PendienteState: cancelar es una transición válida

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | Pedido en estado `PENDIENTE`, acción `Cancelar` |
| **Resultado esperado** | Estado cambia a `"CANCELADO"`, `error=nil` |
| **Nombre del test** | `TestPendienteState_Cancelar` |

---

#### TC-14 — PendienteState: pagar sin procesar es inválido

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | Pedido en estado `PENDIENTE`, acción `Pagar` |
| **Resultado esperado** | Error (transición no permitida) |
| **Nombre del test** | `TestPendienteState_PagarSinProcesar` |

---

#### TC-15 — ConfirmadoState: pagar avanza a PAGADO

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | Pedido en estado `CONFIRMADO`, acción `Pagar` |
| **Resultado esperado** | Estado cambia a `"PAGADO"`, `error=nil` |
| **Nombre del test** | `TestConfirmadoState_Pagar` |

---

#### TC-16 — EntregadoState: no puede cancelarse

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | Pedido en estado `ENTREGADO`, acción `Cancelar` |
| **Resultado esperado** | Error `"no se puede cancelar"` |
| **Nombre del test** | `TestEntregadoState_NoPuedeCancelarse` |

---

#### TC-17 — EnAlmacenState: entregar avanza a ENTREGADO

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | Pedido en estado `EN_ALMACEN`, acción `Entregar` |
| **Resultado esperado** | Estado cambia a `"ENTREGADO"`, `error=nil` |
| **Nombre del test** | `TestEnAlmacenState_Entregar` |

---

### Función: `estadoDesdeNombre(nombre string) EstadoPedido`

---

#### TC-18 — Nombre desconocido devuelve PENDIENTE por defecto

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | `nombre="INEXISTENTE"` |
| **Resultado esperado** | Estado con `GetNombre() == "PENDIENTE"` |
| **Nombre del test** | `TestEstadoDesdeNombre_Default` |

Clase inválida: cualquier nombre no reconocido cae al estado por defecto.

---

#### TC-19 — Nombre "EN_ALMACEN" resuelve correctamente

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | `nombre="EN_ALMACEN"` |
| **Resultado esperado** | Estado con `GetNombre() == "EN_ALMACEN"` |
| **Nombre del test** | `TestEstadoDesdeNombre_EnAlmacen` |

---

### Función: `calcularFaltante(stockActual, stockMinimo, stockMaximo int) int`

---

#### TC-20 — stockMaximo < stockActual → faltante = stockMinimo * 2

| Campo | Detalle |
|---|---|
| **Técnica** | Valor límite |
| **Datos de entrada** | `stockActual=80`, `stockMinimo=10`, `stockMaximo=50` |
| **Resultado esperado** | `20` (= 10 × 2) |
| **Nombre del test** | `TestCalcularFaltante_StockMaximoMenorQueActual` |

Límite especial: cuando la diferencia es negativa se aplica la fórmula alternativa.

---

#### TC-21 — Caso normal: stockMaximo > stockActual

| Campo | Detalle |
|---|---|
| **Técnica** | Partición de equivalencia |
| **Datos de entrada** | `stockActual=10`, `stockMinimo=5`, `stockMaximo=100` |
| **Resultado esperado** | `90` |
| **Nombre del test** | `TestCalcularFaltante_CasoNormal` |

---

## Resumen de cobertura

| Función / Método | Casos | Técnicas usadas |
|---|---|---|
| `calcularNuevoStock` | TC-01 a TC-05 | Equivalencia + Valor límite |
| `calcularUrgencia` | TC-06 a TC-08 | Valor límite |
| `validarProducto` | TC-09 a TC-12 | Equivalencia + Valor límite |
| Patrón State `EstadoPedido` | TC-13 a TC-17 | Equivalencia |
| `estadoDesdeNombre` | TC-18 a TC-19 | Equivalencia |
| `calcularFaltante` | TC-20 a TC-21 | Equivalencia + Valor límite |
| **Total** | **21** | |

---

## Ejecución

```bash
cd tests/unit
go test ./... -v
```

Salida esperada: todos los tests con estado `PASS`.

```
--- PASS: TestCalcularNuevoStock_EntradaValida (0.00s)
--- PASS: TestCalcularNuevoStock_SalidaExactaAlStock (0.00s)
...
PASS
ok      ferreteria_tests    0.002s
```
