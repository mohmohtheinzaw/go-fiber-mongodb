package middleware

import (
	"fmt"
	"net/http"
)

func ExtractTokenFromHeader(token *http.Request) {
	fmt.Println(token)
	// if token == "" {
	// 	return "", fmt.Errorf("Authorization header is empty")
	// }
}
