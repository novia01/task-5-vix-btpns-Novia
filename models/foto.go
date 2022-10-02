package models

type(
	//Foto
	Foto struct {
		ID        	int       	`json:"id"`
		Title      	string       	`json:"title"`
		Caption    	string       	`json:"caption"`
		URL        	string       	`json:"url"`
		User_ID     int		    `json:"userid"`
	}
)