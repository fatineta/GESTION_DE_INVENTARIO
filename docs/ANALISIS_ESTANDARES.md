# Normativas Aplicadas

## ISO 9241-11: Usabilidad

Esta norma define qué significa que un sistema sea usable. La definición oficial tiene tres dimensiones que siempre van juntas:

La eficacia es si el usuario puede completar la tarea correctamente. No importa cuánto tardó ni cómo se sintió ¿lo logró o no? En este caso la pregunta sería: ¿puede un empleado registrar una salida de stock sin cometer errores?

La eficiencia es cuántos recursos consume completar esa tarea tiempo, clics, pasos mentales. Un sistema eficiente permite registrar un producto nuevo en menos de un minuto sin consultar un manual.

La satisfacción es la percepción subjetiva del usuario. ¿Le resultó cómodo? ¿Confía en lo que ve? En una ferretería, si el empleado siente que el sistema "le complica la vida", lo abandona y vuelve al Excel.

Lo más importante de esta norma es que define usabilidad en contexto de uso específico no en abstracto. Un sistema puede ser muy usable para un ingeniero de software y completamente inutilizable para un empleado de mostrador que nunca usó una app web. Por eso la norma exige identificar quién usa el sistema, para qué tarea y en qué entorno antes de evaluar si es usable o no.

---

## ISO 13407: Diseño Centrado en el Humano

Esta norma (hoy actualizada como ISO 9241-210) define un proceso de cuatro pasos para diseñar sistemas que realmente funcionen para las personas que los usan. Es un ciclo iterativo, no una secuencia lineal que se hace una sola vez.

Entender y especificar el contexto de uso. Quién usa el sistema, para qué, con qué equipo, en qué condiciones. En tu proyecto esto fue el Sprint 0: identificar que el usuario principal es un empleado de ferretería sin formación técnica, que trabaja bajo presión de atención al cliente.

Especificar los requisitos del usuario y de la organización. Qué necesita poder hacer el usuario (registrar productos, ver alertas, buscar por categoría) y qué necesita la organización (trazabilidad de movimientos, control de acceso por rol).

Producir soluciones de diseño. Los wireframes de Figma, los prototipos HTML, el diagrama de casos de uso. No es solo hacer pantallas lindas es traducir los requisitos en interfaces concretas.

Evaluar los diseños contra los requisitos. Mostrar el prototipo a usuarios reales y verificar si pueden completar las tareas. Si no pueden, volver al paso 1 o 2. Esta iteración es la esencia de la norma.

La razón por la que esta norma sigue siendo relevante en sistemas críticos es que los errores de interfaz en entornos de alta presión (un médico, un operador de planta) pueden tener consecuencias graves. El diseño centrado en el humano no es solo comodidad es prevención de errores.

---

## ISO/IEC 27001: Seguridad de la Información

Esta norma establece cómo gestionar la seguridad de la información en una organización. Su concepto central es la tríada CIA: Confidencialidad, Integridad y Disponibilidad.

Confidencialidad significa que la información solo es accesible para quien está autorizado. En tu sistema: un vendedor no debería poder ver los precios de costo ni modificar el stock mínimo eso es rol de administrador. La tabla tbltipo_usuario de tu base de datos es exactamente el punto de partida para implementar esto.

Integridad significa que la información no puede ser alterada sin dejar rastro. En tu sistema: la tabla tblmovimientos_stock con stock_anterior y stock_nuevo es una implementación directa de este principio. Nadie puede cambiar el stock de un producto sin que quede registrado quién lo hizo y cuándo.

Disponibilidad significa que el sistema esté accesible cuando se lo necesita. En una ferretería, si el sistema de stock cae en el momento de mayor venta, el negocio opera a ciegas.

La norma no dice qué controles técnicos usar exactamente dice que hay que identificar los riesgos de seguridad, evaluarlos y aplicar controles proporcionales. Para un sistema académico como el tuyo, las prácticas más relevantes son: hashing de contraseñas (nunca guardarlas en texto plano en la columna password de tblusuarios), control de acceso basado en roles, y registro de auditoría de todas las operaciones críticas.

---

## ISA/IEC 62443: Ciberseguridad Industrial

Esta familia de normas existe porque los sistemas industriales tienen un problema de seguridad muy específico: fueron diseñados para funcionar de forma aislada durante décadas, y cuando empezaron a conectarse a internet quedaron extremadamente expuestos.

Se aplica a los llamados Sistemas de Control Industrial (ICS): SCADA (supervisión y control de procesos a distancia), DCS (sistemas de control distribuido), PLC (controladores lógicos programables). Estos sistemas controlan cosas físicas como una válvula de agua, una turbina, una línea de ensamblaje donde un ataque cibernético puede causar daño físico real.

La norma define niveles de seguridad (SL 0 a SL 4) según la sofisticación del atacante que el sistema debe resistir, y zonas de seguridad para segmentar la red industrial de la red corporativa.

Para nuestro proyecto, esta norma no aplica en su contexto actual porque no hay sistemas de control físico. Sin embargo, si el sistema evolucionara para controlar sensores de temperatura en un depósito refrigerado, brazos robóticos de almacenamiento, o básculas automatizadas conectadas a la red, entonces sí entraría en el alcance de esta norma. Es importante entender su límite: es para sistemas donde el software controla hardware físico con consecuencias en el mundo real.

---

## ISO 9001: Calidad en Procesos

Esta es la norma de gestión de calidad más adoptada en el mundo. Su premisa central es que la calidad no es un atributo del producto final sino el resultado de un proceso bien definido y gestionado. Si el proceso es bueno y consistente, el producto será bueno y consistente.

Su herramienta conceptual central es el ciclo PDCA: Planificar (definir qué se va a hacer y cómo), Hacer (ejecutarlo), Verificar (medir si los resultados corresponden a lo planificado) y Actuar (corregir las desviaciones y mejorar el proceso).

En este caso, la ISO 9001 se manifiesta en: los criterios de aceptación de cada tarjeta del tablero (¿cómo sabemos que la tarea está realmente terminada?), los tests unitarios con TDD como evidencia de verificación, y el AI_LOG como registro de proceso que permite auditar cómo se tomaron las decisiones técnicas.
# Análisis de Estándares

## ¿Cuáles son los dos estándares más relevantes para nuestro proyecto?

Los dos estándares más relevantes para el sistema de gestión de inventario de la ferretería son ISO 9241-11 e ISO 9001.

La elección de ISO 9241-11 responde a una característica central del escenario elegido en el Sprint 2: el usuario principal no es un profesional de tecnología. Es un empleado de mostrador que trabaja bajo presión de atención al cliente, que probablemente nunca usó una aplicación web de gestión, y que si encuentra el sistema difícil o confuso simplemente lo abandona y vuelve al Excel o al papel. En ese contexto, la usabilidad no es un detalle estético sino una condición de supervivencia del sistema. ISO 9241-11 es el estándar que nos da el marco conceptual para pensar en eficacia (¿puede el empleado registrar una entrada de stock sin cometer errores?), eficiencia (¿puede hacerlo en el menor tiempo posible mientras hay clientes esperando?) y satisfacción (¿confía en lo que ve en la pantalla?). Sin estas tres dimensiones cubiertas, el sistema técnicamente más correcto del mundo fracasa en producción.

La elección de ISO 9001 responde a la naturaleza del proceso que estamos usando para construir el sistema. Scrum es, en esencia, una implementación del ciclo PDCA que ISO 9001 prescribe: planificamos en el sprint planning, ejecutamos durante el sprint, verificamos con pruebas unitarias y revisión de código, y actuamos en la retrospectiva para mejorar el proceso. Más concretamente, las decisiones de diseño que tomamos en el TP1 son directamente trazables a los principios de esta norma: los tests unitarios con TDD son la evidencia de verificación, los criterios de aceptación de cada tarjeta del tablero son la planificación medible, y el AI_LOG es el registro de proceso que permite auditar cómo se tomaron las decisiones técnicas. ISO 9001 nos da el lenguaje para justificar por qué un proceso bien definido produce un sistema de calidad, independientemente de su escala.

---

## Si el sistema fuera declarado crítico, ¿qué estándares debería cumplir obligatoriamente?

Para responder esta pregunta conviene pensar en qué hace que un sistema sea declarado crítico: es crítico cuando un fallo, un error de datos o una brecha de seguridad produce consecuencias que van más allá del negocio y afectan la vida de personas, la infraestructura pública o activos de alto valor. Si nuestro sistema de inventario, en lugar de gestionar tornillos y pinturas, gestionara el stock de medicamentos en una farmacia hospitalaria, los materiales de una planta química o los fondos de una entidad bancaria, los estándares de cumplimiento cambiarían sustancialmente y en muchos casos serían exigidos por ley.

El primer estándar de cumplimiento obligatorio en ese escenario sería ISO/IEC 27001. Un sistema crítico maneja información sensible cuya alteración, pérdida o exposición puede tener consecuencias graves. ISO 27001 exigiría implementar un Sistema de Gestión de Seguridad de la Información formal, con identificación y clasificación de todos los activos de información, análisis de riesgos documentado y revisado periódicamente, controles de acceso estrictos basados en el principio de mínimo privilegio, cifrado de datos en tránsito y en reposo, y un plan de respuesta a incidentes probado. En nuestro sistema actual, la columna password de tblusuarios ya representa un riesgo si no se hashea correctamente, y la tabla tblmovimientos_stock sería la base del registro de auditoría que esta norma requeriría de forma obligatoria.

El segundo estándar que entraría en juego dependería del dominio específico. Si el sistema controlara infraestructura física automatizada, como sensores de temperatura en un depósito de medicamentos refrigerados o válvulas en una planta industrial, ISA/IEC 62443 sería obligatoria. Esta norma exigiría segmentar la red de control industrial de la red corporativa, definir niveles de seguridad para cada zona del sistema, y garantizar que un atacante que comprometiera la capa web no pudiera llegar a controlar el hardware físico. Si en cambio el sistema operara en el sector de salud, ISO 13485 reemplazaría a ISO 9001 como marco de calidad, con requisitos mucho más estrictos de trazabilidad y validación de software antes de su uso en producción.

En cualquier escenario crítico, ISO 9001 o su equivalente sectorial seguiría siendo obligatoria, porque la trazabilidad completa de cada decisión de diseño, cada cambio de código y cada prueba realizada se convierte en evidencia legal ante una auditoría o ante un incidente.

---

## ¿Qué concepto de ISO 13407 o ISO 9241-11 sigue siendo útil en sistemas críticos?

El concepto más valioso y más frecuentemente subestimado de estas normas en el contexto de sistemas críticos es la definición de usabilidad en función del contexto de uso específico que plantea ISO 9241-11. La norma no define si un sistema es usable o no en abstracto, sino siempre en relación a quién lo usa, para qué tarea concreta y en qué condiciones reales. Esta perspectiva es especialmente poderosa en entornos críticos porque rompe con la idea de que un sistema técnicamente correcto es automáticamente seguro.

Un controlador de tráfico aéreo que opera bajo carga cognitiva extrema y con restricciones de tiempo de décimas de segundo necesita que las alertas del sistema sean absolutamente inequívocas, imposibles de ignorar accidentalmente y que no requieran interpretación. Un médico de urgencias que consulta el stock de un medicamento mientras atiende a un paciente necesita que la información crítica esté a un solo paso, sin menús secundarios ni confirmaciones innecesarias. En ambos casos, un error de interfaz no produce frustración sino consecuencias irreversibles. ISO 9241-11 proporciona exactamente el marco para detectar esos riesgos antes de que el sistema llegue a producción, porque obliga a evaluar la usabilidad con usuarios reales en condiciones reales, no en un laboratorio controlado.

De ISO 13407, el concepto más duradero es la iteración basada en evaluación con usuarios reales como parte obligatoria del proceso de diseño, no como una etapa opcional al final. En sistemas críticos esto se traduce en simulacros con operadores reales bajo condiciones de estrés antes del despliegue, donde se verifica no solo que el sistema funciona correctamente sino que el ser humano que lo opera no comete errores inducidos por el diseño. Esta práctica, que en un sistema de ferretería es una buena costumbre, en un sistema crítico es una barrera de seguridad tan importante como cualquier control técnico.

# Análisis Final de Estándares

## 1. Tabla comparativa de estándares

| Estándar        | Año (aprox.)         | Enfoque principal                                                                 | ¿Aplica a nuestro proyecto? | Justificación |
|----------------|---------------------|----------------------------------------------------------------------------------|-----------------------------|---------------|
| ISO 9241-11    | 1998 (rev. 2018)    | Usabilidad: eficacia, eficiencia y satisfacción del usuario en contexto de uso   | Sí, directamente            | El sistema será usado por empleados de ferretería sin formación técnica. La eficacia (completar tareas sin errores), la eficiencia (hacerlo en el menor tiempo posible) y la satisfacción son críticas para la adopción real del sistema. Se aplica en el diseño de formularios, alertas de stock y navegación. |
| ISO 13407      | 1999 (reemplazada por ISO 9241-210 en 2010) | Proceso de diseño centrado en el humano (DCH): 4 fases iterativas | Sí, como marco de proceso   | Las 4 fases (entender el contexto de uso → especificar requisitos → producir soluciones → evaluarlas) se corresponden directamente con nuestro flujo Scrum: Sprint 0 fue la fase 1 y 2, el prototipo Figma es la fase 3, y las pruebas de usabilidad serán la fase 4. |
| ISO/IEC 27001  | 2005 (rev. 2022)    | Seguridad de la información: confidencialidad, integridad, disponibilidad. Gestión de riesgos. | Parcialmente                | Nuestro sistema maneja datos de productos, precios, proveedores y movimientos de stock. Si bien no es un sistema crítico en esta etapa, las buenas prácticas de 27001 aplican: hashing de contraseñas, control de acceso por rol (Administrador/Vendedor según tbltipo_usuario), y registro de auditoría (tblmovimientos_stock). |
| ISA/IEC 62443  | 2007–2018 (familia de normas) | Ciberseguridad en sistemas de control industrial (SCADA, PLC, automatización) | No aplica en contexto actual | Este estándar está orientado a infraestructura industrial crítica (plantas de energía, sistemas de transporte, manufactura automatizada). Una ferretería mediana no opera sistemas SCADA ni PLC. Sería relevante si el sistema controlara maquinaria, sensores de temperatura de almacén o automatización de depósitos. |
| ISO 9001       | 1987 (rev. 2015)    | Gestión de calidad en procesos: planificación, ejecución, verificación, mejora continua (PDCA) | Sí, como referencia de proceso | El ciclo PDCA de ISO 9001 se alinea con Scrum: planificar (backlog) → hacer (sprint) → verificar (QA/pruebas) → actuar (retrospectiva). La validación de campos en formularios, el control de stock mínimo y los tests unitarios con TDD son aplicaciones concretas del enfoque de calidad de esta norma. |

---
## Los tres componentes de la usabilidad

| Componente   | Definición                                                                 | Ejemplo concreto en el sistema de stock                                                                 |
|--------------|-----------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------|
| Eficacia     | Grado en que los usuarios logran sus objetivos de forma completa y precisa | El empleado puede registrar una entrada de stock sin errores en los datos del producto (nombre, cantidad, motivo) |
| Eficiencia   | Recursos invertidos (tiempo, esfuerzo) en relación al resultado obtenido   | El encargado accede al listado de productos a reponer en menos de 3 clics desde la pantalla de inicio   |
| Satisfacción | Ausencia de incomodidad y actitud positiva hacia el uso del producto       | El empleado no siente frustración al buscar un producto por categoría ni al interpretar el estado del stock (crítico, bajo, normal) |
## 2. ¿Qué estándares serían obligatorios si el sistema fuera crítico?

Si nuestro sistema de gestión de inventario fuera declarado crítico (por ejemplo, si controlara el stock de medicamentos en una farmacia hospitalaria, insumos en una planta nuclear, o materiales peligrosos en una industria química), los estándares de cumplimiento obligatorio cambiarían sustancialmente.

En primer lugar, ISO/IEC 27001 pasaría de ser una buena práctica a un requisito legal en muchas jurisdicciones. Exigiría un Sistema de Gestión de Seguridad de la Información (SGSI) formal, con clasificación de activos de información, análisis de riesgos documentado, controles de acceso auditables y un plan de respuesta a incidentes.

En segundo lugar, ISA/IEC 62443 sería obligatoria si el inventario estuviera integrado con sistemas de control físico (sensores de temperatura, válvulas automáticas, brazos robóticos de depósito). Definiría niveles de seguridad (SL 1 a SL 4) y zonas de segmentación de red.

En tercer lugar, ISO 9001 o su variante sectorial (como ISO 13485 para dispositivos médicos o AS9100 para aeronáutica) sería obligatoria para garantizar trazabilidad completa de cada movimiento de inventario, validación formal del software antes de su uso en producción, y registros inmutables de auditoría.

Finalmente, ISO 9241-11 e ISO 13407 seguirían siendo relevantes incluso en sistemas críticos, ya que los errores de interfaz (un operador que selecciona el producto equivocado bajo presión) pueden tener consecuencias graves. En entornos críticos, la usabilidad no es un "nice to have" sino un requisito de seguridad funcional.

---

## 3. ¿Qué concepto de ISO 13407 / ISO 9241-11 sigue siendo útil en sistemas críticos?

El concepto más valioso y transferible de ISO 9241-11 es la definición de usabilidad en contexto de uso específico. La norma no define usabilidad en abstracto, sino en función de quién usa el sistema, para qué tarea y en qué entorno. Esta perspectiva es especialmente relevante en sistemas críticos:

Un controlador de tráfico aéreo que usa nuestro tipo de sistema bajo alta carga cognitiva necesita que las alertas sean inequívocas e imposibles de ignorar accidentalmente.

Un médico de urgencias que consulta stock de medicamentos necesita eficiencia máxima (mínima cantidad de pasos para llegar a la información crítica).

De ISO 13407 (hoy ISO 9241-210), el concepto más duradero es la iteración basada en evaluación con usuarios reales. En sistemas críticos esto se traduce en pruebas de usabilidad con operadores reales en condiciones simuladas de estrés, antes de desplegar en producción. Esta práctica reduce los errores humanos inducidos por el diseño, que en sistemas críticos pueden tener consecuencias irreversibles.

---

## 4. Conclusión: ¿Bajo qué estándar certificaríamos nuestro sistema?

Si tuviéramos que certificar nuestro sistema bajo un estándar actual, elegiríamos ISO 9001:2015 como marco de calidad de proceso, complementado con las directrices de usabilidad de ISO 9241-11. ISO 9001 es el más aplicable a nuestra escala y contexto: exigiría formalizar la validación de cada funcionalidad antes de liberarla (lo cual ya estamos haciendo con TDD y GitHub Actions CI), documentar los criterios de aceptación por sprint, y establecer un proceso de mejora continua que en nuestro caso ya existe vía retrospectivas Scrum. Los cambios concretos que implicaría en nuestro diseño actual serían:

Formalizar la trazabilidad completa en tblmovimientos_stock para que cada cambio de stock sea auditable con usuario y timestamp, algo que la arquitectura actual ya soporta.

Documentar formalmente los casos de prueba como evidencia de verificación, tarea que corresponde al QA Lead.

Definir criterios de aceptación medibles para cada caso de uso, vinculados directamente a las tarjetas del tablero Kanban.

---

### Relación con las decisiones de diseño del TP1

El patrón Observer favorece el cumplimiento de ISO 9001 y ISO/IEC 27001: al desacoplar la lógica de notificación del registro de movimientos, cada componente es testeable de forma aislada (ISO 9001) y el sistema de alertas puede extenderse para incluir un log de auditoría sin modificar el núcleo (ISO 27001, principio de mínimo impacto en cambios).

El patrón State favorece ISO 9001: encapsular el comportamiento por estado hace que las reglas de negocio sean explícitas, verificables y documentables, lo cual facilita la trazabilidad de por qué una operación fue permitida o rechazada.

La tabla tblmovimientos_stock con stock_anterior y stock_nuevo es en sí misma una implementación práctica del principio de no repudio de ISO/IEC 27001: cada movimiento queda registrado de forma que no puede negarse ni alterarse sin dejar rastro.
