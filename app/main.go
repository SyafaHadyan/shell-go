package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Command    string
	CommandArg []string
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
	c.CommandArg = strings.Split(c.Command, " ")

	switch c.CommandArg[0] {
	case "exit":
		c.Exit()
	default:
		c.NotFound()
	}
}

func (c *Command) Exit() {
	code, _ := strconv.Atoi(c.CommandArg[1])

	os.Exit(code)
}

func (c *Command) NotFound() {
	fmt.Fprintf(os.Stderr, "%s: command not found\n", c.Command)
}
