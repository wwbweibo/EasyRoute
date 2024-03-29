package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"golang.org/x/tools/go/packages"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

var (
	appName         = "rpcgenerator"
	generateComment = "//go:generate"
	buf             = bytes.Buffer{}
)

func main() {
	pkg := parsePackage([]string{"."}, []string{})
	for _, syntax := range pkg.Syntax {
		if isNeedGenerate(syntax.Comments) {
			var imports []string
			if syntax.Imports != nil {
				for _, v := range syntax.Imports {
					imports = append(imports, v.Path.Value)
				}
			}

			generateFileHead(pkg.Name, imports)

			types := parseTypes(syntax.Decls)

			for _, spec := range types {
				generateConstructor(spec.Name.Name, spec.Type.(*ast.StructType).Fields)
			}

			fileBytes := buf.Bytes()
			src, err := format.Source(fileBytes)
			if err != nil {
				err = ioutil.WriteFile(syntax.Name.Name+"_gen.go", fileBytes, 0644)
				if err != nil {
					panic("write content to file error")

				}
			} else {
				err = ioutil.WriteFile(syntax.Name.Name+"_gen.go", src, 0644)
				if err != nil {
					panic("write content to file error")
				}
			}
		}
	}
}

// generateFileHead will generate top file comment, package and imports
func generateFileHead(pkg string, imports []string) {
	bufferPrint("// code generate by codegenerator, DO NOT EDIT;\n\n")
	bufferPrint("package %s\n\n", pkg)
	// will use rpc.Config
	imports = append(imports, "\"github.com/wwbweibo/EasyRoute/rpc\"")
	if len(imports) > 0 {
		bufferPrint("import (\n")
		for _, v := range imports {
			bufferPrint("\t%s\n", v)
		}
		bufferPrint(")\n\n")
	}
}

// generateConstructor will generate constructor for every type
func generateConstructor(typeName string, fields *ast.FieldList) {
	bufferPrint("func New%s(config rpc.Config) *%s {\n", typeName, typeName)
	bufferPrint("\treturn &%s{\n", typeName)

	for _, field := range fields.List {
		if fieldType, ok := (field.Type).(*ast.FuncType); ok {
			bufferPrint("\t\t%s: func(", field.Names[0].Name)

			// generate method args
			if fieldType.Params.NumFields() == 1 {
				f := fieldType.Params.List[0]
				bufferPrint("%s %s", f.Names[0].Name, (f.Type).(*ast.Ident).Name)
			} else {
				for idx, f := range fieldType.Params.List {
					bufferPrint("%s %s", f.Names[0].Name, (f.Type).(*ast.Ident).Name)
					if idx != fieldType.Params.NumFields()-1 {
						bufferPrint(", ")
					}
				}
			}
			bufferPrint(") ")

			resultType := ""

			// generate method return type
			bufferPrint("(")
			for idx, f := range fieldType.Results.List {
				bufferPrint("%s", (f.Type).(*ast.Ident).Name)
				if idx == 0 {
					resultType = (f.Type).(*ast.Ident).Name
				}
				if idx != fieldType.Results.NumFields()-1 {
					bufferPrint(", ")
				}
			}
			bufferPrint(")")

			bufferPrint(" {\n")

			// get http end point information from tag
			tags := parseFieldTag(field)
			method, ok := tags["method"]
			if !ok {
				method = "Get"
			}
			var params []string
			param, ok := tags["param"]
			if ok {
				params = strings.Split(param, ",")
			}
			methodName, ok := tags["route"]
			if !ok {
				methodName = field.Names[0].Name
			}
			// generate http request information
			generateHttpRequest(typeName, methodName, strings.ReplaceAll(method, "\"", ""), params, resultType)

			bufferPrint("\t\t},\n")
		}
	}
	bufferPrint("\t}\n}\n")
}

func generateHttpRequest(controllerName string, methodName string, method string, args []string, responseType string) {
	controllerName = strings.ReplaceAll(strings.ToLower(controllerName), "controller", "")
	if strings.ToLower(method) == "get" {
		if isSimpleValueType(responseType) {
			if isNumberType(responseType) {
				bufferPrint("\t\t\tvar result %s = 0\n", responseType)
			} else if isString(responseType) {
				bufferPrint("\t\t\tresult := \"\"\n")
			}
		} else {
			bufferPrint("\t\t\tresult := %s{}\n", responseType)
		}
		if len(args) == 0 {
			bufferPrint("\t\t\terr := rpc.HttpGet(config, \"%s\",\"%s\", nil, &result)\n", controllerName, methodName)
		} else {
			bufferPrint("\t\t\tparams := make(map[string]string)\n")
			for _, arg := range args {
				bufferPrint("\t\t\tparams[%s] = rpc.JsonSerialize(%s)\n", arg, strings.ReplaceAll(arg, "\"", ""))
			}
			bufferPrint("\t\t\terr := rpc.HttpGet(config, \"%s\",\"%s\", params, &result)\n", controllerName, methodName)
		}

		bufferPrint("\t\t\tif err != nil {return result, err}\n")
		bufferPrint("\t\t\treturn result, nil\n")
	} else if strings.ToLower(method) == "post" {
		if isSimpleValueType(responseType) {
			if isNumberType(responseType) {
				bufferPrint("\t\t\tvar result %s = 0\n", responseType)
			} else if isString(responseType) {
				bufferPrint("\t\t\tresult := \"\"\n")
			}
		} else {
			bufferPrint("\t\t\tresult := %s{}\n", responseType)
		}
		if len(args) == 0 {
			bufferPrint("\t\t\terr := rpc.HttpPost(config, \"%s\" , \"%s\", nil, &result)\n", controllerName, methodName)
		} else {
			bufferPrint("\t\t\tparams := make(map[string]string)\n")
			for _, arg := range args {
				bufferPrint("\t\t\tparams[%s] = rpc.JsonSerialize(%s)\n", arg, strings.ReplaceAll(arg, "\"", ""))
			}
			bufferPrint("\t\t\terr := rpc.HttpGet(config, \"%s\", \"%s\", params, &result)\n", controllerName, methodName)
		}

		bufferPrint("\t\t\tif err != nil {return result, err}\n")
		bufferPrint("\t\t\treturn result, nil\n")
	}
}

func parsePackage(patterns, tags []string) *packages.Package {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports |
			packages.NeedTypes | packages.NeedTypesSizes | packages.NeedSyntax | packages.NeedTypesInfo,
		// TODO: Need to think about constants in test files. Maybe write type_string_test.go
		// in a separate pass? For later.
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	return pkgs[0]
}

func parseTypes(decls []ast.Decl) []*ast.TypeSpec {
	var types []*ast.TypeSpec
	for _, decl := range decls {
		d, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if isNeedGenerate([]*ast.CommentGroup{d.Doc}) {
			for _, spec := range d.Specs {
				if sp, ok := spec.(*ast.TypeSpec); ok {
					if _, ok := sp.Type.(*ast.StructType); ok {
						types = append(types, sp)
					}
				}
			}
		}
	}
	return types
}

func parseFieldTag(field *ast.Field) map[string]string {
	tagMap := make(map[string]string)
	tagStr := strings.ReplaceAll(field.Tag.Value, "`", "")
	tags := strings.Split(tagStr, " ")
	for _, tag := range tags {
		v := strings.Split(tag, ":")
		tagMap[v[0]] = v[1]
	}
	return tagMap
}

// isNeedGenerate will walk all the given comments to check if there is comment like '//go:generate' ,
// and need to generate by this app
func isNeedGenerate(comments []*ast.CommentGroup) bool {
	for _, comment := range comments {
		if comment == nil {
			return false
		}
		for _, c := range comment.List {
			if len(c.Text) > len(generateComment) && c.Text[:len(generateComment)] == generateComment {
				if strings.Trim(strings.Replace(c.Text, generateComment, "", 1), " ")[:len(appName)] == appName {
					return true
				}
			}
		}

	}

	return false
}

func isSimpleValueType(typeName string) bool {
	return isNumberType(typeName) || isString(typeName) || isArray(typeName)
}

func isNumberType(typeName string) bool {
	return typeName == reflect.Int.String() ||
		typeName == reflect.Int8.String() ||
		typeName == reflect.Int16.String() ||
		typeName == reflect.Int32.String() ||
		typeName == reflect.Int64.String() ||
		typeName == reflect.Uint.String() ||
		typeName == reflect.Uint8.String() ||
		typeName == reflect.Uint16.String() ||
		typeName == reflect.Uint32.String() ||
		typeName == reflect.Uint64.String() ||
		typeName == reflect.Float64.String() ||
		typeName == reflect.Float32.String()
}

func isString(typeName string) bool {
	return typeName == reflect.String.String()
}

func isArray(typeName string) bool {
	return typeName[0:2] == "[]"
}

func bufferPrint(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(&buf, format, args...)
}
