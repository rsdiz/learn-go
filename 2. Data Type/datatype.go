package main

import "fmt"

func main() {
	// 2 Type Number Data Type: Integer and Floating Point

	// Integer:
	// int8		= -128					until 128
	// int16	= -32768				until 32768
	// int32	= -2147483648			until 2147483648
	// int64	= -9223372036854775808	until 9223372036854775808

	// Unsigned Integer:
	// uint8	= 0 until 255
	// uint16	= 0	until 65535
	// uint32	= 0	until 4294967295
	// uint64	= 0	until 18446744073709551615

	// Floating Point:
	// float16		= 1.18 x 10 ^ -38	until 3.4 x 10 ^ 38
	// float32		= 2.23 x 10 ^ -308	until 1.80 x 10 ^ 308
	// complex64	= complex number with float32 real and imaginary parts
	// complex128	= complex number with float64 real and imaginary parts

	// Alias:
	// byte	: uint8
	// rune	: int32
	// int	: Minimal int32
	// uint	: Minimal uint32

	fmt.Println("Angka:", -128, "adalah tipe int8")
	fmt.Println("Angka:", -1.523423, "adalah tipe float16")

	// Boolean Data Type
	// bool: true and false

	fmt.Println("Benar:", true)
	fmt.Println("Salah:", false)

	// String Data Type

	fmt.Println("\"Ini merupakan String\", panjangnya adalah:", len("Ini merupakan String"), "karakter")
}
