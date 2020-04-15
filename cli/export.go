package cli

import (
	"github.com/spf13/cobra"

	"tbl-editor/editor"
)

var (
	inFile    string
	outFile   string
	exportCmd = &cobra.Command{
		Use:   "export",
		Short: "Exports a .tbl file into an excel table",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {

			if inFile == "" || outFile == "" {
				cmd.Usage()
				return
			}

			editor.Export(inFile, outFile)
		},
	}
)

func init() {
	exportCmd.Flags().StringVarP(&inFile, "input", "i", "", "Input file")
	exportCmd.Flags().StringVarP(&outFile, "output", "o", "", "Output file")

	rootCmd.AddCommand(exportCmd)
}
