import os

import discord
from discord.ext import commands

from .cogs.general import General
from .cogs.locking import Locking
from .events import init_events

# Load .env file.
from dotenv import load_dotenv

load_dotenv()

intents = discord.Intents.default()
intents.typing = False
intents.presences = False
intents.members = True

PREFIX = "lockerr!"
activity = discord.Game(f"{PREFIX}help")
bot = commands.AutoShardedBot(
    command_prefix=PREFIX, intents=intents, case_insensitive=True, activity=activity
)


class Lockerr:
    def __init__(self) -> None:
        bot.add_cog(General(bot))
        bot.add_cog(Locking(bot))

        init_events(bot)

    def run(self):
        bot.run(os.getenv("BOT_TOKEN"), reconnect=True)


if __name__ == "__main__":
    Lockerr().run()
