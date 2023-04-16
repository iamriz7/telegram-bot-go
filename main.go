package main

import (
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/valyala/fasthttp"
)

func main() {
	// Inisialisasi bot dengan token
	bot, err := tgbotapi.NewBotAPI("TOKEN-LU")
	if err != nil {
		log.Panic(err)
	}
     
	// Log saat bot berhasil terhubung
	log.Printf("Bot %s berhasil terhubung", bot.Self.UserName)

	// Konfigurasi polling untuk mendapatkan update dari Telegram API
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Menerima channel update dari API
	updates, err := bot.GetUpdatesChan(u)

	// Looping untuk mendapatkan pesan dan mengirimkan balasan
	for update := range updates {
		if update.Message == nil { // Pesan yang diterima tidak berupa teks
			continue
		}

		switch {
		case strings.HasPrefix(update.Message.Text, "/start"):
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Halo! Selamat datang di bot ini.")
			_, err = bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
		case strings.HasPrefix(update.Message.Text, "/help"):
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ketik /start untuk memulai bot.")
			_, err = bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
		case strings.HasPrefix(update.Message.Text, "/ping"):
			start := time.Now()
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🏓 Pong!")
			bot.Send(msg)
			elapsed := time.Since(start)
			log.Printf("Waktu respon: %s", elapsed)
		case strings.HasPrefix(update.Message.Text, "/photo"):
            // Request gambar dari URL
            status, body, err := fasthttp.Get(nil, "https://picsum.photos/200/300")
            if err != nil {
                log.Println(err)
                continue
            }
            if status != fasthttp.StatusOK {
                log.Printf("HTTP status code: %d", status)
                continue
            }

            // Kirim gambar ke user
            photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, tgbotapi.FileBytes{Name: "photo.jpg", Bytes: body})
            _, err = bot.Send(photo)
            if err != nil {
                log.Println(err)
                continue
            }
		default:
			log.Printf("Pesan diterima: %s", update.Message.Text)
			if err != nil {
				log.Println(err)
			}
		}
	}

}
