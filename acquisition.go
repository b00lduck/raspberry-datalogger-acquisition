package main

import (
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
	"time"
	"github.com/b00lduck/raspberry-datalogger-dataservice-client"
	"io/ioutil"
)


func main() {

    d1 := []byte("17\n")
    ioutil.WriteFile("/sys/class/gpio/unexport", d1, 0644)

    if err := embd.InitGPIO(); err != nil {
		panic(err)
    }
    defer embd.CloseGPIO()

    led, err := embd.NewDigitalPin(17)
    if err != nil {
		panic(err)
    }
    defer led.Close()
    if err := led.SetDirection(embd.In); err != nil {
		panic(err)
    }

	state := false
	count := 0

    for {
        pin,err := led.Read()
		if  err != nil {
			panic(err)
		}

		if pin == 1 {
			if state == false {
				count++
			}
		} else {
			count = 0
			state = false
		}

		if count >= 3 {
			state = true
			count = 0
			client.SendCounterTick("gas_usage")
		}

		time.Sleep(10 * time.Millisecond)

    }
}

