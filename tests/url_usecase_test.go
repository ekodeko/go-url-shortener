package tests

import (
	"testing"
	"time"
	"github.com/stretchr/testify/require"
	"github.com/yourname/url-shortener/internal/config"
	"github.com/yourname/url-shortener/internal/domain/url"
)

type inMemRepo struct{ m map[string]*url.Entity }

func (r *inMemRepo) Create(e *url.Entity) error { if r.m==nil {r.m=map[string]*url.Entity{}}; r.m[e.Code]=e; return nil }
func (r *inMemRepo) FindByCode(code string) (*url.Entity, error) { if e,ok:=r.m[code]; ok {return e,nil}; return nil, url.ErrNotFound }
func (r *inMemRepo) IncrementClicks(code string) error { if e,ok:=r.m[code]; ok { e.Clicks++; return nil }; return url.ErrNotFound }
func (r *inMemRepo) DeleteByCode(code string) error { if _,ok:=r.m[code]; ok { delete(r.m, code); return nil }; return url.ErrNotFound }

func TestShortenAndResolve(t *testing.T) {
	repo := &inMemRepo{}
	cfg := &config.Config{}
	uc := url.NewUsecase(repo, cfg)

	ent, err := uc.Shorten("https://example.com", nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, ent.Code)

	got, err := uc.Resolve(ent.Code)
	require.NoError(t, err)
	require.Equal(t, "https://example.com", got.Original)
}

func TestExpiry(t *testing.T) {
	repo := &inMemRepo{}
	cfg := &config.Config{}
	uc := url.NewUsecase(repo, cfg)

	d := 1 * time.Nanosecond
	ent, err := uc.Shorten("https://e.com", nil, &d)
	require.NoError(t, err)
	time.Sleep(2 * time.Nanosecond)
	_, err = uc.Resolve(ent.Code)
	require.Error(t, err)
}
