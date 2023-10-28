package bot

import (
	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/irc"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// This is the main entry point for the bot.
// It will start the bot and connect to Twitch Chat
// TODO: Add a way to manage the bot actively being connected to Twitch or not without having to kill the process.
func main() {
	// A channel for listening to system signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT, os.Interrupt, os.Kill)

	// Referenced from https://github.com/Adeithe/go-twitch/blob/d21a5590a5a3d0b4e092daf3d81df3ed327b6be5/.examples/chat/main.go#L24
	reader := twitch.IRC()
	reader.OnShardReconnect(func(shardID int) {
		println("Reconnecting shard", shardID)
	})
	reader.OnShardLatencyUpdate(func(shardID int, latency time.Duration) {
		println("Shard", shardID, "Latency:", latency)
	})
	reader.OnShardMessage(func(shardID int, message irc.ChatMessage) {
		// TODO: Serialize the message into a chat-synthesizer friendly format
		println("Shard", shardID, "Message:", message)
	})
}
