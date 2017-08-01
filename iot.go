package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"gobot.io/x/gobot/platforms/raspi"
)

type Iot struct {
	Dongle *raspi.Adaptor
}

func main() {
	iot := iris.New()
	r := raspi.NewAdaptor()
	i := newIOT(r)
	iot.Get("/", i.toggleFunc)
	iot.Get("/on", i.keepOn)
	iot.Get("/off", i.keepOff)
	/* app will be running on 8080 port*/
	iot.Run(iris.Addr(":8080"))
}

func newIOT(c *raspi.Adaptor) Iot {
	return Iot{Dongle: c}
}

/* this function puts the iot into toggle mode
where the led keeps blinking with a time delay*/
func (iot *Iot) toggleFunc(c context.Context) {}

/* this one puts the iot in constant on mode*/
func (iot *Iot) keepOn(c context.Context) {
	fmt.Println("on")
	iot.Dongle.DigitalWrite("7", byte(1))
}

/* this one puts the iot in constant off mode*/
func (iot *Iot) keepOff(c context.Context) {
	fmt.Println("off")
	iot.Dongle.DigitalWrite("7", byte(0))
}
