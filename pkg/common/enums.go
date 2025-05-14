package common

// AccountType defines the type of bank account.
type AccountType string

const (
	Savings AccountType = "SAVINGS"
	Current AccountType = "CURRENT"
)

// IsValid checks if the AccountType is valid.
func (at AccountType) IsValid() bool {
	switch at {
	case Savings, Current:
		return true
	}
	return false
}

// EqualityType defines the type of comparison for balance queries.
type EqualityType string

const (
	GreaterThan EqualityType = "GREATER_THAN"
	LessThan    EqualityType = "LESS_THAN"
)

// IsValid checks if the EqualityType is valid.
func (et EqualityType) IsValid() bool {
	switch et {
	case GreaterThan, LessThan:
		return true
	}
	return false
}
