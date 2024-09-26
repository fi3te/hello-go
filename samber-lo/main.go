package main

import (
	"fmt"
	"time"

	"github.com/samber/lo"
)

func main() {
	// slices
	fmt.Printf("Filter: %v\n", lo.Filter([]string{"a", "bc", "d", "e"}, func(x string, _ int) bool {
		return len(x) == 1
	}))
	fmt.Printf("Times: %v\n", lo.Times(3, func(_ int) string {
		return "a"
	}))
	fmt.Printf("Uniq: %v\n", lo.Uniq([]string{"a", "b", "b", "c"}))

	// maps
	aMap := map[string]int{"a": 1, "b": 2}
	fmt.Printf("Keys: %v\n", lo.Keys(aMap))
	fmt.Printf("Values: %v\n", lo.Values(aMap))

	// strings
	fmt.Printf("RandomString: %v\n", lo.RandomString(20, lo.AlphanumericCharset))
	fmt.Printf("Substring: %v\n", lo.Substring("abcdef", 3, 2))
	fmt.Printf("ChunkString: %v\n", lo.ChunkString("abcdefghijklmnop", 4))

	// duration
	fmt.Printf("Duration: %v\n", lo.Duration(func() { time.Sleep(20 * time.Millisecond) }))

	// intersections
	fmt.Printf("Contains: %v\n", lo.Contains([]string{"a", "b"}, "b"))
	fmt.Printf("None: %v\n", lo.None([]string{"a", "b"}, []string{"c"}))
}
