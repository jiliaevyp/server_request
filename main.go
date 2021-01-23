package main

import (
	"errors"
	"fmt"
	"net"
)

var (
	ErrInvalidAnswer = errors.New("invalid answer server")
	ErrInvalidListen = errors.New("invalid listen server")
)

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

const (
	_network      = "tcp"
	_port         = ":4545"
	textAnswer    = "Hello, I am a server"
	requestAnswer = "I'm server. It's my answer!"
)

func main() {

	//message := textAnswer   // сообщение сервера после запуска
	listener, err := net.List(_network, _port) // тип сети и порт

	if err != nil {
		fmt.Println(ErrInvalidAnswer)
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
	fmt.Println(textAnswer, "сеть = ", _network, "порт = ", _port)
	for {
		conn, err := listener.Accept() // conn = срез байт из запроса  через _port
		if err != nil {
			fmt.Println(ErrInvalidListen)
			return
		}
		conn.Write([]byte(requestAnswer)) // ответ сервера на клиентский запрос
		conn.Close()

	}
}
