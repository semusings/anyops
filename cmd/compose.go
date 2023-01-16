package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "Docker compose commands",
}

var composeDir = filepath.Join("compose")
var composeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available  docker compose apps",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = filepath.Walk(composeDir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				fmt.Printf("- %s\n", strings.TrimSuffix(info.Name(), ".yaml"))
			}
			return nil
		})
	},
}

type composeFn func(composeFile string) string

func composeUp(composeFile string) string {
	return fmt.Sprintf("docker-compose -f %s up -d --remove-orphans", composeFile)
}

func composeLogs(composeFile string) string {
	return fmt.Sprintf("docker-compose -f %s logs -f", composeFile)
}

func composeDown(composeFile string) string {
	return fmt.Sprintf("docker-compose -f %s down --remove-orphans", composeFile)
}

var composeUpCmd = buildComposeCmd("up", composeUp)
var composeLogsCmd = buildComposeCmd("logs", composeLogs)
var composeDownCmd = buildComposeCmd("down", composeDown)

func buildComposeCmd(command string, f composeFn) *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("%s <app_name>", command),
		Short: fmt.Sprintf("To run docker compose %s for app <app_name>", command),
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires <app_name> arguments")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			composeFile := filepath.Join(composeDir, fmt.Sprintf("%s.yaml", args[0]))
			ExecuteCommand(CommandRequest{Parts: []string{f(composeFile)}})
		},
	}
}

func init() {
	rootCmd.AddCommand(composeCmd)
	composeCmd.AddCommand(composeUpCmd)
	composeCmd.AddCommand(composeLogsCmd)
	composeCmd.AddCommand(composeDownCmd)
	composeCmd.AddCommand(composeListCmd)
}
