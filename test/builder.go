package test

type Builder interface {
	createVehicle()
	addWheel()
	addEngine()
	addDoors()
	getVehicle() *Vehicle
}

type Vehicle struct {
	data map[string]string
}

func (v *Vehicle) setPart(key string, value string) {
	//v.data = make(map[string]string)
	//v.data = map[string]string{key:value}
	v.data[key] = value
	//fmt.Println(v.data)
}

type CarBuilder struct {
	car *Vehicle
}

func (c *CarBuilder) createVehicle() {
	c.car = &Vehicle{}
	c.car.data = make(map[string]string)
}

func (c *CarBuilder) addDoors() {
	c.car.setPart("leftDoor", "门对象")
	c.car.setPart("rightDoor", "门对象")
}

func (c *CarBuilder) addWheel() {
	c.car.setPart("wheelLF", "轮毂对象")
	c.car.setPart("wheelRF", "轮毂对象")
}

func (c *CarBuilder) addEngine() {
	c.car.setPart("engine", "engine对象")
}

func (c *CarBuilder) getVehicle() *Vehicle {
	return c.car
}

type Director struct {
}

func (c *Director) build(builder Builder) *Vehicle {
	builder.createVehicle()
	builder.addDoors()
	builder.addEngine()
	builder.addWheel()
	return builder.getVehicle()
}

func NewBuilder() *Vehicle {
	director := Director{}
	return director.build(&CarBuilder{})
}

//func main() {
//	//car := &CarBuilder{}
//	//director := Director{}
//	//a := director.build(car)
//	//fmt.Println(a)
//	car := NewBuilder()
//	fmt.Println(car)
//}
