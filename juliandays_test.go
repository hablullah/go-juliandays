package juliandays_test

import (
	"math"
	"testing"
	"time"

	"github.com/hablullah/go-juliandays"
)

func Test_JD_FromTime(t *testing.T) {
	for _, entry := range testEntries {
		jd, _ := juliandays.FromTime(entry.Time)
		if math.Round((entry.JD-jd)*100_000) != 0 {
			t.Errorf("%s: want %f got %f\n",
				entry.Time.Format("2006-01-02 15:04:05"),
				entry.JD, jd)
		}
	}
}

func Test_JD_ToTime(t *testing.T) {
	for _, entry := range testEntries {
		if entry.JD == 0 {
			continue
		}

		dt := juliandays.ToTime(entry.JD)
		diff := entry.Time.Sub(dt).Seconds()
		if math.Abs(diff) > 1 {
			t.Errorf("%f: want %s got %s (%v)\n",
				entry.JD,
				entry.Time.Format("2006-01-02 15:04:05"),
				dt.Format("2006-01-02 15:04:05"),
				entry.Time.Sub(dt))
		}
	}
}

func Test_JD_Bidirectional(t *testing.T) {
	dt := time.Date(-4712, 1, 1, 12, 0, 0, 0, time.UTC)
	for dt.Year() <= 2120 {
		// Convert date time to Julian Days
		jd, err := juliandays.FromTime(dt)
		if err != nil {
			dt = dt.AddDate(0, 0, 1)
			continue
		}

		// Convert back Julian Days to date time
		newDt := juliandays.ToTime(jd)

		// Compare original and new date time
		diff := dt.Sub(newDt).Seconds()
		if math.Abs(diff) > 1 {
			t.Errorf("Original %s: JD %f, reverted into %s (%v)\n",
				dt.Format("2006-01-02 15:04:05"), jd,
				newDt.Format("2006-01-02 15:04:05"), diff)
		}

		// Increase date
		dt = dt.AddDate(0, 0, 1)
	}
}
