package config

import ()

const (
	// AI_TEXT_GENERATION_URL is the URL of the AI response model
	AI_TEXT_GENERATION_URL = "https://api.deepseek.com/v1"

	AI_TEXT_GENERATION_MODEL = "deepseek-reasoner"

	// USER_ROLE is the role config of the user for AI response model
	USER_ROLE = "user"

	// SYSTEM_ROLE is the role config of the user for AI response model
	SYSTEM_ROLE = "system"

	// SYSTEM_ROLE_CONTENT is the content config of the system for AI response model
	SYSTEM_ROLE_CONTENT = "You are a technical writer. You are writing a blog post about the latest trends in technology."

	// MAX_TOKENS is the max tokens config of the AI response model
	MAX_TOKENS = 2046*5

	// TEMPERATURE is the temperature config of the AI response model
	TEMPERATURE = 1.5

	// STREAM is the stream config of the AI response model
	STREAM = false
)
