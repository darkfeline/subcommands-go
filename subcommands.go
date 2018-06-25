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

/*
Package subcommands implements dead simple subcommands for Go.
*/
package subcommands

import "fmt"

// Cmd defines the interface for a subcommand to be used in Run.
type Cmd interface {
	Name() string
	Run([]string)
}

// New builds an object that implements Cmd.
func New(n string, f func([]string)) Cmd {
	return &cmd{n, f}
}

type cmd struct {
	name string
	f    func([]string)
}

func (c *cmd) Name() string {
	return c.name
}

func (c *cmd) Run(args []string) {
	c.f(args)
}

// Run dispatches to the matching command using the arguments.  The
// arguments slice should have at least one item.  The first string
// should be the name of the subcommand.  An error is returned if no
// subcommand is matched.
func Run(c []Cmd, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no subcommand argument")
	}
	n := args[0]
	for _, c := range c {
		if c.Name() == n {
			c.Run(args)
			return nil
		}
	}
	return fmt.Errorf("unmatched command %s", n)
}
