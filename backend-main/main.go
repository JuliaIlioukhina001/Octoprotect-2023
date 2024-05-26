package main

import (
	"backend/config"
	"backend/model"
	"backend/router"
	"fmt"
)

func licenseInfo() {
	fmt.Println("Libraries license and copyright notice:")
	fmt.Println("github.com/fsnotify/fsnotify,https://github.com/fsnotify/fsnotify/blob/v1.6.0/LICENSE,BSD-3-Clause\ngithub.com/gabriel-vasile/mimetype,https://github.com/gabriel-vasile/mimetype/blob/v1.4.2/LICENSE,MIT\ngithub.com/gin-contrib/sse,https://github.com/gin-contrib/sse/blob/v0.1.0/LICENSE,MIT\ngithub.com/gin-gonic/gin,https://github.com/gin-gonic/gin/blob/v1.9.1/LICENSE,MIT\ngithub.com/go-playground/locales,https://github.com/go-playground/locales/blob/v0.14.1/LICENSE,MIT\ngithub.com/go-playground/universal-translator,https://github.com/go-playground/universal-translator/blob/v0.18.1/LICENSE,MIT\ngithub.com/go-playground/validator/v10,https://github.com/go-playground/validator/blob/v10.14.0/LICENSE,MIT\ngithub.com/golang-jwt/jwt/v5,https://github.com/golang-jwt/jwt/blob/v5.1.0/LICENSE,MIT\ngithub.com/google/uuid,https://github.com/google/uuid/blob/v1.1.2/LICENSE,BSD-3-Clause\ngithub.com/gorilla/websocket,https://github.com/gorilla/websocket/blob/v1.5.1/LICENSE,BSD-3-Clause\ngithub.com/hashicorp/hcl,https://github.com/hashicorp/hcl/blob/v1.0.0/LICENSE,MPL-2.0\ngithub.com/jackc/pgpassfile,https://github.com/jackc/pgpassfile/blob/v1.0.0/LICENSE,MIT\ngithub.com/jackc/pgservicefile,https://github.com/jackc/pgservicefile/blob/091c0ba34f0a/LICENSE,MIT\ngithub.com/jackc/pgx/v5,https://github.com/jackc/pgx/blob/v5.4.3/LICENSE,MIT\ngithub.com/jinzhu/inflection,https://github.com/jinzhu/inflection/blob/v1.0.0/LICENSE,MIT\ngithub.com/jinzhu/now,https://github.com/jinzhu/now/blob/v1.1.5/License,MIT\ngithub.com/leodido/go-urn,https://github.com/leodido/go-urn/blob/v1.2.4/LICENSE,MIT\ngithub.com/magiconair/properties,https://github.com/magiconair/properties/blob/v1.8.7/LICENSE.md,BSD-2-Clause\ngithub.com/mattn/go-isatty,https://github.com/mattn/go-isatty/blob/v0.0.19/LICENSE,MIT\ngithub.com/mattn/go-sqlite3,https://github.com/mattn/go-sqlite3/blob/v1.14.17/LICENSE,MIT\ngithub.com/mitchellh/mapstructure,https://github.com/mitchellh/mapstructure/blob/v1.5.0/LICENSE,MIT\ngithub.com/pelletier/go-toml/v2,https://github.com/pelletier/go-toml/blob/v2.1.0/LICENSE,MIT\ngithub.com/sagikazarmark/slog-shim,https://github.com/sagikazarmark/slog-shim/blob/v0.1.0/LICENSE,BSD-3-Clause\ngithub.com/sirupsen/logrus,https://github.com/sirupsen/logrus/blob/v1.9.3/LICENSE,MIT\ngithub.com/spf13/afero,https://github.com/spf13/afero/blob/v1.10.0/LICENSE.txt,Apache-2.0\ngithub.com/spf13/cast,https://github.com/spf13/cast/blob/v1.5.1/LICENSE,MIT\ngithub.com/spf13/pflag,https://github.com/spf13/pflag/blob/v1.0.5/LICENSE,BSD-3-Clause\ngithub.com/spf13/viper,https://github.com/spf13/viper/blob/v1.17.0/LICENSE,MIT\ngithub.com/subosito/gotenv,https://github.com/subosito/gotenv/blob/v1.6.0/LICENSE,MIT\ngithub.com/ugorji/go/codec,https://github.com/ugorji/go/blob/codec/v1.2.11/codec/LICENSE,MIT\ngolang.org/x/crypto,https://cs.opensource.google/go/x/crypto/+/v0.14.0:LICENSE,BSD-3-Clause\ngolang.org/x/net,https://cs.opensource.google/go/x/net/+/v0.17.0:LICENSE,BSD-3-Clause\ngolang.org/x/sys/unix,https://cs.opensource.google/go/x/sys/+/v0.13.0:LICENSE,BSD-3-Clause\ngolang.org/x/text,https://cs.opensource.google/go/x/text/+/v0.13.0:LICENSE,BSD-3-Clause\ngoogle.golang.org/protobuf,https://github.com/protocolbuffers/protobuf-go/blob/v1.31.0/LICENSE,BSD-3-Clause\ngopkg.in/ini.v1,https://github.com/go-ini/ini/blob/v1.67.0/LICENSE,Apache-2.0\ngopkg.in/yaml.v3,https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE,MIT\ngorm.io/driver/postgres,https://github.com/go-gorm/postgres/blob/v1.5.4/License,MIT\ngorm.io/driver/sqlite,https://github.com/go-gorm/sqlite/blob/v1.5.4/License,MIT\ngorm.io/gorm,https://github.com/go-gorm/gorm/blob/v1.25.5/LICENSE,MIT\n")
}

func main() {
	licenseInfo()
	config.LoadConfig()
	model.InitDB()
	router.ListenHTTP()
}
