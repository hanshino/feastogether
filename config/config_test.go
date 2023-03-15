package config_test

import (
	"feastogether/config"
	"fmt"
	"log"
	"testing"
)

func TestGetConfig(t *testing.T) {

	if cfg, err := config.GetConfig(".."); err != nil {
		log.Println(err)
	} else {
		fmt.Println(cfg)
	}
}
