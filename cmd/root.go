package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	defaultContextName = "config"
)

func init() {
	rootCmd.Flags().StringP("root", "r", os.Getenv("KUBEROOT"), "the locate of your config root e.g. ~/.kube")
	rootCmd.Flags().StringP("context", "c", "", "the location of your context file (if context is not a abs location it will use root as parent dir)")
	err := rootCmd.MarkFlagRequired("context")
	if err != nil {
		panic(err)
	}
}

var (
	rootCmd = &cobra.Command{
		Use:   "kube-change",
		Short: "Change context for your kubectl",
		Long:  "This is a command line tool to help you change the context of kubectl to easy manage multiple kubernetes cluster",
		Run: func(cmd *cobra.Command, args []string) {
			root, err := cmd.Flags().GetString("root")
			if err != nil {
				panic(fmt.Errorf("root not found %w", err))
			}
			context, err := cmd.Flags().GetString("context")
			if err != nil {
				panic(fmt.Errorf("context not found %w", err))
			}
			ok := filepath.IsAbs(context)
			if !ok {
				context = filepath.Join(root, context)
			}
			_, err = os.Stat(context)
			if err != nil {
				panic(err)
			}
			f, err := os.Open(context)
			if err != nil {
				panic(err)
			}
			bs, err := ioutil.ReadAll(f)
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(filepath.Join(root, defaultContextName), bs, os.ModePerm)
			if err != nil {
				panic(err)
			}
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
