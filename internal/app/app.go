package app

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "log/slog"
    "yourapp/internal/logging"
)

func Run(cmd *cobra.Command, args []string) {
    logging.Logger.Info("Запуск приложения")

    // Чтение переменной окружения
    apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
        logging.Logger.Error("API_KEY не установлен в файле конфигурации")
        os.Exit(1)
    }

    // Получение значения флага
    toggle, _ := cmd.Flags().GetBool("toggle")

    // Основная логика приложения
    logging.Logger.Info("Выполнение основной логики", slog.String("API_KEY", apiKey), slog.Bool("toggle", toggle))

    fmt.Println("API_KEY:", apiKey)
    fmt.Println("Toggle flag is:", toggle)
}
