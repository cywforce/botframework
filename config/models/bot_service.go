package models

//TODO: extends AzureService implements IBotService

type BotService struct {
	/**
	 * MSA App ID for the bot.
	 */
	appId string
}

/**
 * Creates a new BotService instance.
 * @param source (Optional) JSON based service definition.
 */
func constructor(source IBotService) {
super(source, ServiceTypes.Bot);
}

func encrypt(secret string, encryptString string) {
return
}

func decrypt(secret string, decryptString string){
return
}

