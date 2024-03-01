package main

import "errors"

var (
	ErrNotOpenedYet  = errors.New("o portão ainda não abriu")
	ErrClosedAlready = errors.New("o portão já fechou")
)
