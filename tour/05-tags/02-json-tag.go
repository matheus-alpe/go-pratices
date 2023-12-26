package tags

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"-"`                       // hide undesired output
	PreferredFish []string  `json:"preferredFish,omitempty"` // omitempty hides field when has zero value
	CreatedAt     time.Time `json:","`                       // "," doesn't change the field name
}

func JsonTagsExample() {
	u := &User{
		Name:      "Matt",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(out))
}
