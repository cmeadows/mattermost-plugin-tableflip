package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

// Tableflip is a Mattermost plugin that adds a slash command
// to add (╯°□°)╯︵ ┻━┻ to messages like the /shrug

type Plugin struct {
	plugin.MattermostPlugin

	enabled       bool
}

const (
	trigger      string = "tableflip"
	ascii_flip   string = "(╯°□°)╯︵ ┻━┻"
)

// OnActivate handles plugin deactivation
func (p *Plugin) OnActivate() error {
	p.enabled = true

	return p.API.RegisterCommand(&model.Command{
		Trigger:          trigger,
		AutoComplete:     true,
		AutoCompleteHint: "[message]",
		AutoCompleteDesc: "Adds " + ascii_flip + " to your message",
	})
}

// OnDeactivate handles plugin deactivation
func (p *Plugin) OnDeactivate() error {
	p.enabled = false
	return nil
}

// ExecuteCommand handles the core functionality of the plugin
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	if !p.enabled {
		return nil, appError("Cannot execute command while the plugin is disabled.", nil)
	}

	if p.API == nil {
		return nil, appError("Cannot access the plugin API.", nil)
	}

	slashCommand := "/" + trigger

	if !strings.HasPrefix(args.Command, slashCommand) {
		return nil, appError("Expected trigger " + slashCommand + ", but got " + args.Command, nil)
	}

	message := strings.TrimSpace((strings.Replace(args.Command, slashCommand, "", 1)))

	if len(message) == 0{
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
			Text:         fmt.Sprintf(ascii_flip),
		}, nil
	}

	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
		Text:         fmt.Sprintf(message + " " + ascii_flip),
	}, nil
}

func appError(message string, err error) *model.AppError {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	return model.NewAppError("Tableflip Plugin", message, nil, errorMessage, http.StatusBadRequest)
}

func main() {
	plugin.ClientMain(&Plugin{})
}
