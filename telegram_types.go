// Telegram bot API classes and enpoint
package tgbot

// Telegram Bot API                      Twitter   Home  FAQ  Apps  API  Protocol  Schema
// Telegram Bots Telegram Bot API  Telegram Bot API    The Bot API is an HTTP-based interface
// created for developers keen on building bots for Telegram. To learn how to create and set up
// a bot, please consult our Introduction to Bots and Bot FAQ .    This object represents an
// incoming update. At most one of the optional parameters can be present in any given update.
type Update struct {
	// The update's unique identifier. Update identifiers start from a certain positive number
	// and increase sequentially. This identifier becomes especially handy if you're using
	// webhooks, since it allows you to ignore repeated updates or to restore the correct update
	// sequence, should they get out of order. If there are no new updates for at least a week,
	// then identifier of the next update will be chosen randomly instead of sequentially.
	UpdateID int64 `json:"update_id,omitempty"`

	// Optional. New incoming message of any kind - text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`

	// Optional. New version of a message that is known to the bot and was edited. This update
	// may at times be triggered by changes to message fields that are either unavailable or not
	// actively used by your bot.
	EditedMessage *Message `json:"edited_message,omitempty"`

	// Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`

	// Optional. New version of a channel post that is known to the bot and was edited. This
	// update may at times be triggered by changes to message fields that are either unavailable
	// or not actively used by your bot.
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// Optional. A reaction to a message was changed by a user. The bot must be an administrator
	// in the chat and must explicitly specify "message_reaction" in the list of allowed_updates
	// to receive these updates. The update isn't received for reactions set by bots.
	MessageReaction *MessageReactionUpdated `json:"message_reaction,omitempty"`

	// Optional. Reactions to a message with anonymous reactions were changed. The bot must be an
	// administrator in the chat and must explicitly specify "message_reaction_count" in the list
	// of allowed_updates to receive these updates. The updates are grouped and can be sent with
	// delay up to a few minutes.
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`

	// Optional. New incoming inline query
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`

	// Optional. The result of an inline query that was chosen by a user and sent to their chat
	// partner. Please see our documentation on the feedback collecting for details on how to
	// enable these updates for your bot.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`

	// Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

	// Optional. New incoming shipping query. Only for invoices with flexible price
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

	// Optional. New incoming pre-checkout query. Contains full information about checkout
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	// Optional. New poll state. Bots receive only updates about manually stopped polls and
	// polls, which are sent by the bot
	Poll *Poll `json:"poll,omitempty"`

	// Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only
	// in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`

	// Optional. The bot's chat member status was updated in a chat. For private chats, this
	// update is received only when the bot is blocked or unblocked by the user.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`

	// Optional. A chat member's status was updated in a chat. The bot must be an administrator
	// in the chat and must explicitly specify "chat_member" in the list of allowed_updates to
	// receive these updates.
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`

	// Optional. A request to join the chat has been sent. The bot must have the can_invite_users
	// administrator right in the chat to receive these updates.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`

	// Optional. A chat boost was added or changed. The bot must be an administrator in the chat
	// to receive these updates.
	ChatBoost *ChatBoostUpdated `json:"chat_boost,omitempty"`

	// Optional. A boost was removed from a chat. The bot must be an administrator in the chat to
	// receive these updates.
	RemovedChatBoost *ChatBoostRemoved `json:"removed_chat_boost,omitempty"`
}

// Describes the current status of a webhook.
type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	URL string `json:"url,omitempty"`

	// True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate,omitempty"`

	// Number of updates awaiting delivery
	PendingUpdateCount int64 `json:"pending_update_count,omitempty"`

	// Optional. Currently used webhook IP address
	IPAddress string `json:"ip_address,omitempty"`

	// Optional. Unix time for the most recent error that happened when trying to deliver an
	// update via webhook
	LastErrorDate int64 `json:"last_error_date,omitempty"`

	// Optional. Error message in human-readable format for the most recent error that happened
	// when trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message,omitempty"`

	// Optional. Unix time of the most recent error that happened when trying to synchronize
	// available updates with Telegram datacenters
	LastSynchronizationErrorDate int64 `json:"last_synchronization_error_date,omitempty"`

	// Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for
	// update delivery
	MaxConnections int64 `json:"max_connections,omitempty"`

	// Optional. A list of update types the bot is subscribed to. Defaults to all update types
	// except chat_member
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// All types used in the Bot API responses are represented as JSON-objects.  It is safe to use
// 32-bit signed integers for storing all Integer fields unless otherwise noted.   Optional
// fields may be not returned when irrelevant.    This object represents a Telegram user or
// bot.
type User struct {
	// Unique identifier for this user or bot. This number may have more than 32 significant bits
	// and some programming languages may have difficulty/silent defects in interpreting it. But
	// it has at most 52 significant bits, so a 64-bit integer or double-precision float type are
	// safe for storing this identifier.
	ID int64 `json:"id,omitempty"`

	// True, if this user is a bot
	IsBot bool `json:"is_bot,omitempty"`

	// User's or bot's first name
	FirstName string `json:"first_name,omitempty"`

	// Optional. User's or bot's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. User's or bot's username
	Username string `json:"username,omitempty"`

	// Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code,omitempty"`

	// Optional. True, if this user is a Telegram Premium user
	IsPremium bool `json:"is_premium,omitempty"`

	// Optional. True, if this user added the bot to the attachment menu
	AddedToAttachmentMenu bool `json:"added_to_attachment_menu,omitempty"`

	// Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanJoinGroups bool `json:"can_join_groups,omitempty"`

	// Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`

	// Optional. True, if the bot supports inline queries. Returned only in getMe.
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}

// This object represents a chat.
type Chat struct {
	// Unique identifier for this chat. This number may have more than 32 significant bits and
	// some programming languages may have difficulty/silent defects in interpreting it. But it
	// has at most 52 significant bits, so a signed 64-bit integer or double-precision float type
	// are safe for storing this identifier.
	ID int64 `json:"id,omitempty"`

	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type,omitempty"`

	// Optional. Title, for supergroups, channels and group chats
	Title string `json:"title,omitempty"`

	// Optional. Username, for private chats, supergroups and channels if available
	Username string `json:"username,omitempty"`

	// Optional. First name of the other party in a private chat
	FirstName string `json:"first_name,omitempty"`

	// Optional. Last name of the other party in a private chat
	LastName string `json:"last_name,omitempty"`

	// Optional. True, if the supergroup chat is a forum (has topics enabled)
	IsForum bool `json:"is_forum,omitempty"`

	// Optional. Chat photo. Returned only in getChat.
	Photo *ChatPhoto `json:"photo,omitempty"`

	// Optional. If non-empty, the list of all active chat usernames; for private chats,
	// supergroups and channels. Returned only in getChat.
	ActiveUsernames []string `json:"active_usernames,omitempty"`

	// Optional. List of available reactions allowed in the chat. If omitted, then all emoji
	// reactions are allowed. Returned only in getChat.
	AvailableReactions []*ReactionType `json:"available_reactions,omitempty"`

	// Optional. Identifier of the accent color for the chat name and backgrounds of the chat
	// photo, reply header, and link preview. See accent colors for more details. Returned only
	// in getChat. Always returned in getChat.
	AccentColorID int64 `json:"accent_color_id,omitempty"`

	// Optional. Custom emoji identifier of emoji chosen by the chat for the reply header and
	// link preview background. Returned only in getChat.
	BackgroundCustomEmojiID string `json:"background_custom_emoji_id,omitempty"`

	// Optional. Identifier of the accent color for the chat's profile background. See profile
	// accent colors for more details. Returned only in getChat.
	ProfileAccentColorID int64 `json:"profile_accent_color_id,omitempty"`

	// Optional. Custom emoji identifier of the emoji chosen by the chat for its profile
	// background. Returned only in getChat.
	ProfileBackgroundCustomEmojiID string `json:"profile_background_custom_emoji_id,omitempty"`

	// Optional. Custom emoji identifier of the emoji status of the chat or the other party in a
	// private chat. Returned only in getChat.
	EmojiStatusCustomEmojiID string `json:"emoji_status_custom_emoji_id,omitempty"`

	// Optional. Expiration date of the emoji status of the chat or the other party in a private
	// chat, in Unix time, if any. Returned only in getChat.
	EmojiStatusExpirationDate int64 `json:"emoji_status_expiration_date,omitempty"`

	// Optional. Bio of the other party in a private chat. Returned only in getChat.
	Bio string `json:"bio,omitempty"`

	// Optional. True, if privacy settings of the other party in the private chat allows to use
	// tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
	HasPrivateForwards bool `json:"has_private_forwards,omitempty"`

	// Optional. True, if the privacy settings of the other party restrict sending voice and
	// video note messages in the private chat. Returned only in getChat.
	HasRestrictedVoiceAndVideoMessages bool `json:"has_restricted_voice_and_video_messages,omitempty"`

	// Optional. True, if users need to join the supergroup before they can send messages.
	// Returned only in getChat.
	JoinToSendMessages bool `json:"join_to_send_messages,omitempty"`

	// Optional. True, if all users directly joining the supergroup need to be approved by
	// supergroup administrators. Returned only in getChat.
	JoinByRequest bool `json:"join_by_request,omitempty"`

	// Optional. Description, for groups, supergroups and channel chats. Returned only in
	// getChat.
	Description string `json:"description,omitempty"`

	// Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in
	// getChat.
	InviteLink string `json:"invite_link,omitempty"`

	// Optional. The most recent pinned message (by sending date). Returned only in getChat.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Optional. Default chat member permissions, for groups and supergroups. Returned only in
	// getChat.
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// Optional. For supergroups, the minimum allowed delay between consecutive messages sent by
	// each unpriviledged user; in seconds. Returned only in getChat.
	SlowModeDelay int64 `json:"slow_mode_delay,omitempty"`

	// Optional. The time after which all messages sent to the chat will be automatically
	// deleted; in seconds. Returned only in getChat.
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time,omitempty"`

	// Optional. True, if aggressive anti-spam checks are enabled in the supergroup. The field is
	// only available to chat administrators. Returned only in getChat.
	HasAggressiveAntiSpamEnabled bool `json:"has_aggressive_anti_spam_enabled,omitempty"`

	// Optional. True, if non-administrators can only get the list of bots and administrators in
	// the chat. Returned only in getChat.
	HasHiddenMembers bool `json:"has_hidden_members,omitempty"`

	// Optional. True, if messages from the chat can't be forwarded to other chats. Returned only
	// in getChat.
	HasProtectedContent bool `json:"has_protected_content,omitempty"`

	// Optional. True, if new chat members will have access to old messages; available only to
	// chat administrators. Returned only in getChat.
	HasVisibleHistory bool `json:"has_visible_history,omitempty"`

	// Optional. For supergroups, name of group sticker set. Returned only in getChat.
	StickerSetName string `json:"sticker_set_name,omitempty"`

	// Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`

	// Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for
	// a channel and vice versa; for supergroups and channel chats. This identifier may be
	// greater than 32 bits and some programming languages may have difficulty/silent defects in
	// interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-
	// precision float type are safe for storing this identifier. Returned only in getChat.
	LinkedChatID int64 `json:"linked_chat_id,omitempty"`

	// Optional. For supergroups, the location to which the supergroup is connected. Returned
	// only in getChat.
	Location *ChatLocation `json:"location,omitempty"`
}

// This object represents a message.
type Message struct {
	// Unique message identifier inside this chat
	MessageID int64 `json:"message_id,omitempty"`

	// Optional. Unique identifier of a message thread to which the message belongs; for
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Optional. Sender of the message; empty for messages sent to channels. For backward
	// compatibility, the field contains a fake sender user in non-channel chats, if the message
	// was sent on behalf of a chat.
	From *User `json:"from,omitempty"`

	// Optional. Sender of the message, sent on behalf of a chat. For example, the channel itself
	// for channel posts, the supergroup itself for messages from anonymous group administrators,
	// the linked channel for messages automatically forwarded to the discussion group. For
	// backward compatibility, the field from contains a fake sender user in non-channel chats,
	// if the message was sent on behalf of a chat.
	SenderChat *Chat `json:"sender_chat,omitempty"`

	// Date the message was sent in Unix time. It is always a positive number, representing a
	// valid date.
	Date int64 `json:"date,omitempty"`

	// Chat the message belongs to
	Chat *Chat `json:"chat,omitempty"`

	// Optional. Information about the original message for forwarded messages
	ForwardOrigin *MessageOrigin `json:"forward_origin,omitempty"`

	// Optional. True, if the message is sent to a forum topic
	IsTopicMessage bool `json:"is_topic_message,omitempty"`

	// Optional. True, if the message is a channel post that was automatically forwarded to the
	// connected discussion group
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`

	// Optional. For replies in the same chat and message thread, the original message. Note that
	// the Message object in this field will not contain further reply_to_message fields even if
	// it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`

	// Optional. Information about the message that is being replied to, which may come from
	// another chat or forum topic
	ExternalReply *ExternalReplyInfo `json:"external_reply,omitempty"`

	// Optional. For replies that quote part of the original message, the quoted part of the
	// message
	Quote *TextQuote `json:"quote,omitempty"`

	// Optional. Bot through which the message was sent
	ViaBot *User `json:"via_bot,omitempty"`

	// Optional. Date the message was last edited in Unix time
	EditDate int64 `json:"edit_date,omitempty"`

	// Optional. True, if the message can't be forwarded
	HasProtectedContent bool `json:"has_protected_content,omitempty"`

	// Optional. The unique identifier of a media message group this message belongs to
	MediaGroupID string `json:"media_group_id,omitempty"`

	// Optional. Signature of the post author for messages in channels, or the custom title of an
	// anonymous group administrator
	AuthorSignature string `json:"author_signature,omitempty"`

	// Optional. For text messages, the actual UTF-8 text of the message
	Text string `json:"text,omitempty"`

	// Optional. For text messages, special entities like usernames, URLs, bot commands, etc.
	// that appear in the text
	Entities []*MessageEntity `json:"entities,omitempty"`

	// Optional. Options used for link preview generation for the message, if it is a text
	// message and link preview options were changed
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// Optional. Message is an animation, information about the animation. For backward
	// compatibility, when this field is set, the document field will also be set
	Animation *Animation `json:"animation,omitempty"`

	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`

	// Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`

	// Optional. Message is a photo, available sizes of the photo
	Photo []*PhotoSize `json:"photo,omitempty"`

	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`

	// Optional. Message is a forwarded story
	Story *Story `json:"story,omitempty"`

	// Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`

	// Optional. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`

	// Optional. Caption for the animation, audio, document, photo, video or voice
	Caption string `json:"caption,omitempty"`

	// Optional. For messages with a caption, special entities like usernames, URLs, bot
	// commands, etc. that appear in the caption
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`

	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`

	// Optional. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`

	// Optional. Message is a game, information about the game. More about games »
	Game *Game `json:"game,omitempty"`

	// Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`

	// Optional. Message is a venue, information about the venue. For backward compatibility,
	// when this field is set, the location field will also be set
	Venue *Venue `json:"venue,omitempty"`

	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`

	// Optional. New members that were added to the group or supergroup and information about
	// them (the bot itself may be one of these members)
	NewChatMembers []*User `json:"new_chat_members,omitempty"`

	// Optional. A member was removed from the group, information about them (this member may be
	// the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`

	// Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title,omitempty"`

	// Optional. A chat photo was change to this value
	NewChatPhoto []*PhotoSize `json:"new_chat_photo,omitempty"`

	// Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`

	// Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`

	// Optional. Service message: the supergroup has been created. This field can't be received
	// in a message coming through updates, because bot can't be a member of a supergroup when it
	// is created. It can only be found in reply_to_message if someone replies to a very first
	// message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`

	// Optional. Service message: the channel has been created. This field can't be received in a
	// message coming through updates, because bot can't be a member of a channel when it is
	// created. It can only be found in reply_to_message if someone replies to a very first
	// message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`

	// Optional. Service message: auto-delete timer settings changed in the chat
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`

	// Optional. The group has been migrated to a supergroup with the specified identifier. This
	// number may have more than 32 significant bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// Optional. The supergroup has been migrated from a group with the specified identifier.
	// This number may have more than 32 significant bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatID int64 `json:"migrate_from_chat_id,omitempty"`

	// Optional. Specified message was pinned. Note that the Message object in this field will
	// not contain further reply_to_message fields even if it itself is a reply.
	PinnedMessage *MaybeInaccessibleMessage `json:"pinned_message,omitempty"`

	// Optional. Message is an invoice for a payment, information about the invoice. More about
	// payments »
	Invoice *Invoice `json:"invoice,omitempty"`

	// Optional. Message is a service message about a successful payment, information about the
	// payment. More about payments »
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`

	// Optional. Service message: users were shared with the bot
	UsersShared *UsersShared `json:"users_shared,omitempty"`

	// Optional. Service message: a chat was shared with the bot
	ChatShared *ChatShared `json:"chat_shared,omitempty"`

	// Optional. The domain name of the website on which the user has logged in. More about
	// Telegram Login »
	ConnectedWebsite string `json:"connected_website,omitempty"`

	// Optional. Service message: the user allowed the bot to write messages after adding it to
	// the attachment or side menu, launching a Web App from a link, or accepting an explicit
	// request from a Web App sent by the method requestWriteAccess
	WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed,omitempty"`

	// Optional. Telegram Passport data
	PassportData *PassportData `json:"passport_data,omitempty"`

	// Optional. Service message. A user in the chat triggered another user's proximity alert
	// while sharing Live Location.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`

	// Optional. Service message: forum topic created
	ForumTopicCreated *ForumTopicCreated `json:"forum_topic_created,omitempty"`

	// Optional. Service message: forum topic edited
	ForumTopicEdited *ForumTopicEdited `json:"forum_topic_edited,omitempty"`

	// Optional. Service message: forum topic closed
	ForumTopicClosed *ForumTopicClosed `json:"forum_topic_closed,omitempty"`

	// Optional. Service message: forum topic reopened
	ForumTopicReopened *ForumTopicReopened `json:"forum_topic_reopened,omitempty"`

	// Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicHidden *GeneralForumTopicHidden `json:"general_forum_topic_hidden,omitempty"`

	// Optional. Service message: the 'General' forum topic unhidden
	GeneralForumTopicUnhidden *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden,omitempty"`

	// Optional. Service message: a scheduled giveaway was created
	GiveawayCreated *GiveawayCreated `json:"giveaway_created,omitempty"`

	// Optional. The message is a scheduled giveaway message
	Giveaway *Giveaway `json:"giveaway,omitempty"`

	// Optional. A giveaway with public winners was completed
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`

	// Optional. Service message: a giveaway without public winners was completed
	GiveawayCompleted *GiveawayCompleted `json:"giveaway_completed,omitempty"`

	// Optional. Service message: video chat scheduled
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`

	// Optional. Service message: video chat started
	VideoChatStarted *VideoChatStarted `json:"video_chat_started,omitempty"`

	// Optional. Service message: video chat ended
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`

	// Optional. Service message: new participants invited to a video chat
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`

	// Optional. Service message: data sent by a Web App
	WebAppData *WebAppData `json:"web_app_data,omitempty"`

	// Optional. Inline keyboard attached to the message. login_url buttons are represented as
	// ordinary url buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// This object represents a unique message identifier.
type MessageId struct {
	// Unique message identifier
	MessageID int64 `json:"message_id,omitempty"`
}

// This object describes a message that was deleted or is otherwise inaccessible to the bot.
type InaccessibleMessage struct {
	// Chat the message belonged to
	Chat *Chat `json:"chat,omitempty"`

	// Unique message identifier inside the chat
	MessageID int64 `json:"message_id,omitempty"`

	// Always 0. The field can be used to differentiate regular and inaccessible messages.
	Date int64 `json:"date,omitempty"`
}

// This object describes a message that can be inaccessible to the bot. It can be one of
// Message  InaccessibleMessage
type MaybeInaccessibleMessage struct {
}

// This object represents one special entity in a text message. For example, hashtags,
// usernames, URLs, etc.
type MessageEntity struct {
	// Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag),
	// “cashtag” ($USD), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email”
	// (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text),
	// “italic” (italic text), “underline” (underlined text), “strikethrough” (strikethrough
	// text), “spoiler” (spoiler message), “blockquote” (block quotation), “code” (monowidth
	// string), “pre” (monowidth block), “text_link” (for clickable text URLs), “text_mention”
	// (for users without usernames), “custom_emoji” (for inline custom emoji stickers)
	Type string `json:"type,omitempty"`

	// Offset in UTF-16 code units to the start of the entity
	Offset int64 `json:"offset,omitempty"`

	// Length of the entity in UTF-16 code units
	Length int64 `json:"length,omitempty"`

	// Optional. For “text_link” only, URL that will be opened after user taps on the text
	URL string `json:"url,omitempty"`

	// Optional. For “text_mention” only, the mentioned user
	User *User `json:"user,omitempty"`

	// Optional. For “pre” only, the programming language of the entity text
	Language string `json:"language,omitempty"`

	// Optional. For “custom_emoji” only, unique identifier of the custom emoji. Use
	// getCustomEmojiStickers to get full information about the sticker
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// This object contains information about the quoted part of a message that is replied to by
// the given message.
type TextQuote struct {
	// Text of the quoted part of a message that is replied to by the given message
	Text string `json:"text,omitempty"`

	// Optional. Special entities that appear in the quote. Currently, only bold, italic,
	// underline, strikethrough, spoiler, and custom_emoji entities are kept in quotes.
	Entities []*MessageEntity `json:"entities,omitempty"`

	// Approximate quote position in the original message in UTF-16 code units as specified by
	// the sender
	Position int64 `json:"position,omitempty"`

	// Optional. True, if the quote was chosen manually by the message sender. Otherwise, the
	// quote was added automatically by the server.
	IsManual bool `json:"is_manual,omitempty"`
}

// This object contains information about a message that is being replied to, which may come
// from another chat or forum topic.
type ExternalReplyInfo struct {
	// Origin of the message replied to by the given message
	Origin *MessageOrigin `json:"origin,omitempty"`

	// Optional. Chat the original message belongs to. Available only if the chat is a supergroup
	// or a channel.
	Chat *Chat `json:"chat,omitempty"`

	// Optional. Unique message identifier inside the original chat. Available only if the
	// original chat is a supergroup or a channel.
	MessageID int64 `json:"message_id,omitempty"`

	// Optional. Options used for link preview generation for the original message, if it is a
	// text message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// Optional. Message is an animation, information about the animation
	Animation *Animation `json:"animation,omitempty"`

	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`

	// Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`

	// Optional. Message is a photo, available sizes of the photo
	Photo []*PhotoSize `json:"photo,omitempty"`

	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`

	// Optional. Message is a forwarded story
	Story *Story `json:"story,omitempty"`

	// Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`

	// Optional. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`

	// Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`

	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`

	// Optional. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`

	// Optional. Message is a game, information about the game. More about games »
	Game *Game `json:"game,omitempty"`

	// Optional. Message is a scheduled giveaway, information about the giveaway
	Giveaway *Giveaway `json:"giveaway,omitempty"`

	// Optional. A giveaway with public winners was completed
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`

	// Optional. Message is an invoice for a payment, information about the invoice. More about
	// payments »
	Invoice *Invoice `json:"invoice,omitempty"`

	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`

	// Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`

	// Optional. Message is a venue, information about the venue
	Venue *Venue `json:"venue,omitempty"`
}

// Describes reply parameters for the message that is being sent.
type ReplyParameters struct {
	// Identifier of the message that will be replied to in the current chat, or in the chat
	// chat_id if it is specified
	MessageID int64 `json:"message_id,omitempty"`

	// Optional. If the message to be replied to is from a different chat, unique identifier for
	// the chat or username of the channel (in the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Optional. Pass True if the message should be sent even if the specified message to be
	// replied to is not found; can be used only for replies in the same chat and forum topic.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`

	// Optional. Quoted part of the message to be replied to; 0-1024 characters after entities
	// parsing. The quote must be an exact substring of the message to be replied to, including
	// bold, italic, underline, strikethrough, spoiler, and custom_emoji entities. The message
	// will fail to send if the quote isn't found in the original message.
	Quote string `json:"quote,omitempty"`

	// Optional. Mode for parsing entities in the quote. See formatting options for more details.
	QuoteParseMode string `json:"quote_parse_mode,omitempty"`

	// Optional. A JSON-serialized list of special entities that appear in the quote. It can be
	// specified instead of quote_parse_mode.
	QuoteEntities []*MessageEntity `json:"quote_entities,omitempty"`

	// Optional. Position of the quote in the original message in UTF-16 code units
	QuotePosition int64 `json:"quote_position,omitempty"`
}

// The message was originally sent by a known user.
type MessageOriginUser struct {
	// Type of the message origin, always “user”
	Type string `json:"type,omitempty"`

	// Date the message was sent originally in Unix time
	Date int64 `json:"date,omitempty"`

	// User that sent the message originally
	SenderUser *User `json:"sender_user,omitempty"`
}

// The message was originally sent by an unknown user.
type MessageOriginHiddenUser struct {
	// Type of the message origin, always “hidden_user”
	Type string `json:"type,omitempty"`

	// Date the message was sent originally in Unix time
	Date int64 `json:"date,omitempty"`

	// Name of the user that sent the message originally
	SenderUserName string `json:"sender_user_name,omitempty"`
}

// The message was originally sent on behalf of a chat to a group chat.
type MessageOriginChat struct {
	// Type of the message origin, always “chat”
	Type string `json:"type,omitempty"`

	// Date the message was sent originally in Unix time
	Date int64 `json:"date,omitempty"`

	// Chat that sent the message originally
	SenderChat *Chat `json:"sender_chat,omitempty"`

	// Optional. For messages originally sent by an anonymous chat administrator, original
	// message author signature
	AuthorSignature string `json:"author_signature,omitempty"`
}

// The message was originally sent to a channel chat.
type MessageOriginChannel struct {
	// Type of the message origin, always “channel”
	Type string `json:"type,omitempty"`

	// Date the message was sent originally in Unix time
	Date int64 `json:"date,omitempty"`

	// Channel chat to which the message was originally sent
	Chat *Chat `json:"chat,omitempty"`

	// Unique message identifier inside the chat
	MessageID int64 `json:"message_id,omitempty"`

	// Optional. Signature of the original post author
	AuthorSignature string `json:"author_signature,omitempty"`
}

// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id,omitempty"`

	// Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"`

	// Photo width
	Width int64 `json:"width,omitempty"`

	// Photo height
	Height int64 `json:"height,omitempty"`

	// Optional. File size in bytes
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id,omitempty"`

	// Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"`

	// Video width as defined by sender
	Width int64 `json:"width,omitempty"`

	// Video height as defined by sender
	Height int64 `json:"height,omitempty"`

	// Duration of the video in seconds as defined by sender
	Duration int64 `json:"duration,omitempty"`

	// Optional. Animation thumbnail as defined by sender
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// Optional. Original animation filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages
	// may have difficulty/silent defects in interpreting it. But it has at most 52 significant
	// bits, so a signed 64-bit integer or double-precision float type are safe for storing this
	// value.
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id,omitempty"`

	// Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"`

	// Duration of the audio in seconds as defined by sender
	Duration int64 `json:"duration,omitempty"`

	// Optional. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer,omitempty"`

	// Optional. Title of the audio as defined by sender or by audio tags
	Title string `json:"title,omitempty"`

	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages
	// may have difficulty/silent defects in interpreting it. But it has at most 52 significant
	// bits, so a signed 64-bit integer or double-precision float type are safe for storing this
	// value.
	FileSize int64 `json:"file_size,omitempty"`

	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

// This object represents a general file (as opposed to photos , voice messages and audio files
// ).
type Document struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id,omitempty"`

	// Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"`

	// Optional. Document thumbnail as defined by sender
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages
	// may have difficulty/silent defects in interpreting it. But it has at most 52 significant
	// bits, so a signed 64-bit integer or double-precision float type are safe for storing this
	// value.
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents a message about a forwarded story in the chat. Currently holds no
// information.
type Story struct {
}

// This object represents a video file.
type Video struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id,omitempty"`

	// Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"`

	// Video width as defined by sender
	Width int64 `json:"width,omitempty"`

	// Video height as defined by sender
	Height int64 `json:"height,omitempty"`

	// Duration of the video in seconds as defined by sender
	Duration int64 `json:"duration,omitempty"`

	// Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages
	// may have difficulty/silent defects in interpreting it. But it has at most 52 significant
	// bits, so a signed 64-bit integer or double-precision float type are safe for storing this
	// value.
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents a video message (available in Telegram apps as of v.4.0 ).
type VideoNote struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id,omitempty"`

	// Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"`

	// Video width and height (diameter of the video message) as defined by sender
	Length int64 `json:"length,omitempty"`

	// Duration of the video in seconds as defined by sender
	Duration int64 `json:"duration,omitempty"`

	// Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// Optional. File size in bytes
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents a voice note.
type Voice struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id,omitempty"`

	// Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"`

	// Duration of the audio in seconds as defined by sender
	Duration int64 `json:"duration,omitempty"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages
	// may have difficulty/silent defects in interpreting it. But it has at most 52 significant
	// bits, so a signed 64-bit integer or double-precision float type are safe for storing this
	// value.
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents a phone contact.
type Contact struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// Contact's first name
	FirstName string `json:"first_name,omitempty"`

	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. Contact's user identifier in Telegram. This number may have more than 32
	// significant bits and some programming languages may have difficulty/silent defects in
	// interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-
	// precision float type are safe for storing this identifier.
	UserID int64 `json:"user_id,omitempty"`

	// Optional. Additional data about the contact in the form of a vCard
	Vcard string `json:"vcard,omitempty"`
}

// This object represents an animated emoji that displays a random value.
type Dice struct {
	// Emoji on which the dice throw animation is based
	Emoji string `json:"emoji,omitempty"`

	// Value of the dice, 1-6 for “”, “” and “” base emoji, 1-5 for “” and “” base emoji, 1-64
	// for “” base emoji
	Value int64 `json:"value,omitempty"`
}

// This object contains information about one answer option in a poll.
type PollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text,omitempty"`

	// Number of users that voted for this option
	VoterCount int64 `json:"voter_count,omitempty"`
}

// This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	// Unique poll identifier
	PollID string `json:"poll_id,omitempty"`

	// Optional. The chat that changed the answer to the poll, if the voter is anonymous
	VoterChat *Chat `json:"voter_chat,omitempty"`

	// Optional. The user that changed the answer to the poll, if the voter isn't anonymous
	User *User `json:"user,omitempty"`

	// 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
	OptionIds []int64 `json:"option_ids,omitempty"`
}

// This object contains information about a poll.
type Poll struct {
	// Unique poll identifier
	ID string `json:"id,omitempty"`

	// Poll question, 1-300 characters
	Question string `json:"question,omitempty"`

	// List of poll options
	Options []*PollOption `json:"options,omitempty"`

	// Total number of users that voted in the poll
	TotalVoterCount int64 `json:"total_voter_count,omitempty"`

	// True, if the poll is closed
	IsClosed bool `json:"is_closed,omitempty"`

	// True, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// Poll type, currently can be “regular” or “quiz”
	Type string `json:"type,omitempty"`

	// True, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`

	// Optional. 0-based identifier of the correct answer option. Available only for polls in the
	// quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat
	// with the bot.
	CorrectOptionID int64 `json:"correct_option_id,omitempty"`

	// Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp
	// icon in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation,omitempty"`

	// Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the
	// explanation
	ExplanationEntities []*MessageEntity `json:"explanation_entities,omitempty"`

	// Optional. Amount of time in seconds the poll will be active after creation
	OpenPeriod int64 `json:"open_period,omitempty"`

	// Optional. Point in time (Unix timestamp) when the poll will be automatically closed
	CloseDate int64 `json:"close_date,omitempty"`
}

// This object represents a point on the map.
type Location struct {
	// Longitude as defined by sender
	Longitude float32 `json:"longitude,omitempty"`

	// Latitude as defined by sender
	Latitude float32 `json:"latitude,omitempty"`

	// Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`

	// Optional. Time relative to the message sending date, during which the location can be
	// updated; in seconds. For active live locations only.
	LivePeriod int64 `json:"live_period,omitempty"`

	// Optional. The direction in which user is moving, in degrees; 1-360. For active live
	// locations only.
	Heading int64 `json:"heading,omitempty"`

	// Optional. The maximum distance for proximity alerts about approaching another chat member,
	// in meters. For sent live locations only.
	ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
}

// This object represents a venue.
type Venue struct {
	// Venue location. Can't be a live location
	Location *Location `json:"location,omitempty"`

	// Name of the venue
	Title string `json:"title,omitempty"`

	// Address of the venue
	Address string `json:"address,omitempty"`

	// Optional. Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id,omitempty"`

	// Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// Optional. Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// Describes data sent from a Web App to the bot.
type WebAppData struct {
	// The data. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data,omitempty"`

	// Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad
	// client can send arbitrary data in this field.
	ButtonText string `json:"button_text,omitempty"`
}

// This object represents the content of a service message, sent whenever a user in the chat
// triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	// User that triggered the alert
	Traveler *User `json:"traveler,omitempty"`

	// User that set the alert
	Watcher *User `json:"watcher,omitempty"`

	// The distance between the users
	Distance int64 `json:"distance,omitempty"`
}

// This object represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	// New auto-delete time for messages in the chat; in seconds
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time,omitempty"`
}

// This object represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	// Name of the topic
	Name string `json:"name,omitempty"`

	// Color of the topic icon in RGB format
	IconColor int64 `json:"icon_color,omitempty"`

	// Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// This object represents a service message about a forum topic closed in the chat. Currently
// holds no information.
type ForumTopicClosed struct {
}

// This object represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	// Optional. New name of the topic, if it was edited
	Name string `json:"name,omitempty"`

	// Optional. New identifier of the custom emoji shown as the topic icon, if it was edited; an
	// empty string if the icon was removed
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// This object represents a service message about a forum topic reopened in the chat. Currently
// holds no information.
type ForumTopicReopened struct {
}

// This object represents a service message about General forum topic hidden in the chat.
// Currently holds no information.
type GeneralForumTopicHidden struct {
}

// This object represents a service message about General forum topic unhidden in the chat.
// Currently holds no information.
type GeneralForumTopicUnhidden struct {
}

// This object contains information about the users whose identifiers were shared with the bot
// using a KeyboardButtonRequestUsers button.
type UsersShared struct {
	// Identifier of the request
	RequestID int64 `json:"request_id,omitempty"`

	// Identifiers of the shared users. These numbers may have more than 32 significant bits and
	// some programming languages may have difficulty/silent defects in interpreting them. But
	// they have at most 52 significant bits, so 64-bit integers or double-precision float types
	// are safe for storing these identifiers. The bot may not have access to the users and could
	// be unable to use these identifiers, unless the users are already known to the bot by some
	// other means.
	UserIds []int64 `json:"user_ids,omitempty"`
}

// This object contains information about the chat whose identifier was shared with the bot
// using a KeyboardButtonRequestChat button.
type ChatShared struct {
	// Identifier of the request
	RequestID int64 `json:"request_id,omitempty"`

	// Identifier of the shared chat. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at
	// most 52 significant bits, so a 64-bit integer or double-precision float type are safe for
	// storing this identifier. The bot may not have access to the chat and could be unable to
	// use this identifier, unless the chat is already known to the bot by some other means.
	ChatID int64 `json:"chat_id,omitempty"`
}

// This object represents a service message about a user allowing a bot to write messages after
// adding it to the attachment menu, launching a Web App from a link, or accepting an explicit
// request from a Web App sent by the method requestWriteAccess .
type WriteAccessAllowed struct {
	// Optional. True, if the access was granted after the user accepted an explicit request from
	// a Web App sent by the method requestWriteAccess
	FromRequest bool `json:"from_request,omitempty"`

	// Optional. Name of the Web App, if the access was granted when the Web App was launched
	// from a link
	WebAppName string `json:"web_app_name,omitempty"`

	// Optional. True, if the access was granted when the bot was added to the attachment or side
	// menu
	FromAttachmentMenu bool `json:"from_attachment_menu,omitempty"`
}

// This object represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	// Point in time (Unix timestamp) when the video chat is supposed to be started by a chat
	// administrator
	StartDate int64 `json:"start_date,omitempty"`
}

// This object represents a service message about a video chat started in the chat. Currently
// holds no information.
type VideoChatStarted struct {
}

// This object represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	// Video chat duration in seconds
	Duration int64 `json:"duration,omitempty"`
}

// This object represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	// New members that were invited to the video chat
	Users []*User `json:"users,omitempty"`
}

// This object represents a service message about the creation of a scheduled giveaway.
// Currently holds no information.
type GiveawayCreated struct {
}

// This object represents a message about a scheduled giveaway.
type Giveaway struct {
	// The list of chats which the user must join to participate in the giveaway
	Chats []*Chat `json:"chats,omitempty"`

	// Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnersSelectionDate int64 `json:"winners_selection_date,omitempty"`

	// The number of users which are supposed to be selected as winners of the giveaway
	WinnerCount int64 `json:"winner_count,omitempty"`

	// Optional. True, if only users who join the chats after the giveaway started should be
	// eligible to win
	OnlyNewMembers bool `json:"only_new_members,omitempty"`

	// Optional. True, if the list of giveaway winners will be visible to everyone
	HasPublicWinners bool `json:"has_public_winners,omitempty"`

	// Optional. Description of additional giveaway prize
	PrizeDescription string `json:"prize_description,omitempty"`

	// Optional. A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries
	// from which eligible users for the giveaway must come. If empty, then all users can
	// participate in the giveaway. Users with a phone number that was bought on Fragment can
	// always participate in giveaways.
	CountryCodes []string `json:"country_codes,omitempty"`

	// Optional. The number of months the Telegram Premium subscription won from the giveaway
	// will be active for
	PremiumSubscriptionMonthCount int64 `json:"premium_subscription_month_count,omitempty"`
}

// This object represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	// The chat that created the giveaway
	Chat *Chat `json:"chat,omitempty"`

	// Identifier of the messsage with the giveaway in the chat
	GiveawayMessageID int64 `json:"giveaway_message_id,omitempty"`

	// Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnersSelectionDate int64 `json:"winners_selection_date,omitempty"`

	// Total number of winners in the giveaway
	WinnerCount int64 `json:"winner_count,omitempty"`

	// List of up to 100 winners of the giveaway
	Winners []*User `json:"winners,omitempty"`

	// Optional. The number of other chats the user had to join in order to be eligible for the
	// giveaway
	AdditionalChatCount int64 `json:"additional_chat_count,omitempty"`

	// Optional. The number of months the Telegram Premium subscription won from the giveaway
	// will be active for
	PremiumSubscriptionMonthCount int64 `json:"premium_subscription_month_count,omitempty"`

	// Optional. Number of undistributed prizes
	UnclaimedPrizeCount int64 `json:"unclaimed_prize_count,omitempty"`

	// Optional. True, if only users who had joined the chats after the giveaway started were
	// eligible to win
	OnlyNewMembers bool `json:"only_new_members,omitempty"`

	// Optional. True, if the giveaway was canceled because the payment for it was refunded
	WasRefunded bool `json:"was_refunded,omitempty"`

	// Optional. Description of additional giveaway prize
	PrizeDescription string `json:"prize_description,omitempty"`
}

// This object represents a service message about the completion of a giveaway without public
// winners.
type GiveawayCompleted struct {
	// Number of winners in the giveaway
	WinnerCount int64 `json:"winner_count,omitempty"`

	// Optional. Number of undistributed prizes
	UnclaimedPrizeCount int64 `json:"unclaimed_prize_count,omitempty"`

	// Optional. Message with the giveaway that was completed, if it wasn't deleted
	GiveawayMessage *Message `json:"giveaway_message,omitempty"`
}

// Describes the options used for link preview generation.
type LinkPreviewOptions struct {
	// Optional. True, if the link preview is disabled
	IsDisabled bool `json:"is_disabled,omitempty"`

	// Optional. URL to use for the link preview. If empty, then the first URL found in the
	// message text will be used
	URL string `json:"url,omitempty"`

	// Optional. True, if the media in the link preview is suppposed to be shrunk; ignored if the
	// URL isn't explicitly specified or media size change isn't supported for the preview
	PreferSmallMedia bool `json:"prefer_small_media,omitempty"`

	// Optional. True, if the media in the link preview is suppposed to be enlarged; ignored if
	// the URL isn't explicitly specified or media size change isn't supported for the preview
	PreferLargeMedia bool `json:"prefer_large_media,omitempty"`

	// Optional. True, if the link preview must be shown above the message text; otherwise, the
	// link preview will be shown below the message text
	ShowAboveText bool `json:"show_above_text,omitempty"`
}

// This object represent a user's profile pictures.
type UserProfilePhotos struct {
	// Total number of profile pictures the target user has
	TotalCount int64 `json:"total_count,omitempty"`

	// Requested profile pictures (in up to 4 sizes each)
	Photos [][]*PhotoSize `json:"photos,omitempty"`
}

// This object represents a file ready to be downloaded. The file can be downloaded via the
// link https://api.telegram.org/file/bot<token>/<file_path> . It is guaranteed that the link
// will be valid for at least 1 hour. When the link expires, a new one can be requested by
// calling getFile .   The maximum file size to download is 20 MB
type File struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id,omitempty"`

	// Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages
	// may have difficulty/silent defects in interpreting it. But it has at most 52 significant
	// bits, so a signed 64-bit integer or double-precision float type are safe for storing this
	// value.
	FileSize int64 `json:"file_size,omitempty"`

	// Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the
	// file.
	FilePath string `json:"file_path,omitempty"`
}

// Describes a Web App .
type WebAppInfo struct {
	// An HTTPS URL of a Web App to be opened with additional data as specified in Initializing
	// Web Apps
	URL string `json:"url,omitempty"`
}

// This object represents a custom keyboard with reply options (see Introduction to bots for
// details and examples).
type ReplyKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard [][]*KeyboardButton `json:"keyboard,omitempty"`

	// Optional. Requests clients to always show the keyboard when the regular keyboard is
	// hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with
	// a keyboard icon.
	IsPersistent bool `json:"is_persistent,omitempty"`

	// Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make
	// the keyboard smaller if there are just two rows of buttons). Defaults to false, in which
	// case the custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

	// Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard
	// will still be available, but clients will automatically display the usual letter-keyboard
	// in the chat - the user can press a special button in the input field to see the custom
	// keyboard again. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	// Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64
	// characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Optional. Use this parameter if you want to show the keyboard to specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's
	// message is a reply to a message in the same chat and forum topic, sender of the original
	// message.Example: A user requests to change the bot's language, bot replies to the request
	// with a keyboard to select the new language. Other users in the group don't see the
	// keyboard.
	Selective bool `json:"selective,omitempty"`
}

// This object represents one button of the reply keyboard. For simple text buttons, String can
// be used instead of this object to specify the button text. The optional fields web_app ,
// request_users , request_chat , request_contact , request_location , and request_poll are
// mutually exclusive.
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used, it will be sent as a message
	// when the button is pressed
	Text string `json:"text,omitempty"`

	// Optional. If specified, pressing the button will open a list of suitable users.
	// Identifiers of selected users will be sent to the bot in a “users_shared” service message.
	// Available in private chats only.
	RequestUsers *KeyboardButtonRequestUsers `json:"request_users,omitempty"`

	// Optional. If specified, pressing the button will open a list of suitable chats. Tapping on
	// a chat will send its identifier to the bot in a “chat_shared” service message. Available
	// in private chats only.
	RequestChat *KeyboardButtonRequestChat `json:"request_chat,omitempty"`

	// Optional. If True, the user's phone number will be sent as a contact when the button is
	// pressed. Available in private chats only.
	RequestContact bool `json:"request_contact,omitempty"`

	// Optional. If True, the user's current location will be sent when the button is pressed.
	// Available in private chats only.
	RequestLocation bool `json:"request_location,omitempty"`

	// Optional. If specified, the user will be asked to create a poll and send it to the bot
	// when the button is pressed. Available in private chats only.
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`

	// Optional. If specified, the described Web App will be launched when the button is pressed.
	// The Web App will be able to send a “web_app_data” service message. Available in private
	// chats only.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

// Note:  request_users and request_chat options will only work in Telegram versions released
// after 3 February, 2023. Older clients will display unsupported message .   This object
// defines the criteria used to request suitable users. The identifiers of the selected users
// will be shared with the bot when the corresponding button is pressed. More about requesting
// users »
type KeyboardButtonRequestUsers struct {
	// Signed 32-bit identifier of the request that will be received back in the UsersShared
	// object. Must be unique within the message
	RequestID int64 `json:"request_id,omitempty"`

	// Optional. Pass True to request bots, pass False to request regular users. If not
	// specified, no additional restrictions are applied.
	UserIsBot bool `json:"user_is_bot,omitempty"`

	// Optional. Pass True to request premium users, pass False to request non-premium users. If
	// not specified, no additional restrictions are applied.
	UserIsPremium bool `json:"user_is_premium,omitempty"`

	// Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
	MaxQuantity int64 `json:"max_quantity,omitempty"`
}

// This object defines the criteria used to request a suitable chat. The identifier of the
// selected chat will be shared with the bot when the corresponding button is pressed. More
// about requesting chats »
type KeyboardButtonRequestChat struct {
	// Signed 32-bit identifier of the request, which will be received back in the ChatShared
	// object. Must be unique within the message
	RequestID int64 `json:"request_id,omitempty"`

	// Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsChannel bool `json:"chat_is_channel,omitempty"`

	// Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat.
	// If not specified, no additional restrictions are applied.
	ChatIsForum bool `json:"chat_is_forum,omitempty"`

	// Optional. Pass True to request a supergroup or a channel with a username, pass False to
	// request a chat without a username. If not specified, no additional restrictions are
	// applied.
	ChatHasUsername bool `json:"chat_has_username,omitempty"`

	// Optional. Pass True to request a chat owned by the user. Otherwise, no additional
	// restrictions are applied.
	ChatIsCreated bool `json:"chat_is_created,omitempty"`

	// Optional. A JSON-serialized object listing the required administrator rights of the user
	// in the chat. The rights must be a superset of bot_administrator_rights. If not specified,
	// no additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`

	// Optional. A JSON-serialized object listing the required administrator rights of the bot in
	// the chat. The rights must be a subset of user_administrator_rights. If not specified, no
	// additional restrictions are applied.
	BotAdministratorRights *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`

	// Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional
	// restrictions are applied.
	BotIsMember bool `json:"bot_is_member,omitempty"`
}

// This object represents type of a poll, which is allowed to be created and sent when the
// corresponding button is pressed.
type KeyboardButtonPollType struct {
	// Optional. If quiz is passed, the user will be allowed to create only polls in the quiz
	// mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will
	// be allowed to create a poll of any type.
	Type string `json:"type,omitempty"`
}

// Upon receiving a message with this object, Telegram clients will remove the current custom
// keyboard and display the default letter-keyboard. By default, custom keyboards are displayed
// until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are
// hidden immediately after the user presses a button (see ReplyKeyboardMarkup ).
type ReplyKeyboardRemove struct {
	// Requests clients to remove the custom keyboard (user will not be able to summon this
	// keyboard; if you want to hide the keyboard from sight but keep it accessible, use
	// one_time_keyboard in ReplyKeyboardMarkup)
	RemoveKeyboard bool `json:"remove_keyboard,omitempty"`

	// Optional. Use this parameter if you want to remove the keyboard for specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's
	// message is a reply to a message in the same chat and forum topic, sender of the original
	// message.Example: A user votes in a poll, bot returns confirmation message in reply to the
	// vote and removes the keyboard for that user, while still showing the keyboard with poll
	// options to users who haven't voted yet.
	Selective bool `json:"selective,omitempty"`
}

// This object represents an inline keyboard that appears right next to the message it belongs
// to.
type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of InlineKeyboardButton objects
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

// This object represents one button of an inline keyboard. You must use exactly one of the
// optional fields.
type InlineKeyboardButton struct {
	// Label text on the button
	Text string `json:"text,omitempty"`

	// Optional. HTTP or tg:// URL to be opened when the button is pressed. Links
	// tg://user?id=<user_id> can be used to mention a user by their identifier without using a
	// username, if this is allowed by their privacy settings.
	URL string `json:"url,omitempty"`

	// Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64
	// bytes
	CallbackData string `json:"callback_data,omitempty"`

	// Optional. Description of the Web App that will be launched when the user presses the
	// button. The Web App will be able to send an arbitrary message on behalf of the user using
	// the method answerWebAppQuery. Available only in private chats between a user and the bot.
	WebApp *WebAppInfo `json:"web_app,omitempty"`

	// Optional. An HTTPS URL used to automatically authorize the user. Can be used as a
	// replacement for the Telegram Login Widget.
	LoginURL *LoginUrl `json:"login_url,omitempty"`

	// Optional. If set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline query in the input
	// field. May be empty, in which case just the bot's username will be inserted.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

	// Optional. If set, pressing the button will insert the bot's username and the specified
	// inline query in the current chat's input field. May be empty, in which case only the bot's
	// username will be inserted.This offers a quick way for the user to open your bot in inline
	// mode in the same chat - good for selecting something from multiple options.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

	// Optional. If set, pressing the button will prompt the user to select one of their chats of
	// the specified type, open that chat and insert the bot's username and the specified inline
	// query in the input field
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`

	// Optional. Description of the game that will be launched when the user presses the
	// button.NOTE: This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// Optional. Specify True, to send a Pay button.NOTE: This type of button must always be the
	// first button in the first row and can only be used in invoice messages.
	Pay bool `json:"pay,omitempty"`
}

// This object represents a parameter of the inline keyboard button used to automatically
// authorize a user. Serves as a great replacement for the Telegram Login Widget when the user
// is coming from Telegram. All the user needs to do is tap/click a button and confirm that
// they want to log in:     Telegram apps support these buttons as of version 5.7 .   Sample
// bot: @discussbot
type LoginUrl struct {
	// An HTTPS URL to be opened with user authorization data added to the query string when the
	// button is pressed. If the user refuses to provide authorization data, the original URL
	// without information about the user will be opened. The data added is the same as described
	// in Receiving authorization data.NOTE: You must always check the hash of the received data
	// to verify the authentication and the integrity of the data as described in Checking
	// authorization.
	URL string `json:"url,omitempty"`

	// Optional. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`

	// Optional. Username of a bot, which will be used for user authorization. See Setting up a
	// bot for more details. If not specified, the current bot's username will be assumed. The
	// url's domain must be the same as the domain linked with the bot. See Linking your domain
	// to the bot for more details.
	BotUsername string `json:"bot_username,omitempty"`

	// Optional. Pass True to request the permission for your bot to send messages to the user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// This object represents an inline button that switches the current user to inline mode in a
// chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {
	// Optional. The default inline query to be inserted in the input field. If left empty, only
	// the bot's username will be inserted
	Query string `json:"query,omitempty"`

	// Optional. True, if private chats with users can be chosen
	AllowUserChats bool `json:"allow_user_chats,omitempty"`

	// Optional. True, if private chats with bots can be chosen
	AllowBotChats bool `json:"allow_bot_chats,omitempty"`

	// Optional. True, if group and supergroup chats can be chosen
	AllowGroupChats bool `json:"allow_group_chats,omitempty"`

	// Optional. True, if channel chats can be chosen
	AllowChannelChats bool `json:"allow_channel_chats,omitempty"`
}

// This object represents an incoming callback query from a callback button in an inline
// keyboard . If the button that originated the query was attached to a message sent by the
// bot, the field message will be present. If the button was attached to a message sent via the
// bot (in inline mode ), the field inline_message_id will be present. Exactly one of the
// fields data or game_short_name will be present.
type CallbackQuery struct {
	// Unique identifier for this query
	ID string `json:"id,omitempty"`

	// Sender
	From *User `json:"from,omitempty"`

	// Optional. Message sent by the bot with the callback button that originated the query
	Message *MaybeInaccessibleMessage `json:"message,omitempty"`

	// Optional. Identifier of the message sent via the bot in inline mode, that originated the
	// query.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Global identifier, uniquely corresponding to the chat to which the message with the
	// callback button was sent. Useful for high scores in games.
	ChatInstance string `json:"chat_instance,omitempty"`

	// Optional. Data associated with the callback button. Be aware that the message originated
	// the query can contain no callback buttons with this data.
	Data string `json:"data,omitempty"`

	// Optional. Short name of a Game to be returned, serves as the unique identifier for the
	// game
	GameShortName string `json:"game_short_name,omitempty"`
}

// NOTE: After the user presses a callback button, Telegram clients will display a progress bar
// until you call answerCallbackQuery . It is, therefore, necessary to react by calling
// answerCallbackQuery even if no notification to the user is needed (e.g., without specifying
// any of the optional parameters).    Upon receiving a message with this object, Telegram
// clients will display a reply interface to the user (act as if the user has selected the
// bot's message and tapped 'Reply'). This can be extremely useful if you want to create user-
// friendly step-by-step interfaces without having to sacrifice privacy mode .
type ForceReply struct {
	// Shows reply interface to the user, as if they manually selected the bot's message and
	// tapped 'Reply'
	ForceReply bool `json:"force_reply,omitempty"`

	// Optional. The placeholder to be shown in the input field when the reply is active; 1-64
	// characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Optional. Use this parameter if you want to force reply from specific users only. Targets:
	// 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is
	// a reply to a message in the same chat and forum topic, sender of the original message.
	Selective bool `json:"selective,omitempty"`
}

// Example: A poll bot for groups runs in privacy mode (only receives commands, replies to its
// messages and mentions). There could be two ways to create a new poll:   Explain the user how
// to send a command with parameters (e.g. /newpoll question answer1 answer2). May be appealing
// for hardcore users but lacks modern day polish.  Guide the user through a step-by-step
// process. 'Please send me your question', 'Cool, now let's add the first answer option',
// 'Great. Keep adding answer options, then send /done when you're ready'.   The last option is
// definitely more attractive. And if you use ForceReply in your bot's questions, it will
// receive the user's answers even if it only receives replies, commands and mentions - without
// any extra work for the user.    This object represents a chat photo.
type ChatPhoto struct {
	// File identifier of small (160x160) chat photo. This file_id can be used only for photo
	// download and only for as long as the photo is not changed.
	SmallFileID string `json:"small_file_id,omitempty"`

	// Unique file identifier of small (160x160) chat photo, which is supposed to be the same
	// over time and for different bots. Can't be used to download or reuse the file.
	SmallFileUniqueID string `json:"small_file_unique_id,omitempty"`

	// File identifier of big (640x640) chat photo. This file_id can be used only for photo
	// download and only for as long as the photo is not changed.
	BigFileID string `json:"big_file_id,omitempty"`

	// Unique file identifier of big (640x640) chat photo, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	BigFileUniqueID string `json:"big_file_unique_id,omitempty"`
}

// Represents an invite link for a chat.
type ChatInviteLink struct {
	// The invite link. If the link was created by another chat administrator, then the second
	// part of the link will be replaced with “…”.
	InviteLink string `json:"invite_link,omitempty"`

	// Creator of the link
	Creator *User `json:"creator,omitempty"`

	// True, if users joining the chat via the link need to be approved by chat administrators
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`

	// True, if the link is primary
	IsPrimary bool `json:"is_primary,omitempty"`

	// True, if the link is revoked
	IsRevoked bool `json:"is_revoked,omitempty"`

	// Optional. Invite link name
	Name string `json:"name,omitempty"`

	// Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	ExpireDate int64 `json:"expire_date,omitempty"`

	// Optional. The maximum number of users that can be members of the chat simultaneously after
	// joining the chat via this invite link; 1-99999
	MemberLimit int64 `json:"member_limit,omitempty"`

	// Optional. Number of pending join requests created using this link
	PendingJoinRequestCount int64 `json:"pending_join_request_count,omitempty"`
}

// Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// True, if the administrator can access the chat event log, boost list in channels, see
	// channel members, report spam messages, see anonymous administrators in supergroups and
	// ignore slow mode. Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat,omitempty"`

	// True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

	// True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`

	// True, if the administrator can restrict, ban or unban chat members, or access supergroup
	// statistics
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

	// True, if the administrator can add new administrators with a subset of their own
	// privileges or demote administrators that they have promoted, directly or indirectly
	// (promoted by administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`

	// True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// Optional. True, if the administrator can post messages in the channel, or access channel
	// statistics; channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// Optional. True, if the administrator can edit messages of other users and can pin
	// messages; channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// Optional. True, if the administrator can post stories in the channel; channels only
	CanPostStories bool `json:"can_post_stories,omitempty"`

	// Optional. True, if the administrator can edit stories posted by other users; channels only
	CanEditStories bool `json:"can_edit_stories,omitempty"`

	// Optional. True, if the administrator can delete stories posted by other users; channels
	// only
	CanDeleteStories bool `json:"can_delete_stories,omitempty"`

	// Optional. True, if the user is allowed to create, rename, close, and reopen forum topics;
	// supergroups only
	CanManageTopics bool `json:"can_manage_topics,omitempty"`
}

// This object represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	// Chat the user belongs to
	Chat *Chat `json:"chat,omitempty"`

	// Performer of the action, which resulted in the change
	From *User `json:"from,omitempty"`

	// Date the change was done in Unix time
	Date int64 `json:"date,omitempty"`

	// Previous information about the chat member
	OldChatMember *ChatMember `json:"old_chat_member,omitempty"`

	// New information about the chat member
	NewChatMember *ChatMember `json:"new_chat_member,omitempty"`

	// Optional. Chat invite link, which was used by the user to join the chat; for joining by
	// invite link events only.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`

	// Optional. True, if the user joined the chat via a chat folder invite link
	ViaChatFolderInviteLink bool `json:"via_chat_folder_invite_link,omitempty"`
}

// This object contains information about one member of a chat. Currently, the following 6
// types of chat members are supported:   ChatMemberOwner  ChatMemberAdministrator
// ChatMemberMember  ChatMemberRestricted  ChatMemberLeft  ChatMemberBanned
type ChatMember struct {
}

// Represents a chat member that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
	// The member's status in the chat, always “creator”
	Status string `json:"status,omitempty"`

	// Information about the user
	User *User `json:"user,omitempty"`

	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// Optional. Custom title for this user
	CustomTitle string `json:"custom_title,omitempty"`
}

// Represents a chat member that has some additional privileges.
type ChatMemberAdministrator struct {
	// The member's status in the chat, always “administrator”
	Status string `json:"status,omitempty"`

	// Information about the user
	User *User `json:"user,omitempty"`

	// True, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited,omitempty"`

	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// True, if the administrator can access the chat event log, boost list in channels, see
	// channel members, report spam messages, see anonymous administrators in supergroups and
	// ignore slow mode. Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat,omitempty"`

	// True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

	// True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`

	// True, if the administrator can restrict, ban or unban chat members, or access supergroup
	// statistics
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

	// True, if the administrator can add new administrators with a subset of their own
	// privileges or demote administrators that they have promoted, directly or indirectly
	// (promoted by administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`

	// True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// Optional. True, if the administrator can post messages in the channel, or access channel
	// statistics; channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// Optional. True, if the administrator can edit messages of other users and can pin
	// messages; channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// Optional. True, if the administrator can post stories in the channel; channels only
	CanPostStories bool `json:"can_post_stories,omitempty"`

	// Optional. True, if the administrator can edit stories posted by other users; channels only
	CanEditStories bool `json:"can_edit_stories,omitempty"`

	// Optional. True, if the administrator can delete stories posted by other users; channels
	// only
	CanDeleteStories bool `json:"can_delete_stories,omitempty"`

	// Optional. True, if the user is allowed to create, rename, close, and reopen forum topics;
	// supergroups only
	CanManageTopics bool `json:"can_manage_topics,omitempty"`

	// Optional. Custom title for this user
	CustomTitle string `json:"custom_title,omitempty"`
}

// Represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
	// The member's status in the chat, always “member”
	Status string `json:"status,omitempty"`

	// Information about the user
	User *User `json:"user,omitempty"`
}

// Represents a chat member that is under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
	// The member's status in the chat, always “restricted”
	Status string `json:"status,omitempty"`

	// Information about the user
	User *User `json:"user,omitempty"`

	// True, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member,omitempty"`

	// True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners,
	// invoices, locations and venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`

	// True, if the user is allowed to send audios
	CanSendAudios bool `json:"can_send_audios,omitempty"`

	// True, if the user is allowed to send documents
	CanSendDocuments bool `json:"can_send_documents,omitempty"`

	// True, if the user is allowed to send photos
	CanSendPhotos bool `json:"can_send_photos,omitempty"`

	// True, if the user is allowed to send videos
	CanSendVideos bool `json:"can_send_videos,omitempty"`

	// True, if the user is allowed to send video notes
	CanSendVideoNotes bool `json:"can_send_video_notes,omitempty"`

	// True, if the user is allowed to send voice notes
	CanSendVoiceNotes bool `json:"can_send_voice_notes,omitempty"`

	// True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls,omitempty"`

	// True, if the user is allowed to send animations, games, stickers and use inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// True, if the user is allowed to pin messages
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// True, if the user is allowed to create forum topics
	CanManageTopics bool `json:"can_manage_topics,omitempty"`

	// Date when restrictions will be lifted for this user; Unix time. If 0, then the user is
	// restricted forever
	UntilDate int64 `json:"until_date,omitempty"`
}

// Represents a chat member that isn't currently a member of the chat, but may join it
// themselves.
type ChatMemberLeft struct {
	// The member's status in the chat, always “left”
	Status string `json:"status,omitempty"`

	// Information about the user
	User *User `json:"user,omitempty"`
}

// Represents a chat member that was banned in the chat and can't return to the chat or view
// chat messages.
type ChatMemberBanned struct {
	// The member's status in the chat, always “kicked”
	Status string `json:"status,omitempty"`

	// Information about the user
	User *User `json:"user,omitempty"`

	// Date when restrictions will be lifted for this user; Unix time. If 0, then the user is
	// banned forever
	UntilDate int64 `json:"until_date,omitempty"`
}

// Represents a join request sent to a chat.
type ChatJoinRequest struct {
	// Chat to which the request was sent
	Chat *Chat `json:"chat,omitempty"`

	// User that sent the join request
	From *User `json:"from,omitempty"`

	// Identifier of a private chat with the user who sent the join request. This number may have
	// more than 32 significant bits and some programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or
	// double-precision float type are safe for storing this identifier. The bot can use this
	// identifier for 5 minutes to send messages until the join request is processed, assuming no
	// other administrator contacted the user.
	UserChatID int64 `json:"user_chat_id,omitempty"`

	// Date the request was sent in Unix time
	Date int64 `json:"date,omitempty"`

	// Optional. Bio of the user.
	Bio string `json:"bio,omitempty"`

	// Optional. Chat invite link that was used by the user to send the join request
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// Optional. True, if the user is allowed to send text messages, contacts, giveaways,
	// giveaway winners, invoices, locations and venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`

	// Optional. True, if the user is allowed to send audios
	CanSendAudios bool `json:"can_send_audios,omitempty"`

	// Optional. True, if the user is allowed to send documents
	CanSendDocuments bool `json:"can_send_documents,omitempty"`

	// Optional. True, if the user is allowed to send photos
	CanSendPhotos bool `json:"can_send_photos,omitempty"`

	// Optional. True, if the user is allowed to send videos
	CanSendVideos bool `json:"can_send_videos,omitempty"`

	// Optional. True, if the user is allowed to send video notes
	CanSendVideoNotes bool `json:"can_send_video_notes,omitempty"`

	// Optional. True, if the user is allowed to send voice notes
	CanSendVoiceNotes bool `json:"can_send_voice_notes,omitempty"`

	// Optional. True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls,omitempty"`

	// Optional. True, if the user is allowed to send animations, games, stickers and use inline
	// bots
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// Optional. True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// Optional. True, if the user is allowed to change the chat title, photo and other settings.
	// Ignored in public supergroups
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// Optional. True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// Optional. True, if the user is allowed to create forum topics. If omitted defaults to the
	// value of can_pin_messages
	CanManageTopics bool `json:"can_manage_topics,omitempty"`
}

// Represents a location to which a chat is connected.
type ChatLocation struct {
	// The location to which the supergroup is connected. Can't be a live location.
	Location *Location `json:"location,omitempty"`

	// Location address; 1-64 characters, as defined by the chat owner
	Address string `json:"address,omitempty"`
}

// The reaction is based on an emoji.
type ReactionTypeEmoji struct {
	// Type of the reaction, always “emoji”
	Type string `json:"type,omitempty"`

	// Reaction emoji. Currently, it can be one of "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
	Emoji string `json:"emoji,omitempty"`
}

// The reaction is based on a custom emoji.
type ReactionTypeCustomEmoji struct {
	// Type of the reaction, always “custom_emoji”
	Type string `json:"type,omitempty"`

	// Custom emoji identifier
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// Represents a reaction added to a message along with the number of times it was added.
type ReactionCount struct {
	// Type of the reaction
	Type *ReactionType `json:"type,omitempty"`

	// Number of times the reaction was added
	TotalCount int64 `json:"total_count,omitempty"`
}

// This object represents a change of a reaction on a message performed by a user.
type MessageReactionUpdated struct {
	// The chat containing the message the user reacted to
	Chat *Chat `json:"chat,omitempty"`

	// Unique identifier of the message inside the chat
	MessageID int64 `json:"message_id,omitempty"`

	// Optional. The user that changed the reaction, if the user isn't anonymous
	User *User `json:"user,omitempty"`

	// Optional. The chat on behalf of which the reaction was changed, if the user is anonymous
	ActorChat *Chat `json:"actor_chat,omitempty"`

	// Date of the change in Unix time
	Date int64 `json:"date,omitempty"`

	// Previous list of reaction types that were set by the user
	OldReaction []*ReactionType `json:"old_reaction,omitempty"`

	// New list of reaction types that have been set by the user
	NewReaction []*ReactionType `json:"new_reaction,omitempty"`
}

// This object represents reaction changes on a message with anonymous reactions.
type MessageReactionCountUpdated struct {
	// The chat containing the message
	Chat *Chat `json:"chat,omitempty"`

	// Unique message identifier inside the chat
	MessageID int64 `json:"message_id,omitempty"`

	// Date of the change in Unix time
	Date int64 `json:"date,omitempty"`

	// List of reactions that are present on the message
	Reactions []*ReactionCount `json:"reactions,omitempty"`
}

// This object represents a forum topic.
type ForumTopic struct {
	// Unique identifier of the forum topic
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Name of the topic
	Name string `json:"name,omitempty"`

	// Color of the topic icon in RGB format
	IconColor int64 `json:"icon_color,omitempty"`

	// Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// This object represents a bot command.
type BotCommand struct {
	// Text of the command; 1-32 characters. Can contain only lowercase English letters, digits
	// and underscores.
	Command string `json:"command,omitempty"`

	// Description of the command; 1-256 characters.
	Description string `json:"description,omitempty"`
}

// This object represents the scope to which bot commands are applied. Currently, the following
// 7 scopes are supported:   BotCommandScopeDefault  BotCommandScopeAllPrivateChats
// BotCommandScopeAllGroupChats  BotCommandScopeAllChatAdministrators  BotCommandScopeChat
// BotCommandScopeChatAdministrators  BotCommandScopeChatMember
type BotCommandScope struct {
}

// Represents the default scope of bot commands. Default commands are used if no commands with
// a narrower scope are specified for the user.
type BotCommandScopeDefault struct {
	// Scope type, must be default
	Type string `json:"type,omitempty"`
}

// Represents the scope of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {
	// Scope type, must be all_private_chats
	Type string `json:"type,omitempty"`
}

// Represents the scope of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
	// Scope type, must be all_group_chats
	Type string `json:"type,omitempty"`
}

// Represents the scope of bot commands, covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct {
	// Scope type, must be all_chat_administrators
	Type string `json:"type,omitempty"`
}

// Represents the scope of bot commands, covering a specific chat.
type BotCommandScopeChat struct {
	// Scope type, must be chat
	Type string `json:"type,omitempty"`

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Represents the scope of bot commands, covering all administrators of a specific group or
// supergroup chat.
type BotCommandScopeChatAdministrators struct {
	// Scope type, must be chat_administrators
	Type string `json:"type,omitempty"`

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Represents the scope of bot commands, covering a specific member of a group or supergroup
// chat.
type BotCommandScopeChatMember struct {
	// Scope type, must be chat_member
	Type string `json:"type,omitempty"`

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id,omitempty"`
}

// This object represents the bot's name.
type BotName struct {
	// The bot's name
	Name string `json:"name,omitempty"`
}

// This object represents the bot's description.
type BotDescription struct {
	// The bot's description
	Description string `json:"description,omitempty"`
}

// This object represents the bot's short description.
type BotShortDescription struct {
	// The bot's short description
	ShortDescription string `json:"short_description,omitempty"`
}

// This object describes the bot's menu button in a private chat. It should be one of
// MenuButtonCommands  MenuButtonWebApp  MenuButtonDefault   If a menu button other than
// MenuButtonDefault is set for a private chat, then it is applied in the chat. Otherwise the
// default menu button is applied. By default, the menu button opens the list of bot commands.
type MenuButton struct {
}

// Represents a menu button, which opens the bot's list of commands.
type MenuButtonCommands struct {
	// Type of the button, must be commands
	Type string `json:"type,omitempty"`
}

// Represents a menu button, which launches a Web App .
type MenuButtonWebApp struct {
	// Type of the button, must be web_app
	Type string `json:"type,omitempty"`

	// Text on the button
	Text string `json:"text,omitempty"`

	// Description of the Web App that will be launched when the user presses the button. The Web
	// App will be able to send an arbitrary message on behalf of the user using the method
	// answerWebAppQuery.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

// Describes that no specific value for the menu button was set.
type MenuButtonDefault struct {
	// Type of the button, must be default
	Type string `json:"type,omitempty"`
}

// This object describes the source of a chat boost. It can be one of   ChatBoostSourcePremium
// ChatBoostSourceGiftCode  ChatBoostSourceGiveaway
type ChatBoostSource struct {
}

// The boost was obtained by subscribing to Telegram Premium or by gifting a Telegram Premium
// subscription to another user.
type ChatBoostSourcePremium struct {
	// Source of the boost, always “premium”
	Source string `json:"source,omitempty"`

	// User that boosted the chat
	User *User `json:"user,omitempty"`
}

// The boost was obtained by the creation of Telegram Premium gift codes to boost a chat. Each
// such code boosts the chat 4 times for the duration of the corresponding Telegram Premium
// subscription.
type ChatBoostSourceGiftCode struct {
	// Source of the boost, always “gift_code”
	Source string `json:"source,omitempty"`

	// User for which the gift code was created
	User *User `json:"user,omitempty"`
}

// The boost was obtained by the creation of a Telegram Premium giveaway. This boosts the chat
// 4 times for the duration of the corresponding Telegram Premium subscription.
type ChatBoostSourceGiveaway struct {
	// Source of the boost, always “giveaway”
	Source string `json:"source,omitempty"`

	// Identifier of a message in the chat with the giveaway; the message could have been deleted
	// already. May be 0 if the message isn't sent yet.
	GiveawayMessageID int64 `json:"giveaway_message_id,omitempty"`

	// Optional. User that won the prize in the giveaway if any
	User *User `json:"user,omitempty"`

	// Optional. True, if the giveaway was completed, but there was no user to win the prize
	IsUnclaimed bool `json:"is_unclaimed,omitempty"`
}

// This object contains information about a chat boost.
type ChatBoost struct {
	// Unique identifier of the boost
	BoostID string `json:"boost_id,omitempty"`

	// Point in time (Unix timestamp) when the chat was boosted
	AddDate int64 `json:"add_date,omitempty"`

	// Point in time (Unix timestamp) when the boost will automatically expire, unless the
	// booster's Telegram Premium subscription is prolonged
	ExpirationDate int64 `json:"expiration_date,omitempty"`

	// Source of the added boost
	Source *ChatBoostSource `json:"source,omitempty"`
}

// This object represents a boost added to a chat or changed.
type ChatBoostUpdated struct {
	// Chat which was boosted
	Chat *Chat `json:"chat,omitempty"`

	// Infomation about the chat boost
	Boost *ChatBoost `json:"boost,omitempty"`
}

// This object represents a boost removed from a chat.
type ChatBoostRemoved struct {
	// Chat which was boosted
	Chat *Chat `json:"chat,omitempty"`

	// Unique identifier of the boost
	BoostID string `json:"boost_id,omitempty"`

	// Point in time (Unix timestamp) when the boost was removed
	RemoveDate int64 `json:"remove_date,omitempty"`

	// Source of the removed boost
	Source *ChatBoostSource `json:"source,omitempty"`
}

// This object represents a list of boosts added to a chat by a user.
type UserChatBoosts struct {
	// The list of boosts added to the chat by the user
	Boosts []*ChatBoost `json:"boosts,omitempty"`
}

// Describes why a request was unsuccessful.
type ResponseParameters struct {
	// Optional. The group has been migrated to a supergroup with the specified identifier. This
	// number may have more than 32 significant bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// Optional. In case of exceeding flood control, the number of seconds left to wait before
	// the request can be repeated
	RetryAfter int64 `json:"retry_after,omitempty"`
}

// This object represents the content of a media message to be sent. It should be one of
// InputMediaAnimation  InputMediaDocument  InputMediaAudio  InputMediaPhoto  InputMediaVideo
type InputMedia struct {
}

// Represents a photo to be sent.
type InputMediaPhoto struct {
	// Type of the result, must be photo
	Type string `json:"type,omitempty"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name. More information on Sending Files »
	Media string `json:"media,omitempty"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the photo caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// Represents a video to be sent.
type InputMediaVideo struct {
	// Type of the result, must be video
	Type string `json:"type,omitempty"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name. More information on Sending Files »
	Media string `json:"media,omitempty"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file
	// is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in
	// size. A thumbnail's width and height should not exceed 320. Ignored if the file is not
	// uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as
	// a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded
	// using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail,omitempty"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the video caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Video width
	Width int64 `json:"width,omitempty"`

	// Optional. Video height
	Height int64 `json:"height,omitempty"`

	// Optional. Video duration in seconds
	Duration int64 `json:"duration,omitempty"`

	// Optional. Pass True if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`

	// Optional. Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	// Type of the result, must be animation
	Type string `json:"type,omitempty"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name. More information on Sending Files »
	Media string `json:"media,omitempty"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file
	// is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in
	// size. A thumbnail's width and height should not exceed 320. Ignored if the file is not
	// uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as
	// a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded
	// using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail,omitempty"`

	// Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the animation caption. See formatting options for
	// more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Animation width
	Width int64 `json:"width,omitempty"`

	// Optional. Animation height
	Height int64 `json:"height,omitempty"`

	// Optional. Animation duration in seconds
	Duration int64 `json:"duration,omitempty"`

	// Optional. Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	// Type of the result, must be audio
	Type string `json:"type,omitempty"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name. More information on Sending Files »
	Media string `json:"media,omitempty"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file
	// is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in
	// size. A thumbnail's width and height should not exceed 320. Ignored if the file is not
	// uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as
	// a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded
	// using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail,omitempty"`

	// Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the audio caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Duration of the audio in seconds
	Duration int64 `json:"duration,omitempty"`

	// Optional. Performer of the audio
	Performer string `json:"performer,omitempty"`

	// Optional. Title of the audio
	Title string `json:"title,omitempty"`
}

// Represents a general file to be sent.
type InputMediaDocument struct {
	// Type of the result, must be document
	Type string `json:"type,omitempty"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name. More information on Sending Files »
	Media string `json:"media,omitempty"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file
	// is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in
	// size. A thumbnail's width and height should not exceed 320. Ignored if the file is not
	// uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as
	// a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded
	// using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail,omitempty"`

	// Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the document caption. See formatting options for
	// more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Disables automatic server-side content type detection for files uploaded using
	// multipart/form-data. Always True, if the document is sent as part of an album.
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
}

// This object represents the contents of a file to be uploaded. Must be posted using
// multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile struct {
}

// The following methods and objects allow your bot to handle stickers and sticker sets.   This
// object represents a sticker.
type Sticker struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id,omitempty"`

	// Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"`

	// Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”. The type of the
	// sticker is independent from its format, which is determined by the fields is_animated and
	// is_video.
	Type string `json:"type,omitempty"`

	// Sticker width
	Width int64 `json:"width,omitempty"`

	// Sticker height
	Height int64 `json:"height,omitempty"`

	// True, if the sticker is animated
	IsAnimated bool `json:"is_animated,omitempty"`

	// True, if the sticker is a video sticker
	IsVideo bool `json:"is_video,omitempty"`

	// Optional. Sticker thumbnail in the .WEBP or .JPG format
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// Optional. Emoji associated with the sticker
	Emoji string `json:"emoji,omitempty"`

	// Optional. Name of the sticker set to which the sticker belongs
	SetName string `json:"set_name,omitempty"`

	// Optional. For premium regular stickers, premium animation for the sticker
	PremiumAnimation *File `json:"premium_animation,omitempty"`

	// Optional. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// Optional. For custom emoji stickers, unique identifier of the custom emoji
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`

	// Optional. True, if the sticker must be repainted to a text color in messages, the color of
	// the Telegram Premium badge in emoji status, white color on chat photos, or another
	// appropriate color in other places
	NeedsRepainting bool `json:"needs_repainting,omitempty"`

	// Optional. File size in bytes
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents a sticker set.
type StickerSet struct {
	// Sticker set name
	Name string `json:"name,omitempty"`

	// Sticker set title
	Title string `json:"title,omitempty"`

	// Type of stickers in the set, currently one of “regular”, “mask”, “custom_emoji”
	StickerType string `json:"sticker_type,omitempty"`

	// True, if the sticker set contains animated stickers
	IsAnimated bool `json:"is_animated,omitempty"`

	// True, if the sticker set contains video stickers
	IsVideo bool `json:"is_video,omitempty"`

	// List of all set stickers
	Stickers []*Sticker `json:"stickers,omitempty"`

	// Optional. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

// This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed. One of “forehead”,
	// “eyes”, “mouth”, or “chin”.
	Point string `json:"point,omitempty"`

	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to
	// right. For example, choosing -1.0 will place mask just to the left of the default mask
	// position.
	XShift float32 `json:"x_shift,omitempty"`

	// Shift by Y-axis measured in heights of the mask scaled to the face size, from top to
	// bottom. For example, 1.0 will place the mask just below the default mask position.
	YShift float32 `json:"y_shift,omitempty"`

	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float32 `json:"scale,omitempty"`
}

// This object describes a sticker to be added to a sticker set.
type InputSticker struct {
	// The added sticker. Pass a file_id as a String to send a file that already exists on the
	// Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the
	// Internet, upload a new one using multipart/form-data, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name. Animated and video stickers can't be uploaded via HTTP URL. More
	// information on Sending Files »
	Sticker interface{} `json:"sticker,omitempty"`

	// List of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list,omitempty"`

	// Optional. Position where the mask should be placed on faces. For “mask” stickers only.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// Optional. List of 0-20 search keywords for the sticker with total length of up to 64
	// characters. For “regular” and “custom_emoji” stickers only.
	Keywords []string `json:"keywords,omitempty"`
}

// The following methods and objects allow your bot to work in inline mode . Please see our
// Introduction to Inline bots for more details.  To enable this option, send the /setinline
// command to @BotFather and provide the placeholder text that the user will see in the input
// field after typing your bot's name.   This object represents an incoming inline query. When
// the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	// Unique identifier for this query
	ID string `json:"id,omitempty"`

	// Sender
	From *User `json:"from,omitempty"`

	// Text of the query (up to 256 characters)
	Query string `json:"query,omitempty"`

	// Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset,omitempty"`

	// Optional. Type of the chat from which the inline query was sent. Can be either “sender”
	// for a private chat with the inline query sender, “private”, “group”, “supergroup”, or
	// “channel”. The chat type should be always known for requests sent from official clients
	// and most third-party clients, unless the request was sent from a secret chat
	ChatType string `json:"chat_type,omitempty"`

	// Optional. Sender location, only for bots that request user location
	Location *Location `json:"location,omitempty"`
}

// This object represents a button to be shown above inline query results. You must use exactly
// one of the optional fields.
type InlineQueryResultsButton struct {
	// Label text on the button
	Text string `json:"text,omitempty"`

	// Optional. Description of the Web App that will be launched when the user presses the
	// button. The Web App will be able to switch back to the inline mode using the method
	// switchInlineQuery inside the Web App.
	WebApp *WebAppInfo `json:"web_app,omitempty"`

	// Optional. Deep-linking parameter for the /start message sent to the bot when a user
	// presses the button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An
	// inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube
	// account to adapt search results accordingly. To do this, it displays a 'Connect your
	// YouTube account' button above the results, or even before showing any. The user presses
	// the button, switches to a private chat with the bot and, in doing so, passes a start
	// parameter that instructs the bot to return an OAuth link. Once done, the bot can offer a
	// switch_inline button so that the user can easily return to the chat where they wanted to
	// use the bot's inline capabilities.
	StartParameter string `json:"start_parameter,omitempty"`
}

// This object represents one result of an inline query. Telegram clients currently support
// results of the following 20 types:   InlineQueryResultCachedAudio
// InlineQueryResultCachedDocument  InlineQueryResultCachedGif  InlineQueryResultCachedMpeg4Gif
// InlineQueryResultCachedPhoto  InlineQueryResultCachedSticker  InlineQueryResultCachedVideo
// InlineQueryResultCachedVoice  InlineQueryResultArticle  InlineQueryResultAudio
// InlineQueryResultContact  InlineQueryResultGame  InlineQueryResultDocument
// InlineQueryResultGif  InlineQueryResultLocation  InlineQueryResultMpeg4Gif
// InlineQueryResultPhoto  InlineQueryResultVenue  InlineQueryResultVideo
// InlineQueryResultVoice   Note: All URLs passed in inline query results will be available to
// end users and therefore must be assumed to be public .
type InlineQueryResult struct {
}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	// Type of the result, must be article
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 Bytes
	ID string `json:"id,omitempty"`

	// Title of the result
	Title string `json:"title,omitempty"`

	// Content of the message to be sent
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. URL of the result
	URL string `json:"url,omitempty"`

	// Optional. Pass True if you don't want the URL to be shown in the message
	HideURL bool `json:"hide_url,omitempty"`

	// Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Optional. Url of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// Optional. Thumbnail width
	ThumbnailWidth int64 `json:"thumbnail_width,omitempty"`

	// Optional. Thumbnail height
	ThumbnailHeight int64 `json:"thumbnail_height,omitempty"`
}

// Represents a link to a photo. By default, this photo will be sent by the user with optional
// caption. Alternatively, you can use input_message_content to send a message with the
// specified content instead of the photo.
type InlineQueryResultPhoto struct {
	// Type of the result, must be photo
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB
	PhotoURL string `json:"photo_url,omitempty"`

	// URL of the thumbnail for the photo
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// Optional. Width of the photo
	PhotoWidth int64 `json:"photo_width,omitempty"`

	// Optional. Height of the photo
	PhotoHeight int64 `json:"photo_height,omitempty"`

	// Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the photo caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the photo
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to an animated GIF file. By default, this animated GIF file will be sent
// by the user with optional caption. Alternatively, you can use input_message_content to send
// a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	// Type of the result, must be gif
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid URL for the GIF file. File size must not exceed 1MB
	GifURL string `json:"gif_url,omitempty"`

	// Optional. Width of the GIF
	GifWidth int64 `json:"gif_width,omitempty"`

	// Optional. Height of the GIF
	GifHeight int64 `json:"gif_height,omitempty"`

	// Optional. Duration of the GIF in seconds
	GifDuration int64 `json:"gif_duration,omitempty"`

	// URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or
	// “video/mp4”. Defaults to “image/jpeg”
	ThumbnailMimeType string `json:"thumbnail_mime_type,omitempty"`

	// Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default,
// this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you
// can use input_message_content to send a message with the specified content instead of the
// animation.
type InlineQueryResultMpeg4Gif struct {
	// Type of the result, must be mpeg4_gif
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid URL for the MPEG4 file. File size must not exceed 1MB
	Mpeg4URL string `json:"mpeg4_url,omitempty"`

	// Optional. Video width
	Mpeg4Width int64 `json:"mpeg4_width,omitempty"`

	// Optional. Video height
	Mpeg4Height int64 `json:"mpeg4_height,omitempty"`

	// Optional. Video duration in seconds
	Mpeg4Duration int64 `json:"mpeg4_duration,omitempty"`

	// URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or
	// “video/mp4”. Defaults to “image/jpeg”
	ThumbnailMimeType string `json:"thumbnail_mime_type,omitempty"`

	// Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the video animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a page containing an embedded video player or a video file. By default,
// this video file will be sent by the user with an optional caption. Alternatively, you can
// use input_message_content to send a message with the specified content instead of the video.
// If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube), you must
// replace its content using input_message_content .
type InlineQueryResultVideo struct {
	// Type of the result, must be video
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid URL for the embedded video player or video file
	VideoURL string `json:"video_url,omitempty"`

	// MIME type of the content of the video URL, “text/html” or “video/mp4”
	MimeType string `json:"mime_type,omitempty"`

	// URL of the thumbnail (JPEG only) for the video
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the video caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Video width
	VideoWidth int64 `json:"video_width,omitempty"`

	// Optional. Video height
	VideoHeight int64 `json:"video_height,omitempty"`

	// Optional. Video duration in seconds
	VideoDuration int64 `json:"video_duration,omitempty"`

	// Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the video. This field is required
	// if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube
	// video).
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to an MP3 audio file. By default, this audio file will be sent by the
// user. Alternatively, you can use input_message_content to send a message with the specified
// content instead of the audio.
type InlineQueryResultAudio struct {
	// Type of the result, must be audio
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid URL for the audio file
	AudioURL string `json:"audio_url,omitempty"`

	// Title
	Title string `json:"title,omitempty"`

	// Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the audio caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Performer
	Performer string `json:"performer,omitempty"`

	// Optional. Audio duration in seconds
	AudioDuration int64 `json:"audio_duration,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the audio
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a voice recording in an .OGG container encoded with OPUS. By default,
// this voice recording will be sent by the user. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the the voice
// message.
type InlineQueryResultVoice struct {
	// Type of the result, must be voice
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid URL for the voice recording
	VoiceURL string `json:"voice_url,omitempty"`

	// Recording title
	Title string `json:"title,omitempty"`

	// Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the voice message caption. See formatting options
	// for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Recording duration in seconds
	VoiceDuration int64 `json:"voice_duration,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the voice recording
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a file. By default, this file will be sent by the user with an optional
// caption. Alternatively, you can use input_message_content to send a message with the
// specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using
// this method.
type InlineQueryResultDocument struct {
	// Type of the result, must be document
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the document caption. See formatting options for
	// more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// A valid URL for the file
	DocumentURL string `json:"document_url,omitempty"`

	// MIME type of the content of the file, either “application/pdf” or “application/zip”
	MimeType string `json:"mime_type,omitempty"`

	// Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the file
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`

	// Optional. URL of the thumbnail (JPEG only) for the file
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// Optional. Thumbnail width
	ThumbnailWidth int64 `json:"thumbnail_width,omitempty"`

	// Optional. Thumbnail height
	ThumbnailHeight int64 `json:"thumbnail_height,omitempty"`
}

// Represents a location on a map. By default, the location will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified
// content instead of the location.
type InlineQueryResultLocation struct {
	// Type of the result, must be location
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 Bytes
	ID string `json:"id,omitempty"`

	// Location latitude in degrees
	Latitude float32 `json:"latitude,omitempty"`

	// Location longitude in degrees
	Longitude float32 `json:"longitude,omitempty"`

	// Location title
	Title string `json:"title,omitempty"`

	// Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`

	// Optional. Period in seconds for which the location can be updated, should be between 60
	// and 86400.
	LivePeriod int64 `json:"live_period,omitempty"`

	// Optional. For live locations, a direction in which the user is moving, in degrees. Must be
	// between 1 and 360 if specified.
	Heading int64 `json:"heading,omitempty"`

	// Optional. For live locations, a maximum distance for proximity alerts about approaching
	// another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the location
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`

	// Optional. Url of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// Optional. Thumbnail width
	ThumbnailWidth int64 `json:"thumbnail_width,omitempty"`

	// Optional. Thumbnail height
	ThumbnailHeight int64 `json:"thumbnail_height,omitempty"`
}

// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can
// use input_message_content to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	// Type of the result, must be venue
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 Bytes
	ID string `json:"id,omitempty"`

	// Latitude of the venue location in degrees
	Latitude float32 `json:"latitude,omitempty"`

	// Longitude of the venue location in degrees
	Longitude float32 `json:"longitude,omitempty"`

	// Title of the venue
	Title string `json:"title,omitempty"`

	// Address of the venue
	Address string `json:"address,omitempty"`

	// Optional. Foursquare identifier of the venue if known
	FoursquareID string `json:"foursquare_id,omitempty"`

	// Optional. Foursquare type of the venue, if known. (For example,
	// “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// Optional. Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the venue
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`

	// Optional. Url of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// Optional. Thumbnail width
	ThumbnailWidth int64 `json:"thumbnail_width,omitempty"`

	// Optional. Thumbnail height
	ThumbnailHeight int64 `json:"thumbnail_height,omitempty"`
}

// Represents a contact with a phone number. By default, this contact will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified
// content instead of the contact.
type InlineQueryResultContact struct {
	// Type of the result, must be contact
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 Bytes
	ID string `json:"id,omitempty"`

	// Contact's phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// Contact's first name
	FirstName string `json:"first_name,omitempty"`

	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	Vcard string `json:"vcard,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the contact
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`

	// Optional. Url of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// Optional. Thumbnail width
	ThumbnailWidth int64 `json:"thumbnail_width,omitempty"`

	// Optional. Thumbnail height
	ThumbnailHeight int64 `json:"thumbnail_height,omitempty"`
}

// Represents a Game .
type InlineQueryResultGame struct {
	// Type of the result, must be game
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// Short name of the game
	GameShortName string `json:"game_short_name,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Represents a link to a photo stored on the Telegram servers. By default, this photo will be
// sent by the user with an optional caption. Alternatively, you can use input_message_content
// to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	// Type of the result, must be photo
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid file identifier of the photo
	PhotoFileID string `json:"photo_file_id,omitempty"`

	// Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the photo caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the photo
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to an animated GIF file stored on the Telegram servers. By default, this
// animated GIF file will be sent by the user with an optional caption. Alternatively, you can
// use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	// Type of the result, must be gif
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid file identifier for the GIF file
	GifFileID string `json:"gif_file_id,omitempty"`

	// Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the
// Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an
// optional caption. Alternatively, you can use input_message_content to send a message with
// the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	// Type of the result, must be mpeg4_gif
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid file identifier for the MPEG4 file
	Mpeg4FileID string `json:"mpeg4_file_id,omitempty"`

	// Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the video animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will
// be sent by the user. Alternatively, you can use input_message_content to send a message with
// the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	// Type of the result, must be sticker
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid file identifier of the sticker
	StickerFileID string `json:"sticker_file_id,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the sticker
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a file stored on the Telegram servers. By default, this file will be
// sent by the user with an optional caption. Alternatively, you can use input_message_content
// to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
	// Type of the result, must be document
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// Title for the result
	Title string `json:"title,omitempty"`

	// A valid file identifier for the file
	DocumentFileID string `json:"document_file_id,omitempty"`

	// Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the document caption. See formatting options for
	// more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the file
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a video file stored on the Telegram servers. By default, this video
// file will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	// Type of the result, must be video
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid file identifier for the video file
	VideoFileID string `json:"video_file_id,omitempty"`

	// Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the video caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the video
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a voice message stored on the Telegram servers. By default, this voice
// message will be sent by the user. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	// Type of the result, must be voice
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid file identifier for the voice message
	VoiceFileID string `json:"voice_file_id,omitempty"`

	// Voice message title
	Title string `json:"title,omitempty"`

	// Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the voice message caption. See formatting options
	// for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the voice message
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to an MP3 audio file stored on the Telegram servers. By default, this
// audio file will be sent by the user. Alternatively, you can use input_message_content to
// send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	// Type of the result, must be audio
	Type string `json:"type,omitempty"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id,omitempty"`

	// A valid file identifier for the audio file
	AudioFileID string `json:"audio_file_id,omitempty"`

	// Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the audio caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Optional. Content of the message to be sent instead of the audio
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// This object represents the content of a message to be sent as a result of an inline query.
// Telegram clients currently support the following 5 types:   InputTextMessageContent
// InputLocationMessageContent  InputVenueMessageContent  InputContactMessageContent
// InputInvoiceMessageContent
type InputMessageContent struct {
}

// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	// Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text,omitempty"`

	// Optional. Mode for parsing entities in the message text. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in message text, which can be specified
	// instead of parse_mode
	Entities []*MessageEntity `json:"entities,omitempty"`

	// Optional. Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	// Latitude of the location in degrees
	Latitude float32 `json:"latitude,omitempty"`

	// Longitude of the location in degrees
	Longitude float32 `json:"longitude,omitempty"`

	// Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`

	// Optional. Period in seconds for which the location can be updated, should be between 60
	// and 86400.
	LivePeriod int64 `json:"live_period,omitempty"`

	// Optional. For live locations, a direction in which the user is moving, in degrees. Must be
	// between 1 and 360 if specified.
	Heading int64 `json:"heading,omitempty"`

	// Optional. For live locations, a maximum distance for proximity alerts about approaching
	// another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
}

// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	// Latitude of the venue in degrees
	Latitude float32 `json:"latitude,omitempty"`

	// Longitude of the venue in degrees
	Longitude float32 `json:"longitude,omitempty"`

	// Name of the venue
	Title string `json:"title,omitempty"`

	// Address of the venue
	Address string `json:"address,omitempty"`

	// Optional. Foursquare identifier of the venue, if known
	FoursquareID string `json:"foursquare_id,omitempty"`

	// Optional. Foursquare type of the venue, if known. (For example,
	// “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// Optional. Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// Contact's first name
	FirstName string `json:"first_name,omitempty"`

	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	Vcard string `json:"vcard,omitempty"`
}

// Represents the content of an invoice message to be sent as the result of an inline query.
type InputInvoiceMessageContent struct {
	// Product name, 1-32 characters
	Title string `json:"title,omitempty"`

	// Product description, 1-255 characters
	Description string `json:"description,omitempty"`

	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for
	// your internal processes.
	Payload string `json:"payload,omitempty"`

	// Payment provider token, obtained via @BotFather
	ProviderToken string `json:"provider_token,omitempty"`

	// Three-letter ISO 4217 currency code, see more on currencies
	Currency string `json:"currency,omitempty"`

	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount,
	// delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices,omitempty"`

	// Optional. The maximum accepted amount for tips in the smallest units of the currency
	// (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass
	// max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of
	// digits past the decimal point for each currency (2 for the majority of currencies).
	// Defaults to 0
	MaxTipAmount int64 `json:"max_tip_amount,omitempty"`

	// Optional. A JSON-serialized array of suggested amounts of tip in the smallest units of the
	// currency (integer, not float/double). At most 4 suggested tip amounts can be specified.
	// The suggested tip amounts must be positive, passed in a strictly increased order and must
	// not exceed max_tip_amount.
	SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty"`

	// Optional. A JSON-serialized object for data about the invoice, which will be shared with
	// the payment provider. A detailed description of the required fields should be provided by
	// the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// Optional. URL of the product photo for the invoice. Can be a photo of the goods or a
	// marketing image for a service.
	PhotoURL string `json:"photo_url,omitempty"`

	// Optional. Photo size in bytes
	PhotoSize int64 `json:"photo_size,omitempty"`

	// Optional. Photo width
	PhotoWidth int64 `json:"photo_width,omitempty"`

	// Optional. Photo height
	PhotoHeight int64 `json:"photo_height,omitempty"`

	// Optional. Pass True if you require the user's full name to complete the order
	NeedName bool `json:"need_name,omitempty"`

	// Optional. Pass True if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// Optional. Pass True if you require the user's email address to complete the order
	NeedEmail bool `json:"need_email,omitempty"`

	// Optional. Pass True if you require the user's shipping address to complete the order
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// Optional. Pass True if the user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// Optional. Pass True if the user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// Optional. Pass True if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible,omitempty"`
}

// Represents a result of an inline query that was chosen by the user and sent to their chat
// partner.
type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen
	ResultID string `json:"result_id,omitempty"`

	// The user that chose the result
	From *User `json:"from,omitempty"`

	// Optional. Sender location, only for bots that require user location
	Location *Location `json:"location,omitempty"`

	// Optional. Identifier of the sent inline message. Available only if there is an inline
	// keyboard attached to the message. Will be also received in callback queries and can be
	// used to edit the message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// The query that was used to obtain the result
	Query string `json:"query,omitempty"`
}

// Describes an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
	// Optional. Identifier of the sent inline message. Available only if there is an inline
	// keyboard attached to the message.
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// This object represents a portion of the price for goods or services.
type LabeledPrice struct {
	// Portion label
	Label string `json:"label,omitempty"`

	// Price of the product in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in
	// currencies.json, it shows the number of digits past the decimal point for each currency (2
	// for the majority of currencies).
	Amount int64 `json:"amount,omitempty"`
}

// This object contains basic information about an invoice.
type Invoice struct {
	// Product name
	Title string `json:"title,omitempty"`

	// Product description
	Description string `json:"description,omitempty"`

	// Unique bot deep-linking parameter that can be used to generate this invoice
	StartParameter string `json:"start_parameter,omitempty"`

	// Three-letter ISO 4217 currency code
	Currency string `json:"currency,omitempty"`

	// Total price in the smallest units of the currency (integer, not float/double). For
	// example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in
	// currencies.json, it shows the number of digits past the decimal point for each currency (2
	// for the majority of currencies).
	TotalAmount int64 `json:"total_amount,omitempty"`
}

// This object represents a shipping address.
type ShippingAddress struct {
	// Two-letter ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code,omitempty"`

	// State, if applicable
	State string `json:"state,omitempty"`

	// City
	City string `json:"city,omitempty"`

	// First line for the address
	StreetLine1 string `json:"street_line1,omitempty"`

	// Second line for the address
	StreetLine2 string `json:"street_line2,omitempty"`

	// Address post code
	PostCode string `json:"post_code,omitempty"`
}

// This object represents information about an order.
type OrderInfo struct {
	// Optional. User name
	Name string `json:"name,omitempty"`

	// Optional. User's phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// Optional. User email
	Email string `json:"email,omitempty"`

	// Optional. User shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// This object represents one shipping option.
type ShippingOption struct {
	// Shipping option identifier
	ID string `json:"id,omitempty"`

	// Option title
	Title string `json:"title,omitempty"`

	// List of price portions
	Prices []*LabeledPrice `json:"prices,omitempty"`
}

// This object contains basic information about a successful payment.
type SuccessfulPayment struct {
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency,omitempty"`

	// Total price in the smallest units of the currency (integer, not float/double). For
	// example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in
	// currencies.json, it shows the number of digits past the decimal point for each currency (2
	// for the majority of currencies).
	TotalAmount int64 `json:"total_amount,omitempty"`

	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload,omitempty"`

	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`

	// Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id,omitempty"`

	// Provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"`
}

// This object contains information about an incoming shipping query.
type ShippingQuery struct {
	// Unique query identifier
	ID string `json:"id,omitempty"`

	// User who sent the query
	From *User `json:"from,omitempty"`

	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload,omitempty"`

	// User specified shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	// Unique query identifier
	ID string `json:"id,omitempty"`

	// User who sent the query
	From *User `json:"from,omitempty"`

	// Three-letter ISO 4217 currency code
	Currency string `json:"currency,omitempty"`

	// Total price in the smallest units of the currency (integer, not float/double). For
	// example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in
	// currencies.json, it shows the number of digits past the decimal point for each currency (2
	// for the majority of currencies).
	TotalAmount int64 `json:"total_amount,omitempty"`

	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload,omitempty"`

	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

// Telegram Passport is a unified authorization method for services that require personal
// identification. Users can upload their documents once, then instantly share their data with
// services that require real-world ID (finance, ICOs, etc.). Please see the manual for
// details.   Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
	// Array with information about documents and other Telegram Passport elements that was
	// shared with the bot
	Data []*EncryptedPassportElement `json:"data,omitempty"`

	// Encrypted credentials required to decrypt the data
	Credentials *EncryptedCredentials `json:"credentials,omitempty"`
}

// This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport
// files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id,omitempty"`

	// Unique identifier for this file, which is supposed to be the same over time and for
	// different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"`

	// File size in bytes
	FileSize int64 `json:"file_size,omitempty"`

	// Unix time when the file was uploaded
	FileDate int64 `json:"file_date,omitempty"`
}

// Describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	// Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”,
	// “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”,
	// “passport_registration”, “temporary_registration”, “phone_number”, “email”.
	Type string `json:"type,omitempty"`

	// Optional. Base64-encoded encrypted Telegram Passport element data provided by the user,
	// available for “personal_details”, “passport”, “driver_license”, “identity_card”,
	// “internal_passport” and “address” types. Can be decrypted and verified using the
	// accompanying EncryptedCredentials.
	Data string `json:"data,omitempty"`

	// Optional. User's verified phone number, available only for “phone_number” type
	PhoneNumber string `json:"phone_number,omitempty"`

	// Optional. User's verified email address, available only for “email” type
	Email string `json:"email,omitempty"`

	// Optional. Array of encrypted files with documents provided by the user, available for
	// “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and
	// “temporary_registration” types. Files can be decrypted and verified using the accompanying
	// EncryptedCredentials.
	Files []*PassportFile `json:"files,omitempty"`

	// Optional. Encrypted file with the front side of the document, provided by the user.
	// Available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The
	// file can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide *PassportFile `json:"front_side,omitempty"`

	// Optional. Encrypted file with the reverse side of the document, provided by the user.
	// Available for “driver_license” and “identity_card”. The file can be decrypted and verified
	// using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`

	// Optional. Encrypted file with the selfie of the user holding a document, provided by the
	// user; available for “passport”, “driver_license”, “identity_card” and “internal_passport”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie *PassportFile `json:"selfie,omitempty"`

	// Optional. Array of encrypted files with translated versions of documents provided by the
	// user. Available if requested for “passport”, “driver_license”, “identity_card”,
	// “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”,
	// “passport_registration” and “temporary_registration” types. Files can be decrypted and
	// verified using the accompanying EncryptedCredentials.
	Translation []*PassportFile `json:"translation,omitempty"`

	// Base64-encoded element hash for using in PassportElementErrorUnspecified
	Hash string `json:"hash,omitempty"`
}

// Describes data required for decrypting and authenticating EncryptedPassportElement . See the
// Telegram Passport Documentation for a complete description of the data decryption and
// authentication processes.
type EncryptedCredentials struct {
	// Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and
	// secrets required for EncryptedPassportElement decryption and authentication
	Data string `json:"data,omitempty"`

	// Base64-encoded data hash for data authentication
	Hash string `json:"hash,omitempty"`

	// Base64-encoded secret, encrypted with the bot's public RSA key, required for data
	// decryption
	Secret string `json:"secret,omitempty"`
}

// This object represents an error in the Telegram Passport element which was submitted that
// should be resolved by the user. It should be one of:   PassportElementErrorDataField
// PassportElementErrorFrontSide  PassportElementErrorReverseSide  PassportElementErrorSelfie
// PassportElementErrorFile  PassportElementErrorFiles  PassportElementErrorTranslationFile
// PassportElementErrorTranslationFiles  PassportElementErrorUnspecified
type PassportElementError struct {
}

// Represents an issue in one of the data fields that was provided by the user. The error is
// considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	// Error source, must be data
	Source string `json:"source,omitempty"`

	// The section of the user's Telegram Passport which has the error, one of
	// “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”,
	// “address”
	Type string `json:"type,omitempty"`

	// Name of the data field which has the error
	FieldName string `json:"field_name,omitempty"`

	// Base64-encoded data hash
	DataHash string `json:"data_hash,omitempty"`

	// Error message
	Message string `json:"message,omitempty"`
}

// Represents an issue with the front side of a document. The error is considered resolved when
// the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	// Error source, must be front_side
	Source string `json:"source,omitempty"`

	// The section of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type,omitempty"`

	// Base64-encoded hash of the file with the front side of the document
	FileHash string `json:"file_hash,omitempty"`

	// Error message
	Message string `json:"message,omitempty"`
}

// Represents an issue with the reverse side of a document. The error is considered resolved
// when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	// Error source, must be reverse_side
	Source string `json:"source,omitempty"`

	// The section of the user's Telegram Passport which has the issue, one of “driver_license”,
	// “identity_card”
	Type string `json:"type,omitempty"`

	// Base64-encoded hash of the file with the reverse side of the document
	FileHash string `json:"file_hash,omitempty"`

	// Error message
	Message string `json:"message,omitempty"`
}

// Represents an issue with the selfie with a document. The error is considered resolved when
// the file with the selfie changes.
type PassportElementErrorSelfie struct {
	// Error source, must be selfie
	Source string `json:"source,omitempty"`

	// The section of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type,omitempty"`

	// Base64-encoded hash of the file with the selfie
	FileHash string `json:"file_hash,omitempty"`

	// Error message
	Message string `json:"message,omitempty"`
}

// Represents an issue with a document scan. The error is considered resolved when the file
// with the document scan changes.
type PassportElementErrorFile struct {
	// Error source, must be file
	Source string `json:"source,omitempty"`

	// The section of the user's Telegram Passport which has the issue, one of “utility_bill”,
	// “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type,omitempty"`

	// Base64-encoded file hash
	FileHash string `json:"file_hash,omitempty"`

	// Error message
	Message string `json:"message,omitempty"`
}

// Represents an issue with a list of scans. The error is considered resolved when the list of
// files containing the scans changes.
type PassportElementErrorFiles struct {
	// Error source, must be files
	Source string `json:"source,omitempty"`

	// The section of the user's Telegram Passport which has the issue, one of “utility_bill”,
	// “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type,omitempty"`

	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes,omitempty"`

	// Error message
	Message string `json:"message,omitempty"`
}

// Represents an issue with one of the files that constitute the translation of a document. The
// error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	// Error source, must be translation_file
	Source string `json:"source,omitempty"`

	// Type of element of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”,
	// “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type,omitempty"`

	// Base64-encoded file hash
	FileHash string `json:"file_hash,omitempty"`

	// Error message
	Message string `json:"message,omitempty"`
}

// Represents an issue with the translated version of a document. The error is considered
// resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	// Error source, must be translation_files
	Source string `json:"source,omitempty"`

	// Type of element of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”,
	// “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type,omitempty"`

	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes,omitempty"`

	// Error message
	Message string `json:"message,omitempty"`
}

// Represents an issue in an unspecified place. The error is considered resolved when new data
// is added.
type PassportElementErrorUnspecified struct {
	// Error source, must be unspecified
	Source string `json:"source,omitempty"`

	// Type of element of the user's Telegram Passport which has the issue
	Type string `json:"type,omitempty"`

	// Base64-encoded element hash
	ElementHash string `json:"element_hash,omitempty"`

	// Error message
	Message string `json:"message,omitempty"`
}

// This object represents a game. Use BotFather to create and edit games, their short names
// will act as unique identifiers.
type Game struct {
	// Title of the game
	Title string `json:"title,omitempty"`

	// Description of the game
	Description string `json:"description,omitempty"`

	// Photo that will be displayed in the game message in chats.
	Photo []*PhotoSize `json:"photo,omitempty"`

	// Optional. Brief description of the game or high scores included in the game message. Can
	// be automatically edited to include current high scores for the game when the bot calls
	// setGameScore, or manually edited using editMessageText. 0-4096 characters.
	Text string `json:"text,omitempty"`

	// Optional. Special entities that appear in text, such as usernames, URLs, bot commands,
	// etc.
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`

	// Optional. Animation that will be displayed in the game message in chats. Upload via
	// BotFather
	Animation *Animation `json:"animation,omitempty"`
}

// A placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame struct {
}

// This object represents one row of the high scores table for a game.
type GameHighScore struct {
	// Position in high score table for the game
	Position int64 `json:"position,omitempty"`

	// User
	User *User `json:"user,omitempty"`

	// Score
	Score int64 `json:"score,omitempty"`
}

// Oneof type fields are merged into one
// Merged fields of MessageOriginUser, MessageOriginHiddenUser, MessageOriginChat,
// MessageOriginChannel
type MessageOrigin struct {
	// Type of the message origin, always “channel”
	Type string `json:"type,omitempty"`

	// Date the message was sent originally in Unix time
	Date int64 `json:"date,omitempty"`

	// User that sent the message originally
	SenderUser *User `json:"sender_user,omitempty"`

	// Name of the user that sent the message originally
	SenderUserName string `json:"sender_user_name,omitempty"`

	// Chat that sent the message originally
	SenderChat *Chat `json:"sender_chat,omitempty"`

	// Optional. Signature of the original post author
	AuthorSignature string `json:"author_signature,omitempty"`

	// Channel chat to which the message was originally sent
	Chat *Chat `json:"chat,omitempty"`

	// Unique message identifier inside the chat
	MessageID int64 `json:"message_id,omitempty"`
}

// Merged fields of ReactionTypeEmoji, ReactionTypeCustomEmoji
type ReactionType struct {
	// Type of the reaction, always “custom_emoji”
	Type string `json:"type,omitempty"`

	// Reaction emoji. Currently, it can be one of "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
	Emoji string `json:"emoji,omitempty"`

	// Custom emoji identifier
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// Bot request and response types
// Request for API call 'getUpdates'
type GetUpdatesRequest struct {
	// Identifier of the first update to be returned. Must be greater by one than the highest
	// among the identifiers of previously received updates. By default, updates starting with
	// the earliest unconfirmed update are returned. An update is considered confirmed as soon as
	// getUpdates is called with an offset higher than its update_id. The negative offset can be
	// specified to retrieve updates starting from -offset update from the end of the updates
	// queue. All previous updates will be forgotten.
	Offset int64 `json:"offset,omitempty"`

	// Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults
	// to 100.
	Limit int64 `json:"limit,omitempty"`

	// Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be
	// positive, short polling should be used for testing purposes only.
	Timeout int64 `json:"timeout,omitempty"`

	// A JSON-serialized list of the update types you want your bot to receive. For example,
	// specify ["message", "edited_channel_post", "callback_query"] to only receive updates of
	// these types. See Update for a complete list of available update types. Specify an empty
	// list to receive all update types except chat_member, message_reaction, and
	// message_reaction_count (default). If not specified, the previous setting will be
	// used.Please note that this parameter doesn't affect updates created before the call to the
	// getUpdates, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// Response for API call 'getUpdates'
type GetUpdatesResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result []*Update `json:"result,omitempty"`
}

// Request for API call 'setWebhook'
type SetWebhookRequest struct {
	// HTTPS URL to send updates to. Use an empty string to remove webhook integration
	URL string `json:"url,omitempty"`

	// Upload your public key certificate so that the root certificate in use can be checked. See
	// our self-signed guide for details.
	Certificate *InputFile `json:"certificate,omitempty"`

	// The fixed IP address which will be used to send webhook requests instead of the IP address
	// resolved through DNS
	IPAddress string `json:"ip_address,omitempty"`

	// The maximum allowed number of simultaneous HTTPS connections to the webhook for update
	// delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server,
	// and higher values to increase your bot's throughput.
	MaxConnections int64 `json:"max_connections,omitempty"`

	// A JSON-serialized list of the update types you want your bot to receive. For example,
	// specify ["message", "edited_channel_post", "callback_query"] to only receive updates of
	// these types. See Update for a complete list of available update types. Specify an empty
	// list to receive all update types except chat_member, message_reaction, and
	// message_reaction_count (default). If not specified, the previous setting will be
	// used.Please note that this parameter doesn't affect updates created before the call to the
	// setWebhook, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`

	// Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`

	// A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook
	// request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header
	// is useful to ensure that the request comes from a webhook set by you.
	SecretToken string `json:"secret_token,omitempty"`
}

// Response for API call 'setWebhook'
type SetWebhookResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'deleteWebhook'
type DeleteWebhookRequest struct {
	// Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// Response for API call 'deleteWebhook'
type DeleteWebhookResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getWebhookInfo'
type GetWebhookInfoRequest struct {
}

// Response for API call 'getWebhookInfo'
type GetWebhookInfoResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getMe'
type GetMeRequest struct {
}

// Response for API call 'getMe'
type GetMeResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'logOut'
type LogOutRequest struct {
}

// Response for API call 'logOut'
type LogOutResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'close'
type CloseRequest struct {
}

// Response for API call 'close'
type CloseResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'sendMessage'
type SendMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Text of the message to be sent, 1-4096 characters after entities parsing
	Text string `json:"text,omitempty"`

	// Mode for parsing entities in the message text. See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in message text, which can be
	// specified instead of parse_mode
	Entities []*MessageEntity `json:"entities,omitempty"`

	// Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendMessage'
type SendMessageResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'forwardMessage'
type ForwardMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Unique identifier for the chat where the original message was sent (or channel username in
	// the format @channelusername)
	FromChatID string `json:"from_chat_id,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the forwarded message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Message identifier in the chat specified in from_chat_id
	MessageID int64 `json:"message_id,omitempty"`
}

// Response for API call 'forwardMessage'
type ForwardMessageResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'forwardMessages'
type ForwardMessagesRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Unique identifier for the chat where the original messages were sent (or channel username
	// in the format @channelusername)
	FromChatID string `json:"from_chat_id,omitempty"`

	// Identifiers of 1-100 messages in the chat from_chat_id to forward. The identifiers must be
	// specified in a strictly increasing order.
	MessageIds []int64 `json:"message_ids,omitempty"`

	// Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the forwarded messages from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`
}

// Response for API call 'forwardMessages'
type ForwardMessagesResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result []*MessageId `json:"result,omitempty"`
}

// Request for API call 'copyMessage'
type CopyMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Unique identifier for the chat where the original message was sent (or channel username in
	// the format @channelusername)
	FromChatID string `json:"from_chat_id,omitempty"`

	// Message identifier in the chat specified in from_chat_id
	MessageID int64 `json:"message_id,omitempty"`

	// New caption for media, 0-1024 characters after entities parsing. If not specified, the
	// original caption is kept
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the new caption. See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in the new caption, which can be
	// specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'copyMessage'
type CopyMessageResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'copyMessages'
type CopyMessagesRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Unique identifier for the chat where the original messages were sent (or channel username
	// in the format @channelusername)
	FromChatID string `json:"from_chat_id,omitempty"`

	// Identifiers of 1-100 messages in the chat from_chat_id to copy. The identifiers must be
	// specified in a strictly increasing order.
	MessageIds []int64 `json:"message_ids,omitempty"`

	// Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent messages from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Pass True to copy the messages without their captions
	RemoveCaption bool `json:"remove_caption,omitempty"`
}

// Response for API call 'copyMessages'
type CopyMessagesResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result []*MessageId `json:"result,omitempty"`
}

// Request for API call 'sendPhoto'
type SendPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Photo to send. Pass a file_id as String to send a photo that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the
	// Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB
	// in size. The photo's width and height must not exceed 10000 in total. Width and height
	// ratio must be at most 20. More information on Sending Files »
	Photo interface{} `json:"photo,omitempty"`

	// Photo caption (may also be used when resending photos by file_id), 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the photo caption. See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendPhoto'
type SendPhotoResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendAudio'
type SendAudioRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Audio file to send. Pass a file_id as String to send an audio file that exists on the
	// Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio
	// file from the Internet, or upload a new one using multipart/form-data. More information on
	// Sending Files »
	Audio interface{} `json:"audio,omitempty"`

	// Audio caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the audio caption. See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Duration of the audio in seconds
	Duration int64 `json:"duration,omitempty"`

	// Performer
	Performer string `json:"performer,omitempty"`

	// Track name
	Title string `json:"title,omitempty"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in
	// size. A thumbnail's width and height should not exceed 320. Ignored if the file is not
	// uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as
	// a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded
	// using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendAudio'
type SendAudioResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendDocument'
type SendDocumentRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// File to send. Pass a file_id as String to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet,
	// or upload a new one using multipart/form-data. More information on Sending Files »
	Document interface{} `json:"document,omitempty"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in
	// size. A thumbnail's width and height should not exceed 320. Ignored if the file is not
	// uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as
	// a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded
	// using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail,omitempty"`

	// Document caption (may also be used when resending documents by file_id), 0-1024 characters
	// after entities parsing
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the document caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Disables automatic server-side content type detection for files uploaded using
	// multipart/form-data
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendDocument'
type SendDocumentResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendVideo'
type SendVideoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Video to send. Pass a file_id as String to send a video that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the
	// Internet, or upload a new video using multipart/form-data. More information on Sending
	// Files »
	Video interface{} `json:"video,omitempty"`

	// Duration of sent video in seconds
	Duration int64 `json:"duration,omitempty"`

	// Video width
	Width int64 `json:"width,omitempty"`

	// Video height
	Height int64 `json:"height,omitempty"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in
	// size. A thumbnail's width and height should not exceed 320. Ignored if the file is not
	// uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as
	// a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded
	// using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail,omitempty"`

	// Video caption (may also be used when resending videos by file_id), 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the video caption. See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`

	// Pass True if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendVideo'
type SendVideoResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendAnimation'
type SendAnimationRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Animation to send. Pass a file_id as String to send an animation that exists on the
	// Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an
	// animation from the Internet, or upload a new animation using multipart/form-data. More
	// information on Sending Files »
	Animation interface{} `json:"animation,omitempty"`

	// Duration of sent animation in seconds
	Duration int64 `json:"duration,omitempty"`

	// Animation width
	Width int64 `json:"width,omitempty"`

	// Animation height
	Height int64 `json:"height,omitempty"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in
	// size. A thumbnail's width and height should not exceed 320. Ignored if the file is not
	// uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as
	// a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded
	// using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail,omitempty"`

	// Animation caption (may also be used when resending animation by file_id), 0-1024
	// characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the animation caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendAnimation'
type SendAnimationResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendVoice'
type SendVoiceRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Audio file to send. Pass a file_id as String to send a file that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the
	// Internet, or upload a new one using multipart/form-data. More information on Sending Files
	// »
	Voice interface{} `json:"voice,omitempty"`

	// Voice message caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the voice message caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Duration of the voice message in seconds
	Duration int64 `json:"duration,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendVoice'
type SendVoiceResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendVideoNote'
type SendVideoNoteRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Video note to send. Pass a file_id as String to send a video note that exists on the
	// Telegram servers (recommended) or upload a new video using multipart/form-data. More
	// information on Sending Files ». Sending video notes by a URL is currently unsupported
	VideoNote interface{} `json:"video_note,omitempty"`

	// Duration of sent video in seconds
	Duration int64 `json:"duration,omitempty"`

	// Video width and height, i.e. diameter of the video message
	Length int64 `json:"length,omitempty"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in
	// size. A thumbnail's width and height should not exceed 320. Ignored if the file is not
	// uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as
	// a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded
	// using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendVideoNote'
type SendVideoNoteResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendLocation'
type SendLocationRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Latitude of the location
	Latitude float32 `json:"latitude,omitempty"`

	// Longitude of the location
	Longitude float32 `json:"longitude,omitempty"`

	// The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`

	// Period in seconds for which the location will be updated (see Live Locations, should be
	// between 60 and 86400.
	LivePeriod int64 `json:"live_period,omitempty"`

	// For live locations, a direction in which the user is moving, in degrees. Must be between 1
	// and 360 if specified.
	Heading int64 `json:"heading,omitempty"`

	// For live locations, a maximum distance for proximity alerts about approaching another chat
	// member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendLocation'
type SendLocationResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendVenue'
type SendVenueRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Latitude of the venue
	Latitude float32 `json:"latitude,omitempty"`

	// Longitude of the venue
	Longitude float32 `json:"longitude,omitempty"`

	// Name of the venue
	Title string `json:"title,omitempty"`

	// Address of the venue
	Address string `json:"address,omitempty"`

	// Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id,omitempty"`

	// Foursquare type of the venue, if known. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendVenue'
type SendVenueResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendContact'
type SendContactRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Contact's phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// Contact's first name
	FirstName string `json:"first_name,omitempty"`

	// Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Additional data about the contact in the form of a vCard, 0-2048 bytes
	Vcard string `json:"vcard,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendContact'
type SendContactResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendPoll'
type SendPollRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Poll question, 1-300 characters
	Question string `json:"question,omitempty"`

	// A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
	Options []string `json:"options,omitempty"`

	// True, if the poll needs to be anonymous, defaults to True
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// Poll type, “quiz” or “regular”, defaults to “regular”
	Type string `json:"type,omitempty"`

	// True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to
	// False
	AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`

	// 0-based identifier of the correct answer option, required for polls in quiz mode
	CorrectOptionID int64 `json:"correct_option_id,omitempty"`

	// Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a
	// quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	Explanation string `json:"explanation,omitempty"`

	// Mode for parsing entities in the explanation. See formatting options for more details.
	ExplanationParseMode string `json:"explanation_parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in the poll explanation, which can
	// be specified instead of parse_mode
	ExplanationEntities []*MessageEntity `json:"explanation_entities,omitempty"`

	// Amount of time in seconds the poll will be active after creation, 5-600. Can't be used
	// together with close_date.
	OpenPeriod int64 `json:"open_period,omitempty"`

	// Point in time (Unix timestamp) when the poll will be automatically closed. Must be at
	// least 5 and no more than 600 seconds in the future. Can't be used together with
	// open_period.
	CloseDate int64 `json:"close_date,omitempty"`

	// Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
	IsClosed bool `json:"is_closed,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendPoll'
type SendPollResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendDice'
type SendDiceRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Emoji on which the dice throw animation is based. Currently, must be one of “”, “”, “”,
	// “”, “”, or “”. Dice can have values 1-6 for “”, “” and “”, values 1-5 for “” and “”, and
	// values 1-64 for “”. Defaults to “”
	Emoji string `json:"emoji,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendDice'
type SendDiceResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'sendChatAction'
type SendChatActionRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread; supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Type of action to broadcast. Choose one, depending on what the user is about to receive:
	// typing for text messages, upload_photo for photos, record_video or upload_video for
	// videos, record_voice or upload_voice for voice notes, upload_document for general files,
	// choose_sticker for stickers, find_location for location data, record_video_note or
	// upload_video_note for video notes.
	Action string `json:"action,omitempty"`
}

// Response for API call 'sendChatAction'
type SendChatActionResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setMessageReaction'
type SetMessageReactionRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Identifier of the target message. If the message belongs to a media group, the reaction is
	// set to the first non-deleted message in the group instead.
	MessageID int64 `json:"message_id,omitempty"`

	// New list of reaction types to set on the message. Currently, as non-premium users, bots
	// can set up to one reaction per message. A custom emoji reaction can be used if it is
	// either already present on the message or explicitly allowed by chat administrators.
	Reaction []*ReactionType `json:"reaction,omitempty"`

	// Pass True to set the reaction with a big animation
	IsBig bool `json:"is_big,omitempty"`
}

// Response for API call 'setMessageReaction'
type SetMessageReactionResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getUserProfilePhotos'
type GetUserProfilePhotosRequest struct {
	// Unique identifier of the target user
	UserID int64 `json:"user_id,omitempty"`

	// Sequential number of the first photo to be returned. By default, all photos are returned.
	Offset int64 `json:"offset,omitempty"`

	// Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults
	// to 100.
	Limit int64 `json:"limit,omitempty"`
}

// Response for API call 'getUserProfilePhotos'
type GetUserProfilePhotosResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getFile'
type GetFileRequest struct {
	// File identifier to get information about
	FileID string `json:"file_id,omitempty"`
}

// Response for API call 'getFile'
type GetFileResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *File `json:"result,omitempty"`
}

// Request for API call 'banChatMember'
type BanChatMemberRequest struct {
	// Unique identifier for the target group or username of the target supergroup or channel (in
	// the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id,omitempty"`

	// Date when the user will be unbanned; Unix time. If user is banned for more than 366 days
	// or less than 30 seconds from the current time they are considered to be banned forever.
	// Applied for supergroups and channels only.
	UntilDate int64 `json:"until_date,omitempty"`

	// Pass True to delete all messages from the chat for the user that is being removed. If
	// False, the user will be able to see messages in the group that were sent before the user
	// was removed. Always True for supergroups and channels.
	RevokeMessages bool `json:"revoke_messages,omitempty"`
}

// Response for API call 'banChatMember'
type BanChatMemberResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'unbanChatMember'
type UnbanChatMemberRequest struct {
	// Unique identifier for the target group or username of the target supergroup or channel (in
	// the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id,omitempty"`

	// Do nothing if the user is not banned
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}

// Response for API call 'unbanChatMember'
type UnbanChatMemberResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'restrictChatMember'
type RestrictChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id,omitempty"`

	// A JSON-serialized object for new user permissions
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// Pass True if chat permissions are set independently. Otherwise, the
	// can_send_other_messages and can_add_web_page_previews permissions will imply the
	// can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos,
	// can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission
	// will imply the can_send_messages permission.
	UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"`

	// Date when restrictions will be lifted for the user; Unix time. If user is restricted for
	// more than 366 days or less than 30 seconds from the current time, they are considered to
	// be restricted forever
	UntilDate int64 `json:"until_date,omitempty"`
}

// Response for API call 'restrictChatMember'
type RestrictChatMemberResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'promoteChatMember'
type PromoteChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id,omitempty"`

	// Pass True if the administrator's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// Pass True if the administrator can access the chat event log, boost list in channels, see
	// channel members, report spam messages, see anonymous administrators in supergroups and
	// ignore slow mode. Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat,omitempty"`

	// Pass True if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

	// Pass True if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`

	// Pass True if the administrator can restrict, ban or unban chat members, or access
	// supergroup statistics
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

	// Pass True if the administrator can add new administrators with a subset of their own
	// privileges or demote administrators that they have promoted, directly or indirectly
	// (promoted by administrators that were appointed by him)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`

	// Pass True if the administrator can change chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// Pass True if the administrator can invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// Pass True if the administrator can post messages in the channel, or access channel
	// statistics; channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// Pass True if the administrator can edit messages of other users and can pin messages;
	// channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// Pass True if the administrator can pin messages, supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// Pass True if the administrator can post stories in the channel; channels only
	CanPostStories bool `json:"can_post_stories,omitempty"`

	// Pass True if the administrator can edit stories posted by other users; channels only
	CanEditStories bool `json:"can_edit_stories,omitempty"`

	// Pass True if the administrator can delete stories posted by other users; channels only
	CanDeleteStories bool `json:"can_delete_stories,omitempty"`

	// Pass True if the user is allowed to create, rename, close, and reopen forum topics,
	// supergroups only
	CanManageTopics bool `json:"can_manage_topics,omitempty"`
}

// Response for API call 'promoteChatMember'
type PromoteChatMemberResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setChatAdministratorCustomTitle'
type SetChatAdministratorCustomTitleRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id,omitempty"`

	// New custom title for the administrator; 0-16 characters, emoji are not allowed
	CustomTitle string `json:"custom_title,omitempty"`
}

// Response for API call 'setChatAdministratorCustomTitle'
type SetChatAdministratorCustomTitleResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'banChatSenderChat'
type BanChatSenderChatRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target sender chat
	SenderChatID int64 `json:"sender_chat_id,omitempty"`
}

// Response for API call 'banChatSenderChat'
type BanChatSenderChatResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'unbanChatSenderChat'
type UnbanChatSenderChatRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target sender chat
	SenderChatID int64 `json:"sender_chat_id,omitempty"`
}

// Response for API call 'unbanChatSenderChat'
type UnbanChatSenderChatResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setChatPermissions'
type SetChatPermissionsRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// A JSON-serialized object for new default chat permissions
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// Pass True if chat permissions are set independently. Otherwise, the
	// can_send_other_messages and can_add_web_page_previews permissions will imply the
	// can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos,
	// can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission
	// will imply the can_send_messages permission.
	UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"`
}

// Response for API call 'setChatPermissions'
type SetChatPermissionsResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'exportChatInviteLink'
type ExportChatInviteLinkRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'exportChatInviteLink'
type ExportChatInviteLinkResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'createChatInviteLink'
type CreateChatInviteLinkRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Invite link name; 0-32 characters
	Name string `json:"name,omitempty"`

	// Point in time (Unix timestamp) when the link will expire
	ExpireDate int64 `json:"expire_date,omitempty"`

	// The maximum number of users that can be members of the chat simultaneously after joining
	// the chat via this invite link; 1-99999
	MemberLimit int64 `json:"member_limit,omitempty"`

	// True, if users joining the chat via the link need to be approved by chat administrators.
	// If True, member_limit can't be specified
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
}

// Response for API call 'createChatInviteLink'
type CreateChatInviteLinkResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'editChatInviteLink'
type EditChatInviteLinkRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// The invite link to edit
	InviteLink string `json:"invite_link,omitempty"`

	// Invite link name; 0-32 characters
	Name string `json:"name,omitempty"`

	// Point in time (Unix timestamp) when the link will expire
	ExpireDate int64 `json:"expire_date,omitempty"`

	// The maximum number of users that can be members of the chat simultaneously after joining
	// the chat via this invite link; 1-99999
	MemberLimit int64 `json:"member_limit,omitempty"`

	// True, if users joining the chat via the link need to be approved by chat administrators.
	// If True, member_limit can't be specified
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
}

// Response for API call 'editChatInviteLink'
type EditChatInviteLinkResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'revokeChatInviteLink'
type RevokeChatInviteLinkRequest struct {
	// Unique identifier of the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// The invite link to revoke
	InviteLink string `json:"invite_link,omitempty"`
}

// Response for API call 'revokeChatInviteLink'
type RevokeChatInviteLinkResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'approveChatJoinRequest'
type ApproveChatJoinRequestRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id,omitempty"`
}

// Response for API call 'approveChatJoinRequest'
type ApproveChatJoinRequestResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'declineChatJoinRequest'
type DeclineChatJoinRequestRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id,omitempty"`
}

// Response for API call 'declineChatJoinRequest'
type DeclineChatJoinRequestResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setChatPhoto'
type SetChatPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// New chat photo, uploaded using multipart/form-data
	Photo *InputFile `json:"photo,omitempty"`
}

// Response for API call 'setChatPhoto'
type SetChatPhotoResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'deleteChatPhoto'
type DeleteChatPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'deleteChatPhoto'
type DeleteChatPhotoResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setChatTitle'
type SetChatTitleRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// New chat title, 1-128 characters
	Title string `json:"title,omitempty"`
}

// Response for API call 'setChatTitle'
type SetChatTitleResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setChatDescription'
type SetChatDescriptionRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// New chat description, 0-255 characters
	Description string `json:"description,omitempty"`
}

// Response for API call 'setChatDescription'
type SetChatDescriptionResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'pinChatMessage'
type PinChatMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Identifier of a message to pin
	MessageID int64 `json:"message_id,omitempty"`

	// Pass True if it is not necessary to send a notification to all chat members about the new
	// pinned message. Notifications are always disabled in channels and private chats.
	DisableNotification bool `json:"disable_notification,omitempty"`
}

// Response for API call 'pinChatMessage'
type PinChatMessageResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'unpinChatMessage'
type UnpinChatMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Identifier of a message to unpin. If not specified, the most recent pinned message (by
	// sending date) will be unpinned.
	MessageID int64 `json:"message_id,omitempty"`
}

// Response for API call 'unpinChatMessage'
type UnpinChatMessageResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'unpinAllChatMessages'
type UnpinAllChatMessagesRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'unpinAllChatMessages'
type UnpinAllChatMessagesResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'leaveChat'
type LeaveChatRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in
	// the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'leaveChat'
type LeaveChatResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getChat'
type GetChatRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in
	// the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'getChat'
type GetChatResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getChatAdministrators'
type GetChatAdministratorsRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in
	// the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'getChatAdministrators'
type GetChatAdministratorsResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result []*ChatMember `json:"result,omitempty"`
}

// Request for API call 'getChatMemberCount'
type GetChatMemberCountRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in
	// the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'getChatMemberCount'
type GetChatMemberCountResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getChatMember'
type GetChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in
	// the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id,omitempty"`
}

// Response for API call 'getChatMember'
type GetChatMemberResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setChatStickerSet'
type SetChatStickerSetRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// Name of the sticker set to be set as the group sticker set
	StickerSetName string `json:"sticker_set_name,omitempty"`
}

// Response for API call 'setChatStickerSet'
type SetChatStickerSetResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'deleteChatStickerSet'
type DeleteChatStickerSetRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'deleteChatStickerSet'
type DeleteChatStickerSetResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getForumTopicIconStickers'
type GetForumTopicIconStickersRequest struct {
}

// Response for API call 'getForumTopicIconStickers'
type GetForumTopicIconStickersResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result []*Sticker `json:"result,omitempty"`
}

// Request for API call 'createForumTopic'
type CreateForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// Topic name, 1-128 characters
	Name string `json:"name,omitempty"`

	// Color of the topic icon in RGB format. Currently, must be one of 7322096 (0x6FB9F0),
	// 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or
	// 16478047 (0xFB6F5F)
	IconColor int64 `json:"icon_color,omitempty"`

	// Unique identifier of the custom emoji shown as the topic icon. Use
	// getForumTopicIconStickers to get all allowed custom emoji identifiers.
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// Response for API call 'createForumTopic'
type CreateForumTopicResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'editForumTopic'
type EditForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread of the forum topic
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// New topic name, 0-128 characters. If not specified or empty, the current name of the topic
	// will be kept
	Name string `json:"name,omitempty"`

	// New unique identifier of the custom emoji shown as the topic icon. Use
	// getForumTopicIconStickers to get all allowed custom emoji identifiers. Pass an empty
	// string to remove the icon. If not specified, the current icon will be kept
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// Response for API call 'editForumTopic'
type EditForumTopicResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'closeForumTopic'
type CloseForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread of the forum topic
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
}

// Response for API call 'closeForumTopic'
type CloseForumTopicResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'reopenForumTopic'
type ReopenForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread of the forum topic
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
}

// Response for API call 'reopenForumTopic'
type ReopenForumTopicResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'deleteForumTopic'
type DeleteForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread of the forum topic
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
}

// Response for API call 'deleteForumTopic'
type DeleteForumTopicResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'unpinAllForumTopicMessages'
type UnpinAllForumTopicMessagesRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread of the forum topic
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
}

// Response for API call 'unpinAllForumTopicMessages'
type UnpinAllForumTopicMessagesResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'editGeneralForumTopic'
type EditGeneralForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`

	// New topic name, 1-128 characters
	Name string `json:"name,omitempty"`
}

// Response for API call 'editGeneralForumTopic'
type EditGeneralForumTopicResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'closeGeneralForumTopic'
type CloseGeneralForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'closeGeneralForumTopic'
type CloseGeneralForumTopicResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'reopenGeneralForumTopic'
type ReopenGeneralForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'reopenGeneralForumTopic'
type ReopenGeneralForumTopicResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'hideGeneralForumTopic'
type HideGeneralForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'hideGeneralForumTopic'
type HideGeneralForumTopicResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'unhideGeneralForumTopic'
type UnhideGeneralForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'unhideGeneralForumTopic'
type UnhideGeneralForumTopicResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'unpinAllGeneralForumTopicMessages'
type UnpinAllGeneralForumTopicMessagesRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID string `json:"chat_id,omitempty"`
}

// Response for API call 'unpinAllGeneralForumTopicMessages'
type UnpinAllGeneralForumTopicMessagesResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'answerCallbackQuery'
type AnswerCallbackQueryRequest struct {
	// Unique identifier for the query to be answered
	CallbackQueryID string `json:"callback_query_id,omitempty"`

	// Text of the notification. If not specified, nothing will be shown to the user, 0-200
	// characters
	Text string `json:"text,omitempty"`

	// If True, an alert will be shown by the client instead of a notification at the top of the
	// chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`

	// URL that will be opened by the user's client. If you have created a Game and accepted the
	// conditions via @BotFather, specify the URL that opens your game - note that this will only
	// work if the query comes from a callback_game button.Otherwise, you may use links like
	// t.me/your_bot?start=XXXX that open your bot with a parameter.
	URL string `json:"url,omitempty"`

	// The maximum amount of time in seconds that the result of the callback query may be cached
	// client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime int64 `json:"cache_time,omitempty"`
}

// Response for API call 'answerCallbackQuery'
type AnswerCallbackQueryResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getUserChatBoosts'
type GetUserChatBoostsRequest struct {
	// Unique identifier for the chat or username of the channel (in the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id,omitempty"`
}

// Response for API call 'getUserChatBoosts'
type GetUserChatBoostsResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setMyCommands'
type SetMyCommandsRequest struct {
	// A JSON-serialized list of bot commands to be set as the list of the bot's commands. At
	// most 100 commands can be specified.
	Commands []*BotCommand `json:"commands,omitempty"`

	// A JSON-serialized object, describing scope of users for which the commands are relevant.
	// Defaults to BotCommandScopeDefault.
	Scope *BotCommandScope `json:"scope,omitempty"`

	// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from
	// the given scope, for whose language there are no dedicated commands
	LanguageCode string `json:"language_code,omitempty"`
}

// Response for API call 'setMyCommands'
type SetMyCommandsResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'deleteMyCommands'
type DeleteMyCommandsRequest struct {
	// A JSON-serialized object, describing scope of users for which the commands are relevant.
	// Defaults to BotCommandScopeDefault.
	Scope *BotCommandScope `json:"scope,omitempty"`

	// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from
	// the given scope, for whose language there are no dedicated commands
	LanguageCode string `json:"language_code,omitempty"`
}

// Response for API call 'deleteMyCommands'
type DeleteMyCommandsResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getMyCommands'
type GetMyCommandsRequest struct {
	// A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	Scope *BotCommandScope `json:"scope,omitempty"`

	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// Response for API call 'getMyCommands'
type GetMyCommandsResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result []*BotCommand `json:"result,omitempty"`
}

// Request for API call 'setMyName'
type SetMyNameRequest struct {
	// New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the
	// given language.
	Name string `json:"name,omitempty"`

	// A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for
	// whose language there is no dedicated name.
	LanguageCode string `json:"language_code,omitempty"`
}

// Response for API call 'setMyName'
type SetMyNameResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getMyName'
type GetMyNameRequest struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// Response for API call 'getMyName'
type GetMyNameResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setMyDescription'
type SetMyDescriptionRequest struct {
	// New bot description; 0-512 characters. Pass an empty string to remove the dedicated
	// description for the given language.
	Description string `json:"description,omitempty"`

	// A two-letter ISO 639-1 language code. If empty, the description will be applied to all
	// users for whose language there is no dedicated description.
	LanguageCode string `json:"language_code,omitempty"`
}

// Response for API call 'setMyDescription'
type SetMyDescriptionResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getMyDescription'
type GetMyDescriptionRequest struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// Response for API call 'getMyDescription'
type GetMyDescriptionResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setMyShortDescription'
type SetMyShortDescriptionRequest struct {
	// New short description for the bot; 0-120 characters. Pass an empty string to remove the
	// dedicated short description for the given language.
	ShortDescription string `json:"short_description,omitempty"`

	// A two-letter ISO 639-1 language code. If empty, the short description will be applied to
	// all users for whose language there is no dedicated short description.
	LanguageCode string `json:"language_code,omitempty"`
}

// Response for API call 'setMyShortDescription'
type SetMyShortDescriptionResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getMyShortDescription'
type GetMyShortDescriptionRequest struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// Response for API call 'getMyShortDescription'
type GetMyShortDescriptionResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setChatMenuButton'
type SetChatMenuButtonRequest struct {
	// Unique identifier for the target private chat. If not specified, default bot's menu button
	// will be changed
	ChatID int64 `json:"chat_id,omitempty"`

	// A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
	MenuButton *MenuButton `json:"menu_button,omitempty"`
}

// Response for API call 'setChatMenuButton'
type SetChatMenuButtonResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getChatMenuButton'
type GetChatMenuButtonRequest struct {
	// Unique identifier for the target private chat. If not specified, default bot's menu button
	// will be returned
	ChatID int64 `json:"chat_id,omitempty"`
}

// Response for API call 'getChatMenuButton'
type GetChatMenuButtonResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setMyDefaultAdministratorRights'
type SetMyDefaultAdministratorRightsRequest struct {
	// A JSON-serialized object describing new default administrator rights. If not specified,
	// the default administrator rights will be cleared.
	Rights *ChatAdministratorRights `json:"rights,omitempty"`

	// Pass True to change the default administrator rights of the bot in channels. Otherwise,
	// the default administrator rights of the bot for groups and supergroups will be changed.
	ForChannels bool `json:"for_channels,omitempty"`
}

// Response for API call 'setMyDefaultAdministratorRights'
type SetMyDefaultAdministratorRightsResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getMyDefaultAdministratorRights'
type GetMyDefaultAdministratorRightsRequest struct {
	// Pass True to get default administrator rights of the bot in channels. Otherwise, default
	// administrator rights of the bot for groups and supergroups will be returned.
	ForChannels bool `json:"for_channels,omitempty"`
}

// Response for API call 'getMyDefaultAdministratorRights'
type GetMyDefaultAdministratorRightsResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'editMessageText'
type EditMessageTextRequest struct {
	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text,omitempty"`

	// Mode for parsing entities in the message text. See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in message text, which can be
	// specified instead of parse_mode
	Entities []*MessageEntity `json:"entities,omitempty"`

	// Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Response for API call 'editMessageText'
type EditMessageTextResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'editMessageCaption'
type EditMessageCaptionRequest struct {
	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// New caption of the message, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the message caption. See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// A JSON-serialized list of special entities that appear in the caption, which can be
	// specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Response for API call 'editMessageCaption'
type EditMessageCaptionResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'editMessageMedia'
type EditMessageMediaRequest struct {
	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// A JSON-serialized object for a new media content of the message
	Media *InputMedia `json:"media,omitempty"`

	// A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Response for API call 'editMessageMedia'
type EditMessageMediaResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'editMessageLiveLocation'
type EditMessageLiveLocationRequest struct {
	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Latitude of new location
	Latitude float32 `json:"latitude,omitempty"`

	// Longitude of new location
	Longitude float32 `json:"longitude,omitempty"`

	// The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`

	// Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading int64 `json:"heading,omitempty"`

	// The maximum distance for proximity alerts about approaching another chat member, in
	// meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`

	// A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Response for API call 'editMessageLiveLocation'
type EditMessageLiveLocationResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'stopMessageLiveLocation'
type StopMessageLiveLocationRequest struct {
	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the message with live
	// location to stop
	MessageID int64 `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Response for API call 'stopMessageLiveLocation'
type StopMessageLiveLocationResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'editMessageReplyMarkup'
type EditMessageReplyMarkupRequest struct {
	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Response for API call 'editMessageReplyMarkup'
type EditMessageReplyMarkupResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'stopPoll'
type StopPollRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Identifier of the original message with the poll
	MessageID int64 `json:"message_id,omitempty"`

	// A JSON-serialized object for a new message inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Response for API call 'stopPoll'
type StopPollResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Poll `json:"result,omitempty"`
}

// Request for API call 'deleteMessage'
type DeleteMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Identifier of the message to delete
	MessageID int64 `json:"message_id,omitempty"`
}

// Response for API call 'deleteMessage'
type DeleteMessageResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'deleteMessages'
type DeleteMessagesRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Identifiers of 1-100 messages to delete. See deleteMessage for limitations on which
	// messages can be deleted
	MessageIds []int64 `json:"message_ids,omitempty"`
}

// Response for API call 'deleteMessages'
type DeleteMessagesResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'sendSticker'
type SendStickerRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Sticker to send. Pass a file_id as String to send a file that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker
	// from the Internet, or upload a new .WEBP or .TGS sticker using multipart/form-data. More
	// information on Sending Files ». Video stickers can only be sent by a file_id. Animated
	// stickers can't be sent via an HTTP URL.
	Sticker interface{} `json:"sticker,omitempty"`

	// Emoji associated with the sticker; only for just uploaded stickers
	Emoji string `json:"emoji,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom
	// reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// Response for API call 'sendSticker'
type SendStickerResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'getStickerSet'
type GetStickerSetRequest struct {
	// Name of the sticker set
	Name string `json:"name,omitempty"`
}

// Response for API call 'getStickerSet'
type GetStickerSetResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *StickerSet `json:"result,omitempty"`
}

// Request for API call 'getCustomEmojiStickers'
type GetCustomEmojiStickersRequest struct {
	// List of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
	CustomEmojiIds []string `json:"custom_emoji_ids,omitempty"`
}

// Response for API call 'getCustomEmojiStickers'
type GetCustomEmojiStickersResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result []*Sticker `json:"result,omitempty"`
}

// Request for API call 'uploadStickerFile'
type UploadStickerFileRequest struct {
	// User identifier of sticker file owner
	UserID int64 `json:"user_id,omitempty"`

	// A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format. See
	// https://core.telegram.org/stickers for technical requirements. More information on Sending
	// Files »
	Sticker *InputFile `json:"sticker,omitempty"`

	// Format of the sticker, must be one of “static”, “animated”, “video”
	StickerFormat string `json:"sticker_format,omitempty"`
}

// Response for API call 'uploadStickerFile'
type UploadStickerFileResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'createNewStickerSet'
type CreateNewStickerSetRequest struct {
	// User identifier of created sticker set owner
	UserID int64 `json:"user_id,omitempty"`

	// Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can
	// contain only English letters, digits and underscores. Must begin with a letter, can't
	// contain consecutive underscores and must end in "_by_<bot_username>". <bot_username> is
	// case insensitive. 1-64 characters.
	Name string `json:"name,omitempty"`

	// Sticker set title, 1-64 characters
	Title string `json:"title,omitempty"`

	// A JSON-serialized list of 1-50 initial stickers to be added to the sticker set
	Stickers []*InputSticker `json:"stickers,omitempty"`

	// Format of stickers in the set, must be one of “static”, “animated”, “video”
	StickerFormat string `json:"sticker_format,omitempty"`

	// Type of stickers in the set, pass “regular”, “mask”, or “custom_emoji”. By default, a
	// regular sticker set is created.
	StickerType string `json:"sticker_type,omitempty"`

	// Pass True if stickers in the sticker set must be repainted to the color of text when used
	// in messages, the accent color if used as emoji status, white on chat photos, or another
	// appropriate color based on context; for custom emoji sticker sets only
	NeedsRepainting bool `json:"needs_repainting,omitempty"`
}

// Response for API call 'createNewStickerSet'
type CreateNewStickerSetResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'addStickerToSet'
type AddStickerToSetRequest struct {
	// User identifier of sticker set owner
	UserID int64 `json:"user_id,omitempty"`

	// Sticker set name
	Name string `json:"name,omitempty"`

	// A JSON-serialized object with information about the added sticker. If exactly the same
	// sticker had already been added to the set, then the set isn't changed.
	Sticker *InputSticker `json:"sticker,omitempty"`
}

// Response for API call 'addStickerToSet'
type AddStickerToSetResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setStickerPositionInSet'
type SetStickerPositionInSetRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker,omitempty"`

	// New sticker position in the set, zero-based
	Position int64 `json:"position,omitempty"`
}

// Response for API call 'setStickerPositionInSet'
type SetStickerPositionInSetResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'deleteStickerFromSet'
type DeleteStickerFromSetRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker,omitempty"`
}

// Response for API call 'deleteStickerFromSet'
type DeleteStickerFromSetResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setStickerEmojiList'
type SetStickerEmojiListRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker,omitempty"`

	// A JSON-serialized list of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list,omitempty"`
}

// Response for API call 'setStickerEmojiList'
type SetStickerEmojiListResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setStickerKeywords'
type SetStickerKeywordsRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker,omitempty"`

	// A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to
	// 64 characters
	Keywords []string `json:"keywords,omitempty"`
}

// Response for API call 'setStickerKeywords'
type SetStickerKeywordsResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setStickerMaskPosition'
type SetStickerMaskPositionRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker,omitempty"`

	// A JSON-serialized object with the position where the mask should be placed on faces. Omit
	// the parameter to remove the mask position.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// Response for API call 'setStickerMaskPosition'
type SetStickerMaskPositionResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setStickerSetTitle'
type SetStickerSetTitleRequest struct {
	// Sticker set name
	Name string `json:"name,omitempty"`

	// Sticker set title, 1-64 characters
	Title string `json:"title,omitempty"`
}

// Response for API call 'setStickerSetTitle'
type SetStickerSetTitleResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setStickerSetThumbnail'
type SetStickerSetThumbnailRequest struct {
	// Sticker set name
	Name string `json:"name,omitempty"`

	// User identifier of the sticker set owner
	UserID int64 `json:"user_id,omitempty"`

	// A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a
	// width and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes
	// in size (see https://core.telegram.org/stickers#animated-sticker-requirements for animated
	// sticker technical requirements), or a WEBM video with the thumbnail up to 32 kilobytes in
	// size; see https://core.telegram.org/stickers#video-sticker-requirements for video sticker
	// technical requirements. Pass a file_id as a String to send a file that already exists on
	// the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the
	// Internet, or upload a new one using multipart/form-data. More information on Sending Files
	// ». Animated and video sticker set thumbnails can't be uploaded via HTTP URL. If omitted,
	// then the thumbnail is dropped and the first sticker is used as the thumbnail.
	Thumbnail interface{} `json:"thumbnail,omitempty"`
}

// Response for API call 'setStickerSetThumbnail'
type SetStickerSetThumbnailResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setCustomEmojiStickerSetThumbnail'
type SetCustomEmojiStickerSetThumbnailRequest struct {
	// Sticker set name
	Name string `json:"name,omitempty"`

	// Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop
	// the thumbnail and use the first sticker as the thumbnail.
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// Response for API call 'setCustomEmojiStickerSetThumbnail'
type SetCustomEmojiStickerSetThumbnailResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'deleteStickerSet'
type DeleteStickerSetRequest struct {
	// Sticker set name
	Name string `json:"name,omitempty"`
}

// Response for API call 'deleteStickerSet'
type DeleteStickerSetResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'answerInlineQuery'
type AnswerInlineQueryRequest struct {
	// Unique identifier for the answered query
	InlineQueryID string `json:"inline_query_id,omitempty"`

	// A JSON-serialized array of results for the inline query
	Results []*InlineQueryResult `json:"results,omitempty"`

	// The maximum amount of time in seconds that the result of the inline query may be cached on
	// the server. Defaults to 300.
	CacheTime int64 `json:"cache_time,omitempty"`

	// Pass True if results may be cached on the server side only for the user that sent the
	// query. By default, results may be returned to any user who sends the same query.
	IsPersonal bool `json:"is_personal,omitempty"`

	// Pass the offset that a client should send in the next query with the same text to receive
	// more results. Pass an empty string if there are no more results or if you don't support
	// pagination. Offset length can't exceed 64 bytes.
	NextOffset string `json:"next_offset,omitempty"`

	// A JSON-serialized object describing a button to be shown above inline query results
	Button *InlineQueryResultsButton `json:"button,omitempty"`
}

// Response for API call 'answerInlineQuery'
type AnswerInlineQueryResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'answerWebAppQuery'
type AnswerWebAppQueryRequest struct {
	// Unique identifier for the query to be answered
	WebAppQueryID string `json:"web_app_query_id,omitempty"`

	// A JSON-serialized object describing the message to be sent
	Result *InlineQueryResult `json:"result,omitempty"`
}

// Response for API call 'answerWebAppQuery'
type AnswerWebAppQueryResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *SentWebAppMessage `json:"result,omitempty"`
}

// Request for API call 'sendInvoice'
type SendInvoiceRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Product name, 1-32 characters
	Title string `json:"title,omitempty"`

	// Product description, 1-255 characters
	Description string `json:"description,omitempty"`

	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for
	// your internal processes.
	Payload string `json:"payload,omitempty"`

	// Payment provider token, obtained via @BotFather
	ProviderToken string `json:"provider_token,omitempty"`

	// Three-letter ISO 4217 currency code, see more on currencies
	Currency string `json:"currency,omitempty"`

	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount,
	// delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices,omitempty"`

	// The maximum accepted amount for tips in the smallest units of the currency (integer, not
	// float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See
	// the exp parameter in currencies.json, it shows the number of digits past the decimal point
	// for each currency (2 for the majority of currencies). Defaults to 0
	MaxTipAmount int64 `json:"max_tip_amount,omitempty"`

	// A JSON-serialized array of suggested amounts of tips in the smallest units of the currency
	// (integer, not float/double). At most 4 suggested tip amounts can be specified. The
	// suggested tip amounts must be positive, passed in a strictly increased order and must not
	// exceed max_tip_amount.
	SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty"`

	// Unique deep-linking parameter. If left empty, forwarded copies of the sent message will
	// have a Pay button, allowing multiple users to pay directly from the forwarded message,
	// using the same invoice. If non-empty, forwarded copies of the sent message will have a URL
	// button with a deep link to the bot (instead of a Pay button), with the value used as the
	// start parameter
	StartParameter string `json:"start_parameter,omitempty"`

	// JSON-serialized data about the invoice, which will be shared with the payment provider. A
	// detailed description of required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image
	// for a service. People like it better when they see what they are paying for.
	PhotoURL string `json:"photo_url,omitempty"`

	// Photo size in bytes
	PhotoSize int64 `json:"photo_size,omitempty"`

	// Photo width
	PhotoWidth int64 `json:"photo_width,omitempty"`

	// Photo height
	PhotoHeight int64 `json:"photo_height,omitempty"`

	// Pass True if you require the user's full name to complete the order
	NeedName bool `json:"need_name,omitempty"`

	// Pass True if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// Pass True if you require the user's email address to complete the order
	NeedEmail bool `json:"need_email,omitempty"`

	// Pass True if you require the user's shipping address to complete the order
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// Pass True if the user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// Pass True if the user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// Pass True if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button
	// will be shown. If not empty, the first button must be a Pay button.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Response for API call 'sendInvoice'
type SendInvoiceResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'createInvoiceLink'
type CreateInvoiceLinkRequest struct {
	// Product name, 1-32 characters
	Title string `json:"title,omitempty"`

	// Product description, 1-255 characters
	Description string `json:"description,omitempty"`

	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for
	// your internal processes.
	Payload string `json:"payload,omitempty"`

	// Payment provider token, obtained via BotFather
	ProviderToken string `json:"provider_token,omitempty"`

	// Three-letter ISO 4217 currency code, see more on currencies
	Currency string `json:"currency,omitempty"`

	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount,
	// delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices,omitempty"`

	// The maximum accepted amount for tips in the smallest units of the currency (integer, not
	// float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See
	// the exp parameter in currencies.json, it shows the number of digits past the decimal point
	// for each currency (2 for the majority of currencies). Defaults to 0
	MaxTipAmount int64 `json:"max_tip_amount,omitempty"`

	// A JSON-serialized array of suggested amounts of tips in the smallest units of the currency
	// (integer, not float/double). At most 4 suggested tip amounts can be specified. The
	// suggested tip amounts must be positive, passed in a strictly increased order and must not
	// exceed max_tip_amount.
	SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty"`

	// JSON-serialized data about the invoice, which will be shared with the payment provider. A
	// detailed description of required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image
	// for a service.
	PhotoURL string `json:"photo_url,omitempty"`

	// Photo size in bytes
	PhotoSize int64 `json:"photo_size,omitempty"`

	// Photo width
	PhotoWidth int64 `json:"photo_width,omitempty"`

	// Photo height
	PhotoHeight int64 `json:"photo_height,omitempty"`

	// Pass True if you require the user's full name to complete the order
	NeedName bool `json:"need_name,omitempty"`

	// Pass True if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// Pass True if you require the user's email address to complete the order
	NeedEmail bool `json:"need_email,omitempty"`

	// Pass True if you require the user's shipping address to complete the order
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// Pass True if the user's phone number should be sent to the provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// Pass True if the user's email address should be sent to the provider
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// Pass True if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible,omitempty"`
}

// Response for API call 'createInvoiceLink'
type CreateInvoiceLinkResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'answerShippingQuery'
type AnswerShippingQueryRequest struct {
	// Unique identifier for the query to be answered
	ShippingQueryID string `json:"shipping_query_id,omitempty"`

	// Pass True if delivery to the specified address is possible and False if there are any
	// problems (for example, if delivery to the specified address is not possible)
	Ok bool `json:"ok,omitempty"`

	// Required if ok is True. A JSON-serialized array of available shipping options.
	ShippingOptions []*ShippingOption `json:"shipping_options,omitempty"`

	// Required if ok is False. Error message in human readable form that explains why it is
	// impossible to complete the order (e.g. "Sorry, delivery to your desired address is
	// unavailable'). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

// Response for API call 'answerShippingQuery'
type AnswerShippingQueryResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'answerPreCheckoutQuery'
type AnswerPreCheckoutQueryRequest struct {
	// Unique identifier for the query to be answered
	PreCheckoutQueryID string `json:"pre_checkout_query_id,omitempty"`

	// Specify True if everything is alright (goods are available, etc.) and the bot is ready to
	// proceed with the order. Use False if there are any problems.
	Ok bool `json:"ok,omitempty"`

	// Required if ok is False. Error message in human readable form that explains the reason for
	// failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our
	// amazing black T-shirts while you were busy filling out your payment details. Please choose
	// a different color or garment!"). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

// Response for API call 'answerPreCheckoutQuery'
type AnswerPreCheckoutQueryResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'setPassportDataErrors'
type SetPassportDataErrorsRequest struct {
	// User identifier
	UserID int64 `json:"user_id,omitempty"`

	// A JSON-serialized array describing the errors
	Errors []*PassportElementError `json:"errors,omitempty"`
}

// Response for API call 'setPassportDataErrors'
type SetPassportDataErrorsResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'sendGame'
type SendGameRequest struct {
	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id,omitempty"`

	// Unique identifier for the target message thread (topic) of the forum; for forum
	// supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Short name of the game, serves as the unique identifier for the game. Set up your games
	// via @BotFather.
	GameShortName string `json:"game_short_name,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button
	// will be shown. If not empty, the first button must launch the game.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Response for API call 'sendGame'
type SendGameResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result *Message `json:"result,omitempty"`
}

// Request for API call 'setGameScore'
type SetGameScoreRequest struct {
	// User identifier
	UserID int64 `json:"user_id,omitempty"`

	// New score, must be non-negative
	Score int64 `json:"score,omitempty"`

	// Pass True if the high score is allowed to decrease. This can be useful when fixing
	// mistakes or banning cheaters
	Force bool `json:"force,omitempty"`

	// Pass True if the game message should not be automatically edited to include the current
	// scoreboard
	DisableEditMessage bool `json:"disable_edit_message,omitempty"`

	// Required if inline_message_id is not specified. Unique identifier for the target chat
	ChatID int64 `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the sent message
	MessageID int64 `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// Response for API call 'setGameScore'
type SetGameScoreResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`
}

// Request for API call 'getGameHighScores'
type GetGameHighScoresRequest struct {
	// Target user id
	UserID int64 `json:"user_id,omitempty"`

	// Required if inline_message_id is not specified. Unique identifier for the target chat
	ChatID int64 `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the sent message
	MessageID int64 `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// Response for API call 'getGameHighScores'
type GetGameHighScoresResponse struct {
	// Raw response from the server
	Raw []byte `json:"raw,omitempty"`

	// Decoded response from the server
	Result []*GameHighScore `json:"result,omitempty"`
}

// Bot interface
type TelegramApi struct {
	bot TelegramBot
}

func NewTelegramApi(bot TelegramBot) *TelegramApi {
	return &TelegramApi{bot: bot}
}

// Use this method to receive incoming updates using long polling ( wiki ). Returns an Array of
// Update objects.
func (a *TelegramApi) GetUpdates(request *GetUpdatesRequest) (*GetUpdatesResponse, error) {
	apiResponse, err := queryAndUnmarshal[[]*Update](a.bot, "GetUpdates", request)
	if err != nil {
		return nil, err
	}
	return &GetUpdatesResponse{Result: apiResponse.Result}, nil
}

// Notes 1. This method will not work if an outgoing webhook is set up. 2. In order to avoid
// getting duplicate updates, recalculate offset after each server response.    Use this method
// to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an
// update for the bot, we will send an HTTPS POST request to the specified URL, containing a
// JSON-serialized Update . In case of an unsuccessful request, we will give up after a
// reasonable amount of attempts. Returns True on success.  If you'd like to make sure that the
// webhook was set by you, you can specify secret data in the parameter secret_token . If
// specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the
// secret token as content.
func (a *TelegramApi) SetWebhook(request *SetWebhookRequest) (*SetWebhookResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetWebhook", request)
	if err != nil {
		return nil, err
	}
	return &SetWebhookResponse{}, nil
}

// Notes 1. You will not be able to receive updates using getUpdates for as long as an outgoing
// webhook is set up. 2. To use a self-signed certificate, you need to upload your public key
// certificate using certificate parameter. Please upload as InputFile, sending a String will
// not work. 3. Ports currently supported for webhooks : 443, 80, 88, 8443 .  If you're having
// any trouble setting up webhooks, please check out this amazing guide to webhooks .    Use
// this method to remove webhook integration if you decide to switch back to getUpdates .
// Returns True on success.
func (a *TelegramApi) DeleteWebhook(request *DeleteWebhookRequest) (*DeleteWebhookResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "DeleteWebhook", request)
	if err != nil {
		return nil, err
	}
	return &DeleteWebhookResponse{}, nil
}

// Use this method to get current webhook status. Requires no parameters. On success, returns a
// WebhookInfo object. If the bot is using getUpdates , will return an object with the url
// field empty.
func (a *TelegramApi) GetWebhookInfo(request *GetWebhookInfoRequest) (*GetWebhookInfoResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetWebhookInfo", request)
	if err != nil {
		return nil, err
	}
	return &GetWebhookInfoResponse{}, nil
}

// A simple method for testing your bot's authentication token. Requires no parameters. Returns
// basic information about the bot in form of a User object.
func (a *TelegramApi) GetMe(request *GetMeRequest) (*GetMeResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetMe", request)
	if err != nil {
		return nil, err
	}
	return &GetMeResponse{}, nil
}

// Use this method to log out from the cloud Bot API server before launching the bot locally.
// You must log out the bot before running it locally, otherwise there is no guarantee that the
// bot will receive updates. After a successful call, you can immediately log in on a local
// server, but will not be able to log in back to the cloud Bot API server for 10 minutes.
// Returns True on success. Requires no parameters.
func (a *TelegramApi) LogOut(request *LogOutRequest) (*LogOutResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "LogOut", request)
	if err != nil {
		return nil, err
	}
	return &LogOutResponse{}, nil
}

// Use this method to close the bot instance before moving it from one local server to another.
// You need to delete the webhook before calling this method to ensure that the bot isn't
// launched again after server restart. The method will return error 429 in the first 10
// minutes after the bot is launched. Returns True on success. Requires no parameters.
func (a *TelegramApi) Close(request *CloseRequest) (*CloseResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "Close", request)
	if err != nil {
		return nil, err
	}
	return &CloseResponse{}, nil
}

// Use this method to send text messages. On success, the sent Message is returned.
func (a *TelegramApi) SendMessage(request *SendMessageRequest) (*SendMessageResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendMessage", request)
	if err != nil {
		return nil, err
	}
	return &SendMessageResponse{Result: apiResponse.Result}, nil
}

// Use this method to forward messages of any kind. Service messages and messages with
// protected content can't be forwarded. On success, the sent Message is returned.
func (a *TelegramApi) ForwardMessage(request *ForwardMessageRequest) (*ForwardMessageResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "ForwardMessage", request)
	if err != nil {
		return nil, err
	}
	return &ForwardMessageResponse{Result: apiResponse.Result}, nil
}

// Use this method to forward multiple messages of any kind. If some of the specified messages
// can't be found or forwarded, they are skipped. Service messages and messages with protected
// content can't be forwarded. Album grouping is kept for forwarded messages. On success, an
// array of MessageId of the sent messages is returned.
func (a *TelegramApi) ForwardMessages(request *ForwardMessagesRequest) (*ForwardMessagesResponse, error) {
	apiResponse, err := queryAndUnmarshal[[]*MessageId](a.bot, "ForwardMessages", request)
	if err != nil {
		return nil, err
	}
	return &ForwardMessagesResponse{Result: apiResponse.Result}, nil
}

// Use this method to copy messages of any kind. Service messages, giveaway messages, giveaway
// winners messages, and invoice messages can't be copied. A quiz poll can be copied only if
// the value of the field correct_option_id is known to the bot. The method is analogous to the
// method forwardMessage , but the copied message doesn't have a link to the original message.
// Returns the MessageId of the sent message on success.
func (a *TelegramApi) CopyMessage(request *CopyMessageRequest) (*CopyMessageResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "CopyMessage", request)
	if err != nil {
		return nil, err
	}
	return &CopyMessageResponse{}, nil
}

// Use this method to copy messages of any kind. If some of the specified messages can't be
// found or copied, they are skipped. Service messages, giveaway messages, giveaway winners
// messages, and invoice messages can't be copied. A quiz poll can be copied only if the value
// of the field correct_option_id is known to the bot. The method is analogous to the method
// forwardMessages , but the copied messages don't have a link to the original message. Album
// grouping is kept for copied messages. On success, an array of MessageId of the sent messages
// is returned.
func (a *TelegramApi) CopyMessages(request *CopyMessagesRequest) (*CopyMessagesResponse, error) {
	apiResponse, err := queryAndUnmarshal[[]*MessageId](a.bot, "CopyMessages", request)
	if err != nil {
		return nil, err
	}
	return &CopyMessagesResponse{Result: apiResponse.Result}, nil
}

// Use this method to send photos. On success, the sent Message is returned.
func (a *TelegramApi) SendPhoto(request *SendPhotoRequest) (*SendPhotoResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendPhoto", request)
	if err != nil {
		return nil, err
	}
	return &SendPhotoResponse{Result: apiResponse.Result}, nil
}

// Use this method to send audio files, if you want Telegram clients to display them in the
// music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is
// returned. Bots can currently send audio files of up to 50 MB in size, this limit may be
// changed in the future.  For sending voice messages, use the sendVoice method instead.
func (a *TelegramApi) SendAudio(request *SendAudioRequest) (*SendAudioResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendAudio", request)
	if err != nil {
		return nil, err
	}
	return &SendAudioResponse{Result: apiResponse.Result}, nil
}

// Use this method to send general files. On success, the sent Message is returned. Bots can
// currently send files of any type of up to 50 MB in size, this limit may be changed in the
// future.
func (a *TelegramApi) SendDocument(request *SendDocumentRequest) (*SendDocumentResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendDocument", request)
	if err != nil {
		return nil, err
	}
	return &SendDocumentResponse{Result: apiResponse.Result}, nil
}

// Use this method to send video files, Telegram clients support MPEG4 videos (other formats
// may be sent as Document ). On success, the sent Message is returned. Bots can currently send
// video files of up to 50 MB in size, this limit may be changed in the future.
func (a *TelegramApi) SendVideo(request *SendVideoRequest) (*SendVideoResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendVideo", request)
	if err != nil {
		return nil, err
	}
	return &SendVideoResponse{Result: apiResponse.Result}, nil
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On
// success, the sent Message is returned. Bots can currently send animation files of up to 50
// MB in size, this limit may be changed in the future.
func (a *TelegramApi) SendAnimation(request *SendAnimationRequest) (*SendAnimationResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendAnimation", request)
	if err != nil {
		return nil, err
	}
	return &SendAnimationResponse{Result: apiResponse.Result}, nil
}

// Use this method to send audio files, if you want Telegram clients to display the file as a
// playable voice message. For this to work, your audio must be in an .OGG file encoded with
// OPUS (other formats may be sent as Audio or Document ). On success, the sent Message is
// returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be
// changed in the future.
func (a *TelegramApi) SendVoice(request *SendVoiceRequest) (*SendVoiceResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendVoice", request)
	if err != nil {
		return nil, err
	}
	return &SendVoiceResponse{Result: apiResponse.Result}, nil
}

// As of v.4.0 , Telegram clients support rounded square MPEG4 videos of up to 1 minute long.
// Use this method to send video messages. On success, the sent Message is returned.
func (a *TelegramApi) SendVideoNote(request *SendVideoNoteRequest) (*SendVideoNoteResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendVideoNote", request)
	if err != nil {
		return nil, err
	}
	return &SendVideoNoteResponse{Result: apiResponse.Result}, nil
}

// Use this method to send point on the map. On success, the sent Message is returned.
func (a *TelegramApi) SendLocation(request *SendLocationRequest) (*SendLocationResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendLocation", request)
	if err != nil {
		return nil, err
	}
	return &SendLocationResponse{Result: apiResponse.Result}, nil
}

// Use this method to send information about a venue. On success, the sent Message is returned.
func (a *TelegramApi) SendVenue(request *SendVenueRequest) (*SendVenueResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendVenue", request)
	if err != nil {
		return nil, err
	}
	return &SendVenueResponse{Result: apiResponse.Result}, nil
}

// Use this method to send phone contacts. On success, the sent Message is returned.
func (a *TelegramApi) SendContact(request *SendContactRequest) (*SendContactResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendContact", request)
	if err != nil {
		return nil, err
	}
	return &SendContactResponse{Result: apiResponse.Result}, nil
}

// Use this method to send a native poll. On success, the sent Message is returned.
func (a *TelegramApi) SendPoll(request *SendPollRequest) (*SendPollResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendPoll", request)
	if err != nil {
		return nil, err
	}
	return &SendPollResponse{Result: apiResponse.Result}, nil
}

// Use this method to send an animated emoji that will display a random value. On success, the
// sent Message is returned.
func (a *TelegramApi) SendDice(request *SendDiceRequest) (*SendDiceResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendDice", request)
	if err != nil {
		return nil, err
	}
	return &SendDiceResponse{Result: apiResponse.Result}, nil
}

// Use this method when you need to tell the user that something is happening on the bot's
// side. The status is set for 5 seconds or less (when a message arrives from your bot,
// Telegram clients clear its typing status). Returns True on success.   Example: The ImageBot
// needs some time to process a request and upload the image. Instead of sending a text message
// along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with
// action = upload_photo . The user will see a “sending photo” status for the bot.   We only
// recommend using this method when a response from the bot will take a noticeable amount of
// time to arrive.
func (a *TelegramApi) SendChatAction(request *SendChatActionRequest) (*SendChatActionResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SendChatAction", request)
	if err != nil {
		return nil, err
	}
	return &SendChatActionResponse{}, nil
}

// Use this method to change the chosen reactions on a message. Service messages can't be
// reacted to. Automatically forwarded messages from a channel to its discussion group have the
// same available reactions as messages in the channel. Returns True on success.
func (a *TelegramApi) SetMessageReaction(request *SetMessageReactionRequest) (*SetMessageReactionResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetMessageReaction", request)
	if err != nil {
		return nil, err
	}
	return &SetMessageReactionResponse{}, nil
}

// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos
// object.
func (a *TelegramApi) GetUserProfilePhotos(request *GetUserProfilePhotosRequest) (*GetUserProfilePhotosResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetUserProfilePhotos", request)
	if err != nil {
		return nil, err
	}
	return &GetUserProfilePhotosResponse{}, nil
}

// Use this method to get basic information about a file and prepare it for downloading. For
// the moment, bots can download files of up to 20MB in size. On success, a File object is
// returned. The file can then be downloaded via the link
// https://api.telegram.org/file/bot<token>/<file_path> , where <file_path> is taken from the
// response. It is guaranteed that the link will be valid for at least 1 hour. When the link
// expires, a new one can be requested by calling getFile again.
func (a *TelegramApi) GetFile(request *GetFileRequest) (*GetFileResponse, error) {
	apiResponse, err := queryAndUnmarshal[*File](a.bot, "GetFile", request)
	if err != nil {
		return nil, err
	}
	return &GetFileResponse{Result: apiResponse.Result}, nil
}

// Note: This function may not preserve the original file name and MIME type. You should save
// the file's MIME type and name (if available) when the File object is received.   Use this
// method to ban a user in a group, a supergroup or a channel. In the case of supergroups and
// channels, the user will not be able to return to the chat on their own using invite links,
// etc., unless unbanned first. The bot must be an administrator in the chat for this to work
// and must have the appropriate administrator rights. Returns True on success.
func (a *TelegramApi) BanChatMember(request *BanChatMemberRequest) (*BanChatMemberResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "BanChatMember", request)
	if err != nil {
		return nil, err
	}
	return &BanChatMemberResponse{}, nil
}

// Use this method to unban a previously banned user in a supergroup or channel. The user will
// not return to the group or channel automatically, but will be able to join via link, etc.
// The bot must be an administrator for this to work. By default, this method guarantees that
// after the call the user is not a member of the chat, but will be able to join it. So if the
// user is a member of the chat they will also be removed from the chat. If you don't want
// this, use the parameter only_if_banned . Returns True on success.
func (a *TelegramApi) UnbanChatMember(request *UnbanChatMemberRequest) (*UnbanChatMemberResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "UnbanChatMember", request)
	if err != nil {
		return nil, err
	}
	return &UnbanChatMemberResponse{}, nil
}

// Use this method to restrict a user in a supergroup. The bot must be an administrator in the
// supergroup for this to work and must have the appropriate administrator rights. Pass True
// for all permissions to lift restrictions from a user. Returns True on success.
func (a *TelegramApi) RestrictChatMember(request *RestrictChatMemberRequest) (*RestrictChatMemberResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "RestrictChatMember", request)
	if err != nil {
		return nil, err
	}
	return &RestrictChatMemberResponse{}, nil
}

// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an
// administrator in the chat for this to work and must have the appropriate administrator
// rights. Pass False for all boolean parameters to demote a user. Returns True on success.
func (a *TelegramApi) PromoteChatMember(request *PromoteChatMemberRequest) (*PromoteChatMemberResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "PromoteChatMember", request)
	if err != nil {
		return nil, err
	}
	return &PromoteChatMemberResponse{}, nil
}

// Use this method to set a custom title for an administrator in a supergroup promoted by the
// bot. Returns True on success.
func (a *TelegramApi) SetChatAdministratorCustomTitle(request *SetChatAdministratorCustomTitleRequest) (*SetChatAdministratorCustomTitleResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetChatAdministratorCustomTitle", request)
	if err != nil {
		return nil, err
	}
	return &SetChatAdministratorCustomTitleResponse{}, nil
}

// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is
// unbanned , the owner of the banned chat won't be able to send messages on behalf of any of
// their channels . The bot must be an administrator in the supergroup or channel for this to
// work and must have the appropriate administrator rights. Returns True on success.
func (a *TelegramApi) BanChatSenderChat(request *BanChatSenderChatRequest) (*BanChatSenderChatResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "BanChatSenderChat", request)
	if err != nil {
		return nil, err
	}
	return &BanChatSenderChatResponse{}, nil
}

// Use this method to unban a previously banned channel chat in a supergroup or channel. The
// bot must be an administrator for this to work and must have the appropriate administrator
// rights. Returns True on success.
func (a *TelegramApi) UnbanChatSenderChat(request *UnbanChatSenderChatRequest) (*UnbanChatSenderChatResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "UnbanChatSenderChat", request)
	if err != nil {
		return nil, err
	}
	return &UnbanChatSenderChatResponse{}, nil
}

// Use this method to set default chat permissions for all members. The bot must be an
// administrator in the group or a supergroup for this to work and must have the
// can_restrict_members administrator rights. Returns True on success.
func (a *TelegramApi) SetChatPermissions(request *SetChatPermissionsRequest) (*SetChatPermissionsResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetChatPermissions", request)
	if err != nil {
		return nil, err
	}
	return &SetChatPermissionsResponse{}, nil
}

// Use this method to generate a new primary invite link for a chat; any previously generated
// primary link is revoked. The bot must be an administrator in the chat for this to work and
// must have the appropriate administrator rights. Returns the new invite link as String on
// success.
func (a *TelegramApi) ExportChatInviteLink(request *ExportChatInviteLinkRequest) (*ExportChatInviteLinkResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "ExportChatInviteLink", request)
	if err != nil {
		return nil, err
	}
	return &ExportChatInviteLinkResponse{}, nil
}

// Note: Each administrator in a chat generates their own invite links. Bots can't use invite
// links generated by other administrators. If you want your bot to work with invite links, it
// will need to generate its own link using exportChatInviteLink or by calling the getChat
// method. If your bot needs to generate a new primary invite link replacing its previous one,
// use exportChatInviteLink again.    Use this method to create an additional invite link for a
// chat. The bot must be an administrator in the chat for this to work and must have the
// appropriate administrator rights. The link can be revoked using the method
// revokeChatInviteLink . Returns the new invite link as ChatInviteLink object.
func (a *TelegramApi) CreateChatInviteLink(request *CreateChatInviteLinkRequest) (*CreateChatInviteLinkResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "CreateChatInviteLink", request)
	if err != nil {
		return nil, err
	}
	return &CreateChatInviteLinkResponse{}, nil
}

// Use this method to edit a non-primary invite link created by the bot. The bot must be an
// administrator in the chat for this to work and must have the appropriate administrator
// rights. Returns the edited invite link as a ChatInviteLink object.
func (a *TelegramApi) EditChatInviteLink(request *EditChatInviteLinkRequest) (*EditChatInviteLinkResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "EditChatInviteLink", request)
	if err != nil {
		return nil, err
	}
	return &EditChatInviteLinkResponse{}, nil
}

// Use this method to revoke an invite link created by the bot. If the primary link is revoked,
// a new link is automatically generated. The bot must be an administrator in the chat for this
// to work and must have the appropriate administrator rights. Returns the revoked invite link
// as ChatInviteLink object.
func (a *TelegramApi) RevokeChatInviteLink(request *RevokeChatInviteLinkRequest) (*RevokeChatInviteLinkResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "RevokeChatInviteLink", request)
	if err != nil {
		return nil, err
	}
	return &RevokeChatInviteLinkResponse{}, nil
}

// Use this method to approve a chat join request. The bot must be an administrator in the chat
// for this to work and must have the can_invite_users administrator right. Returns True on
// success.
func (a *TelegramApi) ApproveChatJoinRequest(request *ApproveChatJoinRequestRequest) (*ApproveChatJoinRequestResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "ApproveChatJoinRequest", request)
	if err != nil {
		return nil, err
	}
	return &ApproveChatJoinRequestResponse{}, nil
}

// Use this method to decline a chat join request. The bot must be an administrator in the chat
// for this to work and must have the can_invite_users administrator right. Returns True on
// success.
func (a *TelegramApi) DeclineChatJoinRequest(request *DeclineChatJoinRequestRequest) (*DeclineChatJoinRequestResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "DeclineChatJoinRequest", request)
	if err != nil {
		return nil, err
	}
	return &DeclineChatJoinRequestResponse{}, nil
}

// Use this method to set a new profile photo for the chat. Photos can't be changed for private
// chats. The bot must be an administrator in the chat for this to work and must have the
// appropriate administrator rights. Returns True on success.
func (a *TelegramApi) SetChatPhoto(request *SetChatPhotoRequest) (*SetChatPhotoResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetChatPhoto", request)
	if err != nil {
		return nil, err
	}
	return &SetChatPhotoResponse{}, nil
}

// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot
// must be an administrator in the chat for this to work and must have the appropriate
// administrator rights. Returns True on success.
func (a *TelegramApi) DeleteChatPhoto(request *DeleteChatPhotoRequest) (*DeleteChatPhotoResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "DeleteChatPhoto", request)
	if err != nil {
		return nil, err
	}
	return &DeleteChatPhotoResponse{}, nil
}

// Use this method to change the title of a chat. Titles can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate
// administrator rights. Returns True on success.
func (a *TelegramApi) SetChatTitle(request *SetChatTitleRequest) (*SetChatTitleResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetChatTitle", request)
	if err != nil {
		return nil, err
	}
	return &SetChatTitleResponse{}, nil
}

// Use this method to change the description of a group, a supergroup or a channel. The bot
// must be an administrator in the chat for this to work and must have the appropriate
// administrator rights. Returns True on success.
func (a *TelegramApi) SetChatDescription(request *SetChatDescriptionRequest) (*SetChatDescriptionResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetChatDescription", request)
	if err != nil {
		return nil, err
	}
	return &SetChatDescriptionResponse{}, nil
}

// Use this method to add a message to the list of pinned messages in a chat. If the chat is
// not a private chat, the bot must be an administrator in the chat for this to work and must
// have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages'
// administrator right in a channel. Returns True on success.
func (a *TelegramApi) PinChatMessage(request *PinChatMessageRequest) (*PinChatMessageResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "PinChatMessage", request)
	if err != nil {
		return nil, err
	}
	return &PinChatMessageResponse{}, nil
}

// Use this method to remove a message from the list of pinned messages in a chat. If the chat
// is not a private chat, the bot must be an administrator in the chat for this to work and
// must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages'
// administrator right in a channel. Returns True on success.
func (a *TelegramApi) UnpinChatMessage(request *UnpinChatMessageRequest) (*UnpinChatMessageResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "UnpinChatMessage", request)
	if err != nil {
		return nil, err
	}
	return &UnpinChatMessageResponse{}, nil
}

// Use this method to clear the list of pinned messages in a chat. If the chat is not a private
// chat, the bot must be an administrator in the chat for this to work and must have the
// 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator
// right in a channel. Returns True on success.
func (a *TelegramApi) UnpinAllChatMessages(request *UnpinAllChatMessagesRequest) (*UnpinAllChatMessagesResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "UnpinAllChatMessages", request)
	if err != nil {
		return nil, err
	}
	return &UnpinAllChatMessagesResponse{}, nil
}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on
// success.
func (a *TelegramApi) LeaveChat(request *LeaveChatRequest) (*LeaveChatResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "LeaveChat", request)
	if err != nil {
		return nil, err
	}
	return &LeaveChatResponse{}, nil
}

// Use this method to get up to date information about the chat. Returns a Chat object on
// success.
func (a *TelegramApi) GetChat(request *GetChatRequest) (*GetChatResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetChat", request)
	if err != nil {
		return nil, err
	}
	return &GetChatResponse{}, nil
}

// Use this method to get a list of administrators in a chat, which aren't bots. Returns an
// Array of ChatMember objects.
func (a *TelegramApi) GetChatAdministrators(request *GetChatAdministratorsRequest) (*GetChatAdministratorsResponse, error) {
	apiResponse, err := queryAndUnmarshal[[]*ChatMember](a.bot, "GetChatAdministrators", request)
	if err != nil {
		return nil, err
	}
	return &GetChatAdministratorsResponse{Result: apiResponse.Result}, nil
}

// Use this method to get the number of members in a chat. Returns Int on success.
func (a *TelegramApi) GetChatMemberCount(request *GetChatMemberCountRequest) (*GetChatMemberCountResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetChatMemberCount", request)
	if err != nil {
		return nil, err
	}
	return &GetChatMemberCountResponse{}, nil
}

// Use this method to get information about a member of a chat. The method is only guaranteed
// to work for other users if the bot is an administrator in the chat. Returns a ChatMember
// object on success.
func (a *TelegramApi) GetChatMember(request *GetChatMemberRequest) (*GetChatMemberResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetChatMember", request)
	if err != nil {
		return nil, err
	}
	return &GetChatMemberResponse{}, nil
}

// Use this method to set a new group sticker set for a supergroup. The bot must be an
// administrator in the chat for this to work and must have the appropriate administrator
// rights. Use the field can_set_sticker_set optionally returned in getChat requests to check
// if the bot can use this method. Returns True on success.
func (a *TelegramApi) SetChatStickerSet(request *SetChatStickerSetRequest) (*SetChatStickerSetResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetChatStickerSet", request)
	if err != nil {
		return nil, err
	}
	return &SetChatStickerSetResponse{}, nil
}

// Use this method to delete a group sticker set from a supergroup. The bot must be an
// administrator in the chat for this to work and must have the appropriate administrator
// rights. Use the field can_set_sticker_set optionally returned in getChat requests to check
// if the bot can use this method. Returns True on success.
func (a *TelegramApi) DeleteChatStickerSet(request *DeleteChatStickerSetRequest) (*DeleteChatStickerSetResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "DeleteChatStickerSet", request)
	if err != nil {
		return nil, err
	}
	return &DeleteChatStickerSetResponse{}, nil
}

// Use this method to get custom emoji stickers, which can be used as a forum topic icon by any
// user. Requires no parameters. Returns an Array of Sticker objects.
func (a *TelegramApi) GetForumTopicIconStickers(request *GetForumTopicIconStickersRequest) (*GetForumTopicIconStickersResponse, error) {
	apiResponse, err := queryAndUnmarshal[[]*Sticker](a.bot, "GetForumTopicIconStickers", request)
	if err != nil {
		return nil, err
	}
	return &GetForumTopicIconStickersResponse{Result: apiResponse.Result}, nil
}

// Use this method to create a topic in a forum supergroup chat. The bot must be an
// administrator in the chat for this to work and must have the can_manage_topics administrator
// rights. Returns information about the created topic as a ForumTopic object.
func (a *TelegramApi) CreateForumTopic(request *CreateForumTopicRequest) (*CreateForumTopicResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "CreateForumTopic", request)
	if err != nil {
		return nil, err
	}
	return &CreateForumTopicResponse{}, nil
}

// Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must be
// an administrator in the chat for this to work and must have can_manage_topics administrator
// rights, unless it is the creator of the topic. Returns True on success.
func (a *TelegramApi) EditForumTopic(request *EditForumTopicRequest) (*EditForumTopicResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "EditForumTopic", request)
	if err != nil {
		return nil, err
	}
	return &EditForumTopicResponse{}, nil
}

// Use this method to close an open topic in a forum supergroup chat. The bot must be an
// administrator in the chat for this to work and must have the can_manage_topics administrator
// rights, unless it is the creator of the topic. Returns True on success.
func (a *TelegramApi) CloseForumTopic(request *CloseForumTopicRequest) (*CloseForumTopicResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "CloseForumTopic", request)
	if err != nil {
		return nil, err
	}
	return &CloseForumTopicResponse{}, nil
}

// Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an
// administrator in the chat for this to work and must have the can_manage_topics administrator
// rights, unless it is the creator of the topic. Returns True on success.
func (a *TelegramApi) ReopenForumTopic(request *ReopenForumTopicRequest) (*ReopenForumTopicResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "ReopenForumTopic", request)
	if err != nil {
		return nil, err
	}
	return &ReopenForumTopicResponse{}, nil
}

// Use this method to delete a forum topic along with all its messages in a forum supergroup
// chat. The bot must be an administrator in the chat for this to work and must have the
// can_delete_messages administrator rights. Returns True on success.
func (a *TelegramApi) DeleteForumTopic(request *DeleteForumTopicRequest) (*DeleteForumTopicResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "DeleteForumTopic", request)
	if err != nil {
		return nil, err
	}
	return &DeleteForumTopicResponse{}, nil
}

// Use this method to clear the list of pinned messages in a forum topic. The bot must be an
// administrator in the chat for this to work and must have the can_pin_messages administrator
// right in the supergroup. Returns True on success.
func (a *TelegramApi) UnpinAllForumTopicMessages(request *UnpinAllForumTopicMessagesRequest) (*UnpinAllForumTopicMessagesResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "UnpinAllForumTopicMessages", request)
	if err != nil {
		return nil, err
	}
	return &UnpinAllForumTopicMessagesResponse{}, nil
}

// Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot
// must be an administrator in the chat for this to work and must have can_manage_topics
// administrator rights. Returns True on success.
func (a *TelegramApi) EditGeneralForumTopic(request *EditGeneralForumTopicRequest) (*EditGeneralForumTopicResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "EditGeneralForumTopic", request)
	if err != nil {
		return nil, err
	}
	return &EditGeneralForumTopicResponse{}, nil
}

// Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be
// an administrator in the chat for this to work and must have the can_manage_topics
// administrator rights. Returns True on success.
func (a *TelegramApi) CloseGeneralForumTopic(request *CloseGeneralForumTopicRequest) (*CloseGeneralForumTopicResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "CloseGeneralForumTopic", request)
	if err != nil {
		return nil, err
	}
	return &CloseGeneralForumTopicResponse{}, nil
}

// Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must
// be an administrator in the chat for this to work and must have the can_manage_topics
// administrator rights. The topic will be automatically unhidden if it was hidden. Returns
// True on success.
func (a *TelegramApi) ReopenGeneralForumTopic(request *ReopenGeneralForumTopicRequest) (*ReopenGeneralForumTopicResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "ReopenGeneralForumTopic", request)
	if err != nil {
		return nil, err
	}
	return &ReopenGeneralForumTopicResponse{}, nil
}

// Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an
// administrator in the chat for this to work and must have the can_manage_topics administrator
// rights. The topic will be automatically closed if it was open. Returns True on success.
func (a *TelegramApi) HideGeneralForumTopic(request *HideGeneralForumTopicRequest) (*HideGeneralForumTopicResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "HideGeneralForumTopic", request)
	if err != nil {
		return nil, err
	}
	return &HideGeneralForumTopicResponse{}, nil
}

// Use this method to unhide the 'General' topic in a forum supergroup chat. The bot must be an
// administrator in the chat for this to work and must have the can_manage_topics administrator
// rights. Returns True on success.
func (a *TelegramApi) UnhideGeneralForumTopic(request *UnhideGeneralForumTopicRequest) (*UnhideGeneralForumTopicResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "UnhideGeneralForumTopic", request)
	if err != nil {
		return nil, err
	}
	return &UnhideGeneralForumTopicResponse{}, nil
}

// Use this method to clear the list of pinned messages in a General forum topic. The bot must
// be an administrator in the chat for this to work and must have the can_pin_messages
// administrator right in the supergroup. Returns True on success.
func (a *TelegramApi) UnpinAllGeneralForumTopicMessages(request *UnpinAllGeneralForumTopicMessagesRequest) (*UnpinAllGeneralForumTopicMessagesResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "UnpinAllGeneralForumTopicMessages", request)
	if err != nil {
		return nil, err
	}
	return &UnpinAllGeneralForumTopicMessagesResponse{}, nil
}

// Use this method to send answers to callback queries sent from inline keyboards . The answer
// will be displayed to the user as a notification at the top of the chat screen or as an
// alert. On success, True is returned.   Alternatively, the user can be redirected to the
// specified Game URL. For this option to work, you must first create a game for your bot via
// @BotFather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX
// that open your bot with a parameter.
func (a *TelegramApi) AnswerCallbackQuery(request *AnswerCallbackQueryRequest) (*AnswerCallbackQueryResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "AnswerCallbackQuery", request)
	if err != nil {
		return nil, err
	}
	return &AnswerCallbackQueryResponse{}, nil
}

// Use this method to get the list of boosts added to a chat by a user. Requires administrator
// rights in the chat. Returns a UserChatBoosts object.
func (a *TelegramApi) GetUserChatBoosts(request *GetUserChatBoostsRequest) (*GetUserChatBoostsResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetUserChatBoosts", request)
	if err != nil {
		return nil, err
	}
	return &GetUserChatBoostsResponse{}, nil
}

// Use this method to change the list of the bot's commands. See this manual for more details
// about bot commands. Returns True on success.
func (a *TelegramApi) SetMyCommands(request *SetMyCommandsRequest) (*SetMyCommandsResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetMyCommands", request)
	if err != nil {
		return nil, err
	}
	return &SetMyCommandsResponse{}, nil
}

// Use this method to delete the list of the bot's commands for the given scope and user
// language. After deletion, higher level commands will be shown to affected users. Returns
// True on success.
func (a *TelegramApi) DeleteMyCommands(request *DeleteMyCommandsRequest) (*DeleteMyCommandsResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "DeleteMyCommands", request)
	if err != nil {
		return nil, err
	}
	return &DeleteMyCommandsResponse{}, nil
}

// Use this method to get the current list of the bot's commands for the given scope and user
// language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is
// returned.
func (a *TelegramApi) GetMyCommands(request *GetMyCommandsRequest) (*GetMyCommandsResponse, error) {
	apiResponse, err := queryAndUnmarshal[[]*BotCommand](a.bot, "GetMyCommands", request)
	if err != nil {
		return nil, err
	}
	return &GetMyCommandsResponse{Result: apiResponse.Result}, nil
}

// Use this method to change the bot's name. Returns True on success.
func (a *TelegramApi) SetMyName(request *SetMyNameRequest) (*SetMyNameResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetMyName", request)
	if err != nil {
		return nil, err
	}
	return &SetMyNameResponse{}, nil
}

// Use this method to get the current bot name for the given user language. Returns BotName on
// success.
func (a *TelegramApi) GetMyName(request *GetMyNameRequest) (*GetMyNameResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetMyName", request)
	if err != nil {
		return nil, err
	}
	return &GetMyNameResponse{}, nil
}

// Use this method to change the bot's description, which is shown in the chat with the bot if
// the chat is empty. Returns True on success.
func (a *TelegramApi) SetMyDescription(request *SetMyDescriptionRequest) (*SetMyDescriptionResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetMyDescription", request)
	if err != nil {
		return nil, err
	}
	return &SetMyDescriptionResponse{}, nil
}

// Use this method to get the current bot description for the given user language. Returns
// BotDescription on success.
func (a *TelegramApi) GetMyDescription(request *GetMyDescriptionRequest) (*GetMyDescriptionResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetMyDescription", request)
	if err != nil {
		return nil, err
	}
	return &GetMyDescriptionResponse{}, nil
}

// Use this method to change the bot's short description, which is shown on the bot's profile
// page and is sent together with the link when users share the bot. Returns True on success.
func (a *TelegramApi) SetMyShortDescription(request *SetMyShortDescriptionRequest) (*SetMyShortDescriptionResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetMyShortDescription", request)
	if err != nil {
		return nil, err
	}
	return &SetMyShortDescriptionResponse{}, nil
}

// Use this method to get the current bot short description for the given user language.
// Returns BotShortDescription on success.
func (a *TelegramApi) GetMyShortDescription(request *GetMyShortDescriptionRequest) (*GetMyShortDescriptionResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetMyShortDescription", request)
	if err != nil {
		return nil, err
	}
	return &GetMyShortDescriptionResponse{}, nil
}

// Use this method to change the bot's menu button in a private chat, or the default menu
// button. Returns True on success.
func (a *TelegramApi) SetChatMenuButton(request *SetChatMenuButtonRequest) (*SetChatMenuButtonResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetChatMenuButton", request)
	if err != nil {
		return nil, err
	}
	return &SetChatMenuButtonResponse{}, nil
}

// Use this method to get the current value of the bot's menu button in a private chat, or the
// default menu button. Returns MenuButton on success.
func (a *TelegramApi) GetChatMenuButton(request *GetChatMenuButtonRequest) (*GetChatMenuButtonResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetChatMenuButton", request)
	if err != nil {
		return nil, err
	}
	return &GetChatMenuButtonResponse{}, nil
}

// Use this method to change the default administrator rights requested by the bot when it's
// added as an administrator to groups or channels. These rights will be suggested to users,
// but they are free to modify the list before adding the bot. Returns True on success.
func (a *TelegramApi) SetMyDefaultAdministratorRights(request *SetMyDefaultAdministratorRightsRequest) (*SetMyDefaultAdministratorRightsResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetMyDefaultAdministratorRights", request)
	if err != nil {
		return nil, err
	}
	return &SetMyDefaultAdministratorRightsResponse{}, nil
}

// Use this method to get the current default administrator rights of the bot. Returns
// ChatAdministratorRights on success.
func (a *TelegramApi) GetMyDefaultAdministratorRights(request *GetMyDefaultAdministratorRightsRequest) (*GetMyDefaultAdministratorRightsResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "GetMyDefaultAdministratorRights", request)
	if err != nil {
		return nil, err
	}
	return &GetMyDefaultAdministratorRightsResponse{}, nil
}

// Use this method to edit text and game messages. On success, if the edited message is not an
// inline message, the edited Message is returned, otherwise True is returned.
func (a *TelegramApi) EditMessageText(request *EditMessageTextRequest) (*EditMessageTextResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "EditMessageText", request)
	if err != nil {
		return nil, err
	}
	return &EditMessageTextResponse{Result: apiResponse.Result}, nil
}

// Use this method to edit captions of messages. On success, if the edited message is not an
// inline message, the edited Message is returned, otherwise True is returned.
func (a *TelegramApi) EditMessageCaption(request *EditMessageCaptionRequest) (*EditMessageCaptionResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "EditMessageCaption", request)
	if err != nil {
		return nil, err
	}
	return &EditMessageCaptionResponse{Result: apiResponse.Result}, nil
}

// Use this method to edit animation, audio, document, photo, or video messages. If a message
// is part of a message album, then it can be edited only to an audio for audio albums, only to
// a document for document albums and to a photo or a video otherwise. When an inline message
// is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or
// specify a URL. On success, if the edited message is not an inline message, the edited
// Message is returned, otherwise True is returned.
func (a *TelegramApi) EditMessageMedia(request *EditMessageMediaRequest) (*EditMessageMediaResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "EditMessageMedia", request)
	if err != nil {
		return nil, err
	}
	return &EditMessageMediaResponse{Result: apiResponse.Result}, nil
}

// Use this method to edit live location messages. A location can be edited until its
// live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation .
// On success, if the edited message is not an inline message, the edited Message is returned,
// otherwise True is returned.
func (a *TelegramApi) EditMessageLiveLocation(request *EditMessageLiveLocationRequest) (*EditMessageLiveLocationResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "EditMessageLiveLocation", request)
	if err != nil {
		return nil, err
	}
	return &EditMessageLiveLocationResponse{Result: apiResponse.Result}, nil
}

// Use this method to stop updating a live location message before live_period expires. On
// success, if the message is not an inline message, the edited Message is returned, otherwise
// True is returned.
func (a *TelegramApi) StopMessageLiveLocation(request *StopMessageLiveLocationRequest) (*StopMessageLiveLocationResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "StopMessageLiveLocation", request)
	if err != nil {
		return nil, err
	}
	return &StopMessageLiveLocationResponse{Result: apiResponse.Result}, nil
}

// Use this method to edit only the reply markup of messages. On success, if the edited message
// is not an inline message, the edited Message is returned, otherwise True is returned.
func (a *TelegramApi) EditMessageReplyMarkup(request *EditMessageReplyMarkupRequest) (*EditMessageReplyMarkupResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "EditMessageReplyMarkup", request)
	if err != nil {
		return nil, err
	}
	return &EditMessageReplyMarkupResponse{Result: apiResponse.Result}, nil
}

// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll is
// returned.
func (a *TelegramApi) StopPoll(request *StopPollRequest) (*StopPollResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Poll](a.bot, "StopPoll", request)
	if err != nil {
		return nil, err
	}
	return &StopPollResponse{Result: apiResponse.Result}, nil
}

// Use this method to delete a message, including service messages, with the following
// limitations: - A message can only be deleted if it was sent less than 48 hours ago. -
// Service messages about a supergroup, channel, or forum topic creation can't be deleted. - A
// dice message in a private chat can only be deleted if it was sent more than 24 hours ago. -
// Bots can delete outgoing messages in private chats, groups, and supergroups. - Bots can
// delete incoming messages in private chats. - Bots granted can_post_messages permissions can
// delete outgoing messages in channels. - If the bot is an administrator of a group, it can
// delete any message there. - If the bot has can_delete_messages permission in a supergroup or
// a channel, it can delete any message there. Returns True on success.
func (a *TelegramApi) DeleteMessage(request *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "DeleteMessage", request)
	if err != nil {
		return nil, err
	}
	return &DeleteMessageResponse{}, nil
}

// Use this method to delete multiple messages simultaneously. If some of the specified
// messages can't be found, they are skipped. Returns True on success.
func (a *TelegramApi) DeleteMessages(request *DeleteMessagesRequest) (*DeleteMessagesResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "DeleteMessages", request)
	if err != nil {
		return nil, err
	}
	return &DeleteMessagesResponse{}, nil
}

// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success,
// the sent Message is returned.
func (a *TelegramApi) SendSticker(request *SendStickerRequest) (*SendStickerResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendSticker", request)
	if err != nil {
		return nil, err
	}
	return &SendStickerResponse{Result: apiResponse.Result}, nil
}

// Use this method to get a sticker set. On success, a StickerSet object is returned.
func (a *TelegramApi) GetStickerSet(request *GetStickerSetRequest) (*GetStickerSetResponse, error) {
	apiResponse, err := queryAndUnmarshal[*StickerSet](a.bot, "GetStickerSet", request)
	if err != nil {
		return nil, err
	}
	return &GetStickerSetResponse{Result: apiResponse.Result}, nil
}

// Use this method to get information about custom emoji stickers by their identifiers. Returns
// an Array of Sticker objects.
func (a *TelegramApi) GetCustomEmojiStickers(request *GetCustomEmojiStickersRequest) (*GetCustomEmojiStickersResponse, error) {
	apiResponse, err := queryAndUnmarshal[[]*Sticker](a.bot, "GetCustomEmojiStickers", request)
	if err != nil {
		return nil, err
	}
	return &GetCustomEmojiStickersResponse{Result: apiResponse.Result}, nil
}

// Use this method to upload a file with a sticker for later use in the createNewStickerSet and
// addStickerToSet methods (the file can be used multiple times). Returns the uploaded File on
// success.
func (a *TelegramApi) UploadStickerFile(request *UploadStickerFileRequest) (*UploadStickerFileResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "UploadStickerFile", request)
	if err != nil {
		return nil, err
	}
	return &UploadStickerFileResponse{}, nil
}

// Use this method to create a new sticker set owned by a user. The bot will be able to edit
// the sticker set thus created. Returns True on success.
func (a *TelegramApi) CreateNewStickerSet(request *CreateNewStickerSetRequest) (*CreateNewStickerSetResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "CreateNewStickerSet", request)
	if err != nil {
		return nil, err
	}
	return &CreateNewStickerSetResponse{}, nil
}

// Use this method to add a new sticker to a set created by the bot. The format of the added
// sticker must match the format of the other stickers in the set. Emoji sticker sets can have
// up to 200 stickers. Animated and video sticker sets can have up to 50 stickers. Static
// sticker sets can have up to 120 stickers. Returns True on success.
func (a *TelegramApi) AddStickerToSet(request *AddStickerToSetRequest) (*AddStickerToSetResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "AddStickerToSet", request)
	if err != nil {
		return nil, err
	}
	return &AddStickerToSetResponse{}, nil
}

// Use this method to move a sticker in a set created by the bot to a specific position.
// Returns True on success.
func (a *TelegramApi) SetStickerPositionInSet(request *SetStickerPositionInSetRequest) (*SetStickerPositionInSetResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetStickerPositionInSet", request)
	if err != nil {
		return nil, err
	}
	return &SetStickerPositionInSetResponse{}, nil
}

// Use this method to delete a sticker from a set created by the bot. Returns True on success.
func (a *TelegramApi) DeleteStickerFromSet(request *DeleteStickerFromSetRequest) (*DeleteStickerFromSetResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "DeleteStickerFromSet", request)
	if err != nil {
		return nil, err
	}
	return &DeleteStickerFromSetResponse{}, nil
}

// Use this method to change the list of emoji assigned to a regular or custom emoji sticker.
// The sticker must belong to a sticker set created by the bot. Returns True on success.
func (a *TelegramApi) SetStickerEmojiList(request *SetStickerEmojiListRequest) (*SetStickerEmojiListResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetStickerEmojiList", request)
	if err != nil {
		return nil, err
	}
	return &SetStickerEmojiListResponse{}, nil
}

// Use this method to change search keywords assigned to a regular or custom emoji sticker. The
// sticker must belong to a sticker set created by the bot. Returns True on success.
func (a *TelegramApi) SetStickerKeywords(request *SetStickerKeywordsRequest) (*SetStickerKeywordsResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetStickerKeywords", request)
	if err != nil {
		return nil, err
	}
	return &SetStickerKeywordsResponse{}, nil
}

// Use this method to change the mask position of a mask sticker. The sticker must belong to a
// sticker set that was created by the bot. Returns True on success.
func (a *TelegramApi) SetStickerMaskPosition(request *SetStickerMaskPositionRequest) (*SetStickerMaskPositionResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetStickerMaskPosition", request)
	if err != nil {
		return nil, err
	}
	return &SetStickerMaskPositionResponse{}, nil
}

// Use this method to set the title of a created sticker set. Returns True on success.
func (a *TelegramApi) SetStickerSetTitle(request *SetStickerSetTitleRequest) (*SetStickerSetTitleResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetStickerSetTitle", request)
	if err != nil {
		return nil, err
	}
	return &SetStickerSetTitleResponse{}, nil
}

// Use this method to set the thumbnail of a regular or mask sticker set. The format of the
// thumbnail file must match the format of the stickers in the set. Returns True on success.
func (a *TelegramApi) SetStickerSetThumbnail(request *SetStickerSetThumbnailRequest) (*SetStickerSetThumbnailResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetStickerSetThumbnail", request)
	if err != nil {
		return nil, err
	}
	return &SetStickerSetThumbnailResponse{}, nil
}

// Use this method to set the thumbnail of a custom emoji sticker set. Returns True on success.
func (a *TelegramApi) SetCustomEmojiStickerSetThumbnail(request *SetCustomEmojiStickerSetThumbnailRequest) (*SetCustomEmojiStickerSetThumbnailResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetCustomEmojiStickerSetThumbnail", request)
	if err != nil {
		return nil, err
	}
	return &SetCustomEmojiStickerSetThumbnailResponse{}, nil
}

// Use this method to delete a sticker set that was created by the bot. Returns True on
// success.
func (a *TelegramApi) DeleteStickerSet(request *DeleteStickerSetRequest) (*DeleteStickerSetResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "DeleteStickerSet", request)
	if err != nil {
		return nil, err
	}
	return &DeleteStickerSetResponse{}, nil
}

// Use this method to send answers to an inline query. On success, True is returned. No more
// than 50 results per query are allowed.
func (a *TelegramApi) AnswerInlineQuery(request *AnswerInlineQueryRequest) (*AnswerInlineQueryResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "AnswerInlineQuery", request)
	if err != nil {
		return nil, err
	}
	return &AnswerInlineQueryResponse{}, nil
}

// Note: It is necessary to enable inline feedback via @BotFather in order to receive these
// objects in updates.   Use this method to set the result of an interaction with a Web App and
// send a corresponding message on behalf of the user to the chat from which the query
// originated. On success, a SentWebAppMessage object is returned.
func (a *TelegramApi) AnswerWebAppQuery(request *AnswerWebAppQueryRequest) (*AnswerWebAppQueryResponse, error) {
	apiResponse, err := queryAndUnmarshal[*SentWebAppMessage](a.bot, "AnswerWebAppQuery", request)
	if err != nil {
		return nil, err
	}
	return &AnswerWebAppQueryResponse{Result: apiResponse.Result}, nil
}

// Your bot can accept payments from Telegram users. Please see the introduction to payments
// for more details on the process and how to set up payments for your bot.   Use this method
// to send invoices. On success, the sent Message is returned.
func (a *TelegramApi) SendInvoice(request *SendInvoiceRequest) (*SendInvoiceResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendInvoice", request)
	if err != nil {
		return nil, err
	}
	return &SendInvoiceResponse{Result: apiResponse.Result}, nil
}

// Use this method to create a link for an invoice. Returns the created invoice link as String
// on success.
func (a *TelegramApi) CreateInvoiceLink(request *CreateInvoiceLinkRequest) (*CreateInvoiceLinkResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "CreateInvoiceLink", request)
	if err != nil {
		return nil, err
	}
	return &CreateInvoiceLinkResponse{}, nil
}

// If you sent an invoice requesting a shipping address and the parameter is_flexible was
// specified, the Bot API will send an Update with a shipping_query field to the bot. Use this
// method to reply to shipping queries. On success, True is returned.
func (a *TelegramApi) AnswerShippingQuery(request *AnswerShippingQueryRequest) (*AnswerShippingQueryResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "AnswerShippingQuery", request)
	if err != nil {
		return nil, err
	}
	return &AnswerShippingQueryResponse{}, nil
}

// Once the user has confirmed their payment and shipping details, the Bot API sends the final
// confirmation in the form of an Update with the field pre_checkout_query . Use this method to
// respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must
// receive an answer within 10 seconds after the pre-checkout query was sent.
func (a *TelegramApi) AnswerPreCheckoutQuery(request *AnswerPreCheckoutQueryRequest) (*AnswerPreCheckoutQueryResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "AnswerPreCheckoutQuery", request)
	if err != nil {
		return nil, err
	}
	return &AnswerPreCheckoutQueryResponse{}, nil
}

// Informs a user that some of the Telegram Passport elements they provided contains errors.
// The user will not be able to re-submit their Passport to you until the errors are fixed (the
// contents of the field for which you returned the error must change). Returns True on
// success.  Use this if the data submitted by the user doesn't satisfy the standards your
// service requires for any reason. For example, if a birthday date seems invalid, a submitted
// document is blurry, a scan shows evidence of tampering, etc. Supply some details in the
// error message to make sure the user knows how to correct the issues.
func (a *TelegramApi) SetPassportDataErrors(request *SetPassportDataErrorsRequest) (*SetPassportDataErrorsResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetPassportDataErrors", request)
	if err != nil {
		return nil, err
	}
	return &SetPassportDataErrorsResponse{}, nil
}

// Your bot can offer users HTML5 games to play solo or to compete against each other in groups
// and one-on-one chats. Create games via @BotFather using the /newgame command. Please note
// that this kind of power requires responsibility: you will need to accept the terms for each
// game that your bots will be offering.   Games are a new type of content on Telegram,
// represented by the Game and InlineQueryResultGame objects.  Once you've created a game via
// BotFather , you can send games to chats as regular messages using the sendGame method, or
// use inline mode with InlineQueryResultGame .  If you send the game message without any
// buttons, it will automatically have a 'Play GameName ' button. When this button is pressed,
// your bot gets a CallbackQuery with the game_short_name of the requested game. You provide
// the correct URL for this particular user and the app opens the game in the in-app browser.
// You can manually add multiple buttons to your game message. Please note that the first
// button in the first row must always launch the game, using the field callback_game in
// InlineKeyboardButton . You can add extra buttons according to taste: e.g., for a description
// of the rules, or to open the game's official community.  To make your game more attractive,
// you can upload a GIF animation that demostrates the game to the users via BotFather (see
// Lumberjack for example).  A game message will also display high scores for the current chat.
// Use setGameScore to post high scores to the chat with the game, add the edit_message
// parameter to automatically update the message with the current scoreboard.  Use
// getGameHighScores to get data for in-game high score tables.  You can also add an extra
// sharing button for users to share their best score to different chats.  For examples of what
// can be done using this new stuff, check the @gamebot and @gamee bots.    Use this method to
// send a game. On success, the sent Message is returned.
func (a *TelegramApi) SendGame(request *SendGameRequest) (*SendGameResponse, error) {
	apiResponse, err := queryAndUnmarshal[*Message](a.bot, "SendGame", request)
	if err != nil {
		return nil, err
	}
	return &SendGameResponse{Result: apiResponse.Result}, nil
}

// Use this method to set the score of the specified user in a game message. On success, if the
// message is not an inline message, the Message is returned, otherwise True is returned.
// Returns an error, if the new score is not greater than the user's current score in the chat
// and force is False .
func (a *TelegramApi) SetGameScore(request *SetGameScoreRequest) (*SetGameScoreResponse, error) {
	_, err := queryAndUnmarshal[interface{}](a.bot, "SetGameScore", request)
	if err != nil {
		return nil, err
	}
	return &SetGameScoreResponse{}, nil
}

// Use this method to get data for high score tables. Will return the score of the specified
// user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
// This method will currently return scores for the target user, plus two of their closest
// neighbors on each side. Will also return the top three users if the user and their neighbors
// are not among them. Please note that this behavior is subject to change.
func (a *TelegramApi) GetGameHighScores(request *GetGameHighScoresRequest) (*GetGameHighScoresResponse, error) {
	apiResponse, err := queryAndUnmarshal[[]*GameHighScore](a.bot, "GetGameHighScores", request)
	if err != nil {
		return nil, err
	}
	return &GetGameHighScoresResponse{Result: apiResponse.Result}, nil
}
