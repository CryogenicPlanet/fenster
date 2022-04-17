# Fenster - MacOS Space Manager

The main functionality fenster provides is to have good full screen spaces, in the future it might provide more.

<div style="position: relative; padding-bottom: 42.1875%; height: 0;"><iframe src="https://www.loom.com/embed/60d24723621d40b1aa1e08fe29750b71" frameborder="0" webkitallowfullscreen mozallowfullscreen allowfullscreen style="position: absolute; top: 0; left: 0; width: 100%; height: 100%;"></iframe></div>

## Deps
- [yabai](https://github.com/koekeishiya/yabai) - You need `yabai` to move spaces around which is critical for how Fenster works, in the future this might be something directly implemented
    For `yabai` to work with this feature, you need to [Disable System Integrity Protection](https://github.com/koekeishiya/yabai/wiki/Disabling-System-Integrity-Protection)

## Usage

Setup callback
```
yabai -m signal --add event=space_changed action="curl localhost:8090/prepare?oldSpace=\${YABAI_RECENT_SPACE_ID}&newSpace=\${YABAI_SPACE_ID}"

# In the future will be

fenster setup # not implement yet

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
