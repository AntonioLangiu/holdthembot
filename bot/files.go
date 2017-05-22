package bot

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "log"
    "os"
    "io"
    "net/http"
)

func ResendVideo(update tgbotapi.Update, ctx *BotContext) {
    fileID := update.Message.VideoNote.FileID
    createFolder(videoFolder(ctx, fileID))

    getFile(ctx, fileID, videoPath(ctx, fileID))
    sendFile(update, ctx, videoPath(ctx, fileID))

    deleteFile(videoPath(ctx, fileID))
    deleteFile(videoFolder(ctx, fileID))
}

func ResendAudio(update tgbotapi.Update, ctx *BotContext) {
    fileID := update.Message.Voice.FileID
    createFolder(audioFolder(ctx, fileID))

    getFile(ctx, fileID, audioPath(ctx, fileID))
    sendFile(update, ctx, audioPath(ctx, fileID))

    out := "Telegram automatically detect .ogg files as audio messages. "
    out += "This doesn't allows you to download the file, and for this reason "
    out += "the file you received has a <b>.temp</b> suffix. To be able to play "
    out += "the file you need to change the extension from <b>.ogg.temp</b> to "
    out += "<b>.ogg</b>"
    SendText(update, ctx, out)

    deleteFile(audioPath(ctx, fileID))
    deleteFile(audioFolder(ctx, fileID))
}

func createFolder(folder string) {
    err := os.Mkdir(folder, 0777)
    if err != nil {
        log.Fatal(err)
    }
}

func getFile(ctx *BotContext, fileID string, filePath string) {
    file, err := ctx.Bot.GetFileDirectURL(fileID);
    if err != nil {
        log.Fatal(err)
    }
    output, err := os.Create(filePath)
    if err != nil {
        log.Print(err)
    }
    defer output.Close()
    response, err := http.Get(file)
    defer response.Body.Close()

    numBytesWritten, err := io.Copy(output, response.Body)
    if err != nil {
        log.Print(err)
    }
    if numBytesWritten < 0 {
        log.Fatal("numBytesWritten < 0")
    }
}

func sendFile(update tgbotapi.Update, ctx *BotContext, path string) {
    document := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, path)
    ctx.Bot.Send(document)
}

func deleteFile(path string) {
    err := os.Remove(path)
    if err != nil {
        log.Print("error removing file: "+path)
    }
}

func videoPath(ctx *BotContext, fileID string) string {
    return ctx.Config.TempFolder+"/"+fileID+"/video.mp4"
}

func videoFolder(ctx *BotContext, fileID string) string {
    return ctx.Config.TempFolder+"/"+fileID
}

func audioPath(ctx *BotContext, fileID string) string {
    return ctx.Config.TempFolder+"/"+fileID+"/audio.ogg.temp"
}

func audioFolder(ctx *BotContext, fileID string) string {
    return ctx.Config.TempFolder+"/"+fileID
}
