package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	uuidTest()
}
func uuidTest() {
	u := uuid.New()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t.UTC().Format(time.RFC3339Nano)+":", u.String())
		}
	}
}
