package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	result := strings.Split(data, ",")
	if len(result) == 3 {
		steps, err := strconv.Atoi(result[0])
		if err != nil {
			return 0, "", 0, err
		}
		typeWalk := result[1]
		duration, err := time.ParseDuration(result[2])
		if err != nil {
			return 0, "", 0, err
		}
		return steps, typeWalk, duration, nil

	}
	return 0, "", 0, errors.New("")
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	distanceStep := height * stepLengthCoefficient
	resultDistance := (distanceStep * float64(steps)) / mInKm
	return resultDistance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration > 0 {
		distanceRes := distance(steps, height)
		result := distanceRes / duration.Hours()
		return result
	} else {
		return 0
	}

}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, name, duration, err := parseTraining(data)
	if err != nil {
		return fmt.Sprintf("⚠ Ошибка данных: %v", err), nil
	}

	switch strings.TrimSpace(name) {
	case "Бег":
		caloriesRun, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return fmt.Sprintf("⚠ Ошибка подсчёта калорий: %v", err), nil
		}
		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч\nДистанция: %.2f км\nСкорость: %.2f км/ч\nСожгли калорий: %.2f ккал",
			name,
			duration.Hours(),
			distance(steps, height),
			meanSpeed(steps, height, duration),
			caloriesRun,
		), nil

	case "Ходьба":
		caloriesWalk, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return fmt.Sprintf("⚠ Ошибка подсчёта калорий: %v", err), nil
		}
		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч\nДистанция: %.2f км\nСкорость: %.2f км/ч\nСожгли калорий: %.2f ккал",
			name,
			duration.Hours(),
			distance(steps, height),
			meanSpeed(steps, height, duration),
			caloriesWalk,
		), nil

	default:
		return fmt.Sprintf("⚠ Неизвестный тип тренировки: %s", name), nil
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps == 0 || weight == 0 || height == 0 || duration == 0 {
		return 0, errors.New("недостаточно данных для подсчёта")
	} else {
		result := weight * meanSpeed(steps, height, duration) * duration.Minutes() / float64(minInH)
		return result, nil
	}
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	tmp, err := RunningSpentCalories(steps, weight, height, duration)
	if err == nil {
		return tmp * walkingCaloriesCoefficient, nil
	} else {
		return 0, err
	}
}
