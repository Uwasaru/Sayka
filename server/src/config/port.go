package config

import "github.com/Uwasaru/Sayka/utils"

func Port() string {
	return utils.GetEnvOrDefault("PORT", "8080")
}
