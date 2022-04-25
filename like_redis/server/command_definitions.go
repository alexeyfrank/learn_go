package server

import (
	"errors"
	"strconv"
)

type SetCommandDefinition struct {
	Name    string
	RawArgs []string
	Key     string
	Value   int64
}

func NewSetCommandDefinition(args []string) (SetCommandDefinition, error) {
	if len(args) != 2 {
		return SetCommandDefinition{}, errors.New("SET should have format 'SET key value'")
	}

	val, err := strconv.ParseInt(args[1], 0, 64)

	if err != nil {
		return SetCommandDefinition{}, errors.New(err.Error())
	}

	return SetCommandDefinition{
		Name:    "SET",
		RawArgs: args,
		Key:     args[0],
		Value:   val,
	}, nil
}

type GetCommandDefinition struct {
	Name    string
	RawArgs []string
	Key     string
}

func NewGetCommandDefinition(args []string) (GetCommandDefinition, error) {
	if len(args) != 1 {
		return GetCommandDefinition{}, errors.New("GET should have format 'GET key'")
	}

	return GetCommandDefinition{
		Name:    "GET",
		RawArgs: args,
		Key:     args[0],
	}, nil
}

type IncCommandDefinition struct {
	Name    string
	RawArgs []string
	Key     string
	Value   int64
}

func NewIncCommandDefinition(args []string) (IncCommandDefinition, error) {
	if len(args) != 2 {
		return IncCommandDefinition{}, errors.New("INC should have format 'INC key value'")
	}

	val, err := strconv.ParseInt(args[1], 0, 64)

	if err != nil {
		return IncCommandDefinition{}, errors.New(err.Error())
	}

	return IncCommandDefinition{
		Name:    "INC",
		RawArgs: args,
		Key:     args[0],
		Value:   val,
	}, nil
}

type DecCommandDefinition struct {
	Name    string
	RawArgs []string
	Key     string
	Value   int64
}

func NewDecCommandDefinition(args []string) (DecCommandDefinition, error) {
	if len(args) != 2 {
		return DecCommandDefinition{}, errors.New("DEC should have format 'DEC key value'")
	}

	val, err := strconv.ParseInt(args[1], 0, 64)

	if err != nil {
		return DecCommandDefinition{}, errors.New(err.Error())
	}

	return DecCommandDefinition{
		Name:    "DEC",
		RawArgs: args,
		Key:     args[0],
		Value:   val,
	}, nil
}

type NotFoundCommandDefinition struct {
	Name    string
	RawArgs []string
}

func NewNotFoundCommandDefinition(args []string) (NotFoundCommandDefinition, error) {
	return NotFoundCommandDefinition{
		Name:    "DEC",
		RawArgs: args,
	}, nil
}
