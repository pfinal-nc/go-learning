package main

import "fmt"

type Shaper interface {
    Area() float32
}

type Square struct {
    side float32
}

type Rectangle struct {
    length, width float32
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
    // sq1 := new(Square)
    // sq1.side = 5
	r := &Rectangle{5,3}
	q := &Square{5}
	shapes := []Shaper{r, q}
	fmt.Println(shapes)
	for n, _ := range shapes {
        fmt.Println("Shape details: ", shapes[n])
        fmt.Println("Area of this shape is: ", shapes[n].Area())
    }
    // var areaIntf Shaper
    // areaIntf = sq1
    // // shorter,without separate declaration:
    // // areaIntf := Shaper(sq1)
    // // or even:
    // // areaIntf := sq1
    // fmt.Printf("The square has area: %f\n", areaIntf.Area())
}