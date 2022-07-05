package pipeline

import (
	"cli_go.com/b/vault"
	"github.com/spf13/cobra"
)

var user string
var password string
var authCmd = &cobra.Command{
	Use:     "load vault-auth",
	Aliases: []string{"va"},
	Short:   "Seta token de autenticacao das APIs do cofre no ambiente",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vault.GetAuthTokenVault(user, password)
	},
}

var secretsCmd = &cobra.Command{
	Use:     "pipeline-load secrets",
	Aliases: []string{"s"},
	Short:   "Seta secrets definidos no app.yaml para o ambiente",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vault.PopulateSecrets()
	},
}

func init() {
	authCmd.Flags().StringVarP(&user, "user", "u", "", "Usuario para login no cofre")
	authCmd.Flags().StringVarP(&password, "password", "p", "", "Senha para login no cofre")
	authCmd.MarkFlagsRequiredTogether("user", "password")

	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(secretsCmd)
}
