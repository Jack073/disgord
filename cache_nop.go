package disgord

type CacheNop struct{}

var _ CacheUpdater = (*CacheNop)(nil)
var _ CacheGetter = (*CacheNop)(nil)
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

func (c *CacheNop) GetMessages(channel Snowflake, p *GetMessagesParams) ([]*Message, error) {
	return nil, nil
}
func (c *CacheNop) GetMessage(channelID, messageID Snowflake) (*Message, error) { return nil, nil }
func (c *CacheNop) GetChannel(id Snowflake) (*Channel, error)                   { return nil, nil }
func (c *CacheNop) GetGuildEmoji(guildID, emojiID Snowflake) (*Emoji, error)    { return nil, nil }
func (c *CacheNop) GetGuildEmojis(id Snowflake) ([]*Emoji, error)               { return nil, nil }
func (c *CacheNop) GetGuild(id Snowflake) (*Guild, error)                       { return nil, nil }
func (c *CacheNop) GetGuildChannels(id Snowflake) ([]*Channel, error)           { return nil, nil }
func (c *CacheNop) GetMember(guildID, userID Snowflake) (*Member, error)        { return nil, nil }
func (c *CacheNop) GetMembers(guildID Snowflake, p *GetMembersParams) ([]*Member, error) {
	return nil, nil
}
func (c *CacheNop) GetGuildRoles(guildID Snowflake) ([]*Role, error) { return nil, nil }
func (c *CacheNop) GetCurrentUser() (*User, error)                   { return nil, nil }
func (c *CacheNop) GetUser(id Snowflake) (*User, error)              { return nil, nil }
func (c *CacheNop) GetCurrentUserGuilds(p *GetCurrentUserGuildsParams) ([]*PartialGuild, error) {
	return nil, nil
}
