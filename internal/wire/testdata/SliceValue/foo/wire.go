//+build wireinject

package main

import (
	"github.com/google/wire"
)

func NewRouter() *Mux {
	panic(
		wire.Build(
			wire.Slice(
				[]Controller(nil),
				wire.Struct(new(HomeController), "*"),
				wire.Struct(new(DocController), "*"),
				wire.Value(&UploadController{}),
			),
			InitRouter,
		),
	)
}
