package cli

import (
	"encoding/json"
	"fmt"

	"github.com/leg100/otf/internal/resource"
	"github.com/leg100/otf/internal/workspace"
	"github.com/spf13/cobra"
)

func (a *CLI) workspaceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workspaces",
		Short: "Workspace management",
	}

	cmd.AddCommand(a.workspaceListCommand())
	cmd.AddCommand(a.workspaceShowCommand())
	cmd.AddCommand(a.workspaceEditCommand())
	cmd.AddCommand(a.workspaceLockCommand())
	cmd.AddCommand(a.workspaceUnlockCommand())

	return cmd
}

func (a *CLI) workspaceListCommand() *cobra.Command {
	var org string

	cmd := &cobra.Command{
		Use:           "list",
		Short:         "List workspaces",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			list, err := resource.ListAll(func(opts resource.PageOptions) (*resource.Page[*workspace.Workspace], error) {
				return a.ListWorkspaces(cmd.Context(), workspace.ListOptions{
					PageOptions:  opts,
					Organization: &org,
				})
			})
			if err != nil {
				return fmt.Errorf("retrieving existing workspaces: %w", err)
			}
			for _, ws := range list {
				fmt.Fprintln(cmd.OutOrStdout(), ws.Name)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&org, "organization", "", "Organization workspace belongs to")
	cmd.MarkFlagRequired("organization")

	return cmd
}

func (a *CLI) workspaceShowCommand() *cobra.Command {
	var organization string

	cmd := &cobra.Command{
		Use:           "show [name]",
		Short:         "Show a workspace",
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			workspace := args[0]

			ws, err := a.GetWorkspaceByName(cmd.Context(), organization, workspace)
			if err != nil {
				return err
			}
			out, err := json.MarshalIndent(ws, "", "    ")
			if err != nil {
				return err
			}

			fmt.Fprintln(cmd.OutOrStdout(), string(out))

			return nil
		},
	}

	cmd.Flags().StringVar(&organization, "organization", "", "Organization workspace belongs to")
	cmd.MarkFlagRequired("organization")

	return cmd
}

func (a *CLI) workspaceEditCommand() *cobra.Command {
	var (
		organization string
		opts         workspace.UpdateOptions
		mode         *string
	)

	cmd := &cobra.Command{
		Use:           "edit [name]",
		Short:         "Edit a workspace",
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			if mode != nil && *mode != "" {
				opts.ExecutionMode = (*workspace.ExecutionMode)(mode)
			}

			ws, err := a.GetWorkspaceByName(cmd.Context(), organization, name)
			if err != nil {
				return err
			}
			ws, err = a.UpdateWorkspace(cmd.Context(), ws.ID, opts)
			if err != nil {
				return err
			}

			if opts.ExecutionMode != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "updated execution mode: %s\n", ws.ExecutionMode)
			}

			return nil
		},
	}

	mode = cmd.Flags().StringP("execution-mode", "m", "", "Which execution mode to use. Valid values are remote, local, and agent")

	cmd.Flags().StringVar(&organization, "organization", "", "Organization workspace belongs to")
	cmd.MarkFlagRequired("organization")

	return cmd
}

func (a *CLI) workspaceLockCommand() *cobra.Command {
	var organization string

	cmd := &cobra.Command{
		Use:           "lock [name]",
		Short:         "Lock a workspace",
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			workspace := args[0]

			ws, err := a.GetWorkspaceByName(cmd.Context(), organization, workspace)
			if err != nil {
				return err
			}
			ws, err = a.LockWorkspace(cmd.Context(), ws.ID, nil)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Successfully locked workspace %s\n", ws.Name)

			return nil
		},
	}

	cmd.Flags().StringVar(&organization, "organization", "", "Organization workspace belongs to")
	cmd.MarkFlagRequired("organization")

	return cmd
}

func (a *CLI) workspaceUnlockCommand() *cobra.Command {
	var (
		organization string
		force        bool
	)

	cmd := &cobra.Command{
		Use:           "unlock [name]",
		Short:         "Unlock a workspace",
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			workspace := args[0]

			ws, err := a.GetWorkspaceByName(cmd.Context(), organization, workspace)
			if err != nil {
				return err
			}
			ws, err = a.UnlockWorkspace(cmd.Context(), ws.ID, nil, force)
			if err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Successfully unlocked workspace %s\n", ws.Name)

			return nil
		},
	}

	cmd.Flags().StringVar(&organization, "organization", "", "Organization workspace belongs to")
	cmd.Flags().BoolVar(&force, "force", false, "Forceably unlock workspace.")
	cmd.MarkFlagRequired("organization")

	return cmd
}
