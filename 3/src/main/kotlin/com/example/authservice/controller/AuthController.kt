package com.example.authservice.controller

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController
import com.example.authservice.dto.AuthRequest
import com.example.authservice.service.AuthServiceEager

@RestController
@RequestMapping("/api")
class AuthController(
    private val authService: AuthServiceEager
) {

    @PostMapping("/login")
    fun login(@RequestBody request: AuthRequest): Map<String, Any> {
        val isAuthenticated = authService.authenticate(request.username, request.password)
        return if (isAuthenticated) {
            mapOf("success" to true, "message" to "Authentication successful")
        } else {
            mapOf("success" to false, "message" to "Invalid credentials")
        }
    }

    @GetMapping("/users")
    fun getUsers(): List<Map<String, String>> {
        return listOf(
            mapOf("username" to "admin", "role" to "ADMIN"),
            mapOf("username" to "user", "role" to "USER")
        )
    }
}
