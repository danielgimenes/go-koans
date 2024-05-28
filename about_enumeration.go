package go_koans

import "fmt"

func aboutEnumeration() {
	{
		var concatenated string
		var total int

		strings := []string{"hello", " world", "!"}
		for i, v := range strings {
			total += i
			concatenated += v
		}

		assert(concatenated == "hello world!") // for loops have a modern variation
		assert(total == 3)                     // which offers both a value and an index
	}

	{
		var totalLength int

		strings := []string{"hello", " world", "!"}
		for _, v := range strings {
			totalLength += len(v)
			fmt.Println("Size of", v, "is:", len(v))
		}

		assert(totalLength == 12) // although we may omit either value
	}
}
