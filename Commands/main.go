package main

import (
	"os"

	"github.com/spf13/cobra"
	c "../../go_algorithms/Commands/CurrentTime"
	"../../go_algorithms/Commands/JsonToObject"
	"../../go_algorithms/Connections"
	"../../go_algorithms/Commands/DbCommands"
)

func init() {

	var databases Connections.Databases;
	db, err := databases.EstablishConnection(Connections.Database{});

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}

func main() {
	cmd := &cobra.Command{
		Use:          "gc",
		Short:        "Hello Gophers!",
		SilenceUsage: true,
	}

	cmd.AddCommand(c.PrintTimeCmd())
	cmd.AddCommand(JsonToObject.GetObject())
	cmd.AddCommand(DbCommands.Connection())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
