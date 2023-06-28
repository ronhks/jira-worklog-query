package domain

type User struct {
	ID       string `json:"id" gorm:"unique;not null"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	Validate(user *User) error
	ValidateAge(user *User) bool
	Create(user *User) (*User, error)
	FindAll() ([]User, error)
}

type UserRepository interface {
	Save(user *User) (*User, error)
	FindAll() ([]User, error)
	Delete(user *User) error
	Migrate() error
}
