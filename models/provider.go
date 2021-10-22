package models

type Provider interface {
	Provide() []*Car
}
