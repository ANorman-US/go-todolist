package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const help_text string = `+---------------------------+
List of commands and syntax:
    add <list item>
        (adds a new item to the todo list)
    remove <# index>
        (removes the item at # index)
    show
        (prints contents of list)
    help
        (prints this list of commands and syntax)
    quit
        (exits the program)
+---------------------------+`

func add_task(todo_list []string, task string) []string {
	list := append(todo_list, task)
	return list
}

func remove_task(todo_list []string, index int) []string {
	list := append(todo_list[:index-1], todo_list[index:]...)
	return list
}

func show_tasks(todo_list []string) {
	for index := range len(todo_list) {
		fmt.Println(index+1, ":", todo_list[index])
	}
}

func show_help() {
	fmt.Println(help_text)
}

func main() {

	var todo_list []string
	fmt.Println("Welcome to the Todo List! Type \"help\" to get started...")
loop:
	for {
		fmt.Println("Enter Command")

		var token string

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		var user_line []string = strings.SplitN(scanner.Text(), " ", 2)
		token = user_line[0]

		err := scanner.Err()

		if err != nil {
			fmt.Println(err)
		} else {
			switch token {
			case "add":
				if len(user_line) != 2 || user_line[1] == "" {
					fmt.Println("Missing Argument")
				} else {
					todo_list = add_task(todo_list, user_line[1])
				}
			case "remove":
				if len(user_line) != 2 || user_line[1] == "" {
					fmt.Println("Missing Index")
				} else {
					index, err := strconv.Atoi(user_line[1])
					if err != nil {
						fmt.Println("Error converting index")
					} else if index > len(todo_list) || index < 1 {
						fmt.Println("Index out of range")
					} else {
						todo_list = remove_task(todo_list, index)
					}
				}
			case "show":
				show_tasks(todo_list)
			case "help":
				show_help()
			case "quit":
				break loop
			default:
				fmt.Println("Invalid Command")
			}
		}
	}
}
