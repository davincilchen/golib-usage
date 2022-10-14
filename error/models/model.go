package models

import "fmt"

var balance = 12

type ModelIOResult int

type ModelError struct {
	Code ModelIOResult
	Desc string
}

func (t *ModelError) Error() string {

	return fmt.Sprintf("ModelIOResult Code[%d], %s", t.Code, t.Desc)
}

const (
	RetUndefined        ModelIOResult = 1
	RetRecordNotFound   ModelIOResult = 2
	RetBalanceNotEnough ModelIOResult = 3
	RetUpdateFailed     ModelIOResult = 4
	RetAPIFailed        ModelIOResult = 5
	RetDuplicationTxID  ModelIOResult = 10
)

func Deposite(val int) error {

	if val == 0 {
		return &ModelError{
			Code: RetUndefined,
			Desc: "Unknow Error",
		}
	}

	balance += val
	return nil
}

func Withdraw(val int) error {

	if val > balance {
		return &ModelError{
			Code: RetBalanceNotEnough,
			Desc: "Balance is not enough.",
		}
	}

	balance -= val
	return nil
}

func BalanceNotEnough(err error) bool {
	ele, ok := err.(*ModelError)
	if ok {
		if ele.Code == RetBalanceNotEnough {
			return true
		}

	}
	return false
}
