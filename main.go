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

	directions := []uint16{0, 60, 120, 180, 240, 300}
	i := 0

	adaptor := sphero.NewAdaptor("/dev/rfcomm1")
	driver := sphero.NewSpheroDriver(adaptor)

	work := func() {
		gobot.Every(3*time.Second, func() {
			dir := directions[i%len(directions)]
			driver.Roll(200, dir)
			c, colorName := getRandomColor()

			driver.SetRGB(c.R, c.G, c.B)
			log.Printf("Direction: %d R: %d G: %d B: %d - %s", dir, c.R, c.G, c.B, colorName)
			i++
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
