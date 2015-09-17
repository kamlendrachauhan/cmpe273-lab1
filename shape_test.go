package main

import "testing"

func TestAreaAndPerimetere(t *testing.T) {
    
    r := Rectangle{0,2,0,2}
    c := Circle{1,1,1}
    
    cArea := area(&c)
    cPerimeter := perimeter(&c)
    rArea := area(&r)
    rPerimeter := perimeter(&r)
    
    if cArea != 3.141592653589793 {
        t.Error("Expected value for circle area is 3.141592653589793, but got ",cArea)
    }
    if rArea != 4 {
        t.Error("Expected value for Rectangle area 4, but got ",rArea)
    }
    if cPerimeter != 6.283185307179586 {
        t.Error("Expected perimeter value for circle is 6.283185307179586, but got ",cPerimeter)
    }
    if rPerimeter != 8 {
        t.Error("Expected area value for circle area is 8, but got ",rPerimeter)
    }
}