package utils

type Server struct {
	Port string
}

type Config struct {
	Port string
}

type Request struct {
	Title       int    `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type Response struct {
	Message string `json:"message"`
}
