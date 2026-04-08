package com.example.authservice.service

import org.springframework.stereotype.Service

@Service
class AuthServiceEager {
    
    init {
        println("AuthServiceEager initialized")
    }

    fun authenticate(username: String, password: String): Boolean {
        // Mock authentication
        return username == "admin" && password == "admin"
    }
}
