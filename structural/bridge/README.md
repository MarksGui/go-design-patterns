## 桥接模式
桥接模式将抽象和实现解耦，让它们可以独立变化。

设计思想：
- 1.抽象接口；
- 2.实现抽象接口的类（可扩展），类属性为实现的引用；
- 3.具体的实现接口；
- 4.实现具体实现接口的类；

桥接模式和策略模式比较相似。策略模式是封装一系列算法使得算法可以互相替换，从接口上来说策略模式只有一个抽象接口。 而桥接模式是为了应对多个维度的变化，且需要将变化组合在一起来使用的场景。桥接模式会定义两个抽象接口（一个实现、一个抽象）；

示例里面：我们电脑有多个品牌、打印机也有多个品牌。他们都是可以独立变化。但是使用时又必须是两两组合使用。