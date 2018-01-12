package regolithe

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"

	"github.com/Sirupsen/logrus"

	"gopkg.in/src-d/go-git.v4/plumbing"

	"github.com/aporeto-inc/regolithe/spec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	git "gopkg.in/src-d/go-git.v4"
)

// NewCommand generates a new CLI for regolith
func NewCommand(
	name string,
	description string,
	nameConvertFunc spec.AttributeNameConverterFunc,
	typeConvertFunc spec.AttributeTypeConverterFunc,
	typeMappingName string,
	generatorFunc func(*spec.SpecificationSet) error,
) *cobra.Command {

	cobra.OnInitialize(func() {
		viper.SetEnvPrefix(name)
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	})

	var rootCmd = &cobra.Command{
		Use:   name,
		Short: description,
	}

	var cmdFolderGen = &cobra.Command{
		Use:   "folder",
		Short: "Generate the model using a local directory containing the specs.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			specSet, err := spec.NewSpecificationSet(
				viper.GetString("dir"),
				nameConvertFunc,
				typeConvertFunc,
				typeMappingName,
			)
			if err != nil {
				return err
			}

			return generatorFunc(specSet)
		},
	}
	cmdFolderGen.Flags().StringP("dir", "d", "", "Path of the specifications folder.")

	var githubGen = &cobra.Command{
		Use:   "github",
		Short: "Generate the model using a remote github repository.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			if viper.GetString("repo") == "" {
				return errors.New("--repo is required")
			}

			var auth transport.AuthMethod
			if viper.GetString("token") != "" {
				auth = &http.BasicAuth{
					Username: "Bearer",
					Password: viper.GetString("token"),
				}
			}

			tmpFolder, err := ioutil.TempDir("", "")
			if err != nil {
				return err
			}
			defer os.RemoveAll(tmpFolder) // nolint: errcheck

			var (
				ref           plumbing.ReferenceName
				needsCheckout bool
			)

			givenHash := plumbing.NewHash(viper.GetString("ref"))
			if !givenHash.IsZero() {
				ref = plumbing.NewReferenceFromStrings("refs/heads/master", "").Name()
				needsCheckout = true
			} else {
				ref = plumbing.NewReferenceFromStrings("refs/heads/"+viper.GetString("ref"), "").Name()
			}

			logrus.WithFields(logrus.Fields{
				"ref":  viper.GetString("ref"),
				"repo": viper.GetString("repo"),
				"path": viper.GetString("path"),
			}).Info("Retrieving repository")

			repo, err := git.PlainClone(
				tmpFolder,
				false,
				&git.CloneOptions{
					URL:           viper.GetString("repo"),
					Progress:      nil,
					ReferenceName: ref,
					Auth:          auth,
				})
			if err != nil {
				return err
			}

			if needsCheckout {
				wt, e := repo.Worktree()
				if e != nil {
					return e
				}

				if err = wt.Checkout(
					&git.CheckoutOptions{
						Hash: givenHash,
					}); err != nil {
					return err
				}
			}

			specSet, err := spec.NewSpecificationSet(
				path.Join(tmpFolder, viper.GetString("path")),
				nameConvertFunc,
				typeConvertFunc,
				typeMappingName,
			)
			if err != nil {
				return err
			}

			return generatorFunc(specSet)
		},
	}
	githubGen.Flags().StringP("repo", "r", "", "Endpoint for the github api.")
	githubGen.Flags().StringP("path", "p", "", "Internal path to a directory in the repo if not in the root.")
	githubGen.Flags().StringP("ref", "R", "master", "Branch or tag to use.")
	githubGen.Flags().StringP("token", "t", "", "The api token to use.")

	rootCmd.AddCommand(
		cmdFolderGen,
		githubGen,
	)

	return rootCmd
}
