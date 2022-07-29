/*
Copyright 2022 Arda Güçlü.

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
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/cmd/apply"
	"k8s.io/kubectl/pkg/cmd/diff"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/utils/exec"
)

const (
	example = `
 		# Apply the configuration in pod.json to a pod
		kubectl dapply -f ./pod.json
`
	longDesc = `
Apply a configuration to a resource by file name or stdin only after running the
diff command for given resource(s) and user accepts to proceed.
`
)

type DiffApplyFlags struct {
	ConfigFlags  *genericclioptions.ConfigFlags
	ApplyFlags   *apply.ApplyFlags
	ApplyOptions *apply.ApplyOptions
	DiffOptions  *diff.DiffOptions
	genericclioptions.IOStreams
}

func NewDiffApplyFlags(streams genericclioptions.IOStreams) *DiffApplyFlags {
	return &DiffApplyFlags{
		ConfigFlags: genericclioptions.NewConfigFlags(true),
		ApplyFlags:  apply.NewApplyFlags(cmdutil.NewFactory(&genericclioptions.ConfigFlags{}), streams),
		DiffOptions: diff.NewDiffOptions(streams),
		IOStreams:   genericclioptions.IOStreams{},
	}
}

// NewCmdDiffApply provides a cobra command wrapping DiffApplyOptions
func NewCmdDiffApply(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewDiffApplyFlags(streams)

	cmd := &cobra.Command{
		Use:          "dapply [flags]",
		Short:        "show diff and ask user to proceed to apply",
		Long:         longDesc,
		SilenceUsage: true,
		Example:      example,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(c, args); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	o.ApplyFlags.AddFlags(cmd)
	o.ConfigFlags.AddFlags(cmd.Flags())

	return cmd
}

func (o *DiffApplyFlags) Complete(c *cobra.Command, args []string) error {
	o.ApplyFlags.Factory = cmdutil.NewFactory(o.ConfigFlags)
	applyOptions, err := o.ApplyFlags.ToOptions(c, "kubectl", args)
	if err != nil {
		return err
	}
	o.ApplyOptions = applyOptions

	if o.ApplyOptions.Prune {
		return errors.New("prune is not supported")
	}

	o.DiffOptions.FilenameOptions = o.ApplyOptions.DeleteOptions.FilenameOptions
	o.DiffOptions.ServerSideApply = o.ApplyOptions.ServerSideApply
	o.DiffOptions.FieldManager = o.ApplyOptions.FieldManager
	o.DiffOptions.ForceConflicts = o.ApplyOptions.ForceConflicts
	o.DiffOptions.Selector = o.ApplyOptions.Selector
	err = o.DiffOptions.Complete(o.ApplyFlags.Factory, c)
	if err != nil {
		return err
	}

	return nil
}

// Validate ensures that all required arguments and flag values are provided
func (o *DiffApplyFlags) Validate(c *cobra.Command, args []string) error {
	err := o.ApplyOptions.Validate(c, args)
	if err != nil {
		return err
	}

	return nil
}

// Run shows differences and applies if user types yes
func (o *DiffApplyFlags) Run() error {
	err := o.DiffOptions.Run()
	if err == nil {
		fmt.Println("no changes found")
		return nil
	} else if err, ok := err.(exec.ExitError); ok {
		if err.ExitStatus() > 1 {
			return err
		}
	} else {
		return err
	}

	proceed := false
	prompt := &survey.Confirm{
		Message: "Do you want to proceed to apply?",
	}
	err = survey.AskOne(prompt, &proceed)
	if err != nil {
		return err
	}

	if !proceed {
		return nil
	}

	err = o.ApplyOptions.Run()
	if err != nil {
		return err
	}
	return nil
}
