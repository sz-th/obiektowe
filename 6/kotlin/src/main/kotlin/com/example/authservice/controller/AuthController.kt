package com.example.authservice.controller

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController
import com.example.authservice.dto.AuthRequest
import com.example.authservice.dto.AuthResponse
import com.example.authservice.dto.UserDto
import com.example.authservice.service.AuthServiceEager
import com.example.authservice.service.AuthServiceLazy
import org.springframework.context.annotation.Lazy
import jakarta.validation.Valid

@RestController
@RequestMapping("/api")
class AuthController(
	private val authServiceEager: AuthServiceEager,
	@Lazy private val authServiceLazy: AuthServiceLazy
) {

	@PostMapping("/login")
	fun login(@Valid @RequestBody request: AuthRequest): AuthResponse =
		loginResponse(
			authServiceEager.authenticate(request.username, request.password),
			"Authentication successful",
			"Invalid credentials"
		)

	@PostMapping("/login-lazy")
	fun loginLazy(@Valid @RequestBody request: AuthRequest): AuthResponse =
		loginResponse(
			authServiceLazy.authenticate(request.username, request.password),
			"Authentication successful (Lazy)",
			"Invalid credentials (Lazy)"
		)

	@GetMapping("/users")
	fun getUsers(): List<UserDto> = listOf(
		UserDto(username = "admin", role = "ADMIN"),
		UserDto(username = "user", role = "USER")
	)

	private fun loginResponse(
		isAuthenticated: Boolean,
		successMessage: String,
		failureMessage: String
	): AuthResponse =
		if (isAuthenticated) {
			AuthResponse(success = true, message = successMessage)
		} else {
			AuthResponse(success = false, message = failureMessage)
		}
}
