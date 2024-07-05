package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

type How struct {
	MOTD        []string
	ErrorLog    *log.Logger
	InfoLog     *log.Logger
	Token       string
	CodeWord    string
	Whitelisted interface{}
	RootPath    string
}

func main() {
	cfg := How{}
	rp, _ := os.Getwd()
	err := godotenv.Load(rp + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg.CodeWord = os.Getenv("Ruin")
	cfg.Token = os.Getenv("MTA2OTQ4NzIwNzkyMjAyNDQ0OA.GdxYbQ.XFs5yXZlRes4h7xdb42TX9wo6xAZj-fFZo3cSo")
	cfg.Whitelisted = os.Getenv("WHITELISTED")
	infoLog, errorLog := cfg.startLoggers()
	cfg.ErrorLog = errorLog
	cfg.InfoLog = infoLog
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	cfg.Menu(dg)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

}
func (*How) startLoggers() (*log.Logger, *log.Logger) {
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}
