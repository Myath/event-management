package comment

import (
	"event-management/hrm/storage"
	"fmt"
)


type CommentStore interface {
	InsertComment(s storage.Comments) (*storage.Comments, error)
	GetCommentID(id int) (*storage.Comments, error)
	UpdateComment(s storage.Comments) (*storage.Comments, error)
	CommentsOfEvent(eventID int) ([]storage.Comments, error)
}

type CoreComment struct {
	store CommentStore
}

func NewCoreComment (cs CommentStore) *CoreComment{
    return &CoreComment{
    	store: cs,
    }
}

func (cc CoreComment) InsertComment(s storage.Comments) (*storage.Comments, error){
	comment , err := cc.store.InsertComment(s)
	if err != nil {
		return nil, err
	}

	if comment == nil {
		return nil, fmt.Errorf("unable to insert comment")
	}

	return comment, nil
}

func (cc CoreComment) EditComment(s storage.Comments) (*storage.Comments, error){
	comment, err := cc.store.GetCommentID(s.ID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (cc CoreComment) UpdateComment(s storage.Comments) (*storage.Comments, error){
	comment, err := cc.store.UpdateComment(s)
	if err != nil {
		return nil, err
	}

	if comment == nil {
		return nil, fmt.Errorf("unable to update event")
	}

	return comment, nil
}

func (cc CoreComment) CommentsOfEvent(s storage.Comments) ([]storage.Comments, error){
	comment, err :=  cc.store.CommentsOfEvent(s.EventID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}