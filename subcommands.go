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
	n := args[0]
	for _, c := range c {
		if c.Name() == n {
			c.Run(args)
			return nil
		}
	}
	return fmt.Errorf("Unmatched command %s", n)
}
