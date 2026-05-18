package com.example.authservice.controller

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController
import com.example.authservice.dto.AuthRequest
import com.example.authservice.service.AuthServiceEager
import com.example.authservice.service.AuthServiceLazy
import org.springframework.context.annotation.Lazy

@RestController
@RequestMapping("/api")
class AuthController(
	private val authServiceEager: AuthServiceEager,
	@Lazy private val authServiceLazy: AuthServiceLazy
) {

	@PostMapping("/login")
	fun login(@RequestBody request: AuthRequest): Map<String, Any> =
		loginResponse(authServiceEager.authenticate(request.username, request.password), "Authentication successful", "Invalid credentials")

	@PostMapping("/login-lazy")
	fun loginLazy(@RequestBody request: AuthRequest): Map<String, Any> =
		loginResponse(authServiceLazy.authenticate(request.username, request.password), "Authentication successful (Lazy)", "Invalid credentials (Lazy)")

	@GetMapping("/users")
	fun getUsers(): List<Map<String, String>> = listOf(
		mapOf("username" to "admin", "role" to "ADMIN"),
		mapOf("username" to "user", "role" to "USER")
	)

	private fun loginResponse(isAuthenticated: Boolean, successMessage: String, failureMessage: String): Map<String, Any> =
		if (isAuthenticated) {
			mapOf("success" to true, "message" to successMessage)
		} else {
			mapOf("success" to false, "message" to failureMessage)
		}
}
