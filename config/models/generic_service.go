package models

/**
 * Defines a generic service connection.
 */
//TODO:  extends ConnectedService implements IGenericService
type GenericService struct {
	/**
	 * Deep link to service.
	 */
	url string

	/**
	 * Named/value configuration data.
	 */
	configuration string
}

/**
 * Creates a new GenericService instance.
 * @param source (Optional) JSON based service definition.
 */
func constructor(source IGenericService) {
	super(source, ServiceTypes.Generic);
}

func encrypt(secret string, encryptString string) {
	this := new(GenericService)
	if (this.configuration != "") {
		Object.keys(this.configuration).forEach((prop: string) = > {
		that.configuration[prop] = encryptString(that.configuration[prop], secret);
		});
	}
}

func decrypt(secret string, decryptString string) {
	this := new(GenericService)
	if (this.configuration != "") {
		Object.keys(this.configuration).forEach((prop: string) = > {
		that.configuration[prop] = decryptString(that.configuration[prop], secret);
		});
	}
}
