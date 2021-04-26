package Programming

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

//Make the connection to the IRC server and login
func Connect() net.Conn {
	conn, err := net.Dial("tcp", Target+":"+strconv.Itoa(Port))
	if err != nil {
		panic(err)
	}

	//mi loggo
	SendMessage(conn, "USER "+Username+" 8 * :Someone")
	SendMessage(conn, "NICK "+Nick)

	return conn
}

//se premo ctrl+c, chiudo la connessione - INUTILE
func CleanDisconnect(conn net.Conn) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		conn.Close()
	}()
}

//a quanto pare è buona norma salutare prima di chiudere
//la connessione.
func CleanClose(conn net.Conn) {
	SendMessage(conn, "QUIT Bye")
	conn.Close()
}

func ReadWellcomeMessage(conn net.Conn) {
	tp := textproto.NewReader(bufio.NewReader(conn))
	for {
		status, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}
		fmt.Println(status)
		//ogni tanto devo rispondere al ping
		if strings.HasPrefix(status, "PING") {
			SendMessage(conn, "PONG")
		}
		//fine del benvenuto
		if strings.HasPrefix(status, ":"+Nick+"!"+Username) {
			break
		}
	}
}

func GetBotList(conn net.Conn) []string {
	//guardo quali bot sono disponibili
	SendMessage(conn, "BS BOTLIST")
	//creo la lista dei bot
	var botList []string
	tp := textproto.NewReader(bufio.NewReader(conn))
	for {
		status, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}
		//fine della lista
		if strings.Contains(status, " bots disponibles.") {
			break
		}
		//estraggo il nome del bot e lo aggiungo alla lista
		words := strings.Fields(status)
		botList = append(botList, words[4])
	}
	//il primo elemento fa parte del messaggio "adesso ti stampo la lista dei bot"
	return botList[1:]
}

func SendMessage(conn net.Conn, message string) {
	fmt.Fprintf(conn, message+"\r\n")
}

func ReceiveOneLineMessage(conn net.Conn) string {
	tp := textproto.NewReader(bufio.NewReader(conn))
	message, err := tp.ReadLine()
	if err != nil {
		panic(err)
	}
	return message
}
