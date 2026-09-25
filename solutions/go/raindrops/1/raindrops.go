package raindrops
import "strconv"
func Convert(number int) string {
    var res string
    if number%3==0 {
        res=res+"Pling"
    } 
    if number%5==0 {
        res=res+"Plang"
    }
    if number%7==0 {
        res=res+"Plong"
    }
    if !(number%3==0 || number%5==0 || number%7==0) {
        return strconv.Itoa(number)
    } 
    return res
    // panic("Please implement the Convert function")
}
