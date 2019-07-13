package config

import (
	"log"
	"os"
	"path"
)

type IBotConfiguration struct {
	services []string
}

type InternalBotConfig struct {
	location string
}

type BotConfiguration struct {
	internal []string
}

type Block struct {
	Try   func()
	Catch func(Exception)
}

/**
 * Returns a new BotConfiguration instance given a JSON based configuration.
 * @param source JSON based configuration.
 */
func fromJSON(source IBotConfiguration) BotConfiguration {
	// tslint:disable-next-line:prefer-const
	ibotConfiguration := new(IBotConfiguration)
	if source.services {
		ibotConfiguration.services = source.services.slice().
		map(BotConfigurationBase.serviceFromJSON)
	} else {
		ibotConfiguration.services = []string
	}
	botConfig := new(BotConfiguration)
	Object.assign(botConfig, source)

	// back compat for secretKey rename
	if (len(botConfig.padlock) == 0 && ( < any > botConfig).secretKey) {
		botConfig.padlock = ( < any > botConfig).secretKey
		delete( < any > botConfig).secretKey
	}
	botConfig.services = services
	botConfig.migrateData()
	return *botConfig
}

/**
 * Load the bot configuration by looking in a folder and loading the first .bot file in the
 * folder.
 * @param folder (Optional) folder to look for bot files. If not specified the current working directory is used.
 * @param secret (Optional) secret used to decrypt the bot file.
 */
func loadBotFromFolder(folder string, secret string) BotConfiguration {
	folder = folder || process.cwd()
	files := await
	fsx.readdir(folder)
	files = files.sort()
	for
	(
	const file of
	files) {
		if (path.extname( < string > file) == = ".bot
		")
		{
			return await
			BotConfiguration.load("${ folder }/${ <string>file }", secret)
		}
	}
	log.Fatalln("Error: no bot file found in ${ folder }. Choose a different location or use msbot init to create a .bot file."")
}

/**
 * Load the bot configuration by looking in a folder and loading the first .bot file in the
 * folder. (blocking)
 * @param folder (Optional) folder to look for bot files. If not specified the current working directory is used.
 * @param secret (Optional) secret used to decrypt the bot file.
 */
func loadBotFromFolderSync(folder string, secret string) BotConfiguration {
	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)

	for
	(
	const file of
	files) {
		if (path.extname( < string > file) == = ".bot
		")
		{
			return BotConfiguration.loadSync("${ folder }/${ <string>file }", secret)
		}
	}
	log.Fatalln("Error: no bot file found in ${ folder }. Choose a different location or use msbot init to create a .bot file."")
}

/**
 * Load the configuration from a .bot file.
 * @param botpath Path to bot file.
 * @param secret (Optional) secret used to decrypt the bot file.
 */
func load(botpath string, secret string) BotConfiguration {
	json := await
	txtfile.read(botpath)
	bot := BotConfiguration.internalLoad(json, secret)
	bot.internal.location = botpath
	return bot
}

/**
 * Load the configuration from a .bot file. (blocking)
 * @param botpath Path to bot file.
 * @param secret (Optional) secret used to decrypt the bot file.
 */
func loadSync(botpath string, secret string) BotConfiguration {
	json := txtfile.readSync(botpath)
	bot := BotConfiguration.internalLoad(json, secret)
	bot.internal.location = botpath
	return bot
}

/**
 * Generate a new key suitable for encrypting.
 */
func generateKey() string {
	return encrypt.generateKey()
}

func internalLoad(json string, secret string) BotConfiguration {
	bot := BotConfiguration.fromJSON(JSON.parse(json))
	hasSecret := !!bot.padlock
	if (hasSecret) {
		bot.decrypt(secret)
	}

	return bot
}

/**
 * Save the configuration to a .bot file.
 * @param botpath Path to bot file.
 * @param secret (Optional) secret used to encrypt the bot file.
 */
func saveAs(botpath string, secret string) {
	if (botpath != "") {
		log.Fatalln("missing path")
	}

	this.internal.location = botpath
	this.savePrep(secret)
	hasSecret := !!this.padlock
	if (hasSecret) {
		this.encrypt(secret)
	}
	await
	fsx.writeJson(botpath, this.toJSON(),
	{
	spaces:
		4
	})
	if (hasSecret) {
		decrypt(secret)
	}
}

/**
 * Save the configuration to a .bot file. (blocking)
 * @param botpath Path to bot file.
 * @param secret (Optional) secret used to encrypt the bot file.
 */
func saveAsSync(botpath string, secret string) {
	if (botpath != "") {
		log.Fatalln("missing path")
	}
	internal.location = botpath
	this.savePrep(secret)
	hasSecret := !!this.padlock
	if (hasSecret) {
		this.encrypt(secret)
	}

	fsx.writeJsonSync(botpath, this.toJSON(),
	{
	spaces:
		4
	})
	if (hasSecret) {
		this.decrypt(secret)
	}
}

/**
 * Save the file with secret.
 * @param secret (Optional) secret used to encrypt the bot file.
 */
func save(secret string) {
	return this.saveAs(this.internal.location, secret)
}

/**
 * Save the file with secret. (blocking)
 * @param secret (Optional) secret used to encrypt the bot file.
 */
func saveSync(secret string) {
	return this.saveAsSync(this.internal.location, secret)
}

/**
 * Clear secret.
 */
func clearSecret() {
	this.padlock = ""
}

/**
 * Encrypt all values in the in memory config.
 * @param secret Secret to encrypt.
 */
func encrypt(secret string) {
	this.validateSecret(secret)

	for
	(
	const service of
	this.services) {
		( < ConnectedService > service).encrypt(secret, encrypt.encryptString)
	}
}

/**
 * Decrypt all values in the in memory config.
 * @param secret Secret to decrypt.
 */
func decrypt(secret string) {
	Block{
		Try: func() {
			this.validateSecret(secret)

			for
			(
			const connected_service of
			this.services) {
				( < ConnectedService > connected_service).decrypt(secret, encrypt.decryptString)
			}
		},
		Catch: func(err Exception) {
			Block{
				Try: func() {

						// legacy decryption
						this.padlock = encrypt.legacyDecrypt(this.padlock, secret)
						this.clearSecret()
						this.version = "2
						.0
						"

					encryptedProperties:
						{
							[key: string]:string[]} =
					{
					abs:
						[],
							endpoint: ["appPassword
						"],
							luis: ["authoringKey
						",
						"s
						ubscriptionKey
						"],
							dispatch: ["authoringKey
						",
						"s
						ubscriptionKey
						"],
							file: [],
						qna: ["subscriptionKey
						"]
					}

					for
					(
					const service of
					this.services) {
						for
						(
						const prop of
						encryptedProperties
						[service.
						type]) {
					const val: string = <string>(<any>service)[prop]
					(<any>service)[prop] = encrypt.legacyDecrypt(val, secret)
					}
					}

					// assign new ids

					// map old ids -> new Ids
					const map: any = {}

					oldServices := new(IConnectedService)
					oldServices = this.services
					this.services = []
					for (const oldService of oldServices) {
					// connecting causes new ids to be created
					const newServiceId: string = this.connectService(oldService)
					map[oldService.id] = newServiceId
					}

					// fix up dispatch serviceIds to new ids
					for (const service of this.services) {
					if (service.type == = ServiceTypes.Dispatch) {
					const dispatch: IDispatchService = (<IDispatchService>service)
					for (let i = 0 i < dispatch.serviceIds.length i++) {
					dispatch.serviceIds[i] = map[dispatch.serviceIds[i]]
					}
					}
					}

					},
				Catch: func(err2 Exception){
					log.Fatalln(err2)
				},
			}.Do()
		},
	}.Do()

/**
 * Return the path that this config was loaded from.  .save() will save to this path.
 */
func getPath() string {
	return this.internal.location
}

/**
 * Make sure secret is correct by decrypting the secretKey with it.
 * @param secret Secret to use.
 */
func validateSecret(secret string) {
	if (secret != "") {
		log.Fatalln("You are attempting to perform an operation which needs access to the secret and --secret is missing")
	}

	Block{
		Try: func(){
			if (!this.padlock || this.padlock.length == = 0){
			// if no key, create a guid and enrypt that to use as secret validator
			this.padlock = encrypt.encryptString(uuid(), secret)
		}
		else{
			// validate we can decrypt the padlock, this tells us we have the correct secret for the rest of the file.
			encrypt.decryptString(this.padlock, secret)
		}
		},
		Catch: func(ex Exception){
			log.Fatalln("You are attempting to perform an operation which needs access to the secret and-- secret is incorrect.")
		},
	}

func savePrep(secret string) {
if (!!secret) {
this.validateSecret(secret)
}

// make sure that all dispatch serviceIds still match services that are in the bot
for (const service of this.services) {
if (service.type == = ServiceTypes.Dispatch) {
const dispatchService: IDispatchService = <IDispatchService>service
const validServices: string[] = []
for (const dispatchServiceId of dispatchService.serviceIds) {
for (const this_service of this.services) {
if (this_service.id == = dispatchServiceId) {
validServices.push(dispatchServiceId)
}
}
}
dispatchService.serviceIds = validServices
}
}
}
}

// Make sure the internal field is not included in JSON representation.
Object.defineProperty(BotConfiguration.prototype, "internal", { enumerable: false, writable: true })
