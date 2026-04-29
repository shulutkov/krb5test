package krb5test

import "net"

type KDCOption func(*KDC)

func WithTCPListener(l net.Listener) KDCOption {
	return func(kdc *KDC) {
		kdc.TCPListener = l
	}
}
