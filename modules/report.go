package modules

import (
	"anybot/conf"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []discordgo.ApplicationCommand{
		{
			Name:        "report",
			Type:        discordgo.MessageApplicationCommand,
			Description: "Report this message to the moderators.",
		},
	}
)
var (
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"report": reportMessageWithModals,
	}
)

var cmdIDs = make(map[string]string, len(commands))

type ReportMod struct {
	flag uint8
	name string `default:"Template"`
}

func (tempmod *ReportMod) Init(modid int) {
	tempmod.flag = (1 << modid)

	return
}

func (tempmod *ReportMod) Start(discord *discordgo.Session, appID string) {
	// TODO: Store guilds we created Interactions in to later tear down in Stop()
	// Also TODO abstract this code into handlers.go or similar to allow modules to register new ApplicationCommands
	for _, server := range discord.State.Guilds {
		for _, cmd := range commands {
			rcmd, err := discord.ApplicationCommandCreate(appID, server.ID, &cmd)
			if err != nil {
				log.Fatalf("Cannot create slash command %q: %v", cmd.Name, err)
			}

			cmdIDs[rcmd.ID] = rcmd.Name
		}
	}

	return
}

func (tempmod *ReportMod) Stop(discord *discordgo.Session, appID string) {
	for _, server := range discord.State.Guilds {
		for id, name := range cmdIDs {
			err := discord.ApplicationCommandDelete(appID, server.ID, id)
			if err != nil {
				log.Fatalf("Cannot delete slash command %q: %v", name, err)
			}
		}
	}

	return
}

func (tempmod *ReportMod) Name() string {
	return tempmod.name
}

func (tempmod *ReportMod) Flag() uint8 {
	return tempmod.flag
}

func (tempmod *ReportMod) Intents() discordgo.Intent {
	intents := *new(discordgo.Intent)

	return intents
}

func (tempmod *ReportMod) Enabled(serverFlags uint8) bool {
	return tempmod.flag&serverFlags != 0
}

// no-op
// function is run when a user joins the server
func (tempmod *ReportMod) OnNewMember(guildMember *discordgo.GuildMemberAdd, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	return
}

// no-op
// function fires when connecting to a server on boot
func (tempmod *ReportMod) OnGuildConnect(guildConnection *discordgo.GuildCreate, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	return
}

// no-op
// function is run for each visible member when connecting to a server on boot
func (tempmod *ReportMod) OnGuildConnectMember(guildMember *discordgo.Member, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	return
}

// no-op
// function is run when a member's roles or permissions are updated
func (tempmod *ReportMod) OnMemberUpdate(guildMember *discordgo.GuildMemberUpdate, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	return
}

func reportMessageWithModals(discord *discordgo.Session, newInteraction *discordgo.InteractionCreate) {
	err := discord.InteractionRespond(newInteraction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "report_" + newInteraction.Interaction.Member.User.ID + "_" + newInteraction.Interaction.Message.ID,
			Title:    "Report a Message",
			Flags:    discordgo.MessageFlagsIsComponentsV2,
			Components: []discordgo.MessageComponent{
				discordgo.TextInput{
					Label:       "Report Reason",
					Placeholder: "Why are you reporting this message?",
					CustomID:    "reason",
					Style:       discordgo.TextInputParagraph,
					Required:    true,
					MaxLength:   2000,
				},
				discordgo.TextInput{
					Label:       "Is there any context the moderators should know about?",
					Placeholder: "Examples include previous interactions or history with the user, or meanings in the message the moderators may not be aware of.",
					CustomID:    "context",
					Style:       discordgo.TextInputParagraph,
					Required:    false,
					MaxLength:   2000,
				},
			},
		},
	})

	if err != nil {
		log.Fatalf("Cannot complete report command: %v", err)
	}
}
