package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var saFile string

var rootCmd = &cobra.Command{
	Use:   "firease",
	Short: "Firease is a cli tool to interact with firestore database",
	Long:`
------------------------------------------------------
███████╗██╗██████╗ ███████╗ █████╗ ███████╗███████╗
██╔════╝██║██╔══██╗██╔════╝██╔══██╗██╔════╝██╔════╝
█████╗  ██║██████╔╝█████╗  ███████║███████╗█████╗
██╔══╝  ██║██╔══██╗██╔══╝  ██╔══██║╚════██║██╔══╝
██║     ██║██║  ██║███████╗██║  ██║███████║███████╗
╚═╝     ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝
------------------------------------------------------
Interact with firestore interactively
------------------------------------------------------`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&saFile, "safile", "", "Service account authentication json file")
	rootCmd.MarkPersistentFlagRequired("safile")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
