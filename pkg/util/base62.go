package util

import (
	"math"
	"strings"
)

const Base10To62Dict = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Base62Len = 62

var Base62To10Dict = map[string]uint{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"a": 10, "b": 11, "c": 12, "d": 13, "e": 14, "f": 15, "g": 16, "h": 17, "i": 18, "j": 19,
	"k": 20, "l": 21, "m": 22, "n": 23, "o": 24, "p": 25, "q": 26, "r": 27, "s": 28, "t": 29,
	"u": 30, "v": 31, "w": 32, "x": 33, "y": 34, "z": 35, "A": 36, "B": 37, "C": 38, "D": 39,
	"E": 40, "F": 41, "G": 42, "H": 43, "I": 44, "J": 45, "K": 46, "L": 47, "M": 48, "N": 49,
	"O": 50, "P": 51, "Q": 52, "R": 53, "S": 54, "T": 55, "U": 56, "V": 57, "W": 58, "X": 59,
	"Y": 60, "Z": 61,
}

type Base62 struct {
}

func NewBase62() *Base62 {
	return new(Base62)
}

///Encode  编码 整数(10进制) 为 base62 字符串
func (b Base62) Encode(number uint) string {
	if number == 0 {
		return "0"
	}
	result := make([]byte, 0)
	for number > 0 {
		round := number / Base62Len
		remain := number % Base62Len
		result = append(result, Base10To62Dict[remain])
		number = round
	}
	//对结果进行反转  620 举例 求模为0 round 为10 remain为0  result  0,a 实际因为a0 类似十进制的 个位与10位 进位
	return Reverse(string(result))
}

//Decode 解码字符串为整数
func (b Base62) Decode(str string) uint {
	//去空格
	str = strings.TrimSpace(str)
	//反转
	str = Reverse(str)
	result := uint(0)
	for index, char := range []byte(str) {
		result += Base62To10Dict[string(char)] * uint(math.Pow(Base62Len, float64(index)))
	}
	return result
}

func Reverse(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}
