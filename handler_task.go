package main

import (
	"fmt"

	"github.com/EricContino/todocli/internal/config"
)

func handlerList(s *state, cmd command) error {

	fmt.Println("Task Id\tTask Description")
	fmt.Println("=======\t================")
	for index, value := range s.cfg.Tasks {
		fmt.Printf("%d\t%s\n", index, value)
	}
	return nil
}

func handlerAdd(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("missing argument: provide a task description")
	}

	if len(cmd.Args) > 1 {
		return fmt.Errorf("invalid number of argument: ensure description in wrapped in quotes")
	}

	newTask := cmd.Args[0]

	s.cfg.Tasks = append(s.cfg.Tasks, newTask)

	err := config.Write(*s.cfg)
	if err != nil {

	}

	return nil
}
