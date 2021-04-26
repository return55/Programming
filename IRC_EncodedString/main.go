package main

import (
	"Programming"
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	conn := Programming.Connect()
	defer Programming.CleanClose(conn)

	//leggo il messaggio di benvenuto
	Programming.ReadWellcomeMessage(conn)

	//inizio la conversazione col bot della sfida
	Programming.SendMessage(conn, "PRIVMSG "+Programming.BotName+" :!ep2")

	//leggo la risposta
	message := Programming.ReceiveOneLineMessage(conn)

	//estraggo la stringa encodata
	words := strings.Fields(message)
	encodedString := strings.ReplaceAll(words[3], ":", "")
	//decodifico da base 64
	decodedString, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	//invio la risposta
	Programming.SendMessage(conn, "PRIVMSG "+Programming.BotName+" :!ep2 -rep "+string(decodedString))

	//leggo la flag
	fmt.Println(Programming.ReceiveOneLineMessage(conn))

}
