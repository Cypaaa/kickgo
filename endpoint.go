package kickgo

const (
	endpointWSScheme          = "wss"
	endpointWSHost            = "ws-us2.pusher.com"
	endpointWSPath            = "app/eb1d5f283081a78b932c"
	endpointWSRawQuery        = "protocol=7&client=js&version=7.4.0&flash=false"
	endpointKick              = "https://kick.com/"
	endpointApiCurrentVersion = "v1/"
	endpointAPI               = endpointKick + "api/"
	endpointAPIV              = endpointAPI + endpointApiCurrentVersion
	endpointLivestream        = endpointAPIV + "channels/%v"
)
