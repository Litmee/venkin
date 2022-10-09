package core

import (
	"encoding/json"
	"net/http"
	"reflect"
)

// Http request parameter JSON inspection and serialization
func (cI *ControllerImpl) jsonCheckAndSerialize(c Controller) bool {
	// Reflection obtains the real structure type of the incoming controller interface type data
	rV := reflect.TypeOf(c).Elem()
	// Determine whether the data in the request needs to be serialized
	// According to the principle of convention over configuration
	// The control layer structure declared by the framework user needs to combine at least the ControllerImpl structure
	// If you need to get the data in the http request, you need to additionally combine a user-defined structure to receive JSON data
	// And the user's control layer structure must first import the ControllerImpl structure
	// And then import the required JSON serialization structure
	// So if the NumField result of the rV reflection parameter is greater than 1
	// The framework defaults to that the user needs to receive the data in the request and needs JSON serialization
	if rV.NumField() > 1 {
		model := reflect.New(rV.Field(1).Type).Interface()
		var buf [512]byte
		n, _ := cI.r.Body.Read(buf[:])
		if n > 0 {
			err := json.Unmarshal(buf[:n], &model)
			// JSON serialization error returns 500 error
			if err != nil {
				cI.w.WriteHeader(http.StatusInternalServerError)
				cI.w.Write([]byte("An exception occurred in JSON serialization, please confirm the Content-Type type and data structure of the request"))
<<<<<<< HEAD
=======
				logger.LogJsonSerialize(err)
>>>>>>> 28d998696ef8e199d5cc21734002cfaaf9e71bd3
				return false
			}
		}
		// Mount the serialized data back to ControllerImpl
		cI.data = model
		return true
	}
	return true
}
