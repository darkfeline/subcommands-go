package subcommands

import (
	"reflect"
	"testing"
)

func TestRunCallsWithSameArgs(t *testing.T) {
	t.Parallel()
	var args []string
	f := func(a []string) {
		args = a
	}
	c := []Cmd{
		New("foo", f),
	}
	a := []string{"foo", "bar"}
	Run(c, a)
	if !reflect.DeepEqual(a, args) {
		t.Fatalf("Expected args %#v, got %#v", a, args)
	}
}

func TestRunCallsRightCommandOnlyOnce(t *testing.T) {
	t.Parallel()
	var a, b, c int
	fa := func([]string) {
		a++
	}
	fb := func([]string) {
		b++
	}
	fc := func([]string) {
		c++
	}
	cmd := []Cmd{
		New("homura", fa),
		New("hikari", fb),
		New("pneuma", fc),
	}
	Run(cmd, []string{"hikari"})
	if a != 0 {
		t.Errorf("a was called")
	}
	if b != 1 {
		t.Errorf("b was not called only once")
	}
	if c != 0 {
		t.Errorf("c was called")
	}
}
