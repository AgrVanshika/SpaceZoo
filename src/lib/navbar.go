package lib

type Nav struct {
	Name string
	Link string
}

func NavBarFill() []Nav {
	return []Nav{
		{
			Name: "Visualization",
			Link: "visualization",
		},
		{
			Name: "Relaxation",
			Link: "relaxation",
		},
		{
			Name: "Expression",
			Link: "expression",
		},
		{
			Name: "Reflection",
			Link: "reflection",
		},
		{
			Name: "Breathing",
			Link: "breathing",
		},
		{
			Name: "Meditation",
			Link: "mediation",
		},
	}
}
