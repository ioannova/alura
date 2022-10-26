package main

import (
	"fmt"
	// "reflect"
	"os"
	"net/http"
	"time"
	"bufio"
	"io"
	"strings"
	// "strconv"
)
const delay = 1
const loop = 5

func main(){
	exibitIntroducao()
	for {
		switch exibitMenu() {
		case 1:
			monitoramento()
		case 2:
			logs()
		case 0:
			fmt.Println("Encerrando.")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido!")
			os.Exit(-1)
		}
	}

}

func exibitIntroducao() {
	var name string = "Nelson"
	var	age int = 34
	var version float32 = 1.1
	fmt.Println("Olá, Sr.", name, "sua idade é", age)
	fmt.Println("Este programa está na versão", version)
	// fmt.Println("O tipo da variável name é:", reflect.TypeOf(name), "\n")
}

func exibitMenu() int {
	fmt.Println("1 - Monitoramento")
	fmt.Println("2 - Exibir os Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Print("\nComando: ")
	fmt.Scan(&comando)
	// fmt.Println("O ponteiro da variável comando é:", &comando)
	// fmt.Println("O comando escolhido:", comando, "\n")

	return comando
}

func monitoramento(){
	fmt.Println("\nInicitando monitoramento...\n")
	sites := loadSites()
	for x:=0; x<loop; x++ {
		fmt.Println("----- loop", x)
		for _, site := range sites {
			pingSite(site)
		}
		time.Sleep(delay * time.Second)
	}
	fmt.Println("\n")
}

func pingSite(site string)  {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro:", err)
	}
	if res.StatusCode == 200 {
		fmt.Println("online  :", site)
		writeLog(site, true)
	} else {
		fmt.Println("offline :", site, " - status code:", res.StatusCode)
		writeLog(site, false)
	}
}

func loadSites() []string {
	var sites []string
	file, err :=  os.Open("sites.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Erro:", err)
	}

	line := bufio.NewReader(file)

	for {
		site, err := line.ReadString('\n')
		site = strings.TrimSpace(site)
		if site == ""{
			break
		}
		sites = append(sites, site)
		// fmt.Println(site)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Erro:", err)
		}
	}
	fmt.Println(sites)
	return sites
}

func writeLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer arquivo.Close()

	if err != nil {
		fmt.Println(err)
	}
	if status == true {
		// arquivo.WriteString("[" + time.Now().Format("02/01/2006 15:04:05Hs") + "] online (" + strconv.FormatBool(status) + ") " + site + "\n")
		arquivo.WriteString("[" + time.Now().Format("02/01/2006 15:04:05Hs") + "] offline " + site + "\n")
	}
	fmt.Println(arquivo)
}

func logs(){
	fmt.Println("Acessando Logs...")
}
