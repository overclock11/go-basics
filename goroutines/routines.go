package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MyNameDelay(name string) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}

// channel es de tipo canal boleano
func MyNameDelayWithChannel(name string, channel chan bool) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(letter)
	}
	channel <- true
}
