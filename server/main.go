package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
		Trigger:          "figlet",
		DisplayName:      "FIGlet",
		Description:      description,
		AutoComplete:     true,
		AutoCompleteDesc: "/figlet optionalFontName this text will be printed in large letters",
		AutoCompleteHint: "[fontname] text",
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

	f.fontsPath, err = filepath.Abs(filepath.Join(bundlePath, "assets", "fonts"))
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
		Text:         t,
		Username:     args.UserId,
		ChannelId:    args.ChannelId,
		TriggerId:    args.TriggerId,
		ResponseType: "in_channel", // ResponseType is Required.
	}, nil
}

func (f FIGlet) transformText(in string) (out string, err error) {
	s := strings.Fields(in)
	if len(s) < 2 {
		s := "no additional text provided to transform in '%v'"
		return "", fmt.Errorf(s, in)
	}
	in = strings.Replace(in, "/figlet", "", 1)
	in = strings.TrimSpace(in)
	font := strings.ToLower(s[1])
	var cmd *exec.Cmd
	if fontNames[font] {
		s := strings.Replace(in, font, "", 1)
		s = strings.TrimSpace(s)
		cmd = exec.Command(f.binaryPath, "-d", f.fontsPath, "-f", font, s)
	} else {
		cmd = exec.Command(f.binaryPath, "-d", f.fontsPath, in)
	}
	cmd.Stdin = os.Stdin
	b, err := cmd.CombinedOutput()
	if err != nil {
		s := "error occured when running FIGlet and reading output '%v': %v"
		return "", fmt.Errorf(s, string(b), err)
	}
	markdownProof := strings.Replace(string(b), "`", "'", -1)
	return markdownProof, nil
}

var fontNames = map[string]bool{
	"banner":    true,
	"big":       true,
	"black":     true,
	"block":     true,
	"bubble":    true,
	"digital":   true,
	"ivrit":     true,
	"lean":      true,
	"mini":      true,
	"mnemonic":  true,
	"script":    true,
	"shadow":    true,
	"slant":     true,
	"small":     true,
	"smscript":  true,
	"smshadown": true,
	"smslant":   true,
	"standard":  true,
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


with the following named fonts:

banner :

#####    ##   #    # #    # ###### #####
#    #  #  #  ##   # ##   # #      #    #
#####  #    # # #  # # #  # #####  #    #
#    # ###### #  # # #  # # #      #####
#    # #    # #   ## #   ## #      #   #
#####  #    # #    # #    # ###### #    #



big :
 _     _
| |   (_)
| |__  _  __ _
| '_ \| |/ _' |
| |_) | | (_| |
|_.__/|_|\__, |
          __/ |
         |___/


block :

_|        _|                      _|
_|_|_|    _|    _|_|      _|_|_|  _|  _|
_|    _|  _|  _|    _|  _|        _|_|
_|    _|  _|  _|    _|  _|        _|  _|
_|_|_|    _|    _|_|      _|_|_|  _|    _|




bubble :
  _   _   _   _   _   _
 / \ / \ / \ / \ / \ / \
( b | u | b | b | l | e )
 \_/ \_/ \_/ \_/ \_/ \_/


digital :
+-+-+-+-+-+-+-+
|d|i|g|i|t|a|l|
+-+-+-+-+-+-+-+


ivrit :
                                                            _   _            _
                                                           | |_(_)_ ____   _(_)
                                                           | __| | '__\ \ / / |
                                                           | |_| | |   \ V /| |
                                                            \__|_|_|    \_/ |_|



lean :

    _/
   _/    _/_/      _/_/_/  _/_/_/
  _/  _/_/_/_/  _/    _/  _/    _/
 _/  _/        _/    _/  _/    _/
_/    _/_/_/    _/_/_/  _/    _/




mini :

._ _ o._ o
| | ||| ||


script :

               o
 ,   __   ,_        _ _|_
/ \_/    /  |  |  |/ \_|
 \/ \___/   |_/|_/|__/ |_/
                 /|
                 \|


shadow :
      |               |
  __| __ \   _' |  _' |  _ \\ \  \   /
\__ \ | | | (   | (   | (   |\ \  \ /
____/_| |_|\__,_|\__,_|\___/  \_/\_/



slant :
         __            __
   _____/ /___ _____  / /_
  / ___/ / __ '/ __ \/ __/
 (__  ) / /_/ / / / / /_
/____/_/\__,_/_/ /_/\__/



small :
               _ _
 ____ __  __ _| | |
(_-< '  \/ _' | | |
/__/_|_|_\__,_|_|_|



smscript :

 ,           ,   _   ,_  o    _|_
/ \_/|/|/|  / \_/   /  | | |/\_|
 \/  | | |_/ \/ \__/   |/|/|_/ |_/
                          (|


smshadow :
               |              |
(_-<  ' \ (_-<   \   _' |  _' |  _ \\ \  \ /
___/_|_|_|___/_| _|\__,_|\__,_|\___/ \_/\_/



smslant :
                 __          __
  ___ __ _  ___ / /__ ____  / /_
 (_-</  ' \(_-</ / _ '/ _ \/ __/
/___/_/_/_/___/_/\_,_/_//_/\__/



standard :
     _                  _               _
 ___| |_ __ _ _ __   __| | __ _ _ __ __| |
/ __| __/ _' | '_ \ / _' |/ _' | '__/ _' |
\__ \ || (_| | | | | (_| | (_| | | | (_| |
|___/\__\__,_|_| |_|\__,_|\__,_|_|  \__,_|

`
