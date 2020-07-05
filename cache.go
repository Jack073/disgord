package disgord

type Cache interface {
	// Gateway events
	ChannelCreate(data []byte) (interface{}, error)
	ChannelDelete(data []byte) (interface{}, error)
	ChannelPinsUpdate(data []byte) (interface{}, error)
	ChannelUpdate(data []byte) (interface{}, error)
	GuildBanAdd(data []byte) (interface{}, error)
	GuildBanRemove(data []byte) (interface{}, error)
	GuildCreate(data []byte) (interface{}, error)
	GuildDelete(data []byte) (interface{}, error)
	GuildEmojisUpdate(data []byte) (interface{}, error)
	GuildIntegrationsUpdate(data []byte) (interface{}, error)
	GuildMemberAdd(data []byte) (interface{}, error)
	GuildMemberRemove(data []byte) (interface{}, error)
	GuildMemberUpdate(data []byte) (interface{}, error)
	GuildMembersChunk(data []byte) (interface{}, error)
	GuildRoleCreate(data []byte) (interface{}, error)
	GuildRoleDelete(data []byte) (interface{}, error)
	GuildRoleUpdate(data []byte) (interface{}, error)
	GuildUpdate(data []byte) (interface{}, error)
	InviteCreate(data []byte) (interface{}, error)
	InviteDelete(data []byte) (interface{}, error)
	MessageCreate(data []byte) (interface{}, error)
	MessageDelete(data []byte) (interface{}, error)
	MessageDeleteBulk(data []byte) (interface{}, error)
	MessageReactionAdd(data []byte) (interface{}, error)
	MessageReactionRemove(data []byte) (interface{}, error)
	MessageReactionRemoveAll(data []byte) (interface{}, error)
	MessageUpdate(data []byte) (interface{}, error)
	PresenceUpdate(data []byte) (interface{}, error)
	Ready(data []byte) (interface{}, error)
	Resumed(data []byte) (interface{}, error)
	TypingStart(data []byte) (interface{}, error)
	UserUpdate(data []byte) (interface{}, error)
	VoiceServerUpdate(data []byte) (interface{}, error)
	VoiceStateUpdate(data []byte) (interface{}, error)
	WebhooksUpdate(data []byte) (interface{}, error)

	// REST API
	// GetGuildAuditLogs(guildID Snowflake) *guildAuditLogsBuilder // TODO
	GetMessages(channelID Snowflake, params *GetMessagesParams) ([]*Message, error)
	GetMessage(channelID, messageID Snowflake) (ret *Message, err error)
	GetUsersWhoReacted(channelID, messageID Snowflake, emoji interface{}, params URLQueryStringer) (reactors []*User, err error)
	GetPinnedMessages(channelID Snowflake) (ret []*Message, err error)
	GetChannel(id Snowflake) (ret *Channel, err error)
	GetChannelInvites(id Snowflake) (ret []*Invite, err error)
	GetGuildEmoji(guildID, emojiID Snowflake) (*Emoji, error)
	GetGuildEmojis(id Snowflake) ([]*Emoji, error)
	GetGuild(id Snowflake) (*Guild, error)
	GetGuildChannels(id Snowflake) ([]*Channel, error)
	GetMember(guildID, userID Snowflake) (*Member, error)
	GetMembers(guildID Snowflake, params *GetMembersParams) ([]*Member, error)
	GetGuildBans(id Snowflake) ([]*Ban, error)
	GetGuildBan(guildID, userID Snowflake) (*Ban, error)
	GetGuildRoles(guildID Snowflake) ([]*Role, error)
	GetMemberPermissions(guildID, userID Snowflake) (permissions PermissionBit, err error)
	GetGuildVoiceRegions(id Snowflake) ([]*VoiceRegion, error)
	GetGuildInvites(id Snowflake) ([]*Invite, error)
	GetGuildIntegrations(id Snowflake) ([]*Integration, error)
	GetGuildEmbed(guildID Snowflake) (*GuildEmbed, error)
	GetGuildVanityURL(guildID Snowflake) (*PartialInvite, error)
	GetInvite(inviteCode string, params URLQueryStringer) (*Invite, error)
	GetCurrentUser() (*User, error)
	GetUser(id Snowflake) (*User, error)
	GetCurrentUserGuilds(params *GetCurrentUserGuildsParams) (ret []*PartialGuild, err error)
	GetUserDMs() (ret []*Channel, err error)
	GetUserConnections() (ret []*UserConnection, err error)
	GetVoiceRegions() ([]*VoiceRegion, error)
	GetChannelWebhooks(channelID Snowflake) (ret []*Webhook, err error)
	GetGuildWebhooks(guildID Snowflake) (ret []*Webhook, err error)
	GetWebhook(id Snowflake) (ret *Webhook, err error)
}

type CacheNop struct{}

var _ Cache = (*CacheNop)(nil)

func (c *CacheNop) ChannelCreate(_ []byte) (interface{}, error)            { return nil, nil }
func (c *CacheNop) ChannelDelete(_ []byte) (interface{}, error)            { return nil, nil }
func (c *CacheNop) ChannelPinsUpdate(_ []byte) (interface{}, error)        { return nil, nil }
func (c *CacheNop) ChannelUpdate(_ []byte) (interface{}, error)            { return nil, nil }
func (c *CacheNop) GuildBanAdd(_ []byte) (interface{}, error)              { return nil, nil }
func (c *CacheNop) GuildBanRemove(_ []byte) (interface{}, error)           { return nil, nil }
func (c *CacheNop) GuildCreate(_ []byte) (interface{}, error)              { return nil, nil }
func (c *CacheNop) GuildDelete(_ []byte) (interface{}, error)              { return nil, nil }
func (c *CacheNop) GuildEmojisUpdate(_ []byte) (interface{}, error)        { return nil, nil }
func (c *CacheNop) GuildIntegrationsUpdate(_ []byte) (interface{}, error)  { return nil, nil }
func (c *CacheNop) GuildMemberAdd(_ []byte) (interface{}, error)           { return nil, nil }
func (c *CacheNop) GuildMemberRemove(_ []byte) (interface{}, error)        { return nil, nil }
func (c *CacheNop) GuildMemberUpdate(_ []byte) (interface{}, error)        { return nil, nil }
func (c *CacheNop) GuildMembersChunk(_ []byte) (interface{}, error)        { return nil, nil }
func (c *CacheNop) GuildRoleCreate(_ []byte) (interface{}, error)          { return nil, nil }
func (c *CacheNop) GuildRoleDelete(_ []byte) (interface{}, error)          { return nil, nil }
func (c *CacheNop) GuildRoleUpdate(_ []byte) (interface{}, error)          { return nil, nil }
func (c *CacheNop) GuildUpdate(_ []byte) (interface{}, error)              { return nil, nil }
func (c *CacheNop) InviteCreate(_ []byte) (interface{}, error)             { return nil, nil }
func (c *CacheNop) InviteDelete(_ []byte) (interface{}, error)             { return nil, nil }
func (c *CacheNop) MessageCreate(_ []byte) (interface{}, error)            { return nil, nil }
func (c *CacheNop) MessageDelete(_ []byte) (interface{}, error)            { return nil, nil }
func (c *CacheNop) MessageDeleteBulk(_ []byte) (interface{}, error)        { return nil, nil }
func (c *CacheNop) MessageReactionAdd(_ []byte) (interface{}, error)       { return nil, nil }
func (c *CacheNop) MessageReactionRemove(_ []byte) (interface{}, error)    { return nil, nil }
func (c *CacheNop) MessageReactionRemoveAll(_ []byte) (interface{}, error) { return nil, nil }
func (c *CacheNop) MessageUpdate(_ []byte) (interface{}, error)            { return nil, nil }
func (c *CacheNop) PresenceUpdate(_ []byte) (interface{}, error)           { return nil, nil }
func (c *CacheNop) Ready(_ []byte) (interface{}, error)                    { return nil, nil }
func (c *CacheNop) Resumed(_ []byte) (interface{}, error)                  { return nil, nil }
func (c *CacheNop) TypingStart(_ []byte) (interface{}, error)              { return nil, nil }
func (c *CacheNop) UserUpdate(_ []byte) (interface{}, error)               { return nil, nil }
func (c *CacheNop) VoiceServerUpdate(_ []byte) (interface{}, error)        { return nil, nil }
func (c *CacheNop) VoiceStateUpdate(_ []byte) (interface{}, error)         { return nil, nil }
func (c *CacheNop) WebhooksUpdate(_ []byte) (interface{}, error)           { return nil, nil }
