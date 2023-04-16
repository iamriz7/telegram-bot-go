import os
from dotenv import load_dotenv
import telegram
from telegram.ext import Updater, MessageHandler, Filters, CommandHandler
import requests


# Setup bot and OpenAI API credentials
bot_token = os.environ.get("5961433854:AAGwQuLvG2Wm6SQw-dPXUbLPrwewNITYZ0U")
#openai.api_key = os.environ.get("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")
#bot = telegram.Bot(token='5961433854:AAGwQuLvG2Wm6SQw-dPXUbLPrwewNITYZ0U')

# Define greeting message for new users
greeting_message = "Halo! Selamat datang di bot chat. Silakan ajukan pertanyaan atau obrolan tentang apapun dan saya akan mencoba menjawabnya sebaik mungkin."

# Define function to generate response from OpenAI API
def generate_response(prompt):
    response = requests.get('https://mfarels.my.id/api/openai?text=' + prompt)
    message = response.json()["result"]
    return message

# Define function to handle incoming messages from users
def handle_message(update, context):
    # Get user input from message
    user_input = update.message.text
    # Generate response from OpenAI API
    response = generate_response(user_input)
    # Send response back to user
    bot.send_message(chat_id=update.effective_chat.id, text=response)

# Define function to greet new users
def greet_user(update, context):
    # Send greeting message to user
    bot.send_message(chat_id=update.effective_chat.id, text=greeting_message)

# Create Telegram updater and dispatcher
bot = telegram.Bot(token='5961433854:AAGwQuLvG2Wm6SQw-dPXUbLPrwewNITYZ0U')
updater = Updater(bot=bot)

# Register message handler
dispatcher = updater.dispatcher
message_handler = MessageHandler(Filters.text & (~Filters.command), handle_message)
dispatcher.add_handler(message_handler)

# Register greeting handler
greeting_handler = CommandHandler("start", greet_user)
dispatcher.add_handler(greeting_handler)

# Start polling for incoming messages
updater.start_polling()
updater.idle()
