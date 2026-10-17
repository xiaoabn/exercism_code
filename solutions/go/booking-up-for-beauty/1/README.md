# 预约美容

欢迎来到 Exercism Go 轨道中的 预约美容。
如果你在运行测试或提交代码时需要帮助，请查看`HELP.md`。
如果你在练习中卡住了，请查看`HINTS.md`，不过先试着自己解决，不要急着看提示 :)

## Introduction

在 Go 中，[`Time`][time]是一种描述时间点的类型。
日期和时间信息可以通过它的方法进行访问、比较和操作，不过也有一些函数是直接在`time`包上调用的。
当前的日期和时间可以通过 [`time.Now`][now]函数获取。

[`time.Parse`][parse]函数把字符串解析成 `Time` 类型的值。
Go 有一种特殊的方式，用来定义你期望的解析布局。
你需要用这个特殊时间戳里的值来写一个布局示例：
`Mon Jan 2 15:04:05 -0700 MST 2006`。

例如：

```go
import "time"

func parseTime() time.Time {
    date := "Tue, 09/22/1995, 13:00"
    layout := "Mon, 01/02/2006, 15:04"

    t, err := time.Parse(layout,date) // time.Time, error
}

// => 1995-09-22 13:00:00 +0000 UTC
```

[`Time.Format()`][format]方法返回时间的字符串表示。
和 `Parse` 函数一样，目标布局同样要通过一个示例来定义，这个示例使用特殊时间戳中的值。

例如：

```go
import (
    "fmt"
    "time"
)

func main() {
    t := time.Date(1995,time.September,22,13,0,0,0,time.UTC)
    formattedTime := t.Format("Mon, 01/02/2006, 15:04") // string
    fmt.Println(formattedTime)
}

// => Fri, 09/22/1995, 13:00
```

## 布局选项

要自定义布局，可以组合使用这些选项。
Go 中还提供了一些预定义的日期和时间戳[格式常量][const]。

| 时间        | 选项                                        |
| ----------- | ---------------------------------------------- |
| 年        | 2006 ; 06                                      |
| 月        | Jan ; January ; 01 ; 1                         |
| 日        | 02 ; 2 ; \_2（表示前导 0）                 |
| 星期     | Mon ; Monday                                   |
| 小时        | 15（24 小时制） ; 3 ; 03（上午或下午） |
| 分钟      | 04 ; 4                                         |
| 秒        | 05 ; 5                                         |
| 上午/下午标记  | PM                                             |
| 一年中的第几天 | 002 ; \_\_2                                    |

`time.Time`类型有各种方法可以访问特定的时间，例如 Hour：[`Time.Hour()`][hour]，Month：[`Time.Month()`][month]。
关于其工作原理的更多内容，可以在[官方文档][time]中找到。

[`time`][time]还包含另一个类型 [`Duration`][duration]，它表示经过的时间，并支持位置/时区、定时器以及其他相关功能，这些内容会在另一个概念中介绍。

[time]: https://golang.org/pkg/time/#Time
[now]: https://golang.org/pkg/time/#Now
[const]: https://pkg.go.dev/time#pkg-constants
[format]: https://pkg.go.dev/time#Time.Format
[hour]: https://pkg.go.dev/time#Time.Hour
[month]: https://pkg.go.dev/time/#Time.Month
[duration]: https://pkg.go.dev/time#Duration
[parse]: https://golang.org/pkg/time/#Parse

## Instructions

在这个练习中，你将为一个美容院的预约排程程序编写代码，这家美容院于 2012 年 9 月 15 日开业。

这个练习有五个任务，它们都与预约日期有关。

## 1. 解析预约日期

实现 `Schedule` 函数，将预约日期的文本表示解析为对应的 `time.Time` 格式：

```go
Schedule("7/25/2019 13:45:00")
// => 2019-07-25 13:45:00 +0000 UTC
```

## 2. 检查预约是否已经过去

实现 `HasPassed` 函数，它接收一个预约日期，并检查该预约是否在过去：

```go
HasPassed("July 25, 2019 13:45:00")
// => true
```

## 3. 检查预约是否在下午

实现 `IsAfternoonAppointment` 函数，它接收一个预约日期，并检查该预约是否在下午（>= 12:00 且 < 18:00）：

```go
IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
// => true
```

## 4. 描述预约的时间和日期

实现 `Description` 函数，它接收一个预约日期，并返回该日期和时间的描述：

```go
Description("7/25/2019 13:45:00")
// => "You have an appointment on Thursday, July 25, 2019, at 13:45."
```

## 5. 返回美容院开业的周年纪念日

实现 `AnniversaryDate` 函数，它以 UTC 返回美容院开业周年纪念日在当前年份的日期。

假设当前年份是 2020 年：

```go
AnniversaryDate()

// => 2020-09-15 00:00:00 +0000 UTC
```

**注意：**返回值是 `time.Time`，一天中的具体时刻并不重要。

## Source

### Created by

- @jamessouth