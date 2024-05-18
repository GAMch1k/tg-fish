package main

import (
	"os"
	"os/signal"
	"syscall"
	"log"
	"io"
	"strconv"

	back "gamch1k.org/tg-fish/cmd/back-end"
	front "gamch1k.org/tg-fish/cmd/front-end"
	"gamch1k.org/tg-fish/cmd/pkg/utils"

	"github.com/joho/godotenv"
)

func load_env() {
	err := godotenv.Load(".env")
	utils.ErrorHandler(err)
}

func setup_logger() *os.File {
    f, err := os.OpenFile("./logs/logs.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    utils.ErrorHandler(err)
    
    wrt := io.MultiWriter(os.Stdout, f)
    log.SetOutput(wrt)

    return f
}

func main() {
	load_env()
	
	log_to_f, _ := strconv.ParseBool(os.Getenv("LOG_TO_FILE"))

	var f *os.File 
	if log_to_f {
		log.Println("Logging to file ENABLED")
		f = setup_logger()
	} else {
		log.Println("Logging to file DISABLED")
	}


	go front.Start(os.Getenv("FRONT_PORT"))
	go back.Start(os.Getenv("BACK_PORT"))

	if log_to_f {
		defer f.Close()
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}