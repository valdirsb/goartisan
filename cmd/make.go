package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var makeCmd = &cobra.Command{
	Use:   "make [type] [name]",
	Short: "Generate a Go file from a template (e.g., make model User)",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fileType := strings.ToLower(args[0]) // Tipo do arquivo (model, repository, controller)
		name := args[1]                      // Nome do arquivo/estrutura

		// Mapeie o tipo para o arquivo de template correspondente
		templatesDir := "./templates"
		templateFile := filepath.Join(templatesDir, fmt.Sprintf("%s.go.tpl", fileType))

		if _, err := os.Stat(templateFile); os.IsNotExist(err) {
			fmt.Println("Template não encontrado para:", fileType)
			return
		}

		// Leia e parseie o arquivo de template
		tmpl, err := template.ParseFiles(templateFile)
		if err != nil {
			fmt.Println("Erro ao carregar o template:", err)
			return
		}

		// Crie o arquivo de destino
		outputDir := fmt.Sprintf("./%ss", fileType) // Exemplo: models/
		os.MkdirAll(outputDir, os.ModePerm)         // Cria o diretório se não existir
		outputFile := filepath.Join(outputDir, fmt.Sprintf("%s.go", name))

		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Erro ao criar o arquivo:", err)
			return
		}
		defer file.Close()

		// Execute o template e escreva o conteúdo no arquivo
		err = tmpl.Execute(file, struct{ Name string }{Name: name})
		if err != nil {
			fmt.Println("Erro ao preencher o template:", err)
			return
		}

		fmt.Printf("Arquivo '%s' gerado com sucesso!\n", outputFile)
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}
