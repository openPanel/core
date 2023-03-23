package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func ent() {
	log.Println("running ent codegen")

	names := []string{"local", "shared"}
	for _, name := range names {
		err := entc.Generate("./app/db/schema/"+name, &gen.Config{
			Features: []gen.Feature{
				gen.FeatureUpsert,
				gen.FeatureNamedEdges,
				gen.FeatureModifier,
			},
			Target: "./app/generated/db/" + name,
			Schema: "./app/db/schema/" + name,
		})
		if err != nil {
			log.Fatalf("running ent codegen: %v", err)
		}
	}

	log.Println("running ent codegen done")
}
