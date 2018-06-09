// Copyright (C) 2018 Allen Li
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
