# 千层面大师

欢迎来到 Exercism Go 轨道中的 千层面大师。
如果你在运行测试或提交代码时需要帮助，请查看`HELP.md`。
如果你在练习中卡住了，请查看`HINTS.md`，不过先试着自己解决，不要急着看提示 :)

## Introduction

函数让你可以把代码组织成一个可复用的单元。
它由`func`关键字、函数名，以及括号中用逗号分隔的零个或多个形参及其类型组成。

## 函数形参

所有形参都必须显式标注类型；形参没有类型推断。
形参没有默认值，因此所有函数形参都是必需的。

```go
import "fmt"

// No parameters
func PrintHello() {
    fmt.Println("Hello")
}

// Two parameters
func PrintGreetingName(greeting string, name string) {
  fmt.Println(greeting + " " + name)
}
```

类型相同的形参可以一起声明，后面跟一个类型声明。

```go
import "fmt"

func PrintGreetingName(greeting, name string) {
  fmt.Println(greeting + " " + name)
}
```

## 形参与实参

我们快速来说说两个经常被混淆的术语：`parameters`和`arguments`。
函数形参是在函数签名中定义的名称，比如上面`PrintGreetingName`函数中的`greeting`和`name`。
函数实参是我们在调用函数时传递给函数形参的具体值。
例如，在下面的例子中，`"Hello"`和`"Katrina"`就是传给`greeting`和`name`这两个形参的实参：

```go
PrintGreetingName("Hello", "Katrina")
```

## 返回值

函数形参后面是零个或多个返回值，它们同样必须显式标注类型。
单个返回值不加括号，多个返回值则用括号包起来。
使用`return`关键字可以把值从函数返回给调用方代码。
一个函数中可以有多个`return`语句。
函数一执行到其中某个`return`语句就会立即结束。
如果要从函数返回多个值，它们之间用逗号分隔。

```go
func Hello(name string) string {
  return "Hello " + name
}

func HelloAndGoodbye(name string) (string, string) {
  return "Hello " + name, "Goodbye " + name
}
```

## 调用函数

调用函数时，要指定函数名，并在括号中为函数的每个形参传入实参。

```go
import "fmt"

// No parameters, no return value
func PrintHello() {
    fmt.Println("Hello")
}
// Called like this:
PrintHello()

// One parameter, one return value
func Hello(name string) string {
  return "Hello " + name
}
// Called like this:
greeting := Hello("Dave")

// Multiple parameters, multiple return values
func SumAndMultiply(a, b int) (int, int) {
    return a+b, a*b
}
// Called like this:
aplusb, atimesb := SumAndMultiply(a, b)
```

## 具名返回值与裸返回

和形参一样，返回值也可以选择具名。
如果使用了具名返回值，不带实参的`return`语句就会返回这些值。
这被称为“裸”返回。

```go
func SumAndMultiplyThenMinus(a, b, c int) (sum, mult int) {
    sum, mult = a+b, a*b
    sum -= c
    mult -= c
    return
}
```

## 传值与传引用

同样重要的是厘清传值和传引用的概念。

首先，厘清一下传值。
在下面的例子中，当我们把变量`val`传给函数`MultiplyByTwo`时，传递的是`val`的一份副本。
因此，`newVal`得到了更新后的值`4`，但原来的变量`val`仍然是`2`。
在幕后，Go 实际上会复制原始值，所以被函数修改的只有这个副本，也就是`v`。

```go
val := 2
func MultiplyByTwo(v int) int {
    v = v * 2
    return v
}
newVal := MultiplyByTwo(val)
// newval is 4, val is still 2 because only a copy of its value was passed into the function
```

严格来说，Go 中所有实参都是按值传递的，也就是说，会为提供给函数的值或数据创建一份副本。
但如果你不想复制传给函数的数据，而是想在函数中修改这些数据，那就应该把`pointers`作为实参，也就是按引用传递。

## 指针

我们用`pointer`来实现按引用传递。
通过把`pointer`实参传给函数，我们可以修改传入函数的底层数据，而不只是操作数据的一份副本。

目前，你只需知道，在函数签名中，指针类型可以通过类型前面的`*`来识别。

```go
func HandlePointers(x, y *int) {
    // Some logic to handle integer values referenced by pointers x and y
}
```

如果`pointer`这个概念让你感到困惑，不用担心。
教学大纲后面有专门的一节，帮你掌握指针。

## 例外情况

注意，`slices`和`maps`是上述规则的例外。
当我们把`slice`或`map`作为实参传给函数时，即使类型中没有显式的`*`，它们也会被当作指针类型处理。
这意味着，如果我们把一个切片或映射传给函数并修改其底层数据，这些改动会反映到原来的切片或映射上。

## Instructions

在这个练习中，你要再写一些代码，它们与你最爱的食谱里那道绝妙千层面的准备和烹制有关。

一共有四个任务。
第一个与烹制本身有关，另外三个则关于完美的准备工作。

## 1. 估算准备时间

为了准备下一次千层面，你想确保预留出足够的时间，好让自己享受烹饪的过程。
你已经规划好了千层面要有哪些层。
现在，你想据此估算准备需要多长时间。

实现一个函数`PreparationTime`，它接受一个`[]string`类型的层切片，以及每层的平均准备时间（单位：分钟，类型为`int`）。
该函数应根据层的数量，以`int`返回总准备时间的估算值。
Go 的函数没有默认值。
如果传入的平均准备时间为`0`（`int`的默认初始值），则应使用默认值`2`。

```go
layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
PreparationTime(layers, 3)
// => 18
PreparationTime(layers, 0)
// => 12
```

## 2. 计算所需的意面和酱汁用量

除了预留时间，你还要确保有足够的酱汁和面条，来做出你梦寐以求的千层面。
千层面中每一层面条都需要 50 克面条。
每一层酱汁都需要 0.2 升酱汁。

定义函数`Quantities`，它的形参是一个`[]string`类型的层切片。
该函数会确定做这顿饭所需的面条和酱汁用量。
结果应以两个值返回：`noodles`为`int`，`sauce`为`float64`。

```go
Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"})
// => 100, 0.4
```

## 3. 加入秘密食材

不久前，你去一位朋友家做客，在那里吃了千层面。
味道棒极了，而且别有风味。
朋友把食材清单发给了你，并告诉你清单上的最后一项就是让这顿饭如此特别的“秘密食材”。
现在，你也想把这种秘密食材加到自己的菜谱里。

编写一个函数`AddSecretIngredient`，它接受两个`[]string`类型的食材切片作为形参。
第一个形参是朋友发给你的列表，第二个是你自己菜谱的食材列表。
你的食材列表中的最后一个元素总是`"?"`。
该函数应把它替换为朋友列表中的最后一项。
**注意：**`AddSecretIngredient`不返回任何内容，你应当直接修改自己的食材列表。
朋友那份食材列表**不应该**被修改。
此外，由于`slice`是以指针的形式传入函数的，对传入`AddSecretIngredient`的两个`[]string`实参所做的修改会直接生效。

```go
friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
myList := []string{"noodles", "meat", "sauce", "mozzarella","?"}

AddSecretIngredient(friendsList, myList)
// myList => []string{"noodles", "meat", "sauce", "mozzarella", "kampot pepper"}
```

## 4. 按比例缩放菜谱

你的食谱中列出的用量只够做两人份的千层面。
既然下次你想为更多人做饭，就需要计算不同份数下的用量。

实现一个函数`ScaleRecipe`，它接受两个形参。

- 一个`float64`切片，表示做 2 份所需的用量。
- 你想做的份数。

该函数应返回一个`float64`切片，其中是做出所需份数需要的用量。
不过，你希望保留原始菜谱，也就是说，`quantities`这个实参不应在这个函数中被修改。

```go
quantities := []float64{ 1.2, 3.6, 10.5 }
scaledQuantities := ScaleRecipe(quantities, 4)
// => []float64{ 2.4, 7.2, 21 }
```

## Source

### Created by

- @bobtfish

### Contributed to by

- @sougat818