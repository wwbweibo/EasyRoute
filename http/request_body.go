package http

type requestBody struct {
}

func (r requestBody) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func (r requestBody) Close() error {
	panic("implement me")
}
