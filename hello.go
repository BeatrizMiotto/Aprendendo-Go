package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	/*nome, semanas := devolveNomeSemanas() //caso não queira retornar um dos valores basta colocar o "_" exemplo "_,semans :=devolveNomeSemanas()".
	fmt.Println(nome, "ela tem, ", semanas, "semanas")*/

	//exibeNomes()

	inicio()

	for { //No go não tem while, só for
		menu()

		comandolido := leComando()

		switch comandolido {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo logs......")
		case 0:
			fmt.Println("Saindo.....")
			os.Exit(0)
		default:
			fmt.Println("Opção invalida")
			os.Exit(-1)
		}
	}

}

func inicio() {
	var nome string = "Bia"
	var sobrenome = "Miotto"
	fmt.Println("Olá", nome, sobrenome)

}
func menu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")

}
func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("A opção escolhida foi:", comando)
	return comando
}
func iniciaMonitoramento() {
	fmt.Println("Monitorando.....")

	//var sites [4]string
	/*todo array no go tem um tamanho fixo, dentro do [] deve informar o tanhanho,
	usa-se os slices ao inves de arrays pois como eles são fixos e os slices não são*/

	//sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://app.daily.dev/", "https://www.google.com.br"}
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSites(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}
func testaSites(site string) {

	resposta, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "Carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "O site não está disponivel. Status Code:", resposta.StatusCode)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites

} //nil em go é a mesma coisa que null

// Funções com multiplos retornos
/*func devolveNomeSemanas() (string, int) {
	nome := "Larissa"
	semanas := 27
	return nome, semanas
}*/

/*func exibeNomes() {
	nomes := []string{"Larissa", "Beatriz", "Eduardo"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("quantidades de itens:", len(nomes))
	fmt.Println("A capacidade é de:", cap(nomes))

	nomes = append(nomes, "Jasmine", "Docinho")
	fmt.Println("quantidades de itens:", len(nomes))
	fmt.Println("A capacidade é de:", cap(nomes))

}*/
