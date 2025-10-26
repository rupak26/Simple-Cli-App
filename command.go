package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add       string 
    Del       int 
	Edit      string
	Toggle    int 
	List      bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add , "add" , "" , "Add a new todo specify title")
	flag.StringVar(&cf.Edit , "Edit" , "" , "Edit a todo by index & specify an new title. id:new_title")
	flag.IntVar(&cf.Del , "Del" , -1 , "Spacify a todo by index to delete")
	flag.IntVar(&cf.Toggle , "toggel" , -1 , "Spacify a todo by index to toggle")
	flag.BoolVar(&cf.List , "list" , false , "List all todos")
    
	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Excute (todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit , ":" , 2)
		if len(parts) != 2 {
			fmt.Println("Invalid formate to edit . Please put user_id:new_title")
			os.Exit(1)
		}
        index , err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index to edit")
			os.Exit(1)
		}
		todos.edit(index , parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	default:
		fmt.Println("invalid command")
	}
}