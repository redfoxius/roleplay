package combat

import "errors"

var (
	ErrInsufficientSteamPower = errors.New("insufficient steam power")
	ErrInvalidTarget          = errors.New("invalid target")
	ErrInvalidAction          = errors.New("invalid action")
)
