package model

import (
	"geektrust/errors"
)

type Gender int

const (
	Male   = iota
	Female = iota
)

type FamilyMember struct {
	Name   string
	Gender Gender
	Father *FamilyMember
	Mother *FamilyMember
	Spouse *FamilyMember
	Family *Family
}

type Family struct {
	Husband  *FamilyMember
	Wife     *FamilyMember
	Children []*FamilyMember
}

func (f *Family) FindMemberByName(name string) (*FamilyMember, error) {
	if f.Husband.Name == name {
		return f.Husband, nil
	} else if f.Wife.Name == name {
		return f.Wife, nil
	}
	for _, child := range f.Children {
		if child.Name == name {
			return child, nil
		} else if child.Spouse != nil && child.Spouse.Name == name {
			return child.Spouse, nil
		} else if child.Family != nil {
			member, err := child.Family.FindMemberByName(name)
			if err != errors.MemberNotFound {
				return member, nil
			}
		}
	}
	return nil, errors.MemberNotFound
}

func (f *Family) AddChild(parentName string, name string, gender Gender) (*FamilyMember, error) {
	member, err := f.FindMemberByName(parentName)
	if err != nil {
		return nil, err
	}
	if member.Family == nil {
		return nil, errors.MemberCantHaveChildren
	}
	if member.Gender == Male {
		return nil, errors.AdditionFailed
	}
	child := &FamilyMember{
		Name:   name,
		Gender: gender,
		Father: member.Family.Husband,
		Mother: member.Family.Wife,
	}
	member.Family.Children = append(member.Family.Children, child)
	return child, nil
}

func (f *FamilyMember) GetSiblings(genders ...Gender) []*FamilyMember {
	return f.GetSiblingsFrom(Male, genders...)
}

func (f *FamilyMember) GetMotherSideSiblings(genders ...Gender) []*FamilyMember {
	return f.GetSiblingsFrom(Female, genders...)
}

func (f *FamilyMember) GetFatherSideSiblings(genders ...Gender) []*FamilyMember {
	return f.GetSiblingsFrom(Female, genders...)
}

func (f *FamilyMember) GetSiblingsFrom(genderSide Gender, genders ...Gender) []*FamilyMember {
	var siblings []*FamilyMember
	var familyMember *FamilyMember
	if genderSide == Female {
		familyMember = f.Mother
	} else {
		familyMember = f.Father
	}
	if familyMember == nil {
		return siblings
	}
	for _, member := range familyMember.Family.Children {
		//It is okay to compare the references, as the records are not going to change throughout the runtime
		if member != f {
			for _, gender := range genders {
				if gender == member.Gender {
					siblings = append(siblings, member)
				}
			}
		}
	}
	return siblings
}

func (f *FamilyMember) GetDaughters() ([]*FamilyMember, error) {
	return f.GetChildByGender(Female)
}

func (f *FamilyMember) GetSons() ([]*FamilyMember, error) {
	return f.GetChildByGender(Male)
}

func (f *FamilyMember) GetChildByGender(gender Gender) ([]*FamilyMember, error) {
	if f.Family == nil {
		return nil, nil
	}
	var children []*FamilyMember
	if f.Family != nil && f.Family.Children != nil {
		for _, child := range f.Family.Children {
			if child.Gender == gender {
				children = append(children, child)
			}
		}
	}
	return children, nil
}
