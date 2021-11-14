package lib

type Content struct {
	Topic       string // visualization
	Name        string // Visualization
	Description string // Relaxing via visualization ....
}

func DemoContent() Content {
	return Content{
		Topic:       "visualization",
		Name:        "Visualization",
		Description: "Relaxing via visualization techniques, like closing your eyes and imagining a beach.",
	}
}
