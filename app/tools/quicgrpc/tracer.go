package quicgrpc

import (
	"context"
	"net"
	"time"

	"github.com/quic-go/quic-go/logging"
)

var tracer = &quicTracer{}

type quicTracer struct{}

var _ logging.Tracer = (*quicTracer)(nil)

func (q *quicTracer) TracerForConnection(context.Context, logging.Perspective, logging.ConnectionID) logging.ConnectionTracer {
	return &quicConnTracer{}
}

func (q *quicTracer) SentPacket(net.Addr, *logging.Header, logging.ByteCount, []logging.Frame) {
}

func (q *quicTracer) SentVersionNegotiationPacket(net.Addr, logging.ArbitraryLenConnectionID, logging.ArbitraryLenConnectionID, []logging.VersionNumber) {
}

func (q *quicTracer) DroppedPacket(net.Addr, logging.PacketType, logging.ByteCount, logging.PacketDropReason) {
}

type quicConnTracer struct {
	addr net.Addr
}

func (q *quicConnTracer) StartedConnection(_ net.Addr, remote net.Addr, _ logging.ConnectionID, _ logging.ConnectionID) {
	// store addr for future delete
	q.addr = remote
}
func (q *quicConnTracer) NegotiatedVersion(logging.VersionNumber, []logging.VersionNumber, []logging.VersionNumber) {
}
func (q *quicConnTracer) ClosedConnection(error)                                   {}
func (q *quicConnTracer) SentTransportParameters(*logging.TransportParameters)     {}
func (q *quicConnTracer) ReceivedTransportParameters(*logging.TransportParameters) {}
func (q *quicConnTracer) RestoredTransportParameters(*logging.TransportParameters) {}
func (q *quicConnTracer) SentLongHeaderPacket(*logging.ExtendedHeader, logging.ByteCount, *logging.AckFrame, []logging.Frame) {
}
func (q *quicConnTracer) SentShortHeaderPacket(*logging.ShortHeader, logging.ByteCount, *logging.AckFrame, []logging.Frame) {
}
func (q *quicConnTracer) ReceivedVersionNegotiationPacket(logging.ArbitraryLenConnectionID, logging.ArbitraryLenConnectionID, []logging.VersionNumber) {
}
func (q *quicConnTracer) ReceivedRetry(*logging.Header) {}
func (q *quicConnTracer) ReceivedLongHeaderPacket(*logging.ExtendedHeader, logging.ByteCount, []logging.Frame) {
}
func (q *quicConnTracer) ReceivedShortHeaderPacket(*logging.ShortHeader, logging.ByteCount, []logging.Frame) {
}
func (q *quicConnTracer) BufferedPacket(logging.PacketType, logging.ByteCount) {}
func (q *quicConnTracer) DroppedPacket(logging.PacketType, logging.ByteCount, logging.PacketDropReason) {
}
func (q *quicConnTracer) UpdatedMetrics(*logging.RTTStats, logging.ByteCount, logging.ByteCount, int) {
}
func (q *quicConnTracer) AcknowledgedPacket(logging.EncryptionLevel, logging.PacketNumber) {}
func (q *quicConnTracer) LostPacket(logging.EncryptionLevel, logging.PacketNumber, logging.PacketLossReason) {
}
func (q *quicConnTracer) UpdatedCongestionState(logging.CongestionState)                     {}
func (q *quicConnTracer) UpdatedPTOCount(uint32)                                             {}
func (q *quicConnTracer) UpdatedKeyFromTLS(logging.EncryptionLevel, logging.Perspective)     {}
func (q *quicConnTracer) UpdatedKey(logging.KeyPhase, bool)                                  {}
func (q *quicConnTracer) DroppedEncryptionLevel(logging.EncryptionLevel)                     {}
func (q *quicConnTracer) DroppedKey(logging.KeyPhase)                                        {}
func (q *quicConnTracer) SetLossTimer(logging.TimerType, logging.EncryptionLevel, time.Time) {}
func (q *quicConnTracer) LossTimerExpired(logging.TimerType, logging.EncryptionLevel)        {}
func (q *quicConnTracer) LossTimerCanceled()                                                 {}
func (q *quicConnTracer) Close() {
	connCache.Delete(q.addr.String())
}
func (q *quicConnTracer) Debug(string, string) {}
