package main
import (
	"fmt"
	"strconv"
	"bytes"
)
type Persion struct {
	Name string
	Age int
	Sex int
}

func (p *Persion) String() string {
	buffer := bytes.NewBufferString("this is ")
	buffer.WriteString(p.Name + " ")
	if p.Sex == 0 {
		buffer.WriteString("He ")
	} else {
		buffer.WriteString("She ")
	}

	buffer.WriteString("is ")
	buffer.WriteString(strconv.Itoa(p.Age))
	buffer.WriteString(" years old.")
	
	return buffer.String()
}

func (p *Persion) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(p.String()))
		f.Write([]byte(" Persion has three fields."))
	} else {
		// 没有下面的语句什么也不会输出
		f.Write([]byte(fmt.Sprintln(p.String())))
	}
}

func (p  *Persion) GoString() string {
	return "&Persion{Name is "+p.Name+", Age is "+strconv.Itoa(p.Age)+", Sex is "+strconv.Itoa(p.Sex)+"}"
}

func main() {
	p := &Persion{
		Name: "anziguoer",
		Age: 28,
		Sex: 1,
	}
	fmt.Println(p)
	fmt.Printf("%L", p)
}