package main

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	Map map[string]func(*state, command) error
}
