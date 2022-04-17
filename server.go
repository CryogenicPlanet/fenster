package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"sync"
)

type WindowQuery struct {
	Fullscreen bool `json:"is-native-fullscreen"`
}

type SpaceQuery struct {
	ID    int `json:"id"`
	Index int `json:"index"`
}

type CommandInput struct {
	ID      string
	Command string
}

type CommandOutput struct {
	ID     string
	Out    string
	Stderr string
	Err    error
}

const ShellToUse = "bash"

var commandChannel = make(chan CommandInput, 10)
var outputChannel = make(chan CommandOutput, 10)
var shellWg sync.WaitGroup

func runShellSync(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func runShellParallel() {

	for command := range commandChannel {
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		cmd := exec.Command(ShellToUse, "-c", command.Command)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		output := CommandOutput{ID: command.ID, Out: stdout.String(), Stderr: stderr.String(), Err: err}
		outputChannel <- output
		shellWg.Done()
	}
}

func handleSpaceChange(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handle Space Change")

	go runShellParallel()

	quey := req.URL.Query()

	oldSpaceStr := quey.Get("oldSpace")

	oldSpace, err := strconv.Atoi(oldSpaceStr)

	if err != nil {
		log.Printf("error: %v\n", err)
	}

	fmt.Println("GET params were:", req.URL.Query())

	commandChannel <- CommandInput{Command: "yabai -m query --windows --space", ID: "windowQuery"}
	commandChannel <- CommandInput{Command: "yabai -m query --spaces", ID: "spaceQuery"}

	shellWg.Add(2)

	windowQuery := []WindowQuery{}
	spaceQuery := []SpaceQuery{}

	count := 0
	for output := range outputChannel {

		if output.Err != nil {

			log.Printf("error: %v\n", output.Err)
			fmt.Println("Stderr")
			fmt.Println(output.Stderr)

		} else if output.ID == "windowQuery" {
			err = json.Unmarshal([]byte(output.Out), &windowQuery)
			if err != nil {

				log.Printf("error: %v\n", err)
			}
		} else if output.ID == "spaceQuery" {
			err = json.Unmarshal([]byte(output.Out), &spaceQuery)
			if err != nil {

				log.Printf("error: %v\n", err)
			}
		}
		count++

		if count == 2 {
			break
		}
	}

	shellWg.Wait()

	fmt.Println("Query results", windowQuery, spaceQuery)

	spaceIndex := oldSpace

	for _, elm := range spaceQuery {
		if elm.ID == oldSpace {
			spaceIndex = elm.Index
			break
		}
	}

	if len(windowQuery) == 1 {
		elm := windowQuery[0]

		if elm.Fullscreen {
			fmt.Println("Fullscreen")

			cmd := fmt.Sprintf("yabai -m space --move %v", (spaceIndex + 1))

			_, stderr, err := runShellSync(cmd)

			if err != nil {
				log.Printf("error: %v\n", err)

				fmt.Println("Stderr")
				fmt.Println(stderr)
			}
		}
	} else {

		// There is a chance that the app is full screen
		for i := len(windowQuery) - 1; i >= 0; i-- {
			elm := windowQuery[i]

			if elm.Fullscreen {
				fmt.Println("Fullscreen")

				cmd := fmt.Sprintf("yabai -m space --move %v", (spaceIndex + 1))

				_, stderr, err := runShellSync(cmd)

				if err != nil {
					log.Printf("error: %v\n", err)

					fmt.Println("Stderr")
					fmt.Println(stderr)
				}
				break
			}
		}

	}

	fmt.Fprintf(w, "prepare\n")
}

func StartServer() error {

	// Old url
	http.HandleFunc("/prepare", handleSpaceChange)

	fmt.Println("Starting server on PORT", 8090)

	return http.ListenAndServe(":8090", nil)
}
