package observer

import (
	"context"
	"github.com/ledgerwatch/erigon/cmd/utils"
	"github.com/ledgerwatch/erigon/internal/debug"
	"github.com/spf13/cobra"
)

type CommandFlags struct {
	DataDir     string
	Chain       string
	ListenPort  int
	NatDesc     string
	NetRestrict string
}

type Command struct {
	command cobra.Command
	flags CommandFlags
}

func NewCommand() *Command {
	command := cobra.Command{
		Use:     "",
		Short:   "P2P network crawler",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	// default flags
	utils.CobraFlags(&command, append(debug.Flags, utils.MetricFlags...))

	instance := Command{
		command: command,
	}
	instance.withDatadir()
	instance.withChain()
	instance.withListenPort()
	instance.withNAT()
	instance.withNetRestrict()

	return &instance
}

func (command *Command) withDatadir() {
	flag := utils.DataDirFlag
	command.command.Flags().StringVar(&command.flags.DataDir, flag.Name, flag.Value.String(), flag.Usage)
	must(command.command.MarkFlagDirname(utils.DataDirFlag.Name))
}

func (command *Command) withChain() {
	flag := utils.ChainFlag
	command.command.Flags().StringVar(&command.flags.Chain, flag.Name, flag.Value, flag.Usage)
}

func (command *Command) withListenPort() {
	flag := utils.ListenPortFlag
	command.command.Flags().IntVar(&command.flags.ListenPort, flag.Name, flag.Value, flag.Usage)
}

func (command *Command) withNAT() {
	flag := utils.NATFlag
	command.command.Flags().StringVar(&command.flags.NatDesc, flag.Name, flag.Value, flag.Usage)
}

func (command *Command) withNetRestrict() {
	flag := utils.NetrestrictFlag
	command.command.Flags().StringVar(&command.flags.NetRestrict, flag.Name, flag.Value, flag.Usage)
}

func (command *Command) ExecuteContext(ctx context.Context, runFunc func(ctx context.Context, flags CommandFlags) error) error {
	command.command.RunE = func(cmd *cobra.Command, _ []string) error {
		return runFunc(cmd.Context(), command.flags)
	}
	return command.command.ExecuteContext(ctx)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
