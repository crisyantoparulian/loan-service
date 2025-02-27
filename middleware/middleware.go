package middleware

import (
	"net/http"
	"strings"

	"github.com/crisyantoparulian/loansvc/utils/constant"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Request().Header.Get(constant.KEY_HEADER_USER_ID)
		userRole := c.Request().Header.Get(constant.KEY_HEADER_USER_ROLE)

		if userID == "" || userRole == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"success": false, "message": "Unauthorized"})
		}

		if !constant.Role(userRole).IsValid() {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"success": false, "message": "Unauthorized Role"})
		}
		return next(c)
	}
}

// RolePermissions defines which roles can access specific routes and methods
var RolePermissions = map[string]map[string][]string{
	"GET": {
		"/loans":         {"admin", "field_officer"},
		"/loans/:loanID": {"admin", "field_officer", "borrower"},
	},
	"POST": {
		"/loans":                  {"borrower", "admin"},
		"/loans/:loanID/visit":    {"field_validator", "admin"},
		"/loans/:loanID/invest":   {"investor", "admin"},
		"/loans/:loanID/disburse": {"investor", "admin"},
	},
	"PUT": {
		"/loans/:loanID/approve": {"admin"},
	},
}

// RoleMiddleware checks if the user has permission to access the requested path and method.
func RoleMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole := c.Request().Header.Get("login-user-role")
			if userRole == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing user role"})
			}

			// Get request path and method
			path := c.Path() // Echo normalizes routes like "/loans/:loan_id"
			method := c.Request().Method

			// Check if method exists in RolePermissions
			allowedRoutes, methodExists := RolePermissions[method]
			if !methodExists {
				return c.JSON(http.StatusForbidden, map[string]interface{}{"success": false, "error": "method not allowed"})
			}

			// Check if the path exists for this method
			allowedRoles, pathExists := allowedRoutes[path]
			if !pathExists {
				return c.JSON(http.StatusForbidden, map[string]interface{}{"success": false, "error": "access denied"})
			}

			// Check if the user's role is allowed
			for _, role := range allowedRoles {
				if strings.EqualFold(userRole, role) {
					return next(c) // Allow access
				}
			}

			return c.JSON(http.StatusForbidden, map[string]interface{}{"success": false, "error": "access denied"})
		}
	}
}
