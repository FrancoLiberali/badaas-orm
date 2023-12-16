package conditions

import (
	"errors"
	"go/types"

	"github.com/dave/jennifer/jen"
	"github.com/ditrit/badaas-orm/cli/cmd/log"
	"github.com/ditrit/badaas-orm/cli/cmd/version"
)

const badaasORMPath = "github.com/ditrit/badaas/orm"

type File struct {
	destPkg string
	jenFile *jen.File
	name    string
}

func NewConditionsFile(destPkg string, name string) *File {
	// Start a new file in destination package
	f := jen.NewFile(destPkg)

	// Add a package comment, so IDEs detect files as generated
	f.PackageComment("Code generated by badaas-cli v" + version.Version + ", DO NOT EDIT.")

	return &File{
		destPkg: destPkg,
		name:    name,
		jenFile: f,
	}
}

// Add conditions for an object in the file
func (file File) AddConditionsFor(object types.Object) error {
	fields, err := getFields(Type{object.Type()}, "")
	if err != nil {
		return err
	}

	log.Logger.Infof("Generating conditions for type %q in %s", object.Name(), file.name)

	file.addConditionsForEachField(object, fields)
	return nil
}

// Add one condition for each field of the object
func (file File) addConditionsForEachField(object types.Object, fields []Field) {
	conditions := file.generateConditionsForEachField(object, fields)

	for _, condition := range conditions {
		for _, code := range condition.codes {
			file.jenFile.Add(code)
		}
	}
}

// Write generated file
func (file File) Save() error {
	return file.jenFile.Save(file.name)
}

// Generate the conditions for each of the object's fields
func (file File) generateConditionsForEachField(object types.Object, fields []Field) []*Condition {
	conditions := []*Condition{}
	for _, field := range fields {
		log.Logger.Debugf("Generating condition for field %q", field.Name)
		if field.Embedded {
			conditions = append(conditions, file.generateEmbeddedConditions(
				object,
				field,
			)...)
		} else {
			conditions = append(conditions, NewCondition(
				file.destPkg, Type{object.Type()}, field,
			))
		}
	}

	return conditions
}

// Generate conditions for a embedded field
// it will generate a condition for each of the field of the embedded field's type
func (file File) generateEmbeddedConditions(object types.Object, field Field) []*Condition {
	embeddedStructType, ok := field.Type.Underlying().(*types.Struct)
	if !ok {
		panic(errors.New("unreachable! embedded objects are always structs"))
	}

	fields, err := getStructFields(embeddedStructType, field.Tags.getEmbeddedPrefix())
	if err != nil {
		// embedded field's type has not fields
		return []*Condition{}
	}

	return file.generateConditionsForEachField(object, fields)
}
