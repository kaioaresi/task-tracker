package cli

import "flag"

func Create() string {
	var description string

	flag.StringVar(&description, "description", "", "-create 'task description'")

	flag.Parse()

	return description
}
