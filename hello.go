package main

import (
	"fmt"
	"reflect"
)

func main() {

	var nome string = "Bia"
	//nas variaveis: declara o nome e depois coloca o tipo.

	var sobrenome = "Miotto" //Mas você não precisa declarar o tipo da variavel se não quiser.

	var idade int = 33
	/*no tipo int se não for atribuido um valor para a variavel ele retorna zero, o mesmo para o tipo float ele retorna 0.0,
	para toda variavel que não for atribuida valor ela sempre retorna vazia.*/

	var versao float32 = 1.1
	//tipo flutuante tem dois tipo 32 e 64.

	ano := 2023 // podemos simplificar a declaração de uma variavel utilizando := não necessitando utilizar var no inicio.

	fmt.Println("Hello", nome, sobrenome, "sabemos que sua idade é de", idade) //para concatenar usa-se virgula
	fmt.Println("Nosso progarama está na versão", versao, "ano de", ano)
	//se não utilizar uma varivel ele gera um erro " declared and not used", ele não deixa declarar uma variavel e não usar.
	fmt.Println("O tipo da varivel sobrenome é: ", reflect.TypeOf(sobrenome)) //Para descobrir o tipo de varivel usa-se o reflect.Typeof.

	/*//Menu
	fmt.Println("Olá", nome, sobrenome)
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")

	var comando int
	fmt.Scan(&comando) //O & é como se fosse o endereço da varivavel, necessario pois uma variavel declarada sem valor retorna zero ou vazia por padrão
	fmt.Println("A opção escolhida foi:", comando)*/
}
