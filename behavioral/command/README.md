## 命令模式

命令模式本质是把某个对象的方法调用封装到对象中，方便传递、存储、调用。

除了配置灵活外，使用命令模式还可以用作：
* 批处理
* 任务队列
* undo, redo
等等把具体命令封装到对象中使用的场合。

设计思想：
- 1.命令接口 Command；
- 2.接收者 Receiver。执行请求相关的操作 Execute()；
- 3.具体命令的结构体（CommandAttack、CommandEscape），属性为嵌入的 Receiver；调用 Receiver 方法，实现 Command 接口。
- 4.调用者 Invoker。用于添加命令、执行命令；