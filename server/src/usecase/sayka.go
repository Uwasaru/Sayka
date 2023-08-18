package usecase

import (
	"context"

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
}

// ISaykaUsecaseは投稿に関するユースケースを担当します
type ISaykaUsecase interface {
	// GetByIDはIDを指定して投稿を取得します
	GetByID(ctx context.Context, id string) (*entity.Sayka, error)
	// GetByUserIDはUserIDを指定して投稿を取得します
	GetByUserID(ctx context.Context, userID string) (*entity.Saykas, error)
	// GetAllは全ての投稿を取得します
	GetAll(ctx context.Context) (*entity.Saykas, error)
	// GetTimeLineはタイムラインを取得します
	GetTimeLine(ctx context.Context, id string, tag string) (*entity.Saykas, error)
	// Createは投稿を作成します
	CreateSayka(ctx *gin.Context, sayka *entity.Sayka) error
	// UpdateSaykaは投稿を更新します
	UpdateSayka(ctx context.Context, sayka *entity.Sayka) error
	// DeleteSaykaは投稿を削除します
	DeleteSayka(ctx context.Context, id string) error
}

// NewSaykaUsecaseは新しいSaykaUsecaseを初期化し構造体のポインタを返します
func NewSaykaUsecase(pr repository.SaykaRepository, fr repository.FavoriteRepository, cr repository.CommentRepository, ur repository.UserRepository) ISaykaUsecase {
	return &SaykaUsecase{
		pr: pr,
		fr: fr,
		cr: cr,
		ur: ur,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (pu *SaykaUsecase) GetByID(ctx context.Context, id string) (*entity.Sayka, error) {
	return pu.pr.GetByID(ctx, id)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (pu *SaykaUsecase) GetByUserID(ctx context.Context, userID string) (*entity.Saykas, error) {
	saykas, err := pu.pr.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	for _, sayka := range *saykas {
		fovorites, err := pu.fr.GetBySaykaID(ctx, sayka.ID)
		if err != nil {
			return nil, err
		}
		comments, err := pu.cr.GetBySaykaID(ctx, sayka.ID)
		if err != nil {
			return nil, err
		}
		s := set.NewSet(fovorites...)
		sayka.IsFavorite = s.Contains(sayka.UserID)
		sayka.Favorites = len(fovorites)
		sayka.Comments = len(*comments)
		user, err := pu.ur.GetUser(ctx, sayka.UserID)
		if err != nil {
			return nil, err
		}
		sayka.User = user
	}
	return saykas, err
}

// GetAllは全ての投稿を取得します
func (pu *SaykaUsecase) GetAll(ctx context.Context) (*entity.Saykas, error) {
	saykas, err := pu.pr.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, sayka := range *saykas {
		fovorites, err := pu.fr.GetBySaykaID(ctx, sayka.ID)
		if err != nil {
			return nil, err
		}
		comments, err := pu.cr.GetBySaykaID(ctx, sayka.ID)
		if err != nil {
			return nil, err
		}
		s := set.NewSet(fovorites...)
		sayka.IsFavorite = s.Contains(sayka.UserID)
		sayka.Favorites = len(fovorites)
		sayka.Comments = len(*comments)
		user, err := pu.ur.GetUser(ctx, sayka.UserID)
		if err != nil {
			return nil, err
		}
		sayka.User = user
	}
	return saykas, err
}

// GetTimeLineはタイムラインを取得します
func (pu *SaykaUsecase) GetTimeLine(ctx context.Context, id string, tag string) (*entity.Saykas, error) {
	saykas, err := pu.pr.GetTimeLine(ctx, id, tag)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	for _, sayka := range *saykas {
		fovorites, err := pu.fr.GetBySaykaID(ctx, sayka.ID)
		if err != nil {
			return nil, err
		}
		comments, err := pu.cr.GetBySaykaID(ctx, sayka.ID)
		if err != nil {
			return nil, err
		}
		s := set.NewSet(fovorites...)
		sayka.IsFavorite = s.Contains(sayka.UserID)
		sayka.Favorites = len(fovorites)
		sayka.Comments = len(*comments)
		user, err := pu.ur.GetUser(ctx, sayka.UserID)
		if err != nil {
			return nil, err
		}
		sayka.User = user
	}
	return saykas, err
}

// Createは投稿を作成します
func (pu *SaykaUsecase) CreateSayka(ctx *gin.Context, sayka *entity.Sayka) error {
	err := pu.pr.CreateSayka(ctx, sayka)
	if err != nil {
		return err
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
	return pu.pr.UpdateSayka(ctx, sayka)
}

// DeleteSaykaは投稿を削除します
func (pu *SaykaUsecase) DeleteSayka(ctx context.Context, id string) error {
	return pu.pr.DeleteSayka(ctx, id)
}
