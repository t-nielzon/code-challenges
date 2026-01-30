package kata

import (
	"fmt"
	"strings"
)

func Solution(hear <-chan string, speak chan<- string, done <-chan bool) {
	menu := GetMenu()
	var total float32

	for {
		select {
		case <-done:
			return
		case msg := <-hear:
			switch {
			case msg == "place an order":
				total = 0
				speak <- "ok"

			case strings.HasPrefix(msg, "price "):
				item := strings.TrimPrefix(msg, "price ")
				if price, ok := menu[item]; ok {
					speak <- formatPrice(price)
				} else {
					speak <- "unavailable"
				}

			case strings.HasPrefix(msg, "order "):
				item := strings.TrimPrefix(msg, "order ")
				if price, ok := menu[item]; ok {
					total += price
					speak <- "ok"
				} else {
					speak <- "unavailable"
				}

			case msg == "that is all":
				speak <- formatPrice(total)

			case msg == "goodbye":
				speak <- "goodbye"
			}
		}
	}
}

func formatPrice(price float32) string {
	return fmt.Sprintf("%.2f", price)
}