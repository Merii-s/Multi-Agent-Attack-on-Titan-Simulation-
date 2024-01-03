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
		if utils.Contains(positionsBehindObjects, (*agt).Agent().ObjectP().Center()) {
			fmt.Println("Can't see Agent : ", (*agt).Id(), "at", (*agt).Agent().Pos(), "because of an object")
			agentsToRemove = append(agentsToRemove, agt)
		}
	}

	// Remove the objects in the objectsToRemove list from the perceivedObjects list
	//fmt.Println("Agents to remove : ", len(agentsToRemove))
	//fmt.Println("Agents perceived before : ", len(perceptedAgents))
	//fmt.Println("Agents to remove : ", agentsToRemove)
	perceptedAgents = removeAgents(perceptedAgents, agentsToRemove)
	//fmt.Println("Agents perceived after removed : ", len(perceptedAgents))

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
	if !utils.IsOutOfScreen(agt.Object()) && utils.IsOutOfScreen(*dummyObject) {
		return false

	} else if utils.IsWithinWalls(agt.Pos()) && utils.IsOutOfScreen(*dummyObject) {
		return false

	} else if utils.IsWithinWalls(agt.Pos()) && !utils.IsWithinWalls(dummyObject.TL()) {
		return false
	}

	// Verify collisions with other agents of same type (human or titan)
	for i := range e.Agents() {
		if agt.Id() != e.Agents()[i].Id() &&
			utils.DetectCollision(e.Agents()[i].Object(), *dummyObject) {
			// Titan case
			if (agt.Object().GetName() == types.BasicTitan1 || agt.Object().GetName() == types.BasicTitan2) &&
				(e.Agents()[i].Agent().GetName() == types.BasicTitan1 || e.Agents()[i].Agent().GetName() == types.BasicTitan2) {
				fmt.Println("COLLISION DETECTED 2 titans :", e.Agents()[i].Agent().GetName(), "at", e.Agents()[i].Agent().Pos())
				return false
			}
			// Human case
			if (agt.Object().GetName() == types.MaleSoldier || agt.Object().GetName() == types.FemaleSoldier || agt.Object().GetName() == types.MaleCivilian || agt.Object().GetName() == types.FemaleCivilian || agt.Object().GetName() == types.Mikasa || agt.Object().GetName() == types.Eren) &&
				(e.Agents()[i].Agent().GetName() == types.MaleSoldier || e.Agents()[i].Agent().GetName() == types.FemaleSoldier || e.Agents()[i].Agent().GetName() == types.MaleCivilian || e.Agents()[i].Agent().GetName() == types.FemaleCivilian || e.Agents()[i].Agent().GetName() == types.Mikasa || e.Agents()[i].Agent().GetName() == types.Eren) {
				fmt.Println("COLLISION DETECTED 2 humans :", e.Agents()[i].Agent().GetName(), "at", e.Agents()[i].Agent().Pos())
				return false
			}
		}
	}

	for i := range e.Objects() {
		if utils.DetectCollision(*e.Objects()[i], *dummyObject) {
			// Human case
			if (agt.Object().GetName() == types.MaleSoldier || agt.Object().GetName() == types.FemaleSoldier || agt.Object().GetName() == types.MaleCivilian || agt.Object().GetName() == types.FemaleCivilian || agt.Object().GetName() == types.Mikasa || agt.Object().GetName() == types.Eren) &&
				(e.Objects()[i].GetName() == types.Dungeon || e.Objects()[i].GetName() == types.SmallHouse || e.Objects()[i].GetName() == types.BigHouse || e.Objects()[i].GetName() == types.Field) {
				fmt.Println("COLLISION DETECTED 1 titan and 1 object :", e.Objects()[i].GetName(), "at", e.Objects()[i].TL())
				return false
			}
		}
	}

	return true
}
