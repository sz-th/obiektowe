package com.example.authservice.controller

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api")
class AuthController {

    @GetMapping("/users")
    fun getUsers(): List<Map<String, String>> {
        return listOf(
            mapOf("username" to "admin", "role" to "ADMIN"),
            mapOf("username" to "user", "role" to "USER")
        )
    }
}
