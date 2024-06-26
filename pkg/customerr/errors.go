package cerr

import (
	"fmt"
)

// UserType определяет тип пользователя
type UserType string

const (
	User  UserType = "User"
	Admin UserType = "Admin"

	Bread UserType = "BRD"
	Order UserType = "ORD"
)

// LayerType определяет тип слоя
type LayerType string

const (
	Repository LayerType = "REPO"
	Service    LayerType = "SERV"
	Handle     LayerType = "HAND"
)

// ErrorType определяет тип ошибки
type ErrorType string

const (
	Transaction  ErrorType = "transaction error"
	Rollback     ErrorType = "rollback error"
	Commit       ErrorType = "commit error"
	Scan         ErrorType = "scan error"
	Execut       ErrorType = "execution error"
	ExecCon      ErrorType = "transaction.ExecContext error"
	Rows         ErrorType = "rows error"
	NoOneRow     ErrorType = "row count doesnt equals 1"
	InvalidPhone ErrorType = "invalid phone number error"
	InvalidEmail ErrorType = "invalid email error"
	InvalidPWD   ErrorType = "invalid password error"
	InvalidCount ErrorType = "count more that have"
	DiffPWD      ErrorType = "pwd not equal"
	Hash         ErrorType = "error in hashing time"
	NotFound     ErrorType = "this row not found"
)

// CustomError структура для кастомной ошибки
type CustomError struct {
	Who   UserType
	Layer LayerType
	Type  ErrorType
	Err   error
}

// Error метод для реализации интерфейса error
func (e CustomError) Error() error {
	return fmt.Errorf("Error occurred by: %v, Layer: %v, Type: %v, Error: %v", e.Who, e.Layer, e.Type, e.Err)
}

// NewCustomError функция для создания новой кастомной ошибки
func Err(who UserType, layer LayerType, errType ErrorType, err error) CustomError {
	return CustomError{Who: who, Layer: layer, Type: errType, Err: err}
}
