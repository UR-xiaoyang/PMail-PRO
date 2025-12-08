package imap_server

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/Jinnrry/pmail/config"
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapserver"
	log "github.com/sirupsen/logrus"
)

var instanceTLS *imapserver.Server

func Stop() {
	if instanceTLS != nil {
		instanceTLS.Close()
		instanceTLS = nil
	}
}

// StarTLS 启动TLS端口监听，不加密的代码就懒得写了
func StarTLS() {

	var tlsConfig *tls.Config

	if config.Instance.SSLType != config.SSLTypeNone {
		crt, err := tls.LoadX509KeyPair(config.Instance.SSLPublicKeyPath, config.Instance.SSLPrivateKeyPath)
		if err != nil {
			panic(err)
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{crt},
		}
	}

	memServer := NewServer()

	option := &imapserver.Options{
		NewSession: func(conn *imapserver.Conn) (imapserver.Session, *imapserver.GreetingData, error) {
			return memServer.NewSession(), nil, nil
		},
		Caps: imap.CapSet{
			imap.CapIMAP4rev1: {},
		},
		TLSConfig:    tlsConfig,
		InsecureAuth: config.Instance.SSLType == config.SSLTypeNone,
	}

	if config.Instance.LogLevel == "debug" {
		option.DebugWriter = os.Stdout
	}

	instanceTLS = imapserver.New(option)

	if config.Instance.SSLType != config.SSLTypeNone {
		var addr string
		if config.Instance.IMAPSPort == 0 {
			addr = ":993"
		} else {
			addr = fmt.Sprintf(":%d", config.Instance.IMAPSPort)
		}
		log.Infof("IMAP With TLS Server Start On Port %s", addr)
		if err := instanceTLS.ListenAndServeTLS(addr); err != nil {
			panic(err)
		}
	} else {
		var addr string
		if config.Instance.IMAPPort == 0 {
			addr = ":143"
		} else {
			addr = fmt.Sprintf(":%d", config.Instance.IMAPPort)
		}
		log.Infof("IMAP Server Start On Port %s", addr)
		if err := instanceTLS.ListenAndServe(addr); err != nil {
			panic(err)
		}
	}
}
