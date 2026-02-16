package main
import (
	"bufio"
	"fmt"
	"g-cinemapro/inventory"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Función para limpiar la pantalla (Cross-platform)
func CallClear() {
	// Intentamos limpiar la pantalla enviando el comando al sistema operativo
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") // Forzamos el uso de cmd para limpiar
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	// Truco extra: Si el comando falla, imprimimos muchos saltos de línea
	// para "empujar" lo viejo hacia arriba y que no se vea.
	fmt.Print("\033[H\033[2J")
}

// Función para mostrar el logo (la reutilizamos en cada opción)
func ShowLogo() {
	fmt.Println("\033[34m") // Azul
	fmt.Println("##############################################")
	fmt.Println("##                                          ##")
	fmt.Println("##         G - C I N E M A   P R O          ##")
	fmt.Println("##       Streaming Management System        ##")
	fmt.Println("##                                          ##")
	fmt.Println("##                                          ##")
	fmt.Println("##         UNIVERSIDAD INTERNACIONAL        ##")
	fmt.Println("##               DEL ECUADOR                ##")
	fmt.Println("##                                          ##")
	fmt.Println("##############################################")
	fmt.Println("\033[0m") // Reset
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Pantalla de Bienvenida Inicial
	CallClear()
	ShowLogo()
	fmt.Println("\nPresiona [ENTER] para iniciar el sistema...")
	reader.ReadString('\n')

	fmt.Print("Ingresa tu nombre de Usuario: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	for {
		CallClear() // Borra la pantalla anterior
		ShowLogo()  // Muestra el logo arriba de nuevo

		fmt.Printf("\n--- BIENVENIDO %s | MENÚ PRINCIPAL ---\n", strings.ToUpper(nombre))
		fmt.Println("1. Ver Catálogo Completo")
		fmt.Println("2. Registrar Nueva Película")
		fmt.Println("3. Salir")
		fmt.Print("\nSelecciona una opción: ")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			CallClear()
			ShowLogo()
			fmt.Println("\n--- LISTADO DE PELÍCULAS ---")
			listado := inventory.GetList()
			if len(listado) == 0 {
				fmt.Println("[!] El catálogo está vacío.")
			} else {
				for id, m := range listado {
					fmt.Printf("ID: %d | Título: %-15s | Género: %s\n", id, m.Title, m.Genre)
				}
			}
			fmt.Println("\nPresiona [ENTER] para volver al menú...")
			reader.ReadString('\n')

		case 2:
			CallClear()
			ShowLogo()
			fmt.Println("\n--- REGISTRO DE NUEVA PELÍCULA ---")
			var id int
			fmt.Print("ID de la película: ")
			fmt.Scanln(&id)

			fmt.Print("Título: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)

			fmt.Print("Género: ")
			genero, _ := reader.ReadString('\n')
			genero = strings.TrimSpace(genero)

			msg, err := inventory.RegisterMovie(id, titulo, genero)
			if err != nil {
				fmt.Println("\n[X] Error:", err)
			} else {
				fmt.Println("\n[V]", msg)
			}
			fmt.Println("\nPresiona [ENTER] para volver al menú...")
			reader.ReadString('\n')

		case 3:
			fmt.Printf("\nCerrando sesión... ¡Hasta luego %s!\n", nombre)
			return

		default:
			fmt.Println("Opción inválida. Presiona ENTER para reintentar.")
			reader.ReadString('\n')
		}
	}
}
