package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {
	registraLog("site-falso", false)
	readFile()
	exibeIntroducao()
	for {
		exibeMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
			printLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Nao conheco este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	name := "Matheus"
	version := 1.1
	fmt.Println("Ola,", name)
	fmt.Println("Este programa esta na versao:", version)
}

func readCommand() int {
	var readCommand int
	fmt.Scan(&readCommand)
	fmt.Println("O comando escolhido foi", readCommand)

	return readCommand
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites := readFile()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func showLogs() {
	fmt.Println("Exibindo Logs...")
}

func testSite(site string) {
	resp, err := http.Get(site)
	
	if err != nil {
		fmt.Println("Ocorreu um erro")
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}

}

func readFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	reader := bufio.NewReader(file)

	for {
		linha, err := reader.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func registraLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro")
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05")+ " - " + site + " - " + strconv.FormatBool(status) + "\n")
	fmt.Println(file)
	file.Close()
}

func printLog(){
	file, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
