package cowsay

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

type CLI struct {
}

type options struct {
	Help bool `short:"h"`
}

func (c *CLI) Run(argv []string) int {
	if err := c.mow(argv); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return 1
	}
	return 0
}

func (c *CLI) mow(argv []string) error {
	var opts options
	args, err := c.parseOptions(&opts, argv)

	if err != nil {
		return err
	}

	if err := c.mowmow(&opts, args); err != nil {
		return err
	}
	return nil
}

func (c *CLI) parseOptions(opts* options, argv []string) ([]string, error) {
	p := flags.NewParser(opts, Flags.None)
	
}