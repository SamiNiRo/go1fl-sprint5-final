package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	Duration     time.Duration
	TrainingType string
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("Неверный формат данных: ожидаются 3 значения")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("Не удалось определить количество шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("Количество шагов должно быть больше нуля")
	}
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("Неверный формат продолжительности: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("Продолжительность должна быть больше нуля")
	}
	t.Steps, t.Duration = steps, duration
	t.TrainingType = parts[1]
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps, t.Height), spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration), calories), nil
}
