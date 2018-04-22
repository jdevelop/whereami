package resolver

import (
	"fmt"
	"net"
	"strings"

	"github.com/vishvananda/netlink"
)

type IPResolver interface {
	GetAllAddresses() ([]net.IPNet, error)
}

type RouteResolver interface {
	GetDefaultRouteSrc() (net.IP, error)
}

type DefaultResolver struct{}

var def = DefaultResolver{}

func New() (*DefaultResolver, error) {
	return &def, nil
}

func (dr *DefaultResolver) GetAllAddresses() ([]net.IPNet, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	ips := make([]net.IPNet, 0, len(ifaces)*2)
	for _, iface := range ifaces {
		if strings.HasPrefix(iface.Name, "lo") {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				if v.IP.To4() != nil {
					ips = append(ips, *v)
				}
			}
		}
	}
	return ips, nil
}

func (dr *DefaultResolver) GetDefaultRouteSrc() (net.IP, error) {
	h, err := netlink.NewHandle()
	if err != nil {
		return nil, err
	}
	links, err := h.LinkList()
	if err != nil {
		return nil, err
	}
	for _, lnk := range links {
		rte, err := h.RouteList(lnk, netlink.FAMILY_V4)
		if err != nil {
			return nil, err
		}
		for _, rt := range rte {
			if rt.Dst == nil {
				return rt.Src, nil
			}
		}
	}

	return nil, fmt.Errorf("Default route not found")
}

var _ RouteResolver = &def
var _ IPResolver = &def
