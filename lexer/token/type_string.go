// Code generated by "stringer -type=Type"; DO NOT EDIT

package token

import "fmt"

const _Type_name = "IdentifierIntegerParenthesisCurlyBracketSquareBracketArrowCommaAssignPlusMinusSlashPercentEqualNotEqualGreaterLessGreaterEqualLessEqualIllegal"

var _Type_index = [...]uint8{0, 10, 17, 28, 40, 53, 58, 63, 69, 73, 78, 83, 90, 95, 103, 110, 114, 126, 135, 142}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
