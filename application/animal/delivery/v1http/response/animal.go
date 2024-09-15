package response

import "time"

type (
	AnimalReadByIdResponse struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Class string `json:"class"`
		Legs  int    `json:"legs"`
	}
	
	AnimalListResponse struct {
		Id        int       `json:"id"`
		Name      string    `json:"name"`
		Class     string    `json:"class"`
		Legs      int       `json:"legs"`
		CreatedAt time.Time `json:"created_at"`
	}
)
