package env

import (
	"AOT/pkg/obj"
	types "AOT/pkg/types"
	utils "AOT/pkg/utilitaries"
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
	for _, a := range e.Agents() {
		if utils.DetectCollision(a.Object(), *dummyObject) {
			return false
		}
	}
	return true
}
