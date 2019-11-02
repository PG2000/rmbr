/*
Copyright © 2019 Patrick Geschinski

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"strings"
)

type createCmd struct {
	command     string
	description string
	group       string
	manager     *RmbrNotesManager
}

func (cmd createCmd) run() error {
	cmd.manager.Repository.New(RmbrNote{
		Command:     cmd.command,
		Group:       cmd.group,
		Description: cmd.description,
	})
	return nil
}

func CreateCommand(m *RmbrNotesManager) *cobra.Command {
	c := &createCmd{
		command:     "",
		description: "",
		group:       "",
		manager:     m,
	}

	// cmd represents the create command
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "creates a new note to remember",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			m.Logger.InfoCommandMessage("create", strings.Join(args, ","))

			s, _ := cmd.Flags().GetString("command")
			if len(s) < 1 {
				return errors.New("the flag --command is not allowed to be empty")
			}
			return c.run()
		},
	}

	cmd.Flags().StringVarP(&c.command, "command", "c", "", "")
	cmd.Flags().StringVarP(&c.description, "description", "d", "", "")
	cmd.Flags().StringVarP(&c.group, "group", "g", "", "")

	cmd.MarkFlagRequired("command")
	cmd.MarkFlagRequired("description")

	return cmd
}
