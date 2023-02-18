package codegen

//go:generate go run gen.go input_schema.yaml

func CodeGen() string {
	return "arcaflow-engine/codegen.CodeGen"
}
