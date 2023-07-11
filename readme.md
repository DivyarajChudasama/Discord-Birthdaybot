# Discord Bot

This is a simple Discord bot built using the Go programming language. The bot can respond to certain messages and send reminders for birthdays.

## Features

- Greeting: The bot responds with a greeting message when it receives "hello" or "hi".
- Farewell: The bot responds with a goodbye message when it receives "goodbye".
- Birthday Reminders: The bot checks for birthdays and sends reminders to all users in the server.

## Usage

1. Clone the repository: `git clone https://github.com/YourUsername/YourRepository.git`
2. Set up a Discord bot and obtain a token.
3. Configure the bot's token and friend birthdays in the `config.json` file.
4. Build and run the bot: `go run main.go`.
5. Invite the bot to your Discord server using the generated authorization URL.
6. Interact with the bot by sending messages in the server's channels.

Feel free to customize the bot's responses and add more features based on your requirements.

## Dependencies

The bot uses the following dependencies:
- `github.com/bwmarrin/discordgo`: DiscordGo library for interacting with the Discord API.

Please refer to the documentation of each dependency for more information.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
