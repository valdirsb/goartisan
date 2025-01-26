package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func createModulePorts(entityName, projectName string) error {
	// Carrega o conteúdo do template

	//Modules Ports
	content, err := templatesFS.ReadFile(fmt.Sprintf("templates/module_ports.go.tpl"))
	if err != nil {
		return fmt.Errorf("erro ao ler o template: %w", err)
	}

	// Substitui o marcador no template (exemplo: {{.Entity}})
	//Modules Ports
	result := strings.ReplaceAll(string(content), "{{.Entity}}", entityName)
	result2 := strings.ReplaceAll(result, "{{.projectName}}", projectName)

	// Crie o arquivo de destino
	outputDir := fmt.Sprintf("./modules/%s", entityName) // Exemplo: module/User
	os.MkdirAll(outputDir, os.ModePerm)                  // Cria o diretório se não existir

	outputFile := filepath.Join(outputDir, fmt.Sprintf("ports.go"))

	// Escreve o arquivo de saída

	//Modules Ports
	if err := os.WriteFile(outputFile, []byte(result2), 0644); err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %w", err)
	}

	fmt.Println("Arquivo gerado com sucesso:", fmt.Sprintf(outputFile))

	return nil
}

func createModuleUsecase(entityName, projectName string) error {

	// Carrega o conteúdo do template

	//Modules Usecase
	content, err := templatesFS.ReadFile(fmt.Sprintf("templates/module_usecase.go.tpl"))
	if err != nil {
		return fmt.Errorf("erro ao ler o template: %w", err)
	}

	// Substitui o marcador no template (exemplo: {{.Entity}})

	//Modules Usecase
	result := strings.ReplaceAll(string(content), "{{.Entity}}", entityName)
	result2 := strings.ReplaceAll(result, "{{.projectName}}", projectName)

	// Crie o arquivo de destino
	outputDir := fmt.Sprintf("./modules/%s/usecase", entityName) // Exemplo: module/User
	os.MkdirAll(outputDir, os.ModePerm)                          // Cria o diretório se não existir

	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_usecase.go", entityName))

	// Escreve o arquivo de saída

	//Modules Ports
	if err := os.WriteFile(outputFile, []byte(result2), 0644); err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %w", err)
	}

	fmt.Println("Arquivo gerado com sucesso:", fmt.Sprintf(outputFile))

	return nil
}

func createModuleRepository(entityName, projectName string) error {

	// Carrega o conteúdo do template

	fmt.Println(strings.Title(strings.ToLower(projectName)))

	//Modules Usecase
	content, err := templatesFS.ReadFile(fmt.Sprintf("templates/module_repository_interface.go.tpl"))
	if err != nil {
		return fmt.Errorf("erro ao ler o template: %w", err)
	}

	// Substitui o marcador no template (exemplo: {{.Entity}})

	//Modules Usecase
	result := strings.ReplaceAll(string(content), "{{.Entity}}", entityName)

	result2 := strings.ReplaceAll(result, "{{.projectName}}", projectName)

	// Crie o arquivo de destino
	outputDir := fmt.Sprintf("./modules/%s/repository", entityName) // Exemplo: module/User
	os.MkdirAll(outputDir, os.ModePerm)                             // Cria o diretório se não existir

	outputFile := filepath.Join(outputDir, fmt.Sprintf("repository_interface.go"))

	// Escreve o arquivo de saída

	//Modules Ports
	if err := os.WriteFile(outputFile, []byte(result2), 0644); err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %w", err)
	}

	fmt.Println("Arquivo gerado com sucesso:", fmt.Sprintf(outputFile))

	return nil
}

func createModuleDomain(entityName, projectName string) error {

	// Carrega o conteúdo do template

	//Modules Usecase
	content, err := templatesFS.ReadFile(fmt.Sprintf("templates/module_domain.go.tpl"))
	if err != nil {
		return fmt.Errorf("erro ao ler o template: %w", err)
	}

	// Substitui o marcador no template (exemplo: {{.Entity}})

	//Modules Usecase
	result := strings.ReplaceAll(string(content), "{{.Entity}}", entityName)

	result2 := strings.ReplaceAll(result, "{{.projectName}}", projectName)

	// Crie o arquivo de destino
	outputDir := fmt.Sprintf("./modules/%s/domain", entityName) // Exemplo: module/User
	os.MkdirAll(outputDir, os.ModePerm)                         // Cria o diretório se não existir

	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s.go", entityName))

	// Escreve o arquivo de saída

	//Modules Ports
	if err := os.WriteFile(outputFile, []byte(result2), 0644); err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %w", err)
	}

	fmt.Println("Arquivo gerado com sucesso:", fmt.Sprintf(outputFile))

	return nil
}

func createModuleSqlc(entityName, projectName string) error {

	// Carrega o conteúdo do template

	//Modules Usecase
	content, err := templatesFS.ReadFile(fmt.Sprintf("templates/module_sqlc.yaml.tpl"))
	if err != nil {
		return fmt.Errorf("erro ao ler o template: %w", err)
	}

	// Substitui o marcador no template (exemplo: {{.Entity}})

	//Modules Usecase
	result := strings.ReplaceAll(string(content), "{{.Entity}}", entityName)

	result2 := strings.ReplaceAll(result, "{{.projectName}}", projectName)

	// Crie o arquivo de destino
	outputDir := fmt.Sprintf("./modules/%s", entityName) // Exemplo: module/User
	os.MkdirAll(outputDir, os.ModePerm)                  // Cria o diretório se não existir

	outputFile := filepath.Join(outputDir, fmt.Sprintf("sqlc.yaml"))

	// Escreve o arquivo de saída

	//Modules Ports
	if err := os.WriteFile(outputFile, []byte(result2), 0644); err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %w", err)
	}

	fmt.Println("Arquivo gerado com sucesso:", fmt.Sprintf(outputFile))

	return nil
}

var makeCmd3 = &cobra.Command{
	Use:   "make:module",
	Short: "Gera arquivos de módulo",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Uso: make:module <module>")
			return
		}

		entityName := capitalizeFirstLetter(args[0])

		projectName := "app.mentor.space"

		err := createModulePorts(entityName, projectName)
		if err != nil {
			fmt.Printf("Erro ao gerar arquivo: %v\n", err)
			return
		}

		err = createModuleUsecase(entityName, projectName)
		if err != nil {
			fmt.Printf("Erro ao gerar arquivo: %v\n", err)
			return
		}

		err = createModuleRepository(entityName, projectName)
		if err != nil {
			fmt.Printf("Erro ao gerar arquivo: %v\n", err)
			return
		}

		err = createModuleDomain(entityName, projectName)
		if err != nil {
			fmt.Printf("Erro ao gerar arquivo: %v\n", err)
			return
		}

		err = createModuleSqlc(entityName, projectName)
		if err != nil {
			fmt.Printf("Erro ao gerar arquivo: %v\n", err)
			return
		}

	},
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func init() {
	rootCmd.AddCommand(makeCmd3)
}
