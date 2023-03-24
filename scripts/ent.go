package main

import (
	"log"

	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func ent() {
	log.Println("running ent codegen")

	names := []string{"local", "shared"}
	for _, name := range names {
		var hooks = make([]gen.Hook, 0)
		if name == "shared" {
			hooks = append(hooks, entproto.Hook())
		}

		err := entc.Generate("./app/db/schema/"+name, &gen.Config{
			Features: []gen.Feature{
				gen.FeatureUpsert,
				gen.FeatureNamedEdges,
				gen.FeatureModifier,
			},
			Target:  "./app/generated/db/" + name,
			Package: "github.com/openPanel/core/app/generated/db/" + name,
			Schema:  "github.com/openPanel/core/app/db/schema/" + name,
			Hooks:   hooks,
		})
		if err != nil {
			log.Fatalf("running ent codegen: %v", err)
		}
		log.Println("running ent codegen for " + name + " done")
	}

	log.Println("running ent codegen done")
}
