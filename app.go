package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	ps "github.com/mitchellh/go-ps"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
)

// Создаем GaugeVec для хранения состояния процессов с меткой process_name
var (
	processesExists = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "process_exists",
			Help: "Process status: 0 - not running, 1 - running",
		},
		[]string{"process_name"},
	)
)

// Инициализация и регистрация метрик
func init() {
	prometheus.MustRegister(processesExists)
}

// checkProcesses обходит список процессов и проверяет, запущены ли указанные процессы
func checkProcesses(processNames []string) {
	for {
		// Получаем список всех процессов
		processList, err := ps.Processes()
		if err != nil {
			fmt.Println("Error getting process list:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		// Для каждого указанного процесса проверяем, работает ли он
		for _, processName := range processNames {
			found := false
			for _, p := range processList {
				if strings.Contains(p.Executable(), processName) {
					found = true
					break
				}
			}

			// Обновляем значение метрики с соответствующей меткой process_name
			if found {
				processesExists.WithLabelValues(processName).Set(1)
			} else {
				processesExists.WithLabelValues(processName).Set(0)
			}
		}
		// Задержка перед следующей проверкой состояния процессов
		time.Sleep(10 * time.Second)
	}
}

func main() {
	var processNames []string

	// Обработка аргументов командной строки
	pflag.StringArrayVarP(&processNames, "process", "p", []string{}, "Process names to monitor")
	pflag.Parse()

	// Если процессы для мониторинга не указаны, выводим сообщение об ошибке и завершаем работу
	if len(processNames) == 0 {
		fmt.Println("Usage: go run main.go --process <process_name> [--process <another_process_name>]")
		os.Exit(1)
	}

	// Запускаем в фоновом режиме проверку состояния указанных процессов
	go checkProcesses(processNames)

	// Регистрируем обработчик метрик и запускаем HTTP-сервер
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8081", nil)
}
