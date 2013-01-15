/***** BEGIN LICENSE BLOCK *****
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this file,
# You can obtain one at http://mozilla.org/MPL/2.0/.
#
# The Initial Developer of the Original Code is the Mozilla Foundation.
# Portions created by the Initial Developer are Copyright (C) 2012
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Rob Miller (rmiller@mozilla.com)
#   Victor Ng (vng@mozilla.com)
#
# ***** END LICENSE BLOCK *****/

package heka_mozsvc_plugins

import (
	"errors"
	"fmt"
	"github.com/mozilla-services/heka/pipeline"
	"net"
)

// This will be our pipeline.PluginGlobal type
type UdpOutputGlobal struct {
	conn net.Conn
}

// Provides pipeline.PluginGlobal interface
func (self *UdpOutputGlobal) Event(eventType string) {
	if eventType == pipeline.STOP {
		self.conn.Close()
	}
}

// This will be our PluginWithGlobal type
type UdpOutput struct {
	global *UdpOutputGlobal
}

// Initialize UDP connection, store it on the PluginGlobal
func (self *UdpOutput) InitOnce(config interface{}) (pipeline.PluginGlobal, error) {
	conf := config.(*pipeline.PluginConfig)
	addr, ok := (*conf)["address"]
	if !ok {
		return nil, errors.New("UdpOutput: No UDP address")
	}
	addrStr, ok := addr.(string)
	if !ok {
		return nil, errors.New("UdpOutput: UDP address not a string")
	}
	udpAddr, err := net.ResolveUDPAddr("udp", addrStr)
	if err != nil {
		return nil, fmt.Errorf("UdpOutput error resolving UDP address %s: %s",
			addrStr, err.Error())
	}
	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return nil, fmt.Errorf("UdpOutput error dialing UDP address %s: %s",
			addrStr, err.Error())
	}
	return &UdpOutputGlobal{udpConn}, nil
}

// Store a reference to the global for use during pipeline processing
func (self *UdpOutput) Init(global pipeline.PluginGlobal, config interface{}) error {
	self.global = global.(*UdpOutputGlobal) // UDP connection available as self.global.conn
	return nil
}
