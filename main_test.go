package box

import (
	"strconv"
	"testing"
)

var N = 1
var P = "G"

func BenchmarkList(b *testing.B) {

	last := ""
	l := []*A{}
	for i := 0; i < N; i++ {
		last = P + strconv.Itoa(i)
		l = append(l, Action(func() {}).WithName(last))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		//last = P + strconv.Itoa(rand.Intn(N))

		for _, v := range l {
			if last == v.Name {
				//fmt.Println(v)
				break
			}
		}

	}
}

func BenchmarkMap(b *testing.B) {

	m := map[string]*A{}
	last := ""
	for i := 0; i < N; i++ {
		last = P + strconv.Itoa(i)
		m[last] = Action(func() {}).WithName(last)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//last = P + strconv.Itoa(rand.Intn(N))
		_ = m[last]
	}
}
