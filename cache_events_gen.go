package disgord

// Code generated - This file has been automatically generated by generate/events/main.go - DO NOT EDIT.
// Warning: This file is overwritten at "go generate", instead adapt events.go and event/cache.go and run go generate

// cache dispatcher

func cacheDispatcher(c Cache, event string, data []byte) (obj interface{}, err error) {
	switch event {
	case EvtChannelCreate:
		obj, err = c.ChannelCreate(data)
	case EvtChannelDelete:
		obj, err = c.ChannelDelete(data)
	case EvtChannelPinsUpdate:
		obj, err = c.ChannelPinsUpdate(data)
	case EvtChannelUpdate:
		obj, err = c.ChannelUpdate(data)
	case EvtGuildBanAdd:
		obj, err = c.GuildBanAdd(data)
	case EvtGuildBanRemove:
		obj, err = c.GuildBanRemove(data)
	case EvtGuildCreate:
		obj, err = c.GuildCreate(data)
	case EvtGuildDelete:
		obj, err = c.GuildDelete(data)
	case EvtGuildEmojisUpdate:
		obj, err = c.GuildEmojisUpdate(data)
	case EvtGuildIntegrationsUpdate:
		obj, err = c.GuildIntegrationsUpdate(data)
	case EvtGuildMemberAdd:
		obj, err = c.GuildMemberAdd(data)
	case EvtGuildMemberRemove:
		obj, err = c.GuildMemberRemove(data)
	case EvtGuildMemberUpdate:
		obj, err = c.GuildMemberUpdate(data)
	case EvtGuildMembersChunk:
		obj, err = c.GuildMembersChunk(data)
	case EvtGuildRoleCreate:
		obj, err = c.GuildRoleCreate(data)
	case EvtGuildRoleDelete:
		obj, err = c.GuildRoleDelete(data)
	case EvtGuildRoleUpdate:
		obj, err = c.GuildRoleUpdate(data)
	case EvtGuildUpdate:
		obj, err = c.GuildUpdate(data)
	case EvtInviteCreate:
		obj, err = c.InviteCreate(data)
	case EvtInviteDelete:
		obj, err = c.InviteDelete(data)
	case EvtMessageCreate:
		obj, err = c.MessageCreate(data)
	case EvtMessageDelete:
		obj, err = c.MessageDelete(data)
	case EvtMessageDeleteBulk:
		obj, err = c.MessageDeleteBulk(data)
	case EvtMessageReactionAdd:
		obj, err = c.MessageReactionAdd(data)
	case EvtMessageReactionRemove:
		obj, err = c.MessageReactionRemove(data)
	case EvtMessageReactionRemoveAll:
		obj, err = c.MessageReactionRemoveAll(data)
	case EvtMessageUpdate:
		obj, err = c.MessageUpdate(data)
	case EvtPresenceUpdate:
		obj, err = c.PresenceUpdate(data)
	case EvtReady:
		obj, err = c.Ready(data)
	case EvtResumed:
		obj, err = c.Resumed(data)
	case EvtTypingStart:
		obj, err = c.TypingStart(data)
	case EvtUserUpdate:
		obj, err = c.UserUpdate(data)
	case EvtVoiceServerUpdate:
		obj, err = c.VoiceServerUpdate(data)
	case EvtVoiceStateUpdate:
		obj, err = c.VoiceStateUpdate(data)
	case EvtWebhooksUpdate:
		obj, err = c.WebhooksUpdate(data)
	default:
		obj = nil
		err = nil
	}

	return obj, err
}

// updaters

type CacheUpdater interface {
	// Gateway events
	ChannelCreate(data []byte) (*ChannelCreate, error)
	ChannelDelete(data []byte) (*ChannelDelete, error)
	ChannelPinsUpdate(data []byte) (*ChannelPinsUpdate, error)
	ChannelUpdate(data []byte) (*ChannelUpdate, error)
	GuildBanAdd(data []byte) (*GuildBanAdd, error)
	GuildBanRemove(data []byte) (*GuildBanRemove, error)
	GuildCreate(data []byte) (*GuildCreate, error)
	GuildDelete(data []byte) (*GuildDelete, error)
	GuildEmojisUpdate(data []byte) (*GuildEmojisUpdate, error)
	GuildIntegrationsUpdate(data []byte) (*GuildIntegrationsUpdate, error)
	GuildMemberAdd(data []byte) (*GuildMemberAdd, error)
	GuildMemberRemove(data []byte) (*GuildMemberRemove, error)
	GuildMemberUpdate(data []byte) (*GuildMemberUpdate, error)
	GuildMembersChunk(data []byte) (*GuildMembersChunk, error)
	GuildRoleCreate(data []byte) (*GuildRoleCreate, error)
	GuildRoleDelete(data []byte) (*GuildRoleDelete, error)
	GuildRoleUpdate(data []byte) (*GuildRoleUpdate, error)
	GuildUpdate(data []byte) (*GuildUpdate, error)
	InviteCreate(data []byte) (*InviteCreate, error)
	InviteDelete(data []byte) (*InviteDelete, error)
	MessageCreate(data []byte) (*MessageCreate, error)
	MessageDelete(data []byte) (*MessageDelete, error)
	MessageDeleteBulk(data []byte) (*MessageDeleteBulk, error)
	MessageReactionAdd(data []byte) (*MessageReactionAdd, error)
	MessageReactionRemove(data []byte) (*MessageReactionRemove, error)
	MessageReactionRemoveAll(data []byte) (*MessageReactionRemoveAll, error)
	MessageUpdate(data []byte) (*MessageUpdate, error)
	PresenceUpdate(data []byte) (*PresenceUpdate, error)
	Ready(data []byte) (*Ready, error)
	Resumed(data []byte) (*Resumed, error)
	TypingStart(data []byte) (*TypingStart, error)
	UserUpdate(data []byte) (*UserUpdate, error)
	VoiceServerUpdate(data []byte) (*VoiceServerUpdate, error)
	VoiceStateUpdate(data []byte) (*VoiceStateUpdate, error)
	WebhooksUpdate(data []byte) (*WebhooksUpdate, error)
}
