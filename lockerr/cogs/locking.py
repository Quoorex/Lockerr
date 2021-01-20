from discord.ext import commands
import discord

from lockerr.statics import (
    PREFIX,
    PYTHON_VERSION,
    VERSION,
    LOCKED_USERS,
    PERM_LOCKED_USERS,
)
from lockerr.utils import mention_to_member


class Locking(commands.Cog):
    def __init__(self, bot: commands.AutoShardedBot) -> None:
        self.bot = bot

    @commands.command(help="Locks a user into a voice channel until he disconnects.")
    async def lock(self, ctx, user_mention: str):
        member = mention_to_member(self.bot, ctx.guild.id, user_mention)
        if member.voice is None:
            embed = discord.Embed(
                title="",
                description="The mentioned user has to be in a voice channel.",
                color=discord.Color.dark_red(),
            )
            await ctx.send(embed=embed)
        elif isinstance(member.voice, discord.VoiceState):
            LOCKED_USERS[member] = member.voice.channel
            embed = discord.Embed(
                title="",
                description=f"{member} was locked.",
                color=discord.Color.dark_red(),
            )
            await ctx.send(embed=embed)

    @commands.command(help="Locks a user into a voice channel until he is unlocked.")
    async def permlock(self, ctx, user_mention):
        member = mention_to_member(self.bot, ctx.guild.id, user_mention)
        if member.voice is None:
            embed = discord.Embed(
                title="",
                description="The mentioned user has to be in a voice channel.",
                color=discord.Color.dark_red(),
            )
            await ctx.send(embed=embed)
        elif isinstance(member.voice, discord.VoiceState):
            PERM_LOCKED_USERS[member] = member.voice.channel
            embed = discord.Embed(
                title="",
                description=f"{member} was permanently locked.",
                color=discord.Color.dark_red(),
            )
            await ctx.send(embed=embed)

    @commands.command(help="Unlocks a permanently locked user.")
    async def unlock(self, ctx, user_mention):
        member = mention_to_member(self.bot, ctx.guild.id, user_mention)
        PERM_LOCKED_USERS.pop(member)
        embed = discord.Embed(
            title="",
            description=f"{member} was unlocked.",
            color=discord.Color.dark_red(),
        )
        await ctx.send(embed=embed)
