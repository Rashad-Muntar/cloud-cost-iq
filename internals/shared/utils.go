package shared

import ("time"
 "fmt")
func ParseDate(input string) (time.Time, error) {
    formats := []string{
        "02/01/2006",  // 15/04/2025
        "02-01-2006",  // 15-04-2025
        "2006-01-02",  // 2025-04-15 (standard)
        "2006/01/02",  // 2025/04/15
        "01/02/2006",  // 04/15/2025 (US format)
    }

    for _, format := range formats {
        parsed, err := time.Parse(format, input)
        if err == nil {
            return parsed, nil
        }
    }

    return time.Time{}, fmt.Errorf("invalid date format: %s — use DD/MM/YYYY or DD-MM-YYYY", input)
}

func EndOfDay(t time.Time) time.Time {
    return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}