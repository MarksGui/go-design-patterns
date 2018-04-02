package car

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive()
	Stop()
}

type Car struct {
	color Color
	wheel Wheels
	speed Speed
}

func (this *Car) Paint(color Color) *Car{
	return this.Color(color)
}

func (this *Car) Color(color Color) *Car{
	this.color = color
	return this
}
func (this *Car) Wheels(wheel Wheels) *Car{
	this.wheel = wheel
	return this
}
func (this *Car) TopSpeed(speed Speed) *Car{
	this.speed = speed
	return this
}

func (this *Car) Drive() {
	fmt.Println(string(this.color) + "颜色的车," + string(this.wheel) +"轮毂,发车了")
}
func (this *Car) Stop() {
	fmt.Println(string(this.color) + "颜色的车," + string(this.wheel) +"轮毂,停车了")
}

func (this *Car) Build() Interface{
	return this
}

func NewBuilder() *Car{
	return &Car{}
}

