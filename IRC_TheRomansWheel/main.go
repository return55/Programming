package main

import (
	"Programming"

	"fmt"
)

func main() {
	conn := Programming.Connect()
	defer Programming.CleanClose(conn)

	//leggo il messaggio di benvenuto
	Programming.ReadWellcomeMessage(conn)

	fmt.Println(Programming.GetBotList(conn))

	//inizio la conversazione col bot della sfida
	Programming.SendMessage(conn, "PRIVMSG "+Programming.BotName+" :!ep3")

	//leggo la risposta
	message := Programming.ReceiveOneLineMessage(conn)

	//decodifico ROT13
	word = 

}
