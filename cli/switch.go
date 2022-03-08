package cli

import "fmt"

type HttpClient interface {
}

type Switch struct {
	client HttpClient
	api    string
	cmds   map[string]func() func(string) error
}

func createSwitch(uri string) Switch {
	httpClient := NewClient(uri)

	switch1 := Switch{
		client: httpClient,
		api:    uri,
	}

	switch1.cmds = map[string]func() func(string) error{
		"create": switch1.create,
		"edit":   switch1.edit,
		"fetch":  switch1.fetch,
		"delete": switch1.delete,
		"health": switch1.health,
	}

	return switch1
}

func (s Switch) create() func(string) error {
	return func(cmd string) error {
		fmt.Println("create")
		return nil
	}
}

func (s Switch) edit() func(string) error {
	return func(args string) error {
		fmt.Println("switch")
		return nil
	}
}

func (s Switch) fetch() func(string) error {
	return func(args string) error {
		fmt.Println("fetch")
		return nil
	}
}

func (s Switch) delete() func(string) error {
	return func(args string) error {
		fmt.Println("delete")
		return nil
	}
}

func (s Switch) health() func(string) error {
	return func(args string) error {
		fmt.Println("health")
		return nil
	}
}
