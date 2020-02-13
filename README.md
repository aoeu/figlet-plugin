
figlet-plugin is a [Mattermost](https://en.wikipedia.org/wiki/Mattermost) plugin for altering text via [FIGlet](https://en.wikipedia.org/wiki/FIGlet).

This plugin is self-contained and uses the actual FIGlet binary, built from the official distribution on Linux with x86 architecture.  
*Note that this means the plugin will only work if your Mattermost server is also running Linux with x86.*

```
usage: /figlet [list|help|fonts] [font name] text
```

FIGlet is a program for making large letters out of ordinary text

```
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

```

with the following named fonts:

```
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

```
