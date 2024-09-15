package request

type (
	AnimalCreateRequest struct {
		Name  string `json:"name" binding:"required"`
		Class string `json:"class" binding:"required"`
		Legs  int    `json:"legs" binding:"required"`
	}
	
	AnimalUpdateRequest struct {
		Name  string `json:"name"`
		Class string `json:"class"`
		Legs  int    `json:"legs"`
	}
	
	ListAnimalRequest struct {
		Page    int    `form:"page"`
		PerPage int    `form:"per_page"`
		OrderBy string `form:"order_by"`
		Sort    string `form:"sort"`
		Search  string `form:"search"`
	}
)
