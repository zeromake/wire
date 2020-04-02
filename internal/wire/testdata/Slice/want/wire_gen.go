// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func NewRouter() *Mux {
	homeController := NewHome()
	uploadController := NewUpload()
	v := []Controller{
		homeController,
		uploadController,
	}
	mux := InitRouter(v)
	return mux
}