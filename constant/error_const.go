package constant

import "errors"

var MUL_ADDRESS_NUMBER_VERIFY_ERROR = errors.New("failed to verify the number of multi-signed addresses")

var MUL_SCRIPT_NUMBER_ERROR = errors.New("wrong number of public keys in multi-signature script")

var HD_WALLET_GENERATE_ERROR = errors.New("hd wallet generate failure")

var MUL_ADDRESS_GENERATE_ERROR = errors.New("multi-signed addresses generate failure")
