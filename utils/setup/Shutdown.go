package setup

import (
	"os"
	"os/signal"
	"pyncz/go-rest/models"

	"github.com/gofiber/fiber/v2"
)

func Shutdown(app *fiber.App, env *models.AppEnv) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	go func() {
		sig := <-signalChannel
		env.Log.Printf("Received %s signal, graceful shutdown...", sig)
		_ = app.Shutdown()
	}()
}
