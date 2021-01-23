package main

import (
	"errors"
	"fmt"
	//"net/http"
	"os"
	"strconv"
	//"net"
)

var (
	addrWeb, webPort, requestMessage, _network, _port string

	ErrInvalidAddrWeb = errors.New("invalid web server")
	ErrInvalidPort    = errors.New("invalid port number")
	ErrInvalidAnswer = errors.New("invalid answer server")
	ErrInvalidListen = errors.New("invalid listen server")
)


// ****************проверка корректности адреса web сервера *************
// *********** флаги адреса при анализе массива байт date[i]*********
// если пробел встретился - ошибка
// www = 0  счетчик www встретиться только 1 раз  первыми
// multipoint	> "hund" +1 после собаки должен быть хоть 1 символ до встречи с "."
//				если точек после собаки больше 1 ошибка
// если эти флаги установлены то адрес корректен

//func check(data []byte) int { // пока игнорируем
	//err := 0
	// анализ байт адреса
	//_network := string(data[0:4])
	//localhost := string(data[0:9])
	//fmt.Println("www=", www, "localhost=", localhost)
	//if www == "www." || localhost == "localhost" {
	//err = 0
	//} else {
	//err = 1
	//}
	//return err
//}

func inpAddrServer() (string, int) {
	var addrWeb string
	len := 256
	data := make([]byte, len)
	n, err := os.Stdin.Read(data)
	if err != nil {
		return "", 1
	} else {
		err := 0 //check(data)
		if err == 0 {
			addrWeb = string(data[0 : n-1])
			return addrWeb, 0
		}
	}
	return "", 1
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
		return ":"+ webPort, 1
	} else {
		return ":"+ webPort, 0
	}
}

func start() int {
	var komand string
	fmt.Println("-------------------------------------------------")
	fmt.Println("Сервер: ", _network + _port)
	fmt.Println("-------------------------------------------------")
	fmt.Print("Запускаю? (Y)   ", "\n")
	fmt.Scanf(
		"%s\n",
		&komand,
	)
	if komand == "Y" || komand == "y" || komand == "Н" || komand == "н" {
		return 0 // отправляем
	} else {
		return 1
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


func onserver() int {
    _end := 0
	//message := textAnswer   // сообщение сервера после запуска
	//listener, err := net.List(_network, _port) // тип сети и порт

	//if err != nil {
	//	fmt.Println(ErrInvalidAnswer)
	fmt.Scanf(
		"%s\n",
		&_end,
	)
		return _end//1
	//}
	// Вначале в функции net.Listen("tcp", ":4545") устанавливается 4545 порт для прослушивания
	//подключений по протоколу TCP.
	//После вызова этой функции сервер запущен и готов принимать подключения.
	//Затем в бесконечном цикле for получаем входящие подключения с помощью вызова listener.Accept().
	//Этот метод возвращает объект net.Conn, который представляет подключенного клиента.
	//Затем Обработка подключения.
	//Например, с помощью метода Write отправить ему сообщение.
	//Поскольку данный метод принимает срез байтов,
	//то любые сообщения надо транслировать в срез байтов: conn.Write([]byte(message))

	//defer listener.Close() // прослушивание порта _port
	//fmt.Println(textAnswer, "сеть = ", _network, "порт = ", _port)
	//for {
	//	conn, err := listener.Accept() // conn = срез байт из запроса  через _port
	//	if err != nil {
	//		fmt.Println(ErrInvalidListen)
	//		return 1
	//	}
	//	conn.Write([]byte(requestAnswer)) // ответ сервера на клиентский запрос
	//	conn.Close()

	//}
}

func main() {
	var komand string
	var (
		err int
	)
	_network      = "tcp"
	_port         = ":4545"
	textAnswer    := "Hello, I am a server"
	//requestAnswer := "I'm server. It's my answer!"
	addrWeb = ""
	//requestMessage = ""
	komand = "Y"
	for komand == "Y" || komand == "y" || komand == "Н" || komand == "н" {
		fmt.Println("------------------------------------")
		fmt.Println("|  Запуск Go server                |")
		fmt.Println("|  Запускать, не перезапускать!    |")
		fmt.Println("|                                  |")
		fmt.Println("|   (c) jiliaevyp@gmail.com        |")
		fmt.Println("------------------------------------")
		err = 1
		for err == 1 {
			fmt.Print("Введите адрес сервера:	")
			_network, err = inpAddrServer()
			if err == 1 {
				fmt.Println("Aдрес некорректен")
			}
		}
		err = 1
		for err == 1 {
			fmt.Print("Введите номер порта:	")
			_port, err = inpPort()
			if err == 1 {
				fmt.Println("Порт некорректен")
			}
		}
		err =start()
		if err == 0 {
			err = onserver()
			if err == 1 {
				fmt.Print("*** Ошибка при запуске сервера ***", "\n", "\n")
			} else {
				fmt.Print("---------------------------", "\n")
				fmt.Print("\n", textAnswer, ":   ", _network + _port, "\n", "\n")
				fmt.Print("---------------------------", "\n")
			}
		}

		fmt.Print("Продолжить? (Y)   ")
		fmt.Println("Закончить?  (Enter)")
		komand = ""
		fmt.Scanf(
			"%s\n",
			&komand,
			)
		}
	fmt.Println("Рад был с Вами пработать!")
	fmt.Print("Обращайтесь в любое время без колебаний!", "\n", "\n")
}


