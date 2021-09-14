package domain

type (
	Product struct {
		ID          string    	`json:"id"`
		Code       	string    	`json:"code"`
		Name	 	string    	`json:"name"`
		Desc   		string      	`json:"desc"`
		Stock 		int    		`json:"stock"`
		Price 		int    		`json:"price"`
	}

	ProductPayload struct {
		Code       	string    	`json:"code" validate:"required,omitempty,max-10"`
		Name	 	string    	`json:"name" validate:"required,omitempty"`
		Desc   		string      	`json:"desc" validate:"required,omitempty"`
		Stock 		int    		`json:"stock" validate:"required,omitempty"`
		Price 		int    		`json:"price" validate:"required,omitempty"`
	}

	FindProductFilter struct {
		Limit  uint64 `query:"limit"`
		Offset uint64 `query:"offset"`
		Sort   string `query:"sort"`
	}

)

var (
	ProductTableName = "tbl_product"
	ProductTable     = struct {
		ID    string
		Code  string
		Name  string
		Desc  string
		Stock string
		Price string
	}{
		ID:     "id",
		Code:   "code",
		Name:   "name",
		Desc: 	"description",
		Stock: 	"stock",
		Price: 	"price",
	}
)
