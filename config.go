package sebarcore

import "github.com/eaciit/config"
import "os"
import "fmt"

var (
	configLoaded = false
)

func isValid() {
	if !configLoaded {
		fmt.Println("Config is not yet loaded")
		os.Exit(100)
	}
}

func SetConfig(path string) error {
	if err := config.SetConfigFile(path); err != nil {
		configLoaded = false
		return err
	}
	configLoaded = true
	return nil
}
