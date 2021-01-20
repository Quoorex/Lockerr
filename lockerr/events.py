from .statics import LOCKED_USERS, PERM_LOCKED_USERS


def init_events(bot):
    @bot.event
    async def on_connect():
        print(f"Lockerr is now running.")

    @bot.event
    async def on_voice_state_update(member, before, after):
        """
        Runs checks to see if the user is locked and should be moved or
        left a channel and should be unlocked (when temporarely locked).
        """
        # Check if the user is locked at all.
        if not (member in LOCKED_USERS.keys() or member in PERM_LOCKED_USERS.keys()):
            return

        # User disconnects from a channel.
        if after.channel is None:
            # Remove the user from the list of temp locked users.
            LOCKED_USERS.pop(member)

        # Get the locked user's channel
        channel = None
        if member in LOCKED_USERS.keys():
            channel = LOCKED_USERS[member]
        elif member in PERM_LOCKED_USERS.keys():
            channel = PERM_LOCKED_USERS[member]

        # User joined a different voice channel and has to be moved.
        if channel is not None and after.channel.id != channel.id:
            await member.move_to(channel, reason="User is locked by Lockerr.")
