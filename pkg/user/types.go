package user

type Mode string

const (
	ModeUnHiddenOnly            Mode = "unhidden_only"
	ModeHiddenOnly              Mode = "hidden_only"
	ModeHiddenAllowAutoUnhide   Mode = "hidden_allow_auto_unhide"
	ModeHiddenPreventAutoUnhide Mode = "hidden_prevent_auto_unhide"
	ModeAll                     Mode = "all"
)

type State string

const (
	StateJoined  State = "joined"
	StateInvited State = "invited"
)

type SuperMode string

const (
	SuperModeAll      SuperMode = "all"
	SuperModeSuper    SuperMode = "super"
	SuperModeNonSuper SuperMode = "nonsuper"
)
