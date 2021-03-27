package main

import (
  "fmt"
)

func main() {
    var notaA,notaB,media float32
    fmt.Scan(&notaA)
    fmt.Scan(&notaB)
    media = (notaA * 3.5 + notaB * 7.5) / 11.0
    fmt.Printf("MEDIA = %.5f\n", media);
}

