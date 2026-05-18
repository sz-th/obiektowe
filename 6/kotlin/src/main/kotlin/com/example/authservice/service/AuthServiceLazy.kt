package com.example.authservice.service

import org.springframework.context.annotation.Lazy
import org.springframework.stereotype.Service

@Service
@Lazy
class AuthServiceLazy {

	fun authenticate(username: String, password: String): Boolean =
		AuthCredentials.authenticate(username, password)
}
