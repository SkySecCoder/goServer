package customTypes

import (
)

type Config struct {
	Database 					map[string]string	`json:database`
	SecureTransportProtocol		bool 				`json:secureTransportProtocol`	
}
