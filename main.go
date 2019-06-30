// Creates a dummy login terminal with a timeout
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// password for success
const password = "WeezyWinstonWaldo"

// timer to fail
const timer = 10

func main() {
	makeDing(1)
	for {
		start := readInput("Are you ready?")
		if start == "yes" {
			fmt.Printf("You have %v seconds to gain access...\nStarting in ", timer)
			c := 3
			for i := 0; i < 3; i++ {
				fmt.Print(c, "...")
				c--
				time.Sleep(1 * time.Second)
			}
			clearTerm()
		} else {
			clearTerm()
			continue
		}
		checkLoop()
	}
}

// checkLoop is the primary loop
func checkLoop() {

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(timer * time.Second)
		timeout <- true
	}()

	for {
		select {
		case <-timeout:
			clearTerm()
			go makeDing(25)
			for i := 0; i < 25; i++ {
				fmt.Println("You are not worthy :(")
				time.Sleep(150 * time.Millisecond)
			}
			readInput("Press enter to continue...")
			clearTerm()
			return
		default:
			if win := readAttempt(); win {
				return
			}
		}
	}
}

// readAttempt checks the password
func readAttempt() bool {
	pass := readInput("Enter Password")
	if pass != password {
		fmt.Println("Fail! ")
		makeDing(1)
		return false
	}
	go makeDing(50)
	for i := 0; i < 50; i++ {
		fmt.Println("WINNER!")
		time.Sleep(150 * time.Millisecond)
	}
	readInput("Press enter to continue...")
	clearTerm()
	return true

}

// readInput creates an input prompt
func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + ": ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(input)
}

// clearTerm clears the screen
func clearTerm() {
	var c *exec.Cmd
	clear := true
	switch runtime.GOOS {
	case "linux":
		c = exec.Command("clear")
	case "windows":
		c = exec.Command("cmd", "/c", "cls")
	default:
		clear = false
	}
	if clear {
		err := osExec(c)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}

// osExec handles exec calls
func osExec(c *exec.Cmd) error {
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	err := c.Run()
	return err
}

// makeDing makes noise
func makeDing(count int) {
	for i := 0; i < count; i++ {
		_, _ = os.Stdout.Write([]byte("\u0007"))
		time.Sleep(150 * time.Millisecond)
	}
}
