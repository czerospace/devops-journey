package main
import "fmt"
func main() {

     slice := []int{0,1,2,3}
     m := make(map[int]*int)

     for key,val := range slice {
        // m[key] = &val 错误写法
	value := val
	m[key] = &value
     }

    for k,v := range m {
        fmt.Println(k,"->",*v)
    }
}
