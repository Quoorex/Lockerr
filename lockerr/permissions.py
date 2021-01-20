async def is_admin(ctx):
    if ctx.message.author.guild_permissions.administrator:
        return True
    else:
        return False