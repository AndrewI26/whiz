package http

type HttpMethod int

const (
	Get HttpMethod = iota
	Post
	Delete
	Put
	Patch
	Options
	Connect
)
