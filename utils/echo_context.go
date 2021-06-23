package utils

import (
	"github.com/labstack/echo/v4"

	"github.com/Otobook-vn/modules/constant"
)

// GetUserIDFromEchoContext ...
func GetUserIDFromEchoContext(c echo.Context) string {
	userIDInterface := c.Get(constant.EchoContextKeyUserID)
	if userIDInterface == nil {
		return ""
	}
	return userIDInterface.(string)
}

// GetStaffIDFromEchoContext ...
func GetStaffIDFromEchoContext(c echo.Context) string {
	staffIDInterface := c.Get(constant.EchoContextKeyStaffID)
	if staffIDInterface == nil {
		return ""
	}
	return staffIDInterface.(string)
}

// GetLanguageFromEchoContext ...
func GetLanguageFromEchoContext(c echo.Context) string {
	lang := c.Get("lang")
	return lang.(string)
}
