module learn.go

go 1.17

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/spf13/cobra v1.3.0
	learn.go.tools v0.0.0-00010101000000-000000000000
)

require (
	github.com/gin-contrib/pprof v1.3.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.6.2 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.2.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.3.2 // indirect
	gorm.io/gorm v1.23.1 // indirect

)

replace learn.go.tools => ../learn.go.tools //本地的包模块需要去使用

replace github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi

//replace (
//	github.com/spf13/cobra =>github.com/spf13/cobra v1.2.0
//)
