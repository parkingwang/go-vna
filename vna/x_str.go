package vna

import (
	"strings"
	"unicode/utf8"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func starts(s string, v string) bool {
	return strings.HasPrefix(s, v)
}

func ends(s string, v string) bool {
	return strings.HasSuffix(s, v)
}

func sizeOf(s string) int {
	return utf8.RuneCountInString(s)
}
