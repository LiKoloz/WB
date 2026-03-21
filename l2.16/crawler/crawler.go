package crawler

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html/charset"

	"l2.16/downloader"
	"l2.16/parser"
	"l2.16/saver"
	"l2.16/urlutil"
)

// Crawler управляет процессом рекурсивного скачивания сайта (зеркалирования).
type Crawler struct {
	outputDir  string
	maxDepth   int
	parallel   int
	timeout    time.Duration
	client     *http.Client
	visited    map[string]int
	visitedMu  sync.Mutex
	queue      chan *task
	wg         sync.WaitGroup
	saver      *saver.Saver
	downloader *downloader.Downloader
	parser     *parser.Parser
	ctx        context.Context
	cancel     context.CancelFunc
}

// task описывает единицу работы для воркера.
type task struct {
	url    string
	depth  int
	isPage bool
}

// NewCrawler создаёт новый экземпляр краулера с заданными параметрами
func NewCrawler(outputDir string, maxDepth, parallel int, timeout time.Duration) *Crawler {
	client := &http.Client{
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {

			if len(via) >= 10 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Crawler{
		outputDir:  outputDir,
		maxDepth:   maxDepth,
		parallel:   parallel,
		timeout:    timeout,
		client:     client,
		visited:    make(map[string]int),
		queue:      make(chan *task, 1000),
		saver:      saver.NewSaver(outputDir),
		downloader: downloader.NewDownloader(client),
		parser:     parser.NewParser(client),
		ctx:        ctx,
		cancel:     cancel,
	}
}

// Start запускает процесс зеркалирования с указанного rootURL.
func (c *Crawler) Start(rootURL string) error {

	for i := 0; i < c.parallel; i++ {
		c.wg.Add(1)
		go c.worker()
	}

	c.addTask(rootURL, 0, true)

	c.wg.Wait()
	close(c.queue)
	return nil
}

// addTask добавляет новый URL в очередь на обработку, если он ещё не был посещён
func (c *Crawler) addTask(rawURL string, depth int, isPage bool) {

	u, err := urlutil.Normalize(rawURL)
	if err != nil {

		return
	}

	if isPage && depth > c.maxDepth {
		return
	}
	c.visitedMu.Lock()
	defer c.visitedMu.Unlock()

	if oldDepth, ok := c.visited[u]; ok {

		if oldDepth <= depth {
			return
		}

	}
	c.visited[u] = depth
	select {
	case c.queue <- &task{url: u, depth: depth, isPage: isPage}:
	default:

		go func() { c.queue <- &task{url: u, depth: depth, isPage: isPage} }()
	}
}

// worker - основной метод, выполняемый воркерами (горутинами).
func (c *Crawler) worker() {
	defer c.wg.Done()
	for {
		select {
		case <-c.ctx.Done():
			return
		case t, ok := <-c.queue:
			if !ok {
				return
			}
			c.processTask(t)
		}
	}
}

// processTask обрабатывает одну задачу: загружает ресурс, сохраняет его на диск,
func (c *Crawler) processTask(t *task) {

	select {
	case <-c.ctx.Done():
		return
	default:
	}

	resp, err := c.downloader.Get(t.url)
	if err != nil {

		fmt.Printf("Ошибка загрузки %s: %v\n", t.url, err)
		return
	}
	defer resp.Body.Close()

	err = c.saver.Save(resp, t.url)
	if err != nil {
		fmt.Printf("Ошибка сохранения %s: %v\n", t.url, err)
		return
	}

	if t.isPage && t.depth < c.maxDepth {

		body, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
		if err != nil {

			body = resp.Body
		}
		links, err := c.parser.ExtractLinks(body, t.url)
		if err != nil {
			fmt.Printf("Ошибка парсинга %s: %v\n", t.url, err)
			return
		}

		for _, link := range links {

			isPage := !parser.IsResource(link)
			c.addTask(link, t.depth+1, isPage)
		}
	}
}
