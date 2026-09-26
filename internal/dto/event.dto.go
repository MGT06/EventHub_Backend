package dto

import "time"

type Event struct {
	Id            int       `json:"id"`
	OrganizerName string    `json:"organizer"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Image_url     string    `json:"image_url"`
	Start_at      time.Time `json:"start_at"`
	End_at        time.Time `json:"end_at"`
	Format        string    `json:"format"`
	Location      string    `json:"location"`
	Capacity      string    `json:"capacity"`
	Speakers      string    `json:"speakers"`
	Categories    string    `json:"categories"`
}

type JoinEvent struct {
	Id_Event int `json:"id_event"`
}
