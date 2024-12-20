package params

// SequoiaBootnodes are the enode URLs of the P2P bootstrap nodes running
// reliably and availably on the Sequoia network.
// They will be the first point of contact for nodes coming online
// to find peers.
var SequoiaBootnodes = []string{
	"enode://bd0e33eeb371f92b0cabaa487e7c41c80384645c8561c811cb723270ccf829d3a8aa2e27e5a410ca56652a0286f09f291ac5578c6084e48b4e317e4608f2b297@bootsequoia.edevapps.com.br:30303",
}

// Once Sequoia network has DNS discovery set up,
// this value can be configured.
// var SequoiaDNSNetwork = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@example.network"
