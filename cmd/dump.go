// Copyright © 2018 Harry Bagdi <harrybagdi@gmail.com>

package cmd

import (
	"github.com/kong/deck/dump"
	"github.com/kong/deck/file"
	"github.com/spf13/cobra"
)

var dumpCmdKongStateFile string

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Export Kong configuration to a file",
	Long: `A longer description that spans multiple lines
and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO move this out to a common factory
		client, err := GetKongClient(config)
		if err != nil {
			return err
		}

		ks, err := dump.GetState(client)
		if err != nil {
			return err
		}
		if err := file.KongStateToFile(ks, dumpCmdKongStateFile); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
	dumpCmd.Flags().StringVarP(&dumpCmdKongStateFile, "output-file", "o",
		"kong.yaml", "write Kong configuration to FILE")
}