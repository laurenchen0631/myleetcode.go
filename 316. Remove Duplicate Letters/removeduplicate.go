/*
Package removeduplicate Given a string which contains only lowercase letters,
remove duplicate letters so that every letter appear once and only once.
*/
package removeduplicate

// RemoveDuplicateLetters ...
// You must make sure your result is the smallest in lexicographical order among all possible results.
func RemoveDuplicateLetters(s string) string {
	var charCounter [26]int
	for index := range s {
		charCounter[getCharASCII(s[index])]++
	}

	stack := []byte{}
	visited := [26]bool{}
	for index := range s {
		ch := s[index]
		charCode := getCharASCII(ch)

		// decrement number of characters remaining in the string to be analysed
		charCounter[charCode]--

		//if character is already present in stack, dont bother
		if visited[charCode] {
			continue
		}

		// if current character is smaller than last character in stack which occurs later in the string again
		// it can be removed and  added later e.g stack = bc remaining string abc then a can pop b and then c
		for len(stack) > 0 &&
			ch < stack[len(stack)-1] &&
			charCounter[getCharASCII(stack[len(stack)-1])] != 0 {
			topChar := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			visited[getCharASCII(topChar)] = false
		}

		stack = append(stack, ch)
		visited[charCode] = true
	}

	return string(stack)
}

func getCharASCII(c byte) byte {
	return c - 'a'
}
