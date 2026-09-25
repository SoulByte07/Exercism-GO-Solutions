package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
	// panic("Please implement the CollatzConjecture function")
    res:=0
    if n<=0 {
        return 0, errors.New("fuck you")
    }
    for (n!=1) {
    	if n%2==0 {
            n=n/2
        } else {
            n=n*3 + 1
        }
    	res++
	}
	return res, nil
}