package main

import (
  "fmt"
  "math"
)

func main() {
    var x1,x2,y1,y2,x,y,distancia float64
    fmt.Scan(&x1,&y1)
    fmt.Scan(&x2, &y2)
    x = math.Pow((x2-x1),2)
    y = math.Pow((y2-y1),2)
    distancia =  math.Sqrt(x + y)
    fmt.Printf("%.4f\n", distancia);
}

