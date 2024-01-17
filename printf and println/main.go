package main

import "fmt"

func main() {
	a, b, c, d, e, f := 1, 2, 3, 4, 5, 18
	fmt.Printf("%d\t %b\t %x\t %X\t %#x\n", b, b, b, b, b)
	fmt.Printf("%d\t %b\t %x\t %X\t %#x\n", a, a, a, a, a)
	fmt.Printf("%d\t %b\t %x\t %X\t %#x\n", c, c, c, c, c)
	fmt.Printf("%d\t %b\t %x\t %X\t %#x\n", d, d, d, d, d)
	fmt.Printf("%d\t %b\t %x\t %X\t %#x\n", e, e, e, e, e)
	fmt.Printf("%d\t %b\t %x\t %X\t %#x\n", f, f, f, f, f)
}
