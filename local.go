package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const regex = `(#|\/\/|;;|%)((?i) todo|todo)`

/*
//todo -> yes
// TODO -> yes
/& TODO -> no
;;todo -> yes
#todo -> yes
#TODO -> yes
(TODO) -> no
todo -> no
*/

type Todo struct {
	filename string
	word     string
	line     int
}

// outputs a Todo struct
func (todo Todo) output_string() string {
	return fmt.Sprintf("%s:%d: %s",
		todo.filename, todo.line, todo.word)
}

// fills a Todo struct
func (todo *Todo) fill_struct(path string, line int, word string) {
	todo.filename = path
	todo.line = line
	todo.word = word
}

// checks if that file has any provided extension
func check_if_file_has_extension(splited_extensions []string, file_extension string) bool {
	for _, extension := range splited_extensions {
		if file_extension == extension {
			return true
		}
	}
	return false
}

// reads files line by line and match TODO keyword
func walk_on_files(path string, user_file_extension string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Error opening the file")
	}
	defer file.Close()

	splited_extensions := strings.Split(user_file_extension, ",")
	file_extension := filepath.Ext(path)

	todo_slice := []Todo{}
	todo := Todo{}
	line := 0
	if check_if_file_has_extension(splited_extensions, file_extension) {
		scanner := bufio.NewScanner(file)
		line = 1
		for scanner.Scan() {
			match := regexp.MustCompile(regex)
			if match.MatchString(scanner.Text()) {
				todo.fill_struct(path, line, scanner.Text())
				todo_slice = append(todo_slice, todo)
			}
			line++
		}

	}

	for _, todo := range todo_slice {
		fmt.Println(todo.output_string())
	}
}
