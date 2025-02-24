package pokeapi

type Stat struct {
	BaseStat int `json:"base_stat"`
	StatInfo struct {
		Name string `json:"name"`
	} `json:"stat"`
}
type TypeDetail struct {
	Name string `json:"name"`
}

type TypeSlot struct {
	Type TypeDetail `json:"type"`
}

type Pokemon struct {
	Name           string     `json:"name"`
	BaseExperience int        `json:"base_experience"`
	Height         int        `json:"height"`
	Weight         int        `json:"weight"`
	Stats          []Stat     `json:"stats"`
	Types          []TypeSlot `json:"types"`
}
