package user

// UserID uniquely identifies a particular user.
type UserID string

type (
	User struct {
		ID      UserID
		Role    Role
		Address string
	}
)

type Role int

const (
	UserRole     = iota
	OperatorRole
	AdminRole
)

func (r Role) String() string {
	switch r {
	case UserRole:
		return "User"
	case OperatorRole:
		return "Operator"
	case AdminRole:
		return "Admin"
	}

	return ""
}
