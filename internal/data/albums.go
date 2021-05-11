package data

import "time"

type Album struct {
	ID  		int64		`json:"id"`
	CreatedAt   time.Time	`json:"-"`
	Title 		string		`json:"title"`
	Year		int32		`json:"year,omitempty"`
	Runtime 	Runtime		`json:"runtime,omitempty"`
	Artist		[]string	`json:"artist,omitempty"`
	Genres		[]string 	`json:"genres,omitempty"`
	Version 	int32		`json:"version"`
}


