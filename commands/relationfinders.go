package commands

import "geektrust/model"

type RelationFinder func(member *model.FamilyMember) ([]*model.FamilyMember, error)

func GetSiblings(name string, family *model.Family) ([]*model.FamilyMember, error) {
	finder := func(member *model.FamilyMember) ([]*model.FamilyMember, error) {
		return member.GetSiblings(model.Female, model.Male), nil
	}
	return DoIfMemberFound(name, family, finder)
}

func GetPaternalUncle(name string, family *model.Family) ([]*model.FamilyMember, error) {
	finder := func(member *model.FamilyMember) ([]*model.FamilyMember, error) {
		return member.Father.GetFatherSideSiblings(model.Male), nil
	}
	return DoIfMemberFound(name, family, finder)
}

func GetMaternalUncle(name string, family *model.Family) ([]*model.FamilyMember, error) {
	finder := func(member *model.FamilyMember) ([]*model.FamilyMember, error) {
		return member.Mother.GetMotherSideSiblings(model.Male), nil
	}
	return DoIfMemberFound(name, family, finder)
}

func GetPaternalAunt(name string, family *model.Family) ([]*model.FamilyMember, error) {
	finder := func(member *model.FamilyMember) ([]*model.FamilyMember, error) {
		return member.Father.GetFatherSideSiblings(model.Female), nil
	}
	return DoIfMemberFound(name, family, finder)
}

func GetMaternalAunt(name string, family *model.Family) ([]*model.FamilyMember, error) {
	finder := func(member *model.FamilyMember) ([]*model.FamilyMember, error) {
		return member.Mother.GetMotherSideSiblings(model.Female), nil
	}
	return DoIfMemberFound(name, family, finder)
}

func GetSisterInLaw(name string, family *model.Family) ([]*model.FamilyMember, error) {
	return GetXInLaws(name, family, model.Female)
}

func GetBrotherInLaw(name string, family *model.Family) ([]*model.FamilyMember, error) {
	return GetXInLaws(name, family, model.Male)
}

func GetXInLaws(name string, family *model.Family, gender model.Gender) ([]*model.FamilyMember, error) {
	finder := func(member *model.FamilyMember) ([]*model.FamilyMember, error) {
		var inLaws []*model.FamilyMember
		if member.Spouse != nil {
			inLaws = append(inLaws, member.Spouse.GetSiblings(gender)...)
		}
		for _, sibling := range member.GetSiblings(model.Female, model.Male) {
			if sibling.Spouse != nil && sibling.Spouse.Gender == gender {
				inLaws = append(inLaws, sibling.Spouse)
			}
		}
		return inLaws, nil
	}
	return DoIfMemberFound(name, family, finder)
}

func GetDaughters(name string, family *model.Family) ([]*model.FamilyMember, error) {
	finder := func(member *model.FamilyMember) ([]*model.FamilyMember, error) {
		return member.GetDaughters()
	}
	return DoIfMemberFound(name, family, finder)
}

func GetSons(name string, family *model.Family) ([]*model.FamilyMember, error) {
	finder := func(member *model.FamilyMember) ([]*model.FamilyMember, error) {
		return member.GetSons()
	}
	return DoIfMemberFound(name, family, finder)
}

func DoIfMemberFound(name string, family *model.Family, finder RelationFinder) ([]*model.FamilyMember, error) {
	member, err := family.FindMemberByName(name)
	if err != nil {
		return nil, err
	}
	return finder(member)
}
