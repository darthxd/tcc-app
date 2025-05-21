package router

import "time"

func FormatGrade(grade string) string {
	switch grade {
	case "1Y":
		return "Primeiro ano"
	case "2Y":
		return "Segundo ano"
	case "3Y":
		return "Terceiro ano"
	default:
		return "Erro ao formatar ano"
	}
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}
