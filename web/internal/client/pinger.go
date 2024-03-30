package client

import (
	"fmt"
	"net/url"
	"time"

	"github.com/go-ping/ping"
)

type PingResult struct {
	IPAddr                string
	PacketsRecv           int
	PacketsSent           int
	PacketsRecvDuplicates int
	PacketLoss            float64
	Addr                  string
	Rtts                  []time.Duration
	MinRtt                time.Duration
	MaxRtt                time.Duration
	AvgRtt                time.Duration
	StdDevRtt             time.Duration
}

func pingUrlWithTimeOutAndCount(urlS string, count int, maxTimeOut time.Duration) error {

	parsedURL, err := url.Parse(urlS)
	if err != nil {
		return err
	}

	pinger, err := ping.NewPinger(parsedURL.Hostname())
	if err != nil {
		return err
	}
	pinger.Timeout = maxTimeOut // на все пинги, не зависимо от количества
	pinger.Count = count
	err = pinger.Run()
	if err != nil {
		return err
	}
	stats := pinger.Statistics()

	/*	result := &PingResult{
		IPAddr:                stats.IPAddr.String(),
		PacketsRecv:           stats.PacketsRecv,
		PacketsSent:           stats.PacketsSent,
		PacketsRecvDuplicates: stats.PacketsRecvDuplicates,
		PacketLoss:            stats.PacketLoss,
		Addr:                  stats.Addr,
		Rtts:                  stats.Rtts,
		MinRtt:                stats.MinRtt,
		MaxRtt:                stats.MaxRtt,
		AvgRtt:                stats.AvgRtt,
		StdDevRtt:             stats.StdDevRtt,
	}*/

	fmt.Println("Ping Result:")
	fmt.Printf("IP Address: %s\n", stats.IPAddr)
	fmt.Printf("Packets Received: %d\n", stats.PacketsRecv)
	fmt.Printf("Packets Sent: %d\n", stats.PacketsSent)
	fmt.Printf("Packet Loss: %.2f%%\n", stats.PacketLoss)
	fmt.Printf("Address: %s\n", stats.Addr)
	fmt.Printf("Minimum RTT: %s\n", stats.MinRtt)
	fmt.Printf("Maximum RTT: %s\n", stats.MaxRtt)
	fmt.Printf("Average RTT: %s\n", stats.AvgRtt)
	fmt.Printf("Standard Deviation RTT: %s\n", stats.StdDevRtt)

	return nil
}
