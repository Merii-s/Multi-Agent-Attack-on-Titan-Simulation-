package env

import (
	types "AOT/pkg/types"
	utils "AOT/pkg/utilitaries"
)

func removeAgentsBehindPositions(perceptedAgents []AgentI, positionsBehindObjects []types.Position) []AgentI {
	// Filter out positions behind an obstacle if the center of the object is in the positionsBehindObjects list
	objectsToRemove := []AgentI{}
	for _, object := range perceptedAgents {
		if utils.Contains(positionsBehindObjects, object.Pos()) {
			objectsToRemove = append(objectsToRemove, object)
		}
	}

	// Remove the objects in the objectsToRemove list from the perceivedObjects list
	perceptedAgents = removeAgents(perceptedAgents, objectsToRemove)

	return perceptedAgents
}

func removeAgents(perceptedAgents []AgentI, objectsToRemove []AgentI) []AgentI {
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
