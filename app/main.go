package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Command struct {
	Command string
}

func main() {
	start()
}

func start() {
	var command Command
	command.CommandInput()
}

func (c *Command) CommandInput() {
	for {
		_, _ = fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Println(err)
		}

		c.Command = strings.Replace(input, "\n", "", 1)

		c.HandleCommand()
	}
}

func (c *Command) HandleCommand() {
	c.NotFound()
}

func (c *Command) NotFound() {
	fmt.Fprintf(os.Stderr, "%s: command not found\n", c.Command)
}
