package template_method

import "fmt"

// BehaviorInterface 变化的抽象接口
type BehaviorInterface interface {
	// 添加物品
	addObject()
}

// Drink 饮品模板
type Drink struct {
	behavior BehaviorInterface
}

// TemplateMethod 模板方法。定义相关业务逻
func (d *Drink) TemplateMethod() {
	// 烧开水
	d.boilWater()

	// 添加物品
	d.addObject()

	// 倒入开水
	d.pourBoilWater()
}

// BoilWater 烧开水
func (d *Drink) boilWater() {
	fmt.Println("开始烧开水")
}

// AddObject 放入物品
func (d *Drink) addObject() {
	d.behavior.addObject()
}

// PourBoilWater 倒入开水
func (d *Drink) pourBoilWater() {
	fmt.Println("向杯中倒入开水")
}

type Tea struct {
	Drink
}

func (t *Tea) addObject() {
	fmt.Println("杯中放入茶叶")
}

func NewTea() *Tea {
	instance := &Tea{}
	instance.behavior = instance
	return instance
}

type Coffee struct {
	Drink
}

func (c *Coffee) addObject() {
	fmt.Println("杯中放入咖啡")
}

func NewCoffee() *Coffee {
	instance := &Coffee{}
	instance.behavior = instance
	return instance
}
