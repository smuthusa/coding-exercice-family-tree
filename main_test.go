package main

import (
	"geektrust/builder"
	"geektrust/commands"
	"geektrust/parser"
	"geektrust/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

var commandInputs = []string{
	"ADD_CHILD Shan Ganga Female",
	"ADD_SPOUSE Ganga Samy Male",
	"ADD_CHILD Ganga Bhavani Female",
	"ADD_CHILD Ganga Cauvery Female",
	"ADD_SPOUSE Cauvery Krishna Male",
	"ADD_CHILD Cauvery Narmadha Female",
	"ADD_CHILD Cauvery Kumar Male",
}

func Test_AddChild(t *testing.T) {
	shanFamily := builder.BuildRootFamily()
	commandList := parseCommands(t, commandInputs)
	for _, c := range commandList {
		err := c.Execute(shanFamily, util.Println)
		if err != nil {
			t.Error(err)
		}
	}
	var sibling string
	commandList = parseCommands(t, []string{"GET_RELATIONSHIP Narmadha Siblings"})
	commandList[0].Execute(shanFamily, func(siblingRes string) {
		sibling = siblingRes
	})
	assert.Equal(t, "Kumar", sibling)

	commandList = parseCommands(t, []string{"GET_RELATIONSHIP Kumar Siblings"})
	commandList[0].Execute(shanFamily, func(siblingRes string) {
		sibling = siblingRes
	})
	assert.Equal(t, "Narmadha", sibling)

	commandList = parseCommands(t, []string{"GET_RELATIONSHIP Krishna Siblings"})
	commandList[0].Execute(shanFamily, func(siblingRes string) {
		sibling = siblingRes
	})
	assert.Equal(t, "", sibling)

}

func Test_AddSpouse(t *testing.T) {
	shanFamily := builder.BuildRootFamily()
	commandList := parseCommands(t, commandInputs)
	for _, c := range commandList {
		err := c.Execute(shanFamily, util.Println)
		if err != nil {
			t.Error(err)
		}
	}
	var sibling string
	commandList = parseCommands(t, []string{"GET_RELATIONSHIP Narmadha Siblings"})
	commandList[0].Execute(shanFamily, func(siblingRes string) {
		sibling = siblingRes
	})
	assert.Equal(t, "Kumar", sibling)

	commandList = parseCommands(t, []string{"GET_RELATIONSHIP Kumar Siblings"})
	commandList[0].Execute(shanFamily, func(siblingRes string) {
		sibling = siblingRes
	})
	assert.Equal(t, "Narmadha", sibling)
}

func parseCommands(t *testing.T, commandInputs []string) []commands.Command {
	cmdParser := &parser.CommandParser{}
	commandList, err := cmdParser.Parse(commandInputs)
	if err != nil {
		t.Error(err)
	}
	return commandList
}
