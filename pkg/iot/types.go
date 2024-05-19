package iot

type IOTInfo struct {
	Status     string      `json:"status"`
	Message    string      `json:"message,omitempty"`
	Rooms      []Room      `json:"rooms"`
	Devices    []Device    `json:"devices"`
	Households []Household `json:"households"`
}

type Room struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	HouseholdId string   `json:"household_id"`
	Devices     []string `json:"devices"`
}

type Household struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Device struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Room       string     `json:"room"`
	Properties []Property `json:"properties"`

	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type Property struct {
	LastUpdated float64 `json:"last_updated"`
	State       *State  `json:"state"`
}

type State struct {
	Instance string  `json:"instance"`
	Value    float64 `json:"value"`
}
