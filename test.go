package main

import "fmt"

type Duck interface {
	performFly()
	performQuack()
	display()
}

type RedDuck struct {
	FlyBehavior   FlyBehavior
	QuackBehavior QuackBehavior
}

type FlyBehavior interface {
	Fly()
}

type QuackBehavior interface {
	Quack()
}

type FlyWithWings struct {}

type FlyNoWay struct {
}

type Squeak struct {
}



func (d *RedDuck) display() {
	fmt.Println("这是红鸭子")
}

func (d *RedDuck) performFly() {
	d.FlyBehavior.Fly()
}

func (d *RedDuck) performQuack() {
	d.QuackBehavior.Quack()
}

func (f *FlyWithWings) Fly() {
	fmt.Println("会飞")
}
func (f *FlyNoWay) Fly() {
	fmt.Println("不会飞")
}
func (f *Squeak) Quack() {
	fmt.Println("吱吱叫")
}

func main() {
	d := &RedDuck{
		FlyBehavior: &FlyWithWings{},
		QuackBehavior:&Squeak{},
	}
	d.performFly()
	d.performQuack()
	d.display()
}
