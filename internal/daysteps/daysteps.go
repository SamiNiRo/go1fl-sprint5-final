package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("Неверный формат данных: ожидаются 2 значения")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("Не удалось определить количество шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("Количество шагов должно быть больше нуля")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("Неверный формат продолжительности: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("Продолжительность должна быть больше нуля")
	}
	ds.Steps, ds.Duration = steps, duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, spentenergy.Distance(ds.Steps, ds.Height), calories), nil
}
