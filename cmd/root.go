package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
)

var colorYellow = "\x1b[33;1m"
var colorGreen = "\x1b[32;1m"
var colorNormal = "\x1b[0m"

const cliName = "anyops"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: cliName,
}

var workingDir string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(_workingDir string) {
	workingDir = _workingDir
	cobra.CheckErr(rootCmd.Execute())
}

func Finalize(_workingDir string) {
	_ = os.RemoveAll(_workingDir)
}

func init() {
	rootCmd.SilenceErrors = true
	rootCmd.SetUsageTemplate(colorYellow + `Usage:` + colorNormal + `{{if .Runnable}}
` + colorGreen + `{{.UseLine}}` + colorNormal + `{{end}}{{if .HasAvailableSubCommands}}
  ` + colorGreen + `{{.CommandPath}}` + colorNormal + ` [command]{{end}}{{if gt (len .Aliases) 0}}
` + colorYellow + `Aliases:` + colorNormal + `
  {{.NameAndAliases}}{{end}}{{if .HasExample}}
` + colorYellow + `Examples:` + colorNormal + `
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}
` + colorYellow + `Available Commands:` + colorNormal + `{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  ` + colorGreen + `{{rpad .Name .NamePadding }}` + colorNormal + ` {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}
` + colorYellow + `Flags:` + colorNormal + `
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}
` + colorYellow + `Global Flags:` + colorNormal + `
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}
Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`)
}

type CommandRequest struct {
	Args         []string
	ScriptDir    string
	Envs         []string
	DoNotOnError bool
	HideLogs     bool
}

func LogFatal(format string, v ...any) {
	log.Printf(format, v...)
	Finalize(workingDir)
	os.Exit(1)
}

func buildAnyCommands(command string, crs []CommandRequest) *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("%s", command),
		Short: fmt.Sprintf("To run %s", command),
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			for _, cr := range crs {
				ExecuteCommand(cr)
			}
		},
	}
}

func ExecuteCommand(cr CommandRequest) {
	command := strings.Join(cr.Args, " ")
	fmt.Printf(">>>>>>\nCommand:\n%s\n>>>>>>\nEnvs:\n%s\n>>>>>>\n\n", command, cr.Envs)
	cmd := exec.Command("bash", "-c", command)
	if cr.ScriptDir != "" {
		cmd.Dir = cr.ScriptDir
	} else {
		cmd.Dir = workingDir
	}
	if !cr.HideLogs {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, cr.Envs...)
	err := cmd.Run()
	if err != nil && !cr.DoNotOnError {
		LogFatal("cmd.Run() failed with %s\n", err)
	}
}
