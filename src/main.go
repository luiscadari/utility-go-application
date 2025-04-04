package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

    introduction()
    fmt.Println("Escolha uma das opções acima:")
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

func monitore(){
    var url string
    fmt.Println("Digite a URL que deseja monitorar:")
    fmt.Scan(&url)
    resp, err := http.Get(url)
    fmt.Println(err)
    fmt.Println("Status Code:", resp.StatusCode);
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