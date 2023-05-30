package cmd

import "github.com/urfave/cli/v2"

// Build is a method
func (cmd *Command) Build() []*cli.Command {
	cmd.registerCLI(cmd.newDBMigrate())
	cmd.registerCLI(cmd.newDBSeed())

	cmd.registerCLI(cmd.GetListUser())
	cmd.registerCLI(cmd.GetDetailUser())
	cmd.registerCLI(cmd.CreateUser())
	cmd.registerCLI(cmd.UpdateUser())
	cmd.registerCLI(cmd.DeleteUser())

	return cmd.CLI
}
