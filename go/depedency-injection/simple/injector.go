//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func IntializedService(IsErrors bool) (*SimpleService, error) {

	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)

	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseSQL,
		NewDatabaseMysql,
		NewDatabaseRepository,
	)
	return nil
}

var fooset = wire.NewSet(NewFooRepository, NewFooService)
var barset = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {

	wire.Build(
		fooset, barset, NewFooBarService,
	)
	return nil

}

//wrong binding interface
// func InitializedHello() *HelloService {
// 	wire.Build(
// 		NewSayHelloService, NewSayHelloImpl,
// 	)
// 	return nil
// }

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(
		helloSet, NewHelloService,
	)
	return nil
}

func InitializedFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)

	return nil, nil
}
