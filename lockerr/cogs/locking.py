from discord.ext import commands
import discord

from statics import PREFIX, PYTHON_VERSION, VERSION


class Locking(commands.Cog):
    def __init__(self, bot) -> None:
        self.bot = bot

    @commands.command()
    async def lock(self, ctx):
        pass

    @commands.command()
    async def permlock(self, ctx):
        pass

    @commands.command()
    async def unlock(self, ctx):
        pass
