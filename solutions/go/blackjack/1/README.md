# 二十一点

欢迎来到 Exercism Go 轨道中的 二十一点。
如果你在运行测试或提交代码时需要帮助，请查看`HELP.md`。
如果你在练习中卡住了，请查看`HINTS.md`，不过先试着自己解决，不要急着看提示 :)

## Introduction

和其他语言一样，Go 也提供了 `switch` 语句。相比冗长的 `if ... else if` 语句，`switch` 语句写起来更简短。要创建一个 switch 语句，先写关键字 `switch`，后面跟一个值或表达式。然后我们用 `case` 关键字声明每一个条件。我们还可以声明一个 `default` 分支，当之前所有 `case` 条件都不匹配时，它就会运行：

```go
operatingSystem := "windows"

switch operatingSystem {
case "windows":
    // do something if the operating system is windows
case "linux":
    // do something if the operating system is linux
case "macos":
    // do something if the operating system is macos
default:
    // do something if the operating system is none of the above
} 
```

switch 语句有一个有趣的地方：`switch` 关键字后面的值可以省略，这样每个 `case` 都可以是布尔条件：

```go
age := 21

switch {
case age > 20 && age < 30:
    // do something if age is between 20 and 30
case age == 10:
    // do something if age is equal to 10
default:
    // do something else for every other case
}
```

## Instructions

在这个练习中，我们将模拟一局 [Blackjack](https://en.wikipedia.org/wiki/Blackjack) 游戏的第一回合。

你会拿到两张牌，并且能看到庄家的一张明牌。
所有牌都用字符串表示，比如“ace”、“king”、“three”、“two”等。
每张牌的点数如下：

| 牌 | 点数 | 牌 | 点数 |
| :---: | :---: | :-----: | :---: |
|  ace  |  11   | eight   |   8   |
|  two  |   2   | nine    |   9   |
| three |   3   |  ten    |  10   |
| four  |   4   | jack    |  10   |
| five  |   5   | queen   |  10   |
|  six  |   6   | king    |  10   |
| seven |   7   | *other* |   0   |

**注意**：通常 ace 可以算作 1 或 11，但为了简单起见，我们假设它只能算作 11。

根据你的两张牌和庄家的那张牌，第一回合有一套策略，你有以下选择：

- 停牌（S）
- 要牌（H）
- 分牌（P）
- 自动获胜（W）

虽然还不是最优策略，但你会采用你朋友 Alex 一直在完善的策略，具体如下：

- 如果你拿到两张 ace，就必须分牌。
- 如果你拿到 Blackjack（两张牌点数之和为 21），而庄家的牌不是 ace、人头牌（Jack/Queen/King）或 ten，那么你自动获胜。
  如果庄家拿到了这些牌中的任意一张，你就必须停牌，等待另一张牌翻开。
- 如果你的牌点数之和在 [17, 20] 范围内，就应该始终停牌。
- 如果你的牌点数之和在 [12, 16] 范围内，就应该始终停牌，除非庄家有 7 或更大的牌，这时你就应该始终要牌。
- 如果你的牌点数之和为 11 或更小，就应该始终要牌。

## 1. 计算任意一张牌的点数

实现一个函数，计算一张牌的点数：

```go
value := ParseCard("ace")
fmt.Println(value)
// Output: 11
```

## 2. 实现第一回合的决策逻辑

编写一个函数，实现上面描述的决策逻辑：

```go
func FirstTurn(card1, card2, dealerCard string) string
```

下面是一些预期结果的示例：

```go
FirstTurn("ace", "ace", "jack") == "P"
FirstTurn("ace", "king", "ace") == "S"
FirstTurn("five", "queen", "ace") == "H"
```

## Source

### Created by

- @andres-zartab

### Contributed to by

- @tehsphinx
- @andrerfcsantos
- @norbs57