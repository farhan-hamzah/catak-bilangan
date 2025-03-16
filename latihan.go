

package main
import "fmt"

func cetakBilangan(x, y, a, b int) {
    var ratusan, satuan int
    for x <= y {
        ratusan = x/100
        satuan = x%10
        if ratusan == a || satuan == b{
            x++
        }else{
            fmt.Println(x)
            x++
        }
    }
}
func main(){
    var x, y, a, b int
    fmt.Scan(&x, &y, &a, &b)
    cetakBilangan(x, y, a, b)
}