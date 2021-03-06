package main

import (
	"errors"
	"fmt"
	//"io"
	"net"
	"os"
	"strconv"
)

var (
	addrWeb, webPort, requestMessage, _network, _adrServer, _port string
	errServ                                                       int
	ErrInvalidTypeNetwork                                         = errors.New("invalid type network")
	ErrInvalidPort                                                = errors.New("invalid port number")
	ErrInvalidAnswerServer                                        = errors.New("invalid answer server")
	ErrInvalidServerListen                                        = errors.New("invalid listen server")
)

const (
	answerServer  = "Hello, I am a server."
	readyServer   = "I'm ready!"
	requestAnswer = "Server have request."
)

func inpNetwork() (string, int) {
	var typNet string
	len := 256
	data := make([]byte, len)
	n, err := os.Stdin.Read(data)
	typNet = string(data[0 : n-1])
	if err != nil || typNet != "tcp" {
		return typNet, 1
	} else {
		return typNet, 0
	}
}

func inpPort() (string, int) {
	var (
		webPort string
		res     float64
	)
	fmt.Scanf(
		"%s\n",
		&webPort,
	)
	res, err := strconv.ParseFloat(webPort, 16)
	res = res + 1
	if err != nil || len(webPort) != 4 {
		return ":" + webPort, 1
	} else {
		return ":" + webPort, 0
	}
}

// func Listen(network, laddr string) (net.Listener, error)
// network - протокол, по которому приложение будет получать запросы,
//и laddr представляет локальный адрес, по которому будет запускаться сервер.
//Протокол должен представлять одно из значений: "tcp", "tcp4", "tcp6", "unix", "unixpacket".
//Локальный адрес может содержать только номер порта, например, ":8080".
//В этом случае приложение будет обслуживать по всем.
//
// Accept() (принимает входящее подключение) и Close() (закрывает подключение)
//В случае успешного выполнения функция возвращает объект интерфейса net.Listener,
//который представляет функционал для приема входящих подключений.
//В зависимости от типа используемого протокола возвращаемый объект
//Listener может представлять тип net.TCPListener или net.UnixListener
//(оба этих типа реализуют интерфейс net.Listener).

func _server() {
	errServ = 0
	listener, err := net.Listen(_network, _port) // тип сети и порт
	if err != nil {
		fmt.Println(ErrInvalidAnswerServer)
		errServ = 1
		return
	}
	// Вначале в функции net.Listen("tcp", ":4545") устанавливается 4545 порт для прослушивания
	//подключений по протоколу TCP.
	//После вызова этой функции сервер запущен и готов принимать подключения.
	//Затем в бесконечном цикле for получаем входящие подключения с помощью вызова listener.Accept().
	//Этот метод возвращает объект net.Conn, который представляет подключенного клиента.
	//Затем Обработка подключения.
	//Например, с помощью метода Write отправить ему сообщение.
	//Поскольку данный метод принимает срез байтов,
	//то любые сообщения надо транслировать в срез байтов: conn.Write([]byte(message))

	defer listener.Close() // прослушивание порта _port
	fmt.Println(answerServer, " network=", _network, _port)
	fmt.Println(readyServer)
	for {
		conn, err := listener.Accept() // conn = срез байт из запроса  через _port
		if err != nil {
			fmt.Println(ErrInvalidServerListen)
			errServ = 1
			return
		}
		go handleConnect(conn) // запуск обработчика запросов сервера
		//conn.Write([]byte(requestAnswer)) // ответ сервера на клиентский запрос
		//conn.Close()

	}
}

// Обработчик запросов сервера
func handleConnect(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, (1024 * 4)) // считываем полученные в запросе байты данных
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		requestText := string(input[0:n])
		answerText := "На запрос " + requestText + "---> Ответ сервера: Все ништяк!"
		conn.Write([]byte(answerText))
	}
}

// client server
// сервер запущен на локальном компьтере на порте 4545
// клиент подлючается к этому адресу: net.Dial("tcp", ":4545")
// серверу будет отправляться запрос
// с помощью вызова io.Copy(os.Stdout, conn) выводим полученный ответ на консоль.

func _client() int {
	conn, err := net.Dial(_network, _port)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	defer conn.Close()
	for {
		var source string
		fmt.Print("Введи запрос серверу:" + "    (end - выход)")
		fmt.Scanf(
			"%s\n",
			&source,
		)
		if source == "end" || source == "утв" {
			fmt.Print("Конец работы:")
			return 1
		}
		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return 0
		}
		fmt.Println("Ответ сервера: ")
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		fmt.Println(string(buff[0:n]))
	}
	//io.Copy(os.Stdout, conn)
	//fmt.Println("\nDone", "server answer!", "\n")
	return 1
}

func _beg() {
	fmt.Println("------------------------------------")
	fmt.Println("|  Запуск Go server                |")
	fmt.Println("|  Запускать, не перезапускать!    |")
	fmt.Println("|                                  |")
	fmt.Println("|   (c) jiliaevyp@gmail.com        |")
	fmt.Println("------------------------------------")
}

func main() {
	var komand string
	err := 1
	_beg() // заголовок
	for err == 1 {
		fmt.Print("Введите тип сети:	")
		_network, err = inpNetwork()
		if err == 1 {
			fmt.Println(ErrInvalidTypeNetwork)
		}
	}
	err = 1
	for err == 1 {
		fmt.Print("Введите номер порта:	")
		_port, err = inpPort()
		if err == 1 {
			fmt.Println(ErrInvalidPort)
		}
	}
	go _server() // запуск сервера
	if errServ == 0 {
		komand = "Y"
		for komand == "Y" || komand == "y" || komand == "Н" || komand == "н" {
			fmt.Println("Сделать запрос? (Y)   ")
			fmt.Scanf(
				"%s\n",
				&komand,
			)
			fmt.Println()
			if komand == "Y" || komand == "y" || komand == "Н" || komand == "н" {
				err = _client()
			}
			if err == 1 {
				fmt.Println("\n", "Рад был с Вами пработать!")
				fmt.Print("Обращайтесь в любое время без колебаний!", "\n", "\n")
			}
		}
	} else {
		fmt.Println(ErrInvalidServerListen)
	}
}
