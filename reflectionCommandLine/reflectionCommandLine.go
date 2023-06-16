package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

type CommandLine interface {
	Method1()
	Method2()
	ShowCommands()
	executeMethod(i int)
}

type CommandLineTool struct {
	w io.Writer
}

func (c CommandLineTool) Method1() {
	printOut(c.w, "Method1 was called")
}

func (c CommandLineTool) Method2() {
	printOut(c.w, "Method2 was called")
}

func (c CommandLineTool) ShowCommands() {

	val := reflect.TypeOf(c)
	for i := 0; i < val.NumMethod(); i += 1 {
		methodNameWithId := val.Method(i).Name + " id: " + strconv.Itoa(i)
		printOut(c.w, methodNameWithId)
	}
}

func (c CommandLineTool) executeMethod(i int) {
	val := reflect.ValueOf(c)

	if i >= val.NumMethod() {
		printOut(c.w, "This method doesn't exists")
		return
	}

	method := val.Method(i)

	if !method.IsValid() {
		printOut(c.w, "This method doesn't exists")
		return
	}

	method.Call(nil)

}

func printOut(w io.Writer, s string) {
	fmt.Fprintln(w, s)
}

func (c CommandLineTool) instructions() {
	printOut(c.w, "Command Line Tool")
	printOut(c.w, "Enter one of the following methods:")

}

type Reader interface {
	Scan() bool
	Text() string
}

type userInput struct {
	scanner *bufio.Scanner
}

func (u *userInput) Scan() bool {
	return u.scanner.Scan()
}

func (u *userInput) Text() string {
	return u.scanner.Text()
}

func interact(c *CommandLineTool, u Reader) {

	c.instructions()
	c.ShowCommands()

	for u.Scan() {
		methodId := u.Text()
		i, _ := strconv.Atoi(methodId)
		c.executeMethod(i)
	}
}

func main() {

	b := bufio.NewScanner(os.Stdin)
	c := CommandLineTool{
		w: os.Stdout,
	}

	u := userInput{b}

	interact(&c, &u)

}
