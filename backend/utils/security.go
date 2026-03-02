package utils

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

// ValidError struct untuk standarisasi error response di Fiber
type ValidError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value,omitempty"`
}

// Regex patterns tetap sama (Pure Go)
var sqlInjectionPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)\b(union\s+select|insert\s+into|update\s+\w+\s+set|delete\s+from|drop\s+table|create\s+table|alter\s+table|exec\s*\(|execute\s*\()\b`),
	regexp.MustCompile(`(?i)(--|/\*|\*/|;)\s*(\w|$)`),
	regexp.MustCompile(`(?i)\b(or|and)\s+\w+\s*=\s*\w+\s*(--|\s*;|\s*union)`),
	regexp.MustCompile(`(?i)\w+\s*=\s*\w+\s+(or|and)\s+\w+\s*=\s*\w+`),
	regexp.MustCompile(`(?i)\b(sleep|benchmark|waitfor|delay)\s*\(`),
	regexp.MustCompile(`(?i)\'\s*(or|and)\s+\'\w*\'\s*=\s*\'\w*`),
	regexp.MustCompile(`(?i)\'\s*;\s*(drop|delete|insert|update)`),
}

var xssPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)<script[^>]*>.*?</script>`),
	regexp.MustCompile(`(?i)<iframe[^>]*>.*?</iframe>`),
	regexp.MustCompile(`(?i)<object[^>]*>.*?</object>`),
	regexp.MustCompile(`(?i)<embed[^>]*>.*?</embed>`),
	regexp.MustCompile(`(?i)javascript:`),
	regexp.MustCompile(`(?i)vbscript:`),
	regexp.MustCompile(`(?i)on\w+\s*=`),
}

// --- LOGIKA VALIDASI (Pure Go) ---

func CheckSQLInjection(input string) bool {
	input = strings.ToLower(strings.TrimSpace(input))
	for _, pattern := range sqlInjectionPatterns {
		if pattern.MatchString(input) {
			return true
		}
	}
	return false
}

func CheckXSS(input string) bool {
	input = strings.ToLower(strings.TrimSpace(input))
	for _, pattern := range xssPatterns {
		if pattern.MatchString(input) {
			return true
		}
	}
	return false
}

func SanitizeInput(input string) string {
	sanitized := strings.TrimSpace(input)
	sanitized = SanitizeXSS(sanitized)
	spaceRegex := regexp.MustCompile(`\s+`)
	sanitized = spaceRegex.ReplaceAllString(sanitized, " ")
	return sanitized
}

func SanitizeXSS(input string) string {
	// Menghapus tag berbahaya menggunakan regex yang kamu buat
	for _, pattern := range xssPatterns {
		input = pattern.ReplaceAllString(input, "")
	}
	return input
}

func ValidateAndSanitizeText(input string, fieldName string, minLength, maxLength int) (string, []ValidError) {
	var errors []ValidError
	sanitized := SanitizeInput(input)

	if CheckSQLInjection(sanitized) || CheckXSS(sanitized) {
		errors = append(errors, ValidError{
			Field:   fieldName,
			Message: "Input contains potentially dangerous content",
		})
	}

	if len(sanitized) < minLength {
		errors = append(errors, ValidError{Field: fieldName, Message: fmt.Sprintf("Minimum %d characters", minLength)})
	}
	if len(sanitized) > maxLength {
		errors = append(errors, ValidError{Field: fieldName, Message: fmt.Sprintf("Maximum %d characters", maxLength)})
	}

	return sanitized, errors
}

// --- VALIDASI PASSWORD ---
func ValidatePassword(password string) []ValidError {
	var errors []ValidError
	if len(password) < 8 {
		errors = append(errors, ValidError{Field: "password", Message: "Minimum 8 characters"})
	}
	// Tambahkan pengecekan regex uppercase, lowercase, digit, special char (sama seperti kodemu)
	patterns := map[string]string{
		"[A-Z]": "one uppercase letter",
		"[a-z]": "one lowercase letter",
		"[0-9]": "one digit",
		`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~` + "`" + `]`: "one special character",
	}

	for pattern, msg := range patterns {
		if !regexp.MustCompile(pattern).MatchString(password) {
			errors = append(errors, ValidError{Field: "password", Message: "Must contain at least " + msg})
		}
	}
	return errors
}

// --- FILE UPLOAD SECURITY ---

func IsValidImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp":
		return true
	}
	return false
}

func ValidateFileUpload(filename string, fileSize int64, maxSize int64) []ValidError {
	var errors []ValidError
	if fileSize > maxSize {
		errors = append(errors, ValidError{Field: "file", Message: "File size too large"})
	}
	if !IsValidImageFile(filename) {
		errors = append(errors, ValidError{Field: "file", Message: "Invalid file type (Only Images)"})
	}

	dangerousPatterns := []string{".exe", ".bat", ".vbs", ".js", ".php", ".html", "../", "..\\"}
	for _, p := range dangerousPatterns {
		if strings.Contains(strings.ToLower(filename), p) {
			errors = append(errors, ValidError{Field: "file", Message: "Dangerous filename detected"})
			break
		}
	}
	return errors
}
