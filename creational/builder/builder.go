package builder

// 抽象建造者
type Builder interface {
	CreateVehicle()
	AddWheel()
	AddEngine()
	AddDoors()
	GetVehicle() *Vehicle
}

// 父 struct
type Vehicle struct {
	Wheel  string
	Engine string
	Doors  string
}

// 装饰类
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Build 方法中可以做一些参数的校验之类的事情
func (d *Director) Construct() *Vehicle {
	d.builder.CreateVehicle()
	d.builder.AddWheel()
	d.builder.AddEngine()
	d.builder.AddDoors()

	return d.builder.GetVehicle()
}

// 卡车
type TruckBuilder struct {
	vehicle *Vehicle
}

func (t *TruckBuilder) CreateVehicle() {
	t.vehicle = &Vehicle{}
}

func (t *TruckBuilder) AddWheel() {
	t.vehicle.Wheel = "四个卡车轮子"
}

func (t *TruckBuilder) AddEngine() {
	t.vehicle.Engine = "一个卡车发动机"
}

func (t *TruckBuilder) AddDoors() {
	t.vehicle.Doors = "两个卡车门"
}
func (t *TruckBuilder) GetVehicle() *Vehicle {
	return t.vehicle
}

// 轿车
type CarBuilder struct {
	vehicle *Vehicle
}

func (c *CarBuilder) CreateVehicle() {
	c.vehicle = &Vehicle{}
}

func (c *CarBuilder) AddWheel() {
	c.vehicle.Wheel = "四个轿车轮子"
}

func (c *CarBuilder) AddEngine() {
	c.vehicle.Engine = "一个轿车发动机"
}

func (c *CarBuilder) AddDoors() {
	c.vehicle.Doors = "四个轿车车门"
}

func (c *CarBuilder) GetVehicle() *Vehicle {
	return c.vehicle
}
