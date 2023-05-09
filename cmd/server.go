/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"AliGenieServer/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var (
	port      int
	host      string
	authKey   string
	authValue string
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var g errgroup.Group
		gin.SetMode(gin.ReleaseMode)

		r := gin.Default()
		router.Init(r, authKey, authValue)
		// 监听并在 0.0.0.0:8080 上启动服务
		g.Go(func() error {
			return r.Run(fmt.Sprintf("%s:%d", host, port))
		})

		if err := g.Wait(); err != nil {
			logrus.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().IntVarP(&port, "port", "p", 8085, "服务监听的端口")
	serverCmd.Flags().StringVarP(&host, "server", "s", "127.0.0.1", "服务监听的地址")
	serverCmd.Flags().StringVarP(&authKey, "authKey", "k", "authKey", "天猫精灵认证文件的文件名，如：xxxxx.txt")
	serverCmd.Flags().StringVarP(&authValue, "authValue", "v", "authValue", "天猫精灵认证文件的内容")
}
