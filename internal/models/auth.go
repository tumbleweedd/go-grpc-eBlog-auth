package models

type Role string

const (
	ADMIN              Role = "ADMIN"
	USER                    = "USER"
	SUPPORT                 = "SUPPORT"
	IsAccountNonLocked bool = true
	IsAccountLocked         = false
)

type User struct {
	Id                 int    `json:"-" db:"user_id"`
	Name               string `json:"name" db:"name"`
	Lastname           string `json:"lastname" db:"lastname"`
	Username           string `json:"username" db:"username"`
	Email              string `json:"email" db:"email"`
	Password           string `json:"password" db:"password"`
	Role               Role   `json:"-" db:"role"`
	IsAccountNonLocked bool   `json:"-" db:"is_account_non_locked"`
}
