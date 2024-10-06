package model

type StatusData struct {
	Success bool  `json:"success"`
	Website uint8 `json:"website"`
	Api     uint8 `json:"api"`
	Files   uint8 `json:"files"`
}
