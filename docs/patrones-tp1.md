Patrones de Diseño utilizados

Nombre: Observer

Intención: Notificar automáticamente a uno o más objetoss cuando el estado del otro cambia.

Problemas que resuelve: Cuando el stock_actual de un Producto baja del stock_mínimo, el sistema debe alertar. Sin Observer, habría que verificar el stock manualmente en cada operación. 

Justificación: La base de datos ya tiene stock_minimo y punto_reorden en tbl_productos. El Obsrever traduce esa lógica a código orientado a objetos desaccopladdo.

Ejemplo en el código: Producto como Subjet -> AlertaStockObserver como observer concreto.


Nombre: State
Intención: Permitir que un objeto cambie su comportamiento cuando su estado interno cambia.

Problema que resuelve: tblestado tiene 5 estados (Activo, Inactivo, Pendiente, Finalizado, Cancelado). 
Sin State, el código sería un if/else gigante para controlar qué operaciones son válidas en cada estado.

Justificación: Con State, cada estado es una clase. Un Producto Inactivo no puede registrar ventas. Una 
Entrada Cancelada no puede modificar stock.

Ejemplo en el código: ProductoActivo, ProductoInactivo implementando la interfaz EstadoProducto.
