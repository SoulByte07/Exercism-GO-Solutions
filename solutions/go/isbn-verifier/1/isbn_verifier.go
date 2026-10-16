package isbnverifier

func IsValidISBN(isbn string) bool {
	// panic("Please implement the IsValidISBN function")
    digits := make([]int, 0, 10)
    for _ , c:= range isbn {
        if c == '-' {
            continue
        } else if c =='X' && len(digits)==9 {
    		digits=append(digits, 10)
        	continue
        }
        digits=append(digits, int(c-'0'))
    } 
    if len(digits)!=10 {
        return false
    }
    sum :=0
    for i , d := range digits {
        sum += d * (10-i)
    }
    return sum%11==0
    
}
