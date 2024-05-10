package services

import (
	"errors"
	"regexp"
	"trademarkia/repository"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func UserSignup(input repository.UserData) error {
	if !emailRegex.MatchString(input.Email) {
		return errors.New("Invalid email format")
	}
	_, err := repository.FindUserByMail(input.Email)
	if err == nil {
		return errors.New("Given email already exsist")
	}

	_, err = repository.AddUserData(input)
	if err != nil {
		return err
	}

	return nil

}

func UserLogin(email string) (repository.UserData, error) {
	if !emailRegex.MatchString(email) {
		return repository.UserData{}, errors.New("Invalid email format")
	}
	userData, err := repository.FindUserByMail(email)
	if err != nil {
		return repository.UserData{}, errors.New("Given email not exsist")
	}

	return userData, nil
}

func GetUserData(email string) (repository.UserData, error) {
	resposne, err := repository.FindUserByMail(email)
	if err != nil {
		return repository.UserData{}, errors.New("Give email not exsist")
	}
	return resposne, nil
}

func EditUserData(input repository.UserData) error {
	_, err := repository.FindUserByMail(input.Email)
	if err != nil {
		return errors.New("Given email not exsist")
	}
	if err = repository.EditUserData(input); err != nil {
		return err
	}

	return nil
}

func DeleteUserData(email string) error {
	_, err := repository.FindUserByMail(email)
	if err != nil {
		return errors.New("Given email not exsist")
	}
	if err = repository.DeleteUserData(email); err != nil {
		return err
	}

	return nil
}
