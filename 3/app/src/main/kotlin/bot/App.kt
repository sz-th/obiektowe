package bot

import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.coroutines.runBlocking
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.Json

@Serializable
data class DiscordMessage(val content: String)

class DiscordClient(private val token: String) {
    private val client = HttpClient(CIO) {
        install(ContentNegotiation) {
            json(Json {
                ignoreUnknownKeys = true
            })
        }
    }

    suspend fun sendMessage(channelId: String, content: String) {
        val response = client.post("https://discord.com/api/v10/channels/$channelId/messages") {
            header(HttpHeaders.Authorization, "Bot $token")
            contentType(ContentType.Application.Json)
            setBody(DiscordMessage(content))
        }
        println("Message sent. Status: ${response.status}")
    }

    fun close() {
        client.close()
    }
}

fun main() = runBlocking {
    val envFile = java.io.File("../.env")
    val env = if (envFile.exists()) {
        envFile.readLines().filter { it.contains("=") }.associate {
            val (key, value) = it.split("=", limit = 2)
            key.trim() to value.trim()
        }
    } else emptyMap()

    val token = env["DISCORD_BOT_TOKEN"] ?: System.getenv("DISCORD_BOT_TOKEN") ?: error("DISCORD_BOT_TOKEN not set")
    
    val discordClient = DiscordClient(token)
    println("Discord Client initialized.")
    // discordClient.sendMessage("CHANNEL_ID", "Hello from Ktor!")
    discordClient.close()
}
