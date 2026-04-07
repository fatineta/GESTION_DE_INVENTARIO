## 3.2. Patrón N°2: State

### 3.2.1. Intención Original según GoF

El patrón **State** es uno de los patrones de comportamiento definidos por la *Gang of Four (GoF)*.  
Su objetivo principal es permitir que un objeto altere su comportamiento cuando su **estado interno cambia**, dando la impresión de que el objeto ha cambiado de clase.

Este patrón resuelve problemas donde el comportamiento depende fuertemente del estado del objeto y normalmente se implementaría mediante estructuras complejas de `if-else` o `switch-case`.

Cada estado posible se encapsula en una **clase concreta** que implementa una **interfaz común**.  
El objeto principal delega su comportamiento al estado actual en lugar de contener toda la lógica condicional internamente.

De esta manera:

- Se eliminan estructuras condicionales extensas.
- Cada estado posee su propia lógica.
- El sistema se vuelve más extensible y mantenible.

---

### 3.2.2. Problema de Diseño en el Sistema

En el sistema de gestión de inventario, la tabla `tblestado` define cinco estados posibles:

- Activo  
- Inactivo  
- Pendiente  
- Finalizado  
- Cancelado  

Estos estados se aplican a:

- Productos (`tblproductos.estado_id`)
- Movimientos de entrada (`tblentradas.estado_id`)
- Ventas (`tblventas.estado_id`)

El problema surge al controlar **qué operaciones son válidas según el estado actual** de cada objeto.

Ejemplos:

- Un **Producto** en estado *Inactivo* no debería permitir registrar nuevas ventas ni entradas de stock.
- Una **Entrada** en estado *Cancelado* no debería modificar el `stock_actual` del producto asociado.
- Una **Entrada** en estado *Pendiente* debería poder confirmarse o cancelarse, pero no finalizarse directamente.

Sin aplicar el patrón State, toda esta lógica quedaría centralizada dentro de métodos de `Producto`, `Entrada` o `Venta`, utilizando múltiples validaciones condicionales sobre `estado_id`.

Esto genera varios problemas:

- ❌ Violación del **Principio de Responsabilidad Única (SRP)**  
- ❌ Código difícil de leer y mantener  
- ❌ Alta dependencia entre módulos  
- ❌ Dificultad para realizar testing aislado por estado  

Además, esta problemática se refleja directamente en la interfaz del sistema:

- El flujo **v-mov** de tres pasos implica transiciones de estado  
  - Paso 1 → Paso 2 → Paso 3 → Confirmado
- El **Dashboard** muestra badges de estado (*crítico*, *bajo*, *normal*), representando distintos comportamientos del mismo objeto `Producto`.

Esto evidencia que el comportamiento del sistema depende del estado, haciendo necesario un patrón orientado a estados.

---

### 3.2.3. Justificación Técnica y Alternativas Descartadas (Dimensión 3)

#### ¿Por qué State y no otra solución?

**Alternativa 1: uso de `estado_id` con if-else o switch-case**

Se evaluó mantener el campo `estado_id` como un entero y controlar las reglas mediante estructuras condicionales.

Fue descartada porque:

- A medida que aumentan los estados, el código se vuelve inmantenible.
- Cada nueva regla obliga a modificar múltiples clases.
- Viola el **Principio Abierto/Cerrado (OCP)**.

---

**Alternativa 2: tabla de transiciones en base de datos**

Otra opción fue definir en la base de datos una tabla de transiciones válidas y consultarla antes de cada operación.

Si bien resulta útil como validación en la capa de persistencia:

- No organiza la lógica en la capa de objetos.
- El comportamiento dependiente del estado sigue disperso.
- No mejora la estructura del código orientado a objetos.

---

#### Solución elegida: Patrón State

El patrón **State** fue seleccionado porque:

- Encapsula el comportamiento de cada estado en una clase independiente.
- Cada estado conoce:
  - Qué operaciones permite.
  - A qué estados puede transicionar.
- Permite agregar nuevos estados (ej.: *En revisión*) sin modificar clases existentes.
- Facilita el testing aislado de cada estado.
- Se alinea naturalmente con una estrategia **TDD**.

Además, la decisión es coherente con los casos de uso del sistema:

- *Verificar stock mínimo*
- *Señal de alarma*
- Indicadores de stock **normal**, **bajo** o **crítico**

En todos estos escenarios, un mismo objeto `Producto` modifica su comportamiento según su estado, exactamente el problema que el patrón **State** está diseñado para resolver.
