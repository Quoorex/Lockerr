from discord.ext import commands
import discord

from lockerr.statics import (
    LOCKED_USERS,
    PERM_LOCKED_USERS,
)
from lockerr.utils import mention_to_member
from lockerr.permissions import is_admin


class Locking(commands.Cog):
    def __init__(self, bot: commands.AutoShardedBot) -> None:
        self.bot = bot

    @commands.command(help="Locks a user into a voice channel until he disconnects.")
    @commands.check(is_admin)
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
    @commands.check(is_admin)
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

    @commands.command(help="Unlocks a permanently or temporarily locked user.")
    @commands.check(is_admin)
    async def unlock(self, ctx, user_mention):
        member = mention_to_member(self.bot, ctx.guild.id, user_mention)
        if member in PERM_LOCKED_USERS.keys():
            PERM_LOCKED_USERS.pop(member)
        if member in LOCKED_USERS.keys():
            LOCKED_USERS.pop(member)
        embed = discord.Embed(
            title="",
            description=f"{member} was unlocked.",
            color=discord.Color.dark_red(),
        )
        await ctx.send(embed=embed)
