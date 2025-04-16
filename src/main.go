package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos =  3 //definindo constantes para facilitar configuração
const delay = 5 * time.Minute

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
    fmt.Println("Serão realizados ", monitoramentos, " testes a cada ", delay);
    for i := 0; i < monitoramentos; i++ {
        fmt.Println("Iniciando ", (i + 1), "teste...")
        for index, site := range sites {
            fetchWeb(index, site)
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
    }else{
        fmt.Println("O site está fora do ar!");
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
        fmt.Println("Ocorreu um erro: ", err)
    }
    fmt.Println(arquivo)
    return sites
}