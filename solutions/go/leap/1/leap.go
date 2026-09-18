package leap

func IsLeapYear(year int) bool {
	rs := false
	if year%4 == 0 && year%100 != 0 {
		rs = true
	}
	if year%400 == 0{
		rs = true
	}
    
	return rs
}
