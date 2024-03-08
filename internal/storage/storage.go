package storage

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/minism/trk/internal/config"
	"github.com/minism/trk/pkg/model"
	"gopkg.in/yaml.v3"
)

const (
	logEntryDateFormat = "2006-01-02"
	csvDelimeter       = ','
)

func SaveConfig(cfg config.TrkConfig) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(config.GetConfigPath(), data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig() (config.TrkConfig, error) {
	data, err := os.ReadFile(config.GetConfigPath())
	if err != nil {
		return config.TrkConfig{}, err
	}
	var cfg config.TrkConfig
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return config.TrkConfig{}, err
	}
	return cfg, nil
}

// TODO: Obviously revisit the scheme here for performance.
func LoadProjectLogEntries(project model.Project) ([]model.LogEntry, error) {
	file, err := os.Open(project.WorkLogPath())
	if err != nil {
		if os.IsNotExist(err) {
			return []model.LogEntry{}, nil
		}
		return nil, err
	}
	defer file.Close()

	// Parse csv.
	reader := csv.NewReader(file)
	reader.Comma = csvDelimeter
	entries := make([]model.LogEntry, 0)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		entry, err := parseLogEntryCsvRow(row)
		if err != nil {
			return nil, err
		}
		entry.Project = project
		entries = append(entries, entry)
	}

	return entries, nil
}

// For now this is destructive, we will need to consider mutex things for
// the TUI or just multiple async commands.
func SaveProjectLogEntries(projectId string, entries []model.LogEntry) error {
	file, err := os.Create(config.GetProjectWorkLogPath(projectId))
	if err != nil {
		return err
	}
	defer file.Close()

	// Write csv.
	writer := csv.NewWriter(file)
	writer.Comma = csvDelimeter
	for _, entry := range entries {
		row := formatLogEntryCsvRow(entry)
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}
	writer.Flush()

	return writer.Error()
}

func parseLogEntryCsvRow(row []string) (model.LogEntry, error) {
	date, err := time.Parse(logEntryDateFormat, row[0])
	if err != nil {
		return model.LogEntry{}, err
	}
	hours, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		return model.LogEntry{}, err
	}
	note := row[2]
	return model.LogEntry{
		Date:  date,
		Hours: hours,
		Note:  note,
	}, nil
}

func formatLogEntryCsvRow(entry model.LogEntry) []string {
	row := []string{
		entry.Date.Format(logEntryDateFormat),
		strconv.FormatFloat(entry.Hours, 'f', -1, 64),
		entry.Note,
	}
	return row
}
