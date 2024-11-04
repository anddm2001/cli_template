package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "yourapp/internal/app"
    "yourapp/internal/config"
    "yourapp/internal/logging"
)

var rootCmd = &cobra.Command{
    Use:   "yourapp",
    Short: "Описание вашего приложения",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        // Инициализируем логирование
        logging.InitLogger()

        // Загружаем конфигурацию
        configPath, _ := cmd.Flags().GetString("config")
        config.LoadConfig(configPath)
    },
    Run: func(cmd *cobra.Command, args []string) {
        app.Run(cmd, args)
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.PersistentFlags().StringP("config", "c", "", "Путь к файлу конфигурации (по умолчанию .env)")
    rootCmd.Flags().BoolP("toggle", "t", false, "Пример булевого флага")
}
