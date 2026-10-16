package writer_test

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/psyb0t/go-writer-ai"
)

// This is the README usage example. It talks to the live API, so it has no
// Output comment and only has to compile.
//
//nolint:testableexamples // needs a real API key, compile check only
func Example() {
	auth := writer.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			key := os.Getenv("WRITER_API_KEY")
			req.Header.Set("Authorization", "Bearer "+key)

			return nil
		},
	)

	client, err := writer.NewClientWithResponses(
		"https://api.writer.com",
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.ModelsWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	if res.JSON200 == nil {
		panic(res.Status())
	}

	for _, model := range res.JSON200.Models {
		fmt.Println(model.Id)
	}
}
