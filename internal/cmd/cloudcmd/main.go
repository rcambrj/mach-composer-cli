package cloudcmd

import (
	"fmt"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/mach-composer/mach-composer-cli/internal/cloud"
)

var CloudCmd = &cobra.Command{
	Use:   "cloud",
	Short: "Manage your Mach Composer Cloud",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var cloudLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to mach composer cloud",
	RunE: func(cmd *cobra.Command, args []string) error {

		if err := cloud.Login(cmd.Context()); err != nil {
			return err
		}
		cmd.Println("Successfully authenticated to mach composer cloud")
		return nil
	},
}

var cloudConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure mach composer cloud",
	RunE: func(cmd *cobra.Command, args []string) error {
		hasValue := false

		if cmd.Flags().Changed("set-organization") {
			viper.Set("organization", MustGetString(cmd, "set-organization"))
			hasValue = true
		}

		if cmd.Flags().Changed("set-project") {
			viper.Set("project", MustGetString(cmd, "set-project"))
			hasValue = true
		}

		if cmd.Flags().Changed("set-auth-url") {
			viper.Set("auth-url", MustGetString(cmd, "set-auth-url"))
			hasValue = true
		}

		if cmd.Flags().Changed("set-api-url") {
			viper.Set("api-url", MustGetString(cmd, "set-api-url"))
			hasValue = true
		}

		if hasValue {
			if err := viper.WriteConfig(); err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}
		}

		fmt.Println("Auth URL     : ", viper.GetString("auth-url"))
		fmt.Println("API URL      : ", viper.GetString("api-url"))
		fmt.Println("Organization : ", viper.GetString("organization"))
		fmt.Println("Project      : ", viper.GetString("project"))
		return nil
	},
}

func init() {
	// Config command
	CloudCmd.AddCommand(cloudConfigCmd)
	cloudConfigCmd.Flags().String("set-organization", "", "Set default organization")
	cloudConfigCmd.Flags().String("set-project", "", "Set default project")
	cloudConfigCmd.Flags().String("set-api-url", "", "Set api url")
	cloudConfigCmd.Flags().String("set-auth-url", "https://auth.mach.cloud", "Authentication URL")

	// Login
	CloudCmd.AddCommand(cloudLoginCmd)

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	configPath := path.Join(xdg.ConfigHome, "mach-composer")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Warning: encountered while creating configuration file: %s", err))
			return
		}
	}

	viper.SetConfigName("mach-composer")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("MCC")

	err := viper.SafeWriteConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); !ok {
			os.Stderr.WriteString(fmt.Sprintf("Warning: encountered while writing configuration file: %s", err))
			return
		}
	}

	viper.SetDefault("api-url", "https://api.mach.cloud")
	viper.SetDefault("auth-url", "https://auth.mach.cloud")

	if err := viper.ReadInConfig(); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Warning: invalid config file found at: %s\n", viper.GetViper().ConfigFileUsed()))
		return
	}

	// Copy the values from Viper to the matching flag values.
	// TODO: make recursive
	for _, cmd := range CloudCmd.Commands() {
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if cmd.Flags().Changed(f.Name) {
				return
			}

			if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
				Must(cmd.Flags().Set(f.Name, viper.GetString(f.Name)))
			}
		})
	}
}
