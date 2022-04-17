# Fenster - MacOS Space Manager

The main functionality fenster provides is to have good full screen spaces, in the future it might provide more.


https://user-images.githubusercontent.com/10355479/163709379-b5c99dca-0123-41f3-ba43-e93c5b69e66e.mp4



## Deps
- [yabai](https://github.com/koekeishiya/yabai) - You need `yabai` to move spaces around which is critical for how Fenster works, in the future this might be something directly implemented
    For `yabai` to work with this feature, you need to [Disable System Integrity Protection](https://github.com/koekeishiya/yabai/wiki/Disabling-System-Integrity-Protection)

## Usage

Setup callback
```bash
yabai -m signal --add event=space_changed action="curl localhost:8090/prepare?oldSpace=\${YABAI_RECENT_SPACE_ID}&newSpace=\${YABAI_SPACE_ID}"

# In the future will be
# fenster setup # not implement yet
```

Run the fenster server
```
fenster start
```

## Run server on startup

```
crontab -e

@reboot fenster start
```

## Disclaimer

This is really just written for me, so it may not make a lot of sense for you to use it. 

Also the code here was written in like an hour and is really quite shit so don't judge me too much on it