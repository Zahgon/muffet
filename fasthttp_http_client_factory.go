package main

type fasthttpHttpClientFactory struct {
}

func newFasthttpHttpClientFactory() *fasthttpHttpClientFactory {
	_ = "STUB: not implemented"
	return nil
}

func (*fasthttpHttpClientFactory) Create(o httpClientOptions) httpClient {
	_ = "STUB: not implemented"
	return *new(httpClient)
}
