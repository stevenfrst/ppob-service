package errorHelper

import "errors"

var DuplicateEmailRegister = errors.New("email already registered")
var OldPasswordNotMatch = errors.New("old password mismatch")
var DuplicateVoucher = errors.New("voucher already exists")
var ErrRecordNotFound = errors.New("data not found")
var ErrVoucherNotMatch = errors.New("voucher not match")
