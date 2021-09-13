package domain

type (
	User struct {
		ID    string
		Name  string
		Email  string
		Address string
		Price string
	}
)

var (
	UserTableName = "tbl_user"
	UserTable     = struct {
		ID    string
		Name  string
		Email  string
		Address string
		Price string
	}{
		ID:     	"id",
		Name:   	"name",
		Email: 		"email",
		Address: 	"address",
	}
)

