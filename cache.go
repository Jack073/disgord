package disgord

type CacheGetter interface {
	// REST API

	// GetGuildAuditLogs(guildID Snowflake) *guildAuditLogsBuilder // TODO
	GetMessages(channelID Snowflake, params *GetMessagesParams) ([]*Message, error)
	GetMessage(channelID, messageID Snowflake) (ret *Message, err error)
	//GetUsersWhoReacted(channelID, messageID Snowflake, emoji interface{}, params URLQueryStringer) (reactors []*User, err error)
	//GetPinnedMessages(channelID Snowflake) (ret []*Message, err error)
	GetChannel(id Snowflake) (ret *Channel, err error)
	//GetChannelInvites(id Snowflake) (ret []*Invite, err error)
	GetGuildEmoji(guildID, emojiID Snowflake) (*Emoji, error)
	GetGuildEmojis(id Snowflake) ([]*Emoji, error)
	GetGuild(id Snowflake) (*Guild, error)
	GetGuildChannels(id Snowflake) ([]*Channel, error)
	GetMember(guildID, userID Snowflake) (*Member, error)
	GetMembers(guildID Snowflake, params *GetMembersParams) ([]*Member, error)
	//GetGuildBans(id Snowflake) ([]*Ban, error)
	//GetGuildBan(guildID, userID Snowflake) (*Ban, error)
	GetGuildRoles(guildID Snowflake) ([]*Role, error)
	//GetMemberPermissions(guildID, userID Snowflake) (permissions PermissionBit, err error)
	//GetGuildVoiceRegions(id Snowflake) ([]*VoiceRegion, error)
	//GetGuildInvites(id Snowflake) ([]*Invite, error)
	//GetGuildIntegrations(id Snowflake) ([]*Integration, error)
	//GetGuildEmbed(guildID Snowflake) (*GuildEmbed, error)
	//GetGuildVanityURL(guildID Snowflake) (*PartialInvite, error)
	//GetInvite(inviteCode string, params URLQueryStringer) (*Invite, error)
	GetCurrentUser() (*User, error)
	GetUser(id Snowflake) (*User, error)
	GetCurrentUserGuilds(params *GetCurrentUserGuildsParams) (ret []*PartialGuild, err error)
	//GetUserDMs() (ret []*Channel, err error)
	//GetUserConnections() (ret []*UserConnection, err error)
	//GetVoiceRegions() ([]*VoiceRegion, error)
	//GetChannelWebhooks(channelID Snowflake) (ret []*Webhook, err error)
	//GetGuildWebhooks(guildID Snowflake) (ret []*Webhook, err error)
	//GetWebhook(id Snowflake) (ret *Webhook, err error)
}

// Cache interface for event handling and REST methods
// commented out fields are simply not supported yet. PR's are welcome
//
// Note that on events you are expected to return a unmarshalled object. For delete methods
// you should return nil, and a nil error if the objected to be deleted was not found (nop!).
// Note that the error might change to a "CacheMiss" or something similar such that we can
//  get more metrics!
type Cache interface {
	CacheUpdater // generated
	CacheGetter
}
