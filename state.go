package state

import "time"

type roles int

const (
	user roles = iota
	hunters
	admins
)

type State struct {
	Id    int
	Name  string
	Desc  string
	Added time.Time
}

type Transition struct {
	Id        int
	Name      string
	Desc      string
	CurrState int
	ReqState  int
	User      roles
	Added     time.Time
}
