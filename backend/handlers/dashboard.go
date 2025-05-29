package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")[7:] // 去掉 "Bearer "
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
		return
	}

	// 这里应该是从数据库获取用户数据的逻辑
	// 示例中我们返回模拟数据
	response := map[string]interface{}{
		"email":     claims.Email,
		"lastLogin": time.Now().Format(time.RFC3339),
		"stats": map[string]int{
			"projects":    15,
			"teamMembers": 42,
			"tasks":       128,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}