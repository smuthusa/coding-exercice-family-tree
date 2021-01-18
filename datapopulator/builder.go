package datapopulator

import (
	"geektrust/model"
	"geektrust/parser"
	"geektrust/util"
)

func BuildDefaultFamilyTree(shanFamily *model.Family) error {
	memberCommands := getDefaultMembers()
	cmdParser := parser.CommandParser{}
	commandList, err := cmdParser.Parse(memberCommands)
	if err != nil {
		return err
	}
	for _, command := range commandList {
		err := command.Execute(shanFamily, util.NoOpPrint)
		if err != nil {
			return err
		}
	}
	return nil
}

func getDefaultMembers() []string {
	return []string{
		"ADD_CHILD Anga Chit Male",
		"ADD_CHILD Anga Ish Male",
		"ADD_CHILD Anga Vich Male",
		"ADD_CHILD Anga Aras Male",
		"ADD_CHILD Anga Satya Female",
		"ADD_SPOUSE Chit Amba Female",
		"ADD_SPOUSE Vich Lika Female",
		"ADD_SPOUSE Aras Chitra Female",
		"ADD_SPOUSE Satya Vyan Male",
		"ADD_CHILD Amba Dritha Female",
		"ADD_SPOUSE Dritha Jaya Male",
		"ADD_CHILD Amba Tritha Female",
		"ADD_CHILD Amba Vritha Male",
		"ADD_CHILD Lika Vila Female",
		"ADD_CHILD Lika Chika Female",
		"ADD_CHILD Chitra Jnki Female",
		"ADD_SPOUSE Jnki Arit Male",
		"ADD_CHILD Chitra Ahit Male",
		"ADD_CHILD Satya Asva Male",
		"ADD_CHILD Satya Vyas Male",
		"ADD_CHILD Satya Atya Female",
		"ADD_SPOUSE Asva Satvy Female",
		"ADD_SPOUSE Vyas Krpi Female",
		"ADD_CHILD Dritha Yodhan Male",
		"ADD_CHILD Jnki Laki Male",
		"ADD_CHILD Jnki Lavnya Female",
		"ADD_CHILD Satvy Vasa Male",
		"ADD_CHILD Krpi Kriya Male",
		"ADD_CHILD Krpi Krithi Female",
	}
}
