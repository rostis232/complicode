package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rostis232/complicode/internal/tgbot"
	"github.com/rostis232/complicode/view"
	"github.com/rostis232/complicode/view/layout"
	"math/rand"
)

type Handler struct {
	TGBot *tgbot.TGBot
}

func NewHandler(bot *tgbot.TGBot) *Handler {
	return &Handler{TGBot: bot}
}

func (h Handler) Home(c echo.Context) error {
	slogans := []string{
		"HERE WE CREATE CODE THAT SOLVES YOUR PROBLEMS",
		"IN CODE WE TRUST",
		"OUR CODE WILL SOLVE YOUR PROBLEM",
	}

	randNumber := rand.Intn(len(slogans))

	fmt.Println(randNumber)

	return render(c, view.Home(slogans[randNumber], texts, en))
}

func (h Handler) HomeUA(c echo.Context) error {
	slogans := []string{
		"ТУТ МИ СТВОРЮЄМО КОД, ЯКИЙ ВИРІШУЄ ВАШІ ПРОБЛЕМИ",
		"МИ ВІРИМО В КОД",
		"НАШ КОД ВИРІШИТЬ ВАШУ ПРОБЛЕМУ",
	}

	randNumber := rand.Intn(len(slogans))

	fmt.Println(randNumber)

	return render(c, view.Home(slogans[randNumber], texts, ua))
}

func (h Handler) Send(c echo.Context) error {
	name := c.FormValue("name")
	lang := c.FormValue("lang")
	if lang == "" {
		lang = en
	}
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	message := c.FormValue("message")
	text := fmt.Sprintf("<b> Hey, Boss! CompliCode recieved new message!</b>\n<b>Name:</b> %s\n<b>Email:</b> %s\n<b>Phone:</b> %s\n<b>Message:</b> %s", name, email, phone, message)
	h.TGBot.SendMessage(text)
	return render(c, layout.Contact(true, texts, lang))
}
