package client

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/GoSeoTaxi/dh8_msg/internal/config"
)

func StartWorker(ctx context.Context, cfg *config.Config) (errR error) {
	if cfg.CheckPing {
		err := pingUrlWithTimeOutAndCount(cfg.URLPing, 10, time.Duration(5*time.Second))
		if err != nil {
			log.Println(err)
			errR = errors.Join(errR, err)
		}
	}
	if cfg.CheckFile {
		err := checkFile(cfg.URLFile, cfg.Sha256)
		if err != nil {
			log.Println(err)
			errR = errors.Join(errR, err)
		}
	}

	if cfg.CheckLogin {
		err := checkLogin(cfg.URLLogin)
		if err != nil {
			log.Println(err)
			errR = errors.Join(errR, err)
		}
	}

	return errR
}
