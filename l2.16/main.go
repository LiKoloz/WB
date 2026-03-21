// main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"l2.16/crawler"
)

func main() {
	var (
		startURL  string
		depth     int
		parallel  int
		timeout   time.Duration
		outputDir string
	)
	flag.StringVar(&startURL, "url", "", "Начальный URL для зеркалирования")
	flag.IntVar(&depth, "depth", 3, "Максимальная глубина рекурсии (по страницам)")
	flag.IntVar(&parallel, "parallel", 5, "Количество параллельных загрузок")
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "Таймаут HTTP-запросов")
	flag.StringVar(&outputDir, "output", "./mirror", "Директория для сохранения")
	flag.Parse()

	if startURL == "" {
		fmt.Fprintln(os.Stderr, "Необходимо указать -url")
		os.Exit(1)
	}

	parsed, err := url.Parse(startURL)
	if err != nil || !parsed.IsAbs() {
		fmt.Fprintf(os.Stderr, "Некорректный URL: %v\n", err)
		os.Exit(1)
	}
	parsed.Fragment = ""
	startURL = parsed.String()

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal("Не удалось создать выходную директорию:", err)
	}

	c := crawler.NewCrawler(outputDir, depth, parallel, timeout)
	if err := c.Start(startURL); err != nil {
		log.Fatal("Ошибка при зеркалировании:", err)
	}
	fmt.Println("Зеркалирование завершено.")
}
