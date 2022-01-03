package errorHelper

import "errors"

var DuplicateEmailRegister = errors.New("email already registered")
var OldPasswordNotMatch = errors.New("old password mismatch")
