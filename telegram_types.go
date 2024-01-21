package tgbot

// Telegram Bot API                      Twitter   Home  FAQ  Apps  API  Protocol  Schema
// Telegram Bots Telegram Bot API  Telegram Bot API    The Bot API is an HTTP-based interface
// created for developers keen on building bots for Telegram. To learn how to create and set up a
// bot, please consult our Introduction to Bots and Bot FAQ .    This object represents an
// incoming update. At most one of the optional parameters can be present in any given update.
type Update struct {

	// The update's unique identifier. Update identifiers start from a certain positive number and
	// increase sequentially. This identifier becomes especially handy if you're using webhooks, since
	// it allows you to ignore repeated updates or to restore the correct update sequence, should they
	// get out of order. If there are no new updates for at least a week, then identifier of the next
	// update will be chosen randomly instead of sequentially.
	UpdateID int64 `json:"update_id"`

	// Optional. New incoming message of any kind - text, photo, sticker, etc.
	Message *Message `json:"message"`

	// Optional. New version of a message that is known to the bot and was edited. This update may at
	// times be triggered by changes to message fields that are either unavailable or not actively
	// used by your bot.
	EditedMessage *Message `json:"edited_message"`

	// Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post"`

	// Optional. New version of a channel post that is known to the bot and was edited. This update
	// may at times be triggered by changes to message fields that are either unavailable or not
	// actively used by your bot.
	EditedChannelPost *Message `json:"edited_channel_post"`

	// Optional. A reaction to a message was changed by a user. The bot must be an administrator in
	// the chat and must explicitly specify "message_reaction" in the list of allowed_updates to
	// receive these updates. The update isn't received for reactions set by bots.
	MessageReaction *MessageReactionUpdated `json:"message_reaction"`

	// Optional. Reactions to a message with anonymous reactions were changed. The bot must be an
	// administrator in the chat and must explicitly specify "message_reaction_count" in the list of
	// allowed_updates to receive these updates. The updates are grouped and can be sent with delay up
	// to a few minutes.
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count"`

	// Optional. New incoming inline query
	InlineQuery *InlineQuery `json:"inline_query"`

	// Optional. The result of an inline query that was chosen by a user and sent to their chat
	// partner. Please see our documentation on the feedback collecting for details on how to enable
	// these updates for your bot.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`

	// Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query"`

	// Optional. New incoming shipping query. Only for invoices with flexible price
	ShippingQuery *ShippingQuery `json:"shipping_query"`

	// Optional. New incoming pre-checkout query. Contains full information about checkout
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query"`

	// Optional. New poll state. Bots receive only updates about manually stopped polls and polls,
	// which are sent by the bot
	Poll *Poll `json:"poll"`

	// Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in
	// polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer"`

	// Optional. The bot's chat member status was updated in a chat. For private chats, this update is
	// received only when the bot is blocked or unblocked by the user.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member"`

	// Optional. A chat member's status was updated in a chat. The bot must be an administrator in the
	// chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these
	// updates.
	ChatMember *ChatMemberUpdated `json:"chat_member"`

	// Optional. A request to join the chat has been sent. The bot must have the can_invite_users
	// administrator right in the chat to receive these updates.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request"`

	// Optional. A chat boost was added or changed. The bot must be an administrator in the chat to
	// receive these updates.
	ChatBoost *ChatBoostUpdated `json:"chat_boost"`

	// Optional. A boost was removed from a chat. The bot must be an administrator in the chat to
	// receive these updates.
	RemovedChatBoost *ChatBoostRemoved `json:"removed_chat_boost"`
}

// Describes the current status of a webhook.
type WebhookInfo struct {

	// Webhook URL, may be empty if webhook is not set up
	URL string `json:"url"`

	// True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`

	// Number of updates awaiting delivery
	PendingUpdateCount int64 `json:"pending_update_count"`

	// Optional. Currently used webhook IP address
	IPAddress string `json:"ip_address"`

	// Optional. Unix time for the most recent error that happened when trying to deliver an update
	// via webhook
	LastErrorDate int64 `json:"last_error_date"`

	// Optional. Error message in human-readable format for the most recent error that happened when
	// trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message"`

	// Optional. Unix time of the most recent error that happened when trying to synchronize available
	// updates with Telegram datacenters
	LastSynchronizationErrorDate int64 `json:"last_synchronization_error_date"`

	// Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for
	// update delivery
	MaxConnections int64 `json:"max_connections"`

	// Optional. A list of update types the bot is subscribed to. Defaults to all update types except
	// chat_member
	AllowedUpdates []string `json:"allowed_updates"`
}

// All types used in the Bot API responses are represented as JSON-objects.  It is safe to use
// 32-bit signed integers for storing all Integer fields unless otherwise noted.   Optional fields
// may be not returned when irrelevant.    This object represents a Telegram user or bot.
type User struct {

	// Unique identifier for this user or bot. This number may have more than 32 significant bits and
	// some programming languages may have difficulty/silent defects in interpreting it. But it has at
	// most 52 significant bits, so a 64-bit integer or double-precision float type are safe for
	// storing this identifier.
	ID int64 `json:"id"`

	// True, if this user is a bot
	IsBot bool `json:"is_bot"`

	// User's or bot's first name
	FirstName string `json:"first_name"`

	// Optional. User's or bot's last name
	LastName string `json:"last_name"`

	// Optional. User's or bot's username
	Username string `json:"username"`

	// Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code"`

	// Optional. True, if this user is a Telegram Premium user
	IsPremium bool `json:"is_premium"`

	// Optional. True, if this user added the bot to the attachment menu
	AddedToAttachmentMenu bool `json:"added_to_attachment_menu"`

	// Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanJoinGroups bool `json:"can_join_groups"`

	// Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`

	// Optional. True, if the bot supports inline queries. Returned only in getMe.
	SupportsInlineQueries bool `json:"supports_inline_queries"`
}

// This object represents a chat.
type Chat struct {

	// Unique identifier for this chat. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most
	// 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for
	// storing this identifier.
	ID int64 `json:"id"`

	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`

	// Optional. Title, for supergroups, channels and group chats
	Title string `json:"title"`

	// Optional. Username, for private chats, supergroups and channels if available
	Username string `json:"username"`

	// Optional. First name of the other party in a private chat
	FirstName string `json:"first_name"`

	// Optional. Last name of the other party in a private chat
	LastName string `json:"last_name"`

	// Optional. True, if the supergroup chat is a forum (has topics enabled)
	IsForum bool `json:"is_forum"`

	// Optional. Chat photo. Returned only in getChat.
	Photo *ChatPhoto `json:"photo"`

	// Optional. If non-empty, the list of all active chat usernames; for private chats, supergroups
	// and channels. Returned only in getChat.
	ActiveUsernames []string `json:"active_usernames"`

	// Optional. List of available reactions allowed in the chat. If omitted, then all emoji reactions
	// are allowed. Returned only in getChat.
	AvailableReactions []*ReactionType `json:"available_reactions"`

	// Optional. Identifier of the accent color for the chat name and backgrounds of the chat photo,
	// reply header, and link preview. See accent colors for more details. Returned only in getChat.
	// Always returned in getChat.
	AccentColorID int64 `json:"accent_color_id"`

	// Optional. Custom emoji identifier of emoji chosen by the chat for the reply header and link
	// preview background. Returned only in getChat.
	BackgroundCustomEmojiID string `json:"background_custom_emoji_id"`

	// Optional. Identifier of the accent color for the chat's profile background. See profile accent
	// colors for more details. Returned only in getChat.
	ProfileAccentColorID int64 `json:"profile_accent_color_id"`

	// Optional. Custom emoji identifier of the emoji chosen by the chat for its profile background.
	// Returned only in getChat.
	ProfileBackgroundCustomEmojiID string `json:"profile_background_custom_emoji_id"`

	// Optional. Custom emoji identifier of the emoji status of the chat or the other party in a
	// private chat. Returned only in getChat.
	EmojiStatusCustomEmojiID string `json:"emoji_status_custom_emoji_id"`

	// Optional. Expiration date of the emoji status of the chat or the other party in a private chat,
	// in Unix time, if any. Returned only in getChat.
	EmojiStatusExpirationDate int64 `json:"emoji_status_expiration_date"`

	// Optional. Bio of the other party in a private chat. Returned only in getChat.
	Bio string `json:"bio"`

	// Optional. True, if privacy settings of the other party in the private chat allows to use
	// tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
	HasPrivateForwards bool `json:"has_private_forwards"`

	// Optional. True, if the privacy settings of the other party restrict sending voice and video
	// note messages in the private chat. Returned only in getChat.
	HasRestrictedVoiceAndVideoMessages bool `json:"has_restricted_voice_and_video_messages"`

	// Optional. True, if users need to join the supergroup before they can send messages. Returned
	// only in getChat.
	JoinToSendMessages bool `json:"join_to_send_messages"`

	// Optional. True, if all users directly joining the supergroup need to be approved by supergroup
	// administrators. Returned only in getChat.
	JoinByRequest bool `json:"join_by_request"`

	// Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	Description string `json:"description"`

	// Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in
	// getChat.
	InviteLink string `json:"invite_link"`

	// Optional. The most recent pinned message (by sending date). Returned only in getChat.
	PinnedMessage *Message `json:"pinned_message"`

	// Optional. Default chat member permissions, for groups and supergroups. Returned only in
	// getChat.
	Permissions *ChatPermissions `json:"permissions"`

	// Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each
	// unpriviledged user; in seconds. Returned only in getChat.
	SlowModeDelay int64 `json:"slow_mode_delay"`

	// Optional. The time after which all messages sent to the chat will be automatically deleted; in
	// seconds. Returned only in getChat.
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time"`

	// Optional. True, if aggressive anti-spam checks are enabled in the supergroup. The field is only
	// available to chat administrators. Returned only in getChat.
	HasAggressiveAntiSpamEnabled bool `json:"has_aggressive_anti_spam_enabled"`

	// Optional. True, if non-administrators can only get the list of bots and administrators in the
	// chat. Returned only in getChat.
	HasHiddenMembers bool `json:"has_hidden_members"`

	// Optional. True, if messages from the chat can't be forwarded to other chats. Returned only in
	// getChat.
	HasProtectedContent bool `json:"has_protected_content"`

	// Optional. True, if new chat members will have access to old messages; available only to chat
	// administrators. Returned only in getChat.
	HasVisibleHistory bool `json:"has_visible_history"`

	// Optional. For supergroups, name of group sticker set. Returned only in getChat.
	StickerSetName string `json:"sticker_set_name"`

	// Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	CanSetStickerSet bool `json:"can_set_sticker_set"`

	// Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a
	// channel and vice versa; for supergroups and channel chats. This identifier may be greater than
	// 32 bits and some programming languages may have difficulty/silent defects in interpreting it.
	// But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are
	// safe for storing this identifier. Returned only in getChat.
	LinkedChatID int64 `json:"linked_chat_id"`

	// Optional. For supergroups, the location to which the supergroup is connected. Returned only in
	// getChat.
	Location *ChatLocation `json:"location"`
}

// This object represents a message.
type Message struct {

	// Unique message identifier inside this chat
	MessageID int64 `json:"message_id"`

	// Optional. Unique identifier of a message thread to which the message belongs; for supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Optional. Sender of the message; empty for messages sent to channels. For backward
	// compatibility, the field contains a fake sender user in non-channel chats, if the message was
	// sent on behalf of a chat.
	From *User `json:"from"`

	// Optional. Sender of the message, sent on behalf of a chat. For example, the channel itself for
	// channel posts, the supergroup itself for messages from anonymous group administrators, the
	// linked channel for messages automatically forwarded to the discussion group. For backward
	// compatibility, the field from contains a fake sender user in non-channel chats, if the message
	// was sent on behalf of a chat.
	SenderChat *Chat `json:"sender_chat"`

	// Date the message was sent in Unix time. It is always a positive number, representing a valid
	// date.
	Date int64 `json:"date"`

	// Chat the message belongs to
	Chat *Chat `json:"chat"`

	// Optional. Information about the original message for forwarded messages
	ForwardOrigin *MessageOrigin `json:"forward_origin"`

	// Optional. True, if the message is sent to a forum topic
	IsTopicMessage bool `json:"is_topic_message"`

	// Optional. True, if the message is a channel post that was automatically forwarded to the
	// connected discussion group
	IsAutomaticForward bool `json:"is_automatic_forward"`

	// Optional. For replies in the same chat and message thread, the original message. Note that the
	// Message object in this field will not contain further reply_to_message fields even if it itself
	// is a reply.
	ReplyToMessage *Message `json:"reply_to_message"`

	// Optional. Information about the message that is being replied to, which may come from another
	// chat or forum topic
	ExternalReply *ExternalReplyInfo `json:"external_reply"`

	// Optional. For replies that quote part of the original message, the quoted part of the message
	Quote *TextQuote `json:"quote"`

	// Optional. Bot through which the message was sent
	ViaBot *User `json:"via_bot"`

	// Optional. Date the message was last edited in Unix time
	EditDate int64 `json:"edit_date"`

	// Optional. True, if the message can't be forwarded
	HasProtectedContent bool `json:"has_protected_content"`

	// Optional. The unique identifier of a media message group this message belongs to
	MediaGroupID string `json:"media_group_id"`

	// Optional. Signature of the post author for messages in channels, or the custom title of an
	// anonymous group administrator
	AuthorSignature string `json:"author_signature"`

	// Optional. For text messages, the actual UTF-8 text of the message
	Text string `json:"text"`

	// Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that
	// appear in the text
	Entities []*MessageEntity `json:"entities"`

	// Optional. Options used for link preview generation for the message, if it is a text message and
	// link preview options were changed
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"`

	// Optional. Message is an animation, information about the animation. For backward compatibility,
	// when this field is set, the document field will also be set
	Animation *Animation `json:"animation"`

	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio"`

	// Optional. Message is a general file, information about the file
	Document *Document `json:"document"`

	// Optional. Message is a photo, available sizes of the photo
	Photo []*PhotoSize `json:"photo"`

	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker"`

	// Optional. Message is a forwarded story
	Story *Story `json:"story"`

	// Optional. Message is a video, information about the video
	Video *Video `json:"video"`

	// Optional. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note"`

	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice"`

	// Optional. Caption for the animation, audio, document, photo, video or voice
	Caption string `json:"caption"`

	// Optional. For messages with a caption, special entities like usernames, URLs, bot commands,
	// etc. that appear in the caption
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler"`

	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact"`

	// Optional. Message is a dice with random value
	Dice *Dice `json:"dice"`

	// Optional. Message is a game, information about the game. More about games »
	Game *Game `json:"game"`

	// Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll"`

	// Optional. Message is a venue, information about the venue. For backward compatibility, when
	// this field is set, the location field will also be set
	Venue *Venue `json:"venue"`

	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location"`

	// Optional. New members that were added to the group or supergroup and information about them
	// (the bot itself may be one of these members)
	NewChatMembers []*User `json:"new_chat_members"`

	// Optional. A member was removed from the group, information about them (this member may be the
	// bot itself)
	LeftChatMember *User `json:"left_chat_member"`

	// Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title"`

	// Optional. A chat photo was change to this value
	NewChatPhoto []*PhotoSize `json:"new_chat_photo"`

	// Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo"`

	// Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created"`

	// Optional. Service message: the supergroup has been created. This field can't be received in a
	// message coming through updates, because bot can't be a member of a supergroup when it is
	// created. It can only be found in reply_to_message if someone replies to a very first message in
	// a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created"`

	// Optional. Service message: the channel has been created. This field can't be received in a
	// message coming through updates, because bot can't be a member of a channel when it is created.
	// It can only be found in reply_to_message if someone replies to a very first message in a
	// channel.
	ChannelChatCreated bool `json:"channel_chat_created"`

	// Optional. Service message: auto-delete timer settings changed in the chat
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`

	// Optional. The group has been migrated to a supergroup with the specified identifier. This
	// number may have more than 32 significant bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id"`

	// Optional. The supergroup has been migrated from a group with the specified identifier. This
	// number may have more than 32 significant bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatID int64 `json:"migrate_from_chat_id"`

	// Optional. Specified message was pinned. Note that the Message object in this field will not
	// contain further reply_to_message fields even if it itself is a reply.
	PinnedMessage *MaybeInaccessibleMessage `json:"pinned_message"`

	// Optional. Message is an invoice for a payment, information about the invoice. More about
	// payments »
	Invoice *Invoice `json:"invoice"`

	// Optional. Message is a service message about a successful payment, information about the
	// payment. More about payments »
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment"`

	// Optional. Service message: users were shared with the bot
	UsersShared *UsersShared `json:"users_shared"`

	// Optional. Service message: a chat was shared with the bot
	ChatShared *ChatShared `json:"chat_shared"`

	// Optional. The domain name of the website on which the user has logged in. More about Telegram
	// Login »
	ConnectedWebsite string `json:"connected_website"`

	// Optional. Service message: the user allowed the bot to write messages after adding it to the
	// attachment or side menu, launching a Web App from a link, or accepting an explicit request from
	// a Web App sent by the method requestWriteAccess
	WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed"`

	// Optional. Telegram Passport data
	PassportData *PassportData `json:"passport_data"`

	// Optional. Service message. A user in the chat triggered another user's proximity alert while
	// sharing Live Location.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered"`

	// Optional. Service message: forum topic created
	ForumTopicCreated *ForumTopicCreated `json:"forum_topic_created"`

	// Optional. Service message: forum topic edited
	ForumTopicEdited *ForumTopicEdited `json:"forum_topic_edited"`

	// Optional. Service message: forum topic closed
	ForumTopicClosed *ForumTopicClosed `json:"forum_topic_closed"`

	// Optional. Service message: forum topic reopened
	ForumTopicReopened *ForumTopicReopened `json:"forum_topic_reopened"`

	// Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicHidden *GeneralForumTopicHidden `json:"general_forum_topic_hidden"`

	// Optional. Service message: the 'General' forum topic unhidden
	GeneralForumTopicUnhidden *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden"`

	// Optional. Service message: a scheduled giveaway was created
	GiveawayCreated *GiveawayCreated `json:"giveaway_created"`

	// Optional. The message is a scheduled giveaway message
	Giveaway *Giveaway `json:"giveaway"`

	// Optional. A giveaway with public winners was completed
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners"`

	// Optional. Service message: a giveaway without public winners was completed
	GiveawayCompleted *GiveawayCompleted `json:"giveaway_completed"`

	// Optional. Service message: video chat scheduled
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled"`

	// Optional. Service message: video chat started
	VideoChatStarted *VideoChatStarted `json:"video_chat_started"`

	// Optional. Service message: video chat ended
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended"`

	// Optional. Service message: new participants invited to a video chat
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited"`

	// Optional. Service message: data sent by a Web App
	WebAppData *WebAppData `json:"web_app_data"`

	// Optional. Inline keyboard attached to the message. login_url buttons are represented as
	// ordinary url buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// This object represents a unique message identifier.
type MessageId struct {

	// Unique message identifier
	MessageID int64 `json:"message_id"`
}

// This object describes a message that was deleted or is otherwise inaccessible to the bot.
type InaccessibleMessage struct {

	// Chat the message belonged to
	Chat *Chat `json:"chat"`

	// Unique message identifier inside the chat
	MessageID int64 `json:"message_id"`

	// Always 0. The field can be used to differentiate regular and inaccessible messages.
	Date int64 `json:"date"`
}

// This object describes a message that can be inaccessible to the bot. It can be one of
// Message  InaccessibleMessage
type MaybeInaccessibleMessage struct {
}

// This object represents one special entity in a text message. For example, hashtags, usernames,
// URLs, etc.
type MessageEntity struct {

	// Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag), “cashtag”
	// ($USD), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email”
	// (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic”
	// (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler”
	// (spoiler message), “blockquote” (block quotation), “code” (monowidth string), “pre” (monowidth
	// block), “text_link” (for clickable text URLs), “text_mention” (for users without usernames),
	// “custom_emoji” (for inline custom emoji stickers)
	Type string `json:"type"`

	// Offset in UTF-16 code units to the start of the entity
	Offset int64 `json:"offset"`

	// Length of the entity in UTF-16 code units
	Length int64 `json:"length"`

	// Optional. For “text_link” only, URL that will be opened after user taps on the text
	URL string `json:"url"`

	// Optional. For “text_mention” only, the mentioned user
	User *User `json:"user"`

	// Optional. For “pre” only, the programming language of the entity text
	Language string `json:"language"`

	// Optional. For “custom_emoji” only, unique identifier of the custom emoji. Use
	// getCustomEmojiStickers to get full information about the sticker
	CustomEmojiID string `json:"custom_emoji_id"`
}

// This object contains information about the quoted part of a message that is replied to by the
// given message.
type TextQuote struct {

	// Text of the quoted part of a message that is replied to by the given message
	Text string `json:"text"`

	// Optional. Special entities that appear in the quote. Currently, only bold, italic, underline,
	// strikethrough, spoiler, and custom_emoji entities are kept in quotes.
	Entities []*MessageEntity `json:"entities"`

	// Approximate quote position in the original message in UTF-16 code units as specified by the
	// sender
	Position int64 `json:"position"`

	// Optional. True, if the quote was chosen manually by the message sender. Otherwise, the quote
	// was added automatically by the server.
	IsManual bool `json:"is_manual"`
}

// This object contains information about a message that is being replied to, which may come
// from another chat or forum topic.
type ExternalReplyInfo struct {

	// Origin of the message replied to by the given message
	Origin *MessageOrigin `json:"origin"`

	// Optional. Chat the original message belongs to. Available only if the chat is a supergroup or a
	// channel.
	Chat *Chat `json:"chat"`

	// Optional. Unique message identifier inside the original chat. Available only if the original
	// chat is a supergroup or a channel.
	MessageID int64 `json:"message_id"`

	// Optional. Options used for link preview generation for the original message, if it is a text
	// message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"`

	// Optional. Message is an animation, information about the animation
	Animation *Animation `json:"animation"`

	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio"`

	// Optional. Message is a general file, information about the file
	Document *Document `json:"document"`

	// Optional. Message is a photo, available sizes of the photo
	Photo []*PhotoSize `json:"photo"`

	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker"`

	// Optional. Message is a forwarded story
	Story *Story `json:"story"`

	// Optional. Message is a video, information about the video
	Video *Video `json:"video"`

	// Optional. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note"`

	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice"`

	// Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler"`

	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact"`

	// Optional. Message is a dice with random value
	Dice *Dice `json:"dice"`

	// Optional. Message is a game, information about the game. More about games »
	Game *Game `json:"game"`

	// Optional. Message is a scheduled giveaway, information about the giveaway
	Giveaway *Giveaway `json:"giveaway"`

	// Optional. A giveaway with public winners was completed
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners"`

	// Optional. Message is an invoice for a payment, information about the invoice. More about
	// payments »
	Invoice *Invoice `json:"invoice"`

	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location"`

	// Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll"`

	// Optional. Message is a venue, information about the venue
	Venue *Venue `json:"venue"`
}

// Describes reply parameters for the message that is being sent.
type ReplyParameters struct {

	// Identifier of the message that will be replied to in the current chat, or in the chat chat_id
	// if it is specified
	MessageID int64 `json:"message_id"`

	// Optional. If the message to be replied to is from a different chat, unique identifier for the
	// chat or username of the channel (in the format @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Optional. Pass True if the message should be sent even if the specified message to be replied
	// to is not found; can be used only for replies in the same chat and forum topic.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`

	// Optional. Quoted part of the message to be replied to; 0-1024 characters after entities
	// parsing. The quote must be an exact substring of the message to be replied to, including bold,
	// italic, underline, strikethrough, spoiler, and custom_emoji entities. The message will fail to
	// send if the quote isn't found in the original message.
	Quote string `json:"quote"`

	// Optional. Mode for parsing entities in the quote. See formatting options for more details.
	QuoteParseMode string `json:"quote_parse_mode"`

	// Optional. A JSON-serialized list of special entities that appear in the quote. It can be
	// specified instead of quote_parse_mode.
	QuoteEntities []*MessageEntity `json:"quote_entities"`

	// Optional. Position of the quote in the original message in UTF-16 code units
	QuotePosition int64 `json:"quote_position"`
}

// The message was originally sent by a known user.
type MessageOriginUser struct {

	// Type of the message origin, always “user”
	Type string `json:"type"`

	// Date the message was sent originally in Unix time
	Date int64 `json:"date"`

	// User that sent the message originally
	SenderUser *User `json:"sender_user"`
}

// The message was originally sent by an unknown user.
type MessageOriginHiddenUser struct {

	// Type of the message origin, always “hidden_user”
	Type string `json:"type"`

	// Date the message was sent originally in Unix time
	Date int64 `json:"date"`

	// Name of the user that sent the message originally
	SenderUserName string `json:"sender_user_name"`
}

// The message was originally sent on behalf of a chat to a group chat.
type MessageOriginChat struct {

	// Type of the message origin, always “chat”
	Type string `json:"type"`

	// Date the message was sent originally in Unix time
	Date int64 `json:"date"`

	// Chat that sent the message originally
	SenderChat *Chat `json:"sender_chat"`

	// Optional. For messages originally sent by an anonymous chat administrator, original message
	// author signature
	AuthorSignature string `json:"author_signature"`
}

// The message was originally sent to a channel chat.
type MessageOriginChannel struct {

	// Type of the message origin, always “channel”
	Type string `json:"type"`

	// Date the message was sent originally in Unix time
	Date int64 `json:"date"`

	// Channel chat to which the message was originally sent
	Chat *Chat `json:"chat"`

	// Unique message identifier inside the chat
	MessageID int64 `json:"message_id"`

	// Optional. Signature of the original post author
	AuthorSignature string `json:"author_signature"`
}

// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {

	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different
	// bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Photo width
	Width int64 `json:"width"`

	// Photo height
	Height int64 `json:"height"`

	// Optional. File size in bytes
	FileSize int64 `json:"file_size"`
}

// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {

	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different
	// bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Video width as defined by sender
	Width int64 `json:"width"`

	// Video height as defined by sender
	Height int64 `json:"height"`

	// Duration of the video in seconds as defined by sender
	Duration int64 `json:"duration"`

	// Optional. Animation thumbnail as defined by sender
	Thumbnail *PhotoSize `json:"thumbnail"`

	// Optional. Original animation filename as defined by sender
	FileName string `json:"file_name"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size"`
}

// This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {

	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different
	// bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by sender
	Duration int64 `json:"duration"`

	// Optional. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer"`

	// Optional. Title of the audio as defined by sender or by audio tags
	Title string `json:"title"`

	// Optional. Original filename as defined by sender
	FileName string `json:"file_name"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size"`

	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumbnail *PhotoSize `json:"thumbnail"`
}

// This object represents a general file (as opposed to photos , voice messages and audio files
// ).
type Document struct {

	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different
	// bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Optional. Document thumbnail as defined by sender
	Thumbnail *PhotoSize `json:"thumbnail"`

	// Optional. Original filename as defined by sender
	FileName string `json:"file_name"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size"`
}

// This object represents a message about a forwarded story in the chat. Currently holds no
// information.
type Story struct {
}

// This object represents a video file.
type Video struct {

	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different
	// bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Video width as defined by sender
	Width int64 `json:"width"`

	// Video height as defined by sender
	Height int64 `json:"height"`

	// Duration of the video in seconds as defined by sender
	Duration int64 `json:"duration"`

	// Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail"`

	// Optional. Original filename as defined by sender
	FileName string `json:"file_name"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size"`
}

// This object represents a video message (available in Telegram apps as of v.4.0 ).
type VideoNote struct {

	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different
	// bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Video width and height (diameter of the video message) as defined by sender
	Length int64 `json:"length"`

	// Duration of the video in seconds as defined by sender
	Duration int64 `json:"duration"`

	// Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail"`

	// Optional. File size in bytes
	FileSize int64 `json:"file_size"`
}

// This object represents a voice note.
type Voice struct {

	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different
	// bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by sender
	Duration int64 `json:"duration"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size"`
}

// This object represents a phone contact.
type Contact struct {

	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name"`

	// Optional. Contact's user identifier in Telegram. This number may have more than 32 significant
	// bits and some programming languages may have difficulty/silent defects in interpreting it. But
	// it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe
	// for storing this identifier.
	UserID int64 `json:"user_id"`

	// Optional. Additional data about the contact in the form of a vCard
	Vcard string `json:"vcard"`
}

// This object represents an animated emoji that displays a random value.
type Dice struct {

	// Emoji on which the dice throw animation is based
	Emoji string `json:"emoji"`

	// Value of the dice, 1-6 for “”, “” and “” base emoji, 1-5 for “” and “” base emoji, 1-64 for “”
	// base emoji
	Value int64 `json:"value"`
}

// This object contains information about one answer option in a poll.
type PollOption struct {

	// Option text, 1-100 characters
	Text string `json:"text"`

	// Number of users that voted for this option
	VoterCount int64 `json:"voter_count"`
}

// This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {

	// Unique poll identifier
	PollID string `json:"poll_id"`

	// Optional. The chat that changed the answer to the poll, if the voter is anonymous
	VoterChat *Chat `json:"voter_chat"`

	// Optional. The user that changed the answer to the poll, if the voter isn't anonymous
	User *User `json:"user"`

	// 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
	OptionIds []int64 `json:"option_ids"`
}

// This object contains information about a poll.
type Poll struct {

	// Unique poll identifier
	ID string `json:"id"`

	// Poll question, 1-300 characters
	Question string `json:"question"`

	// List of poll options
	Options []*PollOption `json:"options"`

	// Total number of users that voted in the poll
	TotalVoterCount int64 `json:"total_voter_count"`

	// True, if the poll is closed
	IsClosed bool `json:"is_closed"`

	// True, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous"`

	// Poll type, currently can be “regular” or “quiz”
	Type string `json:"type"`

	// True, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`

	// Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz
	// mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the
	// bot.
	CorrectOptionID int64 `json:"correct_option_id"`

	// Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon
	// in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation"`

	// Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the
	// explanation
	ExplanationEntities []*MessageEntity `json:"explanation_entities"`

	// Optional. Amount of time in seconds the poll will be active after creation
	OpenPeriod int64 `json:"open_period"`

	// Optional. Point in time (Unix timestamp) when the poll will be automatically closed
	CloseDate int64 `json:"close_date"`
}

// This object represents a point on the map.
type Location struct {

	// Longitude as defined by sender
	Longitude float32 `json:"longitude"`

	// Latitude as defined by sender
	Latitude float32 `json:"latitude"`

	// Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy"`

	// Optional. Time relative to the message sending date, during which the location can be updated;
	// in seconds. For active live locations only.
	LivePeriod int64 `json:"live_period"`

	// Optional. The direction in which user is moving, in degrees; 1-360. For active live locations
	// only.
	Heading int64 `json:"heading"`

	// Optional. The maximum distance for proximity alerts about approaching another chat member, in
	// meters. For sent live locations only.
	ProximityAlertRadius int64 `json:"proximity_alert_radius"`
}

// This object represents a venue.
type Venue struct {

	// Venue location. Can't be a live location
	Location *Location `json:"location"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Optional. Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id"`

	// Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type"`

	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id"`

	// Optional. Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type"`
}

// Describes data sent from a Web App to the bot.
type WebAppData struct {

	// The data. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data"`

	// Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad
	// client can send arbitrary data in this field.
	ButtonText string `json:"button_text"`
}

// This object represents the content of a service message, sent whenever a user in the chat
// triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {

	// User that triggered the alert
	Traveler *User `json:"traveler"`

	// User that set the alert
	Watcher *User `json:"watcher"`

	// The distance between the users
	Distance int64 `json:"distance"`
}

// This object represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {

	// New auto-delete time for messages in the chat; in seconds
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time"`
}

// This object represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {

	// Name of the topic
	Name string `json:"name"`

	// Color of the topic icon in RGB format
	IconColor int64 `json:"icon_color"`

	// Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
}

// This object represents a service message about a forum topic closed in the chat. Currently
// holds no information.
type ForumTopicClosed struct {
}

// This object represents a service message about an edited forum topic.
type ForumTopicEdited struct {

	// Optional. New name of the topic, if it was edited
	Name string `json:"name"`

	// Optional. New identifier of the custom emoji shown as the topic icon, if it was edited; an
	// empty string if the icon was removed
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
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
	RequestID int64 `json:"request_id"`

	// Identifiers of the shared users. These numbers may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting them. But they have at
	// most 52 significant bits, so 64-bit integers or double-precision float types are safe for
	// storing these identifiers. The bot may not have access to the users and could be unable to use
	// these identifiers, unless the users are already known to the bot by some other means.
	UserIds []int64 `json:"user_ids"`
}

// This object contains information about the chat whose identifier was shared with the bot
// using a KeyboardButtonRequestChat button.
type ChatShared struct {

	// Identifier of the request
	RequestID int64 `json:"request_id"`

	// Identifier of the shared chat. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most
	// 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing
	// this identifier. The bot may not have access to the chat and could be unable to use this
	// identifier, unless the chat is already known to the bot by some other means.
	ChatID int64 `json:"chat_id"`
}

// This object represents a service message about a user allowing a bot to write messages after
// adding it to the attachment menu, launching a Web App from a link, or accepting an explicit
// request from a Web App sent by the method requestWriteAccess .
type WriteAccessAllowed struct {

	// Optional. True, if the access was granted after the user accepted an explicit request from a
	// Web App sent by the method requestWriteAccess
	FromRequest bool `json:"from_request"`

	// Optional. Name of the Web App, if the access was granted when the Web App was launched from a
	// link
	WebAppName string `json:"web_app_name"`

	// Optional. True, if the access was granted when the bot was added to the attachment or side menu
	FromAttachmentMenu bool `json:"from_attachment_menu"`
}

// This object represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {

	// Point in time (Unix timestamp) when the video chat is supposed to be started by a chat
	// administrator
	StartDate int64 `json:"start_date"`
}

// This object represents a service message about a video chat started in the chat. Currently
// holds no information.
type VideoChatStarted struct {
}

// This object represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {

	// Video chat duration in seconds
	Duration int64 `json:"duration"`
}

// This object represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {

	// New members that were invited to the video chat
	Users []*User `json:"users"`
}

// This object represents a service message about the creation of a scheduled giveaway.
// Currently holds no information.
type GiveawayCreated struct {
}

// This object represents a message about a scheduled giveaway.
type Giveaway struct {

	// The list of chats which the user must join to participate in the giveaway
	Chats []*Chat `json:"chats"`

	// Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnersSelectionDate int64 `json:"winners_selection_date"`

	// The number of users which are supposed to be selected as winners of the giveaway
	WinnerCount int64 `json:"winner_count"`

	// Optional. True, if only users who join the chats after the giveaway started should be eligible
	// to win
	OnlyNewMembers bool `json:"only_new_members"`

	// Optional. True, if the list of giveaway winners will be visible to everyone
	HasPublicWinners bool `json:"has_public_winners"`

	// Optional. Description of additional giveaway prize
	PrizeDescription string `json:"prize_description"`

	// Optional. A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from
	// which eligible users for the giveaway must come. If empty, then all users can participate in
	// the giveaway. Users with a phone number that was bought on Fragment can always participate in
	// giveaways.
	CountryCodes []string `json:"country_codes"`

	// Optional. The number of months the Telegram Premium subscription won from the giveaway will be
	// active for
	PremiumSubscriptionMonthCount int64 `json:"premium_subscription_month_count"`
}

// This object represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {

	// The chat that created the giveaway
	Chat *Chat `json:"chat"`

	// Identifier of the messsage with the giveaway in the chat
	GiveawayMessageID int64 `json:"giveaway_message_id"`

	// Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnersSelectionDate int64 `json:"winners_selection_date"`

	// Total number of winners in the giveaway
	WinnerCount int64 `json:"winner_count"`

	// List of up to 100 winners of the giveaway
	Winners []*User `json:"winners"`

	// Optional. The number of other chats the user had to join in order to be eligible for the
	// giveaway
	AdditionalChatCount int64 `json:"additional_chat_count"`

	// Optional. The number of months the Telegram Premium subscription won from the giveaway will be
	// active for
	PremiumSubscriptionMonthCount int64 `json:"premium_subscription_month_count"`

	// Optional. Number of undistributed prizes
	UnclaimedPrizeCount int64 `json:"unclaimed_prize_count"`

	// Optional. True, if only users who had joined the chats after the giveaway started were eligible
	// to win
	OnlyNewMembers bool `json:"only_new_members"`

	// Optional. True, if the giveaway was canceled because the payment for it was refunded
	WasRefunded bool `json:"was_refunded"`

	// Optional. Description of additional giveaway prize
	PrizeDescription string `json:"prize_description"`
}

// This object represents a service message about the completion of a giveaway without public
// winners.
type GiveawayCompleted struct {

	// Number of winners in the giveaway
	WinnerCount int64 `json:"winner_count"`

	// Optional. Number of undistributed prizes
	UnclaimedPrizeCount int64 `json:"unclaimed_prize_count"`

	// Optional. Message with the giveaway that was completed, if it wasn't deleted
	GiveawayMessage *Message `json:"giveaway_message"`
}

// Describes the options used for link preview generation.
type LinkPreviewOptions struct {

	// Optional. True, if the link preview is disabled
	IsDisabled bool `json:"is_disabled"`

	// Optional. URL to use for the link preview. If empty, then the first URL found in the message
	// text will be used
	URL string `json:"url"`

	// Optional. True, if the media in the link preview is suppposed to be shrunk; ignored if the URL
	// isn't explicitly specified or media size change isn't supported for the preview
	PreferSmallMedia bool `json:"prefer_small_media"`

	// Optional. True, if the media in the link preview is suppposed to be enlarged; ignored if the
	// URL isn't explicitly specified or media size change isn't supported for the preview
	PreferLargeMedia bool `json:"prefer_large_media"`

	// Optional. True, if the link preview must be shown above the message text; otherwise, the link
	// preview will be shown below the message text
	ShowAboveText bool `json:"show_above_text"`
}

// This object represent a user's profile pictures.
type UserProfilePhotos struct {

	// Total number of profile pictures the target user has
	TotalCount int64 `json:"total_count"`

	// Requested profile pictures (in up to 4 sizes each)
	Photos [][]*PhotoSize `json:"photos"`
}

// This object represents a file ready to be downloaded. The file can be downloaded via the link
// https://api.telegram.org/file/bot<token>/<file_path> . It is guaranteed that the link will be
// valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile
// .   The maximum file size to download is 20 MB
type File struct {

	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different
	// bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize int64 `json:"file_size"`

	// Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	FilePath string `json:"file_path"`
}

// Describes a Web App .
type WebAppInfo struct {

	// An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web
	// Apps
	URL string `json:"url"`
}

// This object represents a custom keyboard with reply options (see Introduction to bots for
// details and examples).
type ReplyKeyboardMarkup struct {

	// Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard [][]*KeyboardButton `json:"keyboard"`

	// Optional. Requests clients to always show the keyboard when the regular keyboard is hidden.
	// Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard
	// icon.
	IsPersistent bool `json:"is_persistent"`

	// Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the
	// keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the
	// custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard"`

	// Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will
	// still be available, but clients will automatically display the usual letter-keyboard in the
	// chat - the user can press a special button in the input field to see the custom keyboard again.
	// Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard"`

	// Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64
	// characters
	InputFieldPlaceholder string `json:"input_field_placeholder"`

	// Optional. Use this parameter if you want to show the keyboard to specific users only. Targets:
	// 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a
	// reply to a message in the same chat and forum topic, sender of the original message.Example: A
	// user requests to change the bot's language, bot replies to the request with a keyboard to
	// select the new language. Other users in the group don't see the keyboard.
	Selective bool `json:"selective"`
}

// This object represents one button of the reply keyboard. For simple text buttons, String can
// be used instead of this object to specify the button text. The optional fields web_app ,
// request_users , request_chat , request_contact , request_location , and request_poll are
// mutually exclusive.
type KeyboardButton struct {

	// Text of the button. If none of the optional fields are used, it will be sent as a message when
	// the button is pressed
	Text string `json:"text"`

	// Optional. If specified, pressing the button will open a list of suitable users. Identifiers of
	// selected users will be sent to the bot in a “users_shared” service message. Available in
	// private chats only.
	RequestUsers *KeyboardButtonRequestUsers `json:"request_users"`

	// Optional. If specified, pressing the button will open a list of suitable chats. Tapping on a
	// chat will send its identifier to the bot in a “chat_shared” service message. Available in
	// private chats only.
	RequestChat *KeyboardButtonRequestChat `json:"request_chat"`

	// Optional. If True, the user's phone number will be sent as a contact when the button is
	// pressed. Available in private chats only.
	RequestContact bool `json:"request_contact"`

	// Optional. If True, the user's current location will be sent when the button is pressed.
	// Available in private chats only.
	RequestLocation bool `json:"request_location"`

	// Optional. If specified, the user will be asked to create a poll and send it to the bot when the
	// button is pressed. Available in private chats only.
	RequestPoll *KeyboardButtonPollType `json:"request_poll"`

	// Optional. If specified, the described Web App will be launched when the button is pressed. The
	// Web App will be able to send a “web_app_data” service message. Available in private chats only.
	WebApp *WebAppInfo `json:"web_app"`
}

// Note:  request_users and request_chat options will only work in Telegram versions released
// after 3 February, 2023. Older clients will display unsupported message .   This object defines
// the criteria used to request suitable users. The identifiers of the selected users will be
// shared with the bot when the corresponding button is pressed. More about requesting users »
type KeyboardButtonRequestUsers struct {

	// Signed 32-bit identifier of the request that will be received back in the UsersShared object.
	// Must be unique within the message
	RequestID int64 `json:"request_id"`

	// Optional. Pass True to request bots, pass False to request regular users. If not specified, no
	// additional restrictions are applied.
	UserIsBot bool `json:"user_is_bot"`

	// Optional. Pass True to request premium users, pass False to request non-premium users. If not
	// specified, no additional restrictions are applied.
	UserIsPremium bool `json:"user_is_premium"`

	// Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
	MaxQuantity int64 `json:"max_quantity"`
}

// This object defines the criteria used to request a suitable chat. The identifier of the
// selected chat will be shared with the bot when the corresponding button is pressed. More about
// requesting chats »
type KeyboardButtonRequestChat struct {

	// Signed 32-bit identifier of the request, which will be received back in the ChatShared object.
	// Must be unique within the message
	RequestID int64 `json:"request_id"`

	// Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsChannel bool `json:"chat_is_channel"`

	// Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat. If
	// not specified, no additional restrictions are applied.
	ChatIsForum bool `json:"chat_is_forum"`

	// Optional. Pass True to request a supergroup or a channel with a username, pass False to request
	// a chat without a username. If not specified, no additional restrictions are applied.
	ChatHasUsername bool `json:"chat_has_username"`

	// Optional. Pass True to request a chat owned by the user. Otherwise, no additional restrictions
	// are applied.
	ChatIsCreated bool `json:"chat_is_created"`

	// Optional. A JSON-serialized object listing the required administrator rights of the user in the
	// chat. The rights must be a superset of bot_administrator_rights. If not specified, no
	// additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights"`

	// Optional. A JSON-serialized object listing the required administrator rights of the bot in the
	// chat. The rights must be a subset of user_administrator_rights. If not specified, no additional
	// restrictions are applied.
	BotAdministratorRights *ChatAdministratorRights `json:"bot_administrator_rights"`

	// Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional
	// restrictions are applied.
	BotIsMember bool `json:"bot_is_member"`
}

// This object represents type of a poll, which is allowed to be created and sent when the
// corresponding button is pressed.
type KeyboardButtonPollType struct {

	// Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If
	// regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to
	// create a poll of any type.
	Type string `json:"type"`
}

// Upon receiving a message with this object, Telegram clients will remove the current custom
// keyboard and display the default letter-keyboard. By default, custom keyboards are displayed
// until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are
// hidden immediately after the user presses a button (see ReplyKeyboardMarkup ).
type ReplyKeyboardRemove struct {

	// Requests clients to remove the custom keyboard (user will not be able to summon this keyboard;
	// if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in
	// ReplyKeyboardMarkup)
	RemoveKeyboard bool `json:"remove_keyboard"`

	// Optional. Use this parameter if you want to remove the keyboard for specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's
	// message is a reply to a message in the same chat and forum topic, sender of the original
	// message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote
	// and removes the keyboard for that user, while still showing the keyboard with poll options to
	// users who haven't voted yet.
	Selective bool `json:"selective"`
}

// This object represents an inline keyboard that appears right next to the message it belongs
// to.
type InlineKeyboardMarkup struct {

	// Array of button rows, each represented by an Array of InlineKeyboardButton objects
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// This object represents one button of an inline keyboard. You must use exactly one of the
// optional fields.
type InlineKeyboardButton struct {

	// Label text on the button
	Text string `json:"text"`

	// Optional. HTTP or tg:// URL to be opened when the button is pressed. Links
	// tg://user?id=<user_id> can be used to mention a user by their identifier without using a
	// username, if this is allowed by their privacy settings.
	URL string `json:"url"`

	// Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data"`

	// Optional. Description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to send an arbitrary message on behalf of the user using the method
	// answerWebAppQuery. Available only in private chats between a user and the bot.
	WebApp *WebAppInfo `json:"web_app"`

	// Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement
	// for the Telegram Login Widget.
	LoginURL *LoginUrl `json:"login_url"`

	// Optional. If set, pressing the button will prompt the user to select one of their chats, open
	// that chat and insert the bot's username and the specified inline query in the input field. May
	// be empty, in which case just the bot's username will be inserted.
	SwitchInlineQuery string `json:"switch_inline_query"`

	// Optional. If set, pressing the button will insert the bot's username and the specified inline
	// query in the current chat's input field. May be empty, in which case only the bot's username
	// will be inserted.This offers a quick way for the user to open your bot in inline mode in the
	// same chat - good for selecting something from multiple options.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat"`

	// Optional. If set, pressing the button will prompt the user to select one of their chats of the
	// specified type, open that chat and insert the bot's username and the specified inline query in
	// the input field
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat"`

	// Optional. Description of the game that will be launched when the user presses the button.NOTE:
	// This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game"`

	// Optional. Specify True, to send a Pay button.NOTE: This type of button must always be the first
	// button in the first row and can only be used in invoice messages.
	Pay bool `json:"pay"`
}

// This object represents a parameter of the inline keyboard button used to automatically
// authorize a user. Serves as a great replacement for the Telegram Login Widget when the user is
// coming from Telegram. All the user needs to do is tap/click a button and confirm that they want
// to log in:     Telegram apps support these buttons as of version 5.7 .   Sample bot:
// @discussbot
type LoginUrl struct {

	// An HTTPS URL to be opened with user authorization data added to the query string when the
	// button is pressed. If the user refuses to provide authorization data, the original URL without
	// information about the user will be opened. The data added is the same as described in Receiving
	// authorization data.NOTE: You must always check the hash of the received data to verify the
	// authentication and the integrity of the data as described in Checking authorization.
	URL string `json:"url"`

	// Optional. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text"`

	// Optional. Username of a bot, which will be used for user authorization. See Setting up a bot
	// for more details. If not specified, the current bot's username will be assumed. The url's
	// domain must be the same as the domain linked with the bot. See Linking your domain to the bot
	// for more details.
	BotUsername string `json:"bot_username"`

	// Optional. Pass True to request the permission for your bot to send messages to the user.
	RequestWriteAccess bool `json:"request_write_access"`
}

// This object represents an inline button that switches the current user to inline mode in a
// chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {

	// Optional. The default inline query to be inserted in the input field. If left empty, only the
	// bot's username will be inserted
	Query string `json:"query"`

	// Optional. True, if private chats with users can be chosen
	AllowUserChats bool `json:"allow_user_chats"`

	// Optional. True, if private chats with bots can be chosen
	AllowBotChats bool `json:"allow_bot_chats"`

	// Optional. True, if group and supergroup chats can be chosen
	AllowGroupChats bool `json:"allow_group_chats"`

	// Optional. True, if channel chats can be chosen
	AllowChannelChats bool `json:"allow_channel_chats"`
}

// This object represents an incoming callback query from a callback button in an inline
// keyboard . If the button that originated the query was attached to a message sent by the bot,
// the field message will be present. If the button was attached to a message sent via the bot (in
// inline mode ), the field inline_message_id will be present. Exactly one of the fields data or
// game_short_name will be present.
type CallbackQuery struct {

	// Unique identifier for this query
	ID string `json:"id"`

	// Sender
	From *User `json:"from"`

	// Optional. Message sent by the bot with the callback button that originated the query
	Message *MaybeInaccessibleMessage `json:"message"`

	// Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	InlineMessageID string `json:"inline_message_id"`

	// Global identifier, uniquely corresponding to the chat to which the message with the callback
	// button was sent. Useful for high scores in games.
	ChatInstance string `json:"chat_instance"`

	// Optional. Data associated with the callback button. Be aware that the message originated the
	// query can contain no callback buttons with this data.
	Data string `json:"data"`

	// Optional. Short name of a Game to be returned, serves as the unique identifier for the game
	GameShortName string `json:"game_short_name"`
}

// NOTE: After the user presses a callback button, Telegram clients will display a progress bar
// until you call answerCallbackQuery . It is, therefore, necessary to react by calling
// answerCallbackQuery even if no notification to the user is needed (e.g., without specifying any
// of the optional parameters).    Upon receiving a message with this object, Telegram clients
// will display a reply interface to the user (act as if the user has selected the bot's message
// and tapped 'Reply'). This can be extremely useful if you want to create user-friendly
// step-by-step interfaces without having to sacrifice privacy mode .
type ForceReply struct {

	// Shows reply interface to the user, as if they manually selected the bot's message and tapped
	// 'Reply'
	ForceReply bool `json:"force_reply"`

	// Optional. The placeholder to be shown in the input field when the reply is active; 1-64
	// characters
	InputFieldPlaceholder string `json:"input_field_placeholder"`

	// Optional. Use this parameter if you want to force reply from specific users only. Targets: 1)
	// users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply
	// to a message in the same chat and forum topic, sender of the original message.
	Selective bool `json:"selective"`
}

// Example: A poll bot for groups runs in privacy mode (only receives commands, replies to its
// messages and mentions). There could be two ways to create a new poll:   Explain the user how to
// send a command with parameters (e.g. /newpoll question answer1 answer2). May be appealing for
// hardcore users but lacks modern day polish.  Guide the user through a step-by-step process.
// 'Please send me your question', 'Cool, now let's add the first answer option', 'Great. Keep
// adding answer options, then send /done when you're ready'.   The last option is definitely more
// attractive. And if you use ForceReply in your bot's questions, it will receive the user's
// answers even if it only receives replies, commands and mentions - without any extra work for
// the user.    This object represents a chat photo.
type ChatPhoto struct {

	// File identifier of small (160x160) chat photo. This file_id can be used only for photo download
	// and only for as long as the photo is not changed.
	SmallFileID string `json:"small_file_id"`

	// Unique file identifier of small (160x160) chat photo, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	SmallFileUniqueID string `json:"small_file_unique_id"`

	// File identifier of big (640x640) chat photo. This file_id can be used only for photo download
	// and only for as long as the photo is not changed.
	BigFileID string `json:"big_file_id"`

	// Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	BigFileUniqueID string `json:"big_file_unique_id"`
}

// Represents an invite link for a chat.
type ChatInviteLink struct {

	// The invite link. If the link was created by another chat administrator, then the second part of
	// the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`

	// Creator of the link
	Creator *User `json:"creator"`

	// True, if users joining the chat via the link need to be approved by chat administrators
	CreatesJoinRequest bool `json:"creates_join_request"`

	// True, if the link is primary
	IsPrimary bool `json:"is_primary"`

	// True, if the link is revoked
	IsRevoked bool `json:"is_revoked"`

	// Optional. Invite link name
	Name string `json:"name"`

	// Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	ExpireDate int64 `json:"expire_date"`

	// Optional. The maximum number of users that can be members of the chat simultaneously after
	// joining the chat via this invite link; 1-99999
	MemberLimit int64 `json:"member_limit"`

	// Optional. Number of pending join requests created using this link
	PendingJoinRequestCount int64 `json:"pending_join_request_count"`
}

// Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {

	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// True, if the administrator can access the chat event log, boost list in channels, see channel
	// members, report spam messages, see anonymous administrators in supergroups and ignore slow
	// mode. Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat"`

	// True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`

	// True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`

	// True, if the administrator can restrict, ban or unban chat members, or access supergroup
	// statistics
	CanRestrictMembers bool `json:"can_restrict_members"`

	// True, if the administrator can add new administrators with a subset of their own privileges or
	// demote administrators that they have promoted, directly or indirectly (promoted by
	// administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`

	// True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// Optional. True, if the administrator can post messages in the channel, or access channel
	// statistics; channels only
	CanPostMessages bool `json:"can_post_messages"`

	// Optional. True, if the administrator can edit messages of other users and can pin messages;
	// channels only
	CanEditMessages bool `json:"can_edit_messages"`

	// Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages"`

	// Optional. True, if the administrator can post stories in the channel; channels only
	CanPostStories bool `json:"can_post_stories"`

	// Optional. True, if the administrator can edit stories posted by other users; channels only
	CanEditStories bool `json:"can_edit_stories"`

	// Optional. True, if the administrator can delete stories posted by other users; channels only
	CanDeleteStories bool `json:"can_delete_stories"`

	// Optional. True, if the user is allowed to create, rename, close, and reopen forum topics;
	// supergroups only
	CanManageTopics bool `json:"can_manage_topics"`
}

// This object represents changes in the status of a chat member.
type ChatMemberUpdated struct {

	// Chat the user belongs to
	Chat *Chat `json:"chat"`

	// Performer of the action, which resulted in the change
	From *User `json:"from"`

	// Date the change was done in Unix time
	Date int64 `json:"date"`

	// Previous information about the chat member
	OldChatMember *ChatMember `json:"old_chat_member"`

	// New information about the chat member
	NewChatMember *ChatMember `json:"new_chat_member"`

	// Optional. Chat invite link, which was used by the user to join the chat; for joining by invite
	// link events only.
	InviteLink *ChatInviteLink `json:"invite_link"`

	// Optional. True, if the user joined the chat via a chat folder invite link
	ViaChatFolderInviteLink bool `json:"via_chat_folder_invite_link"`
}

// This object contains information about one member of a chat. Currently, the following 6 types
// of chat members are supported:   ChatMemberOwner  ChatMemberAdministrator  ChatMemberMember
// ChatMemberRestricted  ChatMemberLeft  ChatMemberBanned
type ChatMember struct {
}

// Represents a chat member that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {

	// The member's status in the chat, always “creator”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`

	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// Optional. Custom title for this user
	CustomTitle string `json:"custom_title"`
}

// Represents a chat member that has some additional privileges.
type ChatMemberAdministrator struct {

	// The member's status in the chat, always “administrator”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`

	// True, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited"`

	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// True, if the administrator can access the chat event log, boost list in channels, see channel
	// members, report spam messages, see anonymous administrators in supergroups and ignore slow
	// mode. Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat"`

	// True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`

	// True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`

	// True, if the administrator can restrict, ban or unban chat members, or access supergroup
	// statistics
	CanRestrictMembers bool `json:"can_restrict_members"`

	// True, if the administrator can add new administrators with a subset of their own privileges or
	// demote administrators that they have promoted, directly or indirectly (promoted by
	// administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`

	// True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// Optional. True, if the administrator can post messages in the channel, or access channel
	// statistics; channels only
	CanPostMessages bool `json:"can_post_messages"`

	// Optional. True, if the administrator can edit messages of other users and can pin messages;
	// channels only
	CanEditMessages bool `json:"can_edit_messages"`

	// Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages"`

	// Optional. True, if the administrator can post stories in the channel; channels only
	CanPostStories bool `json:"can_post_stories"`

	// Optional. True, if the administrator can edit stories posted by other users; channels only
	CanEditStories bool `json:"can_edit_stories"`

	// Optional. True, if the administrator can delete stories posted by other users; channels only
	CanDeleteStories bool `json:"can_delete_stories"`

	// Optional. True, if the user is allowed to create, rename, close, and reopen forum topics;
	// supergroups only
	CanManageTopics bool `json:"can_manage_topics"`

	// Optional. Custom title for this user
	CustomTitle string `json:"custom_title"`
}

// Represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {

	// The member's status in the chat, always “member”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`
}

// Represents a chat member that is under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {

	// The member's status in the chat, always “restricted”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`

	// True, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member"`

	// True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners,
	// invoices, locations and venues
	CanSendMessages bool `json:"can_send_messages"`

	// True, if the user is allowed to send audios
	CanSendAudios bool `json:"can_send_audios"`

	// True, if the user is allowed to send documents
	CanSendDocuments bool `json:"can_send_documents"`

	// True, if the user is allowed to send photos
	CanSendPhotos bool `json:"can_send_photos"`

	// True, if the user is allowed to send videos
	CanSendVideos bool `json:"can_send_videos"`

	// True, if the user is allowed to send video notes
	CanSendVideoNotes bool `json:"can_send_video_notes"`

	// True, if the user is allowed to send voice notes
	CanSendVoiceNotes bool `json:"can_send_voice_notes"`

	// True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls"`

	// True, if the user is allowed to send animations, games, stickers and use inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages"`

	// True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`

	// True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// True, if the user is allowed to pin messages
	CanPinMessages bool `json:"can_pin_messages"`

	// True, if the user is allowed to create forum topics
	CanManageTopics bool `json:"can_manage_topics"`

	// Date when restrictions will be lifted for this user; Unix time. If 0, then the user is
	// restricted forever
	UntilDate int64 `json:"until_date"`
}

// Represents a chat member that isn't currently a member of the chat, but may join it
// themselves.
type ChatMemberLeft struct {

	// The member's status in the chat, always “left”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`
}

// Represents a chat member that was banned in the chat and can't return to the chat or view
// chat messages.
type ChatMemberBanned struct {

	// The member's status in the chat, always “kicked”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`

	// Date when restrictions will be lifted for this user; Unix time. If 0, then the user is banned
	// forever
	UntilDate int64 `json:"until_date"`
}

// Represents a join request sent to a chat.
type ChatJoinRequest struct {

	// Chat to which the request was sent
	Chat *Chat `json:"chat"`

	// User that sent the join request
	From *User `json:"from"`

	// Identifier of a private chat with the user who sent the join request. This number may have more
	// than 32 significant bits and some programming languages may have difficulty/silent defects in
	// interpreting it. But it has at most 52 significant bits, so a 64-bit integer or
	// double-precision float type are safe for storing this identifier. The bot can use this
	// identifier for 5 minutes to send messages until the join request is processed, assuming no
	// other administrator contacted the user.
	UserChatID int64 `json:"user_chat_id"`

	// Date the request was sent in Unix time
	Date int64 `json:"date"`

	// Optional. Bio of the user.
	Bio string `json:"bio"`

	// Optional. Chat invite link that was used by the user to send the join request
	InviteLink *ChatInviteLink `json:"invite_link"`
}

// Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {

	// Optional. True, if the user is allowed to send text messages, contacts, giveaways, giveaway
	// winners, invoices, locations and venues
	CanSendMessages bool `json:"can_send_messages"`

	// Optional. True, if the user is allowed to send audios
	CanSendAudios bool `json:"can_send_audios"`

	// Optional. True, if the user is allowed to send documents
	CanSendDocuments bool `json:"can_send_documents"`

	// Optional. True, if the user is allowed to send photos
	CanSendPhotos bool `json:"can_send_photos"`

	// Optional. True, if the user is allowed to send videos
	CanSendVideos bool `json:"can_send_videos"`

	// Optional. True, if the user is allowed to send video notes
	CanSendVideoNotes bool `json:"can_send_video_notes"`

	// Optional. True, if the user is allowed to send voice notes
	CanSendVoiceNotes bool `json:"can_send_voice_notes"`

	// Optional. True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls"`

	// Optional. True, if the user is allowed to send animations, games, stickers and use inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages"`

	// Optional. True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`

	// Optional. True, if the user is allowed to change the chat title, photo and other settings.
	// Ignored in public supergroups
	CanChangeInfo bool `json:"can_change_info"`

	// Optional. True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
	CanPinMessages bool `json:"can_pin_messages"`

	// Optional. True, if the user is allowed to create forum topics. If omitted defaults to the value
	// of can_pin_messages
	CanManageTopics bool `json:"can_manage_topics"`
}

// Represents a location to which a chat is connected.
type ChatLocation struct {

	// The location to which the supergroup is connected. Can't be a live location.
	Location *Location `json:"location"`

	// Location address; 1-64 characters, as defined by the chat owner
	Address string `json:"address"`
}

// The reaction is based on an emoji.
type ReactionTypeEmoji struct {

	// Type of the reaction, always “emoji”
	Type string `json:"type"`

	// Reaction emoji. Currently, it can be one of "", "", "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", ""
	Emoji string `json:"emoji"`
}

// The reaction is based on a custom emoji.
type ReactionTypeCustomEmoji struct {

	// Type of the reaction, always “custom_emoji”
	Type string `json:"type"`

	// Custom emoji identifier
	CustomEmojiID string `json:"custom_emoji_id"`
}

// Represents a reaction added to a message along with the number of times it was added.
type ReactionCount struct {

	// Type of the reaction
	Type *ReactionType `json:"type"`

	// Number of times the reaction was added
	TotalCount int64 `json:"total_count"`
}

// This object represents a change of a reaction on a message performed by a user.
type MessageReactionUpdated struct {

	// The chat containing the message the user reacted to
	Chat *Chat `json:"chat"`

	// Unique identifier of the message inside the chat
	MessageID int64 `json:"message_id"`

	// Optional. The user that changed the reaction, if the user isn't anonymous
	User *User `json:"user"`

	// Optional. The chat on behalf of which the reaction was changed, if the user is anonymous
	ActorChat *Chat `json:"actor_chat"`

	// Date of the change in Unix time
	Date int64 `json:"date"`

	// Previous list of reaction types that were set by the user
	OldReaction []*ReactionType `json:"old_reaction"`

	// New list of reaction types that have been set by the user
	NewReaction []*ReactionType `json:"new_reaction"`
}

// This object represents reaction changes on a message with anonymous reactions.
type MessageReactionCountUpdated struct {

	// The chat containing the message
	Chat *Chat `json:"chat"`

	// Unique message identifier inside the chat
	MessageID int64 `json:"message_id"`

	// Date of the change in Unix time
	Date int64 `json:"date"`

	// List of reactions that are present on the message
	Reactions []*ReactionCount `json:"reactions"`
}

// This object represents a forum topic.
type ForumTopic struct {

	// Unique identifier of the forum topic
	MessageThreadID int64 `json:"message_thread_id"`

	// Name of the topic
	Name string `json:"name"`

	// Color of the topic icon in RGB format
	IconColor int64 `json:"icon_color"`

	// Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
}

// This object represents a bot command.
type BotCommand struct {

	// Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and
	// underscores.
	Command string `json:"command"`

	// Description of the command; 1-256 characters.
	Description string `json:"description"`
}

// This object represents the scope to which bot commands are applied. Currently, the following
// 7 scopes are supported:   BotCommandScopeDefault  BotCommandScopeAllPrivateChats
// BotCommandScopeAllGroupChats  BotCommandScopeAllChatAdministrators  BotCommandScopeChat
// BotCommandScopeChatAdministrators  BotCommandScopeChatMember
type BotCommandScope struct {
}

// Represents the default scope of bot commands. Default commands are used if no commands with a
// narrower scope are specified for the user.
type BotCommandScopeDefault struct {

	// Scope type, must be default
	Type string `json:"type"`
}

// Represents the scope of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {

	// Scope type, must be all_private_chats
	Type string `json:"type"`
}

// Represents the scope of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {

	// Scope type, must be all_group_chats
	Type string `json:"type"`
}

// Represents the scope of bot commands, covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct {

	// Scope type, must be all_chat_administrators
	Type string `json:"type"`
}

// Represents the scope of bot commands, covering a specific chat.
type BotCommandScopeChat struct {

	// Scope type, must be chat
	Type string `json:"type"`

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`
}

// Represents the scope of bot commands, covering all administrators of a specific group or
// supergroup chat.
type BotCommandScopeChatAdministrators struct {

	// Scope type, must be chat_administrators
	Type string `json:"type"`

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`
}

// Represents the scope of bot commands, covering a specific member of a group or supergroup
// chat.
type BotCommandScopeChatMember struct {

	// Scope type, must be chat_member
	Type string `json:"type"`

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// This object represents the bot's name.
type BotName struct {

	// The bot's name
	Name string `json:"name"`
}

// This object represents the bot's description.
type BotDescription struct {

	// The bot's description
	Description string `json:"description"`
}

// This object represents the bot's short description.
type BotShortDescription struct {

	// The bot's short description
	ShortDescription string `json:"short_description"`
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
	Type string `json:"type"`
}

// Represents a menu button, which launches a Web App .
type MenuButtonWebApp struct {

	// Type of the button, must be web_app
	Type string `json:"type"`

	// Text on the button
	Text string `json:"text"`

	// Description of the Web App that will be launched when the user presses the button. The Web App
	// will be able to send an arbitrary message on behalf of the user using the method
	// answerWebAppQuery.
	WebApp *WebAppInfo `json:"web_app"`
}

// Describes that no specific value for the menu button was set.
type MenuButtonDefault struct {

	// Type of the button, must be default
	Type string `json:"type"`
}

// This object describes the source of a chat boost. It can be one of   ChatBoostSourcePremium
// ChatBoostSourceGiftCode  ChatBoostSourceGiveaway
type ChatBoostSource struct {
}

// The boost was obtained by subscribing to Telegram Premium or by gifting a Telegram Premium
// subscription to another user.
type ChatBoostSourcePremium struct {

	// Source of the boost, always “premium”
	Source string `json:"source"`

	// User that boosted the chat
	User *User `json:"user"`
}

// The boost was obtained by the creation of Telegram Premium gift codes to boost a chat. Each
// such code boosts the chat 4 times for the duration of the corresponding Telegram Premium
// subscription.
type ChatBoostSourceGiftCode struct {

	// Source of the boost, always “gift_code”
	Source string `json:"source"`

	// User for which the gift code was created
	User *User `json:"user"`
}

// The boost was obtained by the creation of a Telegram Premium giveaway. This boosts the chat 4
// times for the duration of the corresponding Telegram Premium subscription.
type ChatBoostSourceGiveaway struct {

	// Source of the boost, always “giveaway”
	Source string `json:"source"`

	// Identifier of a message in the chat with the giveaway; the message could have been deleted
	// already. May be 0 if the message isn't sent yet.
	GiveawayMessageID int64 `json:"giveaway_message_id"`

	// Optional. User that won the prize in the giveaway if any
	User *User `json:"user"`

	// Optional. True, if the giveaway was completed, but there was no user to win the prize
	IsUnclaimed bool `json:"is_unclaimed"`
}

// This object contains information about a chat boost.
type ChatBoost struct {

	// Unique identifier of the boost
	BoostID string `json:"boost_id"`

	// Point in time (Unix timestamp) when the chat was boosted
	AddDate int64 `json:"add_date"`

	// Point in time (Unix timestamp) when the boost will automatically expire, unless the booster's
	// Telegram Premium subscription is prolonged
	ExpirationDate int64 `json:"expiration_date"`

	// Source of the added boost
	Source *ChatBoostSource `json:"source"`
}

// This object represents a boost added to a chat or changed.
type ChatBoostUpdated struct {

	// Chat which was boosted
	Chat *Chat `json:"chat"`

	// Infomation about the chat boost
	Boost *ChatBoost `json:"boost"`
}

// This object represents a boost removed from a chat.
type ChatBoostRemoved struct {

	// Chat which was boosted
	Chat *Chat `json:"chat"`

	// Unique identifier of the boost
	BoostID string `json:"boost_id"`

	// Point in time (Unix timestamp) when the boost was removed
	RemoveDate int64 `json:"remove_date"`

	// Source of the removed boost
	Source *ChatBoostSource `json:"source"`
}

// This object represents a list of boosts added to a chat by a user.
type UserChatBoosts struct {

	// The list of boosts added to the chat by the user
	Boosts []*ChatBoost `json:"boosts"`
}

// Describes why a request was unsuccessful.
type ResponseParameters struct {

	// Optional. The group has been migrated to a supergroup with the specified identifier. This
	// number may have more than 32 significant bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id"`

	// Optional. In case of exceeding flood control, the number of seconds left to wait before the
	// request can be repeated
	RetryAfter int64 `json:"retry_after"`
}

// This object represents the content of a media message to be sent. It should be one of
// InputMediaAnimation  InputMediaDocument  InputMediaAudio  InputMediaPhoto  InputMediaVideo
type InputMedia struct {
}

// Represents a photo to be sent.
type InputMediaPhoto struct {

	// Type of the result, must be photo
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name. More information on Sending Files »
	Media string `json:"media"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the photo caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler"`
}

// Represents a video to be sent.
type InputMediaVideo struct {

	// Type of the result, must be video
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name. More information on Sending Files »
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A
	// thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you
	// can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data
	// under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the video caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Video width
	Width int64 `json:"width"`

	// Optional. Video height
	Height int64 `json:"height"`

	// Optional. Video duration in seconds
	Duration int64 `json:"duration"`

	// Optional. Pass True if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming"`

	// Optional. Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler"`
}

// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {

	// Type of the result, must be animation
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name. More information on Sending Files »
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A
	// thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you
	// can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data
	// under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail"`

	// Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the animation caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Animation width
	Width int64 `json:"width"`

	// Optional. Animation height
	Height int64 `json:"height"`

	// Optional. Animation duration in seconds
	Duration int64 `json:"duration"`

	// Optional. Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler"`
}

// Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {

	// Type of the result, must be audio
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name. More information on Sending Files »
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A
	// thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you
	// can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data
	// under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail"`

	// Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the audio caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Duration of the audio in seconds
	Duration int64 `json:"duration"`

	// Optional. Performer of the audio
	Performer string `json:"performer"`

	// Optional. Title of the audio
	Title string `json:"title"`
}

// Represents a general file to be sent.
type InputMediaDocument struct {

	// Type of the result, must be document
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name. More information on Sending Files »
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A
	// thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you
	// can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data
	// under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail"`

	// Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the document caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Disables automatic server-side content type detection for files uploaded using
	// multipart/form-data. Always True, if the document is sent as part of an album.
	DisableContentTypeDetection bool `json:"disable_content_type_detection"`
}

// This object represents the contents of a file to be uploaded. Must be posted using
// multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile struct {
}

// The following methods and objects allow your bot to handle stickers and sticker sets.   This
// object represents a sticker.
type Sticker struct {

	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different
	// bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”. The type of the
	// sticker is independent from its format, which is determined by the fields is_animated and
	// is_video.
	Type string `json:"type"`

	// Sticker width
	Width int64 `json:"width"`

	// Sticker height
	Height int64 `json:"height"`

	// True, if the sticker is animated
	IsAnimated bool `json:"is_animated"`

	// True, if the sticker is a video sticker
	IsVideo bool `json:"is_video"`

	// Optional. Sticker thumbnail in the .WEBP or .JPG format
	Thumbnail *PhotoSize `json:"thumbnail"`

	// Optional. Emoji associated with the sticker
	Emoji string `json:"emoji"`

	// Optional. Name of the sticker set to which the sticker belongs
	SetName string `json:"set_name"`

	// Optional. For premium regular stickers, premium animation for the sticker
	PremiumAnimation *File `json:"premium_animation"`

	// Optional. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position"`

	// Optional. For custom emoji stickers, unique identifier of the custom emoji
	CustomEmojiID string `json:"custom_emoji_id"`

	// Optional. True, if the sticker must be repainted to a text color in messages, the color of the
	// Telegram Premium badge in emoji status, white color on chat photos, or another appropriate
	// color in other places
	NeedsRepainting bool `json:"needs_repainting"`

	// Optional. File size in bytes
	FileSize int64 `json:"file_size"`
}

// This object represents a sticker set.
type StickerSet struct {

	// Sticker set name
	Name string `json:"name"`

	// Sticker set title
	Title string `json:"title"`

	// Type of stickers in the set, currently one of “regular”, “mask”, “custom_emoji”
	StickerType string `json:"sticker_type"`

	// True, if the sticker set contains animated stickers
	IsAnimated bool `json:"is_animated"`

	// True, if the sticker set contains video stickers
	IsVideo bool `json:"is_video"`

	// List of all set stickers
	Stickers []*Sticker `json:"stickers"`

	// Optional. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
	Thumbnail *PhotoSize `json:"thumbnail"`
}

// This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {

	// The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”,
	// “mouth”, or “chin”.
	Point string `json:"point"`

	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For
	// example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift float32 `json:"x_shift"`

	// Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom.
	// For example, 1.0 will place the mask just below the default mask position.
	YShift float32 `json:"y_shift"`

	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float32 `json:"scale"`
}

// This object describes a sticker to be added to a sticker set.
type InputSticker struct {

	// The added sticker. Pass a file_id as a String to send a file that already exists on the
	// Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet,
	// upload a new one using multipart/form-data, or pass “attach://<file_attach_name>” to upload a
	// new one using multipart/form-data under <file_attach_name> name. Animated and video stickers
	// can't be uploaded via HTTP URL. More information on Sending Files »
	Sticker interface{} `json:"sticker"`

	// List of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list"`

	// Optional. Position where the mask should be placed on faces. For “mask” stickers only.
	MaskPosition *MaskPosition `json:"mask_position"`

	// Optional. List of 0-20 search keywords for the sticker with total length of up to 64
	// characters. For “regular” and “custom_emoji” stickers only.
	Keywords []string `json:"keywords"`
}

// The following methods and objects allow your bot to work in inline mode . Please see our
// Introduction to Inline bots for more details.  To enable this option, send the /setinline
// command to @BotFather and provide the placeholder text that the user will see in the input
// field after typing your bot's name.   This object represents an incoming inline query. When the
// user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {

	// Unique identifier for this query
	ID string `json:"id"`

	// Sender
	From *User `json:"from"`

	// Text of the query (up to 256 characters)
	Query string `json:"query"`

	// Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`

	// Optional. Type of the chat from which the inline query was sent. Can be either “sender” for a
	// private chat with the inline query sender, “private”, “group”, “supergroup”, or “channel”. The
	// chat type should be always known for requests sent from official clients and most third-party
	// clients, unless the request was sent from a secret chat
	ChatType string `json:"chat_type"`

	// Optional. Sender location, only for bots that request user location
	Location *Location `json:"location"`
}

// This object represents a button to be shown above inline query results. You must use exactly
// one of the optional fields.
type InlineQueryResultsButton struct {

	// Label text on the button
	Text string `json:"text"`

	// Optional. Description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to switch back to the inline mode using the method switchInlineQuery
	// inside the Web App.
	WebApp *WebAppInfo `json:"web_app"`

	// Optional. Deep-linking parameter for the /start message sent to the bot when a user presses the
	// button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that
	// sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt
	// search results accordingly. To do this, it displays a 'Connect your YouTube account' button
	// above the results, or even before showing any. The user presses the button, switches to a
	// private chat with the bot and, in doing so, passes a start parameter that instructs the bot to
	// return an OAuth link. Once done, the bot can offer a switch_inline button so that the user can
	// easily return to the chat where they wanted to use the bot's inline capabilities.
	StartParameter string `json:"start_parameter"`
}

// This object represents one result of an inline query. Telegram clients currently support
// results of the following 20 types:   InlineQueryResultCachedAudio
// InlineQueryResultCachedDocument  InlineQueryResultCachedGif  InlineQueryResultCachedMpeg4Gif
// InlineQueryResultCachedPhoto  InlineQueryResultCachedSticker  InlineQueryResultCachedVideo
// InlineQueryResultCachedVoice  InlineQueryResultArticle  InlineQueryResultAudio
// InlineQueryResultContact  InlineQueryResultGame  InlineQueryResultDocument
// InlineQueryResultGif  InlineQueryResultLocation  InlineQueryResultMpeg4Gif
// InlineQueryResultPhoto  InlineQueryResultVenue  InlineQueryResultVideo  InlineQueryResultVoice
// Note: All URLs passed in inline query results will be available to end users and therefore must
// be assumed to be public .
type InlineQueryResult struct {
}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {

	// Type of the result, must be article
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`

	// Title of the result
	Title string `json:"title"`

	// Content of the message to be sent
	InputMessageContent *InputMessageContent `json:"input_message_content"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. URL of the result
	URL string `json:"url"`

	// Optional. Pass True if you don't want the URL to be shown in the message
	HideURL bool `json:"hide_url"`

	// Optional. Short description of the result
	Description string `json:"description"`

	// Optional. Url of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url"`

	// Optional. Thumbnail width
	ThumbnailWidth int64 `json:"thumbnail_width"`

	// Optional. Thumbnail height
	ThumbnailHeight int64 `json:"thumbnail_height"`
}

// Represents a link to a photo. By default, this photo will be sent by the user with optional
// caption. Alternatively, you can use input_message_content to send a message with the specified
// content instead of the photo.
type InlineQueryResultPhoto struct {

	// Type of the result, must be photo
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB
	PhotoURL string `json:"photo_url"`

	// URL of the thumbnail for the photo
	ThumbnailURL string `json:"thumbnail_url"`

	// Optional. Width of the photo
	PhotoWidth int64 `json:"photo_width"`

	// Optional. Height of the photo
	PhotoHeight int64 `json:"photo_height"`

	// Optional. Title for the result
	Title string `json:"title"`

	// Optional. Short description of the result
	Description string `json:"description"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the photo caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the photo
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by
// the user with optional caption. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the animation.
type InlineQueryResultGif struct {

	// Type of the result, must be gif
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid URL for the GIF file. File size must not exceed 1MB
	GifURL string `json:"gif_url"`

	// Optional. Width of the GIF
	GifWidth int64 `json:"gif_width"`

	// Optional. Height of the GIF
	GifHeight int64 `json:"gif_height"`

	// Optional. Duration of the GIF in seconds
	GifDuration int64 `json:"gif_duration"`

	// URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url"`

	// Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”.
	// Defaults to “image/jpeg”
	ThumbnailMimeType string `json:"thumbnail_mime_type"`

	// Optional. Title for the result
	Title string `json:"title"`

	// Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default,
// this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you
// can use input_message_content to send a message with the specified content instead of the
// animation.
type InlineQueryResultMpeg4Gif struct {

	// Type of the result, must be mpeg4_gif
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid URL for the MPEG4 file. File size must not exceed 1MB
	Mpeg4URL string `json:"mpeg4_url"`

	// Optional. Video width
	Mpeg4Width int64 `json:"mpeg4_width"`

	// Optional. Video height
	Mpeg4Height int64 `json:"mpeg4_height"`

	// Optional. Video duration in seconds
	Mpeg4Duration int64 `json:"mpeg4_duration"`

	// URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url"`

	// Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”.
	// Defaults to “image/jpeg”
	ThumbnailMimeType string `json:"thumbnail_mime_type"`

	// Optional. Title for the result
	Title string `json:"title"`

	// Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the video animation
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to a page containing an embedded video player or a video file. By default,
// this video file will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the video.   If
// an InlineQueryResultVideo message contains an embedded video (e.g., YouTube), you must replace
// its content using input_message_content .
type InlineQueryResultVideo struct {

	// Type of the result, must be video
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid URL for the embedded video player or video file
	VideoURL string `json:"video_url"`

	// MIME type of the content of the video URL, “text/html” or “video/mp4”
	MimeType string `json:"mime_type"`

	// URL of the thumbnail (JPEG only) for the video
	ThumbnailURL string `json:"thumbnail_url"`

	// Title for the result
	Title string `json:"title"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the video caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Video width
	VideoWidth int64 `json:"video_width"`

	// Optional. Video height
	VideoHeight int64 `json:"video_height"`

	// Optional. Video duration in seconds
	VideoDuration int64 `json:"video_duration"`

	// Optional. Short description of the result
	Description string `json:"description"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the video. This field is required if
	// InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to an MP3 audio file. By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content
// instead of the audio.
type InlineQueryResultAudio struct {

	// Type of the result, must be audio
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid URL for the audio file
	AudioURL string `json:"audio_url"`

	// Title
	Title string `json:"title"`

	// Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the audio caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Performer
	Performer string `json:"performer"`

	// Optional. Audio duration in seconds
	AudioDuration int64 `json:"audio_duration"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the audio
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to a voice recording in an .OGG container encoded with OPUS. By default,
// this voice recording will be sent by the user. Alternatively, you can use input_message_content
// to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {

	// Type of the result, must be voice
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid URL for the voice recording
	VoiceURL string `json:"voice_url"`

	// Recording title
	Title string `json:"title"`

	// Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the voice message caption. See formatting options for
	// more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Recording duration in seconds
	VoiceDuration int64 `json:"voice_duration"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the voice recording
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to a file. By default, this file will be sent by the user with an optional
// caption. Alternatively, you can use input_message_content to send a message with the specified
// content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {

	// Type of the result, must be document
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// Title for the result
	Title string `json:"title"`

	// Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the document caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// A valid URL for the file
	DocumentURL string `json:"document_url"`

	// MIME type of the content of the file, either “application/pdf” or “application/zip”
	MimeType string `json:"mime_type"`

	// Optional. Short description of the result
	Description string `json:"description"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the file
	InputMessageContent *InputMessageContent `json:"input_message_content"`

	// Optional. URL of the thumbnail (JPEG only) for the file
	ThumbnailURL string `json:"thumbnail_url"`

	// Optional. Thumbnail width
	ThumbnailWidth int64 `json:"thumbnail_width"`

	// Optional. Thumbnail height
	ThumbnailHeight int64 `json:"thumbnail_height"`
}

// Represents a location on a map. By default, the location will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content
// instead of the location.
type InlineQueryResultLocation struct {

	// Type of the result, must be location
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`

	// Location latitude in degrees
	Latitude float32 `json:"latitude"`

	// Location longitude in degrees
	Longitude float32 `json:"longitude"`

	// Location title
	Title string `json:"title"`

	// Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy"`

	// Optional. Period in seconds for which the location can be updated, should be between 60 and
	// 86400.
	LivePeriod int64 `json:"live_period"`

	// Optional. For live locations, a direction in which the user is moving, in degrees. Must be
	// between 1 and 360 if specified.
	Heading int64 `json:"heading"`

	// Optional. For live locations, a maximum distance for proximity alerts about approaching another
	// chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int64 `json:"proximity_alert_radius"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the location
	InputMessageContent *InputMessageContent `json:"input_message_content"`

	// Optional. Url of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url"`

	// Optional. Thumbnail width
	ThumbnailWidth int64 `json:"thumbnail_width"`

	// Optional. Thumbnail height
	ThumbnailHeight int64 `json:"thumbnail_height"`
}

// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can
// use input_message_content to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {

	// Type of the result, must be venue
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`

	// Latitude of the venue location in degrees
	Latitude float32 `json:"latitude"`

	// Longitude of the venue location in degrees
	Longitude float32 `json:"longitude"`

	// Title of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Optional. Foursquare identifier of the venue if known
	FoursquareID string `json:"foursquare_id"`

	// Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type"`

	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id"`

	// Optional. Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the venue
	InputMessageContent *InputMessageContent `json:"input_message_content"`

	// Optional. Url of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url"`

	// Optional. Thumbnail width
	ThumbnailWidth int64 `json:"thumbnail_width"`

	// Optional. Thumbnail height
	ThumbnailHeight int64 `json:"thumbnail_height"`
}

// Represents a contact with a phone number. By default, this contact will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content
// instead of the contact.
type InlineQueryResultContact struct {

	// Type of the result, must be contact
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`

	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name"`

	// Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	Vcard string `json:"vcard"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the contact
	InputMessageContent *InputMessageContent `json:"input_message_content"`

	// Optional. Url of the thumbnail for the result
	ThumbnailURL string `json:"thumbnail_url"`

	// Optional. Thumbnail width
	ThumbnailWidth int64 `json:"thumbnail_width"`

	// Optional. Thumbnail height
	ThumbnailHeight int64 `json:"thumbnail_height"`
}

// Represents a Game .
type InlineQueryResultGame struct {

	// Type of the result, must be game
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// Short name of the game
	GameShortName string `json:"game_short_name"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// Represents a link to a photo stored on the Telegram servers. By default, this photo will be
// sent by the user with an optional caption. Alternatively, you can use input_message_content to
// send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {

	// Type of the result, must be photo
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid file identifier of the photo
	PhotoFileID string `json:"photo_file_id"`

	// Optional. Title for the result
	Title string `json:"title"`

	// Optional. Short description of the result
	Description string `json:"description"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the photo caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the photo
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to an animated GIF file stored on the Telegram servers. By default, this
// animated GIF file will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {

	// Type of the result, must be gif
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid file identifier for the GIF file
	GifFileID string `json:"gif_file_id"`

	// Optional. Title for the result
	Title string `json:"title"`

	// Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the
// Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an
// optional caption. Alternatively, you can use input_message_content to send a message with the
// specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {

	// Type of the result, must be mpeg4_gif
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid file identifier for the MPEG4 file
	Mpeg4FileID string `json:"mpeg4_file_id"`

	// Optional. Title for the result
	Title string `json:"title"`

	// Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the video animation
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will
// be sent by the user. Alternatively, you can use input_message_content to send a message with
// the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {

	// Type of the result, must be sticker
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid file identifier of the sticker
	StickerFileID string `json:"sticker_file_id"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the sticker
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to a file stored on the Telegram servers. By default, this file will be
// sent by the user with an optional caption. Alternatively, you can use input_message_content to
// send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {

	// Type of the result, must be document
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// Title for the result
	Title string `json:"title"`

	// A valid file identifier for the file
	DocumentFileID string `json:"document_file_id"`

	// Optional. Short description of the result
	Description string `json:"description"`

	// Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the document caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the file
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to a video file stored on the Telegram servers. By default, this video file
// will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {

	// Type of the result, must be video
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid file identifier for the video file
	VideoFileID string `json:"video_file_id"`

	// Title for the result
	Title string `json:"title"`

	// Optional. Short description of the result
	Description string `json:"description"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the video caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the video
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to a voice message stored on the Telegram servers. By default, this voice
// message will be sent by the user. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {

	// Type of the result, must be voice
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid file identifier for the voice message
	VoiceFileID string `json:"voice_file_id"`

	// Voice message title
	Title string `json:"title"`

	// Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the voice message caption. See formatting options for
	// more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the voice message
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// Represents a link to an MP3 audio file stored on the Telegram servers. By default, this audio
// file will be sent by the user. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {

	// Type of the result, must be audio
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// A valid file identifier for the audio file
	AudioFileID string `json:"audio_file_id"`

	// Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the audio caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified instead
	// of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the audio
	InputMessageContent *InputMessageContent `json:"input_message_content"`
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
	MessageText string `json:"message_text"`

	// Optional. Mode for parsing entities in the message text. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in message text, which can be specified instead
	// of parse_mode
	Entities []*MessageEntity `json:"entities"`

	// Optional. Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"`
}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {

	// Latitude of the location in degrees
	Latitude float32 `json:"latitude"`

	// Longitude of the location in degrees
	Longitude float32 `json:"longitude"`

	// Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy"`

	// Optional. Period in seconds for which the location can be updated, should be between 60 and
	// 86400.
	LivePeriod int64 `json:"live_period"`

	// Optional. For live locations, a direction in which the user is moving, in degrees. Must be
	// between 1 and 360 if specified.
	Heading int64 `json:"heading"`

	// Optional. For live locations, a maximum distance for proximity alerts about approaching another
	// chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int64 `json:"proximity_alert_radius"`
}

// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {

	// Latitude of the venue in degrees
	Latitude float32 `json:"latitude"`

	// Longitude of the venue in degrees
	Longitude float32 `json:"longitude"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Optional. Foursquare identifier of the venue, if known
	FoursquareID string `json:"foursquare_id"`

	// Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type"`

	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id"`

	// Optional. Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type"`
}

// Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {

	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name"`

	// Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	Vcard string `json:"vcard"`
}

// Represents the content of an invoice message to be sent as the result of an inline query.
type InputInvoiceMessageContent struct {

	// Product name, 1-32 characters
	Title string `json:"title"`

	// Product description, 1-255 characters
	Description string `json:"description"`

	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your
	// internal processes.
	Payload string `json:"payload"`

	// Payment provider token, obtained via @BotFather
	ProviderToken string `json:"provider_token"`

	// Three-letter ISO 4217 currency code, see more on currencies
	Currency string `json:"currency"`

	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount,
	// delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices"`

	// Optional. The maximum accepted amount for tips in the smallest units of the currency (integer,
	// not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See
	// the exp parameter in currencies.json, it shows the number of digits past the decimal point for
	// each currency (2 for the majority of currencies). Defaults to 0
	MaxTipAmount int64 `json:"max_tip_amount"`

	// Optional. A JSON-serialized array of suggested amounts of tip in the smallest units of the
	// currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The
	// suggested tip amounts must be positive, passed in a strictly increased order and must not
	// exceed max_tip_amount.
	SuggestedTipAmounts []int64 `json:"suggested_tip_amounts"`

	// Optional. A JSON-serialized object for data about the invoice, which will be shared with the
	// payment provider. A detailed description of the required fields should be provided by the
	// payment provider.
	ProviderData string `json:"provider_data"`

	// Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing
	// image for a service.
	PhotoURL string `json:"photo_url"`

	// Optional. Photo size in bytes
	PhotoSize int64 `json:"photo_size"`

	// Optional. Photo width
	PhotoWidth int64 `json:"photo_width"`

	// Optional. Photo height
	PhotoHeight int64 `json:"photo_height"`

	// Optional. Pass True if you require the user's full name to complete the order
	NeedName bool `json:"need_name"`

	// Optional. Pass True if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number"`

	// Optional. Pass True if you require the user's email address to complete the order
	NeedEmail bool `json:"need_email"`

	// Optional. Pass True if you require the user's shipping address to complete the order
	NeedShippingAddress bool `json:"need_shipping_address"`

	// Optional. Pass True if the user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider"`

	// Optional. Pass True if the user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider"`

	// Optional. Pass True if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible"`
}

// Represents a result of an inline query that was chosen by the user and sent to their chat
// partner.
type ChosenInlineResult struct {

	// The unique identifier for the result that was chosen
	ResultID string `json:"result_id"`

	// The user that chose the result
	From *User `json:"from"`

	// Optional. Sender location, only for bots that require user location
	Location *Location `json:"location"`

	// Optional. Identifier of the sent inline message. Available only if there is an inline keyboard
	// attached to the message. Will be also received in callback queries and can be used to edit the
	// message.
	InlineMessageID string `json:"inline_message_id"`

	// The query that was used to obtain the result
	Query string `json:"query"`
}

// Describes an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {

	// Optional. Identifier of the sent inline message. Available only if there is an inline keyboard
	// attached to the message.
	InlineMessageID string `json:"inline_message_id"`
}

// This object represents a portion of the price for goods or services.
type LabeledPrice struct {

	// Portion label
	Label string `json:"label"`

	// Price of the product in the smallest units of the currency (integer, not float/double). For
	// example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of
	// currencies).
	Amount int64 `json:"amount"`
}

// This object contains basic information about an invoice.
type Invoice struct {

	// Product name
	Title string `json:"title"`

	// Product description
	Description string `json:"description"`

	// Unique bot deep-linking parameter that can be used to generate this invoice
	StartParameter string `json:"start_parameter"`

	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`

	// Total price in the smallest units of the currency (integer, not float/double). For example, for
	// a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the
	// number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int64 `json:"total_amount"`
}

// This object represents a shipping address.
type ShippingAddress struct {

	// Two-letter ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`

	// State, if applicable
	State string `json:"state"`

	// City
	City string `json:"city"`

	// First line for the address
	StreetLine1 string `json:"street_line1"`

	// Second line for the address
	StreetLine2 string `json:"street_line2"`

	// Address post code
	PostCode string `json:"post_code"`
}

// This object represents information about an order.
type OrderInfo struct {

	// Optional. User name
	Name string `json:"name"`

	// Optional. User's phone number
	PhoneNumber string `json:"phone_number"`

	// Optional. User email
	Email string `json:"email"`

	// Optional. User shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// This object represents one shipping option.
type ShippingOption struct {

	// Shipping option identifier
	ID string `json:"id"`

	// Option title
	Title string `json:"title"`

	// List of price portions
	Prices []*LabeledPrice `json:"prices"`
}

// This object contains basic information about a successful payment.
type SuccessfulPayment struct {

	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`

	// Total price in the smallest units of the currency (integer, not float/double). For example, for
	// a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the
	// number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int64 `json:"total_amount"`

	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id"`

	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info"`

	// Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

	// Provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}

// This object contains information about an incoming shipping query.
type ShippingQuery struct {

	// Unique query identifier
	ID string `json:"id"`

	// User who sent the query
	From *User `json:"from"`

	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// User specified shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {

	// Unique query identifier
	ID string `json:"id"`

	// User who sent the query
	From *User `json:"from"`

	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`

	// Total price in the smallest units of the currency (integer, not float/double). For example, for
	// a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the
	// number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int64 `json:"total_amount"`

	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id"`

	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info"`
}

// Telegram Passport is a unified authorization method for services that require personal
// identification. Users can upload their documents once, then instantly share their data with
// services that require real-world ID (finance, ICOs, etc.). Please see the manual for details.
// Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {

	// Array with information about documents and other Telegram Passport elements that was shared
	// with the bot
	Data []*EncryptedPassportElement `json:"data"`

	// Encrypted credentials required to decrypt the data
	Credentials *EncryptedCredentials `json:"credentials"`
}

// This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport
// files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {

	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different
	// bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// File size in bytes
	FileSize int64 `json:"file_size"`

	// Unix time when the file was uploaded
	FileDate int64 `json:"file_date"`
}

// Describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {

	// Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”,
	// “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”,
	// “passport_registration”, “temporary_registration”, “phone_number”, “email”.
	Type string `json:"type"`

	// Optional. Base64-encoded encrypted Telegram Passport element data provided by the user,
	// available for “personal_details”, “passport”, “driver_license”, “identity_card”,
	// “internal_passport” and “address” types. Can be decrypted and verified using the accompanying
	// EncryptedCredentials.
	Data string `json:"data"`

	// Optional. User's verified phone number, available only for “phone_number” type
	PhoneNumber string `json:"phone_number"`

	// Optional. User's verified email address, available only for “email” type
	Email string `json:"email"`

	// Optional. Array of encrypted files with documents provided by the user, available for
	// “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and
	// “temporary_registration” types. Files can be decrypted and verified using the accompanying
	// EncryptedCredentials.
	Files []*PassportFile `json:"files"`

	// Optional. Encrypted file with the front side of the document, provided by the user. Available
	// for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be
	// decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide *PassportFile `json:"front_side"`

	// Optional. Encrypted file with the reverse side of the document, provided by the user. Available
	// for “driver_license” and “identity_card”. The file can be decrypted and verified using the
	// accompanying EncryptedCredentials.
	ReverseSide *PassportFile `json:"reverse_side"`

	// Optional. Encrypted file with the selfie of the user holding a document, provided by the user;
	// available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file
	// can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie *PassportFile `json:"selfie"`

	// Optional. Array of encrypted files with translated versions of documents provided by the user.
	// Available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”,
	// “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and
	// “temporary_registration” types. Files can be decrypted and verified using the accompanying
	// EncryptedCredentials.
	Translation []*PassportFile `json:"translation"`

	// Base64-encoded element hash for using in PassportElementErrorUnspecified
	Hash string `json:"hash"`
}

// Describes data required for decrypting and authenticating EncryptedPassportElement . See the
// Telegram Passport Documentation for a complete description of the data decryption and
// authentication processes.
type EncryptedCredentials struct {

	// Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and
	// secrets required for EncryptedPassportElement decryption and authentication
	Data string `json:"data"`

	// Base64-encoded data hash for data authentication
	Hash string `json:"hash"`

	// Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
	Secret string `json:"secret"`
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
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the error, one of “personal_details”,
	// “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”
	Type string `json:"type"`

	// Name of the data field which has the error
	FieldName string `json:"field_name"`

	// Base64-encoded data hash
	DataHash string `json:"data_hash"`

	// Error message
	Message string `json:"message"`
}

// Represents an issue with the front side of a document. The error is considered resolved when
// the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {

	// Error source, must be front_side
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`

	// Base64-encoded hash of the file with the front side of the document
	FileHash string `json:"file_hash"`

	// Error message
	Message string `json:"message"`
}

// Represents an issue with the reverse side of a document. The error is considered resolved
// when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {

	// Error source, must be reverse_side
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the issue, one of “driver_license”,
	// “identity_card”
	Type string `json:"type"`

	// Base64-encoded hash of the file with the reverse side of the document
	FileHash string `json:"file_hash"`

	// Error message
	Message string `json:"message"`
}

// Represents an issue with the selfie with a document. The error is considered resolved when
// the file with the selfie changes.
type PassportElementErrorSelfie struct {

	// Error source, must be selfie
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`

	// Base64-encoded hash of the file with the selfie
	FileHash string `json:"file_hash"`

	// Error message
	Message string `json:"message"`
}

// Represents an issue with a document scan. The error is considered resolved when the file with
// the document scan changes.
type PassportElementErrorFile struct {

	// Error source, must be file
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the issue, one of “utility_bill”,
	// “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// Base64-encoded file hash
	FileHash string `json:"file_hash"`

	// Error message
	Message string `json:"message"`
}

// Represents an issue with a list of scans. The error is considered resolved when the list of
// files containing the scans changes.
type PassportElementErrorFiles struct {

	// Error source, must be files
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the issue, one of “utility_bill”,
	// “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`

	// Error message
	Message string `json:"message"`
}

// Represents an issue with one of the files that constitute the translation of a document. The
// error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {

	// Error source, must be translation_file
	Source string `json:"source"`

	// Type of element of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”,
	// “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// Base64-encoded file hash
	FileHash string `json:"file_hash"`

	// Error message
	Message string `json:"message"`
}

// Represents an issue with the translated version of a document. The error is considered
// resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {

	// Error source, must be translation_files
	Source string `json:"source"`

	// Type of element of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”,
	// “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`

	// Error message
	Message string `json:"message"`
}

// Represents an issue in an unspecified place. The error is considered resolved when new data
// is added.
type PassportElementErrorUnspecified struct {

	// Error source, must be unspecified
	Source string `json:"source"`

	// Type of element of the user's Telegram Passport which has the issue
	Type string `json:"type"`

	// Base64-encoded element hash
	ElementHash string `json:"element_hash"`

	// Error message
	Message string `json:"message"`
}

// This object represents a game. Use BotFather to create and edit games, their short names will
// act as unique identifiers.
type Game struct {

	// Title of the game
	Title string `json:"title"`

	// Description of the game
	Description string `json:"description"`

	// Photo that will be displayed in the game message in chats.
	Photo []*PhotoSize `json:"photo"`

	// Optional. Brief description of the game or high scores included in the game message. Can be
	// automatically edited to include current high scores for the game when the bot calls
	// setGameScore, or manually edited using editMessageText. 0-4096 characters.
	Text string `json:"text"`

	// Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	TextEntities []*MessageEntity `json:"text_entities"`

	// Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
	Animation *Animation `json:"animation"`
}

// A placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame struct {
}

// This object represents one row of the high scores table for a game.
type GameHighScore struct {

	// Position in high score table for the game
	Position int64 `json:"position"`

	// User
	User *User `json:"user"`

	// Score
	Score int64 `json:"score"`
}

// Oneof type fields are merged into one
// Merged fields of MessageOriginUser, MessageOriginHiddenUser, MessageOriginChat,
// MessageOriginChannel
type MessageOrigin struct {

	// Type of the message origin, always “channel”
	Type string `json:"type"`

	// Date the message was sent originally in Unix time
	Date int64 `json:"date"`

	// User that sent the message originally
	SenderUser *User `json:"sender_user"`

	// Name of the user that sent the message originally
	SenderUserName string `json:"sender_user_name"`

	// Chat that sent the message originally
	SenderChat *Chat `json:"sender_chat"`

	// Optional. Signature of the original post author
	AuthorSignature string `json:"author_signature"`

	// Channel chat to which the message was originally sent
	Chat *Chat `json:"chat"`

	// Unique message identifier inside the chat
	MessageID int64 `json:"message_id"`
}

// Merged fields of ReactionTypeEmoji, ReactionTypeCustomEmoji
type ReactionType struct {

	// Type of the reaction, always “custom_emoji”
	Type string `json:"type"`

	// Reaction emoji. Currently, it can be one of "", "", "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	// "", "", "", "", "", "", "", "", "", "", "", ""
	Emoji string `json:"emoji"`

	// Custom emoji identifier
	CustomEmojiID string `json:"custom_emoji_id"`
}

// Bot request and response types
// Request for API call 'getUpdates'
type GetUpdatesRequest struct {

	// Identifier of the first update to be returned. Must be greater by one than the highest among
	// the identifiers of previously received updates. By default, updates starting with the earliest
	// unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is
	// called with an offset higher than its update_id. The negative offset can be specified to
	// retrieve updates starting from -offset update from the end of the updates queue. All previous
	// updates will be forgotten.
	Offset int64 `json:"offset"`

	// Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to
	// 100.
	Limit int64 `json:"limit"`

	// Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be
	// positive, short polling should be used for testing purposes only.
	Timeout int64 `json:"timeout"`

	// A JSON-serialized list of the update types you want your bot to receive. For example, specify
	// ["message", "edited_channel_post", "callback_query"] to only receive updates of these types.
	// See Update for a complete list of available update types. Specify an empty list to receive all
	// update types except chat_member, message_reaction, and message_reaction_count (default). If not
	// specified, the previous setting will be used.Please note that this parameter doesn't affect
	// updates created before the call to the getUpdates, so unwanted updates may be received for a
	// short period of time.
	AllowedUpdates []string `json:"allowed_updates"`
}

// Response for API call 'getUpdates'
type GetUpdatesResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result []*Update `json:"result"`
}

// Request for API call 'setWebhook'
type SetWebhookRequest struct {

	// HTTPS URL to send updates to. Use an empty string to remove webhook integration
	URL string `json:"url"`

	// Upload your public key certificate so that the root certificate in use can be checked. See our
	// self-signed guide for details.
	Certificate *InputFile `json:"certificate"`

	// The fixed IP address which will be used to send webhook requests instead of the IP address
	// resolved through DNS
	IPAddress string `json:"ip_address"`

	// The maximum allowed number of simultaneous HTTPS connections to the webhook for update
	// delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and
	// higher values to increase your bot's throughput.
	MaxConnections int64 `json:"max_connections"`

	// A JSON-serialized list of the update types you want your bot to receive. For example, specify
	// ["message", "edited_channel_post", "callback_query"] to only receive updates of these types.
	// See Update for a complete list of available update types. Specify an empty list to receive all
	// update types except chat_member, message_reaction, and message_reaction_count (default). If not
	// specified, the previous setting will be used.Please note that this parameter doesn't affect
	// updates created before the call to the setWebhook, so unwanted updates may be received for a
	// short period of time.
	AllowedUpdates []string `json:"allowed_updates"`

	// Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates"`

	// A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook
	// request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is
	// useful to ensure that the request comes from a webhook set by you.
	SecretToken string `json:"secret_token"`
}

// Response for API call 'setWebhook'
type SetWebhookResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'deleteWebhook'
type DeleteWebhookRequest struct {

	// Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates"`
}

// Response for API call 'deleteWebhook'
type DeleteWebhookResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getWebhookInfo'
type GetWebhookInfoRequest struct {
}

// Response for API call 'getWebhookInfo'
type GetWebhookInfoResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getMe'
type GetMeRequest struct {
}

// Response for API call 'getMe'
type GetMeResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'logOut'
type LogOutRequest struct {
}

// Response for API call 'logOut'
type LogOutResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'close'
type CloseRequest struct {
}

// Response for API call 'close'
type CloseResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'sendMessage'
type SendMessageRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Text of the message to be sent, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// Mode for parsing entities in the message text. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// A JSON-serialized list of special entities that appear in message text, which can be specified
	// instead of parse_mode
	Entities []*MessageEntity `json:"entities"`

	// Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendMessage'
type SendMessageResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'forwardMessage'
type ForwardMessageRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Unique identifier for the chat where the original message was sent (or channel username in the
	// format @channelusername)
	FromChatID interface{} `json:"from_chat_id"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the forwarded message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Message identifier in the chat specified in from_chat_id
	MessageID int64 `json:"message_id"`
}

// Response for API call 'forwardMessage'
type ForwardMessageResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'forwardMessages'
type ForwardMessagesRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Unique identifier for the chat where the original messages were sent (or channel username in
	// the format @channelusername)
	FromChatID interface{} `json:"from_chat_id"`

	// Identifiers of 1-100 messages in the chat from_chat_id to forward. The identifiers must be
	// specified in a strictly increasing order.
	MessageIds []int64 `json:"message_ids"`

	// Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the forwarded messages from forwarding and saving
	ProtectContent bool `json:"protect_content"`
}

// Response for API call 'forwardMessages'
type ForwardMessagesResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result []*MessageId `json:"result"`
}

// Request for API call 'copyMessage'
type CopyMessageRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Unique identifier for the chat where the original message was sent (or channel username in the
	// format @channelusername)
	FromChatID interface{} `json:"from_chat_id"`

	// Message identifier in the chat specified in from_chat_id
	MessageID int64 `json:"message_id"`

	// New caption for media, 0-1024 characters after entities parsing. If not specified, the original
	// caption is kept
	Caption string `json:"caption"`

	// Mode for parsing entities in the new caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// A JSON-serialized list of special entities that appear in the new caption, which can be
	// specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'copyMessage'
type CopyMessageResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'copyMessages'
type CopyMessagesRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Unique identifier for the chat where the original messages were sent (or channel username in
	// the format @channelusername)
	FromChatID interface{} `json:"from_chat_id"`

	// Identifiers of 1-100 messages in the chat from_chat_id to copy. The identifiers must be
	// specified in a strictly increasing order.
	MessageIds []int64 `json:"message_ids"`

	// Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent messages from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Pass True to copy the messages without their captions
	RemoveCaption bool `json:"remove_caption"`
}

// Response for API call 'copyMessages'
type CopyMessagesResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result []*MessageId `json:"result"`
}

// Request for API call 'sendPhoto'
type SendPhotoRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or
	// upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The
	// photo's width and height must not exceed 10000 in total. Width and height ratio must be at most
	// 20. More information on Sending Files »
	Photo interface{} `json:"photo"`

	// Photo caption (may also be used when resending photos by file_id), 0-1024 characters after
	// entities parsing
	Caption string `json:"caption"`

	// Mode for parsing entities in the photo caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// A JSON-serialized list of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendPhoto'
type SendPhotoResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendAudio'
type SendAudioRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the
	// Internet, or upload a new one using multipart/form-data. More information on Sending Files »
	Audio interface{} `json:"audio"`

	// Audio caption, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Mode for parsing entities in the audio caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// A JSON-serialized list of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Duration of the audio in seconds
	Duration int64 `json:"duration"`

	// Performer
	Performer string `json:"performer"`

	// Track name
	Title string `json:"title"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you
	// can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data
	// under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendAudio'
type SendAudioResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendDocument'
type SendDocumentRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// File to send. Pass a file_id as String to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or
	// upload a new one using multipart/form-data. More information on Sending Files »
	Document interface{} `json:"document"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you
	// can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data
	// under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail"`

	// Document caption (may also be used when resending documents by file_id), 0-1024 characters
	// after entities parsing
	Caption string `json:"caption"`

	// Mode for parsing entities in the document caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// A JSON-serialized list of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Disables automatic server-side content type detection for files uploaded using
	// multipart/form-data
	DisableContentTypeDetection bool `json:"disable_content_type_detection"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendDocument'
type SendDocumentResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendVideo'
type SendVideoRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Video to send. Pass a file_id as String to send a video that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or
	// upload a new video using multipart/form-data. More information on Sending Files »
	Video interface{} `json:"video"`

	// Duration of sent video in seconds
	Duration int64 `json:"duration"`

	// Video width
	Width int64 `json:"width"`

	// Video height
	Height int64 `json:"height"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you
	// can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data
	// under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail"`

	// Video caption (may also be used when resending videos by file_id), 0-1024 characters after
	// entities parsing
	Caption string `json:"caption"`

	// Mode for parsing entities in the video caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// A JSON-serialized list of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler"`

	// Pass True if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendVideo'
type SendVideoResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendAnimation'
type SendAnimationRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Animation to send. Pass a file_id as String to send an animation that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the
	// Internet, or upload a new animation using multipart/form-data. More information on Sending
	// Files »
	Animation interface{} `json:"animation"`

	// Duration of sent animation in seconds
	Duration int64 `json:"duration"`

	// Animation width
	Width int64 `json:"width"`

	// Animation height
	Height int64 `json:"height"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you
	// can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data
	// under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail"`

	// Animation caption (may also be used when resending animation by file_id), 0-1024 characters
	// after entities parsing
	Caption string `json:"caption"`

	// Mode for parsing entities in the animation caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// A JSON-serialized list of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendAnimation'
type SendAnimationResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendVoice'
type SendVoiceRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or
	// upload a new one using multipart/form-data. More information on Sending Files »
	Voice interface{} `json:"voice"`

	// Voice message caption, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Mode for parsing entities in the voice message caption. See formatting options for more
	// details.
	ParseMode string `json:"parse_mode"`

	// A JSON-serialized list of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// Duration of the voice message in seconds
	Duration int64 `json:"duration"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendVoice'
type SendVoiceResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendVideoNote'
type SendVideoNoteRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Video note to send. Pass a file_id as String to send a video note that exists on the Telegram
	// servers (recommended) or upload a new video using multipart/form-data. More information on
	// Sending Files ». Sending video notes by a URL is currently unsupported
	VideoNote interface{} `json:"video_note"`

	// Duration of sent video in seconds
	Duration int64 `json:"duration"`

	// Video width and height, i.e. diameter of the video message
	Length int64 `json:"length"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's
	// width and height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you
	// can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data
	// under <file_attach_name>. More information on Sending Files »
	Thumbnail interface{} `json:"thumbnail"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendVideoNote'
type SendVideoNoteResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendLocation'
type SendLocationRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Latitude of the location
	Latitude float32 `json:"latitude"`

	// Longitude of the location
	Longitude float32 `json:"longitude"`

	// The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy"`

	// Period in seconds for which the location will be updated (see Live Locations, should be between
	// 60 and 86400.
	LivePeriod int64 `json:"live_period"`

	// For live locations, a direction in which the user is moving, in degrees. Must be between 1 and
	// 360 if specified.
	Heading int64 `json:"heading"`

	// For live locations, a maximum distance for proximity alerts about approaching another chat
	// member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int64 `json:"proximity_alert_radius"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendLocation'
type SendLocationResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendVenue'
type SendVenueRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Latitude of the venue
	Latitude float32 `json:"latitude"`

	// Longitude of the venue
	Longitude float32 `json:"longitude"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id"`

	// Foursquare type of the venue, if known. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type"`

	// Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id"`

	// Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendVenue'
type SendVenueResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendContact'
type SendContactRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Contact's last name
	LastName string `json:"last_name"`

	// Additional data about the contact in the form of a vCard, 0-2048 bytes
	Vcard string `json:"vcard"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendContact'
type SendContactResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendPoll'
type SendPollRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Poll question, 1-300 characters
	Question string `json:"question"`

	// A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
	Options []string `json:"options"`

	// True, if the poll needs to be anonymous, defaults to True
	IsAnonymous bool `json:"is_anonymous"`

	// Poll type, “quiz” or “regular”, defaults to “regular”
	Type string `json:"type"`

	// True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`

	// 0-based identifier of the correct answer option, required for polls in quiz mode
	CorrectOptionID int64 `json:"correct_option_id"`

	// Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a
	// quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	Explanation string `json:"explanation"`

	// Mode for parsing entities in the explanation. See formatting options for more details.
	ExplanationParseMode string `json:"explanation_parse_mode"`

	// A JSON-serialized list of special entities that appear in the poll explanation, which can be
	// specified instead of parse_mode
	ExplanationEntities []*MessageEntity `json:"explanation_entities"`

	// Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together
	// with close_date.
	OpenPeriod int64 `json:"open_period"`

	// Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5
	// and no more than 600 seconds in the future. Can't be used together with open_period.
	CloseDate int64 `json:"close_date"`

	// Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
	IsClosed bool `json:"is_closed"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendPoll'
type SendPollResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendDice'
type SendDiceRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Emoji on which the dice throw animation is based. Currently, must be one of “”, “”, “”, “”, “”,
	// or “”. Dice can have values 1-6 for “”, “” and “”, values 1-5 for “” and “”, and values 1-64
	// for “”. Defaults to “”
	Emoji string `json:"emoji"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendDice'
type SendDiceResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'sendChatAction'
type SendChatActionRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread; supergroups only
	MessageThreadID int64 `json:"message_thread_id"`

	// Type of action to broadcast. Choose one, depending on what the user is about to receive: typing
	// for text messages, upload_photo for photos, record_video or upload_video for videos,
	// record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker
	// for stickers, find_location for location data, record_video_note or upload_video_note for video
	// notes.
	Action string `json:"action"`
}

// Response for API call 'sendChatAction'
type SendChatActionResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setMessageReaction'
type SetMessageReactionRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Identifier of the target message. If the message belongs to a media group, the reaction is set
	// to the first non-deleted message in the group instead.
	MessageID int64 `json:"message_id"`

	// New list of reaction types to set on the message. Currently, as non-premium users, bots can set
	// up to one reaction per message. A custom emoji reaction can be used if it is either already
	// present on the message or explicitly allowed by chat administrators.
	Reaction []*ReactionType `json:"reaction"`

	// Pass True to set the reaction with a big animation
	IsBig bool `json:"is_big"`
}

// Response for API call 'setMessageReaction'
type SetMessageReactionResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getUserProfilePhotos'
type GetUserProfilePhotosRequest struct {

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// Sequential number of the first photo to be returned. By default, all photos are returned.
	Offset int64 `json:"offset"`

	// Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to
	// 100.
	Limit int64 `json:"limit"`
}

// Response for API call 'getUserProfilePhotos'
type GetUserProfilePhotosResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getFile'
type GetFileRequest struct {

	// File identifier to get information about
	FileID string `json:"file_id"`
}

// Response for API call 'getFile'
type GetFileResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *File `json:"result"`
}

// Request for API call 'banChatMember'
type BanChatMemberRequest struct {

	// Unique identifier for the target group or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// Date when the user will be unbanned; Unix time. If user is banned for more than 366 days or
	// less than 30 seconds from the current time they are considered to be banned forever. Applied
	// for supergroups and channels only.
	UntilDate int64 `json:"until_date"`

	// Pass True to delete all messages from the chat for the user that is being removed. If False,
	// the user will be able to see messages in the group that were sent before the user was removed.
	// Always True for supergroups and channels.
	RevokeMessages bool `json:"revoke_messages"`
}

// Response for API call 'banChatMember'
type BanChatMemberResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'unbanChatMember'
type UnbanChatMemberRequest struct {

	// Unique identifier for the target group or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// Do nothing if the user is not banned
	OnlyIfBanned bool `json:"only_if_banned"`
}

// Response for API call 'unbanChatMember'
type UnbanChatMemberResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'restrictChatMember'
type RestrictChatMemberRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// A JSON-serialized object for new user permissions
	Permissions *ChatPermissions `json:"permissions"`

	// Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and
	// can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios,
	// can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and
	// can_send_voice_notes permissions; the can_send_polls permission will imply the
	// can_send_messages permission.
	UseIndependentChatPermissions bool `json:"use_independent_chat_permissions"`

	// Date when restrictions will be lifted for the user; Unix time. If user is restricted for more
	// than 366 days or less than 30 seconds from the current time, they are considered to be
	// restricted forever
	UntilDate int64 `json:"until_date"`
}

// Response for API call 'restrictChatMember'
type RestrictChatMemberResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'promoteChatMember'
type PromoteChatMemberRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// Pass True if the administrator's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// Pass True if the administrator can access the chat event log, boost list in channels, see
	// channel members, report spam messages, see anonymous administrators in supergroups and ignore
	// slow mode. Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat"`

	// Pass True if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`

	// Pass True if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`

	// Pass True if the administrator can restrict, ban or unban chat members, or access supergroup
	// statistics
	CanRestrictMembers bool `json:"can_restrict_members"`

	// Pass True if the administrator can add new administrators with a subset of their own privileges
	// or demote administrators that they have promoted, directly or indirectly (promoted by
	// administrators that were appointed by him)
	CanPromoteMembers bool `json:"can_promote_members"`

	// Pass True if the administrator can change chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// Pass True if the administrator can invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// Pass True if the administrator can post messages in the channel, or access channel statistics;
	// channels only
	CanPostMessages bool `json:"can_post_messages"`

	// Pass True if the administrator can edit messages of other users and can pin messages; channels
	// only
	CanEditMessages bool `json:"can_edit_messages"`

	// Pass True if the administrator can pin messages, supergroups only
	CanPinMessages bool `json:"can_pin_messages"`

	// Pass True if the administrator can post stories in the channel; channels only
	CanPostStories bool `json:"can_post_stories"`

	// Pass True if the administrator can edit stories posted by other users; channels only
	CanEditStories bool `json:"can_edit_stories"`

	// Pass True if the administrator can delete stories posted by other users; channels only
	CanDeleteStories bool `json:"can_delete_stories"`

	// Pass True if the user is allowed to create, rename, close, and reopen forum topics, supergroups
	// only
	CanManageTopics bool `json:"can_manage_topics"`
}

// Response for API call 'promoteChatMember'
type PromoteChatMemberResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setChatAdministratorCustomTitle'
type SetChatAdministratorCustomTitleRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// New custom title for the administrator; 0-16 characters, emoji are not allowed
	CustomTitle string `json:"custom_title"`
}

// Response for API call 'setChatAdministratorCustomTitle'
type SetChatAdministratorCustomTitleResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'banChatSenderChat'
type BanChatSenderChatRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target sender chat
	SenderChatID int64 `json:"sender_chat_id"`
}

// Response for API call 'banChatSenderChat'
type BanChatSenderChatResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'unbanChatSenderChat'
type UnbanChatSenderChatRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target sender chat
	SenderChatID int64 `json:"sender_chat_id"`
}

// Response for API call 'unbanChatSenderChat'
type UnbanChatSenderChatResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setChatPermissions'
type SetChatPermissionsRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// A JSON-serialized object for new default chat permissions
	Permissions *ChatPermissions `json:"permissions"`

	// Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and
	// can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios,
	// can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and
	// can_send_voice_notes permissions; the can_send_polls permission will imply the
	// can_send_messages permission.
	UseIndependentChatPermissions bool `json:"use_independent_chat_permissions"`
}

// Response for API call 'setChatPermissions'
type SetChatPermissionsResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'exportChatInviteLink'
type ExportChatInviteLinkRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'exportChatInviteLink'
type ExportChatInviteLinkResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'createChatInviteLink'
type CreateChatInviteLinkRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Invite link name; 0-32 characters
	Name string `json:"name"`

	// Point in time (Unix timestamp) when the link will expire
	ExpireDate int64 `json:"expire_date"`

	// The maximum number of users that can be members of the chat simultaneously after joining the
	// chat via this invite link; 1-99999
	MemberLimit int64 `json:"member_limit"`

	// True, if users joining the chat via the link need to be approved by chat administrators. If
	// True, member_limit can't be specified
	CreatesJoinRequest bool `json:"creates_join_request"`
}

// Response for API call 'createChatInviteLink'
type CreateChatInviteLinkResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'editChatInviteLink'
type EditChatInviteLinkRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// The invite link to edit
	InviteLink string `json:"invite_link"`

	// Invite link name; 0-32 characters
	Name string `json:"name"`

	// Point in time (Unix timestamp) when the link will expire
	ExpireDate int64 `json:"expire_date"`

	// The maximum number of users that can be members of the chat simultaneously after joining the
	// chat via this invite link; 1-99999
	MemberLimit int64 `json:"member_limit"`

	// True, if users joining the chat via the link need to be approved by chat administrators. If
	// True, member_limit can't be specified
	CreatesJoinRequest bool `json:"creates_join_request"`
}

// Response for API call 'editChatInviteLink'
type EditChatInviteLinkResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'revokeChatInviteLink'
type RevokeChatInviteLinkRequest struct {

	// Unique identifier of the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// The invite link to revoke
	InviteLink string `json:"invite_link"`
}

// Response for API call 'revokeChatInviteLink'
type RevokeChatInviteLinkResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'approveChatJoinRequest'
type ApproveChatJoinRequestRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// Response for API call 'approveChatJoinRequest'
type ApproveChatJoinRequestResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'declineChatJoinRequest'
type DeclineChatJoinRequestRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// Response for API call 'declineChatJoinRequest'
type DeclineChatJoinRequestResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setChatPhoto'
type SetChatPhotoRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// New chat photo, uploaded using multipart/form-data
	Photo *InputFile `json:"photo"`
}

// Response for API call 'setChatPhoto'
type SetChatPhotoResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'deleteChatPhoto'
type DeleteChatPhotoRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'deleteChatPhoto'
type DeleteChatPhotoResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setChatTitle'
type SetChatTitleRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// New chat title, 1-128 characters
	Title string `json:"title"`
}

// Response for API call 'setChatTitle'
type SetChatTitleResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setChatDescription'
type SetChatDescriptionRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// New chat description, 0-255 characters
	Description string `json:"description"`
}

// Response for API call 'setChatDescription'
type SetChatDescriptionResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'pinChatMessage'
type PinChatMessageRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Identifier of a message to pin
	MessageID int64 `json:"message_id"`

	// Pass True if it is not necessary to send a notification to all chat members about the new
	// pinned message. Notifications are always disabled in channels and private chats.
	DisableNotification bool `json:"disable_notification"`
}

// Response for API call 'pinChatMessage'
type PinChatMessageResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'unpinChatMessage'
type UnpinChatMessageRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Identifier of a message to unpin. If not specified, the most recent pinned message (by sending
	// date) will be unpinned.
	MessageID int64 `json:"message_id"`
}

// Response for API call 'unpinChatMessage'
type UnpinChatMessageResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'unpinAllChatMessages'
type UnpinAllChatMessagesRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'unpinAllChatMessages'
type UnpinAllChatMessagesResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'leaveChat'
type LeaveChatRequest struct {

	// Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'leaveChat'
type LeaveChatResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getChat'
type GetChatRequest struct {

	// Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'getChat'
type GetChatResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getChatAdministrators'
type GetChatAdministratorsRequest struct {

	// Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'getChatAdministrators'
type GetChatAdministratorsResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result []*ChatMember `json:"result"`
}

// Request for API call 'getChatMemberCount'
type GetChatMemberCountRequest struct {

	// Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'getChatMemberCount'
type GetChatMemberCountResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getChatMember'
type GetChatMemberRequest struct {

	// Unique identifier for the target chat or username of the target supergroup or channel (in the
	// format @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// Response for API call 'getChatMember'
type GetChatMemberResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setChatStickerSet'
type SetChatStickerSetRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// Name of the sticker set to be set as the group sticker set
	StickerSetName string `json:"sticker_set_name"`
}

// Response for API call 'setChatStickerSet'
type SetChatStickerSetResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'deleteChatStickerSet'
type DeleteChatStickerSetRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'deleteChatStickerSet'
type DeleteChatStickerSetResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getForumTopicIconStickers'
type GetForumTopicIconStickersRequest struct {
}

// Response for API call 'getForumTopicIconStickers'
type GetForumTopicIconStickersResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result []*Sticker `json:"result"`
}

// Request for API call 'createForumTopic'
type CreateForumTopicRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// Topic name, 1-128 characters
	Name string `json:"name"`

	// Color of the topic icon in RGB format. Currently, must be one of 7322096 (0x6FB9F0), 16766590
	// (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047
	// (0xFB6F5F)
	IconColor int64 `json:"icon_color"`

	// Unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to
	// get all allowed custom emoji identifiers.
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
}

// Response for API call 'createForumTopic'
type CreateForumTopicResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'editForumTopic'
type EditForumTopicRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread of the forum topic
	MessageThreadID int64 `json:"message_thread_id"`

	// New topic name, 0-128 characters. If not specified or empty, the current name of the topic will
	// be kept
	Name string `json:"name"`

	// New unique identifier of the custom emoji shown as the topic icon. Use
	// getForumTopicIconStickers to get all allowed custom emoji identifiers. Pass an empty string to
	// remove the icon. If not specified, the current icon will be kept
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
}

// Response for API call 'editForumTopic'
type EditForumTopicResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'closeForumTopic'
type CloseForumTopicRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread of the forum topic
	MessageThreadID int64 `json:"message_thread_id"`
}

// Response for API call 'closeForumTopic'
type CloseForumTopicResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'reopenForumTopic'
type ReopenForumTopicRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread of the forum topic
	MessageThreadID int64 `json:"message_thread_id"`
}

// Response for API call 'reopenForumTopic'
type ReopenForumTopicResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'deleteForumTopic'
type DeleteForumTopicRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread of the forum topic
	MessageThreadID int64 `json:"message_thread_id"`
}

// Response for API call 'deleteForumTopic'
type DeleteForumTopicResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'unpinAllForumTopicMessages'
type UnpinAllForumTopicMessagesRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread of the forum topic
	MessageThreadID int64 `json:"message_thread_id"`
}

// Response for API call 'unpinAllForumTopicMessages'
type UnpinAllForumTopicMessagesResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'editGeneralForumTopic'
type EditGeneralForumTopicRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`

	// New topic name, 1-128 characters
	Name string `json:"name"`
}

// Response for API call 'editGeneralForumTopic'
type EditGeneralForumTopicResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'closeGeneralForumTopic'
type CloseGeneralForumTopicRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'closeGeneralForumTopic'
type CloseGeneralForumTopicResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'reopenGeneralForumTopic'
type ReopenGeneralForumTopicRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'reopenGeneralForumTopic'
type ReopenGeneralForumTopicResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'hideGeneralForumTopic'
type HideGeneralForumTopicRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'hideGeneralForumTopic'
type HideGeneralForumTopicResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'unhideGeneralForumTopic'
type UnhideGeneralForumTopicRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'unhideGeneralForumTopic'
type UnhideGeneralForumTopicResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'unpinAllGeneralForumTopicMessages'
type UnpinAllGeneralForumTopicMessagesRequest struct {

	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @supergroupusername)
	ChatID interface{} `json:"chat_id"`
}

// Response for API call 'unpinAllGeneralForumTopicMessages'
type UnpinAllGeneralForumTopicMessagesResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'answerCallbackQuery'
type AnswerCallbackQueryRequest struct {

	// Unique identifier for the query to be answered
	CallbackQueryID string `json:"callback_query_id"`

	// Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	Text string `json:"text"`

	// If True, an alert will be shown by the client instead of a notification at the top of the chat
	// screen. Defaults to false.
	ShowAlert bool `json:"show_alert"`

	// URL that will be opened by the user's client. If you have created a Game and accepted the
	// conditions via @BotFather, specify the URL that opens your game - note that this will only work
	// if the query comes from a callback_game button.Otherwise, you may use links like
	// t.me/your_bot?start=XXXX that open your bot with a parameter.
	URL string `json:"url"`

	// The maximum amount of time in seconds that the result of the callback query may be cached
	// client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime int64 `json:"cache_time"`
}

// Response for API call 'answerCallbackQuery'
type AnswerCallbackQueryResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getUserChatBoosts'
type GetUserChatBoostsRequest struct {

	// Unique identifier for the chat or username of the channel (in the format @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`
}

// Response for API call 'getUserChatBoosts'
type GetUserChatBoostsResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setMyCommands'
type SetMyCommandsRequest struct {

	// A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100
	// commands can be specified.
	Commands []*BotCommand `json:"commands"`

	// A JSON-serialized object, describing scope of users for which the commands are relevant.
	// Defaults to BotCommandScopeDefault.
	Scope *BotCommandScope `json:"scope"`

	// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the
	// given scope, for whose language there are no dedicated commands
	LanguageCode string `json:"language_code"`
}

// Response for API call 'setMyCommands'
type SetMyCommandsResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'deleteMyCommands'
type DeleteMyCommandsRequest struct {

	// A JSON-serialized object, describing scope of users for which the commands are relevant.
	// Defaults to BotCommandScopeDefault.
	Scope *BotCommandScope `json:"scope"`

	// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the
	// given scope, for whose language there are no dedicated commands
	LanguageCode string `json:"language_code"`
}

// Response for API call 'deleteMyCommands'
type DeleteMyCommandsResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getMyCommands'
type GetMyCommandsRequest struct {

	// A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	Scope *BotCommandScope `json:"scope"`

	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code"`
}

// Response for API call 'getMyCommands'
type GetMyCommandsResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result []*BotCommand `json:"result"`
}

// Request for API call 'setMyName'
type SetMyNameRequest struct {

	// New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given
	// language.
	Name string `json:"name"`

	// A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose
	// language there is no dedicated name.
	LanguageCode string `json:"language_code"`
}

// Response for API call 'setMyName'
type SetMyNameResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getMyName'
type GetMyNameRequest struct {

	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code"`
}

// Response for API call 'getMyName'
type GetMyNameResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setMyDescription'
type SetMyDescriptionRequest struct {

	// New bot description; 0-512 characters. Pass an empty string to remove the dedicated description
	// for the given language.
	Description string `json:"description"`

	// A two-letter ISO 639-1 language code. If empty, the description will be applied to all users
	// for whose language there is no dedicated description.
	LanguageCode string `json:"language_code"`
}

// Response for API call 'setMyDescription'
type SetMyDescriptionResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getMyDescription'
type GetMyDescriptionRequest struct {

	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code"`
}

// Response for API call 'getMyDescription'
type GetMyDescriptionResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setMyShortDescription'
type SetMyShortDescriptionRequest struct {

	// New short description for the bot; 0-120 characters. Pass an empty string to remove the
	// dedicated short description for the given language.
	ShortDescription string `json:"short_description"`

	// A two-letter ISO 639-1 language code. If empty, the short description will be applied to all
	// users for whose language there is no dedicated short description.
	LanguageCode string `json:"language_code"`
}

// Response for API call 'setMyShortDescription'
type SetMyShortDescriptionResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getMyShortDescription'
type GetMyShortDescriptionRequest struct {

	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code"`
}

// Response for API call 'getMyShortDescription'
type GetMyShortDescriptionResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setChatMenuButton'
type SetChatMenuButtonRequest struct {

	// Unique identifier for the target private chat. If not specified, default bot's menu button will
	// be changed
	ChatID int64 `json:"chat_id"`

	// A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
	MenuButton *MenuButton `json:"menu_button"`
}

// Response for API call 'setChatMenuButton'
type SetChatMenuButtonResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getChatMenuButton'
type GetChatMenuButtonRequest struct {

	// Unique identifier for the target private chat. If not specified, default bot's menu button will
	// be returned
	ChatID int64 `json:"chat_id"`
}

// Response for API call 'getChatMenuButton'
type GetChatMenuButtonResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setMyDefaultAdministratorRights'
type SetMyDefaultAdministratorRightsRequest struct {

	// A JSON-serialized object describing new default administrator rights. If not specified, the
	// default administrator rights will be cleared.
	Rights *ChatAdministratorRights `json:"rights"`

	// Pass True to change the default administrator rights of the bot in channels. Otherwise, the
	// default administrator rights of the bot for groups and supergroups will be changed.
	ForChannels bool `json:"for_channels"`
}

// Response for API call 'setMyDefaultAdministratorRights'
type SetMyDefaultAdministratorRightsResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getMyDefaultAdministratorRights'
type GetMyDefaultAdministratorRightsRequest struct {

	// Pass True to get default administrator rights of the bot in channels. Otherwise, default
	// administrator rights of the bot for groups and supergroups will be returned.
	ForChannels bool `json:"for_channels"`
}

// Response for API call 'getMyDefaultAdministratorRights'
type GetMyDefaultAdministratorRightsResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'editMessageText'
type EditMessageTextRequest struct {

	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id"`

	// New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// Mode for parsing entities in the message text. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// A JSON-serialized list of special entities that appear in message text, which can be specified
	// instead of parse_mode
	Entities []*MessageEntity `json:"entities"`

	// Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// Response for API call 'editMessageText'
type EditMessageTextResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'editMessageCaption'
type EditMessageCaptionRequest struct {

	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id"`

	// New caption of the message, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Mode for parsing entities in the message caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// A JSON-serialized list of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// Response for API call 'editMessageCaption'
type EditMessageCaptionResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'editMessageMedia'
type EditMessageMediaRequest struct {

	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id"`

	// A JSON-serialized object for a new media content of the message
	Media *InputMedia `json:"media"`

	// A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// Response for API call 'editMessageMedia'
type EditMessageMediaResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'editMessageLiveLocation'
type EditMessageLiveLocationRequest struct {

	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id"`

	// Latitude of new location
	Latitude float32 `json:"latitude"`

	// Longitude of new location
	Longitude float32 `json:"longitude"`

	// The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy"`

	// Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading int64 `json:"heading"`

	// The maximum distance for proximity alerts about approaching another chat member, in meters.
	// Must be between 1 and 100000 if specified.
	ProximityAlertRadius int64 `json:"proximity_alert_radius"`

	// A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// Response for API call 'editMessageLiveLocation'
type EditMessageLiveLocationResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'stopMessageLiveLocation'
type StopMessageLiveLocationRequest struct {

	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Required if inline_message_id is not specified. Identifier of the message with live location to
	// stop
	MessageID int64 `json:"message_id"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id"`

	// A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// Response for API call 'stopMessageLiveLocation'
type StopMessageLiveLocationResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'editMessageReplyMarkup'
type EditMessageReplyMarkupRequest struct {

	// Required if inline_message_id is not specified. Unique identifier for the target chat or
	// username of the target channel (in the format @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// Response for API call 'editMessageReplyMarkup'
type EditMessageReplyMarkupResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'stopPoll'
type StopPollRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Identifier of the original message with the poll
	MessageID int64 `json:"message_id"`

	// A JSON-serialized object for a new message inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// Response for API call 'stopPoll'
type StopPollResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Poll `json:"result"`
}

// Request for API call 'deleteMessage'
type DeleteMessageRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Identifier of the message to delete
	MessageID int64 `json:"message_id"`
}

// Response for API call 'deleteMessage'
type DeleteMessageResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'deleteMessages'
type DeleteMessagesRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Identifiers of 1-100 messages to delete. See deleteMessage for limitations on which messages
	// can be deleted
	MessageIds []int64 `json:"message_ids"`
}

// Response for API call 'deleteMessages'
type DeleteMessagesResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'sendSticker'
type SendStickerRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the
	// Internet, or upload a new .WEBP or .TGS sticker using multipart/form-data. More information on
	// Sending Files ». Video stickers can only be sent by a file_id. Animated stickers can't be sent
	// via an HTTP URL.
	Sticker interface{} `json:"sticker"`

	// Emoji associated with the sticker; only for just uploaded stickers
	Emoji string `json:"emoji"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply
	// keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup"`
}

// Response for API call 'sendSticker'
type SendStickerResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'getStickerSet'
type GetStickerSetRequest struct {

	// Name of the sticker set
	Name string `json:"name"`
}

// Response for API call 'getStickerSet'
type GetStickerSetResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *StickerSet `json:"result"`
}

// Request for API call 'getCustomEmojiStickers'
type GetCustomEmojiStickersRequest struct {

	// List of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
	CustomEmojiIds []string `json:"custom_emoji_ids"`
}

// Response for API call 'getCustomEmojiStickers'
type GetCustomEmojiStickersResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result []*Sticker `json:"result"`
}

// Request for API call 'uploadStickerFile'
type UploadStickerFileRequest struct {

	// User identifier of sticker file owner
	UserID int64 `json:"user_id"`

	// A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format. See
	// https://core.telegram.org/stickers for technical requirements. More information on Sending
	// Files »
	Sticker *InputFile `json:"sticker"`

	// Format of the sticker, must be one of “static”, “animated”, “video”
	StickerFormat string `json:"sticker_format"`
}

// Response for API call 'uploadStickerFile'
type UploadStickerFileResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'createNewStickerSet'
type CreateNewStickerSetRequest struct {

	// User identifier of created sticker set owner
	UserID int64 `json:"user_id"`

	// Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain
	// only English letters, digits and underscores. Must begin with a letter, can't contain
	// consecutive underscores and must end in "_by_<bot_username>". <bot_username> is case
	// insensitive. 1-64 characters.
	Name string `json:"name"`

	// Sticker set title, 1-64 characters
	Title string `json:"title"`

	// A JSON-serialized list of 1-50 initial stickers to be added to the sticker set
	Stickers []*InputSticker `json:"stickers"`

	// Format of stickers in the set, must be one of “static”, “animated”, “video”
	StickerFormat string `json:"sticker_format"`

	// Type of stickers in the set, pass “regular”, “mask”, or “custom_emoji”. By default, a regular
	// sticker set is created.
	StickerType string `json:"sticker_type"`

	// Pass True if stickers in the sticker set must be repainted to the color of text when used in
	// messages, the accent color if used as emoji status, white on chat photos, or another
	// appropriate color based on context; for custom emoji sticker sets only
	NeedsRepainting bool `json:"needs_repainting"`
}

// Response for API call 'createNewStickerSet'
type CreateNewStickerSetResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'addStickerToSet'
type AddStickerToSetRequest struct {

	// User identifier of sticker set owner
	UserID int64 `json:"user_id"`

	// Sticker set name
	Name string `json:"name"`

	// A JSON-serialized object with information about the added sticker. If exactly the same sticker
	// had already been added to the set, then the set isn't changed.
	Sticker *InputSticker `json:"sticker"`
}

// Response for API call 'addStickerToSet'
type AddStickerToSetResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setStickerPositionInSet'
type SetStickerPositionInSetRequest struct {

	// File identifier of the sticker
	Sticker string `json:"sticker"`

	// New sticker position in the set, zero-based
	Position int64 `json:"position"`
}

// Response for API call 'setStickerPositionInSet'
type SetStickerPositionInSetResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'deleteStickerFromSet'
type DeleteStickerFromSetRequest struct {

	// File identifier of the sticker
	Sticker string `json:"sticker"`
}

// Response for API call 'deleteStickerFromSet'
type DeleteStickerFromSetResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setStickerEmojiList'
type SetStickerEmojiListRequest struct {

	// File identifier of the sticker
	Sticker string `json:"sticker"`

	// A JSON-serialized list of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list"`
}

// Response for API call 'setStickerEmojiList'
type SetStickerEmojiListResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setStickerKeywords'
type SetStickerKeywordsRequest struct {

	// File identifier of the sticker
	Sticker string `json:"sticker"`

	// A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to 64
	// characters
	Keywords []string `json:"keywords"`
}

// Response for API call 'setStickerKeywords'
type SetStickerKeywordsResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setStickerMaskPosition'
type SetStickerMaskPositionRequest struct {

	// File identifier of the sticker
	Sticker string `json:"sticker"`

	// A JSON-serialized object with the position where the mask should be placed on faces. Omit the
	// parameter to remove the mask position.
	MaskPosition *MaskPosition `json:"mask_position"`
}

// Response for API call 'setStickerMaskPosition'
type SetStickerMaskPositionResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setStickerSetTitle'
type SetStickerSetTitleRequest struct {

	// Sticker set name
	Name string `json:"name"`

	// Sticker set title, 1-64 characters
	Title string `json:"title"`
}

// Response for API call 'setStickerSetTitle'
type SetStickerSetTitleResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setStickerSetThumbnail'
type SetStickerSetThumbnailRequest struct {

	// Sticker set name
	Name string `json:"name"`

	// User identifier of the sticker set owner
	UserID int64 `json:"user_id"`

	// A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a width
	// and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes in size
	// (see https://core.telegram.org/stickers#animated-sticker-requirements for animated sticker
	// technical requirements), or a WEBM video with the thumbnail up to 32 kilobytes in size; see
	// https://core.telegram.org/stickers#video-sticker-requirements for video sticker technical
	// requirements. Pass a file_id as a String to send a file that already exists on the Telegram
	// servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a
	// new one using multipart/form-data. More information on Sending Files ». Animated and video
	// sticker set thumbnails can't be uploaded via HTTP URL. If omitted, then the thumbnail is
	// dropped and the first sticker is used as the thumbnail.
	Thumbnail interface{} `json:"thumbnail"`
}

// Response for API call 'setStickerSetThumbnail'
type SetStickerSetThumbnailResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setCustomEmojiStickerSetThumbnail'
type SetCustomEmojiStickerSetThumbnailRequest struct {

	// Sticker set name
	Name string `json:"name"`

	// Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the
	// thumbnail and use the first sticker as the thumbnail.
	CustomEmojiID string `json:"custom_emoji_id"`
}

// Response for API call 'setCustomEmojiStickerSetThumbnail'
type SetCustomEmojiStickerSetThumbnailResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'deleteStickerSet'
type DeleteStickerSetRequest struct {

	// Sticker set name
	Name string `json:"name"`
}

// Response for API call 'deleteStickerSet'
type DeleteStickerSetResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'answerInlineQuery'
type AnswerInlineQueryRequest struct {

	// Unique identifier for the answered query
	InlineQueryID string `json:"inline_query_id"`

	// A JSON-serialized array of results for the inline query
	Results []*InlineQueryResult `json:"results"`

	// The maximum amount of time in seconds that the result of the inline query may be cached on the
	// server. Defaults to 300.
	CacheTime int64 `json:"cache_time"`

	// Pass True if results may be cached on the server side only for the user that sent the query. By
	// default, results may be returned to any user who sends the same query.
	IsPersonal bool `json:"is_personal"`

	// Pass the offset that a client should send in the next query with the same text to receive more
	// results. Pass an empty string if there are no more results or if you don't support pagination.
	// Offset length can't exceed 64 bytes.
	NextOffset string `json:"next_offset"`

	// A JSON-serialized object describing a button to be shown above inline query results
	Button *InlineQueryResultsButton `json:"button"`
}

// Response for API call 'answerInlineQuery'
type AnswerInlineQueryResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'answerWebAppQuery'
type AnswerWebAppQueryRequest struct {

	// Unique identifier for the query to be answered
	WebAppQueryID string `json:"web_app_query_id"`

	// A JSON-serialized object describing the message to be sent
	Result *InlineQueryResult `json:"result"`
}

// Response for API call 'answerWebAppQuery'
type AnswerWebAppQueryResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *SentWebAppMessage `json:"result"`
}

// Request for API call 'sendInvoice'
type SendInvoiceRequest struct {

	// Unique identifier for the target chat or username of the target channel (in the format
	// @channelusername)
	ChatID interface{} `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Product name, 1-32 characters
	Title string `json:"title"`

	// Product description, 1-255 characters
	Description string `json:"description"`

	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your
	// internal processes.
	Payload string `json:"payload"`

	// Payment provider token, obtained via @BotFather
	ProviderToken string `json:"provider_token"`

	// Three-letter ISO 4217 currency code, see more on currencies
	Currency string `json:"currency"`

	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount,
	// delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices"`

	// The maximum accepted amount for tips in the smallest units of the currency (integer, not
	// float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the
	// exp parameter in currencies.json, it shows the number of digits past the decimal point for each
	// currency (2 for the majority of currencies). Defaults to 0
	MaxTipAmount int64 `json:"max_tip_amount"`

	// A JSON-serialized array of suggested amounts of tips in the smallest units of the currency
	// (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested
	// tip amounts must be positive, passed in a strictly increased order and must not exceed
	// max_tip_amount.
	SuggestedTipAmounts []int64 `json:"suggested_tip_amounts"`

	// Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a
	// Pay button, allowing multiple users to pay directly from the forwarded message, using the same
	// invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep
	// link to the bot (instead of a Pay button), with the value used as the start parameter
	StartParameter string `json:"start_parameter"`

	// JSON-serialized data about the invoice, which will be shared with the payment provider. A
	// detailed description of required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data"`

	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for
	// a service. People like it better when they see what they are paying for.
	PhotoURL string `json:"photo_url"`

	// Photo size in bytes
	PhotoSize int64 `json:"photo_size"`

	// Photo width
	PhotoWidth int64 `json:"photo_width"`

	// Photo height
	PhotoHeight int64 `json:"photo_height"`

	// Pass True if you require the user's full name to complete the order
	NeedName bool `json:"need_name"`

	// Pass True if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number"`

	// Pass True if you require the user's email address to complete the order
	NeedEmail bool `json:"need_email"`

	// Pass True if you require the user's shipping address to complete the order
	NeedShippingAddress bool `json:"need_shipping_address"`

	// Pass True if the user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider"`

	// Pass True if the user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider"`

	// Pass True if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be
	// shown. If not empty, the first button must be a Pay button.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// Response for API call 'sendInvoice'
type SendInvoiceResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'createInvoiceLink'
type CreateInvoiceLinkRequest struct {

	// Product name, 1-32 characters
	Title string `json:"title"`

	// Product description, 1-255 characters
	Description string `json:"description"`

	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your
	// internal processes.
	Payload string `json:"payload"`

	// Payment provider token, obtained via BotFather
	ProviderToken string `json:"provider_token"`

	// Three-letter ISO 4217 currency code, see more on currencies
	Currency string `json:"currency"`

	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount,
	// delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices"`

	// The maximum accepted amount for tips in the smallest units of the currency (integer, not
	// float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the
	// exp parameter in currencies.json, it shows the number of digits past the decimal point for each
	// currency (2 for the majority of currencies). Defaults to 0
	MaxTipAmount int64 `json:"max_tip_amount"`

	// A JSON-serialized array of suggested amounts of tips in the smallest units of the currency
	// (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested
	// tip amounts must be positive, passed in a strictly increased order and must not exceed
	// max_tip_amount.
	SuggestedTipAmounts []int64 `json:"suggested_tip_amounts"`

	// JSON-serialized data about the invoice, which will be shared with the payment provider. A
	// detailed description of required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data"`

	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for
	// a service.
	PhotoURL string `json:"photo_url"`

	// Photo size in bytes
	PhotoSize int64 `json:"photo_size"`

	// Photo width
	PhotoWidth int64 `json:"photo_width"`

	// Photo height
	PhotoHeight int64 `json:"photo_height"`

	// Pass True if you require the user's full name to complete the order
	NeedName bool `json:"need_name"`

	// Pass True if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number"`

	// Pass True if you require the user's email address to complete the order
	NeedEmail bool `json:"need_email"`

	// Pass True if you require the user's shipping address to complete the order
	NeedShippingAddress bool `json:"need_shipping_address"`

	// Pass True if the user's phone number should be sent to the provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider"`

	// Pass True if the user's email address should be sent to the provider
	SendEmailToProvider bool `json:"send_email_to_provider"`

	// Pass True if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible"`
}

// Response for API call 'createInvoiceLink'
type CreateInvoiceLinkResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'answerShippingQuery'
type AnswerShippingQueryRequest struct {

	// Unique identifier for the query to be answered
	ShippingQueryID string `json:"shipping_query_id"`

	// Pass True if delivery to the specified address is possible and False if there are any problems
	// (for example, if delivery to the specified address is not possible)
	Ok bool `json:"ok"`

	// Required if ok is True. A JSON-serialized array of available shipping options.
	ShippingOptions []*ShippingOption `json:"shipping_options"`

	// Required if ok is False. Error message in human readable form that explains why it is
	// impossible to complete the order (e.g. "Sorry, delivery to your desired address is
	// unavailable'). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message"`
}

// Response for API call 'answerShippingQuery'
type AnswerShippingQueryResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'answerPreCheckoutQuery'
type AnswerPreCheckoutQueryRequest struct {

	// Unique identifier for the query to be answered
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`

	// Specify True if everything is alright (goods are available, etc.) and the bot is ready to
	// proceed with the order. Use False if there are any problems.
	Ok bool `json:"ok"`

	// Required if ok is False. Error message in human readable form that explains the reason for
	// failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing
	// black T-shirts while you were busy filling out your payment details. Please choose a different
	// color or garment!"). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message"`
}

// Response for API call 'answerPreCheckoutQuery'
type AnswerPreCheckoutQueryResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'setPassportDataErrors'
type SetPassportDataErrorsRequest struct {

	// User identifier
	UserID int64 `json:"user_id"`

	// A JSON-serialized array describing the errors
	Errors []*PassportElementError `json:"errors"`
}

// Response for API call 'setPassportDataErrors'
type SetPassportDataErrorsResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'sendGame'
type SendGameRequest struct {

	// Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups
	// only
	MessageThreadID int64 `json:"message_thread_id"`

	// Short name of the game, serves as the unique identifier for the game. Set up your games via
	// @BotFather.
	GameShortName string `json:"game_short_name"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`

	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters"`

	// A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be
	// shown. If not empty, the first button must launch the game.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// Response for API call 'sendGame'
type SendGameResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result *Message `json:"result"`
}

// Request for API call 'setGameScore'
type SetGameScoreRequest struct {

	// User identifier
	UserID int64 `json:"user_id"`

	// New score, must be non-negative
	Score int64 `json:"score"`

	// Pass True if the high score is allowed to decrease. This can be useful when fixing mistakes or
	// banning cheaters
	Force bool `json:"force"`

	// Pass True if the game message should not be automatically edited to include the current
	// scoreboard
	DisableEditMessage bool `json:"disable_edit_message"`

	// Required if inline_message_id is not specified. Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	// Required if inline_message_id is not specified. Identifier of the sent message
	MessageID int64 `json:"message_id"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id"`
}

// Response for API call 'setGameScore'
type SetGameScoreResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`
}

// Request for API call 'getGameHighScores'
type GetGameHighScoresRequest struct {

	// Target user id
	UserID int64 `json:"user_id"`

	// Required if inline_message_id is not specified. Unique identifier for the target chat
	ChatID int64 `json:"chat_id"`

	// Required if inline_message_id is not specified. Identifier of the sent message
	MessageID int64 `json:"message_id"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID string `json:"inline_message_id"`
}

// Response for API call 'getGameHighScores'
type GetGameHighScoresResponse struct {

	// Raw response from the server
	Raw []byte `json:"raw"`

	// Decoded response from the server
	Result []*GameHighScore `json:"result"`
}

// Bot interface
type TelegramBot interface {
	// Use this method to receive incoming updates using long polling ( wiki ). Returns an Array of
	// Update objects.
	GetUpdates(request *GetUpdatesRequest) (*GetUpdatesResponse, error)

	// Notes 1. This method will not work if an outgoing webhook is set up. 2. In order to avoid
	// getting duplicate updates, recalculate offset after each server response.    Use this method to
	// specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update
	// for the bot, we will send an HTTPS POST request to the specified URL, containing a
	// JSON-serialized Update . In case of an unsuccessful request, we will give up after a reasonable
	// amount of attempts. Returns True on success.  If you'd like to make sure that the webhook was
	// set by you, you can specify secret data in the parameter secret_token . If specified, the
	// request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as
	// content.
	SetWebhook(request *SetWebhookRequest) (*SetWebhookResponse, error)

	// Notes 1. You will not be able to receive updates using getUpdates for as long as an outgoing
	// webhook is set up. 2. To use a self-signed certificate, you need to upload your public key
	// certificate using certificate parameter. Please upload as InputFile, sending a String will not
	// work. 3. Ports currently supported for webhooks : 443, 80, 88, 8443 .  If you're having any
	// trouble setting up webhooks, please check out this amazing guide to webhooks .    Use this
	// method to remove webhook integration if you decide to switch back to getUpdates . Returns True
	// on success.
	DeleteWebhook(request *DeleteWebhookRequest) (*DeleteWebhookResponse, error)

	// Use this method to get current webhook status. Requires no parameters. On success, returns a
	// WebhookInfo object. If the bot is using getUpdates , will return an object with the url field
	// empty.
	GetWebhookInfo(request *GetWebhookInfoRequest) (*GetWebhookInfoResponse, error)

	// A simple method for testing your bot's authentication token. Requires no parameters. Returns
	// basic information about the bot in form of a User object.
	GetMe(request *GetMeRequest) (*GetMeResponse, error)

	// Use this method to log out from the cloud Bot API server before launching the bot locally. You
	// must log out the bot before running it locally, otherwise there is no guarantee that the bot
	// will receive updates. After a successful call, you can immediately log in on a local server,
	// but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on
	// success. Requires no parameters.
	LogOut(request *LogOutRequest) (*LogOutResponse, error)

	// Use this method to close the bot instance before moving it from one local server to another.
	// You need to delete the webhook before calling this method to ensure that the bot isn't launched
	// again after server restart. The method will return error 429 in the first 10 minutes after the
	// bot is launched. Returns True on success. Requires no parameters.
	Close(request *CloseRequest) (*CloseResponse, error)

	//  Use this method to send text messages. On success, the sent Message is returned.
	SendMessage(request *SendMessageRequest) (*SendMessageResponse, error)

	// Use this method to forward messages of any kind. Service messages and messages with protected
	// content can't be forwarded. On success, the sent Message is returned.
	ForwardMessage(request *ForwardMessageRequest) (*ForwardMessageResponse, error)

	// Use this method to forward multiple messages of any kind. If some of the specified messages
	// can't be found or forwarded, they are skipped. Service messages and messages with protected
	// content can't be forwarded. Album grouping is kept for forwarded messages. On success, an array
	// of MessageId of the sent messages is returned.
	ForwardMessages(request *ForwardMessagesRequest) (*ForwardMessagesResponse, error)

	// Use this method to copy messages of any kind. Service messages, giveaway messages, giveaway
	// winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the
	// value of the field correct_option_id is known to the bot. The method is analogous to the method
	// forwardMessage , but the copied message doesn't have a link to the original message. Returns
	// the MessageId of the sent message on success.
	CopyMessage(request *CopyMessageRequest) (*CopyMessageResponse, error)

	// Use this method to copy messages of any kind. If some of the specified messages can't be
	// found or copied, they are skipped. Service messages, giveaway messages, giveaway winners
	// messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of
	// the field correct_option_id is known to the bot. The method is analogous to the method
	// forwardMessages , but the copied messages don't have a link to the original message. Album
	// grouping is kept for copied messages. On success, an array of MessageId of the sent messages is
	// returned.
	CopyMessages(request *CopyMessagesRequest) (*CopyMessagesResponse, error)

	//   Use this method to send photos. On success, the sent Message is returned.
	SendPhoto(request *SendPhotoRequest) (*SendPhotoResponse, error)

	// Use this method to send audio files, if you want Telegram clients to display them in the
	// music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is
	// returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed
	// in the future.  For sending voice messages, use the sendVoice method instead.
	SendAudio(request *SendAudioRequest) (*SendAudioResponse, error)

	// Use this method to send general files. On success, the sent Message is returned. Bots can
	// currently send files of any type of up to 50 MB in size, this limit may be changed in the
	// future.
	SendDocument(request *SendDocumentRequest) (*SendDocumentResponse, error)

	// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may
	// be sent as Document ). On success, the sent Message is returned. Bots can currently send video
	// files of up to 50 MB in size, this limit may be changed in the future.
	SendVideo(request *SendVideoRequest) (*SendVideoResponse, error)

	// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On
	// success, the sent Message is returned. Bots can currently send animation files of up to 50 MB
	// in size, this limit may be changed in the future.
	SendAnimation(request *SendAnimationRequest) (*SendAnimationResponse, error)

	// Use this method to send audio files, if you want Telegram clients to display the file as a
	// playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS
	// (other formats may be sent as Audio or Document ). On success, the sent Message is returned.
	// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the
	// future.
	SendVoice(request *SendVoiceRequest) (*SendVoiceResponse, error)

	// As of v.4.0 , Telegram clients support rounded square MPEG4 videos of up to 1 minute long.
	// Use this method to send video messages. On success, the sent Message is returned.
	SendVideoNote(request *SendVideoNoteRequest) (*SendVideoNoteResponse, error)

	//   Use this method to send point on the map. On success, the sent Message is returned.
	SendLocation(request *SendLocationRequest) (*SendLocationResponse, error)

	//   Use this method to send information about a venue. On success, the sent Message is returned.
	SendVenue(request *SendVenueRequest) (*SendVenueResponse, error)

	//   Use this method to send phone contacts. On success, the sent Message is returned.
	SendContact(request *SendContactRequest) (*SendContactResponse, error)

	//   Use this method to send a native poll. On success, the sent Message is returned.
	SendPoll(request *SendPollRequest) (*SendPollResponse, error)

	// Use this method to send an animated emoji that will display a random value. On success, the
	// sent Message is returned.
	SendDice(request *SendDiceRequest) (*SendDiceResponse, error)

	// Use this method when you need to tell the user that something is happening on the bot's side.
	// The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients
	// clear its typing status). Returns True on success.   Example: The ImageBot needs some time to
	// process a request and upload the image. Instead of sending a text message along the lines of
	// “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo .
	// The user will see a “sending photo” status for the bot.   We only recommend using this method
	// when a response from the bot will take a noticeable amount of time to arrive.
	SendChatAction(request *SendChatActionRequest) (*SendChatActionResponse, error)

	// Use this method to change the chosen reactions on a message. Service messages can't be
	// reacted to. Automatically forwarded messages from a channel to its discussion group have the
	// same available reactions as messages in the channel. Returns True on success.
	SetMessageReaction(request *SetMessageReactionRequest) (*SetMessageReactionResponse, error)

	// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos
	// object.
	GetUserProfilePhotos(request *GetUserProfilePhotosRequest) (*GetUserProfilePhotosResponse, error)

	// Use this method to get basic information about a file and prepare it for downloading. For the
	// moment, bots can download files of up to 20MB in size. On success, a File object is returned.
	// The file can then be downloaded via the link
	// https://api.telegram.org/file/bot<token>/<file_path> , where <file_path> is taken from the
	// response. It is guaranteed that the link will be valid for at least 1 hour. When the link
	// expires, a new one can be requested by calling getFile again.
	GetFile(request *GetFileRequest) (*GetFileResponse, error)

	// Note: This function may not preserve the original file name and MIME type. You should save the
	// file's MIME type and name (if available) when the File object is received.   Use this method to
	// ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the
	// user will not be able to return to the chat on their own using invite links, etc., unless
	// unbanned first. The bot must be an administrator in the chat for this to work and must have the
	// appropriate administrator rights. Returns True on success.
	BanChatMember(request *BanChatMemberRequest) (*BanChatMemberResponse, error)

	// Use this method to unban a previously banned user in a supergroup or channel. The user will
	// not return to the group or channel automatically, but will be able to join via link, etc. The
	// bot must be an administrator for this to work. By default, this method guarantees that after
	// the call the user is not a member of the chat, but will be able to join it. So if the user is a
	// member of the chat they will also be removed from the chat. If you don't want this, use the
	// parameter only_if_banned . Returns True on success.
	UnbanChatMember(request *UnbanChatMemberRequest) (*UnbanChatMemberResponse, error)

	// Use this method to restrict a user in a supergroup. The bot must be an administrator in the
	// supergroup for this to work and must have the appropriate administrator rights. Pass True for
	// all permissions to lift restrictions from a user. Returns True on success.
	RestrictChatMember(request *RestrictChatMemberRequest) (*RestrictChatMemberResponse, error)

	// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an
	// administrator in the chat for this to work and must have the appropriate administrator rights.
	// Pass False for all boolean parameters to demote a user. Returns True on success.
	PromoteChatMember(request *PromoteChatMemberRequest) (*PromoteChatMemberResponse, error)

	// Use this method to set a custom title for an administrator in a supergroup promoted by the
	// bot. Returns True on success.
	SetChatAdministratorCustomTitle(request *SetChatAdministratorCustomTitleRequest) (*SetChatAdministratorCustomTitleResponse, error)

	// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is
	// unbanned , the owner of the banned chat won't be able to send messages on behalf of any of
	// their channels . The bot must be an administrator in the supergroup or channel for this to work
	// and must have the appropriate administrator rights. Returns True on success.
	BanChatSenderChat(request *BanChatSenderChatRequest) (*BanChatSenderChatResponse, error)

	// Use this method to unban a previously banned channel chat in a supergroup or channel. The bot
	// must be an administrator for this to work and must have the appropriate administrator rights.
	// Returns True on success.
	UnbanChatSenderChat(request *UnbanChatSenderChatRequest) (*UnbanChatSenderChatResponse, error)

	// Use this method to set default chat permissions for all members. The bot must be an
	// administrator in the group or a supergroup for this to work and must have the
	// can_restrict_members administrator rights. Returns True on success.
	SetChatPermissions(request *SetChatPermissionsRequest) (*SetChatPermissionsResponse, error)

	// Use this method to generate a new primary invite link for a chat; any previously generated
	// primary link is revoked. The bot must be an administrator in the chat for this to work and must
	// have the appropriate administrator rights. Returns the new invite link as String on success.
	ExportChatInviteLink(request *ExportChatInviteLinkRequest) (*ExportChatInviteLinkResponse, error)

	// Note: Each administrator in a chat generates their own invite links. Bots can't use invite
	// links generated by other administrators. If you want your bot to work with invite links, it
	// will need to generate its own link using exportChatInviteLink or by calling the getChat method.
	// If your bot needs to generate a new primary invite link replacing its previous one, use
	// exportChatInviteLink again.    Use this method to create an additional invite link for a chat.
	// The bot must be an administrator in the chat for this to work and must have the appropriate
	// administrator rights. The link can be revoked using the method revokeChatInviteLink . Returns
	// the new invite link as ChatInviteLink object.
	CreateChatInviteLink(request *CreateChatInviteLinkRequest) (*CreateChatInviteLinkResponse, error)

	// Use this method to edit a non-primary invite link created by the bot. The bot must be an
	// administrator in the chat for this to work and must have the appropriate administrator rights.
	// Returns the edited invite link as a ChatInviteLink object.
	EditChatInviteLink(request *EditChatInviteLinkRequest) (*EditChatInviteLinkResponse, error)

	// Use this method to revoke an invite link created by the bot. If the primary link is revoked,
	// a new link is automatically generated. The bot must be an administrator in the chat for this to
	// work and must have the appropriate administrator rights. Returns the revoked invite link as
	// ChatInviteLink object.
	RevokeChatInviteLink(request *RevokeChatInviteLinkRequest) (*RevokeChatInviteLinkResponse, error)

	// Use this method to approve a chat join request. The bot must be an administrator in the chat
	// for this to work and must have the can_invite_users administrator right. Returns True on
	// success.
	ApproveChatJoinRequest(request *ApproveChatJoinRequestRequest) (*ApproveChatJoinRequestResponse, error)

	// Use this method to decline a chat join request. The bot must be an administrator in the chat
	// for this to work and must have the can_invite_users administrator right. Returns True on
	// success.
	DeclineChatJoinRequest(request *DeclineChatJoinRequestRequest) (*DeclineChatJoinRequestResponse, error)

	// Use this method to set a new profile photo for the chat. Photos can't be changed for private
	// chats. The bot must be an administrator in the chat for this to work and must have the
	// appropriate administrator rights. Returns True on success.
	SetChatPhoto(request *SetChatPhotoRequest) (*SetChatPhotoResponse, error)

	// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot
	// must be an administrator in the chat for this to work and must have the appropriate
	// administrator rights. Returns True on success.
	DeleteChatPhoto(request *DeleteChatPhotoRequest) (*DeleteChatPhotoResponse, error)

	// Use this method to change the title of a chat. Titles can't be changed for private chats. The
	// bot must be an administrator in the chat for this to work and must have the appropriate
	// administrator rights. Returns True on success.
	SetChatTitle(request *SetChatTitleRequest) (*SetChatTitleResponse, error)

	// Use this method to change the description of a group, a supergroup or a channel. The bot must
	// be an administrator in the chat for this to work and must have the appropriate administrator
	// rights. Returns True on success.
	SetChatDescription(request *SetChatDescriptionRequest) (*SetChatDescriptionResponse, error)

	// Use this method to add a message to the list of pinned messages in a chat. If the chat is not
	// a private chat, the bot must be an administrator in the chat for this to work and must have the
	// 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator
	// right in a channel. Returns True on success.
	PinChatMessage(request *PinChatMessageRequest) (*PinChatMessageResponse, error)

	// Use this method to remove a message from the list of pinned messages in a chat. If the chat
	// is not a private chat, the bot must be an administrator in the chat for this to work and must
	// have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages'
	// administrator right in a channel. Returns True on success.
	UnpinChatMessage(request *UnpinChatMessageRequest) (*UnpinChatMessageResponse, error)

	// Use this method to clear the list of pinned messages in a chat. If the chat is not a private
	// chat, the bot must be an administrator in the chat for this to work and must have the
	// 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator
	// right in a channel. Returns True on success.
	UnpinAllChatMessages(request *UnpinAllChatMessagesRequest) (*UnpinAllChatMessagesResponse, error)

	// Use this method for your bot to leave a group, supergroup or channel. Returns True on
	// success.
	LeaveChat(request *LeaveChatRequest) (*LeaveChatResponse, error)

	// Use this method to get up to date information about the chat. Returns a Chat object on
	// success.
	GetChat(request *GetChatRequest) (*GetChatResponse, error)

	// Use this method to get a list of administrators in a chat, which aren't bots. Returns an
	// Array of ChatMember objects.
	GetChatAdministrators(request *GetChatAdministratorsRequest) (*GetChatAdministratorsResponse, error)

	//   Use this method to get the number of members in a chat. Returns Int on success.
	GetChatMemberCount(request *GetChatMemberCountRequest) (*GetChatMemberCountResponse, error)

	// Use this method to get information about a member of a chat. The method is only guaranteed to
	// work for other users if the bot is an administrator in the chat. Returns a ChatMember object on
	// success.
	GetChatMember(request *GetChatMemberRequest) (*GetChatMemberResponse, error)

	// Use this method to set a new group sticker set for a supergroup. The bot must be an
	// administrator in the chat for this to work and must have the appropriate administrator rights.
	// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot
	// can use this method. Returns True on success.
	SetChatStickerSet(request *SetChatStickerSetRequest) (*SetChatStickerSetResponse, error)

	// Use this method to delete a group sticker set from a supergroup. The bot must be an
	// administrator in the chat for this to work and must have the appropriate administrator rights.
	// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot
	// can use this method. Returns True on success.
	DeleteChatStickerSet(request *DeleteChatStickerSetRequest) (*DeleteChatStickerSetResponse, error)

	// Use this method to get custom emoji stickers, which can be used as a forum topic icon by any
	// user. Requires no parameters. Returns an Array of Sticker objects.
	GetForumTopicIconStickers(request *GetForumTopicIconStickersRequest) (*GetForumTopicIconStickersResponse, error)

	// Use this method to create a topic in a forum supergroup chat. The bot must be an administrator
	// in the chat for this to work and must have the can_manage_topics administrator rights. Returns
	// information about the created topic as a ForumTopic object.
	CreateForumTopic(request *CreateForumTopicRequest) (*CreateForumTopicResponse, error)

	// Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must be
	// an administrator in the chat for this to work and must have can_manage_topics administrator
	// rights, unless it is the creator of the topic. Returns True on success.
	EditForumTopic(request *EditForumTopicRequest) (*EditForumTopicResponse, error)

	// Use this method to close an open topic in a forum supergroup chat. The bot must be an
	// administrator in the chat for this to work and must have the can_manage_topics administrator
	// rights, unless it is the creator of the topic. Returns True on success.
	CloseForumTopic(request *CloseForumTopicRequest) (*CloseForumTopicResponse, error)

	// Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an
	// administrator in the chat for this to work and must have the can_manage_topics administrator
	// rights, unless it is the creator of the topic. Returns True on success.
	ReopenForumTopic(request *ReopenForumTopicRequest) (*ReopenForumTopicResponse, error)

	// Use this method to delete a forum topic along with all its messages in a forum supergroup
	// chat. The bot must be an administrator in the chat for this to work and must have the
	// can_delete_messages administrator rights. Returns True on success.
	DeleteForumTopic(request *DeleteForumTopicRequest) (*DeleteForumTopicResponse, error)

	// Use this method to clear the list of pinned messages in a forum topic. The bot must be an
	// administrator in the chat for this to work and must have the can_pin_messages administrator
	// right in the supergroup. Returns True on success.
	UnpinAllForumTopicMessages(request *UnpinAllForumTopicMessagesRequest) (*UnpinAllForumTopicMessagesResponse, error)

	// Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot
	// must be an administrator in the chat for this to work and must have can_manage_topics
	// administrator rights. Returns True on success.
	EditGeneralForumTopic(request *EditGeneralForumTopicRequest) (*EditGeneralForumTopicResponse, error)

	// Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be
	// an administrator in the chat for this to work and must have the can_manage_topics administrator
	// rights. Returns True on success.
	CloseGeneralForumTopic(request *CloseGeneralForumTopicRequest) (*CloseGeneralForumTopicResponse, error)

	// Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must
	// be an administrator in the chat for this to work and must have the can_manage_topics
	// administrator rights. The topic will be automatically unhidden if it was hidden. Returns True
	// on success.
	ReopenGeneralForumTopic(request *ReopenGeneralForumTopicRequest) (*ReopenGeneralForumTopicResponse, error)

	// Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an
	// administrator in the chat for this to work and must have the can_manage_topics administrator
	// rights. The topic will be automatically closed if it was open. Returns True on success.
	HideGeneralForumTopic(request *HideGeneralForumTopicRequest) (*HideGeneralForumTopicResponse, error)

	// Use this method to unhide the 'General' topic in a forum supergroup chat. The bot must be an
	// administrator in the chat for this to work and must have the can_manage_topics administrator
	// rights. Returns True on success.
	UnhideGeneralForumTopic(request *UnhideGeneralForumTopicRequest) (*UnhideGeneralForumTopicResponse, error)

	// Use this method to clear the list of pinned messages in a General forum topic. The bot must
	// be an administrator in the chat for this to work and must have the can_pin_messages
	// administrator right in the supergroup. Returns True on success.
	UnpinAllGeneralForumTopicMessages(request *UnpinAllGeneralForumTopicMessagesRequest) (*UnpinAllGeneralForumTopicMessagesResponse, error)

	// Use this method to send answers to callback queries sent from inline keyboards . The answer
	// will be displayed to the user as a notification at the top of the chat screen or as an alert.
	// On success, True is returned.   Alternatively, the user can be redirected to the specified Game
	// URL. For this option to work, you must first create a game for your bot via @BotFather and
	// accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot
	// with a parameter.
	AnswerCallbackQuery(request *AnswerCallbackQueryRequest) (*AnswerCallbackQueryResponse, error)

	// Use this method to get the list of boosts added to a chat by a user. Requires administrator
	// rights in the chat. Returns a UserChatBoosts object.
	GetUserChatBoosts(request *GetUserChatBoostsRequest) (*GetUserChatBoostsResponse, error)

	// Use this method to change the list of the bot's commands. See this manual for more details
	// about bot commands. Returns True on success.
	SetMyCommands(request *SetMyCommandsRequest) (*SetMyCommandsResponse, error)

	// Use this method to delete the list of the bot's commands for the given scope and user
	// language. After deletion, higher level commands will be shown to affected users. Returns True
	// on success.
	DeleteMyCommands(request *DeleteMyCommandsRequest) (*DeleteMyCommandsResponse, error)

	// Use this method to get the current list of the bot's commands for the given scope and user
	// language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is
	// returned.
	GetMyCommands(request *GetMyCommandsRequest) (*GetMyCommandsResponse, error)

	//   Use this method to change the bot's name. Returns True on success.
	SetMyName(request *SetMyNameRequest) (*SetMyNameResponse, error)

	// Use this method to get the current bot name for the given user language. Returns BotName on
	// success.
	GetMyName(request *GetMyNameRequest) (*GetMyNameResponse, error)

	// Use this method to change the bot's description, which is shown in the chat with the bot if
	// the chat is empty. Returns True on success.
	SetMyDescription(request *SetMyDescriptionRequest) (*SetMyDescriptionResponse, error)

	// Use this method to get the current bot description for the given user language. Returns
	// BotDescription on success.
	GetMyDescription(request *GetMyDescriptionRequest) (*GetMyDescriptionResponse, error)

	// Use this method to change the bot's short description, which is shown on the bot's profile
	// page and is sent together with the link when users share the bot. Returns True on success.
	SetMyShortDescription(request *SetMyShortDescriptionRequest) (*SetMyShortDescriptionResponse, error)

	// Use this method to get the current bot short description for the given user language. Returns
	// BotShortDescription on success.
	GetMyShortDescription(request *GetMyShortDescriptionRequest) (*GetMyShortDescriptionResponse, error)

	// Use this method to change the bot's menu button in a private chat, or the default menu
	// button. Returns True on success.
	SetChatMenuButton(request *SetChatMenuButtonRequest) (*SetChatMenuButtonResponse, error)

	// Use this method to get the current value of the bot's menu button in a private chat, or the
	// default menu button. Returns MenuButton on success.
	GetChatMenuButton(request *GetChatMenuButtonRequest) (*GetChatMenuButtonResponse, error)

	// Use this method to change the default administrator rights requested by the bot when it's
	// added as an administrator to groups or channels. These rights will be suggested to users, but
	// they are free to modify the list before adding the bot. Returns True on success.
	SetMyDefaultAdministratorRights(request *SetMyDefaultAdministratorRightsRequest) (*SetMyDefaultAdministratorRightsResponse, error)

	// Use this method to get the current default administrator rights of the bot. Returns
	// ChatAdministratorRights on success.
	GetMyDefaultAdministratorRights(request *GetMyDefaultAdministratorRightsRequest) (*GetMyDefaultAdministratorRightsResponse, error)

	// Use this method to edit text and game messages. On success, if the edited message is not an
	// inline message, the edited Message is returned, otherwise True is returned.
	EditMessageText(request *EditMessageTextRequest) (*EditMessageTextResponse, error)

	// Use this method to edit captions of messages. On success, if the edited message is not an
	// inline message, the edited Message is returned, otherwise True is returned.
	EditMessageCaption(request *EditMessageCaptionRequest) (*EditMessageCaptionResponse, error)

	// Use this method to edit animation, audio, document, photo, or video messages. If a message is
	// part of a message album, then it can be edited only to an audio for audio albums, only to a
	// document for document albums and to a photo or a video otherwise. When an inline message is
	// edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify
	// a URL. On success, if the edited message is not an inline message, the edited Message is
	// returned, otherwise True is returned.
	EditMessageMedia(request *EditMessageMediaRequest) (*EditMessageMediaResponse, error)

	// Use this method to edit live location messages. A location can be edited until its
	// live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation . On
	// success, if the edited message is not an inline message, the edited Message is returned,
	// otherwise True is returned.
	EditMessageLiveLocation(request *EditMessageLiveLocationRequest) (*EditMessageLiveLocationResponse, error)

	// Use this method to stop updating a live location message before live_period expires. On
	// success, if the message is not an inline message, the edited Message is returned, otherwise
	// True is returned.
	StopMessageLiveLocation(request *StopMessageLiveLocationRequest) (*StopMessageLiveLocationResponse, error)

	// Use this method to edit only the reply markup of messages. On success, if the edited message
	// is not an inline message, the edited Message is returned, otherwise True is returned.
	EditMessageReplyMarkup(request *EditMessageReplyMarkupRequest) (*EditMessageReplyMarkupResponse, error)

	// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll is
	// returned.
	StopPoll(request *StopPollRequest) (*StopPollResponse, error)

	// Use this method to delete a message, including service messages, with the following
	// limitations: - A message can only be deleted if it was sent less than 48 hours ago. - Service
	// messages about a supergroup, channel, or forum topic creation can't be deleted. - A dice
	// message in a private chat can only be deleted if it was sent more than 24 hours ago. - Bots can
	// delete outgoing messages in private chats, groups, and supergroups. - Bots can delete incoming
	// messages in private chats. - Bots granted can_post_messages permissions can delete outgoing
	// messages in channels. - If the bot is an administrator of a group, it can delete any message
	// there. - If the bot has can_delete_messages permission in a supergroup or a channel, it can
	// delete any message there. Returns True on success.
	DeleteMessage(request *DeleteMessageRequest) (*DeleteMessageResponse, error)

	// Use this method to delete multiple messages simultaneously. If some of the specified messages
	// can't be found, they are skipped. Returns True on success.
	DeleteMessages(request *DeleteMessagesRequest) (*DeleteMessagesResponse, error)

	// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the
	// sent Message is returned.
	SendSticker(request *SendStickerRequest) (*SendStickerResponse, error)

	//   Use this method to get a sticker set. On success, a StickerSet object is returned.
	GetStickerSet(request *GetStickerSetRequest) (*GetStickerSetResponse, error)

	// Use this method to get information about custom emoji stickers by their identifiers. Returns
	// an Array of Sticker objects.
	GetCustomEmojiStickers(request *GetCustomEmojiStickersRequest) (*GetCustomEmojiStickersResponse, error)

	// Use this method to upload a file with a sticker for later use in the createNewStickerSet and
	// addStickerToSet methods (the file can be used multiple times). Returns the uploaded File on
	// success.
	UploadStickerFile(request *UploadStickerFileRequest) (*UploadStickerFileResponse, error)

	// Use this method to create a new sticker set owned by a user. The bot will be able to edit the
	// sticker set thus created. Returns True on success.
	CreateNewStickerSet(request *CreateNewStickerSetRequest) (*CreateNewStickerSetResponse, error)

	// Use this method to add a new sticker to a set created by the bot. The format of the added
	// sticker must match the format of the other stickers in the set. Emoji sticker sets can have up
	// to 200 stickers. Animated and video sticker sets can have up to 50 stickers. Static sticker
	// sets can have up to 120 stickers. Returns True on success.
	AddStickerToSet(request *AddStickerToSetRequest) (*AddStickerToSetResponse, error)

	// Use this method to move a sticker in a set created by the bot to a specific position. Returns
	// True on success.
	SetStickerPositionInSet(request *SetStickerPositionInSetRequest) (*SetStickerPositionInSetResponse, error)

	//   Use this method to delete a sticker from a set created by the bot. Returns True on success.
	DeleteStickerFromSet(request *DeleteStickerFromSetRequest) (*DeleteStickerFromSetResponse, error)

	// Use this method to change the list of emoji assigned to a regular or custom emoji sticker.
	// The sticker must belong to a sticker set created by the bot. Returns True on success.
	SetStickerEmojiList(request *SetStickerEmojiListRequest) (*SetStickerEmojiListResponse, error)

	// Use this method to change search keywords assigned to a regular or custom emoji sticker. The
	// sticker must belong to a sticker set created by the bot. Returns True on success.
	SetStickerKeywords(request *SetStickerKeywordsRequest) (*SetStickerKeywordsResponse, error)

	// Use this method to change the mask position of a mask sticker. The sticker must belong to a
	// sticker set that was created by the bot. Returns True on success.
	SetStickerMaskPosition(request *SetStickerMaskPositionRequest) (*SetStickerMaskPositionResponse, error)

	//   Use this method to set the title of a created sticker set. Returns True on success.
	SetStickerSetTitle(request *SetStickerSetTitleRequest) (*SetStickerSetTitleResponse, error)

	// Use this method to set the thumbnail of a regular or mask sticker set. The format of the
	// thumbnail file must match the format of the stickers in the set. Returns True on success.
	SetStickerSetThumbnail(request *SetStickerSetThumbnailRequest) (*SetStickerSetThumbnailResponse, error)

	//   Use this method to set the thumbnail of a custom emoji sticker set. Returns True on success.
	SetCustomEmojiStickerSetThumbnail(request *SetCustomEmojiStickerSetThumbnailRequest) (*SetCustomEmojiStickerSetThumbnailResponse, error)

	// Use this method to delete a sticker set that was created by the bot. Returns True on success.
	//
	DeleteStickerSet(request *DeleteStickerSetRequest) (*DeleteStickerSetResponse, error)

	// Use this method to send answers to an inline query. On success, True is returned. No more
	// than 50 results per query are allowed.
	AnswerInlineQuery(request *AnswerInlineQueryRequest) (*AnswerInlineQueryResponse, error)

	// Note: It is necessary to enable inline feedback via @BotFather in order to receive these
	// objects in updates.   Use this method to set the result of an interaction with a Web App and
	// send a corresponding message on behalf of the user to the chat from which the query originated.
	// On success, a SentWebAppMessage object is returned.
	AnswerWebAppQuery(request *AnswerWebAppQueryRequest) (*AnswerWebAppQueryResponse, error)

	// Your bot can accept payments from Telegram users. Please see the introduction to payments for
	// more details on the process and how to set up payments for your bot.   Use this method to send
	// invoices. On success, the sent Message is returned.
	SendInvoice(request *SendInvoiceRequest) (*SendInvoiceResponse, error)

	// Use this method to create a link for an invoice. Returns the created invoice link as String
	// on success.
	CreateInvoiceLink(request *CreateInvoiceLinkRequest) (*CreateInvoiceLinkResponse, error)

	// If you sent an invoice requesting a shipping address and the parameter is_flexible was
	// specified, the Bot API will send an Update with a shipping_query field to the bot. Use this
	// method to reply to shipping queries. On success, True is returned.
	AnswerShippingQuery(request *AnswerShippingQueryRequest) (*AnswerShippingQueryResponse, error)

	// Once the user has confirmed their payment and shipping details, the Bot API sends the final
	// confirmation in the form of an Update with the field pre_checkout_query . Use this method to
	// respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must
	// receive an answer within 10 seconds after the pre-checkout query was sent.
	AnswerPreCheckoutQuery(request *AnswerPreCheckoutQueryRequest) (*AnswerPreCheckoutQueryResponse, error)

	// Informs a user that some of the Telegram Passport elements they provided contains errors. The
	// user will not be able to re-submit their Passport to you until the errors are fixed (the
	// contents of the field for which you returned the error must change). Returns True on success.
	// Use this if the data submitted by the user doesn't satisfy the standards your service requires
	// for any reason. For example, if a birthday date seems invalid, a submitted document is blurry,
	// a scan shows evidence of tampering, etc. Supply some details in the error message to make sure
	// the user knows how to correct the issues.
	SetPassportDataErrors(request *SetPassportDataErrorsRequest) (*SetPassportDataErrorsResponse, error)

	// Your bot can offer users HTML5 games to play solo or to compete against each other in groups
	// and one-on-one chats. Create games via @BotFather using the /newgame command. Please note that
	// this kind of power requires responsibility: you will need to accept the terms for each game
	// that your bots will be offering.   Games are a new type of content on Telegram, represented by
	// the Game and InlineQueryResultGame objects.  Once you've created a game via BotFather , you can
	// send games to chats as regular messages using the sendGame method, or use inline mode with
	// InlineQueryResultGame .  If you send the game message without any buttons, it will
	// automatically have a 'Play GameName ' button. When this button is pressed, your bot gets a
	// CallbackQuery with the game_short_name of the requested game. You provide the correct URL for
	// this particular user and the app opens the game in the in-app browser.  You can manually add
	// multiple buttons to your game message. Please note that the first button in the first row must
	// always launch the game, using the field callback_game in InlineKeyboardButton . You can add
	// extra buttons according to taste: e.g., for a description of the rules, or to open the game's
	// official community.  To make your game more attractive, you can upload a GIF animation that
	// demostrates the game to the users via BotFather (see Lumberjack for example).  A game message
	// will also display high scores for the current chat. Use setGameScore to post high scores to the
	// chat with the game, add the edit_message parameter to automatically update the message with the
	// current scoreboard.  Use getGameHighScores to get data for in-game high score tables.  You can
	// also add an extra sharing button for users to share their best score to different chats.  For
	// examples of what can be done using this new stuff, check the @gamebot and @gamee bots.    Use
	// this method to send a game. On success, the sent Message is returned.
	SendGame(request *SendGameRequest) (*SendGameResponse, error)

	// Use this method to set the score of the specified user in a game message. On success, if the
	// message is not an inline message, the Message is returned, otherwise True is returned. Returns
	// an error, if the new score is not greater than the user's current score in the chat and force
	// is False .
	SetGameScore(request *SetGameScoreRequest) (*SetGameScoreResponse, error)

	// Use this method to get data for high score tables. Will return the score of the specified
	// user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
	// This method will currently return scores for the target user, plus two of their closest
	// neighbors on each side. Will also return the top three users if the user and their neighbors
	// are not among them. Please note that this behavior is subject to change.
	GetGameHighScores(request *GetGameHighScoresRequest) (*GetGameHighScoresResponse, error)
}