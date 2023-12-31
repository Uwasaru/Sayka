package usecase

import (
	"context"
	"log"

	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
	set "github.com/deckarep/golang-set/v2"
	"github.com/gin-gonic/gin"
)

var _ ISaykaUsecase = &SaykaUsecase{}

// SaykaUsecaseは投稿に関するユースケースを担当します
type SaykaUsecase struct {
	pr repository.SaykaRepository
	fr repository.FavoriteRepository
	cr repository.CommentRepository
	ur repository.UserRepository
	tr repository.TagRepository
}

// ISaykaUsecaseは投稿に関するユースケースを担当します
type ISaykaUsecase interface {
	// GetByIDはIDを指定して投稿を取得します
	GetByID(ctx context.Context, id string, userId string) (*entity.Sayka, error)
	// GetByUserIDはUserIDを指定して投稿を取得します
	GetByUserID(ctx context.Context, userID string, myID string) (*entity.Saykas, error)
	// GetMeは自分の投稿数、いいね数、コメント数を取得します
	GetMe(ctx context.Context, userID string) (*entity.Me, error)
	// GetAllは全ての投稿を取得します
	GetAll(ctx context.Context, mysId string) (*entity.Saykas, error)
	// GetTimeLineはタイムラインを取得します
	GetTimeLine(ctx context.Context, id string, tag string, myId string) (*entity.Saykas, error)
	// GetAllFavoriteSaykaはいいねした全ての投稿を取得します
	GetAllFavoriteSayka(ctx context.Context, userId string, myId string) (*entity.Saykas, error)
	// GetAllCommentSaykaはコメントした全ての投稿を取得します
	GetAllCommentSayka(ctx context.Context, userId string, myId string) (*entity.Saykas, error)
	// Createは投稿を作成します
	CreateSayka(ctx *gin.Context, sayka *entity.Sayka) error
	// UpdateSaykaは投稿を更新します
	UpdateSayka(ctx context.Context, sayka *entity.Sayka) error
	// DeleteSaykaは投稿を削除します
	DeleteSayka(ctx context.Context, id string) error
}

// NewSaykaUsecaseは新しいSaykaUsecaseを初期化し構造体のポインタを返します
func NewSaykaUsecase(pr repository.SaykaRepository, fr repository.FavoriteRepository, cr repository.CommentRepository, ur repository.UserRepository, tr repository.TagRepository) ISaykaUsecase {
	return &SaykaUsecase{
		pr: pr,
		fr: fr,
		cr: cr,
		ur: ur,
		tr: tr,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (pu *SaykaUsecase) GetByID(ctx context.Context, id string, userId string) (*entity.Sayka, error) {
	sayka, err := pu.pr.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	tags, err := pu.tr.GetBySaykaID(ctx, sayka.ID)
	if err != nil {
		return nil, err
	}
	tagNames := make([]string, 0, len(*tags))
	for _, tag := range *tags {
		tagNames = append(tagNames, tag.Name)
	}
	sayka.Tags = tagNames
	fovorites, err := pu.fr.GetBySaykaID(ctx, sayka.ID)
	if err != nil {
		return nil, err
	}
	comments, err := pu.cr.GetBySaykaID(ctx, sayka.ID)
	if err != nil {
		return nil, err
	}
	s := set.NewSet(fovorites...)
	sayka.IsFavorite = s.Contains(userId)
	sayka.Favorites = len(fovorites)
	sayka.Comments = len(*comments)
	user, err := pu.ur.GetUser(ctx, sayka.UserID)
	if err != nil {
		return nil, err
	}
	sayka.User = user
	sayka.IsMe = sayka.UserID == userId
	return sayka, err
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (pu *SaykaUsecase) GetByUserID(ctx context.Context, userID string, myId string) (*entity.Saykas, error) {
	saykas, err := pu.pr.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	for _, sayka := range *saykas {
		pu.CompleteSayka(ctx, sayka, myId)
	}
	return saykas, err
}

// GetMeは自分の投稿数、いいね数、コメント数を取得します
func (pu *SaykaUsecase) GetMe(ctx context.Context, userID string) (*entity.Me, error) {
	saykas, err := pu.pr.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	favorited_count, err := pu.fr.GetAllCountByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	comment_count, err := pu.cr.GetCountByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	user, err := pu.ur.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &entity.Me{
		User:           user,
		SaykaCount:     len(*saykas),
		FavoritedCount: favorited_count,
		CommentCount:   comment_count,
	}, nil
}

// GetAllは全ての投稿を取得します
func (pu *SaykaUsecase) GetAll(ctx context.Context, myId string) (*entity.Saykas, error) {
	saykas, err := pu.pr.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, sayka := range *saykas {
		pu.CompleteSayka(ctx, sayka, myId)
	}
	return saykas, err
}

// GetTimeLineはタイムラインを取得します
func (pu *SaykaUsecase) GetTimeLine(ctx context.Context, id string, tag string, myId string) (*entity.Saykas, error) {
	saykas, err := pu.pr.GetTimeLine(ctx, id, tag)
	if err != nil {
		return nil, err
	}
	for _, sayka := range *saykas {
		pu.CompleteSayka(ctx, sayka, myId)
	}
	return saykas, err
}

// GetAllFavoriteSaykaはいいねした全ての投稿を取得します
func (pu *SaykaUsecase) GetAllFavoriteSayka(ctx context.Context, userId string, myId string) (*entity.Saykas, error) {
	saykas, err := pu.pr.GetAllFavoriteSayka(ctx, userId)
	if err != nil {
		return nil, err
	}
	for _, sayka := range *saykas {
		pu.CompleteSayka(ctx, sayka, myId)
	}
	return saykas, err
}

// GetAllCommentSaykaはコメントした全ての投稿を取得します
func (pu *SaykaUsecase) GetAllCommentSayka(ctx context.Context, userId string, myId string) (*entity.Saykas, error) {
	saykas, err := pu.pr.GetAllCommentSayka(ctx, myId)
	if err != nil {
		return nil, err
	}
	//重複する投稿を削除
	for _, sayka := range *saykas {
		pu.CompleteSayka(ctx, sayka, myId)
	}
	return saykas, err
}

// Createは投稿を作成します
func (pu *SaykaUsecase) CreateSayka(ctx *gin.Context, sayka *entity.Sayka) error {
	err := pu.pr.CreateSayka(ctx, sayka)
	if err != nil {
		return err
	}
	log.Println(sayka.ID)
	tags := sayka.Tags
	for _, tagName := range tags {
		tag := &entity.Tag{
			SaykaID: sayka.ID,
			Name:    tagName,
		}
		err := pu.tr.CreateTag(ctx, tag)
		if err != nil {
			return err
		}
	}
	user, err := pu.ur.GetUser(ctx, sayka.UserID)
	if err != nil {
		return err
	}
	sayka.User = user
	return nil
}

// UpdateSaykaは投稿を更新します
func (pu *SaykaUsecase) UpdateSayka(ctx context.Context, sayka *entity.Sayka) error {
	err := pu.pr.UpdateSayka(ctx, sayka)
	if err != nil {
		return err
	}
	err = pu.tr.DeleteBySaykaID(ctx, sayka.ID)
	if err != nil {
		return err
	}
	tags := sayka.Tags
	for _, tagName := range tags {
		tag := &entity.Tag{
			SaykaID: sayka.ID,
			Name:    tagName,
		}
		err := pu.tr.CreateTag(ctx, tag)
		if err != nil {
			return err
		}
	}
	user, err := pu.ur.GetUser(ctx, sayka.UserID)
	if err != nil {
		return err
	}
	sayka.User = user
	return nil
}

// DeleteSaykaは投稿を削除します
func (pu *SaykaUsecase) DeleteSayka(ctx context.Context, id string) error {
	return pu.pr.DeleteSayka(ctx, id)
}

func (pu *SaykaUsecase) CompleteSayka(ctx context.Context, sayka *entity.Sayka, myId string) error {
	fovorites, err := pu.fr.GetBySaykaID(ctx, sayka.ID)
	if err != nil {
		return err
	}
	comments, err := pu.cr.GetBySaykaID(ctx, sayka.ID)
	if err != nil {
		return err
	}
	tags, err := pu.tr.GetBySaykaID(ctx, sayka.ID)
	if err != nil {
		return err
	}
	tagNames := make([]string, 0, len(*tags))
	for _, tag := range *tags {
		tagNames = append(tagNames, tag.Name)
	}
	sayka.Tags = tagNames
	s := set.NewSet(fovorites...)
	sayka.IsFavorite = s.Contains(myId)
	sayka.Favorites = len(fovorites)
	sayka.Comments = len(*comments)
	sayka.IsMe = sayka.UserID == myId
	user, err := pu.ur.GetUser(ctx, sayka.UserID)
	if err != nil {
		return err
	}
	sayka.User = user
	return nil
}