package funfacts

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string) []string {
	funFacts := FunFacts{
		Sun: []string{
			"The Sun is over 4.5 billion years old.",
			"The Sun's light reaches the Earth in eight minutes.",
			"The Sun is made from Hydrogen and Helium.",
		},
		Luna: []string{
			"The Moon always shows Earth the same face.",
			"The Moon's surface is actually dark.",
			"The Moon has quakes too.",
		},
		Terra: []string{
			"Earth has never been perfectly round",
			"Earth's gravity isn't uniform.",
			"In the past, sea levels were very different.",
		},
	}

	switch about {
	case "sun":
		return funFacts.Sun
	case "luna":
		return funFacts.Luna
	case "terra":
		return funFacts.Terra
	default:
		return []string{}
	}
}
