### 策略模式
策略模式定义一系列算法，让这些算法在运行时可以互换，使得分离算法，符合开闭原则

设计思想：
- 1.定义一个 Operator 抽象接口;
- 2.定义一个包含属性为 Operator 抽象接口的 struct；
- 3.实现结构体自身的方法(Operate 方法中调用抽象接口)；
- 4.定义其余结构体，实现 Operator 抽象接口（实现各自的业务逻辑）；