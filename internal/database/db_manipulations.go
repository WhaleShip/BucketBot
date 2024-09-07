package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/WhaleShip/BucketBot/internal/database/models"
	"github.com/jackc/pgx/v5"
)

func getNoteName(noteText string) string {
	newlinePos := strings.Index(noteText, "\n")
	if newlinePos != -1 {
		noteText = noteText[:newlinePos]
	}

	if len(noteText) >= 8 {
		return noteText[:6] + "…"
	}

	return noteText
}

func AddNewNote(session *pgx.Conn, noteText string, userID int) error {
	var noteID int
	tx, err := session.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("error creating transition: %w", err)
	}
	defer tx.Rollback(context.Background())

	noteName := getNoteName(noteText)
	err = tx.QueryRow(context.Background(), "INSERT INTO Notes (text, name) VALUES ($1, $2) RETURNING id", noteText, noteName).Scan(&noteID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(context.Background(), "INSERT INTO UserNotes (user_id, note_id) VALUES ($1, $2)", userID, noteID)
	if err != nil {
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func GetNotesByUserID(conn *pgx.Conn, userID int) ([]models.Note, error) {
	query := `
	 SELECT n.id, n.text, n.name
	 FROM UserNotes un
	 INNER JOIN Notes n ON n.id = un.note_id
	 WHERE un.user_id = $1`

	rows, err := conn.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Name, &note.Text); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		notes = append(notes, note)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to iterate over rows: %w", rows.Err())
	}
	log.Println(userID)
	return notes, nil
}
