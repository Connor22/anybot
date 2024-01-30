# ANyBot

Anybot is a Discord bot written using [discordgo](https://pkg.go.dev/github.com/bwmarrin/discordgo) for the AnimeNorth server

## Currently implemented features

- Resolves conflicts with Join Roles and Verify Roles
- Assigns users a JoinRole when they first join a server
- Does not store unnecessary data, so does not need restarts or maintenance
- Handles downtime and restarts

## Planned features

- Database to enable dynamically editing join and verify roles
- Sync roles to google sheets