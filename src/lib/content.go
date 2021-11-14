package lib

type Content struct {
	Topic       string
	Name        string
	Description string
}

func DemoContent() Content {
	return Content{
		Topic:       "visualization",
		Name:        "Visualization",
		Description: "Relaxing via visualization techniques, like closing your eyes and imagining a beach.",
	}
}
