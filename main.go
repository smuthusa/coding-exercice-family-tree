package main

import (
	"fmt"
	"geektrust/builder"
	"geektrust/datapopulator"
	"geektrust/model"
	"geektrust/parser"
	"geektrust/util"
	"os"
)

var logger = util.Println
var noOpLogger = util.NoOpPrint

func main() {
	if len(os.Args) != 2 {
		logger(fmt.Sprintf("missing file name argument"))
		os.Exit(-1)
	}
	shanFamily := builder.BuildRootFamily()
	err := datapopulator.BuildDefaultFamilyTree(shanFamily)
	if err != nil {
		logger(fmt.Sprintf("Error populating Shan family Tree!, error: %s", err.Error()))
	}
	fileBackedCommandParser := parser.NewFileBackedCommandParser(os.Args[1], &parser.CommandParser{})
	executeCommands(err, fileBackedCommandParser, shanFamily, logger)
}

func executeCommands(err error, fileBackedCommandParser *parser.FileBackedCommandParser, shanFamily *model.Family, logger func(str string)) {
	commands, err := fileBackedCommandParser.Parse()
	if err == nil {
		for _, command := range commands {
			err := command.Execute(shanFamily, logger)
			if err != nil {
				noOpLogger(fmt.Sprintf("Error execuring command %s - error: %s", command.String(), err.Error()))
			}
		}
	} else {
		noOpLogger(fmt.Sprintf("Error parsing a file %v", err))
	}
}
