package doc

type fileInfo struct {
	dir string

	name string

	field string

	module string

	md5sum string
}

type Option interface {
	apply(s *fileInfo)
}

type funcOption struct {
	f func(*fileInfo)
}

func (fdo *funcOption) apply(do *fileInfo) {
	fdo.f(do)
}

func newOption(f func(*fileInfo)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithDocument(document File) Option {
	return newOption(func(s *fileInfo) {
		s.module = document.Module()
		s.field = document.Field()
	})
}

func WithModule(module string) Option {
	return newOption(func(s *fileInfo) {
		s.module = module
	})
}

func WithField(field string) Option {
	return newOption(func(s *fileInfo) {
		s.field = field
	})
}
