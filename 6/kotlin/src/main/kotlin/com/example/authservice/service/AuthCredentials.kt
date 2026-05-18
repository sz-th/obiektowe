package com.example.authservice.service

import org.springframework.beans.factory.annotation.Value
import org.springframework.stereotype.Component

@Component
class AuthCredentials(
	@Value("\${app.auth.demo-user}") private val demoUser: String,
	@Value("\${app.auth.demo-password}") private val demoPassword: String
) {
	fun authenticate(username: String, password: String): Boolean =
		username == demoUser && password == demoPassword
}
