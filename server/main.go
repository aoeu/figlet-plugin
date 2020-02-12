package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// FIGlet is a mattermost plugin for transforming provided text and writing it back
// to the channel in which the plugin was invoked.
type FIGlet struct {
	plugin.MattermostPlugin
	binaryPath string
	fontsPath  string
}

const (
	triggerWord = "figlet"
	botID       = "unknown"
	token       = "unknown"
	pluginURL   = "https://shouting.online/figlet"
)

func main() {
	plugin.ClientMain(&FIGlet{})
}

// OnActivate is a function from the mattermost plugin framework that is called
// automatically once the plugin has loaded.
func (f *FIGlet) OnActivate() error {
	if err := f.init(); err != nil {
		return fmt.Errorf("could not init FIGlet state upon plugin activation: %v", err)
	}
	c := &model.Command{
		Trigger:          triggerWord,
		DisplayName:      "FIGlet",
		Description:      description,
		AutoComplete:     true,
		AutoCompleteDesc: "/figlet this text will be printed in large letters",
		AutoCompleteHint: "[text]",
	}
	if err := f.API.RegisterCommand(c); err != nil {
		return fmt.Errorf("could not initialize FIGlet plugin due to error: %v", err)
	}
	return nil
}

func (f *FIGlet) init() error {
	bundlePath, err := f.API.GetBundlePath()
	if err != nil {
		return fmt.Errorf("could not initialize FIGlet plugin with bundle path: %v", err)
	}

	f.binaryPath, err = filepath.Abs(filepath.Join(bundlePath, "assets", "figlet"))
	if err != nil {
		return fmt.Errorf("could not initialize FIGlet plugin with binary path: %v", err)
	}

	f.fontsPath, err = filepath.Abs(filepath.Join(bundlePath, "essets", "fonts"))
	if err != nil {
		return fmt.Errorf("could not initialize FIGlet plugin with fonts path: %v", err)
	}
	return nil
}

// ExecuteCommand runs the FIGlet binary on the provided text to transform
// it, and the transformed text is posted back to the mattermost server in
// the same channel in which this plugin was invoked.
func (f *FIGlet) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	// TODO(aoeu): Token authentication.
	t, err := f.transformText(args.Command)
	if err != nil {
		return nil, &model.AppError{
			DetailedError: err.Error(),
		}
	}
	return &model.CommandResponse{
		Username:     botID,
		Text:         t,
		ChannelId:    args.ChannelId,
		TriggerId:    args.TriggerId,
		ResponseType: "in_channel", // ResponseType is Required.
	}, nil
}

func (f FIGlet) transformText(in string) (out string, err error) {
	cmd := exec.Command(f.binaryPath, "-d", f.fontsPath, in)
	cmd.Stdin = os.Stdin
	b, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not run and read output of FIGlet: %v", err)
	}
	return string(b), nil
}

const description = `FIGlet is a program for making large letters out of ordinary text

 _ _ _          _   _     _
| (_) | _____  | |_| |__ (_)___
| | | |/ / _ \ | __|  _ \| / __|
| | |   <  __/ | |_| | | | \__ \
|_|_|_|\_\___|  \__|_| |_|_|___/


                          .   oooo         o8o
                        .o8   '888         '"'
 .ooooo.  oooo d8b    .o888oo  888 .oo.   oooo   .oooo.o
d88' '88b '888""8P      888    888P"Y88b  '888  d88(  "8
888   888  888          888    888   888   888  '"Y88b.
888   888  888          888 .  888   888   888  o.  )88b
'Y8bod8P' d888b         "888" o888o o888o o888o 8""888P'
`
