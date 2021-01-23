package client

import (
	"fmt"
	"io"
	"net"
	"os"
)

// client server
// сервер запущен на локальном компьтере на порте 4545
// клиент подлючается к этому адресу: net.Dial("tcp", "127.0.0.1:4545")
// серверу будет отправляться запрос
// с помощью вызова io.Copy(os.Stdout, conn) выводим полученный ответ на консоль.

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	io.Copy(os.Stdout, conn)
	fmt.Println("\nDone", "server answer!")
}
