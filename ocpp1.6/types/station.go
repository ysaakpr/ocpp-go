package types

import "context"

type Station interface {
	ID() string
	Get(string) interface{}
	GetString(string) (string, bool)
	Context() context.Context
	SetContext(context.Context)
}

type station struct {
	id  string
	ctx context.Context
}

func NewStation(id string, ctx context.Context) Station {
	return &station{
		id:  id,
		ctx: ctx,
	}
}

func (s *station) ID() string {
	return s.id
}

func (s *station) Context() context.Context {
	return s.ctx
}

func (s *station) Get(key string) interface{} {
	return s.ctx.Value(key)
}

func (s *station) GetString(key string) (string, bool) {
	v, ok := s.ctx.Value(key).(string)
	return v, ok
}

func (s *station) SetContext(ctx context.Context) {
	s.ctx = ctx
}
