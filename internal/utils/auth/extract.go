package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

// // ExtractBearerToken extracts the Bearer token from the request
// func ExtractBearerToken(ctx *gin.Context) (string, error) {
// 	authHeader := ctx.GetHeader("Authorization")
// 	if authHeader == "" {
// 		return "", errors.New("authorization header is required")
// 	}

// 	parts := strings.SplitN(authHeader, " ", 2)
// 	if !(len(parts) == 2 && parts[0] == "Bearer") {
// 		return "", errors.New("invalid authorization header format")
// 	}

// 	return parts[1], nil
// }

// ExtractUserID extracts the user ID from the context
// In a real implementation, this would decrypt and validate the JWT token
func ExtractUserID(ctx context.Context) (string, error) {
	// If using Gin context
	if ginCtx, ok := ctx.(*gin.Context); ok {
		// Get user ID from claims or context
		userID, exists := ginCtx.Get("user_id")
		if !exists {
			return "", errors.New("user ID not found in context")
		}
		
		// Convert to string
		if id, ok := userID.(string); ok {
			return id, nil
		}
		
		// If it's an int or uint, convert to string
		if id, ok := userID.(uint); ok {
			return fmt.Sprintf("%d", id), nil
		}
		if id, ok := userID.(int); ok {
			return fmt.Sprintf("%d", id), nil
		}
		
		return "", errors.New("user ID is of invalid type")
	}
	
	// For testing or non-Gin contexts, you might have a different way to extract
	// the userID directly from the context
	if userID, ok := ctx.Value("user_id").(string); ok {
		return userID, nil
	}
	
	return "", errors.New("user ID not found in context")
} 