package hash

import "fmt"

// ErrSampleNonce is returned if the random nonce can't be sampled.
var ErrSampleNonce = fmt.Errorf("unable to sample random nonce")
