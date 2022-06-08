package main

import "fmt"

func decode(utf8 []byte) []rune {
	var utf32 []rune
	var sym rune
	for i := 0; i < len(utf8); i++ {
		if utf8[i]&0x00000080 == 0 {
			utf32 = append(utf32, rune(utf8[i]))
		} else if utf8[i]&0x000000E0 == 0x000000C0 {
			sym = ((rune(utf8[i]) & 0x0000001F) << 6) | (rune(utf8[i+1]) & 0x0000003F)
			utf32 = append(utf32, sym)
			i++
		} else if utf8[i]&0x000000F0 == 0x000000E0 {
			sym = ((rune(utf8[i]) & 0x0000000F) << 6) | (rune(utf8[i+1]) & 0x0000003F)
			i++
			sym = (sym << 6) | (rune(utf8[i+1]) & 0x0000003F)
			utf32 = append(utf32, sym)
			i++
		} else {
			sym = ((rune(utf8[i]) & 0x00000007) << 6) | ((rune(utf8[i+1])) & 0x0000003F)
			i++
			sym = (sym << 6) | ((rune(utf8[i+1])) & 0x0000003F)
			i++
			sym = (sym << 6) | ((rune(utf8[i+1])) & 0x0000003F)
			utf32 = append(utf32, sym)
			i++
		}
	}
	return utf32
}

func encode(utf32 []rune) []byte {
	var utf8 []byte
	for _, x := range utf32 {
		if x < 0x0000007F {
			utf8 = append(utf8, byte(x))
		} else if x < 0x000007FF {
			b1 := byte(((x & 0x000007C0) >> 6) | 0x000000C0)
			b2 := byte((x & 0x0000003F) | 0x00000080)
			utf8 = append(utf8, b1, b2)
		} else if x < 0x0000FFFF {
			b1 := byte(((x & 0x0000F000) >> 12) | 0x000000E0)
			b2 := byte(((x & 0x00000FC0) >> 6) | 0x00000080)
			b3 := byte((x & 0x0000003F) | 0x00000080)
			utf8 = append(utf8, b1, b2, b3)
		} else {
			b1 := byte(((x & 0x00180000) >> 18) | 0x03C00000)
			b2 := byte(((x & 0x0003F000) >> 12) | 0x00080000)
			b3 := byte(((x & 0x00000FC0) >> 6) | 0x00002000)
			b4 := byte((x & 0x0000003F) | 0x00000080)
			utf8 = append(utf8, b1, b2, b3, b4)
		}
	}
	return utf8
}

func main() {
	example := []rune("")
	result1 := encode(example)
	result2 := decode(result1)
	fmt.Print(result1, result2)
}
