package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			// Проверяем, не EOF ли это (Ctrl+D в Unix, Ctrl+Z в Windows)
			if err.Error() == "EOF" {
				fmt.Println("Получен Ctrl+D, завершение.")
				os.Exit(0)
			}
			fmt.Fprintf(os.Stderr, "Ошибка чтения: %v\n", err)
			os.Exit(1)
		}

		str := strings.TrimSpace(input)
		args := strings.Split(str, " ")

		// Обработка команд
		for i := 0; i < len(args); i++ {
			select {
			case <-sigChan:
				continue
			default:
				switch args[i] {
				case "cd":
					i++
					if i >= len(args) {
						fmt.Println("cd: требуется аргумент")
						continue
					}
					if err := os.Chdir(args[i]); err != nil {
						fmt.Printf("cd: %v\n", err)
					}
				case "pwd":
					dir, err := os.Getwd()
					if err != nil {
						fmt.Printf("pwd: %v\n", err)
					} else {
						fmt.Println(dir)
					}
				case "echo":
					i++
					if i < len(args) {
						fmt.Println(args[i])
					} else {
						fmt.Println()
					}
				case "kill":
					i++
					if i < len(args) {
						cmd := exec.Command("taskkill", "/F", "/PID", args[i])
						cmd.Start()
					}
				case "ps":
					cmd := exec.Command("tasklist")
					output, _ := cmd.Output()
					fmt.Println(string(output))
				case "exec":
					i++
					if i+1 < len(args) {
						cmd := exec.Command(args[i], args[i+1])
						var out bytes.Buffer
						cmd.Stdout = &out
						if err := cmd.Run(); err != nil {
							fmt.Printf("exec: %v\n", err)
						} else {
							fmt.Print(out.String())
						}
						i++
					} else {
						fmt.Println("exec: требуется команда и аргумент")
					}
				default:
					fmt.Printf("%s: команда не найдена\n", args[i])
				}
			}
		}
	}

}
