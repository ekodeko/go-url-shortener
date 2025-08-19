package url

import (
	"errors"
	"fmt"
	neturl "net/url"
	"time"
	"github.com/yourname/url-shortener/internal/config"
	"github.com/yourname/url-shortener/internal/pkg/shortid"
)

type Usecase interface {
	Shorten(original string, customAlias *string, ttl *time.Duration) (*Entity, error)
	Resolve(code string) (*Entity, error)
	Stats(code string) (*Entity, error)
	Delete(code string) error
}

type usecase struct {
	repo Repository
	cfg  *config.Config
}

func NewUsecase(r Repository, cfg *config.Config) Usecase { return &usecase{repo: r, cfg: cfg} }

func (u *usecase) Shorten(original string, customAlias *string, ttl *time.Duration) (*Entity, error) {
	if _, err := neturl.ParseRequestURI(original); err != nil { return nil, fmt.Errorf("invalid_url: %w", err) }
	code := ""
	if customAlias != nil && *customAlias != "" { code = *customAlias } else { code = shortid.Generate(7) }
	var exp *time.Time
	if ttl != nil && *ttl > 0 { t := time.Now().Add(*ttl); exp = &t } else if u.cfg.DefaultTTL > 0 { t := time.Now().Add(u.cfg.DefaultTTL); exp = &t }
	ent := &Entity{ Code: code, Original: original, ExpiresAt: exp }
	if err := u.repo.Create(ent); err != nil { return nil, err }
	return ent, nil
}

func (u *usecase) Resolve(code string) (*Entity, error) {
	ent, err := u.repo.FindByCode(code)
	if err != nil { return nil, err }
	if ent.ExpiresAt != nil && time.Now().After(*ent.ExpiresAt) { return nil, errors.New("expired") }
	_ = u.repo.IncrementClicks(code)
	return ent, nil
}

func (u *usecase) Stats(code string) (*Entity, error) { return u.repo.FindByCode(code) }

func (u *usecase) Delete(code string) error { return u.repo.DeleteByCode(code) }
