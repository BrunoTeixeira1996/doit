package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

func usage() {
	help := map[string]string{
		"help": "[shows help]",
		"list": "[lists all todos of a dir recursively being .md the default extension] [-type][-help]",
	}

	keys := make([]string, 0)

	for k, _ := range help {
		keys = append(keys, k)
	}

	// sorting the map into a slice
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, help[k])
	}
}

func main() {

	list_command := flag.NewFlagSet("list", flag.ExitOnError)

	list_text_ptr := list_command.String("type", ".md", "List file by extension. If more than one, use comma [.py,.go,.c]")
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		list_command.Parse(os.Args[2:])
	case "help":
		usage()
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
	todo_slice := []Todo{}
	if list_command.Parsed() {
		if string((*list_text_ptr)[0]) != "." {
			list_command.PrintDefaults()
			os.Exit(1)
		} else {
			// anonymous func inside WalkDir to walk on directories recursively
			// and call walk_on_files func to read files line by line
			filepath.WalkDir(".", func(s string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if !d.IsDir() {
					walk_on_files(s, *list_text_ptr, &todo_slice)
				}
				return nil
			})
			//sort slice of todos by urgency
			sort.SliceStable(todo_slice, func(i, j int) bool {
				return todo_slice[i].urgency > todo_slice[j].urgency
			})
			output_todo(todo_slice)
		}
	}
}
