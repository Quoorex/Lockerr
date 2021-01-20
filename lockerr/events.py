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
        print(member)
