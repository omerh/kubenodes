package cmd

import (
	"fmt"
	"kubenodes/pkg/client"
	"kubenodes/pkg/render"
	"kubenodes/pkg/resource"
	"os"
	"time"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var version = "v0.0.1"

var rootCmd = &cobra.Command{
	Use:     "kubenodes",
	Version: version,
	Short:   "Top down view from nodes to pods in a namespace",
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		labelSlice, _ := cmd.Flags().GetStringSlice("label")
		kubeConfigPath, _ := cmd.Flags().GetString("kubeconfig")
		refresh, _ := cmd.Flags().GetInt("refresh")
		compact, _ := cmd.Flags().GetBool("compact")

		// Kubernetes client
		apiClient := client.LoadAPIClient(kubeConfigPath)

		// Retrieve running pods according to the deployment tag of app=
		pods := resource.GetPods(apiClient, namespace, labelSlice)
		// pods returns v1.PodList, converting to a unique map per node name from pod.spec
		podsOnNodesMap := resource.MakeUniquePodsOnNode(pods)

		// Convert map to nodeSlice with pod and node from GetPods
		nodesInfo := resource.NodeMapToNodes(podsOnNodesMap)
		// Update the instance info
		nodesInfo = resource.UpdateNodeInfoSlice(apiClient, nodesInfo)

		table := render.NodesPodsFullRender(nodesInfo, compact)

		app := tview.NewApplication()
		flex := tview.NewFlex().SetDirection(tview.FlexRow)
		flex.AddItem(table, 0, 1, false)

		timer := time.NewTicker(time.Duration(refresh) * time.Second)
		done := make(chan interface{})
		go func() {
			defer close(done)
			for {
				select {
				case <-timer.C:
					pods := resource.GetPods(apiClient, namespace, labelSlice)
					podsOnNodesMap := resource.MakeUniquePodsOnNode(pods)
					nodesInfo := resource.NodeMapToNodes(podsOnNodesMap)
					nodesInfo = resource.UpdateNodeInfoSlice(apiClient, nodesInfo)

					table.Clear()
					table = render.NodesPodsFullRender(nodesInfo, compact)

					app.QueueUpdateDraw(func() {
						app.SetRoot(table, true)
					})
				case <-done:
					return
				}
			}
		}()

		if err := app.SetRoot(table, true).Run(); err != nil {
			panic(err)
		}
		timer.Stop()
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
	rootCmd.PersistentFlags().StringP("kubeconfig", "k", "", "kubeconfig path")
	rootCmd.PersistentFlags().StringP("namespace", "n", "default", "kubernetes namespace")
	rootCmd.PersistentFlags().StringSliceP("label", "l", []string{}, "app pod label, looks for app=[deployment_name], -l a,b")
	rootCmd.PersistentFlags().IntP("refresh", "r", 5, "application refresh interval")
	rootCmd.PersistentFlags().Bool("compact", false, "how to see pod listing in the node view")
}
