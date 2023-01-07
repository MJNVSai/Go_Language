package main
import "fmt"

func main(){
	// Integer DataTypes
	var x int = 500
	var x1 = 200
	x2 := 4500
	fmt.Printf("Type: %T, value: %v\n",x,x) // %T and %v works on only Printf Function..
	fmt.Printf("Type: %T, value: %v\n",x1,x1)
	fmt.Printf("Type: %T, value: %v\n",x2,x2)

	// Float Datatypes
	fmt.Println()
	var y float32 = 12.56 // or you can use float64 type also
	var y1 = 19.999
	y2 := 10.15
	fmt.Printf("Type: %T, value: %v\n",y,y)
	fmt.Printf("Type: %T, value: %v\n",y1,y1)
	fmt.Printf("Type: %T, value: %v\n",y2,y2)

	// String Datatypes
	fmt.Println()
	var z string = "Hello World!"
	var z1 = "Welcome To The Go lang World"
	z3 := 'A'
	z2 := "I am Faster than Python and Efficient Than Java"
	fmt.Printf("Type: %T, value: %v\n",z,z)
	fmt.Printf("Type: %T, value: %v\n",z1,z1)
	fmt.Printf("Type: %T, value: %v\n",z2,z2)
	fmt.Printf("Type: %T, value: %v\n",z3,z3)

	// Boolean DataTypes
	fmt.Println()
	var a bool = true
	var a1 = false
	a2 := true
	fmt.Printf("Type: %T, value: %v\n",a,a)
	fmt.Printf("Type: %T, value: %v\n",a1,a1)
	fmt.Printf("Type: %T, value: %v\n",a2,a2)

}