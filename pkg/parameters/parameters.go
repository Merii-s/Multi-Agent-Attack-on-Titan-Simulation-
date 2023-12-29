package params

const (
	//SIMULATION PARAMETERS
	NB_BASIC_TITANS   = 10
	NB_SPECIAL_TITANS = 0
	NB_TITANS         = NB_BASIC_TITANS + NB_SPECIAL_TITANS

	NB_VILLAGERS = 50
	NB_SOLDIERS  = 50
	NB_HUMANS    = NB_VILLAGERS + NB_SOLDIERS + 2

	NB_AGENTS = NB_TITANS + NB_HUMANS

	EREN_LIFE           = 100
	MIKASA_LIFE         = 100
	VILLAGER_LIFE       = 100
	SOLDIER_LIFE        = 200
	BASIC_TITAN_LIFE    = 100
	COLOSSAL_TITAN_LIFE = 100
	ARMORED_TITAN_LIFE  = 100
	FEMALE_TITAN_LIFE   = 100
	JAW_TITAN_LIFE      = 100
	BEAST_TITAN_LIFE    = 100

	GRASS_LIFE       = 100000000000000
	WALL_LIFE        = 300
	BIG_HOUSE_LIFE   = 150
	SMALL_HOUSE_LIFE = 100
	DUNGEON_LIFE     = 500
	FIELD_LIFE       = 200
	FIELD_RESERVE    = 100

	//GUI PARAMETERS (SPRITE DIMENSIONS)
	//Screen Dimensions
	ScreenHeight = 700
	ScreenWidth  = 1000

	//Sprite Dimensions
	CWall = 20

	WField = 40
	HField = 34

	CGrass = 50

	WBHouse1 = 55
	HBHouse1 = 46

	WBHouse2 = 42
	HBHouse2 = 55

	WSHouse = 43
	HSHouse = 40

	WDungeon = 29
	HDungeon = 52

	WEren = 14
	HEren = 33

	WMikasa = 15
	HMikasa = 33

	WMaleVillager = 7
	HMaleVillager = 18

	WFemaleVillager = 7
	HFemaleVillager = 17

	WBasicTitanF = 21
	HBasicTitanF = 40

	WBasicTitanM = 22
	HBasicTitanM = 40

	WArmoredTitan = 20
	HArmoredTitan = 49

	WArmoredTitanHuman = 12
	HArmoredTitanHuman = 32

	WBeastTitan = 31
	HBeastTitan = 64

	HBeastTitanHuman = 33
	WBeastTitanHuman = 15

	WColossalTitan = 28
	HColossalTitan = 65

	WColossalTitanHuman = 15
	HColossalTitanHuman = 33

	WErenTitan = 20
	HErenTitan = 50

	WFemaleTitan = 19
	HFemaleTitan = 50

	WFemaleTitanHuman = 12
	HFemaleTitanHuman = 33

	WJawTitan = 32
	HJawTitan = 34

	WJawTitanHuman = 16
	HJawTitanHuman = 33

	WSoldierM = 17
	HSoldierM = 25

	WSoldierF = 20
	HSoldierF = 22
)
