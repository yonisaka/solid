package cmd

import "github.com/urfave/cli/v2"

// Build is a method
func (cmd *Command) Build() []*cli.Command {
	cmd.registerCLI(cmd.newDBMigrate())
	cmd.registerCLI(cmd.newDBSeed())

	cmd.registerCLI(cmd.getListUser())
	cmd.registerCLI(cmd.getDetailUser())
	cmd.registerCLI(cmd.createUser())
	cmd.registerCLI(cmd.updateUser())
	cmd.registerCLI(cmd.deleteUser())

	return cmd.CLI
}
