package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Reduction(x, y int) int {
	return x - y
}

func Multiply(x, y float64) float64 {
	return x * y
}

func Divider(x, y float64) float64 {
	return x / y
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

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

	// Another Example

	fmt.Println("2 + 2 =", Add(2, 2))
	fmt.Println("47 - 19 =", Reduction(47, 19))
	fmt.Println("4.7 * 3.2 =", Multiply(4.7, 3.2))
	fmt.Println("7 / 2 =", Divider(7, 2))

	// Array Data Type
	// Array contains a collection of data of the same type
	// when create an array, we need to determine the amount of data that the array can hold
	// the capacity of array cannot be increased after the array is created

	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Rosyid"
	names[2] = "Izzulkhaq"

	fmt.Println(names)
	fmt.Println(len(names)) // to show how much length/amount data of array

	values := [5]int{
		80,
		75,
		90,
		67,
		72,
	}

	fmt.Println(values)

	fmt.Println(Sum(values[:]))

	// Slice Data Type
	// Slice data type is a slice of array data type
	// Slice are similar with Arrays, the difference is that the size of Slice can change
	// Slice and Arrays are always connected, where Slice are data that access some or all the data in the Array

	// The Detail from Slice Data Type
	// Slice data type have three data, pointer, length, and capacity
	// Pointer is the first data pointer in the slice array
	// Length is length of the slice
	// and Capacity is the capacity of the slice, where the length cannot more than the capacity
	// If we changed data in array, the data in the slice will also change (vice versa)
	// because data in slice is reference from data in array

	// Example:

	month := [...]string{ // if we don't know how much length the data, we can use [...]
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	semester1 := month[0:6]  // pointer = 0, length = 6, capacity = 12 (from pointer/0 until end of array/12)
	semester2 := month[6:12] // pointer = 6, length = 6, capacity = 6 (from pointer/6 until end of array/12)

	fmt.Printf("Semester 1: %q, Capacity: %d, Length: %d\n", semester1, cap(semester1), len(semester1))
	fmt.Printf("Semester 2: %q, Capacity: %d, Length: %d\n", semester2, cap(semester2), len(semester2))

	// Append Function in Slice
	// append is function to create new slice from existing slice and add/replace data to the array
	// if capacity from existing slice is full, it will create new array,
	// if capacity is available, it will replace the existing data
	// and the new slice will not reference from existing array

	semester3 := append(semester1, "Zulhijah") // Juli in month will replace with Zulhijah
	fmt.Printf("Semester 3: %q, Capacity: %d, Length: %d\n", semester3, cap(semester3), len(semester3))

	semester4 := append(semester2, "Syakban") // This will create new array
	fmt.Printf("Semester 4: %q, Capacity: %d, Length: %d\n", semester4, cap(semester4), len(semester4))

	fmt.Println("Value month:", month)

	// Make function in slice
	// make is function to create slice without existing array

	day := make([]string, 6, 7)
	day[0], day[1], day[2], day[3], day[4], day[5] = "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"
	weekday := append(day[5:], "Minggu")

	fmt.Printf("Weekday: %q\n", weekday)
	fmt.Printf("Day: %q, Len: %d, Cap: %d\n", day, len(day), cap(day)) // Minggu will not show because the length is 6

	day = day[0:7]
	fmt.Printf("Full Day: %q, Len: %d, Cap: %d\n", day, len(day), cap(day))

	// Please be careful when create array or slice
	thisArray := [...]int{90, 86, 79} // use ... or specific number to create array
	thisSlice := []int{10, 20, 30}

	fmt.Println(thisArray)
	fmt.Println(thisSlice)

	// Map Data Type
	// Map is like array and slice, this data type contains a collection of the data, but we can change the index with
	// some value
	// Map is collection of key-value, the key is unique
	// The different between array is Map has no limit in capacity, we can store data as much as possible, as long
	// as the keywords are different, if the key are same, then previous data will be replaced with the new data

	identity := map[string]string{ // map[TypeKey]TypeValue{...}
		"name":    "Rosyid",
		"address": "Temanggung",
	}

	identity["title"] = "Backend Developer"

	fmt.Println("Identity:", identity)
	fmt.Println("Name:", identity["name"])
	fmt.Println("Address:", identity["address"])

	// Function in map
	// len(map) = return how much data in map
	// map[key] = return value based on the given key
	// map[key] = update value based on the given key
	// make(map[TypeKey]TypeValue) = create map
	// delete(map, key) = delete data from map based on the given key

	delete(identity, "address")
	fmt.Println("Identity:", identity)
}
