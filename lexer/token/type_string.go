// Code generated by "stringer -type=Type"; DO NOT EDIT

package token

import "fmt"

const _Type_name = "IdentifierKeywordIntegerParenthesisOpenParenthesisCloseCurlyBracketOpenCurlyBracketCloseSquareBracketOpenSquareBracketCloseArrowCommaAssignPlusMinusSlashMultiplyPercentEqualNotEqualGreaterLessGreaterEqualLessEqualIllegal"

var _Type_index = [...]uint8{0, 10, 17, 24, 39, 55, 71, 88, 105, 123, 128, 133, 139, 143, 148, 153, 161, 168, 173, 181, 188, 192, 204, 213, 220}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
