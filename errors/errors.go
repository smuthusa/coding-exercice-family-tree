package errors

import "errors"

var MemberNotFound = errors.New("PERSON_NOT_FOUND")
var MemberCantHaveChildren = errors.New("PERSON_NOT_FOUND")
var AdditionFailed = errors.New("CHILD_ADDITION_FAILED")
var InvalidInput = errors.New("invalid_line")
var InvalidGender = errors.New("invalid_gender_value")