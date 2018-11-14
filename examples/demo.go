package main

import (
	"fmt"
	"github.com/johannesostling/state"
)

func main() {	
	service := state.NewService()

	stateReady := &state.State{Name: "Ready"}
	stateRiding := &state.State{Name: "Riding"}
	stateBatteryLow := &state.State{Name: "Battery-low"}
	stateBounty := &state.State{Name: "Bounty"}
	stateCollected := &state.State{Name: "Collected"}
	stateDropped := &state.State{Name: "Dropped"}
	stateUnknown := &state.State{Name: "Unknown"}

	service.AddState(stateReady)
	service.AddState(stateRiding)
	service.AddState(stateBatteryLow)
	service.AddState(stateBounty)
	service.AddState(stateCollected)
	service.AddState(stateDropped)
	service.AddState(stateUnknown)

	transReadyRidingUser :=
		state.Transition{stateReady, stateRiding, state.User}
	transReadyRidingHunter :=
		state.Transition{stateReady, stateRiding, state.Hunter}
	service.AddTransition(transReadyRidingUser)
	service.AddTransition(transReadyRidingHunter)

	transRidingReadyUser :=
		state.Transition{stateRiding, stateReady, state.User}
	transRidingReadyHunter :=
		state.Transition{stateRiding, stateReady, state.Hunter}
	service.AddTransition(transRidingReadyUser)
	service.AddTransition(transRidingReadyHunter)

	transRidingBatteryLowInternal :=
		state.Transition{stateRiding, stateBatteryLow, state.Internal}
	service.AddTransition(transRidingBatteryLowInternal)

	transBatteryLowBountyInternal :=
		state.Transition{stateBatteryLow, stateBounty, state.Internal}
	service.AddTransition(transBatteryLowBountyInternal)

	transBountyCollectedHunter :=
		state.Transition{stateBounty, stateCollected, state.Hunter}
	service.AddTransition(transBountyCollectedHunter)

	transCollectedDroppedHunter :=
		state.Transition{stateCollected, stateDropped, state.Hunter}
	service.AddTransition(transCollectedDroppedHunter)

	transDroppedReadyHunter :=
		state.Transition{stateDropped, stateReady, state.Hunter}
	service.AddTransition(transDroppedReadyHunter)

	transReadyBountyInternal :=
		state.Transition{stateReady,stateBounty,state.Internal}
	service.AddTransition(transReadyBountyInternal)

	transReadyUnknownInternal :=
		state.Transition{stateReady,stateUnknown,state.Internal}
	service.AddTransition(transReadyUnknownInternal)




	myExample1 :=
		state.Transition{stateBounty,stateCollected,state.User}
	err := service.IsValid(myExample1)
	fmt.Println(err)

	myExample2 :=
		state.Transition{stateBounty,stateCollected,state.Hunter}
	err = service.IsValid(myExample2)
	fmt.Println(err)
}
