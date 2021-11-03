package server

import (
	"TasQ/app/serializers"
)

type Server struct {
	sz serializers.Serializer
}

// functional option
type ServerOption func(s *Server)

// set global serializer
func WithCustomSerializer(sz serializers.Serializer) ServerOption {
	return func(s *Server) {
		s.sz = sz
	}
}
