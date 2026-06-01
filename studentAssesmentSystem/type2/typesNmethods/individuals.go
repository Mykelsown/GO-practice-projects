package type2

import (
	"fmt"
	"strings"
)

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
	bar, percent := progressBar(std.Scores)
	return fmt.Sprintf("[ %s ] %s · %s · %s %.1f%", std.EnrollmentID, std.Name, std.Course, bar, percent)
}

func (TA TeachingAssistant) String() string {
	for i:= 0; i>TA.SessionLed; i++{
		
	}

	bar, percent := progressBar(TA.Scores)
	return fmt.Sprintf("[ %s ] %s · %s · %s %.1f% [+%d sessions]", TA.EnrollmentID, TA.Name, TA.Course, bar, percent, TA.SessionLed)
}

func progressBar(score []float64) (string, float64) {
	bar := "░░░░░░░░░░"
	numScore := len(score) - 1
	totalScore := 0.0
	for _, scr := range score {
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
