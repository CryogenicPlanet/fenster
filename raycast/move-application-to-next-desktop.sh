#!/bin/bash

# Required parameters:
# @raycast.schemaVersion 1
# @raycast.title Move application to next desktop
# @raycast.mode silent

# Optional parameters:
# @raycast.icon 🤖
# @raycast.packageName Window management

# Documentation:
# @raycast.author Rahul Tarak
# @raycast.authorURL https://github.com/CryogenicPlanet

yabai -m window --display next && yabai -m display --focus next

