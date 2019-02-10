package disgord

// Warning: This file has been automatically generated by generate/events/main.go
// Do NOT make changes here, instead adapt events.go and event/events.go and run go generate

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// NewDispatch construct a Dispatch object for reacting to web socket events
// from discord
func NewDispatch(activateEventChannels bool, evtChanSize int) *Dispatch {
	dispatcher := &Dispatch{
		activateEventChannels: activateEventChannels,

		listeners:      make(map[string][]interface{}),
		listenOnceOnly: make(map[string][]int),

		shutdown: make(chan struct{}),
	}

	if activateEventChannels {

		dispatcher.channelCreateChan = make(chan *ChannelCreate, evtChanSize)
		dispatcher.channelDeleteChan = make(chan *ChannelDelete, evtChanSize)
		dispatcher.channelPinsUpdateChan = make(chan *ChannelPinsUpdate, evtChanSize)
		dispatcher.channelUpdateChan = make(chan *ChannelUpdate, evtChanSize)
		dispatcher.guildBanAddChan = make(chan *GuildBanAdd, evtChanSize)
		dispatcher.guildBanRemoveChan = make(chan *GuildBanRemove, evtChanSize)
		dispatcher.guildCreateChan = make(chan *GuildCreate, evtChanSize)
		dispatcher.guildDeleteChan = make(chan *GuildDelete, evtChanSize)
		dispatcher.guildEmojisUpdateChan = make(chan *GuildEmojisUpdate, evtChanSize)
		dispatcher.guildIntegrationsUpdateChan = make(chan *GuildIntegrationsUpdate, evtChanSize)
		dispatcher.guildMemberAddChan = make(chan *GuildMemberAdd, evtChanSize)
		dispatcher.guildMemberRemoveChan = make(chan *GuildMemberRemove, evtChanSize)
		dispatcher.guildMemberUpdateChan = make(chan *GuildMemberUpdate, evtChanSize)
		dispatcher.guildMembersChunkChan = make(chan *GuildMembersChunk, evtChanSize)
		dispatcher.guildRoleCreateChan = make(chan *GuildRoleCreate, evtChanSize)
		dispatcher.guildRoleDeleteChan = make(chan *GuildRoleDelete, evtChanSize)
		dispatcher.guildRoleUpdateChan = make(chan *GuildRoleUpdate, evtChanSize)
		dispatcher.guildUpdateChan = make(chan *GuildUpdate, evtChanSize)
		dispatcher.messageCreateChan = make(chan *MessageCreate, evtChanSize)
		dispatcher.messageDeleteChan = make(chan *MessageDelete, evtChanSize)
		dispatcher.messageDeleteBulkChan = make(chan *MessageDeleteBulk, evtChanSize)
		dispatcher.messageReactionAddChan = make(chan *MessageReactionAdd, evtChanSize)
		dispatcher.messageReactionRemoveChan = make(chan *MessageReactionRemove, evtChanSize)
		dispatcher.messageReactionRemoveAllChan = make(chan *MessageReactionRemoveAll, evtChanSize)
		dispatcher.messageUpdateChan = make(chan *MessageUpdate, evtChanSize)
		dispatcher.presenceUpdateChan = make(chan *PresenceUpdate, evtChanSize)
		dispatcher.presencesReplaceChan = make(chan *PresencesReplace, evtChanSize)
		dispatcher.readyChan = make(chan *Ready, evtChanSize)
		dispatcher.resumedChan = make(chan *Resumed, evtChanSize)
		dispatcher.typingStartChan = make(chan *TypingStart, evtChanSize)
		dispatcher.userUpdateChan = make(chan *UserUpdate, evtChanSize)
		dispatcher.voiceServerUpdateChan = make(chan *VoiceServerUpdate, evtChanSize)
		dispatcher.voiceStateUpdateChan = make(chan *VoiceStateUpdate, evtChanSize)
		dispatcher.webhooksUpdateChan = make(chan *WebhooksUpdate, evtChanSize)
	}

	return dispatcher
}

// Dispatch holds all the channels and internal state for all registered
// observers
type Dispatch struct {
	channelCreateChan            chan *ChannelCreate
	channelDeleteChan            chan *ChannelDelete
	channelPinsUpdateChan        chan *ChannelPinsUpdate
	channelUpdateChan            chan *ChannelUpdate
	guildBanAddChan              chan *GuildBanAdd
	guildBanRemoveChan           chan *GuildBanRemove
	guildCreateChan              chan *GuildCreate
	guildDeleteChan              chan *GuildDelete
	guildEmojisUpdateChan        chan *GuildEmojisUpdate
	guildIntegrationsUpdateChan  chan *GuildIntegrationsUpdate
	guildMemberAddChan           chan *GuildMemberAdd
	guildMemberRemoveChan        chan *GuildMemberRemove
	guildMemberUpdateChan        chan *GuildMemberUpdate
	guildMembersChunkChan        chan *GuildMembersChunk
	guildRoleCreateChan          chan *GuildRoleCreate
	guildRoleDeleteChan          chan *GuildRoleDelete
	guildRoleUpdateChan          chan *GuildRoleUpdate
	guildUpdateChan              chan *GuildUpdate
	messageCreateChan            chan *MessageCreate
	messageDeleteChan            chan *MessageDelete
	messageDeleteBulkChan        chan *MessageDeleteBulk
	messageReactionAddChan       chan *MessageReactionAdd
	messageReactionRemoveChan    chan *MessageReactionRemove
	messageReactionRemoveAllChan chan *MessageReactionRemoveAll
	messageUpdateChan            chan *MessageUpdate
	presenceUpdateChan           chan *PresenceUpdate
	presencesReplaceChan         chan *PresencesReplace
	readyChan                    chan *Ready
	resumedChan                  chan *Resumed
	typingStartChan              chan *TypingStart
	userUpdateChan               chan *UserUpdate
	voiceServerUpdateChan        chan *VoiceServerUpdate
	voiceStateUpdateChan         chan *VoiceStateUpdate
	webhooksUpdateChan           chan *WebhooksUpdate

	activateEventChannels bool

	listeners      map[string][]interface{}
	listenOnceOnly map[string][]int

	shutdown chan struct{}

	listenersLock sync.RWMutex
}

// EventChan ... TODO
func (d *Dispatch) EventChan(event string) (channel interface{}, err error) {
	if !d.activateEventChannels {
		return nil, errors.New("usage of event channels have not been activated. See disgord.Config")
	}

	switch event {

	case EventChannelCreate:
		channel = d.ChannelCreate()
	case EventChannelDelete:
		channel = d.ChannelDelete()
	case EventChannelPinsUpdate:
		channel = d.ChannelPinsUpdate()
	case EventChannelUpdate:
		channel = d.ChannelUpdate()
	case EventGuildBanAdd:
		channel = d.GuildBanAdd()
	case EventGuildBanRemove:
		channel = d.GuildBanRemove()
	case EventGuildCreate:
		channel = d.GuildCreate()
	case EventGuildDelete:
		channel = d.GuildDelete()
	case EventGuildEmojisUpdate:
		channel = d.GuildEmojisUpdate()
	case EventGuildIntegrationsUpdate:
		channel = d.GuildIntegrationsUpdate()
	case EventGuildMemberAdd:
		channel = d.GuildMemberAdd()
	case EventGuildMemberRemove:
		channel = d.GuildMemberRemove()
	case EventGuildMemberUpdate:
		channel = d.GuildMemberUpdate()
	case EventGuildMembersChunk:
		channel = d.GuildMembersChunk()
	case EventGuildRoleCreate:
		channel = d.GuildRoleCreate()
	case EventGuildRoleDelete:
		channel = d.GuildRoleDelete()
	case EventGuildRoleUpdate:
		channel = d.GuildRoleUpdate()
	case EventGuildUpdate:
		channel = d.GuildUpdate()
	case EventMessageCreate:
		channel = d.MessageCreate()
	case EventMessageDelete:
		channel = d.MessageDelete()
	case EventMessageDeleteBulk:
		channel = d.MessageDeleteBulk()
	case EventMessageReactionAdd:
		channel = d.MessageReactionAdd()
	case EventMessageReactionRemove:
		channel = d.MessageReactionRemove()
	case EventMessageReactionRemoveAll:
		channel = d.MessageReactionRemoveAll()
	case EventMessageUpdate:
		channel = d.MessageUpdate()
	case EventPresenceUpdate:
		channel = d.PresenceUpdate()
	case EventPresencesReplace:
		channel = d.PresencesReplace()
	case EventReady:
		channel = d.Ready()
	case EventResumed:
		channel = d.Resumed()
	case EventTypingStart:
		channel = d.TypingStart()
	case EventUserUpdate:
		channel = d.UserUpdate()
	case EventVoiceServerUpdate:
		channel = d.VoiceServerUpdate()
	case EventVoiceStateUpdate:
		channel = d.VoiceStateUpdate()
	case EventWebhooksUpdate:
		channel = d.WebhooksUpdate()
	default:
		err = errors.New("no event channel exists for given event: " + event)
	}

	return
}

func (d *Dispatch) triggerChan(ctx context.Context, evtName string, session Session, box interface{}) {
	if !d.activateEventChannels {
		return
	}

	switch evtName {

	case EventChannelCreate:
		d.channelCreateChan <- box.(*ChannelCreate)
	case EventChannelDelete:
		d.channelDeleteChan <- box.(*ChannelDelete)
	case EventChannelPinsUpdate:
		d.channelPinsUpdateChan <- box.(*ChannelPinsUpdate)
	case EventChannelUpdate:
		d.channelUpdateChan <- box.(*ChannelUpdate)
	case EventGuildBanAdd:
		d.guildBanAddChan <- box.(*GuildBanAdd)
	case EventGuildBanRemove:
		d.guildBanRemoveChan <- box.(*GuildBanRemove)
	case EventGuildCreate:
		d.guildCreateChan <- box.(*GuildCreate)
	case EventGuildDelete:
		d.guildDeleteChan <- box.(*GuildDelete)
	case EventGuildEmojisUpdate:
		d.guildEmojisUpdateChan <- box.(*GuildEmojisUpdate)
	case EventGuildIntegrationsUpdate:
		d.guildIntegrationsUpdateChan <- box.(*GuildIntegrationsUpdate)
	case EventGuildMemberAdd:
		d.guildMemberAddChan <- box.(*GuildMemberAdd)
	case EventGuildMemberRemove:
		d.guildMemberRemoveChan <- box.(*GuildMemberRemove)
	case EventGuildMemberUpdate:
		d.guildMemberUpdateChan <- box.(*GuildMemberUpdate)
	case EventGuildMembersChunk:
		d.guildMembersChunkChan <- box.(*GuildMembersChunk)
	case EventGuildRoleCreate:
		d.guildRoleCreateChan <- box.(*GuildRoleCreate)
	case EventGuildRoleDelete:
		d.guildRoleDeleteChan <- box.(*GuildRoleDelete)
	case EventGuildRoleUpdate:
		d.guildRoleUpdateChan <- box.(*GuildRoleUpdate)
	case EventGuildUpdate:
		d.guildUpdateChan <- box.(*GuildUpdate)
	case EventMessageCreate:
		d.messageCreateChan <- box.(*MessageCreate)
	case EventMessageDelete:
		d.messageDeleteChan <- box.(*MessageDelete)
	case EventMessageDeleteBulk:
		d.messageDeleteBulkChan <- box.(*MessageDeleteBulk)
	case EventMessageReactionAdd:
		d.messageReactionAddChan <- box.(*MessageReactionAdd)
	case EventMessageReactionRemove:
		d.messageReactionRemoveChan <- box.(*MessageReactionRemove)
	case EventMessageReactionRemoveAll:
		d.messageReactionRemoveAllChan <- box.(*MessageReactionRemoveAll)
	case EventMessageUpdate:
		d.messageUpdateChan <- box.(*MessageUpdate)
	case EventPresenceUpdate:
		d.presenceUpdateChan <- box.(*PresenceUpdate)
	case EventPresencesReplace:
		d.presencesReplaceChan <- box.(*PresencesReplace)
	case EventReady:
		d.readyChan <- box.(*Ready)
	case EventResumed:
		d.resumedChan <- box.(*Resumed)
	case EventTypingStart:
		d.typingStartChan <- box.(*TypingStart)
	case EventUserUpdate:
		d.userUpdateChan <- box.(*UserUpdate)
	case EventVoiceServerUpdate:
		d.voiceServerUpdateChan <- box.(*VoiceServerUpdate)
	case EventVoiceStateUpdate:
		d.voiceStateUpdateChan <- box.(*VoiceStateUpdate)
	case EventWebhooksUpdate:
		d.webhooksUpdateChan <- box.(*WebhooksUpdate)
	default:
		// if we land at this stage, the channel is either full or a unknown event has come through
		// empty the channel
		d.emptyChannel(evtName)
	}
}

func (d *Dispatch) emptyChannel(evtName string) {
	if !d.activateEventChannels {
		return
	}

	switch evtName {

	case EventChannelCreate:
		for range d.channelCreateChan {
		}
	case EventChannelDelete:
		for range d.channelDeleteChan {
		}
	case EventChannelPinsUpdate:
		for range d.channelPinsUpdateChan {
		}
	case EventChannelUpdate:
		for range d.channelUpdateChan {
		}
	case EventGuildBanAdd:
		for range d.guildBanAddChan {
		}
	case EventGuildBanRemove:
		for range d.guildBanRemoveChan {
		}
	case EventGuildCreate:
		for range d.guildCreateChan {
		}
	case EventGuildDelete:
		for range d.guildDeleteChan {
		}
	case EventGuildEmojisUpdate:
		for range d.guildEmojisUpdateChan {
		}
	case EventGuildIntegrationsUpdate:
		for range d.guildIntegrationsUpdateChan {
		}
	case EventGuildMemberAdd:
		for range d.guildMemberAddChan {
		}
	case EventGuildMemberRemove:
		for range d.guildMemberRemoveChan {
		}
	case EventGuildMemberUpdate:
		for range d.guildMemberUpdateChan {
		}
	case EventGuildMembersChunk:
		for range d.guildMembersChunkChan {
		}
	case EventGuildRoleCreate:
		for range d.guildRoleCreateChan {
		}
	case EventGuildRoleDelete:
		for range d.guildRoleDeleteChan {
		}
	case EventGuildRoleUpdate:
		for range d.guildRoleUpdateChan {
		}
	case EventGuildUpdate:
		for range d.guildUpdateChan {
		}
	case EventMessageCreate:
		for range d.messageCreateChan {
		}
	case EventMessageDelete:
		for range d.messageDeleteChan {
		}
	case EventMessageDeleteBulk:
		for range d.messageDeleteBulkChan {
		}
	case EventMessageReactionAdd:
		for range d.messageReactionAddChan {
		}
	case EventMessageReactionRemove:
		for range d.messageReactionRemoveChan {
		}
	case EventMessageReactionRemoveAll:
		for range d.messageReactionRemoveAllChan {
		}
	case EventMessageUpdate:
		for range d.messageUpdateChan {
		}
	case EventPresenceUpdate:
		for range d.presenceUpdateChan {
		}
	case EventPresencesReplace:
		for range d.presencesReplaceChan {
		}
	case EventReady:
		for range d.readyChan {
		}
	case EventResumed:
		for range d.resumedChan {
		}
	case EventTypingStart:
		for range d.typingStartChan {
		}
	case EventUserUpdate:
		for range d.userUpdateChan {
		}
	case EventVoiceServerUpdate:
		for range d.voiceServerUpdateChan {
		}
	case EventVoiceStateUpdate:
		for range d.voiceStateUpdateChan {
		}
	case EventWebhooksUpdate:
		for range d.webhooksUpdateChan {
		}
	}
}

func (d *Dispatch) triggerHandlers(ctx context.Context, evtName string, session Session, box interface{}) {
	d.listenersLock.RLock()
	switch evtName {

	case EventChannelCreate:
		for _, listener := range d.listeners[EventChannelCreate] {
			if cb, ok := listener.(ChannelCreateHandler); ok {
				cb(session, box.(*ChannelCreate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: ChannelCreate")
			}
		}
	case EventChannelDelete:
		for _, listener := range d.listeners[EventChannelDelete] {
			if cb, ok := listener.(ChannelDeleteHandler); ok {
				cb(session, box.(*ChannelDelete))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: ChannelDelete")
			}
		}
	case EventChannelPinsUpdate:
		for _, listener := range d.listeners[EventChannelPinsUpdate] {
			if cb, ok := listener.(ChannelPinsUpdateHandler); ok {
				cb(session, box.(*ChannelPinsUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: ChannelPinsUpdate")
			}
		}
	case EventChannelUpdate:
		for _, listener := range d.listeners[EventChannelUpdate] {
			if cb, ok := listener.(ChannelUpdateHandler); ok {
				cb(session, box.(*ChannelUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: ChannelUpdate")
			}
		}
	case EventGuildBanAdd:
		for _, listener := range d.listeners[EventGuildBanAdd] {
			if cb, ok := listener.(GuildBanAddHandler); ok {
				cb(session, box.(*GuildBanAdd))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildBanAdd")
			}
		}
	case EventGuildBanRemove:
		for _, listener := range d.listeners[EventGuildBanRemove] {
			if cb, ok := listener.(GuildBanRemoveHandler); ok {
				cb(session, box.(*GuildBanRemove))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildBanRemove")
			}
		}
	case EventGuildCreate:
		for _, listener := range d.listeners[EventGuildCreate] {
			if cb, ok := listener.(GuildCreateHandler); ok {
				cb(session, box.(*GuildCreate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildCreate")
			}
		}
	case EventGuildDelete:
		for _, listener := range d.listeners[EventGuildDelete] {
			if cb, ok := listener.(GuildDeleteHandler); ok {
				cb(session, box.(*GuildDelete))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildDelete")
			}
		}
	case EventGuildEmojisUpdate:
		for _, listener := range d.listeners[EventGuildEmojisUpdate] {
			if cb, ok := listener.(GuildEmojisUpdateHandler); ok {
				cb(session, box.(*GuildEmojisUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildEmojisUpdate")
			}
		}
	case EventGuildIntegrationsUpdate:
		for _, listener := range d.listeners[EventGuildIntegrationsUpdate] {
			if cb, ok := listener.(GuildIntegrationsUpdateHandler); ok {
				cb(session, box.(*GuildIntegrationsUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildIntegrationsUpdate")
			}
		}
	case EventGuildMemberAdd:
		for _, listener := range d.listeners[EventGuildMemberAdd] {
			if cb, ok := listener.(GuildMemberAddHandler); ok {
				cb(session, box.(*GuildMemberAdd))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildMemberAdd")
			}
		}
	case EventGuildMemberRemove:
		for _, listener := range d.listeners[EventGuildMemberRemove] {
			if cb, ok := listener.(GuildMemberRemoveHandler); ok {
				cb(session, box.(*GuildMemberRemove))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildMemberRemove")
			}
		}
	case EventGuildMemberUpdate:
		for _, listener := range d.listeners[EventGuildMemberUpdate] {
			if cb, ok := listener.(GuildMemberUpdateHandler); ok {
				cb(session, box.(*GuildMemberUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildMemberUpdate")
			}
		}
	case EventGuildMembersChunk:
		for _, listener := range d.listeners[EventGuildMembersChunk] {
			if cb, ok := listener.(GuildMembersChunkHandler); ok {
				cb(session, box.(*GuildMembersChunk))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildMembersChunk")
			}
		}
	case EventGuildRoleCreate:
		for _, listener := range d.listeners[EventGuildRoleCreate] {
			if cb, ok := listener.(GuildRoleCreateHandler); ok {
				cb(session, box.(*GuildRoleCreate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildRoleCreate")
			}
		}
	case EventGuildRoleDelete:
		for _, listener := range d.listeners[EventGuildRoleDelete] {
			if cb, ok := listener.(GuildRoleDeleteHandler); ok {
				cb(session, box.(*GuildRoleDelete))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildRoleDelete")
			}
		}
	case EventGuildRoleUpdate:
		for _, listener := range d.listeners[EventGuildRoleUpdate] {
			if cb, ok := listener.(GuildRoleUpdateHandler); ok {
				cb(session, box.(*GuildRoleUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildRoleUpdate")
			}
		}
	case EventGuildUpdate:
		for _, listener := range d.listeners[EventGuildUpdate] {
			if cb, ok := listener.(GuildUpdateHandler); ok {
				cb(session, box.(*GuildUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: GuildUpdate")
			}
		}
	case EventMessageCreate:
		for _, listener := range d.listeners[EventMessageCreate] {
			if cb, ok := listener.(MessageCreateHandler); ok {
				cb(session, box.(*MessageCreate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: MessageCreate")
			}
		}
	case EventMessageDelete:
		for _, listener := range d.listeners[EventMessageDelete] {
			if cb, ok := listener.(MessageDeleteHandler); ok {
				cb(session, box.(*MessageDelete))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: MessageDelete")
			}
		}
	case EventMessageDeleteBulk:
		for _, listener := range d.listeners[EventMessageDeleteBulk] {
			if cb, ok := listener.(MessageDeleteBulkHandler); ok {
				cb(session, box.(*MessageDeleteBulk))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: MessageDeleteBulk")
			}
		}
	case EventMessageReactionAdd:
		for _, listener := range d.listeners[EventMessageReactionAdd] {
			if cb, ok := listener.(MessageReactionAddHandler); ok {
				cb(session, box.(*MessageReactionAdd))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: MessageReactionAdd")
			}
		}
	case EventMessageReactionRemove:
		for _, listener := range d.listeners[EventMessageReactionRemove] {
			if cb, ok := listener.(MessageReactionRemoveHandler); ok {
				cb(session, box.(*MessageReactionRemove))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: MessageReactionRemove")
			}
		}
	case EventMessageReactionRemoveAll:
		for _, listener := range d.listeners[EventMessageReactionRemoveAll] {
			if cb, ok := listener.(MessageReactionRemoveAllHandler); ok {
				cb(session, box.(*MessageReactionRemoveAll))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: MessageReactionRemoveAll")
			}
		}
	case EventMessageUpdate:
		for _, listener := range d.listeners[EventMessageUpdate] {
			if cb, ok := listener.(MessageUpdateHandler); ok {
				cb(session, box.(*MessageUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: MessageUpdate")
			}
		}
	case EventPresenceUpdate:
		for _, listener := range d.listeners[EventPresenceUpdate] {
			if cb, ok := listener.(PresenceUpdateHandler); ok {
				cb(session, box.(*PresenceUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: PresenceUpdate")
			}
		}
	case EventPresencesReplace:
		for _, listener := range d.listeners[EventPresencesReplace] {
			if cb, ok := listener.(PresencesReplaceHandler); ok {
				cb(session, box.(*PresencesReplace))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: PresencesReplace")
			}
		}
	case EventReady:
		for _, listener := range d.listeners[EventReady] {
			if cb, ok := listener.(ReadyHandler); ok {
				cb(session, box.(*Ready))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: Ready")
			}
		}
	case EventResumed:
		for _, listener := range d.listeners[EventResumed] {
			if cb, ok := listener.(ResumedHandler); ok {
				cb(session, box.(*Resumed))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: Resumed")
			}
		}
	case EventTypingStart:
		for _, listener := range d.listeners[EventTypingStart] {
			if cb, ok := listener.(TypingStartHandler); ok {
				cb(session, box.(*TypingStart))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: TypingStart")
			}
		}
	case EventUserUpdate:
		for _, listener := range d.listeners[EventUserUpdate] {
			if cb, ok := listener.(UserUpdateHandler); ok {
				cb(session, box.(*UserUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: UserUpdate")
			}
		}
	case EventVoiceServerUpdate:
		for _, listener := range d.listeners[EventVoiceServerUpdate] {
			if cb, ok := listener.(VoiceServerUpdateHandler); ok {
				cb(session, box.(*VoiceServerUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: VoiceServerUpdate")
			}
		}
	case EventVoiceStateUpdate:
		for _, listener := range d.listeners[EventVoiceStateUpdate] {
			if cb, ok := listener.(VoiceStateUpdateHandler); ok {
				cb(session, box.(*VoiceStateUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: VoiceStateUpdate")
			}
		}
	case EventWebhooksUpdate:
		for _, listener := range d.listeners[EventWebhooksUpdate] {
			if cb, ok := listener.(WebhooksUpdateHandler); ok {
				cb(session, box.(*WebhooksUpdate))
			} else if cb, ok := listener.(SimpleHandler); ok {
				cb(session)
			} else if cb, ok := listener.(SimplestHandler); ok {
				cb()
			} else {
				fmt.Println("ERROR! Incorrect handler type given for event: WebhooksUpdate")
			}
		}
		//default:
		//	fmt.Printf("------\nTODO\nImplement handler for `%s`\n------\n\n", evtName)
	}
	d.listenersLock.RUnlock()

	// remove the run only once listeners
	d.listenersLock.Lock()
	defer d.listenersLock.Unlock()

	for _, index := range d.listenOnceOnly[evtName] {
		// https://github.com/golang/go/wiki/SliceTricks#delete-without-preserving-order
		d.listeners[evtName][index] = d.listeners[evtName][len(d.listeners[evtName])-1]
		d.listeners[evtName][len(d.listeners[evtName])-1] = nil
		d.listeners[evtName] = d.listeners[evtName][:len(d.listeners[evtName])-1]

		if len(d.listeners[evtName]) == 0 {
			// TODO: call removeEvent from socket pkg
		}
	}

	// remove the once only register
	if _, exists := d.listenOnceOnly[evtName]; exists {
		delete(d.listenOnceOnly, evtName)
	}
}

// ChannelCreate gives access to channelCreateChan for ChannelCreate events
func (d *Dispatch) ChannelCreate() <-chan *ChannelCreate {
	return d.channelCreateChan
}

// ChannelDelete gives access to channelDeleteChan for ChannelDelete events
func (d *Dispatch) ChannelDelete() <-chan *ChannelDelete {
	return d.channelDeleteChan
}

// ChannelPinsUpdate gives access to channelPinsUpdateChan for ChannelPinsUpdate events
func (d *Dispatch) ChannelPinsUpdate() <-chan *ChannelPinsUpdate {
	return d.channelPinsUpdateChan
}

// ChannelUpdate gives access to channelUpdateChan for ChannelUpdate events
func (d *Dispatch) ChannelUpdate() <-chan *ChannelUpdate {
	return d.channelUpdateChan
}

// GuildBanAdd gives access to guildBanAddChan for GuildBanAdd events
func (d *Dispatch) GuildBanAdd() <-chan *GuildBanAdd {
	return d.guildBanAddChan
}

// GuildBanRemove gives access to guildBanRemoveChan for GuildBanRemove events
func (d *Dispatch) GuildBanRemove() <-chan *GuildBanRemove {
	return d.guildBanRemoveChan
}

// GuildCreate gives access to guildCreateChan for GuildCreate events
func (d *Dispatch) GuildCreate() <-chan *GuildCreate {
	return d.guildCreateChan
}

// GuildDelete gives access to guildDeleteChan for GuildDelete events
func (d *Dispatch) GuildDelete() <-chan *GuildDelete {
	return d.guildDeleteChan
}

// GuildEmojisUpdate gives access to guildEmojisUpdateChan for GuildEmojisUpdate events
func (d *Dispatch) GuildEmojisUpdate() <-chan *GuildEmojisUpdate {
	return d.guildEmojisUpdateChan
}

// GuildIntegrationsUpdate gives access to guildIntegrationsUpdateChan for GuildIntegrationsUpdate events
func (d *Dispatch) GuildIntegrationsUpdate() <-chan *GuildIntegrationsUpdate {
	return d.guildIntegrationsUpdateChan
}

// GuildMemberAdd gives access to guildMemberAddChan for GuildMemberAdd events
func (d *Dispatch) GuildMemberAdd() <-chan *GuildMemberAdd {
	return d.guildMemberAddChan
}

// GuildMemberRemove gives access to guildMemberRemoveChan for GuildMemberRemove events
func (d *Dispatch) GuildMemberRemove() <-chan *GuildMemberRemove {
	return d.guildMemberRemoveChan
}

// GuildMemberUpdate gives access to guildMemberUpdateChan for GuildMemberUpdate events
func (d *Dispatch) GuildMemberUpdate() <-chan *GuildMemberUpdate {
	return d.guildMemberUpdateChan
}

// GuildMembersChunk gives access to guildMembersChunkChan for GuildMembersChunk events
func (d *Dispatch) GuildMembersChunk() <-chan *GuildMembersChunk {
	return d.guildMembersChunkChan
}

// GuildRoleCreate gives access to guildRoleCreateChan for GuildRoleCreate events
func (d *Dispatch) GuildRoleCreate() <-chan *GuildRoleCreate {
	return d.guildRoleCreateChan
}

// GuildRoleDelete gives access to guildRoleDeleteChan for GuildRoleDelete events
func (d *Dispatch) GuildRoleDelete() <-chan *GuildRoleDelete {
	return d.guildRoleDeleteChan
}

// GuildRoleUpdate gives access to guildRoleUpdateChan for GuildRoleUpdate events
func (d *Dispatch) GuildRoleUpdate() <-chan *GuildRoleUpdate {
	return d.guildRoleUpdateChan
}

// GuildUpdate gives access to guildUpdateChan for GuildUpdate events
func (d *Dispatch) GuildUpdate() <-chan *GuildUpdate {
	return d.guildUpdateChan
}

// MessageCreate gives access to messageCreateChan for MessageCreate events
func (d *Dispatch) MessageCreate() <-chan *MessageCreate {
	return d.messageCreateChan
}

// MessageDelete gives access to messageDeleteChan for MessageDelete events
func (d *Dispatch) MessageDelete() <-chan *MessageDelete {
	return d.messageDeleteChan
}

// MessageDeleteBulk gives access to messageDeleteBulkChan for MessageDeleteBulk events
func (d *Dispatch) MessageDeleteBulk() <-chan *MessageDeleteBulk {
	return d.messageDeleteBulkChan
}

// MessageReactionAdd gives access to messageReactionAddChan for MessageReactionAdd events
func (d *Dispatch) MessageReactionAdd() <-chan *MessageReactionAdd {
	return d.messageReactionAddChan
}

// MessageReactionRemove gives access to messageReactionRemoveChan for MessageReactionRemove events
func (d *Dispatch) MessageReactionRemove() <-chan *MessageReactionRemove {
	return d.messageReactionRemoveChan
}

// MessageReactionRemoveAll gives access to messageReactionRemoveAllChan for MessageReactionRemoveAll events
func (d *Dispatch) MessageReactionRemoveAll() <-chan *MessageReactionRemoveAll {
	return d.messageReactionRemoveAllChan
}

// MessageUpdate gives access to messageUpdateChan for MessageUpdate events
func (d *Dispatch) MessageUpdate() <-chan *MessageUpdate {
	return d.messageUpdateChan
}

// PresenceUpdate gives access to presenceUpdateChan for PresenceUpdate events
func (d *Dispatch) PresenceUpdate() <-chan *PresenceUpdate {
	return d.presenceUpdateChan
}

// PresencesReplace gives access to presencesReplaceChan for PresencesReplace events
func (d *Dispatch) PresencesReplace() <-chan *PresencesReplace {
	return d.presencesReplaceChan
}

// Ready gives access to readyChan for Ready events
func (d *Dispatch) Ready() <-chan *Ready {
	return d.readyChan
}

// Resumed gives access to resumedChan for Resumed events
func (d *Dispatch) Resumed() <-chan *Resumed {
	return d.resumedChan
}

// TypingStart gives access to typingStartChan for TypingStart events
func (d *Dispatch) TypingStart() <-chan *TypingStart {
	return d.typingStartChan
}

// UserUpdate gives access to userUpdateChan for UserUpdate events
func (d *Dispatch) UserUpdate() <-chan *UserUpdate {
	return d.userUpdateChan
}

// VoiceServerUpdate gives access to voiceServerUpdateChan for VoiceServerUpdate events
func (d *Dispatch) VoiceServerUpdate() <-chan *VoiceServerUpdate {
	return d.voiceServerUpdateChan
}

// VoiceStateUpdate gives access to voiceStateUpdateChan for VoiceStateUpdate events
func (d *Dispatch) VoiceStateUpdate() <-chan *VoiceStateUpdate {
	return d.voiceStateUpdateChan
}

// WebhooksUpdate gives access to webhooksUpdateChan for WebhooksUpdate events
func (d *Dispatch) WebhooksUpdate() <-chan *WebhooksUpdate {
	return d.webhooksUpdateChan
}
