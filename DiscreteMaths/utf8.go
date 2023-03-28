package main

import "fmt"

func encode(utf32 []rune) []byte {
	utf8 := make([]byte, 0)
	for _, elem := range utf32 {
		if elem>>7 == 0 {
			utf8 = append(utf8, byte(elem))
		} else if elem>>11 == 0 {
			byte1 := 3<<6 + byte(elem>>6)
			byte2 := 1<<7 + byte(elem&63)
			utf8 = append(utf8, byte1, byte2)
		} else if elem>>16 == 0 {
			byte1 := 7<<5 + byte(elem>>12)
			byte2 := 1<<7 + byte(elem>>6&63)
			byte3 := 1<<7 + byte(elem&63)
			utf8 = append(utf8, byte1, byte2, byte3)
		} else {
			byte1 := 15<<4 + byte(elem>>18)
			byte2 := 1<<7 + byte(elem>>12&63)
			byte3 := 1<<7 + byte(elem>>6&63)
			byte4 := 1<<7 + byte(elem&63)
			utf8 = append(utf8, byte1, byte2, byte3, byte4)
		}
	}
	return utf8
}

func decode(utf8 []byte) []rune {
	utf32 := make([]rune, 0)
	for idx := 0; idx < len(utf8); idx++ {
		if utf8[idx]>>7 == 0 {
			utf32 = append(utf32, rune(utf8[idx]))
		} else if utf8[idx]>>5 == 6 {
			byte1 := rune((utf8[idx+1]<<1)>>1) + rune(utf8[idx]<<2)<<4
			utf32 = append(utf32, byte1)
			idx += 1
		} else if utf8[idx]>>4 == 14 {
			byte1 := rune((utf8[idx+2]<<1)>>1) + rune(utf8[idx+1]<<2)<<4 + rune(utf8[idx]<<4)<<8
			utf32 = append(utf32, byte1)
			idx += 2
		} else {
			byte1 := rune((utf8[idx+3]<<1)>>1) + rune(utf8[idx+2]<<2)<<4 + rune(utf8[idx+1]<<2)<<10 + rune(utf8[idx]<<5)<<13
			utf32 = append(utf32, byte1)
			idx += 3
		}
	}
	return utf32
}

func main() {
	example := ([]rune)("😀")
	fmt.Println(example)
	fmt.Println(decode(encode(example)))
	fmt.Println(encode(example))
}
