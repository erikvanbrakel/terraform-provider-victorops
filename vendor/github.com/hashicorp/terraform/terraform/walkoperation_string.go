// Code generated by "stringer -type=walkOperation graph_walk_operation.go"; DO NOT EDIT.

package terraform

import "strconv"

const _walkOperation_name = "walkInvalidwalkInputwalkApplywalkPlanwalkPlanDestroywalkRefreshwalkValidatewalkDestroywalkImport"

var _walkOperation_index = [...]uint8{0, 11, 20, 29, 37, 52, 63, 75, 86, 96}

func (i walkOperation) String() string {
	if i >= walkOperation(len(_walkOperation_index)-1) {
		return "walkOperation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _walkOperation_name[_walkOperation_index[i]:_walkOperation_index[i+1]]
}
