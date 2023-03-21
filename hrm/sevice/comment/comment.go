package comment

import (
	"context"
	commentpb "event-management/gunk/v1/comment"
	"event-management/hrm/storage"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CommentCore interface {
	InsertComment(s storage.Comments) (*storage.Comments, error)
	EditComment(s storage.Comments) (*storage.Comments, error)
	UpdateComment(s storage.Comments) (*storage.Comments, error)
	CommentsOfEvent(s storage.Comments) ([]storage.Comments, error)
}

type CommentSvc struct {
	commentpb.UnimplementedCommentServiceServer
	core CommentCore
}

func NewCommentSvc(cc CommentCore) *CommentSvc {
	return &CommentSvc{
		core: cc,
	}
}

func (cs CommentSvc) CreateComment(ctx context.Context, r *commentpb.CreateCommentRequest) (*commentpb.CreateCommentResponse, error) {
	comment := storage.Comments{
		UserId:  int(r.GetUserId()),
		EventID: int(r.GetEventID()),
		Comment: r.GetComment(),
	}


	if err := comment.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	ct, err := cs.core.InsertComment(comment)
	if err != nil {
		return nil, err
	}

	return &commentpb.CreateCommentResponse{
		Comments: &commentpb.Comments{
			ID:      int32(ct.ID),
			UserId:  int32(ct.UserId),
			EventID: int32(ct.EventID),
			Comment: ct.Comment,
		},
	}, nil
}

func (cs CommentSvc) EditComment(ctx context.Context, r *commentpb.EditCommentRequest) (*commentpb.EditCommentResponse, error) {
	commentID := storage.Comments{
		ID: int(r.GetID()),
	}

	ct, err := cs.core.EditComment(commentID)
	if err != nil {
		return nil, err
	}

	return &commentpb.EditCommentResponse{
		Comments: &commentpb.Comments{
			ID:      int32(ct.ID),
			UserId:  int32(ct.UserId),
			EventID: int32(ct.EventID),
			Comment: ct.Comment,
		},
	}, nil
}

func (cs CommentSvc) UpdateComment(xtc context.Context, r *commentpb.UpdateCommentRequest) (*commentpb.UpdateCommentResponse, error) {
	comment := storage.Comments{
		ID:      int(r.Comments.ID),
		UserId:  int(r.Comments.GetUserId()),
		EventID: int(r.Comments.GetEventID()),
		Comment: r.Comments.GetComment(),
	}

	if err := comment.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	ct, err := cs.core.UpdateComment(comment)
	if err != nil {
		return nil, err
	}

	return &commentpb.UpdateCommentResponse{
		Comments: &commentpb.Comments{
			ID:      int32(ct.ID),
			UserId:  int32(ct.UserId),
			EventID: int32(ct.EventID),
			Comment: ct.Comment,
		},
	}, nil

}

func (cs CommentSvc) CommentListOfEVents(ctx context.Context, r *commentpb.CommentListOfEVentsRequest) (*commentpb.CommentListOfEVentsResponse, error) {
	eventId := storage.Comments{
		EventID:       int(r.GetEventID()),
	}

	ct, err := cs.core.CommentsOfEvent(eventId)
	if err != nil {
		return nil, err
	}

	var allCommentsOfEvent []*commentpb.CommentList
	for _, sec := range ct{
		commentlist := &commentpb.CommentList{
			ID:            int32(sec.ID),
			UserId:        int32(sec.UserId),
			Username:      sec.Username,
			EventID:       int32(sec.EventID),
			EventTypeName: sec.EventTypeName,
			EventName:     sec.EventTypeName,
			Description:   sec.Description,
			Location:      sec.Location,
			StartAt:       timestamppb.New(sec.StartAt),
			EndAt:         timestamppb.New(sec.EndAt),
			PublishedAt:   timestamppb.New(sec.PublishedAt.Time),
			Comment:       sec.Comment,
		}
		allCommentsOfEvent = append(allCommentsOfEvent, commentlist)
	}
	return &commentpb.CommentListOfEVentsResponse{
		CmtOfEv: allCommentsOfEvent,
	}, nil
}
