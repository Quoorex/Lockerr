from discord.ext import commands
import discord

from statics import PREFIX, PYTHON_VERSION, VERSION


class General(commands.Cog):
    def __init__(self, bot) -> None:
        self.bot = bot

    @commands.command()
    async def about(self, ctx):
        embed = discord.Embed(
            title=f"Lockerr Discord Bot v{VERSION}",
            description=f"made by Zoore#7255 using Discord.py by Rapptz\n Use {PREFIX}help for a list of all commands.\n Python Version: {PYTHON_VERSION}",
            color=discord.Color.dark_red(),
        )
        await ctx.send(embed=embed)
