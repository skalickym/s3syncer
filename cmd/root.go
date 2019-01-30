package cmd

import (
	"fmt"
	"github.com/WanderaOrg/s3syncer/pkg/sync"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var loglevel, folderToWatch, s3Path, s3Region string
var rootCmd = &cobra.Command{
	Use:               "s3syncer",
	DisableAutoGenTag: true,
	Short:             "Syncing folder to S3",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		lvl, err := log.ParseLevel(loglevel)
		if err != nil {
			return err
		}

		log.SetLevel(lvl)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return sync.RunSync(folderToWatch, s3Region, s3Path)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&loglevel, "log-level", "l", "info", fmt.Sprintf("command log level (options: %s)", log.AllLevels))
	rootCmd.Flags().StringVarP(&folderToWatch, "folder", "f", "", "folder to watch")
	rootCmd.Flags().StringVarP(&s3Path, "s3-path", "p", "", "S3 path (s3://<bucket name>/<path>)")
	rootCmd.Flags().StringVarP(&s3Region, "s3-region", "r", "eu-west-1", "S3 region")

	_ = rootCmd.MarkFlagRequired("folder")
	_ = rootCmd.MarkFlagRequired("s3-path")
}

//Execute run root command (main entrypoint)
func Execute() error {
	return rootCmd.Execute()
}