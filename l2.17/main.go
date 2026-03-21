package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	// Создаем reader для чтения из стандартного ввода
	reader := bufio.NewReader(os.Stdin)

	// Читаем до символа новой строки
	input, _ := reader.ReadString('\n')

	// Удаляем символ новой строки
	str := strings.TrimSpace(input)

	a := strings.Split(str, " ")
	timeout := 10

	if a[len(a)-2] == "--timeout" {
		timeout, _ = strconv.Atoi(a[len(a)-1])
		a = a[:len(a)-2]
	}
	var (
		path = a[0]
		port = a[1]
	)
	listener, err := net.Listen("tcp", path+":"+port)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
	defer listener.Close()

	tcpListener := listener.(*net.TCPListener)
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	tcpListener.SetDeadline(deadline)
	var wg sync.WaitGroup

	stop := make(chan struct{})

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
		}
		close(stop)
	}()
	select {
	case <-stop:
		fmt.Println("Получен Ctrl+D, завершение.")
		return
	default:
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка соединения:", err)
			continue
		}

		wg.Add(2)
		go func() {
			defer wg.Done()

			for {
				buffer := make([]byte, 1024)
				n, err := conn.Read(buffer)
				if err != nil {
					fmt.Println("Ошибка чтения:", err)
					return
				}
				fmt.Println("Получено:", string(buffer[:n]))
			}
		}()

		go func() {

			defer wg.Done()
			for {
				// Создаем reader для чтения из стандартного ввода
				reader2 := bufio.NewReader(os.Stdin)

				// Читаем до символа новой строки
				input2, _ := reader2.ReadString('\n')

				// Удаляем символ новой строки
				str2 := strings.TrimSpace(input2)
				conn.Write([]byte(str2))
			}
		}()
	}
	wg.Wait()
}
