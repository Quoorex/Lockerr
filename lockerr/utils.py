import discord
from discord.ext import commands


def mention_to_member(
    bot: commands.AutoShardedBot, guild_id: int, user_mention: str
) -> discord.User:
    """
    Turns a discord user mention into a discord.py member object.
    """
    user_id = int(user_mention[3:].split(">")[0])
    return bot.get_guild(guild_id).get_member(user_id)
