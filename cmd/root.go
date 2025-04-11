package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"main/models"
	"main/models/dns"
	"main/resolver"
	"os"
)

var recordType dns.RecordType
var verbose bool

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: Domain is required")
			os.Exit(1)
		}
		if !verbose {
			log.SetLevel(log.PanicLevel)
		}
		domain, err := models.ParseDomain(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		log.Info("Query: ", recordType.String(), ".", domain)
		dnsQuery := dns.Query{RecordType: recordType, Domain: *domain}
		err = resolver.Run(&dnsQuery, dns.DefaultServer())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().VarP(&recordType, "type", "t", "Type of record")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	err := rootCmd.MarkFlagRequired("type")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
