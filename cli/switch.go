package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type idsFlag []string

func (list idsFlag) String() string {
	return strings.Join(list, ",")
}

func (list *idsFlag) Set(value string) error {
	*list = append(*list, value)
	return nil
}

type HttpClient interface {
	Create(title, message string, duration time.Duration) ([]byte, error)
	Edit(id, title, message string, duration time.Duration) ([]byte, error)
	Fetch(ids []string) ([]byte, error)
	Delete(ids []string) error
	Healthy(host string) bool
}

type Switch struct {
	client HttpClient
	api    string
	cmds   map[string]func() func(string) error
}

func CreateSwitch(uri string) Switch {
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

func (s Switch) Switch() error {
	name := os.Args[1]
	cmd, ok := s.cmds[name]
	if !ok {
		return fmt.Errorf("command '%s' not found\n|", name)
	}
	return cmd()(name)
}

/*
* Prints out list of commands and their usages
 */
func (s Switch) Help() {
	var help string
	for name := range s.cmds {
		help += name + "\t--help\n"
	}
	fmt.Printf("Usage: %s <command> [<args>]\n%s", os.Args[0], help)
}

func (s Switch) create() func(string) error {
	return func(cmd string) error {
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, msg, dur := s.reminderFlags(createCmd)

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(createCmd); err != nil {
			return err
		}

		res, err := s.client.Create(*t, *msg, *dur)

		if err != nil {
			return wrapError("Could not create task! Error:", err)
		}

		fmt.Printf("New Task Created!\n%s\n", string(res))
		return nil
	}
}

func (s Switch) edit() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		editCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		editCmd.Var(&ids, "id", "Tasks to edit")

		t, msg, dur := s.reminderFlags(editCmd)

		if err := s.checkArgs(2); err != nil {
			return err
		}

		if err := s.parseCmd(editCmd); err != nil {
			return err
		}

		lastID := ids[len(ids)-1]

		res, err := s.client.Edit(lastID, *t, *msg, *dur)

		if err != nil {
			return wrapError("Could not edit task! Error:", err)
		}

		fmt.Printf("Task edited succesfully!\n%s\n", string(res))
		return nil
	}
}

func (s Switch) fetch() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		fetchCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		fetchCmd.Var(&ids, "id", "Tasks to fetch")

		if err := s.checkArgs(1); err != nil {
			return err
		}

		if err := s.parseCmd(fetchCmd); err != nil {
			return err
		}

		res, err := s.client.Fetch(ids)

		if err != nil {
			return wrapError("Could not fetch task! Error:", err)
		}

		fmt.Printf("Task fetched succesfully!\n%s\n", string(res))
		return nil
	}
}

func (s Switch) delete() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		deleteCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		deleteCmd.Var(&ids, "id", "Ids of tasks to delete")

		if err := s.checkArgs(1); err != nil {
			return err
		}

		err := s.client.Delete(ids)

		if err != nil {
			return wrapError("Could not delete task! Error:", err)
		}

		fmt.Printf("Task deleted succesfully!\n%v\n", ids)
		return nil
	}
}

func (s Switch) health() func(string) error {
	return func(cmd string) error {
		var host string
		healthCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		healthCmd.StringVar(&host, "host", s.api, "Host to check health of")

		if err := s.parseCmd(healthCmd); err != nil {
			return err
		}

		if !s.client.Healthy(host) {
			fmt.Printf("Host %s is not healthy!\n", host)
		} else {
			fmt.Printf("Host %s is healthy!\n", host)
		}

		return nil
	}
}

/*
* Reminder flags (Helper Function)
 */
func (s Switch) reminderFlags(f *flag.FlagSet) (*string, *string, *time.Duration) {
	t, msg, dur := "", "", time.Duration(0)
	f.StringVar(&t, "title", "", "Task Title")
	f.StringVar(&t, "t", "", "Task Title")
	f.StringVar(&msg, "message", "", "Task Message")
	f.StringVar(&msg, "m", "", "Task Message")
	f.DurationVar(&dur, "duration", 0, "Task Duration")
	f.DurationVar(&dur, "d", 0, "Task Duration")
	return &t, &msg, &dur
}

func (s Switch) parseCmd(cmd *flag.FlagSet) error {
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return wrapError("error parsing command"+cmd.Name(), err)
	}
	return nil
}

func (s Switch) checkArgs(minArgs int) error {
	if len(os.Args) == 3 && os.Args[2] == "--help" {
		return nil
	}

	if len(os.Args)-2 < minArgs {
		fmt.Printf("incorrect use of %s\n%s %s --help\n", os.Args[1], os.Args[0], os.Args[1])
		return fmt.Errorf("%s expects at least %d arguments", os.Args[1], minArgs)
	}

	return nil
}
