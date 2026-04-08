package com.example.authservice.controller

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController
import com.example.authservice.dto.AuthRequest

@RestController
@RequestMapping("/api")
class AuthController {

    @PostMapping("/login")
    fun login(@RequestBody request: AuthRequest): Map<String, Any> {
        // To be implemented with service
        return mapOf("success" to false, "message" to "Not implemented yet")
    }

    @GetMapping("/users")
    fun getUsers(): List<Map<String, String>> {
        return listOf(
            mapOf("username" to "admin", "role" to "ADMIN"),
            mapOf("username" to "user", "role" to "USER")
        )
    }
}
