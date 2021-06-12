package main // Potrzebne - nie wiem co robi

import ( // Importowanie jakiegoś shitpostu
	"bufio"     // Scaner który CZYTA spacje
	"fmt"       // Podstawowy output
	"io/ioutil" // To coś przetworzyło mi body na cyfry - jest git
	"net/http"  // Operowanie HTTP
	"os"        // Coś związanego z systemem operacyjnym ( ? )
	"strings"   // Działania ze stringami
)

func main() { // główna funckja
	fmt.Println("Napsiz tytuł ksiąki") // Output

	var x string = "" // Zmienna string

	scanner := bufio.NewScanner(os.Stdin) // Scaner który skanuje spacje
	if scanner.Scan() {                   // Jezeli zeskanował
		x = scanner.Text() // to zapisuje to w zmiennej
	}

	x = strings.ReplaceAll(x, " ", "%20") // Podmienia wszystkie spacje na %20

	url := "http://biblioteka.dkonto.pl/index.php?book=" + x // Tworzenie linku

	fmt.Println(url) // Wyświetlanie linku

	resp, err := http.Get(url) // Pobieranie danych z linku

	fmt.Println("") // Spacja między liniami

	fmt.Println(resp) // Wyświetlanie odpowiedzi

	if err != nil { // Sprawdzanie czy jest błąd
		fmt.Println(err) // Wyświetlanie błędu
	}

	fmt.Println("") // Kolejna przerwa

	content, err := ioutil.ReadAll(resp.Body) // Czytanie body w taki sposób ze mam char'y w int'ach

	fmt.Println(string(content)) // Konwertowanie int'ów na chary i wyświetlanie ich

}
