package util

import "github.com/bwmarrin/discordgo"

func StringToLocale(s string) discordgo.Locale {
	locale := discordgo.Locale(s)

	// Check if the converted locale is valid
	if _, ok := discordgo.Locales[locale]; ok {
		return locale
	}

	return discordgo.Unknown
}
