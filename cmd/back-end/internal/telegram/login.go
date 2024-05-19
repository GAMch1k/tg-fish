package telegram

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gamch1k.org/tg-fish/cmd/pkg/utils"

	pebbledb "github.com/cockroachdb/pebble"
	boltstor "github.com/gotd/contrib/bbolt"
	"github.com/gotd/contrib/middleware/floodwait"
	"github.com/gotd/contrib/middleware/ratelimit"
	"github.com/gotd/contrib/pebble"
	"github.com/gotd/contrib/storage"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/message/peer"
	"github.com/gotd/td/telegram/updates"
	"github.com/gotd/td/tg"
	"go.etcd.io/bbolt"
	"golang.org/x/time/rate"
)


func Login(phone string, ctx context.Context) {
	sessionDir := filepath.Join("sessions", utils.GetSessionFolder(phone))
	err := os.MkdirAll(sessionDir, 0700)
	utils.ErrorHandler(err)

	sessionStorage := &telegram.FileSessionStorage{
		Path: filepath.Join(sessionDir, "session.json"),
	}

	db, err := pebbledb.Open(filepath.Join(sessionDir, "peers.peeble.db"), &pebbledb.Options{})
	utils.ErrorHandler(err)

	peerDB := pebble.NewPeerStorage(db)

	dispatcher := tg.NewUpdateDispatcher()
	updateteHandler := storage.UpdateHook(dispatcher, peerDB)

	boltdb, err := bbolt.Open(filepath.Join(sessionDir, "updates.bolt.db"), 0666, nil)
	utils.ErrorHandler(err)

	updatesRecovery := updates.New(updates.Config{
		Handler: updateteHandler,
		Storage: boltstor.NewStateStorage(boltdb),
	})

	waiter := floodwait.NewWaiter().WithCallback(func(ctx context.Context, wait floodwait.FloodWait) {
		log.Println("Flood wait", wait.Duration)
	})

	options := telegram.Options{
		SessionStorage: sessionStorage,
		UpdateHandler: updatesRecovery,
		Middlewares: []telegram.Middleware{
			waiter,
			ratelimit.New(rate.Every(time.Millisecond * 100), 5),
		},
	}

	appId, _ := strconv.Atoi(os.Getenv("APP_ID"))
	client := telegram.NewClient(
		appId,
		os.Getenv("APP_HASH"),
		options,
	)

	api := client.API()

	resolver := storage.NewResolverCache(peer.Plain(api), peerDB)
	_ = resolver


	flow := auth.NewFlow(LoginData{PhoneNumber: phone}, auth.SendCodeOptions{})

	waiter.Run(ctx, func(ctx context.Context) error {

		// Continue auth here
		_ = flow
		return nil

	})


	log.Println("Finished login")
}