// Code generated by "stringer -type=Type"; DO NOT EDIT

package token

import "fmt"

const _Type_name = "IdentifierIntegerFloatParenthesisCurlyBracketSquareBracketAssignPlusMinusSlashPercentEqualNotEqualGreaterLessGreaterEqualLessEqualIllegal"

var _Type_index = [...]uint8{0, 10, 17, 22, 33, 45, 58, 64, 68, 73, 78, 85, 90, 98, 105, 109, 121, 130, 137}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
