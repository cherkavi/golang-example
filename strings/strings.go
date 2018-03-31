package main

import (
	"fmt"
	"strings"
)

func main() {
	first := "first element"
	second := "second element"
	fmt.Printf("compare [%v] with [%v] = %v \n", first, second, strings.Compare(first, second))
	fmt.Printf("does [%v] contains [%v] = %v \n", first, "abcd", strings.ContainsAny(first, "abcd"))
	fmt.Printf("does [%v] containsAny [%v] (space !!!) = %v \n", first, "a b", strings.ContainsAny(first, "a b"))
	fmt.Printf("does [%v] containsRune [%v] (space !!!) = %v \n", first, 0x20, strings.ContainsRune(first, 0x20))
	fmt.Printf("count into [%v] of [%v] = %v \n", first, "e", strings.Count(first, "e"))
	fmt.Printf("equals ignore case  [%v] and [%v] = %v \n", "abc", "aBc", strings.EqualFold("abc", "aBc"))
	fmt.Printf("equals ignore case  [%v] and [%v] = %v \n", "abc", "abcd", strings.EqualFold("abc", "abcd"))
	fmt.Printf("split string  [%v] to [%#v] \n", "this is a message ", strings.Fields("this is a message"))
	fmt.Printf("split string  [%v] to [%#v] \n", "this is a message ", strings.FieldsFunc("this is a message", func(symbol rune) bool {
		return symbol == ' '
	}))
	fmt.Printf("has  [%v] prefix [%v]  ? [%v] \n", first, "first", strings.HasPrefix(first, "first"))
	fmt.Printf("has  [%v] suffix [%v]  ? [%v] \n", first, "element", strings.HasSuffix(first, "element"))
	fmt.Printf("before [%v] after Map [%v] \n", first, strings.Map(func(symbol rune) rune {
		if symbol == 'e' {
			return 'E'
		}
		return symbol
	}, first))
	fmt.Printf("replace of the string : %v \n", strings.Replace("this is my $personal string", "$personal", "RED", -1))

	fmt.Printf("repeat symbol [%v]\n", strings.Repeat("*_*|", 5))

	var stringBuilder strings.Builder
	fmt.Fprintf(&stringBuilder, " %v ", first)
	fmt.Fprintf(&stringBuilder, " %v ", second)
	fmt.Fprintf(&stringBuilder, " \n ")
	fmt.Print(stringBuilder.String())

	var reader = strings.NewReader(first)
	fmt.Printf(" len: [%v] of [%v] \n", reader.Len(), first)

	if nextByte, notOk := reader.ReadByte(); notOk == nil {
		fmt.Printf("first byte of [%v]  is %q \n", first, nextByte)
	}

	s := "abcdefgh"
	fmt.Printf(" original string [%v] substring 0:3 [%v]  substring 3:len(s) [%v]\n ", s, s[:3], s[3:])

}
