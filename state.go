package state
import (
	"errors"
	"fmt"
)

type roles int

//NOTE: role Internal added to make explicit regular service generated
//		transitions.
const (
	User roles = iota
	Hunter
	Admin
	Internal
)

type Service struct {
	States 				[]*State
	Transitions 		map[Transition]bool
}

func NewService() *Service {
	t := make(map[Transition]bool)
	s := &Service{Transitions: t}
	return s
}

func (service *Service) IsValid(t Transition) error {
	//super users bypass this logic
	if t.Role == Admin { return nil }

	//known allowed transition
	if _,ok := service.Transitions[t]; ok {	return nil }
	
	//TODO: enum value of role not descriptive enough
	//unknown transition
	return errors.New(fmt.Sprintf(
		"Transition from state \"%s\" to \"%s\" using role %v failed",
		t.CurrState.Name,t.ReqState.Name, t.Role))	
}

type State struct {
	Name  string
}

func (service *Service) AddState(state *State) {
	service.States=append(service.States,state)
}

type Transition struct {
	CurrState	*State
	ReqState	*State
	Role 		roles
}

func (service *Service) AddTransition(t Transition) {
	service.Transitions[t] = true
}