package env

import (
	"AOT/pkg/obj"
	params "AOT/pkg/parameters"
	types "AOT/pkg/types"
	utils "AOT/pkg/utilitaries"
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
					agentsToRemove = append(agentsToRemove, perceivedAgents[i])
				}
			}

		}
	}
	perceivedAgents = removeAgents(perceivedAgents, agentsToRemove)

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
		return false

	} else if utils.IsWithinWalls(agt.Pos()) && utils.IsOutOfScreen(*dummyObject) {
		return false

	} else if agt.Object().GetName() != types.BasicTitan1 || agt.Object().GetName() != types.BasicTitan2 {
		if utils.IsWithinWalls(agt.Pos()) && !utils.IsWithinWalls(dummyObject.TL()) {
			return false
		}
	}
	// Verify collisions with other agents of same type (human or titan)
	for i := range e.Agents() {
		if agt.Id() != e.Agents()[i].Id() &&
			e.Agents()[i].Agent().ObjectP().Life() > 0 &&
			utils.DetectCollision(e.Agents()[i].Object(), *dummyObject) {
			// Titan case
			if (agt.Object().GetName() == types.BasicTitan1 || agt.Object().GetName() == types.BasicTitan2) &&
				(e.Agents()[i].Agent().GetName() == types.BasicTitan1 || e.Agents()[i].Agent().GetName() == types.BasicTitan2) {
				return false
			}
			// Human case
			if (agt.Object().GetName() == types.MaleSoldier || agt.Object().GetName() == types.FemaleSoldier || agt.Object().GetName() == types.MaleCivilian || agt.Object().GetName() == types.FemaleCivilian || agt.Object().GetName() == types.Mikasa || agt.Object().GetName() == types.Eren) &&
				(e.Agents()[i].Agent().GetName() == types.MaleSoldier || e.Agents()[i].Agent().GetName() == types.FemaleSoldier || e.Agents()[i].Agent().GetName() == types.MaleCivilian || e.Agents()[i].Agent().GetName() == types.FemaleCivilian || e.Agents()[i].Agent().GetName() == types.Mikasa || e.Agents()[i].Agent().GetName() == types.Eren) {
				return false
			}
		}
	}

	for i := range e.Objects() {
		if utils.DetectCollision(*e.Objects()[i], *dummyObject) {
			// Human case
			if (agt.Object().GetName() == types.MaleSoldier || agt.Object().GetName() == types.FemaleSoldier || agt.Object().GetName() == types.MaleCivilian || agt.Object().GetName() == types.FemaleCivilian || agt.Object().GetName() == types.Mikasa || agt.Object().GetName() == types.Eren) &&
				(e.Objects()[i].GetName() == types.Dungeon || e.Objects()[i].GetName() == types.SmallHouse || e.Objects()[i].GetName() == types.BigHouse || e.Objects()[i].GetName() == types.Field) {
				return false
			}
		}
	}

	return true
}

func FirstValidPositionToCityCenter(agt AgentI, e *Environment) types.Position {
	neighbours := utils.GetNeighbors(agt.Pos(), agt.Agent().Speed())
	cityCenter := types.Position{X: (params.WallTLX + params.WallBRX) / 2, Y: (params.WallTLY + params.WallBRY) / 2}
	if len(neighbours) > 0 {
		neighbour := cityCenter.ClosestPosition(neighbours)
		agt.Agent().SetNextPos(neighbour)
		if IsNextPositionValid(agt, e) {
			return neighbour
		}
		neighbours = utils.RemovePosition(neighbours, neighbour)
	}
	return agt.Pos()
}
