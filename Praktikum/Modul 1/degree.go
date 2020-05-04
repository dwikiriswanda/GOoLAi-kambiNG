package main
import"fmt"
func main(){
var c, r, f, k float64

fmt.Scanln(&c)

r = c * 4 / 5
fmt.Printf("Dalam Reamur " + "%.15f\n",r)
f = c * 9 / 5 + 32
fmt.Printf("Dalam Fahrenheit " + "%.14f\n",f)
k = c + 273.15
fmt.Printf("Dalam Kelvin " + "%.14f\n",k)

fmt.Scanln(&f)
k = (f + 459.67) *5 / 9 
fmt.Printf("Dalam Kelvin " + "%.14f\n",k)

fmt.Scanln(&f)
k = (f + 459 + 2 / 3) *5 / 9 
fmt.Printf("Dalam Kelvin " + "%.14f\n",k)




}