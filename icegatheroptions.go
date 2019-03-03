package webrtc

// ICEGatherOptions provides options relating to the gathering of ICE candidates.
type ICEGatherOptions struct {
	ICEServers []ICEServer

	// IgnoreIPv6 is an optional configuration for disabling support for ipv6
	// server reflexive candidates. Default is false.
	IgnoreIPv6 bool
}
