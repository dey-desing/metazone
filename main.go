package main

import (
	"metazone/services"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
	"net/http"

	"metazone/db"
	"metazone/controllers"
	"metazone/models"
)

func main() {

	//conectar a la base de datos
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos %v", err)
	}

	defer dbConn.Close()

	controllers.InitRoutes()

	fmt.Println("Conexión a la base de datos establecida exitosamente")

	//levantar un servidor para verificar la conexion con la base de datos
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== BIENVENIDO A METAZONE =====")
		fmt.Println("1. Registrar usuario")
		fmt.Println("2. Lista de usuarios")
		fmt.Println("3. Agregar producto")
		fmt.Println("4. Lista de productos")
		fmt.Println("5. Realizar pago")
		fmt.Println("6. Listar pagos")
		fmt.Println("7. Salir")
		fmt.Print("Seleccione una opción: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			createUserMenu(reader)
		case "2":
			listUsersMenu()
		case "3":
			createProductMenu(reader)
		case "4":
			listProductsMenu()
		case "5":
			createPaymentMenu(reader)	
		case "6":
			listPaymentsMenu()
		case "7":
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}

//Crear usuario
func createUserMenu(reader *bufio.Reader) {
	fmt.Print("Nombre: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Correo: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Contraseña: ")
	password, _ := reader.ReadString('\n')

	user, err := services.CreateUser(
		models.User{
			Name:     strings.TrimSpace(name),
			Email:    strings.TrimSpace(email),
			Password: strings.TrimSpace(password),
		},
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Usuario creado con éxito:", user.Name)
}

//Listar usuarios
func listUsersMenu() {
	users := services.GetUsers()

	if len(users) == 0 {
		fmt.Println("No hay usuarios registrados.")
		return
	}

	fmt.Println("\nUsuarios registrados:")
	for _, u := range users {
		fmt.Printf("ID: %d | Nombre: %s | Email: %s\n", u.ID, u.Name, u.Email)
	}
}

//Crear producto
func createProductMenu(reader *bufio.Reader) {
	fmt.Print("Nombre del producto: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Descripción del producto: ")
	description, _ := reader.ReadString('\n')
	
	fmt.Print("Precio del producto: ")
	priceStr, _ := reader.ReadString('\n')
	price, _ := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)

	fmt.Print("Stock: ")
	stockStr, _ := reader.ReadString('\n')
	stock, _ := strconv.Atoi(strings.TrimSpace(stockStr))

	services.CreateProduct(
		strings.TrimSpace(name),
		strings.TrimSpace(description),
		price,
		stock,
	)

	fmt.Println("Producto creado con éxito:", name)

}

//Listar productos
func listProductsMenu() {
	products := services.GetProducts()
	for _, p := range products {
		fmt.Printf("ID: %d | Nombre: %s | Descripción: %s | Precio: %.2f | Stock: %d\n", p.ID, p.Name, p.Description, p.Price, p.GetStock())
	}	
}

// Crear pago
func createPaymentMenu(reader *bufio.Reader) {
	fmt.Print("Pedido: ")
	orderStr, _ := reader.ReadString('\n')
	orderID, _ := strconv.Atoi(strings.TrimSpace(orderStr))

	fmt.Print("Monto a pagar: ")
	amountStr, _ := reader.ReadString('\n')
	amount, _ := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)

	fmt.Print("Método de pago (efectivo/tarjeta): ")
	method, _ := reader.ReadString('\n')

	payment, err := services.CreatePayment(
		orderID,
		amount,
		strings.TrimSpace(method),
	)

	if err != nil {
		fmt.Println("Error en el pago:", err)
		return
	}

	fmt.Println("Pago realizado con éxito")
	fmt.Println("Estado:", payment.Status)
}

//Listar pagos
func listPaymentsMenu() {
	payments := services.GetPayments()
	for _, p := range payments {
		fmt.Printf("ID: %d | Pedido ID: %d | Monto: %.2f | Método: %s | Estado: %s\n", p.ID, p.OrderID, p.Amount, p.Method, p.Status)
	}	
}

