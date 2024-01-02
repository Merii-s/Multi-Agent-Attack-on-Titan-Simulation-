package env

import (
	"AOT/pkg/obj"
	types "AOT/pkg/types"
	utils "AOT/pkg/utilitaries"
	"fmt"
)

func removeAgentsBehindPositions(perceptedAgents []*AgentI, positionsBehindObjects []types.Position) []*AgentI {
	// Filter out positions behind an obstacle if the center of the object is in the positionsBehindObjects list
	agentsToRemove := []*AgentI{}

	for _, agt := range perceptedAgents {
		if utils.Contains(positionsBehindObjects, (*agt).Pos()) {
			agentsToRemove = append(agentsToRemove, agt)
		}
	}

	// Remove the objects in the objectsToRemove list from the perceivedObjects list
	perceptedAgents = removeAgents(perceptedAgents, agentsToRemove)

	return perceptedAgents
}

func removeAgents(perceptedAgents []*AgentI, objectsToRemove []*AgentI) []*AgentI {
	// Remove the objects in the objectsToRemove list from the perceptedObjects list
	for _, object := range objectsToRemove {
		for i, perceptedObject := range perceptedAgents {
			if perceptedObject == object {
				perceptedAgents = append(perceptedAgents[:i], perceptedAgents[i+1:]...)
			}
		}
	}

	return perceptedAgents
}

func IsNextPositionValid(agt AgentI, e *Environment) bool {
	dummyObject := obj.NewObject(agt.Agent().GetName(), agt.Agent().NextPosition(), agt.Agent().ObjectP().Life())
	if !utils.IsOutOfScreen(agt.Pos()) && utils.IsOutOfScreen(dummyObject.TL()) {
		return false
	}
	for i := range e.Agents() {
		if agt.Id() != e.Agents()[i].Id() && utils.DetectCollision(e.Agents()[i].Object(), *dummyObject) {
			fmt.Println("COLLISION DETECTED with", e.Agents()[i].Agent().GetName(), "at", e.Agents()[i].Agent().Pos())
			return false
		}
	}
	return true
}
