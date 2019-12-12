package main

import "fmt"

//var n = flag.Bool("n", false, "omit trailing newline")
//var sep  = flag.String("s", " ", "separator")
func main() {
	//flag.Parse()
	//fmt.Print(strings.Join(flag.Args(), *sep) + "\n")
	//if !*n {
	//	fmt.Println(*n)
	//}
	medals := []string{"gold", "silver", "bronze"}
	fmt.Printf("%T", medals)
	p := new(int)
	s := new(string)

	fmt.Println(p, *p, s, *s)
	*p = 2
	*s = "Hello world"

	fmt.Println(*p, *s)

	fmt.Println(fbi(10));
}

func fbi(n int) int  {
	x, y := 0, 1
	for i := 0; i<n; i++  {
		x, y = y, x+y
		fmt.Println(x, y)
	}

	return y;
}
