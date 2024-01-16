package main

import (
	"canvas_users_service/cmd"
	"fmt"
	"net/http"
)

func main() {
	db := cmd.Init()
	defer db.Close()

	// TODO: Register the routes and outline + define the API, separate RPC's in the URI path
	http.HandleFunc("/user_service", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
        Global Params:

        CANVAS_URL and CANVAS_KEY can either be loaded from .env file, or passed in body of request as JSON.

        UNLOCKED_URL and UNLOCKED_KEY can either be loaded from .env file, or passed in body of request as JSON.

        "Export all users: POST /user_service/export_users" (returns an array of CanvasUsers)

        "Import users: POST /user_service/import_users" (returns an array of UnlockedUsers and an array of ProviderUserMapping)

        Params:
        all: true (imports all users from canvas instance to Unlocked) 
        users: [1, 2, 3] (imports specific users with ID's 1, 2, 3)
        accounts: [1, 2, 3] (imports all users from specific accounts with ID's 1, 2, 3)

        "Match users: POST /user_service/match_users" (returns an array of ProviderUserMapping)
        Params:
        all: true (matches all users from canvas instance to Unlocked) 
        users: [1, 2, 3] (matches specific users with ID's 1, 2, 3)
        accounts: [1, 2, 3] (matches all users from specific accounts with ID's 1, 2, 3)
        `)
	})

	// Start the server on port 8080
	http.ListenAndServe(":8090", nil)
}
