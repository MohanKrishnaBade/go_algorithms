package JsonToObject

import (
	"github.com/spf13/cobra"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func GetObject() *cobra.Command {
	return &cobra.Command{
		Use: "json-to-object",
		RunE: func(cmd *cobra.Command, args []string) error {

			plan, _ := ioutil.ReadFile(args[0])
			var data interface{}
			err := json.Unmarshal(plan, &data)
			if err != nil {
				panic(err)
			}

			result := data.(interface{}).(map[string]interface{})["sdgs"]
			fmt.Println(result);
			return nil
		},
	}
}
