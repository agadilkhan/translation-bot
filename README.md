# Translator Bot
### Simple Discord bot for translation

***

## Getting started
### Installation

1. Go to the [Discord developer portal](https://discord.com/developers)
2. Create a new application
3. Add a bot user to the application
4. Get the token for the bot
5. Clone the repository

   ```sh
   git clone https://github.com/agadilkhan/translation-bot.git
   ```

6. Install dependencies

   ```sh
   go mod tidy
   ```
7. Create the environment variables file `.env` from the `.env.example` file in the root folder and add the following:
   ```dotenv
    DISCORD_TOKEN=<Your discord token>
    ```
   
### Usage
To run the discord bot from root directory, execute the following command
```shell
go run main.go
```
![Screenshot 2023-04-03 181545] (https://github.com/agadilkhan/translation-bot/img/example_screen.png)