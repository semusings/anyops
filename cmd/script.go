package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Bash script commands",
}

var scriptDir = filepath.Join("script")
var scriptListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available script",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = filepath.Walk(scriptDir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				fmt.Printf("- %s\n", strings.TrimSuffix(info.Name(), ".sh"))
			}
			return nil
		})
	},
}

var scriptRunCmd = buildScriptCmd("run")

func buildScriptCmd(command string) *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("%s", command),
		Short: fmt.Sprintf("To %s script", command),
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			scriptFile := filepath.Join(scriptDir, fmt.Sprintf("%s.sh", args[0]))
			ExecuteCommand(CommandRequest{Args: []string{readScript(scriptFile)}, Envs: args[1:]})
		},
	}
}

func readScript(scriptFile string) string {
	data, err := os.ReadFile(workingDir + "/" + scriptFile)
	if err != nil {
		LogFatal("%s not found: %s", scriptFile, err)
	}
	return string(data)
}

func init() {
	rootCmd.AddCommand(scriptCmd)
	scriptCmd.AddCommand(scriptRunCmd)
	scriptCmd.AddCommand(scriptListCmd)
}
