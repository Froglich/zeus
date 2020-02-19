package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
)

type task struct {
	Name    string   `json:"name"`
	CMD     string   `json:"cmd"`
	Args    []string `json:"args"`
	Restart bool     `json:"restart"`
}

func (t *task) start() {
	fmt.Printf("Starting %s\n", t.Name)
	cmd := exec.Command(t.CMD, t.Args...)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\033[31;1m%s: %v\033[0m\n", t.Name, err)
	}
}

func (t *task) run(wg *sync.WaitGroup) {
	defer wg.Done()

	for t.Restart {
		t.start()
	}

	t.start()
}

func main() {
	c := os.Args[1]
	f, err := os.Open(c)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	config, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	tasks := make([]task, 0)
	err = json.Unmarshal(config, &tasks)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for x := range tasks {
		fmt.Printf("Adding handler for %s\n", tasks[x].Name)
		wg.Add(1)
		go tasks[x].run(&wg)
	}

	wg.Wait()
	fmt.Println("Zeus is exiting.")
}
