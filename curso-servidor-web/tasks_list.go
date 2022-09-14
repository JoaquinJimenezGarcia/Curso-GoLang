package main

import "fmt"

type tasksList struct {
	tasks []*task
}

type task struct {
	name        string
	description string
	completed   bool
}

func (my_list *tasksList) addToList(task *task) {
	my_list.tasks = append(my_list.tasks, task)
}

func (my_list *tasksList) deleteFromList(taskIndex int) {
	my_list.tasks = append(my_list.tasks[:taskIndex], my_list.tasks[taskIndex+1:]...)
}

func (my_task *task) markAsComplete() {
	my_task.completed = true
}

func (my_task *task) updateDescription(description string) {
	my_task.description = description
}

func (my_task *task) updateName(name string) {
	my_task.name = name
}

func main() {
	my_task := &task{
		name:        "Finish my GO course",
		description: "Getting my GO certification this week",
	}

	my_task2 := &task{
		name:        "Continue learning GO",
		description: "It's really interesting",
	}

	my_task3 := &task{
		name:        "See how can I mix it with networking",
		description: "Networking and security are my passion.",
	}

	fmt.Printf("%+v\n", my_task)

	my_task.markAsComplete()
	my_task.updateDescription("I got it!")

	fmt.Printf("%+v\n", my_task)

	fmt.Println("Showing now list")

	my_tasksLists := &tasksList{
		tasks: []*task{
			my_task,
		},
	}

	my_tasksLists.addToList(my_task2)
	my_tasksLists.addToList(my_task3)

	fmt.Printf("%+v\n", my_tasksLists.tasks)

	my_tasksLists.deleteFromList(1)

	fmt.Printf("%+v\n", my_tasksLists.tasks)
}
