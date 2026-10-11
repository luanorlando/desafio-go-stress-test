package cmd

import (
	"fmt"
	"os"
	"stress-test/loadtest"

	"github.com/spf13/cobra"
)

var (
	url         string
	requests    int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "load-tester",
	Short: "CLI para testes de carga em serviços web",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" || requests <= 0 || concurrency <= 0 {
			fmt.Println("Erro: As flags --url, --requests e --concurrency são obrigatórias e devem ser maiores que zero.")
			os.Exit(1)
		}

		fmt.Printf("Iniciando teste de carga para: %s\n", url)
		fmt.Printf("Requisições: %d | Concorrência: %d\n\n", requests, concurrency)

		metrics := loadtest.Run(loadtest.Config{
			URL:         url,
			Requests:    requests,
			Concurrency: concurrency,
		})

		// Exibição do Relatório
		fmt.Println("==========================================")
		fmt.Println("           RELATÓRIO DE TESTE             ")
		fmt.Println("==========================================")
		fmt.Printf("Tempo total gasto:       %v\n", metrics.TotalDuration)
		fmt.Printf("Total de requisições:    %d\n", metrics.TotalRequests)
		fmt.Printf("Requisições com 200 OK:  %d\n", metrics.Success200)
		fmt.Println("Distribuição de Status HTTP:")
		for code, count := range metrics.StatusDist {
			if code == 0 {
				fmt.Printf("  - Erros de Conexão:    %d\n", count)
			} else {
				fmt.Printf("  - Status %d:            %d\n", code, count)
			}
		}
		fmt.Println("==========================================")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&url, "url", "", "URL do serviço a ser testado (obrigatório)")
	rootCmd.Flags().IntVar(&requests, "requests", 0, "Número total de requisições (obrigatório)")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 0, "Número de chamadas simultâneas (obrigatório)")

	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
	rootCmd.MarkFlagRequired("concurrency")
}
