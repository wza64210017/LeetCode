package main

func reverseString(s []byte) {
	length := len(s)-1
	if length <= 0 {
		return
	}

	for i := 0; i <= length/2; i++ {
		s[i], s[length-i] = s[length-i], s[i]
	}
}

func main() {
	s := []byte("hello")
	reverseString(s)
}
