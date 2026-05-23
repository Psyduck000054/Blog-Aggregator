package functions

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Map map[string]func(*State, Command) error
}
