package models

/**
 * Defines a dispatch service connection.
 */
//TODO:  extends LuisService implements IDispatchService
type DispatchService struct {
	/**
	 * Service IDs that the dispatch model will dispatch across.
	 */
	serviceIds []string
}

/**
 * Creates a new DispatchService instance.
 * @param source (Optional) JSON based service definition.
 */
func constructor(source IDispatchService) {
super(source, ServiceTypes.Dispatch);
dispatchService := new(DispatchService)
if (len(dispatchService.serviceIds) != 0){
	dispatchService.serviceIds = dispatchService.serviceIds
}
else{
	dispatchService.serviceIds = []string
	}
}


