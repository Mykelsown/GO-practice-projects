package type2

import (
	"fmt"
	"strings"
)

type Learner interface {
	Progress() (string, float64)
	Grade() string
}

type Student struct {
	BaseProfile
	Scores []float64
	Course string
}

type TeachingAssistant struct {
	BaseProfile
	Scores     []float64
	Course     string
	SessionLed int
}

// Formatting types outputs
func (std Student) String() string {
	bar, percent := std.Progress()
	return fmt.Sprintf("[ %s ] %s · %s · %s %.1f%", std.EnrollmentID, std.Name, std.Course, bar, percent)
}

func (TA TeachingAssistant) String() string {
	bar, percent := TA.Progress()
	return fmt.Sprintf("[ %s ] %s · %s · %s %.1f% [+%d sessions]", TA.EnrollmentID, TA.Name, TA.Course, bar, percent, TA.SessionLed)
}

// Exposing Student and TeachingAAssistant to Learner
func (std Student) Progress() (string, float64) {
	bar := "░░░░░░░░░░"
	numScore := len(std.Scores)
	totalScore := 0.0
	for _, scr := range std.Scores {
		if scr > 100.00 {
			scr = 100
		}
		totalScore += scr
	}

	scorePercent := totalScore / float64(numScore)
	repeatBarCount := int(scorePercent) / 10
	if scorePercent < 10 {
		repeatBarCount = 1
	}
	bar = strings.Replace(bar, "░", "▓", repeatBarCount)

	return bar, scorePercent
}

func (TA TeachingAssistant) Progress() (string, float64) {
	bonusScore := 0
	for i := TA.SessionLed; i > 0; i-- {
		bonusScore += 5
		if bonusScore == 100 {
			break
		}
	}

	TA.Scores = append(TA.Scores, float64(bonusScore))

	bar := "░░░░░░░░░░"
	numScore := len(TA.Scores)
	totalScore := 0.0
	for _, scr := range TA.Scores {
		if scr > 100.00 {
			scr = 100
		}
		totalScore += scr
	}

	scorePercent := totalScore / float64(numScore)
	repeatBarCount := int(scorePercent) / 10
	if scorePercent < 10 {
		repeatBarCount = 1
	}
	bar = strings.Replace(bar, "░", "▓", repeatBarCount)

	return bar, scorePercent
}

func (std Student) Grade() string {
	_, percent := std.Progress()
	badge := ""
	switch {
	case percent >= 85.0:
		badge = "🏅 (A)"
	case percent >= 70.0:
		badge = "⭐ (B)"
	case percent >= 55.0:
		badge = "✅ (C)"
	case percent >= 40.0:
		badge = "⚠️ (D)"
	case percent < 40.0:
		badge = "❌ (F)"
	default:
		badge = "!!!"
	}

	return badge
}

func (TA TeachingAssistant) Grade() string {
	_, percent := TA.Progress()
	badge := ""
	switch {
	case percent >= 85.0:
		badge = "🏅"
	case percent >= 70.0:
		badge = "⭐"
	case percent >= 55.0:
		badge = "✅"
	case percent >= 40.0:
		badge = "⚠️"
	case percent < 40.0:
		badge = "❌"
	default:
		badge = "???"
	}

	return badge
}