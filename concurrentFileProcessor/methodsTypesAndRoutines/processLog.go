package MTR

import (
	"math/rand"
	"strings"
	"time"
)

var rndSeed = rand.New(rand.NewSource(int64(time.Now().Second())))

func ProcessLog(filename string) LogResult {
	time.Sleep(time.Duration(rand.Intn(400)+100) * time.Microsecond)

	if strings.Contains(filename, "error") {
		errMsg := "failed to parse " + filename + ": critical error detected"
		return LogResult{
			FileName:  filename,
			LineCount: 0,
			HasError:  true,
			ErrorMsg:  errMsg,
		}
	}

	count := rndSeed.Intn(900) + 100
	return LogResult{
		FileName:  filename,
		LineCount: count,
		HasError:  false,
		ErrorMsg:  "",
	}
}
