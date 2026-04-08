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
    fun login(@RequestBody request: AuthRequest): Map<String, Any> {
        val isAuthenticated = authServiceEager.authenticate(request.username, request.password)
        return if (isAuthenticated) {
            mapOf("success" to true, "message" to "Authentication successful")
        } else {
            mapOf("success" to false, "message" to "Invalid credentials")
        }
    }

    @PostMapping("/login-lazy")
    fun loginLazy(@RequestBody request: AuthRequest): Map<String, Any> {
        val isAuthenticated = authServiceLazy.authenticate(request.username, request.password)
        return if (isAuthenticated) {
            mapOf("success" to true, "message" to "Authentication successful (Lazy)")
        } else {
            mapOf("success" to false, "message" to "Invalid credentials (Lazy)")
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
