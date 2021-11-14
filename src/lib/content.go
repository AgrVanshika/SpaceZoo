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

func FillContent() []Content {
	return []Content{
		{
			Topic:       "visualization",
			Name:        "Visualization",
			Description: "Relaxing via visualization techniques, like closing your eyes and imagining a beach.",
		},
		{
			Topic:       "relaxation",
			Name:        "Relaxation",
			Description: "Techniques meant to envoke relaxation like muscle relaxation, sand tray therapy goes here.",
		},
		{
			Topic:       "expression",
			Name:        "Expression",
			Description: "Releasing stress by creating something, painting, writing and breathing.",
		},
		{
			Topic:       "reflection",
			Name:        "Reflection",
			Description: "Reflecting on parts of yourself, whether its your goals, what you're thankful for, or your priorities.",
		},
		{
			Topic:       "breathing",
			Name:        "Breathing",
			Description: "Using breathing tecniques such as Lion's breath, Focus breath and Progressive muscle relaxation.",
		},
		{
			Topic:       "meditation",
			Name:        "Meditation",
			Description: "Using Meditation techniques for certain time intervals to reduce and releive from stress.",
		},
	}
}
