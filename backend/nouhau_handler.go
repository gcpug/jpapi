package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NouhauHandler(w http.ResponseWriter, r *http.Request) {
	sample := []string{
		"app-engine/note/be-careful-at-designing",
		"app-engine/note/gradle-plugin-throw-npe",
		"app-engine/example/sync-once-example",
		"spanner/note/benchmark/sinmetal-data-migration",
		"spanner/note/mutation-count",
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Cache-Control", "max-age=900")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(sample); err != nil {
		fmt.Println(err)
	}
}
