package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "log/slog"
    "yourapp/internal/logging"
)

var helloCmd = &cobra.Command{
    Use:   "hello",
    Short: "Приветственная команда",
    Run: func(cmd *cobra.Command, args []string) {
        name, _ := cmd.Flags().GetString("name")
        logging.Logger.Info("Выполнение команды hello", slog.String("name", name))
        fmt.Printf("Привет, %s!\n", name)
    },
}

func init() {
    rootCmd.AddCommand(helloCmd)
    helloCmd.Flags().StringP("name", "n", "Мир", "Имя для приветствия")
}
