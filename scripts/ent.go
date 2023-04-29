package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func ent() {
	log.Println("running ent codegen")

	protoExt, _ := entproto.NewExtension(entproto.WithProtoDir(""))

	names := []string{"local", "shared"}
	for _, name := range names {
		var exts []entc.Extension
		if name == "shared" {
			exts = append(exts, protoExt)
		}

		targetDir, err := filepath.Abs("./app/generated/db/" + name)
		if err != nil {
			log.Fatalf("failed to get absolute path: %v", err)
		}

		err = entc.Generate("./app/db/schema/"+name, &gen.Config{
			Features: []gen.Feature{
				gen.FeatureUpsert,
				gen.FeatureNamedEdges,
				gen.FeatureModifier,
			},
			Target:  targetDir,
			Package: "github.com/openPanel/core/app/generated/db/" + name,
			Schema:  "github.com/openPanel/core/app/db/schema/" + name,
		}, entc.Extensions(exts...))
		if err != nil {
			log.Fatalf("running ent codegen: %v", err)
		}

		if name == "shared" {
			moveGenerateProto()
		}

		log.Println("running ent codegen for " + name + " done")
	}

	log.Println("running ent codegen done")
}

func moveGenerateProto() {
	proto, err := os.ReadFile("./app/generated/db/shared/proto/entpb/entpb.proto")
	if err != nil {
		log.Fatalf("failed to read proto file: %v", err)
	}
	protoStr := string(proto)
	var processed string
	for _, line := range strings.Split(protoStr, "\n") {
		if strings.HasPrefix(line, "option go_package") {
			continue
		}
		processed += line + "\n"
	}
	err = os.WriteFile("./app/protos/entpb.proto", []byte(processed), 0644)
	if err != nil {
		log.Fatalf("failed to write proto file: %v", err)
	}

	err = os.RemoveAll("./app/generated/db/shared/proto")
	if err != nil {
		log.Fatalf("failed to remove proto dir: %v", err)
	}
}
