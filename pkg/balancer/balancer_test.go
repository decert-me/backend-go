package balancer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSW_Next(t *testing.T) {
	w := NewSmoothRoundrobin()

	s := w.Next()
	//assert.Nil(t, s)
	//done(balancer.DoneInfo{})

	w.Add("server1", 5)
	s = w.Next()
	assert.Equal(t, "server1", s.Item)

	w.Reset()
	s = w.Next()
	assert.Nil(t, s)

	w.Add("server1", 5)
	s = w.Next()
	assert.Equal(t, "server1", s.Item)

	w.Add("server2", 2)
	w.Add("server3", 3)

	results := make(map[string]int)

	for i := 0; i < 1000; i++ {
		s := w.Next()
		results[s.Item]++
	}

	if results["server1"] != 500 || results["server2"] != 200 || results["server3"] != 300 {
		t.Error("the algorithm is wrong")
	}

	w.items[0].EffectiveWeight = w.items[0].CurrentWeight - 1
	s = w.Next()
	assert.Equal(t, "server3", s.Item)
}

func Test_nextSmoothWeighted(t *testing.T) {
	w := NewSmoothRoundrobin()

	w.Add("server1", 0)
	w.Add("server2", 0)
	w.Add("server3", 0)
	fmt.Println(w.Next())
	fmt.Println(w.Next())
	w.Next().OnInvokeFault()
	//items[0].EffectiveWeight--
	for i := 0; i < 20; i++ {
		fmt.Println(w.Next())
	}
}
