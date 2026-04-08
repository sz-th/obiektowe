package com.example.authservice.service

import org.springframework.context.annotation.Lazy
import org.springframework.stereotype.Service

@Service
@Lazy
class AuthServiceLazy {
    
    init {
        println("AuthServiceLazy initialized")
    }

    fun authenticate(username: String, password: String): Boolean {
        // Mock authentication
        return username == "admin" && password == "admin"
    }
}
