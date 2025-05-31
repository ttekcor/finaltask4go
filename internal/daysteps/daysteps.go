package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

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
		timeWalk, err = time.ParseDuration(result[1])
		if err != nil {
			return 0, 0, fmt.Errorf("ошибка парсинга времени: %v", err)
		}
		return steps, timeWalk, nil
	} else {
		return 0, 0, errors.New("")
	}
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	
	steps, duration, err := parsePackage(data)
	if err != nil && steps == 0{
		return fmt.Sprintf("⚠ Ошибка данных: %v", err)
	} else {
		distance := (float64(steps) * stepLength) / mInKm
		calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
		return fmt.Sprintf("⚠ Ошибка подсчёта калорий: %v", err)
	}
		return fmt.Sprintf("Количество шагов: %d\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, calories)
	
}}