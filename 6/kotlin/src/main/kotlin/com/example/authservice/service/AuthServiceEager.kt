package com.example.authservice.service

import org.springframework.stereotype.Service

@Service
class AuthServiceEager(
	private val authCredentials: AuthCredentials
) {
	fun authenticate(username: String, password: String): Boolean =
		authCredentials.authenticate(username, password)
}
