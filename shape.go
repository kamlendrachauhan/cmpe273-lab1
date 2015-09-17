package main

import ("fmt"; "math")

type Circle struct{
    x, y, r float64
}

type Rectangle struct{
    x1, x2, y1, y2 float64
}

type Shape interface{
    area() float64
    perimeter() float64
}

func distance(x1,y1,x2,y2 float64) float64{
    a := x2-x1
    b := y2-y1
    return math.Sqrt(a*a + b*b)
}
func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func (r *Rectangle) perimeter() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    
    return 2*l+2*w
}
func (c *Circle) area() float64 {
    return math.Pi *c.r*c.r
}

func (c *Circle) perimeter() float64{
    return 2* math.Pi*c.r
}

func performOperation(shape Shape){
    fmt.Println("Area is : ", shape.area())
    fmt.Println("Perimeter is : :",shape.perimeter())
}
func main() {
    r := Rectangle{1,2,4,1}
    c := Circle{1,1,10}
    performOperation(&r)
    performOperation(&c)
    
}
