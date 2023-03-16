package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubenodes",
	Short: "Top down view from nodes to pods in a namespace",
	Run: func(cmd *cobra.Command, args []string) {
		// if len(args) == 0 {
		// 	cmd.Help()
		// 	os.Exit(0)
		// }
		// app := view.NewApp(LoadAPIClient())
		flag.CommandLine.Parse([]string{})
		namespace, _ := cmd.Flags().GetString("namespace")
		deployment, _ := cmd.Flags().GetString("deployment")

		apiClient := LoadAPIClient()
		pods := GetPods(apiClient, namespace, deployment)

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error executing app: ", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("namespace", "n", "default", "kubernetes namespace")
	rootCmd.PersistentFlags().StringP("deployment", "d", "", "kubernetes deployment, looks for app=[deployment_name]")
	rootCmd.PersistentFlags().StringP("pod", "p", "", "kubernetes pods")
}
