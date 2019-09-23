package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", blink)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func blink (c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Use gpioreg GPIO pin registry to find a GPIO pin by name.
	p := gpioreg.ByName("GPIO6")
	if p == nil {
		log.Fatal("Failed to find GPIO6")
	}

	// Set the pin as output High.
	for i := 0; i < 5; i ++ {
		if err := p.Out(gpio.High); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second / 5)
		if err := p.Out(gpio.Low); err != nil {
			log.Fatal(err)
		}
	}
}