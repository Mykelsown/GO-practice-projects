package MTR

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

func ProcessLog(filename string) LogResult {
	time.Sleep(500)

	if strings.Contains(filename, "error") {
		errMsg := "failed to parse " + filename + ": critical error detected"
		return LogResult{
			FileName:  filename,
			LineCount: 0,
			HasError:  true,
			ErrorMsg:  errMsg,
		}
	}

	count := math.Max(100, float64(rand.Intn(1000)))
	return LogResult{
		FileName:  filename,
		LineCount: int(count),
		HasError:  false,
		ErrorMsg:  "",
	}
}
