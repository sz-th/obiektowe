package com.example.authservice.service

import org.springframework.stereotype.Service

@Service
class AuthServiceEager {

	fun authenticate(username: String, password: String): Boolean =
		AuthCredentials.authenticate(username, password)
}
