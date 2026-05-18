package com.example.authservice.dto

import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.Size

data class AuthRequest(
	@field:NotBlank
	@field:Size(max = 64)
	val username: String,

	@field:NotBlank
	@field:Size(max = 128)
	val password: String
)
