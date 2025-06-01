package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

// извиняюсь немножко, я случайно сначала коммиттил в main
const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	result := strings.Split(data, ",")
	var steps int
	var timeWalk time.Duration
	var err error
	if len(result) == 2 {
		steps, err = strconv.Atoi(result[0])
		if err != nil {
			return 0, 0, fmt.Errorf("ошибка парсинга шагов: %v", err)
		}
		if steps <= 0 {
			log.Printf("отрицательные шаги: %v", err)
			return 0, 0, fmt.Errorf("сделайте пожалуйста шаги: %v", err)
		}
		timeWalk, err = time.ParseDuration(result[1])
		if err != nil {

			return 0, 0, fmt.Errorf("ошибка парсинга времени: %v", err)
		}
		if timeWalk.Hours()/60+timeWalk.Minutes() <= 0 {
			log.Printf("отрицательныое время: %v", err)
			return 0, 0, fmt.Errorf("Встань и иди")
		}
		return steps, timeWalk, nil
	} else {
		log.Printf("пустая строка: %v", err)
		return 0, 0, fmt.Errorf("")
	}
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if duration.Hours()/60+duration.Minutes() <= 0 && steps <= 0 {
		return fmt.Sprintf("ввод: %d,%s, вес: %.2f, рост: %.2f", steps, duration, weight, height)
	} else {
		distance := (float64(steps) * stepLength) / mInKm
		calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return fmt.Sprintf("")
		}
		return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)

	}
}
