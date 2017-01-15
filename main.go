package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"golang.org/x/image/colornames"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/sphero"
)

func getRandomColor() (color.RGBA, string) {
	n := rand.Intn(len(colornames.Names))
	c := colornames.Names[n]

	return colornames.Map[c], c
}

func main() {

	adaptor := sphero.NewAdaptor("/dev/rfcomm0")
	driver := sphero.NewSpheroDriver(adaptor)

	work := func() {
		gobot.Every(3*time.Second, func() {
			driver.Roll(150, uint16(gobot.Rand(360)))
			c, colorName := getRandomColor()

			driver.SetRGB(c.R, c.G, c.B)
			log.Printf("R: %d G: %d B: %d - %s", c.R, c.G, c.B, colorName)
		})
	}

	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
		work,
	)

	robot.Start()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
