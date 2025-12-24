package cmd

import (
	"Blog/pkg/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  `Application will be served on host and port defined in config.yml file`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}


func serve() {
	config.Set()

	configs:=config.Get()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.GetString("App.Name"),
			"version":  "1.0.0",
		})
	})

	r.Run(fmt.Sprintf("%s:%d", configs.Server.Host, configs.Server.Port))
}
