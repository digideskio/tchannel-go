// generated by stringer -type=RetryOn; DO NOT EDIT

package tchannel

import "fmt"

const _RetryOn_name = "RetryDefaultRetryConnectionErrorRetryNeverRetryNonIdempotentRetryUnexpected"

var _RetryOn_index = [...]uint8{0, 12, 32, 42, 60, 75}

func (i RetryOn) String() string {
	if i < 0 || i+1 >= RetryOn(len(_RetryOn_index)) {
		return fmt.Sprintf("RetryOn(%d)", i)
	}
	return _RetryOn_name[_RetryOn_index[i]:_RetryOn_index[i+1]]
}