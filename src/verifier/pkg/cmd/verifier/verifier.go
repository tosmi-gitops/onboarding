package verifier

import "fmt"
import "github.com/spf13/cobra"

func Execute(name string) error {
	fmt.Printf("Hello, World %s!\n", name)

	rootCmd := &cobra.Command{
		Use:   "verifier",
		Short: "Verify cluster onboarding configuration",
		Long: `verifier checks a standardized kustomize onboarding
configuration for obvious errors.`,
	}

	var root string
	rootCmd.PersistentFlags().StringVar(&root, "root", ".", "root directory of onboarding configuration")

	return rootCmd.Execute()
}
