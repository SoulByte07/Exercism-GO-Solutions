package luhn

func Valid(id string) bool {
	// panic("Please implement the Valid function")
	digits :=make([]byte,0,len(id))
    for i:=0;i<len(id);i++ {
        ch:=id[i]
        if ch==' ' {
            continue
        }
        if ch<'0' || ch>'9' {
            return false
        }
        digits=append(digits, ch-'0')
    }
    if len(digits)<=1 {
        return false
    }
    sum:=0
    double:=false
    for i:=len(digits)-1;i>=0;i-- {
		digit:=int(digits[i])
        if double {
            digit*=2
            if digit>9 {
                digit-=9
            }
        }
		sum+=digit
        double=!double
    }
    return sum%10==0
}
