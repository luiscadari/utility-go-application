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

const monitoramentos =  3 //definindo constantes para facilitar configuração
const delay = 5 * time.Second

func main() {
    for {
        introduction()
        command := inputCommand();
        switch command {
        case 0: {
            fmt.Println("Saindo do programa...")
            os.Exit(0)
        }
        case 1: {
            fmt.Println("Iniciando o monitoramento...")
            monitore()
        }
        case 2: {
            printLog()
        }
        default: {
            fmt.Println("Comando inválido!")   
            os.Exit(-1)
        }
        }
    }


}

func monitore(){
    sites := fileInput() // slice
    fmt.Println("sites que foram encontrados: ", sites)
    fmt.Println("Serão realizados ", monitoramentos, " testes a cada ", delay);
    for i := 0; i < monitoramentos; i++ {
        fmt.Println("Iniciando ", (i + 1), "teste...")
        for index, _ := range sites {
            fetchWeb(index, sites[index])
        }
        time.Sleep(delay)
    }

}

func fetchWeb(index int ,site string){
    resp, err := http.Get(site)
    if err != nil {
        fmt.Println("Ocorreu um erro: ", err)
    }
    if resp.StatusCode == 200 {
        fmt.Println(index," - ", "O site: ", site, "está no ar!");
        createLog(site, true)
    }else{
        fmt.Println("O site está fora do ar!");
        createLog(site, false)
    }
}

func introduction() {
    version := 1.1
    fmt.Println("Este programa está na versão", version)
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}

func inputCommand() int {
    var command int
    fmt.Println("Escolha uma das opções acima:")
    fmt.Scan(&command)
    return command
}

func fileInput()[]string{
    var sites []string

    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)
    for {
        linha, err := leitor.ReadString('\n')
        if err == io.EOF {
            break
        }
        linha = strings.TrimSpace(linha)

        sites = append(sites, linha)



    }

    arquivo.Close()
    return sites
}

func createLog(site string, status bool){
    arquivo, err := os.OpenFile("logs.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil{
        fmt.Println(err)
    }
    println(arquivo)
    arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05 - ") + site + " - online:" + strconv.FormatBool(status) + "\n")
    arquivo.Close()
}

func printLog(){
    arquivo, err := os.ReadFile("logs.txt");
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(arquivo))
}