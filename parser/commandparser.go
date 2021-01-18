package parser

import (
	"fmt"
	"geektrust/commands"
	"geektrust/errors"
	"geektrust/util"
	"strings"
)

const AddChild = "ADD_CHILD"
const GetRelationship = "GET_RELATIONSHIP"

type CommandParser struct {
	logger util.Logger
}

func (c *CommandParser) Parse(lines []string) ([]commands.Command, error) {
	var commandList []commands.Command
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		lineSplits := strings.Fields(line)
		if lineSplits[0] == AddChild {
			//ADD_CHILD Chitra Aria Female
			command, err := c.ParseAddCommand(lineSplits)
			if err != nil {
				c.logger(fmt.Sprintf("Error parsing line %v", lineSplits))
				continue
			}
			commandList = append(commandList, command)
		} else if GetRelationship == lineSplits[0] {
			//GET_RELATIONSHIP Lavnya Maternal-Aunt
			command := c.ParseGetRelationCommand(lineSplits)
			commandList = append(commandList, command)
		} else if "ADD_SPOUSE" == lineSplits[0] {
			command, err := c.ParseAddSpouse(lineSplits)
			if err != nil {
				c.logger(fmt.Sprintf("Error parsing line %v", lineSplits))
				continue
			}
			commandList = append(commandList, command)
		}
	}
	return commandList, nil
}

func (c *CommandParser) ParseAddCommand(lineColumns []string) (commands.Command, error) {
	if len(lineColumns) != 4 {
		c.logger(fmt.Sprintf("Invalid record %v", lineColumns))
		return nil, errors.InvalidInput
	}
	genderEnum, err := util.GetGenderEnumFromString(lineColumns[3])
	if err != nil {
		c.logger(fmt.Sprintf("Invalid gender value. %v", lineColumns))
		return nil, errors.InvalidInput
	}
	return &commands.AddChild{
		ParentName: lineColumns[1],
		ChildName:  lineColumns[2],
		Gender:     genderEnum,
	}, nil
}

func (c *CommandParser) ParseGetRelationCommand(lineColumns []string) commands.Command {
	return &commands.GetRelationship{
		ReferenceMemberName: lineColumns[1],
		Relationship:        lineColumns[2],
	}
}

func (c *CommandParser) ParseAddSpouse(lineColumns []string) (commands.Command, error) {
	if len(lineColumns) != 4 {
		c.logger(fmt.Sprintf("Invalid line %v", lineColumns))
		return nil, errors.InvalidInput
	}
	genderEnum, err := util.GetGenderEnumFromString(lineColumns[3])
	if err != nil {
		c.logger(fmt.Sprintf("Invalid line %v", lineColumns))
		return nil, errors.InvalidInput
	}
	return &commands.AddSpouse{
		PartnerName: lineColumns[1],
		SpouseName:  lineColumns[2],
		Gender:      genderEnum,
	}, nil
}
