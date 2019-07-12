package config

/**
 * @private
 * This is class which allows you to manipulate in memory representations of bot configuration
 * with no nodejs depedencies.
 */

type IResource struct {
	// unique Id for the service in the bot
	id string

	// Friendly name for the service
	name string
}

type BotRecipe struct {
	/**
	 * Version of the recipe.
	 */
	version string `json:"version"`

	/**
	 *
	 */
	resources []string `json:"resources"`
	/**
	 * Creates a new BotRecipe instance.
	 */
}

func Init() {
	// noop
}

func fromJSON(botRecipe BotRecipe) BotRecipe {
	recipe := &BotRecipe{
		version: "1.0",

	}
	if len(recipe.resources) != 0 {
		botRecipe.resources = recipe.resources
	} else {
		botRecipe.resources = botRecipe.resources
	}
	if recipe.version != "" {
		botRecipe.version = recipe.version
	} else {
		botRecipe.version = botRecipe.version
	}

	return botRecipe;
}

func toJSON() BotRecipe {
	recipe := new(BotRecipe)
	return *recipe
}
