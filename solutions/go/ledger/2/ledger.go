package ledger

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

/*
Refactoring log:
- Replaced the custom sorting loop with sort.Slice.
- Removed per-entry goroutines and the result channel.
- Separated validation, date formatting, description formatting, and amount formatting.
- Added clear errors for unsupported locales, currencies, and invalid dates.
- Used integer arithmetic for amounts, including safe handling of the minimum int.
- Built the output with strings.Builder.
*/
func FormatLedger(currency, locale string, entries []Entry) (string, error) {
	if locale != "en-US" && locale != "nl-NL" {
		return "", fmt.Errorf("unsupported locale: %q", locale)
	}
	if currency != "USD" && currency != "EUR" {
		return "", fmt.Errorf("unsupported currency: %q", currency)
	}

	sortedEntries := append([]Entry(nil), entries...)
	for _, entry := range sortedEntries {
		if _, err := time.Parse("2006-01-02", entry.Date); err != nil {
			return "", fmt.Errorf("invalid date %q: %w", entry.Date, err)
		}
	}

	sort.Slice(sortedEntries, func(i, j int) bool {
		a, b := sortedEntries[i], sortedEntries[j]
		if a.Date != b.Date {
			return a.Date < b.Date
		}
		if a.Description != b.Description {
			return a.Description < b.Description
		}
		return a.Change < b.Change
	})

	dateHeader, descriptionHeader, changeHeader := "Date", "Description", "Change"
	if locale == "nl-NL" {
		dateHeader, descriptionHeader, changeHeader = "Datum", "Omschrijving", "Verandering"
	}

	var output strings.Builder
	fmt.Fprintf(&output, "%-10s | %-25s | %-13s\n",
		dateHeader, descriptionHeader, changeHeader)

	for _, entry := range sortedEntries {
		date, _ := time.Parse("2006-01-02", entry.Date)
		dateLayout := "01/02/2006"
		if locale == "nl-NL" {
			dateLayout = "02-01-2006"
		}

		description := formatDescription(entry.Description)
		amount := formatAmount(entry.Change, currency, locale)

		fmt.Fprintf(&output, "%-10s | %-25s | %13s\n",
			date.Format(dateLayout), description, amount)
	}

	return output.String(), nil
}

func formatDescription(description string) string {
	runes := []rune(description)
	if len(runes) > 25 {
		return string(runes[:22]) + "..."
	}
	return description + strings.Repeat(" ", 25-len(runes))
}

func formatAmount(change int, currency, locale string) string {
	value := int64(change)
	negative := value < 0

	// Convert to an unsigned magnitude without overflowing for the minimum int.
	var magnitude uint64
	if negative {
		magnitude = uint64(-(value + 1)) + 1
	} else {
		magnitude = uint64(value)
	}

	major := strconv.FormatUint(magnitude/100, 10)
	minor := fmt.Sprintf("%02d", magnitude%100)

	symbol := "$"
	if currency == "EUR" {
		symbol = "€"
	}

	if locale == "nl-NL" {
		major = groupDigits(major, ".")
		sign := ""
		if negative {
			sign = "-"
		}
		return symbol + " " + sign + major + "," + minor + " "
	}

	major = groupDigits(major, ",")
	amount := symbol + major + "." + minor
	if negative {
		return "(" + amount + ")"
	}
	return amount + " "
}

func groupDigits(digits, separator string) string {
	if len(digits) <= 3 {
		return digits
	}

	firstGroup := len(digits) % 3
	if firstGroup == 0 {
		firstGroup = 3
	}

	var grouped strings.Builder
	grouped.WriteString(digits[:firstGroup])
	for i := firstGroup; i < len(digits); i += 3 {
		grouped.WriteString(separator)
		grouped.WriteString(digits[i : i+3])
	}
	return grouped.String()
}
