package cmd

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed templates/*.tpl
var templatesFS embed.FS

func createFileFromTemplate(templateName, entityName string) error {
	// Carrega o conteúdo do template
	content, err := templatesFS.ReadFile(fmt.Sprintf("templates/%s.go.tpl", templateName))
	if err != nil {
		return fmt.Errorf("erro ao ler o template: %w", err)
	}

	// Substitui o marcador no template (exemplo: {{.Entity}})
	result := strings.ReplaceAll(string(content), "{{.Entity}}", entityName)

	// Crie o arquivo de destino
	outputDir := fmt.Sprintf("./%ss", templateName) // Exemplo: models/
	os.MkdirAll(outputDir, os.ModePerm)             // Cria o diretório se não existir
	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s.go", entityName))

	// Escreve o arquivo de saída
	if err := os.WriteFile(outputFile, []byte(result), 0644); err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %w", err)
	}

	return nil
}

var makeCmd2 = &cobra.Command{
	Use:   "make2",
	Short: "Gera arquivos base usando templates embutidos",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Uso: make <template> <entity>")
			return
		}

		templateName := args[0]
		entityName := args[1]

		err := createFileFromTemplate(templateName, entityName)
		if err != nil {
			fmt.Printf("Erro ao gerar arquivo: %v\n", err)
			return
		}

		fmt.Println("Arquivo gerado com sucesso:", fmt.Sprintf("%s.go", entityName))
	},
}

func init() {
	rootCmd.AddCommand(makeCmd2)
}
