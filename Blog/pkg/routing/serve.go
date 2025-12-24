package routing

import (
	"Blog/pkg/config"
	"fmt"
	"log"
)

func Serve() {
	r := GetRouter()
	configs := config.Get()

	err := r.Run(fmt.Sprintf("%s:%d", configs.Server.Host, configs.Server.Port))
	if err != nil {
		log.Fatal("Error running server:", err)
	}
}
