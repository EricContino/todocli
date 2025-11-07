package main

import (
	"fmt"
	"strconv"

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

func handlerDone(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("missing argument: provide a task id")
	}

	if len(cmd.Args) > 1 {
		return fmt.Errorf("invalid arguments: too many arguments provided, only provide task id")
	}

	// Check arg1
	taskId, err := strconv.Atoi(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid argument: '%s' is not an integer: %v", cmd.Args[0], err)
	}

	if taskId < 0 || taskId >= len(s.cfg.Tasks) {
		return fmt.Errorf("invalid argument: %d is not a valid task id", taskId)
	}

	fmt.Println(s.cfg.Tasks)
	s.cfg.Tasks = append(s.cfg.Tasks[:taskId], s.cfg.Tasks[taskId+1:]...)
	fmt.Println(s.cfg.Tasks)
	config.Write(*s.cfg)

	return nil
}
