# todoist-go

[![Actions Status](https://github.com/ides15/todoist/workflows/Go/badge.svg)](https://github.com/ides15/todoist/actions)

Golang client library for the V8 Todoist Sync API. This repository is in development.

Influenced by Google's Github API Golang Client: https://github.com/google/go-github

## Installation

```sh
go get -u github.com/ides15/todoist
```

## Creating a Client

All that is required to set up a client is your Todoist API token, which can be found at https://todoist.com/prefs/integrations

```go
client, err := todoist.NewClient("<YOUR_TODOIST_API_TOKEN>", false)
if err != nil {
    panic(err)
}
```

---

**NOTE**

The second parameter to `todoist.NewClient` is a debug flag. If you want logs to be shown in the console, set this to true.

---

<br/>

## Working with Resources

Through `todoist.Client`, you can work with any Todoist resource (projects, notes, items, etc.)

### Projects

```go
package todoist_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ides15/todoist/todoist"
)

var (
    // Get the Todoist API token from an environment variable
	apiToken = os.Getenv("TODOIST_API_TOKEN")
)

func getFirstFromMap(m map[string]int64) string {
	for key := range m {
		return key
	}

	return ""
}

func Test_Projects(t *testing.T) {
    // Create the client to interact with Todoist
	client, err := todoist.NewClient(apiToken, true)
	if err != nil {
		panic(err)
	}

    // List all projects
	projects, _, err := client.Projects.List(context.Background(), "")
	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		fmt.Println(*project.ID, *project.Name)
	}

    // Add a new project
	projects, resp, err := client.Projects.Add(context.Background(), "", &todoist.AddProject{
		Name: "not another new project...",
	})
	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		fmt.Println(*project.ID, *project.Name)
	}

    // Update the project we just added
	projects, _, err = client.Projects.Update(context.Background(), "", &todoist.UpdateProject{
		ID:   getFirstFromMap(resp.TempIDMapping), // get the temp ID of the project we just added so we can update the title
		Name: "an *updated* project!!!",
	})
	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		fmt.Println(*project.ID, *project.Name)
	}
}
```
