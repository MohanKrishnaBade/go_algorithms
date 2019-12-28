package DbCommands

import (
	"github.com/spf13/cobra"
	"../../../go_algorithms/Connections"
)

func Connection() *cobra.Command {
	return &cobra.Command{
		Use: "DbConnection",
		RunE: func(cmd *cobra.Command, args []string) error {

			var databases Connections.Databases;
			db, err := databases.EstablishConnection(Connections.Database{});

			// if there is an error opening the connection, handle it
			if err != nil {
				panic(err.Error())
			}
			// defer the close till after the main function has finished
			// executing
			defer db.Close()
			return nil
		},
	}
}
