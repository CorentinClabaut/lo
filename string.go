package lo

import (
	"unicode/utf8"
)

// Substring return part of a string.
func Substring[T ~string](str T, offset int, length uint) T {
	size := len(str)

	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}

	if offset > size {
		return Empty[T]()
	}

	if length > uint(size)-uint(offset) {
		length = uint(size - offset)
	}

	return str[offset : offset+int(length)]
}

// ChunkString returns an array of strings split into groups the length of size. If array can't be split evenly,
// the final chunk will be the remaining elements.
func ChunkString(str string, size int) []string {
	if size <= 0 {
		panic("lo.ChunkString: Size parameter must be greater than 0")
	}

	if len(str) == 0 {
		return []string{""}
	}

	if size >= len(str) {
		return []string{str}
	}

	var chunks []string = make([]string, 0, ((len(str)-1)/size)+1)
	currentLen := 0
	currentStart := 0
	for i := range str {
		if currentLen == size {
			chunks = append(chunks, str[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, str[currentStart:])
	return chunks
}

// RuneLength is an alias to utf8.RuneCountInString which returns the number of runes in string.
func RuneLength(str string) int {
	return utf8.RuneCountInString(str)
}
