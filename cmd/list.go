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
	"fmt"
	. "github.com/ahmetb/go-linq"
	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"strings"
)

type listCmd struct {
	manager *RmbrNotesManager
	query   string
}

func (cmd listCmd) run() error {

	notes, _ := cmd.manager.Repository.Load()

	r := From(notes.Notes).WhereT(func(note RmbrNote) bool {
		if len(cmd.query) > 0 {
			return strings.Contains(strings.ToLower(note.Description), strings.ToLower(cmd.query))
		}

		return true
	}).GroupByT(func(group RmbrNote) interface{} {
		return group.Group
	}, func(group RmbrNote) RmbrNote {
		return group
	}).Results()

	fmt.Printf("%s\n", Bold(Magenta("rmbr")))
	for _, v := range r {
		fmt.Printf("|__%s\n", Yellow(v.(Group).Key))
		for _, b := range v.(Group).Group {
			fmt.Printf("   |  %s\n", Bold(b.(RmbrNote).Command))
			fmt.Printf("   |  Description: \t%s\n", b.(RmbrNote).Description)
			fmt.Printf("   |--\n")
		}
	}
	return nil
}

func ListCommand(m *RmbrNotesManager) *cobra.Command {
	c := &listCmd{
		manager: m,
	}

	// cmd represents the create command
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "creates a new note to remember",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			m.Logger.Debug("create", strings.Join(args, ","))
			return c.run()
		},
	}

	cmd.Flags().StringVarP(&c.query, "query", "q", "", "")

	return cmd
}
