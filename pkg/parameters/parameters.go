package params

const (

	//SIMULATION PARAMETERS

	MaxStep     = 10000000000000000
	MaxDuration = 10000000000000000

	//Agent numbers
	NB_BASIC_TITANS   = 3
	NB_SPECIAL_TITANS = 0
	NB_TITANS         = NB_BASIC_TITANS + NB_SPECIAL_TITANS

	NB_CIVILIANS = 0
	NB_SOLDIERS  = 0
	NB_HUMANS    = NB_CIVILIANS + NB_SOLDIERS //+ 2

	NB_AGENTS = NB_TITANS + NB_HUMANS

	//Agent Lives
	EREN_LIFE           = 100
	MIKASA_LIFE         = 100
	CIVILIAN_LIFE       = 150
	SOLDIER_LIFE        = 200
	BASIC_TITAN_LIFE    = 100
	COLOSSAL_TITAN_LIFE = 100
	ARMORED_TITAN_LIFE  = 100
	FEMALE_TITAN_LIFE   = 100
	JAW_TITAN_LIFE      = 100
	BEAST_TITAN_LIFE    = 100

	//Object Lives
	GRASS_LIFE       = 100000000000000
	WALL_LIFE        = 300
	BIG_HOUSE_LIFE   = 150
	SMALL_HOUSE_LIFE = 100
	DUNGEON_LIFE     = 500
	FIELD_LIFE       = 200
	FIELD_RESERVE    = 100

	//Agent stats
	EREN_SPEED           = 2
	MIKASA_SPEED         = 3
	CIVILIAN_SPEED       = 1
	SOLDIER_SPEED        = 2
	BASIC_TITAN_SPEED    = 30
	COLOSSAL_TITAN_SPEED = 1
	ARMORED_TITAN_SPEED  = 1
	BEAST_TITAN_SPEED    = 1
	FEMALE_TITAN_SPEED   = 1
	JAW_TITAN_SPEED      = 1

	EREN_STRENGTH           = 10
	MIKASA_STRENGTH         = 10
	CIVILIAN_STRENGTH       = 10
	SOLDIER_STRENGTH        = 10
	BASIC_TITAN_STRENGTH    = 10
	COLOSSAL_TITAN_STRENGTH = 10
	ARMORED_TITAN_STRENGTH  = 10
	BEAST_TITAN_STRENGTH    = 10
	FEMALE_TITAN_STRENGTH   = 10
	JAW_TITAN_STRENGTH      = 10

	EREN_REACH           = 1
	MIKASA_REACH         = 1
	CIVILIAN_REACH       = 1
	SOLDIER_REACH        = 1
	BASIC_TITAN_REACH    = 1
	COLOSSAL_TITAN_REACH = 1
	ARMORED_TITAN_REACH  = 1
	BEAST_TITAN_REACH    = 1
	FEMALE_TITAN_REACH   = 1
	JAW_TITAN_REACH      = 1

	EREN_VISION           = 10
	MIKASA_VISION         = 10
	CIVILIAN_VISION       = 10
	SOLDIER_VISION        = 10
	BASIC_TITAN_VISION    = 100
	BEAST_TITAN_VISION    = 10
	COLOSSAL_TITAN_VISION = 10
	ARMORED_TITAN_VISION  = 10
	FEMALE_TITAN_VISION   = 10
	JAW_TITAN_VISION      = 10

	BASIC_TITAN_REGEN    = 1
	BEAST_TITAN_REGEN    = 1
	COLOSSAL_TITAN_REGEN = 1
	ARMORED_TITAN_REGEN  = 1
	FEMALE_TITAN_REGEN   = 1
	JAW_TITAN_REGEN      = 1

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

	WMaleCivilian = 7
	HMaleCivilian = 18

	WFemaleCivilian = 7
	HFemaleCivilian = 17

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
