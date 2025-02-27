package constant

type Role string

const (
	ROLE_ADMIN           Role = Role("admin")
	ROLE_FIELD_OFFICER   Role = Role("field_officer")
	ROLE_FIELD_VALIDATOR Role = Role("field_validator")
	ROLE_FIELD_BORROWER  Role = Role("field_borrower")
	ROLE_FIELD_INVESTOR  Role = Role("investor")
)

// RoleSet for fast validation lookup
var validRoles = map[Role]bool{
	ROLE_ADMIN:           true,
	ROLE_FIELD_OFFICER:   true,
	ROLE_FIELD_VALIDATOR: true,
	ROLE_FIELD_BORROWER:  true,
	ROLE_FIELD_INVESTOR:  true,
}

// IsValid checks if a given role is valid
func (r Role) IsValid() bool {
	_, exists := validRoles[r]
	return exists
}

func (r Role) String() string {
	return string(r)
}
