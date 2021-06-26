## golang-design-pattern


### 设计原则

做什么事都需要遵循一些准则，设计模式也不例外。我们在设计一些设计模式时，一般遵循如下6项基本原则，它们分别是:

* 单一职责原则（Single Responsibility Principle）
* 开闭原则（Open Closed Principle）
* 里氏替换原则（Liskov Substitution Principle）
* 迪米特法则（Law of Demeter），又叫“最少知道法则”
* 接口隔离原则（Interface Segregation Principle）
* 依赖倒置原则（Dependence Inversion Principle）

把这 6 个原则的首字母（里氏替换原则和迪米特法则的首字母重复，只取一个）联合起来就是：SOLID（稳定的），其代表的含义也就是把这 6 个原则结合使用的好处：建立稳定、灵活、健壮的设计

|  原则  | 描述  |  表达式  |
| ----  |----   | ----   |
|SRP单一职责原则 |一个类或者一个函数只负责一个功能或者职责，实现高内聚、低耦合的基础  |There should never be more than one reason for a class to change |
|OCP开闭原则 |开闭原则并不是完全拒绝修改，而是以最小力度的修改完成新功能的开发，具备扩展意识、抽象意识、封装意识 |Software entities (classes, modules, functions) should be open for extension but closed for modification |
|LSP里式替换原则 |子类对象可以替换程序中父类出现的任何地方，并且保证原来的逻辑行为不变及正确性不被破坏（就是子类对父类做了增强，还不改变父类的逻辑）  |Functions that use pointers of references to base classes must be able to use  objects of  derived classes without knowing it  |
|LOD迪米特法则 |如果两个软件实体无须直接通信，那么就不应当发生直接的相互调用，可以通过第三方转发该调用。其目的是降低类之间的耦合度，提高模块的相对独立性  |Each unit should have only limited knowledge about other units: only units “closely” related to the current unit. Or: Each unit should only talk to its friends;Don’t talk to strangers. |
|ISP接口隔离原则 |客户端不应该强迫依赖它不需要的接口。其中的"客户端"可以理解为调用者和使用者  |Clients should not be forced to depend upon interfaces that they don't use。/ The dependency of one class to another one should depend on the smallest possible interface |
|DIP依赖倒置原则 |高层模块不应该直接依赖于底层模块的具体实现，而应该依赖于底层的抽象。(面向接口编程)|High level modules should not depend upon low level modules. Both should depend upon abstractions. Abstractions should not depend upon details. Details should depend upon abstractions |

### 设计模式
