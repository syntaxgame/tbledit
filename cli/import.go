package cli

import (
	"github.com/spf13/cobra"

	"tbl-editor/editor"
)

var (
	importCmd = &cobra.Command{
		Use:   "import",
		Short: "Imports an excel file into .tbl file",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {

			if inFile == "" || outFile == "" {
				cmd.Usage()
				return
			}

			editor.Import(inFile, outFile)
		},
	}
)

func init() {
	importCmd.Flags().StringVarP(&inFile, "input", "i", "", "Input file")
	importCmd.Flags().StringVarP(&outFile, "output", "o", "", "Output file")

	rootCmd.AddCommand(importCmd)
}
