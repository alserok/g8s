package memcached

import (
	"context"
	"encoding/json"
	"time"

	"github.com/alserok/g8s/internal/cache"
	"github.com/alserok/g8s/internal/utils/errors"
	"github.com/bradfitz/gomemcache/memcache"
)

func New(host string, exp time.Duration) cache.Repository {
	return &memcached{
		cl:  memcache.New(host),
		exp: exp,
	}
}

type memcached struct {
	cl  *memcache.Client
	exp time.Duration
}

func (m memcached) Set(ctx context.Context, key string, val any) error {
	b, err := json.Marshal(val)
	if err != nil {
		return errors.New(err.Error(), errors.ErrInternal)
	}

	if err = m.cl.Add(&memcache.Item{Value: b, Key: key, Expiration: int32(m.exp.Seconds())}); err != nil {
		return errors.New(err.Error(), errors.ErrInternal)
	}

	return nil
}

func (m memcached) Get(ctx context.Context, key string, target any) error {
	val, err := m.cl.Get(key)
	if err != nil {
		return errors.New(err.Error(), errors.ErrInternal)
	}

	if err = json.Unmarshal(val.Value, target); err != nil {
		return errors.New(err.Error(), errors.ErrInternal)
	}

	return nil
}

func (m memcached) GetBytes(ctx context.Context, key string) ([]byte, error) {
	val, err := m.cl.Get(key)
	if err != nil {
		return nil, errors.New(err.Error(), errors.ErrInternal)
	}

	return val.Value, nil
}

func (m memcached) Close() error {
	return m.cl.Close()
}
