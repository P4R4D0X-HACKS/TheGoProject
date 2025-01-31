package main

import "fmt"

type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = ByteSize(1 << (10 * iota))
    MB
    GB
    TB
    PB
    EB
)

func main() {
    fmt.Println(KB, MB, GB, TB, PB, EB)
}
