package main

import (
	"error/models"
	"fmt"

	"github.com/pkg/errors"
)

func main() {

	err := models.Deposite(0)
	if err != nil {
		ret := models.BalanceNotEnough(err)

		fmt.Printf("Deposite 0, models.BalanceNotEnough %t \n", ret)

	}
	fmt.Println("0.0 ====================== ")

	for i := 0; i < 2; i++ {
		err := TryWithdraw(5)

		if err != nil {
			s := fmt.Sprintf("%+v", err)
			fmt.Println(s)
		}
	}
	//多次warp 會有多次print call stack
	fmt.Println("1.0 ====================== ")

	err = TryWithdraw(5)

	if err != nil {
		s := fmt.Sprintf("%v", err) //不使用+
		fmt.Println(s)
	}
	fmt.Println("2.0 ====================== ")
	for i := 0; i < 3; i++ {
		err := CorrectWithdraw(5) //只有最上層用warp 其他用withMessage

		if err != nil {
			s := fmt.Sprintf("%+v", err) //有使用+
			fmt.Println(s)

			s = fmt.Sprintf("!!printf%%s  %s", err) //
			fmt.Println(s)
		}
	}
	fmt.Println("3.0 ====================== ")
	err = models.Withdraw(5)
	if err != nil {
		ret := models.BalanceNotEnough(err)

		fmt.Printf("Deposite 0, models.BalanceNotEnough %t \n", ret)

	}
	fmt.Println("5.0 ====================== ")
}

func TryWithdraw(val int) error {

	err := tryWithdraw(val)
	if err != nil {
		return errors.Wrap(err, "Failed to TryWithdraw")
	}

	return nil

}
func tryWithdraw(val int) error {

	err := models.Withdraw(val)
	if err != nil {
		return errors.Wrap(err, "(Failed to tryWithdraw)")
	}

	return nil
}

// =========

func CorrectWithdraw(val int) error {

	err := correctWithdraw(val)
	if err != nil {
		return errors.Wrap(err, "Failed to CorrectWithdraw")
	}

	return nil

}
func correctWithdraw(val int) error {

	err := models.Withdraw(val)
	if err != nil {
		//return errors.Wrap(err, "(Failed to correctWithdraw)")
		return errors.WithMessage(err, "(Failed to correctWithdraw)")
	}

	return nil
}
