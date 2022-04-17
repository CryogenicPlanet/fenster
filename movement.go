package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// This file is to assist with movement
// Instead of writing bash I'd rather it just be in go, even if it a bit less performant

// ctrl + shift + right -> move focused window to the next avaliable space
func HandleRight() error {

	_, _, err := runShellSync("yabai -m window --space next && yabai -m --space --focus next")

	if err != nil {

		// Have to do some manual work now

		// Find the next avaliable space and move there

		out, _, err := runShellSync("yabai -m query --spaces")

		if err != nil {

			log.Printf("error: %v\n", err)
		}

		spaces := []SpaceQuery{}

		err = json.Unmarshal([]byte(out), &spaces)
		if err != nil {

			log.Printf("error: %v\n", err)
		}

		display := -1
		focusedElement := -1
		for idx, space := range spaces {
			if space.Focus {
				// This the focused space
				focusedElement = space.Index
				display = space.Display
			} else if focusedElement != -1 {
				if (idx > focusedElement) && (space.Display == display) && !space.Fullscreen {
					avaliableSpace := space.Index

					cmd := fmt.Sprintf("yabai -m window --space %v && yabai -m space --focus %v", avaliableSpace, avaliableSpace)

					fmt.Println(cmd)

					_, _, err := runShellSync(cmd)

					if err != nil {

						log.Printf("error: %v\n", err)
						return err
					}

					return nil
				}
			}
		}
		return fmt.Errorf("no avaliable space")
	}
	return nil

}

// ctrl + shift + right -> move focused window to the previous avaliable space
func HandleLeft() error {

	_, _, err := runShellSync("yabai -m window --space previous && yabai -m --space --focus previous")

	if err != nil {

		// Have to do some manual work now

		// Find the next avaliable space and move there

		out, _, err := runShellSync("yabai -m query --spaces")

		if err != nil {

			log.Printf("error: %v\n", err)
		}

		spaces := []SpaceQuery{}

		err = json.Unmarshal([]byte(out), &spaces)
		if err != nil {

			log.Printf("error: %v\n", err)
		}

		display := -1
		focusedElement := -1

		for idx := len(spaces) - 1; idx >= 0; idx-- {
			space := spaces[idx]

			fmt.Println(space.Index, focusedElement)

			if space.Focus {
				// This the focused space
				focusedElement = space.Index
				display = space.Display
			} else if focusedElement != -1 {
				if idx < focusedElement && space.Display == display && !space.Fullscreen {
					avaliableSpace := space.Index

					cmd := fmt.Sprintf("yabai -m window --space %v && yabai -m space --focus %v", avaliableSpace, avaliableSpace)

					fmt.Println(cmd)

					_, _, err := runShellSync(cmd)

					if err != nil {

						log.Printf("error: %v\n", err)
						return err
					}

					return nil
				}
			}
		}
		return fmt.Errorf("no avaliable space")
	}
	return nil

}
