package state

type roles int

const (
	user roles = iota
	hunters
	admins
)

type Service struct {
	States 				[]*State
	Transitions 		map[Transition]*Transition
}

func NewService() *Service {
	s := &Service{
		
	}	
}

func (s *Service) IsValid(t *Transition) (string,error) {

}

type State struct {
	Name  string
}

func (s *Service) AddState() {

}

type Transition struct {
	CurrState	*State
	ReqState	*State
	User 		roles
}

func (s *Service) AddTransition() {

}







