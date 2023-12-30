package agtutils

import (
	env "AOT/agt/env"
	hagt "AOT/agt/humans"
	tagt "AOT/agt/titans"
	obj "AOT/pkg/obj"
	params "AOT/pkg/parameters"
	types "AOT/pkg/types"
	utils "AOT/pkg/utilitaries"
)

func PlaceHuman(objs []obj.Object, humans []env.AgentI, human env.AgentI, tl_village types.Position, br_village types.Position) []env.AgentI {
	end := false

	for !end {
		counter := 0
		nb_grass := 0
		for _, object := range objs {
			if object.Name() != types.Grass {
				if utils.DetectCollision(human.Object(), object) {
					x, y := utils.GetRandomCoords(tl_village, br_village)
					human.SetPos(types.Position{X: x, Y: y})
					break
				} else {
					counter++
				}
			} else {
				nb_grass++
			}
		}

		for _, hu := range humans {
			if utils.DetectCollision(human.Object(), hu.Object()) {
				x, y := utils.GetRandomCoords(tl_village, br_village)
				human.SetPos(types.Position{X: x, Y: y})
				break
			} else {
				counter++
			}
		}

		if counter == len(objs)+len(humans)-nb_grass {
			end = true
		}
	}
	humans = append(humans, human)

	return humans
}

func PlaceTitan(titan env.AgentI, titans []env.AgentI, W int, H int, dir int) []env.AgentI {
	tl_screen := types.Position{X: 0, Y: 0}
	br_screen := types.Position{X: W, Y: H}
	counter := 0
	end := false
	for !end {
		for _, ti := range titans {
			if utils.DetectCollision(titan.Object(), ti.Object()) {
				x, y := utils.GetRandomCoords(tl_screen, br_screen)
				if dir == 0 {
					y = y - H
				} else if dir == 1 {
					x = x - W
				} else {
					x = x + W
				}
				titan.SetPos(types.Position{X: x, Y: y})
				break
			} else {
				counter++
			}
		}

		if counter == len(titans) {
			end = true
		}
	}

	titans = append(titans, titan)
	return titans
}

func CreateHuman(agt_nb int, humans []env.AgentI, objs []obj.Object, tl_village types.Position, br_village types.Position, objectType types.ObjectName, life int) []env.AgentI {
	var human env.AgentI
	var w, h int

	switch objectType {
	case types.MaleCivilian:
		w, h = params.WMaleCivilian, params.HMaleCivilian
	case types.FemaleCivilian:
		w, h = params.WFemaleCivilian, params.HFemaleCivilian
	case types.MaleSoldier:
		w, h = params.WSoldierM, params.HSoldierM
	case types.FemaleSoldier:
		w, h = params.WSoldierF, params.HSoldierF
	case types.Eren:
		w, h = params.WEren, params.HEren
	case types.Mikasa:
		w, h = params.WMikasa, params.HMikasa
	}

	x, y := utils.GetRandomCoords(tl_village, types.Position{X: br_village.X - w, Y: br_village.Y - h})
	agtId := utils.CreateAgentID(agt_nb, objectType)
	//, reach int, strength int, speed int, vision int,
	if objectType == types.FemaleCivilian || objectType == types.MaleCivilian {
		human = hagt.NewCivilian(agtId, types.Position{X: x, Y: y}, params.CIVILIAN_LIFE, params.CIVILIAN_REACH, params.CIVILIAN_STRENGTH, params.CIVILIAN_SPEED, params.CIVILIAN_VISION, objectType)
	} else if objectType == types.FemaleSoldier || objectType == types.MaleSoldier {
		human = hagt.NewSoldier(agtId, types.Position{X: x, Y: y}, params.SOLDIER_LIFE, params.SOLDIER_REACH, params.SOLDIER_STRENGTH, params.SOLDIER_SPEED, params.SOLDIER_VISION, objectType)
	} else if objectType == types.Eren {
		human = hagt.NewEren(agtId, types.Position{X: x, Y: y}, params.EREN_LIFE, params.EREN_REACH, params.EREN_STRENGTH, params.EREN_SPEED, params.EREN_VISION, objectType)
	} else if objectType == types.Mikasa {
		human = hagt.NewMikasa(agtId, types.Position{X: x, Y: y}, params.MIKASA_LIFE, params.MIKASA_REACH, params.MIKASA_STRENGTH, params.MIKASA_SPEED, params.MIKASA_VISION, objectType)
	}
	humans = PlaceHuman(objs, humans, human, tl_village, types.Position{X: br_village.X - w, Y: br_village.Y - h})
	return humans
}

func CreateHumans(objs []obj.Object, tl_village types.Position, br_village types.Position) []env.AgentI {
	humans := make([]env.AgentI, 0)

	for i := 0; i < params.NB_HUMANS; i++ {
		if i < params.NB_CIVILIANS {
			if i < params.NB_CIVILIANS/2 {
				humans = CreateHuman(i, humans, objs, tl_village, br_village, types.MaleCivilian, params.CIVILIAN_LIFE)
			} else {
				humans = CreateHuman(i, humans, objs, tl_village, br_village, types.FemaleCivilian, params.CIVILIAN_LIFE)
			}
		} else if i < params.NB_CIVILIANS+params.NB_SOLDIERS {
			if i < params.NB_CIVILIANS+params.NB_SOLDIERS/2 {
				humans = CreateHuman(i, humans, objs, tl_village, br_village, types.MaleSoldier, params.SOLDIER_LIFE)
			} else {
				humans = CreateHuman(i, humans, objs, tl_village, br_village, types.FemaleSoldier, params.SOLDIER_LIFE)
			}
			//} else if i < params.NB_CIVILIANS+params.NB_SOLDIERS+1 {
			//	humans = CreateHuman(i, humans, objs, tl_village, br_village, types.Eren, params.EREN_LIFE)
			//} else {
			//	humans = CreateHuman(i, humans, objs, tl_village, br_village, types.Mikasa, params.MIKASA_LIFE)
		}
	}
	return humans
}

func CreateTitans(H int, W int) []env.AgentI {
	var titan env.AgentI
	tl_screen := types.Position{X: 0, Y: 0}
	br_screen := types.Position{X: W, Y: H}

	dir := 0

	titans := make([]env.AgentI, 0)

	for i := 0; i < params.NB_TITANS; i++ {
		x, y := utils.GetRandomCoords(tl_screen, br_screen)

		if dir == 0 {
			y = y - H - 100
			dir = 1
		} else if dir == 1 {
			x = x - W - 50
			dir = 2
		} else {
			x = x + W + 50
			dir = 0
		}

		if i < params.NB_BASIC_TITANS {
			if i < params.NB_BASIC_TITANS/2 {
				titan = tagt.NewBasicTitan(utils.CreateAgentID(i+params.NB_HUMANS, types.BasicTitan1), types.Position{X: x, Y: y}, params.BASIC_TITAN_LIFE, params.BASIC_TITAN_REACH, params.BASIC_TITAN_STRENGTH, params.BASIC_TITAN_SPEED, params.BASIC_TITAN_VISION, types.BasicTitan1, params.BASIC_TITAN_REGEN)
			} else {
				titan = tagt.NewBasicTitan(utils.CreateAgentID(i+params.NB_HUMANS, types.BasicTitan2), types.Position{X: x, Y: y}, params.BASIC_TITAN_LIFE, params.BASIC_TITAN_REACH, params.BASIC_TITAN_STRENGTH, params.BASIC_TITAN_SPEED, params.BASIC_TITAN_VISION, types.BasicTitan2, params.BASIC_TITAN_REGEN)
			}
		}
		// } else if i < params.NB_BASIC_TITANS+params.NB_SPECIAL_TITANS {
		// 	if i < params.NB_BASIC_TITANS+1 {
		// 		//A modifier quand le construteur d'humain sera pret
		// 		titan = obj.NewObject(types.ColossalTitan, types.Position{X: x, Y: y}, params.COLOSSAL_TITAN_LIFE)
		// 	} else if i < params.NB_BASIC_TITANS+2 {
		// 		//A modifier quand le construteur d'humain sera pret
		// 		titan = obj.NewObject(types.BeastTitan, types.Position{X: x, Y: y}, params.BEAST_TITAN_LIFE)
		// 	} else if i < params.NB_BASIC_TITANS+3 {
		// 		titan = obj.NewObject(types.FemaleTitan, types.Position{X: x, Y: y}, params.FEMALE_TITAN_LIFE)
		// 	} else if i < params.NB_BASIC_TITANS+4 {
		// 		titan = obj.NewObject(types.JawTitan, types.Position{X: x, Y: y}, params.JAW_TITAN_LIFE)
		// 	} else if i < params.NB_BASIC_TITANS+5 {
		// 		titan = obj.NewObject(types.ArmoredTitan, types.Position{X: x, Y: y}, params.ARMORED_TITAN_LIFE)
		// 	}
		// }

		//Place l'humain dans des coordonnees aleatoires valides (i.e sans collisions) dans le village
		titans = PlaceTitan(titan, titans, H, W, dir)
	}
	return titans
}

func CreateAgents(H int, W int, objects []obj.Object) []env.AgentI {
	humans := CreateHumans(objects, types.Position{X: int(0.2*float32(W)) + params.CWall, Y: int(0.2*float32(H)) + params.CWall}, types.Position{X: int(0.8 * float32(W)), Y: H})
	titans := CreateTitans(H, W)
	all_agents := make([]env.AgentI, 0)
	all_agents = append(all_agents, humans...)
	all_agents = append(all_agents, titans...)
	return all_agents
}

func GetAvoidancePositions(agentAttributes *env.Agent) []types.Position {
	toAvoid := []types.Position{}

	for _, object := range agentAttributes.PerceivedObjects() {
		hitboxStart, hitboxEnd := object.TL(), object.Hitbox()[1]
		for _, pos := range utils.GetPositionsInHitbox(hitboxStart, hitboxEnd) {
			toAvoid = append(toAvoid, pos)
		}
	}

	for _, agt := range agentAttributes.PerceivedAgents() {
		hitboxStart, hitboxEnd := agt.Agent().ObjectP().TL(), agt.Agent().ObjectP().Hitbox()[1]
		for _, pos := range utils.GetPositionsInHitbox(hitboxStart, hitboxEnd) {
			toAvoid = append(toAvoid, pos)
		}
	}

	return toAvoid
}
