### 简单工厂模式

当我们有很多类需要实例化，并且可能未来需要扩展。我们就可以使用简单工厂模式。
这里要注意简单工厂模式适合创建类很简单的，比如就是一行代码：NewJsonParser()。

例子中我们需要针对不同的配置文件，然后解析到同一个结构体对象上。
