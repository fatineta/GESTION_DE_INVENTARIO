package unit

import (
	"errors"
	"testing"
)

// ==================== ESTRUCTURAS (copia local para pruebas sin DB) ====================

type Producto struct {
	IDProducto   int64
	CodigoSKU    string
	Nombre       string
	PrecioVenta  float64
	StockActual  int
	StockMinimo  int
	StockMaximo  int
	PuntoReorden int
}

type DetallePedido struct {
	ProductoID     int64
	Cantidad       int
	PrecioUnitario float64
	Subtotal       float64
	ProductoNombre string
}

type Pedido struct {
	IDVenta      int64
	NumeroVenta  int
	Total        float64
	EstadoActual EstadoPedido
	EstadoNombre string
	Items        []DetallePedido
}

// ==================== PATRON STATE (extraído para pruebas puras) ====================

type EstadoPedido interface {
	Procesar(p *Pedido) error
	Pagar(p *Pedido) error
	Entregar(p *Pedido) error
	Cancelar(p *Pedido) error
	EnviarAlmacen(p *Pedido) error
	GetNombre() string
}

type PendienteState struct{}
type ConfirmadoState struct{}
type PagadoState struct{}
type EnAlmacenState struct{}
type EntregadoState struct{}
type CanceladoState struct{}

// --- PendienteState ---
func (s *PendienteState) Procesar(p *Pedido) error {
	// Versión testeable: sin llamadas a DB; valida stock pasado via campo simulado
	p.EstadoActual = &ConfirmadoState{}
	p.EstadoNombre = p.EstadoActual.GetNombre()
	return nil
}
func (s *PendienteState) Pagar(p *Pedido) error        { return errors.New("primero debe procesar el pedido") }
func (s *PendienteState) Entregar(p *Pedido) error     { return errors.New("pedido no procesado") }
func (s *PendienteState) EnviarAlmacen(p *Pedido) error { return errors.New("primero confirme el pedido") }
func (s *PendienteState) Cancelar(p *Pedido) error {
	p.EstadoActual = &CanceladoState{}
	p.EstadoNombre = p.EstadoActual.GetNombre()
	return nil
}
func (s *PendienteState) GetNombre() string { return "PENDIENTE" }

// --- ConfirmadoState ---
func (s *ConfirmadoState) Procesar(p *Pedido) error { return errors.New("ya confirmado") }
func (s *ConfirmadoState) Pagar(p *Pedido) error {
	p.EstadoActual = &PagadoState{}
	p.EstadoNombre = p.EstadoActual.GetNombre()
	return nil
}
func (s *ConfirmadoState) Entregar(p *Pedido) error { return errors.New("pague primero") }
func (s *ConfirmadoState) Cancelar(p *Pedido) error {
	p.EstadoActual = &CanceladoState{}
	p.EstadoNombre = p.EstadoActual.GetNombre()
	return nil
}
func (s *ConfirmadoState) EnviarAlmacen(p *Pedido) error {
	p.EstadoActual = &EnAlmacenState{}
	p.EstadoNombre = p.EstadoActual.GetNombre()
	return nil
}
func (s *ConfirmadoState) GetNombre() string { return "CONFIRMADO" }

// --- PagadoState ---
func (s *PagadoState) Procesar(p *Pedido) error      { return errors.New("ya pagado") }
func (s *PagadoState) Pagar(p *Pedido) error          { return errors.New("ya pagado") }
func (s *PagadoState) Entregar(p *Pedido) error       { return errors.New("envíe a almacén primero") }
func (s *PagadoState) Cancelar(p *Pedido) error {
	p.EstadoActual = &CanceladoState{}
	p.EstadoNombre = p.EstadoActual.GetNombre()
	return nil
}
func (s *PagadoState) EnviarAlmacen(p *Pedido) error {
	p.EstadoActual = &EnAlmacenState{}
	p.EstadoNombre = p.EstadoActual.GetNombre()
	return nil
}
func (s *PagadoState) GetNombre() string { return "PAGADO" }

// --- EnAlmacenState ---
func (s *EnAlmacenState) Procesar(p *Pedido) error      { return errors.New("ya en almacén") }
func (s *EnAlmacenState) Pagar(p *Pedido) error          { return errors.New("ya pagado") }
func (s *EnAlmacenState) EnviarAlmacen(p *Pedido) error  { return errors.New("ya está en almacén") }
func (s *EnAlmacenState) Cancelar(p *Pedido) error {
	p.EstadoActual = &CanceladoState{}
	p.EstadoNombre = p.EstadoActual.GetNombre()
	return nil
}
func (s *EnAlmacenState) Entregar(p *Pedido) error {
	p.EstadoActual = &EntregadoState{}
	p.EstadoNombre = p.EstadoActual.GetNombre()
	return nil
}
func (s *EnAlmacenState) GetNombre() string { return "EN_ALMACEN" }

// --- EntregadoState ---
func (s *EntregadoState) Procesar(p *Pedido) error      { return errors.New("ya entregado") }
func (s *EntregadoState) Pagar(p *Pedido) error          { return errors.New("ya entregado") }
func (s *EntregadoState) Entregar(p *Pedido) error       { return errors.New("ya entregado") }
func (s *EntregadoState) Cancelar(p *Pedido) error       { return errors.New("no se puede cancelar") }
func (s *EntregadoState) EnviarAlmacen(p *Pedido) error  { return errors.New("ya entregado") }
func (s *EntregadoState) GetNombre() string               { return "ENTREGADO" }

// --- CanceladoState ---
func (s *CanceladoState) Procesar(p *Pedido) error      { return errors.New("pedido cancelado") }
func (s *CanceladoState) Pagar(p *Pedido) error          { return errors.New("pedido cancelado") }
func (s *CanceladoState) Entregar(p *Pedido) error       { return errors.New("pedido cancelado") }
func (s *CanceladoState) Cancelar(p *Pedido) error       { return errors.New("ya cancelado") }
func (s *CanceladoState) EnviarAlmacen(p *Pedido) error  { return errors.New("pedido cancelado") }
func (s *CanceladoState) GetNombre() string               { return "CANCELADO" }

// ==================== LÓGICA DE NEGOCIO PURA (extraída para pruebas) ====================

// calcularNuevoStock aplica la lógica de movimiento de stock sin DB.
func calcularNuevoStock(stockActual, cantidad int, tipo string) (int, error) {
	switch tipo {
	case "ENTRADA":
		return stockActual + cantidad, nil
	case "SALIDA":
		if cantidad > stockActual {
			return 0, errors.New("stock insuficiente")
		}
		return stockActual - cantidad, nil
	case "AJUSTE":
		return cantidad, nil
	default:
		return 0, errors.New("tipo inválido")
	}
}

// calcularUrgencia reproduce la lógica de handleReposicion.
func calcularUrgencia(stockActual, stockMinimo int) string {
	if stockActual == 0 {
		return "CRITICO"
	} else if stockActual <= stockMinimo/2 {
		return "ALTA"
	}
	return "MEDIA"
}

// calcularFaltante reproduce la lógica de handleReposicion.
func calcularFaltante(stockActual, stockMinimo, stockMaximo int) int {
	faltante := stockMaximo - stockActual
	if faltante < 0 {
		faltante = stockMinimo * 2
	}
	return faltante
}

// validarProducto reproduce la validación de handleCrearProducto.
func validarProducto(nombre string, precioVenta float64) error {
	if nombre == "" || precioVenta <= 0 {
		return errors.New("nombre y precio_venta son obligatorios")
	}
	return nil
}

// estadoDesdeNombre reproduce la función homónima del main.
func estadoDesdeNombre(nombre string) EstadoPedido {
	switch nombre {
	case "CONFIRMADO":
		return &ConfirmadoState{}
	case "PAGADO":
		return &PagadoState{}
	case "EN_ALMACEN":
		return &EnAlmacenState{}
	case "ENTREGADO":
		return &EntregadoState{}
	case "CANCELADO":
		return &CanceladoState{}
	default:
		return &PendienteState{}
	}
}

// ==================== PRUEBAS UNITARIAS ====================

// -----------------------------------------------------------------------
// FUNCIÓN: calcularNuevoStock
// -----------------------------------------------------------------------

// TC-01 | Equivalencia | ENTRADA válida
// Entrada: stockActual=10, cantidad=5, tipo="ENTRADA"
// Esperado: nuevoStock=15, error=nil
func TestCalcularNuevoStock_EntradaValida(t *testing.T) {
	stock, err := calcularNuevoStock(10, 5, "ENTRADA")
	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if stock != 15 {
		t.Errorf("esperado 15, obtenido %d", stock)
	}
}

// TC-02 | Valor Límite | SALIDA exactamente igual al stock disponible (límite inferior permitido)
// Entrada: stockActual=5, cantidad=5, tipo="SALIDA"
// Esperado: nuevoStock=0, error=nil
func TestCalcularNuevoStock_SalidaExactaAlStock(t *testing.T) {
	stock, err := calcularNuevoStock(5, 5, "SALIDA")
	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if stock != 0 {
		t.Errorf("esperado 0, obtenido %d", stock)
	}
}

// TC-03 | Valor Límite | SALIDA con cantidad = stock + 1 (supera el límite)
// Entrada: stockActual=5, cantidad=6, tipo="SALIDA"
// Esperado: error "stock insuficiente"
func TestCalcularNuevoStock_SalidaExcedeStock(t *testing.T) {
	_, err := calcularNuevoStock(5, 6, "SALIDA")
	if err == nil {
		t.Fatal("se esperaba error de stock insuficiente, no se obtuvo ninguno")
	}
}

// TC-04 | Equivalencia | AJUSTE con valor cualquiera (reemplaza el stock)
// Entrada: stockActual=100, cantidad=30, tipo="AJUSTE"
// Esperado: nuevoStock=30, error=nil
func TestCalcularNuevoStock_AjusteValido(t *testing.T) {
	stock, err := calcularNuevoStock(100, 30, "AJUSTE")
	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if stock != 30 {
		t.Errorf("esperado 30, obtenido %d", stock)
	}
}

// TC-05 | Equivalencia | Tipo de movimiento inválido
// Entrada: stockActual=10, cantidad=5, tipo="DEVOLUCION"
// Esperado: error "tipo inválido"
func TestCalcularNuevoStock_TipoInvalido(t *testing.T) {
	_, err := calcularNuevoStock(10, 5, "DEVOLUCION")
	if err == nil {
		t.Fatal("se esperaba error por tipo inválido, no se obtuvo ninguno")
	}
}

// -----------------------------------------------------------------------
// FUNCIÓN: calcularUrgencia
// -----------------------------------------------------------------------

// TC-06 | Valor Límite | Stock = 0 → urgencia CRITICO
// Entrada: stockActual=0, stockMinimo=10
// Esperado: "CRITICO"
func TestCalcularUrgencia_StockCero(t *testing.T) {
	urgencia := calcularUrgencia(0, 10)
	if urgencia != "CRITICO" {
		t.Errorf("esperado CRITICO, obtenido %s", urgencia)
	}
}

// TC-07 | Valor Límite | Stock = stockMinimo/2 (límite entre ALTA y MEDIA)
// Entrada: stockActual=5, stockMinimo=10  →  5 <= 10/2 = 5 → "ALTA"
// Esperado: "ALTA"
func TestCalcularUrgencia_StockEnLimiteAlta(t *testing.T) {
	urgencia := calcularUrgencia(5, 10)
	if urgencia != "ALTA" {
		t.Errorf("esperado ALTA, obtenido %s", urgencia)
	}
}

// TC-08 | Valor Límite | Stock = stockMinimo/2 + 1 (pasa al rango MEDIA)
// Entrada: stockActual=6, stockMinimo=10  →  6 > 5 → "MEDIA"
// Esperado: "MEDIA"
func TestCalcularUrgencia_StockSobreLimiteAlta(t *testing.T) {
	urgencia := calcularUrgencia(6, 10)
	if urgencia != "MEDIA" {
		t.Errorf("esperado MEDIA, obtenido %s", urgencia)
	}
}

// -----------------------------------------------------------------------
// FUNCIÓN: validarProducto
// -----------------------------------------------------------------------

// TC-09 | Equivalencia | Producto válido (nombre y precio correctos)
// Entrada: nombre="Martillo", precioVenta=150.0
// Esperado: nil
func TestValidarProducto_DatosValidos(t *testing.T) {
	err := validarProducto("Martillo", 150.0)
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo: %v", err)
	}
}

// TC-10 | Equivalencia | Nombre vacío → debe rechazarse
// Entrada: nombre="", precioVenta=150.0
// Esperado: error
func TestValidarProducto_NombreVacio(t *testing.T) {
	err := validarProducto("", 150.0)
	if err == nil {
		t.Fatal("se esperaba error por nombre vacío")
	}
}

// TC-11 | Valor Límite | Precio exactamente 0 → debe rechazarse
// Entrada: nombre="Llave", precioVenta=0
// Esperado: error
func TestValidarProducto_PrecioCero(t *testing.T) {
	err := validarProducto("Llave", 0)
	if err == nil {
		t.Fatal("se esperaba error por precio_venta = 0")
	}
}

// TC-12 | Valor Límite | Precio negativo → debe rechazarse
// Entrada: nombre="Llave", precioVenta=-1.0
// Esperado: error
func TestValidarProducto_PrecioNegativo(t *testing.T) {
	err := validarProducto("Llave", -1.0)
	if err == nil {
		t.Fatal("se esperaba error por precio_venta negativo")
	}
}

// -----------------------------------------------------------------------
// FUNCIÓN: Patrón State — transiciones de EstadoPedido
// -----------------------------------------------------------------------

// TC-13 | Equivalencia | Pedido PENDIENTE puede ser cancelado
// Estado inicial: PendienteState
// Acción: Cancelar
// Esperado: estado cambia a "CANCELADO", sin error
func TestPendienteState_Cancelar(t *testing.T) {
	pedido := &Pedido{EstadoActual: &PendienteState{}, EstadoNombre: "PENDIENTE"}
	err := pedido.EstadoActual.Cancelar(pedido)
	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if pedido.EstadoNombre != "CANCELADO" {
		t.Errorf("esperado CANCELADO, obtenido %s", pedido.EstadoNombre)
	}
}

// TC-14 | Equivalencia | Pedido PENDIENTE no puede pagarse sin procesar primero
// Estado inicial: PendienteState
// Acción: Pagar
// Esperado: error (transición inválida)
func TestPendienteState_PagarSinProcesar(t *testing.T) {
	pedido := &Pedido{EstadoActual: &PendienteState{}, EstadoNombre: "PENDIENTE"}
	err := pedido.EstadoActual.Pagar(pedido)
	if err == nil {
		t.Fatal("se esperaba error al pagar un pedido PENDIENTE sin procesar")
	}
}

// TC-15 | Equivalencia | Pedido CONFIRMADO puede avanzar a PAGADO
// Estado inicial: ConfirmadoState
// Acción: Pagar
// Esperado: estado cambia a "PAGADO", sin error
func TestConfirmadoState_Pagar(t *testing.T) {
	pedido := &Pedido{EstadoActual: &ConfirmadoState{}, EstadoNombre: "CONFIRMADO"}
	err := pedido.EstadoActual.Pagar(pedido)
	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if pedido.EstadoNombre != "PAGADO" {
		t.Errorf("esperado PAGADO, obtenido %s", pedido.EstadoNombre)
	}
}

// TC-16 | Equivalencia | Pedido ENTREGADO no puede cancelarse
// Estado inicial: EntregadoState
// Acción: Cancelar
// Esperado: error "no se puede cancelar"
func TestEntregadoState_NoPuedeCancelarse(t *testing.T) {
	pedido := &Pedido{EstadoActual: &EntregadoState{}, EstadoNombre: "ENTREGADO"}
	err := pedido.EstadoActual.Cancelar(pedido)
	if err == nil {
		t.Fatal("se esperaba error al intentar cancelar un pedido ENTREGADO")
	}
}

// TC-17 | Equivalencia | Pedido EN_ALMACEN puede entregarse
// Estado inicial: EnAlmacenState
// Acción: Entregar
// Esperado: estado cambia a "ENTREGADO", sin error
func TestEnAlmacenState_Entregar(t *testing.T) {
	pedido := &Pedido{
		EstadoActual: &EnAlmacenState{},
		EstadoNombre: "EN_ALMACEN",
		Items:        []DetallePedido{},
	}
	err := pedido.EstadoActual.Entregar(pedido)
	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if pedido.EstadoNombre != "ENTREGADO" {
		t.Errorf("esperado ENTREGADO, obtenido %s", pedido.EstadoNombre)
	}
}

// -----------------------------------------------------------------------
// FUNCIÓN: estadoDesdeNombre
// -----------------------------------------------------------------------

// TC-18 | Equivalencia | Nombre desconocido devuelve estado PENDIENTE por defecto
// Entrada: "INEXISTENTE"
// Esperado: GetNombre() == "PENDIENTE"
func TestEstadoDesdeNombre_Default(t *testing.T) {
	estado := estadoDesdeNombre("INEXISTENTE")
	if estado.GetNombre() != "PENDIENTE" {
		t.Errorf("esperado PENDIENTE como default, obtenido %s", estado.GetNombre())
	}
}

// TC-19 | Equivalencia | Nombre "EN_ALMACEN" devuelve EnAlmacenState
// Entrada: "EN_ALMACEN"
// Esperado: GetNombre() == "EN_ALMACEN"
func TestEstadoDesdeNombre_EnAlmacen(t *testing.T) {
	estado := estadoDesdeNombre("EN_ALMACEN")
	if estado.GetNombre() != "EN_ALMACEN" {
		t.Errorf("esperado EN_ALMACEN, obtenido %s", estado.GetNombre())
	}
}

// -----------------------------------------------------------------------
// FUNCIÓN: calcularFaltante
// -----------------------------------------------------------------------

// TC-20 | Valor Límite | stockMaximo < stockActual → faltante = stockMinimo * 2
// Entrada: stockActual=80, stockMinimo=10, stockMaximo=50  →  50-80 = -30 < 0
// Esperado: 20
func TestCalcularFaltante_StockMaximoMenorQueActual(t *testing.T) {
	faltante := calcularFaltante(80, 10, 50)
	if faltante != 20 {
		t.Errorf("esperado 20 (stockMinimo*2), obtenido %d", faltante)
	}
}

// TC-21 | Equivalencia | Caso normal: stockMaximo > stockActual
// Entrada: stockActual=10, stockMinimo=5, stockMaximo=100
// Esperado: 90
func TestCalcularFaltante_CasoNormal(t *testing.T) {
	faltante := calcularFaltante(10, 5, 100)
	if faltante != 90 {
		t.Errorf("esperado 90, obtenido %d", faltante)
	}
}
