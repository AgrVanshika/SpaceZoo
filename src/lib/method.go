package lib

/*
Categories:

expression
relaxation
visualization
reflection
meditation
breathing
*/

type Method struct {
	Topic       string
	Name        string
	Description string
	Source      string
}

type Technique struct {
	Name        string
	Description string
	Source      string
}

func ListMethods(category string) []Technique {

	switch category {
	case "expression":
		return []Technique{
			{
				Name:        "Art therapy",
				Description: "Draw your emotions or express your feelings in a painting.",
				Source:      "https://www.healthline.com/health/mental-health/anxiety-drawing#:~:text=Art%20therapy%20refers%20to%20any,have%20trouble%20putting%20into%20words.",
			},
			{
				Name:        "Therapeutic journaling",
				Description: "Writing your way through your thoughts is a great way to reflect on your emotions and experiences.",
				Source:      "https://www.va.gov/WHOLEHEALTHLIBRARY/tools/therapeutic-journaling.asp",
			},
			{
				Name:        "Music therapy",
				Description: "Singing a song or playing an instrument is a great way to express your emotions and release stress.",
				Source:      "https://www.healthline.com/health/sound-healing",
			},
		}
	case "relaxation":
		return []Technique{
			{
				Name:        "Sand tray therapy",
				Description: "Raking the sand of a mini zen garden gives you a way to express yourself and reduce tension.",
				Source:      "https://www.healthyplace.com/blogs/anxiety-schmanxiety/2015/07/hey-kids-and-adults-play-with-sand-to-reduce-anxiety",
			},
			{
				Name:        "Progressive muscle relaxation",
				Description: "Relax your muscles one group at a time, releasing your stress.",
				Source:      "https://www.uofmhealth.org/health-library/uz2225",
			},
			{
				Name:        "Relaxing soundscapes",
				Description: "A relaxing soundscape, whether it''s the beach, traffic, or birds chirping in the woods, can help put your mind at ease.",
				Source:      "https://www.drjohnlapuma.com/naturetherapy/what-is-sound-or-music-therapy-what-are-soundscapes/",
			},
		}
	case "visualization":
		return []Technique{
			{
				Name:        "Perfect day",
				Description: "What's your idea of a perfect day? Reflecting on and imagining it can help you reduce stress.",
				Source:      "https://projectlifemastery.com/design-your-perfect-day/",
			},
			{
				Name:        "Happy Place visualization",
				Description: "Imagine a peaceful place, like a beach or the woods. Imagine with all your senses: how it sounds, how it feels, how it smells.",
				Source:      "https://www.mondaycampaigns.org/destress-monday/find-happy-place",
			},
		}
	case "reflection":
		return []Technique{
			{
				Name:        "Reflect on your good traits",
				Description: "What are you good at? What do other people see in you? Reflect on your good traits to improve your self-image.",
				Source:      "https://www.healthyplace.com/blogs/buildingselfesteem/2015/10/identifying-your-good-qualities-when-you-believe-youre-worthless",
			},
			{
				Name:        "Reflect on your gratitute",
				Description: "What are you thankful for? How can you express it?",
				Source:      "https://resources.wellcertified.com/articles/practicing-gratitude-5-tips-for-positive-reflection/",
			},
			{
				Name:        "Self-reflection",
				Description: "Reflecting on what your priorities are and how you are working towards them can help you identify sources of stress and work towards reducing them.",
				Source:      "https://www.skillsyouneed.com/ps/reflective-practice.html",
			},
		}
	case "meditation":
		return []Technique{
			{
				Name:        "Meditation",
				Description: "Meditation is one of the best ways to take a break a relax.",
				Source:      "https://www.mindful.org/how-to-meditate/",
			},
			{
				Name:        "Visiualization meditation",
				Description: "Picturing positive images and relaxing your body is a good way to meditate if you struggle with traditional meditation.",
				Source:      "https://www.wellandgood.com/visualization-meditation/",
			},
		}
	case "breathing":
		return []Technique{
			{
				Name:        "Roll breathing",
				Description: "Lay on your back with one hand on your belly and the other on your chest. Feel your breath move through you like a wave as you breathe in and out.",
				Source:      "https://myhealth.alberta.ca/Health/Pages/conditions.aspx?hwid=ug4360",
			},
			{
				Name:        "Square breathing",
				Description: "Inhale through your nose for 4 seconds, hold for 4 seconds, exhale through your mouth for 4 seconds, hold for 4 seconds, and repeat.",
				Source:      "https://blog.zencare.co/square-breathing/",
			},
			{
				Name:        "Resonant breathing",
				Description: "Lie down, eyes shut, and breathe in slowly through your nose. Exhale slowly and repeat.",
				Source:      "https://www.rosalbacourtney.com/resonance-frequency-breathing/#:~:text=Resonance%20frequency%20breathing%20is%20a,such%20as%20the%20circulatory%20system",
			},
		}
	}

	return []Technique{
		{
			Name:        "Resonant breathing",
			Description: "Lie down, eyes shut, and breathe in slowly through your nose. Exhale slowly and repeat.",
			Source:      "https://www.rosalbacourtney.com/resonance-frequency-breathing/#:~:text=Resonance%20frequency%20breathing%20is%20a,such%20as%20the%20circulatory%20system",
		},
	}
}
