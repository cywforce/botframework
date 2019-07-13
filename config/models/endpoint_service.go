package models

/**
 * Defines an endpoint service connection.
 */
//TODO:  extends ConnectedService implements IEndpointService
type EndpointService struct {
	/**
	 * MSA App ID.
	 */
	appId string

	/**
	 * MSA app password for the bot.
	 */
	appPassword string

	/**
	 * Endpoint of localhost service.
	 */
	endpoint string

	/**
	 * The channel service (Azure or US Government Azure) for the bot.
	 * A value of "https://botframework.azure.us" means the bot will be talking to a US Government Azure data center.
	 * An undefined or null value means the bot will be talking to public Azure
	 */
	channelService string
}

/**
 * Creates a new EndpointService instance.
 * @param source JSON based service definition.
 */
func constructor(source IEndpointService) {
super(source, ServiceTypes.Endpoint);
}

func encrypt(secret string, encryptString string) {
if (this.appPassword && this.appPassword.length > 0) {
this.appPassword = encryptString(this.appPassword, secret);
}
}

func decrypt(secret string, decryptString string){
if (this.appPassword && this.appPassword.length > 0) {
this.appPassword = decryptString(this.appPassword, secret);
}
}


