package kata

import (
	"fmt"
	"strings"
)

func TakeOrders(hear <-chan string, say chan<- string, done <-chan bool) {
	menu := GetMenu()
	var total float64
	for {
		select {
		case <-done:
			return
		case msg := <-hear:
			switch {
			case msg == "place an order":
				total = 0
				say <- "ok"
			case msg == "that is all":
				say <- fmt.Sprintf("%.2f", total)
			case msg == "goodbye":
				say <- "goodbye"
			case strings.HasPrefix(msg, "price "):
				item := strings.TrimPrefix(msg, "price ")
				if p, ok := menu[item]; ok {
					say <- fmt.Sprintf("%.2f", p)
				} else {
					say <- "unavailable"
				}
			case strings.HasPrefix(msg, "order "):
				item := strings.TrimPrefix(msg, "order ")
				if p, ok := menu[item]; ok {
					total += float64(p)
					say <- "ok"
				} else {
					say <- "unavailable"
				}
			}
		}
	}
}