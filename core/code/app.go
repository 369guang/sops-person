package code

import "fmt"

var appFolder = []string{"apis", "models", "serializers"}

var appFile = []string{"README.md", "routes.go"}

func CreateFolder(name string) error {
	fmt.Println(appFolder)
	fmt.Println(appFile)
	return nil
}

func TouchFile() error {
	return nil
}

func TemplateReadme(name string) error {
	return nil
}

func TemplateRoutes(name string) error {
	return nil
}
