package parser

import (
	"geektrust/commands"
	"geektrust/util"
	"io/ioutil"
	"os"
	"strings"
)

type FileBackedCommandParser struct {
	filePath      string
	commandParser *CommandParser
}

func NewFileBackedCommandParser(filePath string, commandProcessor *CommandParser) *FileBackedCommandParser {
	return &FileBackedCommandParser{
		filePath:      filePath,
		commandParser: commandProcessor,
	}
}

func (p *FileBackedCommandParser) Parse() ([]commands.Command, error) {
	if _, err := os.Stat(p.filePath); p.filePath == "" || err != nil {
		return nil, os.ErrNotExist
	}
	data, err := ioutil.ReadFile(p.filePath)
	if err != nil {
		return nil, err
	}
	contents := string(data)
	lines := strings.Split(contents, util.LineBreak)
	return p.commandParser.Parse(lines)
}
