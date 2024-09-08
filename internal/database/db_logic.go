package database

import (
	"context"
	"fmt"
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

func GetSomeUserNotes(conn *pgx.Conn, userID, limit, offset int) ([]models.Note, error) {
	rows, err := conn.Query(context.Background(), `
        SELECT n.id, n.name, n.text
        FROM Notes n
        JOIN UserNotes un ON n.id = un.note_id
        WHERE un.user_id = $1
        LIMIT $2 OFFSET $3
    `, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []models.Note{}
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Name, &note.Text); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return notes, nil
}

func GetNoteByIDForOwner(session *pgx.Conn, noteID, requesterID int) (*models.Note, error) {
	query :=
		`SELECT UserNotes.user_id, Notes.id, Notes.name, Notes.text FROM UserNotes
	INNER JOIN Notes ON Notes.id = UserNotes.note_id
	WHERE note_id = $1 AND user_id = $2`
	row := session.QueryRow(context.Background(), query, noteID, requesterID)

	var ownerID int
	var note models.Note
	err := row.Scan(&ownerID, &note.ID, &note.Name, &note.Text)
	if err != nil {
		return nil, fmt.Errorf("unable to get note: %v", err)
	}
	if ownerID != requesterID {
		return nil, fmt.Errorf("requester and owner id are not equal. requesterID=%d, ownerID=%d", requesterID, ownerID)
	}

	return &note, nil
}
