package tgbot

import (
	"encoding/json"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
	"os"
)

const filePath = "./data/chat_numbers.json"

type TGBot struct {
	Bot         *telego.Bot
	chatNumbers []int64
	key         string
}

func NewBot(token, key string) (*TGBot, error) {
	bot, err := telego.NewBot(token, telego.WithDefaultLogger(false, true))
	if err != nil {
		return nil, err
	}

	tgBot := TGBot{
		Bot: bot,
		key: key,
	}

	err = tgBot.ReadFromFile()
	if err != nil {
		fmt.Println(err)
	}

	return &tgBot, nil
}

func (b *TGBot) ReadFromFile() error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("помилка відкриття файлу: %s", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&b.chatNumbers); err != nil {
		return fmt.Errorf("помилка декодування JSON: %s", err)
	}
	return nil
}

func (b *TGBot) WriteToFile() error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("помилка відкриття файлу: %s", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(b.chatNumbers); err != nil {
		return fmt.Errorf("помилка кодування в JSON: %s", err)
	}

	return nil
}

func (b *TGBot) ChatNumberExists(chatNumber int64) bool {
	for _, n := range b.chatNumbers {
		if n == chatNumber {
			return true
		}
	}
	return false
}

func (b *TGBot) ListenTelegramForKey() {
	updates, err := b.Bot.UpdatesViaLongPolling(nil)
	if err != nil {
		log.Println(err)
	}

	for update := range updates {
		if update.Message.Text == b.key && !b.ChatNumberExists(update.Message.Chat.ID) {
			b.chatNumbers = append(b.chatNumbers, update.Message.Chat.ID)
			err = b.WriteToFile()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (b *TGBot) SendMessage(text string) {
	for _, chatID := range b.chatNumbers {
		msg := telego.SendMessageParams{
			ChatID:    tu.ID(chatID),
			Text:      text,
			ParseMode: "HTML",
		}

		_, err := b.Bot.SendMessage(&msg)
		if err != nil {
			log.Println(err)
		}
	}
}
