package clock

import "fmt"

type Clock int

const minutesPerDay = 24 * 60 // 1440

// New 根据小时和分钟创建一个 Clock。
// 小时和分钟可以是任意整数，会自动归一化到一天之内。
func New(h, m int) Clock {
	total := h*60 + m
	return Clock(normalize(total))
}

// Add 在当前时间上增加 m 分钟，返回新的 Clock。
func (c Clock) Add(m int) Clock {
	return Clock(normalize(int(c) + m))
}

// Subtract 在当前时间上减少 m 分钟，返回新的 Clock。
func (c Clock) Subtract(m int) Clock {
	return Clock(normalize(int(c) - m))
}

// String 返回 "HH:MM" 格式的字符串。
func (c Clock) String() string {
	h := int(c) / 60
	m := int(c) % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}

// normalize 把总分钟数调整到 0 到 1439 之间。
func normalize(total int) int {
	total %= minutesPerDay
	if total < 0 {
		total += minutesPerDay
	}
	return total
}

