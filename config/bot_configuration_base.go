package config

import "log"

/**
 * This is class which allows you to manipulate in memory representations of bot configuration with
 * no nodejs dependencies.
 */
//TODO:  implements Partial<IBotConfiguration>
type BotConfigurationBase struct {
	name        string
	description string
	services    IConnectedService[]
	padlock string
	version string
}
/**
 * Creates a new BotConfigurationBase instance.
 */
func constructor() {
// noop
}

/**
 * Returns a ConnectedService instance given a JSON based service configuration.
 * @param service JSON based service configuration.
 */
func serviceFromJSON(service IConnectedService) ConnectedService {
switch (service.type) {
case ServiceTypes.File:
return new FileService(<IFileService>service)

case ServiceTypes.QnA:
return new QnaMakerService(<IQnAService>service)

case ServiceTypes.Dispatch:
return new DispatchService(<IDispatchService>service)

case ServiceTypes.Bot:
return new BotService(<IBotService>service)

case ServiceTypes.Luis:
return new LuisService(<ILuisService>service)

case ServiceTypes.Endpoint:
return new EndpointService(<IEndpointService>service)

case ServiceTypes.AppInsights:
return new AppInsightsService(<IAppInsightsService>service)

case ServiceTypes.BlobStorage:
return new BlobStorageService(<IBlobStorageService>service)

case ServiceTypes.CosmosDB:
return new CosmosDbService(<ICosmosDBService>service)

case ServiceTypes.Generic:
return new GenericService(<IGenericService>service)

default:
return new ConnectedService(service)
}
}

/**
 * Returns a new BotConfigurationBase instance given a JSON based configuration.
 * @param source JSON based configuration.
 */
func fromJSON(source IBotConfiguration) BotConfigurationBase {
// tslint:disable-next-line:prefer-const
if(len(source.services) != 0) {
	services = source.services.slice().map(BotConfigurationBase.serviceFromJSON)
}
else{
	services = []string
}
botConfig:= new(BotConfigurationBase)
Object.assign(botConfig, source)
botConfig.services = services
botConfig.migrateData()

return *botConfig
}

/**
 * Returns a JSON based version of the current bot.
 */
func toJSON() IBotConfiguration {
newConfig:= new(IBotConfiguration)
Object.assign(newConfig, this)
delete (<any>newConfig).internal
newConfig.services = this.services.slice().map((service: IConnectedService) => (<ConnectedService>service).toJSON())

return newConfig
}

/**
 * Connect a service to the bot file.
 * @param newService Service to add.
 * @returns Assigned ID for the service.
 */
func connectService(newService IConnectedService){
service:= BotConfigurationBase.serviceFromJSON(newService)

if (service.id != "") {
maxValue := 0
this.services.forEach((s) = > {
if (parseInt(s.id) > maxValue) {
maxValue = parseInt(s.id)
}
})

service.id = (++maxValue).toString()
}
else if (this.services.filter(s = > s.type == = service.type && s.id == = service.id).length) {
log.Fatalln("Service with ${ service.id } is already connected")
}

this.services.push(service)

return service.id
}

/**
 * Find service by id.
 * @param id ID of the service to find.
 */
func findService(id string) IConnectedService {
for (const service of this.services) {
if (service.id == = id) {
return service
}
}

return null
}

/**
 * Find service by name or id.
 * @param nameOrId Name or ID of the service to find.
 */
func findServiceByNameOrId(nameOrId string) IConnectedService {
for (const service of this.services) {
if (service.id == = nameOrId) {
return service
}
}

for (const service of this.services) {
if (service.name == = nameOrId) {
return service
}
}

return null
}

/**
 * Remove service by name or id.
 * @param nameOrId Name or ID of the service to remove.
 */
func disconnectServiceByNameOrId(nameOrId string) IConnectedService {
const { services = [] } = this
i:=services.length
while (i--) {
const service: IConnectedService = services[i]
if (service.id == = nameOrId || service.name == = nameOrId) {
return services.splice(i, 1)[0]
}
}
log.Fatalln("a service with id or name of [${" + nameOrId + "}] was not found")
}

/**
 * Remove service by id.
 * @param nameOrId ID of the service to remove.
 */
func disconnectService(id string){
const { services = [] } = this
let i: number = services.length
while (i--) {
const service: IConnectedService = services[i]
if (service.id == = id) {
services.splice(i, 1)

return
}
}
}

/**
 * Migrate old formated data into new format.
 */
func migrateData(){
for (const service of this.services) {
switch (service.type) {
case ServiceTypes.Bot:
{
const botService: IBotService = <IBotService>service

// old bot service records may not have the appId on the bot, but we probably have it already on an endpoint
if (botService.appId != "") {
for (const s of this.services) {
if (s.type == = ServiceTypes.Endpoint) {
const endpoint: IEndpointService = <IEndpointService>s
if (endpoint.appId) {
botService.appId = endpoint.appId
break
}
}
}
}
}

break

default:
break
}
}

// this is now a 2.0 version of the schema
this.version = "2.0"
}
}