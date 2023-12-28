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

// A modifier quand les constructeurs seront pret
func PlaceHuman(objs []obj.Object, humans []env.AgentI, human env.AgentI, tl_village types.Position, br_village types.Position) []env.AgentI {
	end := false

	for !end {
		counter := 0
		nb_grass := 0
		for _, obj := range objs {
			if obj.Name() != types.Grass {
				if utils.DetectCollision2(human.Object(), obj) {
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
			if utils.DetectCollision2(human.Object(), hu.Object()) {
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
			if utils.DetectCollision2(titan.Object(), ti.Object()) {
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

func CreateHuman(humans []env.AgentI, objs []obj.Object, tl_village types.Position, br_village types.Position, objectType types.ObjectName, life int) []env.AgentI {
	var human env.AgentI
	var w, h int

	switch objectType {
	case types.MaleVillager:
		w, h = params.WMaleVillager, params.HMaleVillager
	case types.FemaleVillager:
		w, h = params.WFemaleVillager, params.HFemaleVillager
	case types.MaleSoldier:
		w, h = params.WSoldierM, params.HSoldierM
	case types.FemaleSoldier:
		w, h = params.WSoldierF, params.HSoldierF
	case types.Eren:
		w, h = params.WEren, params.HEren
	case types.Mikasa:
		w, h = params.WMikasa, params.HMikasa
	}

	//x, y := utils.GetRandomCoords(tl_village, types.Position{X: br_village.X - w, Y: br_village.Y - h})

	//Condition impossible pour eviter d'entrer dans la boucle
	if objectType == types.FemaleVillager || objectType == types.MaleSoldier {
		human = hagt.NewCivilian("", types.Position{X: x, Y: y}, params.VILLAGER_LIFE, 0, 0, 0, 0, objectType)
	} else if objectType == types.FemaleSoldier || objectType == types.MaleSoldier {
		human = hagt.NewSoldier("", types.Position{X: x, Y: y}, params.SOLDIER_LIFE, 0, 0, 0, 0, objectType)
	} else if objectType == types.Eren {
		human = hagt.NewEren("", types.Position{X: x, Y: y}, params.EREN_LIFE, 0, 0, 0, 0, objectType)
	} else if objectType == types.Mikasa {
		human = hagt.NewMikasa("", types.Position{X: x, Y: y}, params.MIKASA_LIFE, 0, 0, 0, 0, objectType)
	}

	humans = PlaceHuman(objs, humans, human, tl_village, types.Position{X: br_village.X - w, Y: br_village.Y - h})
	return humans
}

func CreateHumans(objs []obj.Object, tl_village types.Position, br_village types.Position) []env.AgentI {
	humans := make([]env.AgentI, 0)

	for i := 0; i < params.NB_HUMANS; i++ {
		if i < params.NB_VILLAGERS {
			if i < params.NB_VILLAGERS/2 {
				humans = CreateHuman(humans, objs, tl_village, br_village, types.MaleVillager, params.VILLAGER_LIFE)
				//titan = tagt.NewBasicTitan("", types.Position{X: x, Y: y}, params.BASIC_TITAN_LIFE, 0, 0, 0, 0, types.BasicTitan1, 0)
			} else {
				humans = CreateHuman(humans, objs, tl_village, br_village, types.FemaleVillager, params.VILLAGER_LIFE)
			}
		} else if i < params.NB_VILLAGERS+params.NB_SOLDIERS {
			if i < params.NB_VILLAGERS+params.NB_SOLDIERS/2 {
				humans = CreateHuman(humans, objs, tl_village, br_village, types.MaleSoldier, params.SOLDIER_LIFE)
			} else {
				humans = CreateHuman(humans, objs, tl_village, br_village, types.FemaleSoldier, params.SOLDIER_LIFE)
			}
		} else if i < params.NB_VILLAGERS+params.NB_SOLDIERS+1 {
			humans = CreateHuman(humans, objs, tl_village, br_village, types.Eren, params.EREN_LIFE)
		} else {
			humans = CreateHuman(humans, objs, tl_village, br_village, types.Mikasa, params.MIKASA_LIFE)
		}

	}
	return humans
}

func CreateTitans(H int, W int) []env.AgentI {
	var titan env.AgentI
	tl_screen := types.Position{X: 0, Y: 0}
	br_screen := types.Position{X: W, Y: H}

	dir := 0

	//A modifier quand le construteur de titan sera pret
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
				//A modifier quand le construteur d'humain sera pret
				titan = tagt.NewBasicTitan("", types.Position{X: x, Y: y}, params.BASIC_TITAN_LIFE, 0, 0, 0, 0, types.BasicTitan1, 0)
			} else {
				//A modifier quand le construteur d'humain sera pret
				titan = tagt.NewBasicTitan("", types.Position{X: x, Y: y}, params.BASIC_TITAN_LIFE, 0, 0, 0, 0, types.BasicTitan2, 0)
			}

		} /*else if i < params.NB_BASIC_TITANS+params.NB_SPECIAL_TITANS {
			if i < params.NB_BASIC_TITANS+1 {
				//A modifier quand le construteur d'humain sera pret
				titan = obj.NewObject(types.ColossalTitan, types.Position{X: x, Y: y}, params.COLOSSAL_TITAN_LIFE)
			} else if i < params.NB_BASIC_TITANS+2 {
				//A modifier quand le construteur d'humain sera pret
				titan = obj.NewObject(types.BeastTitan, types.Position{X: x, Y: y}, params.BEAST_TITAN_LIFE)
			} else if i < params.NB_BASIC_TITANS+3 {
				titan = obj.NewObject(types.FemaleTitan, types.Position{X: x, Y: y}, params.FEMALE_TITAN_LIFE)
			} else if i < params.NB_BASIC_TITANS+4 {
				titan = obj.NewObject(types.JawTitan, types.Position{X: x, Y: y}, params.JAW_TITAN_LIFE)
			} else if i < params.NB_BASIC_TITANS+5 {
				titan = obj.NewObject(types.ArmoredTitan, types.Position{X: x, Y: y}, params.ARMORED_TITAN_LIFE)
			}
		}*/

		//Place l'humain dans des coordonnees aleatoires valides (i.e sans collisions) dans le village
		titans = PlaceTitan(titan, titans, H, W, dir)
	}
	return titans

}
