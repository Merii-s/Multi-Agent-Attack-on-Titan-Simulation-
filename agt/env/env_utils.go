package env

import (
	"AOT/pkg/obj"
	types "AOT/pkg/types"
	utils "AOT/pkg/utilitaries"
	"fmt"
	"math/rand"
	"time"
)

func RemoveNoSeeableAgents(perceivedAgents []*AgentI, noSeeableSquaresBehindObjects map[*obj.Object][]types.Position) []*AgentI {
	// Filter out positions behind an obstacle if the center of the object is in the positionsBehindObjects list
	if len(noSeeableSquaresBehindObjects) == 0 {
		return perceivedAgents
	}
	agentsToRemove := []*AgentI{}
	for i, agt := range perceivedAgents {
		for _, noSeeableBox := range noSeeableSquaresBehindObjects {
			if len(noSeeableBox) > 0 {
				if utils.IntersectSquare(noSeeableBox[0], noSeeableBox[1], (*agt).Agent().ObjectP().Hitbox()[0], (*agt).Agent().ObjectP().Hitbox()[1]) {
					fmt.Println("Can't see :", (*agt).Agent().ObjectP().Name(), "at", (*agt).Agent().ObjectP().TL(), "because of", noSeeableBox)
					agentsToRemove = append(agentsToRemove, perceivedAgents[i])
				}
			}

		}
	}
	perceivedAgents = removeAgents(perceivedAgents, agentsToRemove)
	//fmt.Println("Agents perceived after removed : ", len(perceptedAgents))

	return perceivedAgents
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
		fmt.Println(agt.Id(), " can't go out of screen")
		return false

	} else if utils.IsWithinWalls(agt.Pos()) && utils.IsOutOfScreen(*dummyObject) {
		fmt.Println(agt.Id(), " can't go out of screen when is in walls")
		return false

	} else if agt.Object().GetName() != types.BasicTitan1 || agt.Object().GetName() != types.BasicTitan2 {
		if utils.IsWithinWalls(agt.Pos()) && !utils.IsWithinWalls(dummyObject.TL()) {
			fmt.Println(agt.Id(), " can't go out of walls when is in")
			return false
		}
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

func FirstValidPosition(agt AgentI, e *Environment) types.Position {
	neighbours := utils.GetNeighbors(agt.Pos(), agt.Agent().Speed())
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	neighbour := neighbours[random.Intn(len(neighbours))]
	for i := 0; i < 10; i++ {
		agt.Agent().SetNextPos(neighbour)
		if IsNextPositionValid(agt, e) {
			fmt.Println(agt.Id(), " New valid position", neighbour)
			return neighbour
		}
		neighbour = neighbours[random.Intn(len(neighbours))]
	}
	return agt.Pos()
}
