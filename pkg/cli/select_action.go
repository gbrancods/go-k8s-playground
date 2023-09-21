package cli

import (
	"fmt"
	"go-k8s/pkg/controller"

	"github.com/manifoldco/promptui"
	"k8s.io/client-go/rest"
)

func SelectAction(cfg *rest.Config) {
	prompt := promptui.Select{
		Label: "Select your option",
		Items: []string{"Crud Demo",
			"Watch Demo",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Crud Demo":
		controller.Crud(cfg)
	case "Watch Demo":
		controller.Watch(cfg)
	}
}
