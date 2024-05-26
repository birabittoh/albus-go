package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

const tokenEnvName = "TELEGRAM_BOT_TOKEN"

func main() {
	godotenv.Load()

	tgToken := os.Getenv(tokenEnvName)
	if tgToken == "" {
		println("Please set the", tokenEnvName, "environment variable.")
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	bot.Debug = true

	println("Authorized on account", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ReplyParameters.MessageID = update.Message.MessageID

		messageDocument := update.Message.Document
		if messageDocument == nil {
			continue
		}

		msgFileNameFull := messageDocument.FileName
		msgFileId := messageDocument.FileID

		msgFile, err := bot.GetFile(tgbotapi.FileConfig{FileID: msgFileId})
		if err != nil {
			println("Error getting file:", err)
			continue
		}

		msgFileName, msgFileExtension := SplitFilename(msgFileNameFull)
		uniqueFileName := msgFileId + msgFileExtension

		err = DownloadFile(uniqueFileName, msgFile.Link(bot.Token))
		if err != nil {
			println("Error downloading file:", err)
			continue
		}

		outFileName, outExtension := ConvertFile(uniqueFileName)
		os.Remove(uniqueFileName)
		if outFileName == "" && outExtension == "" {
			println("Error converting file:", uniqueFileName)
			continue
		}
		outFileNameFull := JoinFilename(outFileName, outExtension)
		finalFileName := JoinFilename(msgFileName, outExtension)

		file, err := os.Open(outFileNameFull)
		if err != nil {
			println("Error opening new file: ", err)
		}

		fileReader := tgbotapi.FileReader{
			Name:   finalFileName,
			Reader: file,
		}

		newMsg := tgbotapi.NewDocument(update.Message.Chat.ID, fileReader)

		if _, err := bot.Send(newMsg); err != nil {
			println("Could not send message.")
			os.Exit(1)
		}

		file.Close()
		os.Remove(outFileNameFull)
	}
}
