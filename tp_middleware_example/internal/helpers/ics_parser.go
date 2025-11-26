package helpers

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
    "middleware/example/internal/models"
)

// ParseICS parse un fichier ICS et retourne une liste d'événements
func ParseICS(filePath string) ([]models.Event, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("could not open file: %v", err)
    }
    defer file.Close()

    var events []models.Event
    var currentEvent models.Event

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        // Détection des sections de l'event
        if strings.HasPrefix(line, "BEGIN:VEVENT") {
            currentEvent = models.Event{}
        }

        if strings.HasPrefix(line, "DTSTART:") {
            startTime, _ := time.Parse("20060102T150405Z", strings.TrimPrefix(line, "DTSTART:"))
            currentEvent.Start = startTime
        }

        if strings.HasPrefix(line, "DTEND:") {
            endTime, _ := time.Parse("20060102T150405Z", strings.TrimPrefix(line, "DTEND:"))
            currentEvent.End = endTime
        }

        if strings.HasPrefix(line, "SUMMARY:") {
            currentEvent.Name = strings.TrimPrefix(line, "SUMMARY:")
        }

        if strings.HasPrefix(line, "LOCATION:") {
            currentEvent.Location = strings.TrimPrefix(line, "LOCATION:")
        }

        if strings.HasPrefix(line, "DESCRIPTION:") {
            currentEvent.Description = strings.TrimPrefix(line, "DESCRIPTION:")
        }

        if strings.HasPrefix(line, "UID:") {
            currentEvent.UID = strings.TrimPrefix(line, "UID:")
        }

        if strings.HasPrefix(line, "LAST-MODIFIED:") {
            lastModified, _ := time.Parse("20060102T150405Z", strings.TrimPrefix(line, "LAST-MODIFIED:"))
            currentEvent.LastUpdate = lastModified
        }

        // Fin d'un événement
        if strings.HasPrefix(line, "END:VEVENT") {
            events = append(events, currentEvent)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error while reading file: %v", err)
    }

    return events, nil
}
