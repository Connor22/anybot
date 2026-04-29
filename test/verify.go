package modules

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type VerificationMod struct {
	flag uint8
	name string `default:"VerificationMod"`
}

func init() {
	verifyInteractionCommand = discordgo.ApplicationCommand{
		Name:        "verify",
		Description: "Mark user as trusted and assign the Attendee role",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "User to verify",
				Required:    true,
			},
		},
		DefaultMemberPermissions: discordgo.PermissionManageMessages,
	}
}

func (vermod *VerificationMod) slashHandler(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	data := interaction.ApplicationCommandData()
	if data.name == "verify" {
		// Access options in the order provided by the user.
		options := data().Options

		// Or convert the slice into a map
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		// This example stores the provided arguments in an []interface{}
		// which will be used to format the bot's response
		margs := make([]interface{}, 0, len(options))

		if opt, ok := optionMap["user-option"]; ok {
			margs = append(margs, opt.UserValue(nil).ID)
			msgformat += "> verified: <@%s>\n"
		}

		discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(
					msgformat,
					margs...,
				),
			},
		})
	}
}

func (vermod *VerificationMod) init(discord *discordgo.Session) {
	var verifyInteractionCommand = discordgo.ApplicationCommand{
		Name:        "verify",
		Description: "Mark user as trusted and assign the Attendee role",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "User to verify",
				Required:    true,
			},
		},
		DefaultMemberPermissions: int64(&discordgo.PermissionManageMessages),
	}

	discord.ApplicationCommandCreate(discord.State.User.ID, *GuildID, verifyInteractionCommand)
}
