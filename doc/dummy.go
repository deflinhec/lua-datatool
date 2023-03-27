package doc

type dummyFile struct {
}

func NewDummy() File {
	return &dummyFile{}
}

func (f *dummyFile) Path() string {
	return ""
}

func (f *dummyFile) Module() string {
	return ""
}

func (f *dummyFile) Field() string {
	return ""
}

func (f *dummyFile) Md5Sum() string {
	return ""
}

func (f *dummyFile) read() (interface{}, error) {
	return nil, nil
}

func (f *dummyFile) write(value interface{}) error {
	return nil
}
