package types

type Error struct {
	msg string
}

func (x *Error) Error() string { return x.msg }

func (x *Error) Is(err error) bool {
	if e, ok := err.(*Error); ok {
		return x.msg == e.msg
	}
	return false
}

func newError(msg string) *Error {
	return &Error{
		msg: msg,
	}
}

var (
	ErrConfigNoActionID   = newError("no action ID")
	ErrConfigNoActionName = newError("no action name")
	ErrConfigNoProbeID    = newError("no probe ID")
	ErrConfigNoProbeName  = newError("no probe name")

	ErrConfigConflictActionID = newError("conflict action ID")
	ErrConfigConflictProbeID  = newError("conflict probe ID")
	ErrConfigNoPolicyPath     = newError("no policy path")

	ErrActionInvalidConfig = newError("invalid action config")

	ErrNoSuchActionName = newError("no such action name")
	ErrNoSuchActionID   = newError("no such action ID")

	ErrInvalidScenario = newError("invalid play scenario")

	ErrInvalidHTTPRequest = newError("invalid HTTP request")

	ErrPolicyClientFailed = newError("policy client failed")
	ErrNoPolicyData       = newError("no policy data")
	ErrNoPolicyResult     = newError("no policy result")

	// runtime errors
	ErrMaxStackDepth = newError("max stack depth")
)
