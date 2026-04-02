package keybindings

import (
	"log"

	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/config"
)

func GetKey(label string) gocui.Key {
	key, ok := config.KeyFromLabel(label)
	if !ok {
		log.Fatalf("Unrecognized key %s, this should have been caught by user config validation", label)
	}

	return key
}
