# MetaZone — Sistema de E-commerce en Go

> Proyecto académico de gestión de e-commerce desarrollado en **Go (Golang)** con base de datos **MySQL (Laragon)** y plantillas HTML usando **Bulma CSS**.

---

## ¿Qué es MetaZone?

MetaZone es un sistema de gestión de e-commerce completo que permite:
- Registrar y autenticar usuarios
- Gestionar un catálogo de productos con stock
- Agregar productos al carrito de compras
- Seleccionar método de pago y finalizar compras
- Ver historial de pedidos
- Panel de administración para gestionar usuarios, productos y pedidos

---

## Tecnologías utilizadas

| **Go (Golang)** | Lenguaje backend | Compilado, rápido, excelente para APIs y servidores web |
| **MySQL** | Base de datos | Relacional, gratuito, integrado en Laragon |
| **Laragon** | Entorno de desarrollo | Incluye MySQL + phpMyAdmin en Windows |
| **Bulma CSS** | Estilos frontend | Framework CSS moderno, sin JavaScript requerido |
| **gorilla/mux** | Router HTTP | Rutas con parámetros y métodos HTTP específicos |
| **gorilla/sessions** | Sesiones | Cookies firmadas para autenticación |
| **bcrypt** | Seguridad | Hash de contraseñas seguro |
| **html/template** | Plantillas | Generación de HTML en el servidor (SSR) |

---

## Servicios implementados

### 1. Registro de usuario
- **Ruta:** `GET /registro` · `POST /registro`
- **Función:** El usuario completa un formulario con nombre, apellido, email y contraseña.
- **Seguridad:** La contraseña se hashea con **bcrypt** antes de guardarse. Nunca se almacena en texto plano.
- **Auto-login:** Tras registrarse, se crea una sesión automáticamente.
- **Archivos:** `handlers.go → RegisterForm/RegisterSubmit`, `services.go → RegistrarUsuario`

### 2. Login de usuario
- **Ruta:** `GET /login` · `POST /login`
- **Función:** Verifica email y contraseña contra los datos almacenados.
- **Sesión:** Usa cookies firmadas (`gorilla/sessions`) para mantener la sesión entre páginas.
- **Seguridad:** No se revela si el email existe o no (mismo mensaje de error para ambos casos).
- **Archivos:** `handlers.go → LoginForm/LoginSubmit`, `services.go → AutenticarUsuario`

### 3. Registro de producto (Admin)
- **Ruta:** `GET /admin/productos/nuevo` · `POST /admin/productos/nuevo`
- **Función:** El administrador puede agregar productos con nombre, descripción, precio, stock e imagen.
- **Protección:** Solo accesible con rol `admin` (verificado en middleware).
- **Archivos:** `handlers.go → AdminProductCreate`, `services.go → CrearProducto`

### 4. Stock de productos (Admin)
- **Ruta:** `GET /admin/stock` · `POST /admin/stock/{id}`
- **Función:** Panel para ver el inventario actual y actualizar unidades disponibles.
- **Indicadores visuales:** Verde = OK, Amarillo = pocas unidades, Rojo = agotado.
- **Archivos:** `handlers.go → AdminStock/AdminStockUpdate`, `services.go → ActualizarStock`

### 5. 🛒 Agregar al carrito
- **Ruta:** `POST /carrito/agregar/{id}`
- **Función:** Agrega un producto al carrito del usuario con la cantidad especificada.
- **Verificación:** Comprueba que haya stock suficiente antes de agregar.
- **Persistencia:** El carrito se guarda en la base de datos (no en la sesión).
- **Archivos:** `handlers.go → CartAdd`, `services.go → AgregarAlCarrito`

### 6. Selección de método de pago
- **Ruta:** `GET /pago`
- **Función:** El usuario elige entre 3 métodos: tarjeta, transferencia bancaria o contra entrega.
- **UI dinámica:** Al seleccionar "tarjeta", aparecen los campos del número de tarjeta (JavaScript).
- **Archivos:** `handlers.go → CheckoutForm`, `templates/pages/checkout.html`

### 7. Finalizar compra
- **Ruta:** `POST /pago`
- **Función:** Crea el pedido, descuenta el stock y vacía el carrito. Todo en una **transacción SQL**.
- **Transacción:** Si algo falla (stock insuficiente, error de DB), todo se revierte automáticamente.
- **Archivos:** `handlers.go → CheckoutSubmit`, `services.go → CrearPedido`

### 8. Búsqueda de productos
- **Ruta:** `GET /buscar?q=texto`
- **Función:** Filtra productos por nombre, descripción o categoría.
- **SQL:** Usa `LIKE` con `%` para búsqueda parcial.
- **Archivos:** `handlers.go → Search`, `services.go → BuscarProductos`

### Servicios adicionales:
- **Ver historial de pedidos** (`GET /pedidos`) — Lista todas las compras del usuario
- **Detalle de pedido** (`GET /pedidos/{id}`) — Muestra todos los productos de una compra
- **Perfil de usuario** (`GET/POST /perfil`) — El usuario puede editar su nombre y apellido
- **Panel de administración** — Dashboard con estadísticas, gestión de usuarios, productos y pedidos

---

## Requisitos previos

1. **Go 1.21+** → https://go.dev/dl/
2. **Laragon** (Windows) → https://laragon.org/download/
3. **Git** → https://git-scm.com/

---

## Configuración de la base de datos

### Paso 1: Iniciar Laragon
Abre Laragon y haz clic en **"Start All"** (inicia Apache + MySQL).

### Paso 2: Crear la base de datos
1. Abre `http://localhost/phpmyadmin`
2. Haz clic en **"Nueva"** (columna izquierda)
3. Nombre: `metazone`
4. Cotejamiento: `utf8mb4_unicode_ci`
5. Haz clic en **"Crear"**

Las tablas se crean automáticamente cuando se ejecuta el proyecto.

### Configuración por defecto (Laragon)
```
Host:     localhost
Puerto:   3306
Usuario:  root
Password: (vacío)
Base de datos: metazone
```

---

### Para crear un usuario administrador:
1. Regístrate normalmente en `/registro`
2. En phpMyAdmin, ejecuta:
```sql
UPDATE usuarios SET rol = 'admin' WHERE email = 'tu@email.com';
```
3. Ahora tendrás acceso a `/admin`

---

## Cómo funciona el código

### Arquitectura en capas

```
┌─────────────────────────────────────────────────┐
│               NAVEGADOR (Browser)               │
│         GET /productos, POST /carrito...        │
└────────────────────┬────────────────────────────┘
                     │ HTTP Request
┌────────────────────▼────────────────────────────┐
│           ROUTER (gorilla/mux)                  │
│   Decide qué handler llama según la URL         │
└────────────────────┬────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────┐
│           MIDDLEWARE                            │
│   Auth → ¿está logueado?                       │
│   AdminOnly → ¿es admin?                       │
└────────────────────┬────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────┐
│           HANDLER (handlers.go)                 │
│   Lee el request, llama al servicio,           │
│   renderiza el template HTML                   │
└────────────────────┬────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────┐
│           SERVICE (services.go)                 │
│   Lógica de negocio: validaciones,             │
│   consultas SQL, transacciones                 │
└────────────────────┬────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────┐
│           BASE DE DATOS (MySQL/Laragon)         │
│   Almacenamiento persistente de datos          │
└─────────────────────────────────────────────────┘
```

---

## Flujo de una petición HTTP

**Ejemplo: Usuario agrega un producto al carrito**

1. **Browser:** Envía `POST /carrito/agregar/3` con `cantidad=2`
2. **Router (gorilla/mux):** Reconoce la ruta, extrae `{id}=3`, pasa al middleware
3. **Middleware Auth:** Verifica la cookie de sesión. ¿Hay `usuario_id`? Si no → redirige a `/login`
4. **Handler CartAdd:**
   - Lee el ID del producto de la URL con `mux.Vars(r)`
   - Lee la cantidad del formulario con `r.FormValue("cantidad")`
   - Lee el ID del usuario de la sesión
   - Llama a `services.AgregarAlCarrito(db, usuarioID, productoID, cantidad)`
5. **Service AgregarAlCarrito:**
   - Consulta MySQL para verificar el stock
   - Si hay stock, inserta/actualiza `carrito_items`
   - Si no hay stock, devuelve `ErrStockInsuficiente`
6. **Handler:** Si hubo error → redirige con parámetro de error. Si OK → redirige a `/carrito`
7. **Browser:** Recibe el redirect y carga la página del carrito

---

## Conceptos clave de Go usados

### Structs y métodos
```go
// Una struct es como una clase sin herencia
type Producto struct {
    ID     int
    Nombre string
    Precio float64
    Stock  int
}

// Los métodos se definen fuera de la struct
func (p *Producto) TieneStock() bool {
    return p.Stock > 0
}
```

### Interfaces
```go
// http.Handler es una interfaz: cualquier tipo que tenga ServeHTTP() la implementa
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

### Defer (diferido)
```go
// defer ejecuta la función cuando la función actual termina
rows, _ := db.Query(...)
defer rows.Close()  // Se cierra al final, sin importar cómo termina la función
```

### Errores como valores
```go
// En Go, los errores son valores normales que se devuelven
usuario, err := services.AutenticarUsuario(db, email, pass)
if err != nil {
    // Manejar el error
}
```

### Transacciones SQL
```go
// Una transacción garantiza que todo ocurre o nada
tx, _ := db.Begin()
defer tx.Rollback()    // Si algo falla, revertir

tx.Exec("INSERT ...")  // Operación 1
tx.Exec("UPDATE ...")  // Operación 2

tx.Commit()            // Confirmar todo
```

---

