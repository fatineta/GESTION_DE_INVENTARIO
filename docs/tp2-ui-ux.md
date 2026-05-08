## A2. Análisis de usuario, tarea y contexto

### ¿Quiénes son los usuarios objetivo del sistema?

El sistema de control de stock está diseñado para ser utilizado por trabajadores de ferreterías pequeñas o medianas con una única sucursal física.  

Los perfiles de usuario contemplados son dos:

- **Encargado de ventas**: utiliza el sistema principalmente para consultar el stock disponible en tiempo real, lo que le permite responder con precisión las consultas de los clientes y confirmar la disponibilidad de productos durante la atención en mostrador.  
- **Encargado de compras**: utiliza el módulo de reposición para identificar qué productos necesitan ser pedidos a los proveedores. Este módulo presenta un sistema de **alarmas clasificadas por urgencia**, que se configura al momento de dar de alta cada producto definiendo su **stock mínimo aceptable**.

---

### ¿Qué tareas principales realizan y en qué contexto (dispositivo, entorno, restricciones)?

En cuanto al entorno de uso, el sistema opera en un contexto de **atención al público**, donde la velocidad de respuesta es importante.

- El **encargado de ventas** trabaja bajo presión, por lo que el sistema debe permitirle obtener información de stock en pocos segundos y sin pasos innecesarios.  
- El **encargado de compras** trabaja en un contexto más administrativo, revisando el estado del inventario de forma periódica para planificar pedidos.

En establecimientos pequeños, ambas funciones pueden recaer sobre la misma persona, por lo que el sistema fue diseñado para que ambos flujos sean accesibles desde una **única interfaz**, sin necesidad de configuraciones adicionales.

---

### Dispositivos y requisitos técnicos

- Se requiere al menos una **computadora de escritorio o notebook** con acceso a un navegador web.  
- Si las tareas están divididas entre personas distintas, se recomienda contar con **dos equipos independientes** para evitar interrupciones.  
- Se sugiere el uso de un **lector de código de barras compatible con entrada por teclado**, lo que permite agilizar el registro de productos sin ingresar códigos manualmente.

No se requiere formación técnica avanzada, ya que la interfaz fue diseñada siguiendo criterios de **usabilidad**, orientados a usuarios sin experiencia en software de gestión.

# A3. Auditoría de Usabilidad según ISO 9241-11

## Criterio 1: Eficacia

### Definición en contexto
La eficacia mide si el usuario puede completar la tarea correctamente, sin errores y sin ayuda externa. En el sistema de ferretería, la tarea crítica es registrar una entrada o salida de stock. Un empleado que comete un error en esa operación (producto equivocado, cantidad incorrecta, movimiento duplicado) genera un dato falso en el inventario que puede no detectarse hasta el próximo recuento físico.

### Métrica definida
Porcentaje de usuarios que completan el flujo de registro de movimiento de stock sin cometer ningún error en el primer intento, sin consultar ayuda y sin abandonar el formulario a mitad de camino.

**Meta:** que el 80% de los usuarios nuevos logren completar la tarea correctamente en su primera interacción con el sistema.

### Simulación en el prototipo actual
Para evaluar esta métrica sobre el prototipo Figma, se le pide a una persona que no participó del desarrollo que complete el siguiente flujo sin explicaciones previas:

1. Ingresar al sistema.
2. Encontrar el producto **"Tornillo 6mm"**.
3. Registrar una salida de 5 unidades.
4. Confirmar la operación.

Se observa sin intervenir y se registra si logra completarlo sin errores.

#### Posibles puntos de fricción
- Ambigüedad entre **"registrar entrada"** y **"registrar salida"** si ambas opciones están en el mismo botón o menú.
- Ausencia de confirmación explícita antes de guardar el movimiento, lo que puede llevar a envíos accidentales.

### Mejora concreta identificada
- Separar visualmente las acciones de entrada y salida de stock desde el primer nivel de navegación.
- Utilizar:
  - Flecha hacia arriba para entrada.
  - Flecha hacia abajo para salida.
- Aplicar colores diferenciados:
  - Verde para entrada.
  - Rojo para salida.

Agregar un paso de confirmación antes de guardar que muestre un resumen:

> "Vas a registrar una SALIDA de 5 unidades de Tornillo 6mm. Stock actual: 10 → Stock resultante: 5. ¿Confirmar?"

Este paso adicional reduce los errores de envío accidental sin agregar fricción innecesaria al flujo.

Además, modificar el selector de productos para mostrar junto al nombre el stock actual y la unidad de medida en texto secundario, por ejemplo:

> "Tornillo 6×50mm · Stock actual: 10 u."

Esto elimina la ambigüedad cuando existen productos con nombres similares en el catálogo.

---

## Criterio 2: Eficiencia

### Definición en contexto
La eficiencia mide cuántos recursos consume el usuario para completar la tarea correctamente: tiempo, cantidad de pasos, clics y carga cognitiva.

En una ferretería, el empleado registra movimientos de stock mientras atiende clientes en el mostrador. Cada segundo adicional que tarda en completar la operación es tiempo que le quita a la atención al cliente.

Un sistema ineficiente no se abandona de inmediato, pero genera resistencia acumulada hasta que el empleado busca un atajo, generalmente omitir el registro.

### Métrica definida
Cantidad de clics necesarios para completar el flujo completo de registro de una salida de stock, desde la pantalla de inicio hasta la confirmación.

**Meta:** que ningún flujo principal supere los 5 clics.

Esta métrica es objetiva, fácil de contar en el prototipo y directamente mejorable a través del diseño.

### Simulación en el prototipo actual
Se recorre el flujo completo en el prototipo Figma contando cada clic necesario.

#### Flujo típico actual
1. Clic en **"Movimientos"** en el menú.
2. Clic en **"Nueva salida"**.
3. Seleccionar el producto del listado.
4. Ingresar la cantidad.
5. Clic en **"Guardar"**.
6. Clic en **"Confirmar"** en el popup de confirmación.

El flujo tiene **6 pasos**, uno por encima de la meta.

Si además el empleado necesita buscar el producto porque el listado no está ordenado o no tiene buscador rápido, pueden sumarse entre 2 y 3 clics adicionales.

### Mejora concreta identificada
Agregar un buscador de producto con autocompletado directamente en el formulario de nuevo movimiento.

De esta manera, el empleado puede escribir las primeras letras del nombre y seleccionar el producto sin tener que navegar a una pantalla de listado separada.

#### Beneficios
- Elimina al menos 2 clics del flujo.
- Concentra toda la operación en una sola pantalla.
- Reduce el tiempo de búsqueda.

Además, mostrar los últimos 5 productos movidos como accesos rápidos disminuye aún más el tiempo de operación para productos de alta rotación.

---

# Alineación con ISO 13407: Diseño Centrado en el Usuario

ISO 13407 (actualizada como ISO 9241-210) define un proceso iterativo de cuatro fases para diseñar sistemas centrados en las personas usuarias.

El proceso seguido por el equipo durante el proyecto se alinea con esas cuatro fases.

## 1. Entender y especificar el contexto de uso
Esta fase ocurrió en el Sprint 0 cuando el equipo eligió el escenario de ferretería e identificó:

- Empleado de mostrador.
- Encargado de compras.

como usuarios principales.

Esto definió restricciones de diseño como:

- Usuario sin formación técnica.
- Entorno de alta presión.
- Uso en computadora de escritorio dentro del local.

## 2. Especificar requisitos del usuario y de la organización
Se materializó en:

- El diagrama de casos de uso del TP1.
- Las tarjetas del backlog.

Cada funcionalidad mínima de la consigna:

- Registrar productos.
- Alertar stock bajo.
- Buscar por categoría.

representa un requisito de usuario que guió las decisiones de diseño.

## 3. Producir soluciones de diseño
Esta fase corresponde al desarrollo del prototipo en Figma.

Las pantallas navegables son la traducción de los requisitos en interfaces concretas que el usuario puede utilizar.

## 4. Evaluar los diseños contra los requisitos
Esta auditoría de usabilidad corresponde a la fase de evaluación.

Las métricas de eficacia y eficiencia definidas funcionan como criterios de evaluación.

Si el prototipo no cumple esos criterios, se vuelve a la etapa de diseño aplicando mejoras.

---

## Conclusión
La iteración constante es la esencia de ISO 13407 y es lo que diferencia un diseño centrado en el usuario de simplemente crear pantallas sin validación.

Los problemas detectados durante la auditoría no representan fracasos, sino insumos para el ciclo de mejora continua que alimenta los siguientes sprints del proyecto.
