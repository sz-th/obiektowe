package bot

import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.plugins.websocket.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.serialization.kotlinx.*
import io.ktor.serialization.kotlinx.json.*
import io.ktor.websocket.*
import kotlinx.coroutines.*
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.*

@Serializable
data class DiscordMessage(val content: String)

class DiscordClient(private val token: String) {
    private val client = HttpClient(CIO) {
        install(ContentNegotiation) {
            json(Json { ignoreUnknownKeys = true })
        }
        install(WebSockets) {
            contentConverter = KotlinxWebsocketSerializationConverter(Json { ignoreUnknownKeys = true })
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

    suspend fun startGateway() {
        val gatewayUrl = "wss://gateway.discord.gg/?v=10&encoding=json"
        
        client.webSocket(gatewayUrl) {
            var heartbeatInterval = 41250L
            var seqNum: Int? = null

            // Start reading messages
            for (frame in incoming) {
                if (frame !is Frame.Text) continue
                val text = frame.readText()
                val json = Json.parseToJsonElement(text).jsonObject
                
                val op = json["op"]?.jsonPrimitive?.intOrNull
                val t = json["t"]?.jsonPrimitive?.contentOrNull
                if (json["s"]?.jsonPrimitive?.intOrNull != null) {
                    seqNum = json["s"]?.jsonPrimitive?.intOrNull
                }

                when (op) {
                    10 -> { // Hello
                        heartbeatInterval = json["d"]?.jsonObject?.get("heartbeat_interval")?.jsonPrimitive?.longOrNull ?: 41250L
                        // Launch heartbeat coroutine
                        launch {
                            while (isActive) {
                                delay(heartbeatInterval)
                                val heartbeatPayload = buildJsonObject {
                                    put("op", 1)
                                    put("d", seqNum?.let { JsonPrimitive(it) } ?: JsonNull)
                                }
                                send(Frame.Text(heartbeatPayload.toString()))
                            }
                        }
                        
                        // Send Identify
                        val identifyPayload = buildJsonObject {
                            put("op", 2)
                            put("d", buildJsonObject {
                                put("token", token)
                                put("intents", 33280) // GUILD_MESSAGES (512) | MESSAGE_CONTENT (32768)
                                put("properties", buildJsonObject {
                                    put("os", "windows")
                                    put("browser", "ktor")
                                    put("device", "ktor")
                                })
                            })
                        }
                        send(Frame.Text(identifyPayload.toString()))
                        println("Sent Identify to Gateway")
                    }
                    0 -> { // Dispatch
                        if (t == "MESSAGE_CREATE") {
                            val data = json["d"]?.jsonObject
                            val author = data?.get("author")?.jsonObject
                            val isBot = author?.get("bot")?.jsonPrimitive?.booleanOrNull == true
                            if (!isBot) {
                                val content = data?.get("content")?.jsonPrimitive?.content
                                val channelId = data?.get("channel_id")?.jsonPrimitive?.content
                                println("Received message from channel $channelId: $content")
                            }
                        }
                    }
                }
            }
        }
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
    
    // Launch gateway to listen for messages
    launch {
        try {
            discordClient.startGateway()
        } catch (e: Exception) {
            e.printStackTrace()
        }
    }
    
    // Keep alive
    delay(Long.MAX_VALUE)
    discordClient.close()
}
