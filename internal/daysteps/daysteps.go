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

			return 0, 0, err
		}
		timeWalk, err = time.ParseDuration(result[1])
		if err != nil {
			return 0, 0, err
		}
		return steps, timeWalk, nil
	}
	return 0, 0, errors.New("")
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if steps != 0 && err == nil {
		distance := (float64(steps) * stepLength) / mInKm
		calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Количество шагов: %d\nДистанция составила %b\nВы сожгли %b", steps, distance, calories)
	} else if err != nil {
		return err.Error()
	} else {
		return ""
	}

}
