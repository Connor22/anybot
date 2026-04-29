var (
	commands = []discordgo.ApplicationCommand{
		Name: "report",
		Type: discordgo.MessageApplicationCommand,
		Description: "Report this message to the moderators.",
	},
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"report": reportMessageWithModals,
		}	
	}
)


func init(){
	for id, name := range cmdIDs {
		err := s.ApplicationCommandDelete(*AppID, *GuildID, id)
		if err != nil {
			log.Fatalf("Cannot delete slash command %q: %v", name, err)
		}
	}
}

func 
	cmdIDs := make(map[string]string, len(commands))

	for _, cmd := range commands {
		rcmd, err := s.ApplicationCommandCreate(*AppID, *GuildID, &cmd)
		if err != nil {
			log.Fatalf("Cannot create slash command %q: %v", cmd.Name, err)
		}

		cmdIDs[rcmd.ID] = rcmd.Name

	}
}

func reportMessageWithModals(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "report_" + i.Interaction.Member.ID + "_" + i.Interaction.Message.ID,
			Title:    "Report a Message",
			Flags:    discordgo.MessageFlagsIsComponentsV2,
			Components: []discordgo.MessageComponent{
				discordgo.Label{
					Label:       "Report Reason",
					Description: "Why are you reporting this message?",
					Component: discordgo.TextInput{
						CustomID:  "reason",
						Style:     discordgo.TextInputParagraph,
						Required:  true,
						MaxLength: 2000,
					},
				},
				discordgo.Label{
					Label:       "Is there any context the moderators should know about?",
					Description: "Examples include previous interactions or history with the user, or meanings in the message the moderators may not be aware of.",
					Component: discordgo.TextInput{
						CustomID:  "context",
						Style:     discordgo.TextInputParagraph,
						Required:  false,
						MaxLength: 2000,
					},
				},
			},
		},
	})
}