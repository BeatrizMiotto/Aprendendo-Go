package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	/*nome, semanas := devolveNomeSemanas() //caso não queira retornar um dos valores basta colocar o "_" exemplo "_,semans :=devolveNomeSemanas()".
	fmt.Println(nome, "ela tem, ", semanas, "semanas")*/

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
	fmt.Println("Melhorando.....")
	site := "https://www.alura.com.br"
	resposta, _ := http.Get(site)
	//fmt.Println(resposta)
	if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "Carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "O site não está disponivel. Status Code:", resposta.StatusCode)
	}

}

// Funções com multiplos retornos
/*func devolveNomeSemanas() (string, int) {
	nome := "Larissa"
	semanas := 27
	return nome, semanas
}*/
