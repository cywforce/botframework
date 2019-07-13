package models

/**
 * Base class for all connected service definitions.
 */

//TODO: implement IConnectedService

type ConnectedService struct {
	/**
	 * Unique Id for the service.
	 */
	id string

	/**
	 * Friendly name for the service.
	 */
	name string
}

/**
 * Creates a new ConnectedService instance.
 * @param source (Optional) JSON based service definition.
 * @param type (Optional) type of service being defined.
 */
func constructor(source IConnectedService, type ServiceTypes) {
Object.assign(this, source);
if (type) {
this.type = type;
}
}

/**
 * Returns a JSON based version of the model for saving to disk.
 */
func toJSON() IConnectedService {
return <IConnectedService>Object.assign({}, this);
}

/**
 * Encrypt properties on this service.
 * @param secret Secret to use to encrypt the keys in this service.
 * @param encryptString Function called to encrypt an individual value.
 */
func encrypt(secret string, encryptString string) {
// noop
}

/**
 * Decrypt properties on this service.
 * @param secret Secret to use to decrypt the keys in this service.
 * @param decryptString Function called to decrypt an individual value.
 */
func decrypt(secret string, decryptString string) {
// noop
}
