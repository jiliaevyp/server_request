package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

mport (
"errors"
"fmt"
"io"
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
		conn.Write([]byte(requestAnswer)) // ответ сервера на клиентский запрос
		conn.Close()
		//_ready =
	}
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
	var komand, _stop string
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
	go _server()
	for errServ == 0 {
		fmt.Println("Для останова нажмите любую клавишу")
		fmt.Scanf(
			"%s\n",
			&_stop)
		fmt.Println("Сервер остановлен ")
		fmt.Println("\n", "Рад был с Вами пработать!")
		fmt.Print("Обращайтесь в любое время без колебаний!", "\n", "\n")
		return
	}
		fmt.Println(ErrInvalidServerListen)
		fmt.Println("\n", "Выход по ошибке!")
		fmt.Println("\n", "Рад был с Вами пработать!")
		fmt.Print("Обращайтесь в любое время без колебаний!", "\n", "\n")
	}
}

