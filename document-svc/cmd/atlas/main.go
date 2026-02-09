package main

import (
	"fmt"
	"log"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	document "document-svc/internal"
)

func main() {
	loader := gormschema.New("postgres")
	schema, err := loader.Load(document.Models()...)
	if err != nil {
		log.Fatalf("schema load error: %v", err)
	}
	if _, err := fmt.Fprint(os.Stdout, schema); err != nil {
		log.Fatalf("schema write error: %v", err)
	}
}
