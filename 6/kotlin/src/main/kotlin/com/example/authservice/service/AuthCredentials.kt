package com.example.authservice.service

object AuthCredentials {
	private const val ADMIN_USER = "admin"
	private const val ADMIN_PASSWORD = "admin"

	fun authenticate(username: String, password: String): Boolean =
		username == ADMIN_USER && password == ADMIN_PASSWORD
}
