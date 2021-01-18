package commands

import (
	"fmt"
	"geektrust/util"
	"strings"

	"geektrust/model"
)

type Command interface {
	String() string
	Execute(family *model.Family, logger util.Logger) error
}

type GetRelationship struct {
	ReferenceMemberName string
	Relationship        string
}

type AddChild struct {
	ParentName string
	ChildName  string
	Gender     model.Gender
}

type AddSpouse struct {
	PartnerName string
	SpouseName  string
	Gender      model.Gender
}

func (s *AddSpouse) Execute(family *model.Family, logger util.Logger) error {
	member, err := family.FindMemberByName(s.PartnerName)
	if err != nil {
		logger("PERSON_NOT_FOUND")
		return err
	}
	var husband, wife *model.FamilyMember
	spouse := &model.FamilyMember{
		Name:   s.SpouseName,
		Gender: s.Gender,
		Spouse: member,
	}

	if s.Gender == model.Male {
		husband = spouse
		wife = member
	} else {
		husband = member
		wife = spouse
	}
	family = &model.Family{
		Husband: husband,
		Wife:    wife,
	}
	member.Spouse = spouse
	member.Family = family
	spouse.Family = family
	return nil
}

func (s *AddSpouse) String() string {
	return fmt.Sprintf("{ADD_SPOUSE: %s - %s}", s.PartnerName, s.SpouseName)
}

func (g *GetRelationship) Execute(family *model.Family, logger util.Logger) error {
	callback := func(nodes []*model.FamilyMember, err error) {
		if err != nil {
			logger(err.Error())
			return
		}
		var relationNames []string
		for _, node := range nodes {
			relationNames = append(relationNames, node.Name)
		}
		if len(relationNames) == 0 {
			logger("NONE")
		} else {
			logger(fmt.Sprintf("%s", strings.Join(relationNames, " ")))
		}
	}

	var relations []*model.FamilyMember
	var err error

	switch strings.ToUpper(g.Relationship) {
	case "PATERNAL-UNCLE":
		relations, err = GetPaternalUncle(g.ReferenceMemberName, family)
	case "MATERNAL-UNCLE":
		relations, err = GetMaternalUncle(g.ReferenceMemberName, family)
	case "MATERNAL-AUNT":
		relations, err = GetMaternalAunt(g.ReferenceMemberName, family)
	case "PATERNAL-AUNT":
		relations, err = GetPaternalAunt(g.ReferenceMemberName, family)
	case "SISTER-IN-LAW":
		relations, err = GetSisterInLaw(g.ReferenceMemberName, family)
	case "BROTHER-IN-LAW":
		relations, err = GetBrotherInLaw(g.ReferenceMemberName, family)
	case "SON":
		relations, err = GetSons(g.ReferenceMemberName, family)
	case "DAUGHTER":
		relations, err = GetDaughters(g.ReferenceMemberName, family)
	case "SIBLINGS":
		relations, err = GetSiblings(g.ReferenceMemberName, family)
	}
	callback(relations, err)
	return err
}

func (g *GetRelationship) String() string {
	return fmt.Sprintf("{GET_RELATIONSHIP: %s - %s}", g.ReferenceMemberName, g.Relationship)
}

func (c *AddChild) Execute(family *model.Family, logger util.Logger) error {
	_, err := family.AddChild(c.ParentName, c.ChildName, c.Gender)
	if err == nil {
		logger("CHILD_ADDITION_SUCCEEDED")
	} else {
		logger(err.Error())
	}
	return err
}

func (c *AddChild) String() string {
	return fmt.Sprintf("{ADD_CHILD: %s - %s - %v}", c.ParentName, c.ChildName, c.Gender)
}
