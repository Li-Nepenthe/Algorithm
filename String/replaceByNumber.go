package String

import (
	"fmt"
	"strings"
)

func replaceByNumber(s string) string {
	st := strings.Builder{}
	b := []byte(s)
	for i := 0; i < len(b); i += 1 {
		if b[i] >= '0' && b[i] <= '9' {
			st.WriteString("number")
		} else {
			st.WriteByte(b[i])
		}
	}
	return st.String()
}

func main() {
	var s string
	fmt.Scan(&s)
	replaceByNumber(s)
	fmt.Println(s)
}
