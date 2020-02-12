package main

import (
	"fmt"
	"log"
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

const triggerWord = "figlet"

func main() {
	f := &FIGlet{}
	if err := f.init(); err != nil {
		log.Fatal("could not complete main function due to error:", err)
	}
	plugin.ClientMain(f)
}

func (f *FIGlet) init() error {
	// TODO(aoeu): What fields are needed in the Command struct? The godoc doesn't specify.
	c := &model.Command{Trigger: triggerWord}
	if err := f.API.RegisterCommand(c); err != nil {
		return fmt.Errorf("could not initialize FIGlet plugin due to error: %v", err)
	}

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
