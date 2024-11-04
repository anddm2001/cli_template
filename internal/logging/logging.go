package logging

import (
    "log"
    "os"
    "path/filepath"

    "log/slog"
)

var Logger *slog.Logger

func InitLogger() {
    // Получаем путь к директории исполняемого файла
    exePath, err := os.Executable()
    if err != nil {
        log.Fatalf("Ошибка получения пути к исполняемому файлу: %v", err)
    }
    exeDir := filepath.Dir(exePath)

    logFilePath := filepath.Join(exeDir, "yourapp.log")

    logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Ошибка при создании файла лога: %v", err)
    }

    // Создаем новый обработчик, который записывает в файл
    handler := slog.NewTextHandler(logFile, nil)

    // Инициализируем глобальный логгер
    Logger = slog.New(handler)
}
