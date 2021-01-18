package builder

import "geektrust/model"

func BuildRootFamily() *model.Family {
	kingShan := &model.FamilyMember{
		Name:   "Shan",
		Gender: model.Male,
	}
	queenAnga := &model.FamilyMember{
		Name:   "Anga",
		Gender: model.Female,
		Spouse: kingShan,
	}
	kingShan.Spouse = queenAnga
	shanFamily := &model.Family{
		Husband: kingShan,
		Wife:    queenAnga,
	}
	kingShan.Family = shanFamily
	queenAnga.Family = shanFamily
	return shanFamily
}