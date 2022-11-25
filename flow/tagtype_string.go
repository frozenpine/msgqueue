// Code generated by "stringer -type TagType -linecomment"; DO NOT EDIT.

package flow

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Size8_T-1]
	_ = x[Size16_T-2]
	_ = x[Size32_T-4]
	_ = x[Size64_T-8]
	_ = x[FixedSize_T-16]
	_ = x[VariantSize_T-32]
}

const (
	_TagType_name_0 = "one byte typetwo byte type"
	_TagType_name_1 = "four byte type"
	_TagType_name_2 = "eight byte type"
	_TagType_name_3 = "fixed len bytes"
	_TagType_name_4 = "variant len bytes"
)

var (
	_TagType_index_0 = [...]uint8{0, 13, 26}
)

func (i TagType) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _TagType_name_0[_TagType_index_0[i]:_TagType_index_0[i+1]]
	case i == 4:
		return _TagType_name_1
	case i == 8:
		return _TagType_name_2
	case i == 16:
		return _TagType_name_3
	case i == 32:
		return _TagType_name_4
	default:
		return "TagType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
