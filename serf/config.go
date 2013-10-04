package serf

import (
	"github.com/hashicorp/memberlist"
	"time"
)

type Config struct {
	Hostname string // Node name (FQDN)
	Role     string // Role in the gossip pool

	MaxCoalesceTime  time.Duration // Maximum period of event coalescing for updates
	MinQuiescentTime time.Duration // Minimum period of quiescence for updates. This has lower precedence then MaxCoalesceTime

	LeaveTimeout time.Duration // Timeout for leaving

	PartitionCount    int           // If PartitionCount nodes fail in PartitionInvernal, it is considered a partition
	PartitionInterval time.Duration // ParitionInterval must be < MaxCoalesceTime

	ReconnectInterval time.Duration // How often do we attempt to reconnect to failed nodes
	ReconnectTimeout  time.Duration // How long do we keep retrying to connect to a failed node before giving up

	GossipBindAddr   string        // Binding address
	GossipPort       int           // TCP and UDP ports for gossip
	GossipTCPTimeout time.Duration // TCP timeout
	IndirectChecks   int           // Number of indirect checks to use
	RetransmitMult   int           // Retransmits = RetransmitMult * log(N+1)
	SuspicionMult    int           // Suspicion time = SuspcicionMult * log(N+1) * Interval
	PushPullInterval time.Duration // How often we do a Push/Pull update
	RTT              time.Duration // 99% precentile of round-trip-time
	ProbeInterval    time.Duration // Failure probing interval length
	GossipNodes      int           // Number of nodes to gossip to per GossipInterval
	GossipInterval   time.Duration // Gossip interval for non-piggyback messages (only if GossipNodes > 0)

	Delegate EventDelegate // Notified on member events
}

// memberlistConfig constructs the memberlist configuration from our configuration
func memberlistConfig(conf *Config) *memberlist.Config {
	mc := &memberlist.Config{}
	mc.Name = conf.Hostname
	mc.BindAddr = conf.GossipBindAddr
	mc.UDPPort = conf.GossipPort
	mc.TCPPort = conf.GossipPort
	mc.TCPTimeout = conf.GossipTCPTimeout
	mc.IndirectChecks = conf.IndirectChecks
	mc.RetransmitMult = conf.RetransmitMult
	mc.SuspicionMult = conf.SuspicionMult
	mc.PushPullInterval = conf.PushPullInterval
	mc.RTT = conf.RTT
	mc.ProbeInterval = conf.ProbeInterval
	mc.GossipNodes = conf.GossipNodes
	mc.GossipInterval = conf.GossipInterval
	return mc
}

// DefaultConfig is used to return a default set of sane configurations
func DefaultConfig() *Config {
	c := &Config{}

	// Copy the memberlist configs
	defaultMb := memberlist.DefaultConfig()
	c.Hostname = defaultMb.Name
	c.GossipBindAddr = defaultMb.BindAddr
	c.GossipPort = defaultMb.UDPPort
	c.GossipTCPTimeout = defaultMb.TCPTimeout
	c.IndirectChecks = defaultMb.IndirectChecks
	c.RetransmitMult = defaultMb.RetransmitMult
	c.SuspicionMult = defaultMb.SuspicionMult
	c.PushPullInterval = defaultMb.PushPullInterval
	c.RTT = defaultMb.RTT
	c.ProbeInterval = defaultMb.ProbeInterval
	c.GossipNodes = defaultMb.GossipNodes
	c.GossipInterval = defaultMb.GossipInterval

	// Set our own defaults
	c.PartitionCount = 2
	c.PartitionInterval = 30 * time.Second
	c.MinQuiescentTime = 5 * time.Second
	c.MaxCoalesceTime = 60 * time.Second
	c.ReconnectInterval = 30 * time.Second
	c.ReconnectTimeout = 24 * time.Hour

	return c
}
