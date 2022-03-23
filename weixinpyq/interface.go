package main

import "learn.go/pkg/apis"

type ServerInterface interface {
	RegisterPersonInformation(pi *apis.PersonInformation) error
	UpdatePersonInformation(pi *apis.PersonInformation) (*apis.PersonInformationFatRate, error)
	GetFateRate(name string) (*apis.PersonRank, error)
	GetTop() ([]*apis.PersonRank, error)
	DeletePersonPyq(id int) error
	ShowPersonPyq() ([]*apis.PersonInformation, error)
}
type RankInitInterface interface {
	Init() error
}
