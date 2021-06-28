// This file is part of ezBastion.

//     ezBastion is free software: you can redistribute it and/or modify
//     it under the terms of the GNU Affero General Public License as published by
//     the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.

//     ezBastion is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU Affero General Public License for more details.

//     You should have received a copy of the GNU Affero General Public License
//     along with ezBastion.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"ezBastion/cmd/ezb_sta/ctrl"
	"ezBastion/cmd/ezb_sta/middleware"
	"ezBastion/pkg/logmanager"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
)

// Must implement Mainservice interface from servicemanager package
type mainService struct{}

func (sm mainService) StartMainService(serverchan *chan bool) {
	logmanager.Debug("#### Main service started #####")
	// Pushing current conf to controllers
	server := gin.Default()

	server.Use(func(c *gin.Context) {
		c.Set("configuration", conf)
		c.Set("exPath", exePath)
	})

	server.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	})

	server.OPTIONS("*a", func(c *gin.Context) {
		c.AbortWithStatus(200)
	})

	// Middleware
	server.Use(middleware.EzbAuthJWT)
	server.Use(middleware.EzbAuthform)
	server.Use(middleware.EzbAuthbasic)
	// token endpoint
	server.POST("/token", ctrl.Createtoken)
	//route.POST("/token", middleware.EzbCache)
	server.GET("/token", ctrl.Createtoken)
	server.GET("/renew", ctrl.Renewtoken)

	server.RunTLS(":"+strconv.Itoa(conf.EZBSTA.Network.Port), path.Join(exePath, conf.TLS.PublicCert), path.Join(exePath, conf.TLS.PrivateKey))
}