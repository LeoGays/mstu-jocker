package presenters_test

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"LeoGays/new-year-cards-slack/pkg/presenters"
)

func ExampleCropper() {
	cropped := presenters.NewCropper(10)
	src := bytes.NewBufferString("hello hello - is very lo-o-o-ong string!")
	if _, err := io.Copy(cropped, src); err != nil {
		log.Fatal(err)
	}

	fmt.Println(cropped)

	// Output:
	// hello hell...
}
