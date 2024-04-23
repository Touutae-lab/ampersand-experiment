package ampersand_experiment

import "github.com/amp-labs/connectors"

func main() {
	client, err := connectors.Mock()

	if err != nil {
		panic(err)
	}

	client.Close()
}
