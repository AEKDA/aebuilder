package builder

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

func (a *App) ReadAst() error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, a.Source, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	return a.WriteAst(fset, f)
}

func (a *App) WriteAst(fset *token.FileSet, node *ast.File) error {
	foundStruct, fileNode := findStruct(node, a.Name)
	if foundStruct == nil {
		return errors.New("Структура не найдена")
	}

	typeSpec, ok := foundStruct.Specs[0].(*ast.TypeSpec)
	if ok {
		removeFieldsFromStruct(typeSpec, []string{"status", "state", "sizeCache", "unknownFields"})
		removeTagsFromStruct(typeSpec)
	}

	// Откройте файл для записи
	dstFile, err := os.Create(a.Output)
	if err != nil {
		log.Fatal(err)
	}

	defer dstFile.Close()

	packageName := getPackageName(a.Output)

	file := &ast.File{
		Name: ast.NewIdent(packageName),
		Decls: []ast.Decl{
			foundStruct,
		},
		Comments: fileNode.Comments,
		Scope:    fileNode.Scope,
	}

	fset2 := token.NewFileSet()
	// Преобразуйте AST-дерево обратно в исходный код и запишите его в файл
	if err := printer.Fprint(dstFile, fset2, file); err != nil {
		return err
	}
	return nil
}

// Функция для удаления всех тегов из структуры
func removeTagsFromStruct(typeSpec *ast.TypeSpec) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return
	}
	for _, field := range structType.Fields.List {
		field.Tag = nil
	}
}

// Функция для поиска структуры по имени
func findStruct(node ast.Node, structName string) (*ast.GenDecl, *ast.File) {
	var (
		foundStruct *ast.GenDecl
		fileNode    *ast.File
	)
	ast.Inspect(node, func(n ast.Node) bool {
		if file, ok := n.(*ast.File); ok {
			fileNode = file
		}
		if genDecl, ok := n.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if typeSpec.Name.Name == structName {
						foundStruct = genDecl
						return false
					}
				}
			}
		}
		return true
	})
	return foundStruct, fileNode
}

// Функция для удаления определенных полей из структуры
func removeFieldsFromStruct(typeSpec *ast.TypeSpec, fieldsToRemove []string) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return
	}
	var newFields []*ast.Field
	for _, field := range structType.Fields.List {
		keep := true
		for _, fieldName := range fieldsToRemove {
			for _, name := range field.Names {
				if name.Name == fieldName {
					keep = false
					break
				}
			}
		}
		if keep {
			newFields = append(newFields, field)
		}
	}
	structType.Fields.List = newFields
}

// Функция для получения имени пакета из пути
func getPackageName(dstPath string) string {
	// Получаем родительскую директорию файла
	dir := filepath.Dir(dstPath)
	// Имя пакета равно имени родительской директории
	packageName := filepath.Base(dir)
	return packageName
}
