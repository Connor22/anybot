//go:build exclude

package modules

import (
	"anybot/conf"

	"github.com/bwmarrin/discordgo"
)

type LogMod struct {
	flag uint8
	name string `default:"Template"`
}

func (logmod *LogMod) Init(modid int) {
	logmod.flag = (1 << modid)

	return
}

func (logmod *LogMod) Start() {
	return
}

func (logmod *LogMod) Name() string {
	return logmod.name
}

func (logmod *LogMod) Flag() uint8 {
	return logmod.flag
}

func (logmod *LogMod) Intents() discordgo.Intent {
	intents := *new(discordgo.Intent)

	return intents
}

func (logmod *LogMod) Enabled(serverFlags uint8) bool {
	return logmod.flag&serverFlags != 0
}

// no-op
func (logmod *LogMod) OnNewMember(guildMember *discordgo.GuildMemberAdd, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	return
}

// no-op
func (logmod *LogMod) OnGuildConnect(guildConnection *discordgo.GuildCreate, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	return
}

// no-op
func (logmod *LogMod) OnGuildConnectMember(guildMember *discordgo.Member, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	return
}

// no-op
func (logmod *LogMod) OnMemberUpdate(guildMember *discordgo.GuildMemberUpdate, discord *discordgo.Session, serverConfig *conf.AnyGuild) {
	return
}
