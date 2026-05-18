package com.example.authservice.service

import org.springframework.context.annotation.Lazy
import org.springframework.stereotype.Service

@Service
@Lazy
class AuthServiceLazy(
	private val authCredentials: AuthCredentials
) {
	fun authenticate(username: String, password: String): Boolean =
		authCredentials.authenticate(username, password)
}
