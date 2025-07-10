package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	
}

func NewUserDomain(
	email, password, name string, age int8,
) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (userDomain *userDomain) GetEmail() {
	return userDomain.email
}

func (userDomain *userDomain) GetPassword() {
	return userDomain.password
}

func (userDomain *userDomain) GetName() {
	return userDomain.name
}

func (userDomain *userDomain) GetAge() {
	return userDomain.age
}

func (userDomain *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(userDomain.password))
	userDomain.password = hex.EncodeToString(hash.Sum(nil))
}
