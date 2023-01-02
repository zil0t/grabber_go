package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/opesun/goquery"
)

var (
	WORKERS       int    = 2            //кол-во "потоков"
	REPORT_PERIOD int    = 10           //частота отчетов (сек)
	DUP_TO_STOP   int    = 500          //максимум повторов до останова
	HASH_FILE     string = "hash.bin"   //файл с хешами
	QUOTES_FILE   string = "quotes.txt" //файл с цитатами
)

func init() {
	//Задаем правила разбора:
	flag.IntVar(&WORKERS, "w", WORKERS, "количество потоков")
	flag.IntVar(&REPORT_PERIOD, "r", REPORT_PERIOD, "частота отчетов (сек)")
	flag.IntVar(&DUP_TO_STOP, "d", DUP_TO_STOP, "кол-во дубликатов для остановки")
	flag.StringVar(&HASH_FILE, "hf", HASH_FILE, "файл хешей")
	flag.StringVar(&QUOTES_FILE, "qf", QUOTES_FILE, "файл записей")
	//И запускаем разбор аргументов
	flag.Parse()
}
func grab() <-chan string { //функция вернет канал, из которого мы будем читать данные типа string
	c := make(chan string)
	for i := 0; i < WORKERS; i++ { //в цикле создадим нужное нам количество гоурутин - worker'oв
		go func() {
			for { //в вечном цикле собираем данные
				x, err := goquery.ParseUrl("http://vpustotu.ru/moderation/")
				if err == nil {
					if s := strings.TrimSpace(x.Find(".fi_text").Text()); s != "" {
						c <- s //и отправляем их в канал
					}
				}
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}
	fmt.Println("Запущено потоков: ", WORKERS)
	return c
}

func main() {
	quote_chan := grab()
	for i := 0; i < 5; i++ { //получаем 5 цитат и закругляемся
		fmt.Println(<-quote_chan)
	}
}
