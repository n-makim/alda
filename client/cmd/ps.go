package cmd

import (
	"fmt"
	"time"

	log "alda.io/client/logging"
	"alda.io/client/system"
	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List background processes",
	RunE: func(_ *cobra.Command, args []string) error {
		playerStates, err := system.ReadPlayerStates()
		if err != nil {
			return err
		}

		replServerStates, err := system.ReadREPLServerStates()
		if err != nil {
			return err
		}

		fmt.Println("id\tport\tstate\texpiry\ttype")

		for _, state := range playerStates {
			if state.ReadError != nil {
				log.Warn().Err(state.ReadError).Msg("Failed to read player state")
				continue
			}

			expiry := humanize.Time(time.Unix(state.Expiry/1000, 0))

			fmt.Printf(
				"%s\t%d\t%s\t%s\t%s\n",
				state.ID, state.Port, state.State, expiry, "player",
			)
		}

		for _, state := range replServerStates {
			if state.ReadError != nil {
				log.Warn().Err(state.ReadError).Msg("Failed to read REPL server state")
				continue
			}

			fmt.Printf(
				"%s\t%d\t%s\t%s\t%s\n",
				state.ID, state.Port, "-", "-", "repl-server",
			)
		}

		return nil
	},
}
