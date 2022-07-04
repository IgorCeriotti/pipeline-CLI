package pipeline

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "pipeline - CLI para executar estágios de esteiras de deploy",
	Long: `pipeline será uma CLI completa para execução agnóstica de ambiente e cloud provider para CI/CD pipelines.
			Também pode ser possível utilizar como self-service.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Falha ao executar cmd! '%s'", err)
		os.Exit(1)
	}
}
