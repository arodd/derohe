//go:build !wasm
// +build !wasm

// Copyright 2017-2021 DERO Project. All rights reserved.
// Use of this source code in any form is governed by RESEARCH license.
// license can be found in the LICENSE file.
// GPG: 0F39 E425 8C65 3947 702A  8234 08B2 0360 A03A 9DE8
//
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
// STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF
// THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package walletapi

// this file needs  serious improvements but have extremely limited time
/* this file handles communication with the daemon
 * this includes receiving output information
 *
 * *
 */
//import "io"
//import "os"
//import "fmt"

//import "net/url"
import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/deroproject/derohe/glue/rwc"
	"github.com/gorilla/websocket"
)

// there should be no global variables, so multiple wallets can run at the same time with different assset

var netClient *http.Client

type Client struct {
	WS  *websocket.Conn
	RPC *jrpc2.Client
}

var rpc_client = &Client{}

// this is as simple as it gets
// single threaded communication to get the daemon status and height
// this will tell whether the wallet can connection successfully to  daemon or not
func Connect(endpoint string) (err error) {

	Daemon_Endpoint_Active = get_daemon_address()

	logger.V(1).Info("Daemon endpoint ", "address", Daemon_Endpoint_Active)

	u, err := url.Parse(Daemon_Endpoint_Active)
	if err != nil {
		logger.Info(fmt.Sprintf("Error parsing endpoint %s", err))
	}

	// Check Urls
	is_url := regexp.MustCompile("^http")

	wsSchema := "ws://"
	if is_url.Match([]byte(Daemon_Endpoint_Active)) {
		if u.Scheme == "https" {
			wsSchema = "wss://"
			Daemon_Endpoint_Active = u.Hostname()
			if u.Port() != "" {
				Daemon_Endpoint_Active = Daemon_Endpoint_Active + ":" + u.Port()
			}
		}
	}

	dialer := websocket.DefaultDialer
	dialer.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	rpc_client.WS, _, err = dialer.Dial(wsSchema+Daemon_Endpoint_Active+"/ws", nil)

	// notify user of any state change
	// if daemon connection breaks or comes live again
	if err == nil {
		if !Connected {
			logger.V(1).Info("Connection to RPC server successful", "address", wsSchema+Daemon_Endpoint_Active+"/ws")
			Connected = true
		}
	} else {

		if Connected {
			logger.V(1).Error(err, "Connection to RPC server Failed", "endpoint", wsSchema+Daemon_Endpoint_Active+"/ws")
		}
		Connected = false
		return
	}

	input_output := rwc.New(rpc_client.WS)
	rpc_client.RPC = jrpc2.NewClient(channel.RawJSON(input_output, input_output), &jrpc2.ClientOptions{OnNotify: Notify_broadcaster})

	return test_connectivity()
}
