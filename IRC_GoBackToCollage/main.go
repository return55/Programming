package main

import (
	"Programming"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	conn := Programming.Connect()
	defer Programming.CleanClose(conn)

	//leggo il messaggio di benvenuto
	Programming.ReadWellcomeMessage(conn)

	//inizio la conversazione col bot della sfida
	Programming.SendMessage(conn, "PRIVMSG "+Programming.BotName+" :!ep1")

	//leggo la risposta
	message := Programming.ReceiveOneLineMessage(conn)

	//estraggo i numeri
	words := strings.Fields(message)
	num1, err := strconv.Atoi(strings.ReplaceAll(words[3], ":", ""))
	if err != nil {
		panic(err)
	}
	num2, err := strconv.Atoi(words[5])
	//calcoli
	res := math.Sqrt(float64(num1)) * float64(num2)
	//invio il risultato
	message = fmt.Sprintf("PRIVMSG "+Programming.BotName+" :!ep1 -rep %.2f\r\n", res)
	Programming.SendMessage(conn, message)

	//leggo la flag
	fmt.Println(Programming.ReceiveOneLineMessage(conn))

}
