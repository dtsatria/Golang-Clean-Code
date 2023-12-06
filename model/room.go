package model

import "time"

type Room struct {
	Id          string       `json:"id"`
	RoomType    string       `json:"roomType"`
	MaxCapacity int          `json:"maxcapacity"`
	Facility    RoomFacility `json:"facility"`
	Status      string       `json:"status"` //untuk status hanya ada dua yaitu Available atau Booked
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type RoomFacility struct {
	Id               string    `json:"id"`
	RoomDescription  string    `json:"description"`
	Fwifi            string    `json:"wifi"`
	FsoundSystem     string    `json:"soundSystem"`
	Fprojector       string    `json:"projector"`
	FscreenProjector string    `json:"screenProjector"`
	Fchairs          string    `json:"chairs"`
	Ftables          string    `json:"tables"`
	FsoundProof      string    `json:"soundProof"`
	FsmonkingArea    string    `json:"smokingArea"`
	Ftelevison       string    `json:"television"`
	FAc              string    `json:"ac"`
	Fbathroom        string    `json:"bathroom"`
	FcoffeMaker      string    `json:"coffeMaker"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
