// App contains helpers and structs for discord objects
// all structs were generated from the discord object definition tables using chatgpt and modified after
package app

// struct for ApplicationCommand,
// adheres to https://discord.com/developers/docs/interactions/application-commands#application-command-object
// and its respective structure
// from the docs: "Given that all Discord IDs are snowflakes, you should always expect a string."
type ApplicationCommand struct {
	ID                       string                     `json:"id"`
	Type                     *int                       `json:"type,omitempty"`
	ApplicationID            string                     `json:"application_id"`
	GuildID                  *string                    `json:"guild_id,omitempty"`
	Name                     string                     `json:"name"`
	NameLocalizations        map[string]string          `json:"name_localizations,omitempty"`
	Description              string                     `json:"description"`
	DescriptionLocalizations map[string]string          `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions *string                    `json:"default_member_permissions,omitempty"`
	DMPermission             *bool                      `json:"dm_permission,omitempty"`
	DefaultPermission        *bool                      `json:"default_permission,omitempty"`
	NSFW                     *bool                      `json:"nsfw,omitempty"`
	IntegrationTypes         []int                      `json:"integration_types,omitempty"`
	Contexts                 []string                   `json:"contexts,omitempty"`
	Version                  string                     `json:"version"`
	Handler                  *string                    `json:"handler,omitempty"`
}

// ApplicationCommandType represents the type of an application command.
// it adheres to https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types
type ApplicationCommandType int

const (
	CHAT_INPUT          ApplicationCommandType = 1 // Slash commands
	USER                ApplicationCommandType = 2 // User context menu commands
	MESSAGE             ApplicationCommandType = 3 // Message context menu commands
	PRIMARY_ENTRY_POINT ApplicationCommandType = 4 // Primary activity invocation
)

// String provides a human-readable representation of the command type.
func (t ApplicationCommandType) String() string {
	switch t {
	case CHAT_INPUT:
		return "CHAT_INPUT"
	case USER:
		return "USER"
	case MESSAGE:
		return "MESSAGE"
	case PRIMARY_ENTRY_POINT:
		return "PRIMARY_ENTRY_POINT"
	default:
		return "UNKNOWN"
	}
}

// ApplicationCommandOption represents an option for an application command.
type ApplicationCommandOption struct {
	Type                     ApplicationCommandOptionType     `json:"type"`
	Name                     string                           `json:"name"`
	NameLocalizations        map[string]string                `json:"name_localizations,omitempty"`
	Description              string                           `json:"description"`
	DescriptionLocalizations map[string]string                `json:"description_localizations,omitempty"`
	Required                 *bool                            `json:"required,omitempty"`
	Choices                  []ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options                  []ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes             []int                            `json:"channel_types,omitempty"`
	MinValue                 *float64                         `json:"min_value,omitempty"`
	MaxValue                 *float64                         `json:"max_value,omitempty"`
	MinLength                *int                             `json:"min_length,omitempty"`
	MaxLength                *int                             `json:"max_length,omitempty"`
	Autocomplete             *bool                            `json:"autocomplete,omitempty"`
}

// ApplicationCommandOptionType represents the type of an application command option.
type ApplicationCommandOptionType int

// Constants for valid application command option types.
const (
	SUB_COMMAND       ApplicationCommandOptionType = 1  // Represents a sub-command
	SUB_COMMAND_GROUP ApplicationCommandOptionType = 2  // Represents a sub-command group
	STRING            ApplicationCommandOptionType = 3  // String input (1-6000 characters)
	INTEGER           ApplicationCommandOptionType = 4  // Integer input (-2^53 to 2^53)
	BOOLEAN           ApplicationCommandOptionType = 5  // True/false value
	USER_OT           ApplicationCommandOptionType = 6  // A Discord user
	CHANNEL           ApplicationCommandOptionType = 7  // A Discord channel (including categories)
	ROLE              ApplicationCommandOptionType = 8  // A Discord role
	MENTIONABLE       ApplicationCommandOptionType = 9  // Either a user or a role
	NUMBER            ApplicationCommandOptionType = 10 // Floating-point number (-2^53 to 2^53)
	ATTACHMENT        ApplicationCommandOptionType = 11 // An attachment object
)

// String provides a readable representation of the command option type.
func (t ApplicationCommandOptionType) String() string {
	switch t {
	case SUB_COMMAND:
		return "SUB_COMMAND"
	case SUB_COMMAND_GROUP:
		return "SUB_COMMAND_GROUP"
	case STRING:
		return "STRING"
	case INTEGER:
		return "INTEGER"
	case BOOLEAN:
		return "BOOLEAN"
	case USER_OT:
		return "USER"
	case CHANNEL:
		return "CHANNEL"
	case ROLE:
		return "ROLE"
	case MENTIONABLE:
		return "MENTIONABLE"
	case NUMBER:
		return "NUMBER"
	case ATTACHMENT:
		return "ATTACHMENT"
	default:
		return "UNKNOWN"
	}
}

// ApplicationCommandOptionChoice represents a choice for an application command option.
type ApplicationCommandOptionChoice struct {
	Name              string            `json:"name"`                         // 1-100 character choice name
	NameLocalizations map[string]string `json:"name_localizations,omitempty"` // Localization dictionary for name field
	Value             any               `json:"value"`                        // String, integer, or double value
}
