package dialog

import (
	"github.com/FrancescoIlario/docker-images-interactive-eraser/internal/images"
	"github.com/manifoldco/promptui"
)

func askPrune(img *images.Image, tags []images.Tag) (bool, error) {
	pruneConfirm := promptui.Prompt{
		Label:     "Do you want to prune the children of the image",
		IsConfirm: true,
	}
	pruneResult, err := pruneConfirm.Run()
	if err != nil {
		if err == promptui.ErrAbort {
			return false, nil
		}
		return false, err
	}
	return pruneResult == "y", nil
}
