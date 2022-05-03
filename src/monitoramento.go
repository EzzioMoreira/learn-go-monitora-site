package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 3

func main() {
	exibeIntroducao()
	lerSitesDoArquivo()
	for {
		exibeMenu()
		comandoLido := lerComando()

		switch comandoLido {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Printf("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do comando...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando.")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Ezzio Moreira"
	versao := 0.1
	fmt.Println("Olá, Sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento:")
	fmt.Println("2- Exibir Logs:")
	fmt.Println("0- Sair do Programa:")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando digitado foi:", comandoLido)

	return comandoLido
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://go.dev/", "https://github.com/"}

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site:", i, ":", site)
			testaSite(site)
		}
		time.Sleep(monitoramento * time.Second)
	}
	fmt.Println("")
	fmt.Println("#############################")
}

func testaSite(site string) {
	//Capturando o segundo parâmetro err
	resp, err := http.Get(site)

	if err != nil {
		//Verificando se houve erro: se err for diferente de nill > print
		println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Print("Site", site, "esta com problema. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
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
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format(time.RFC3339) + site + "- online  - " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("logs.log")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
