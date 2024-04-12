module github.com/forgeutah/meetup-replacement-project/server

go 1.21

require (
	github.com/catalystsquad/app-utils-go v1.0.7
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.16.0
)

replace github.com/forgeutah/meetup-replacement-project/protos => ./../protos