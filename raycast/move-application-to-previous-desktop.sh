#!/bin/bash

# Required parameters:
# @raycast.schemaVersion 1
# @raycast.title Move application to previous window
# @raycast.mode silent

# Optional parameters:
# @raycast.icon 🤖
# @raycast.packageName Window management

# Documentation:
# @raycast.author Rahul Tarak
# @raycast.authorURL https://github.com/CryogenicPlanet

yabai -m window --display prev && yabai -m display --focus prev

