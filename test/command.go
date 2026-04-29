var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "options",
		Description: "Command for demonstrating options",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user-option",
				Description: "User option",
				Required:    false,
			},
		}
	}
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"options": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options

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
				msgformat += "> user-option: <@%s>\n"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						msgformat,
						margs...,
					),
				},
			})
		},
	},
	"responses": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Responses to a command are very important.
		// First of all, because you need to react to the interaction
		// by sending the response in 3 seconds after receiving, otherwise
		// interaction will be considered invalid and you can no longer
		// use the interaction token and ID for responding to the user's request

		content := ""
		// As you can see, the response type names used here are pretty self-explanatory,
		// but for those who want more information see the official documentation
		switch i.ApplicationCommandData().Options[0].IntValue() {
		case int64(discordgo.InteractionResponseChannelMessageWithSource):
			content =
				"You just responded to an interaction, sent a message and showed the original one. " +
					"Congratulations!"
			content +=
				"\nAlso... you can edit your response, wait 5 seconds and this message will be changed"
		default:
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseType(i.ApplicationCommandData().Options[0].IntValue()),
			})
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: "Something went wrong",
				})
			}
			return
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseType(i.ApplicationCommandData().Options[0].IntValue()),
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
		if err != nil {
			s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "Something went wrong",
			})
			return
		}
		time.AfterFunc(time.Second*5, func() {
			content := content + "\n\nWell, now you know how to create and edit responses. " +
				"But you still don't know how to delete them... so... wait 10 seconds and this " +
				"message will be deleted."
			_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: &content,
			})
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: "Something went wrong",
				})
				return
			}
			time.Sleep(time.Second * 10)
			s.InteractionResponseDelete(i.Interaction)
		})
	},
	"followups": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Followup messages are basically regular messages (you can create as many of them as you wish)
		// but work as they are created by webhooks and their functionality
		// is for handling additional messages after sending a response.

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				// Note: this isn't documented, but you can use that if you want to.
				// This flag just allows you to create messages visible only for the caller of the command
				// (user who triggered the command)
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Surprise!",
			},
		})
		msg, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Followup message has been created, after 5 seconds it will be edited",
		})
		if err != nil {
			s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "Something went wrong",
			})
			return
		}
		time.Sleep(time.Second * 5)

		content := "Now the original message is gone and after 10 seconds this message will ~~self-destruct~~ be deleted."
		s.FollowupMessageEdit(i.Interaction, msg.ID, &discordgo.WebhookEdit{
			Content: &content,
		})

		time.Sleep(time.Second * 10)

		s.FollowupMessageDelete(i.Interaction, msg.ID)

		s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "For those, who didn't skip anything and followed tutorial along fairly, " +
				"take a unicorn :unicorn: as reward!\n" +
				"Also, as bonus... look at the original interaction response :D",
		})
	},
}
