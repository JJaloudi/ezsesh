package main

import (
	"fmt"
	"github.com/jjaloudi/ezsesh"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"log"
	"net/http"
)

func main() {
	db, err := sqlx.Open("postgres", "user=postgres password=password search_path=server_status sslmode=disable host=localhost port=5432")
	if err != nil {
		log.Fatal(err)
	}
	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	options := &ezsesh.EzOptions{
		Table:       "session",
		Association: "user_id",
		Lifetime:    5000,
		SingleToken: true,

		CookieName: "test",
		SameSite:   http.SameSiteStrictMode,
		HttpOnly:   true,
		Secure:     true,
	}
	store := ezsesh.CreateEZStore(options, db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		store.Generate(w, "c6eeae8e-002e-468e-8767-691897bc2fec")

		c, e := r.Cookie("session-cookie")

		if e != nil {
			return
		}

		fmt.Printf(c.Value)
		w.Write([]byte(c.Value))
	})

	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
