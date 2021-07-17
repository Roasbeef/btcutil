package bech32

// Bech32Const is a type that represents the currently defined bech32 checksum
// constants.
type Bech32Const int

const (
	// Version0Const is the original constant used in the checksum
	// verification for bech32.
	Version0Const Bech32Const = 1

	// VersionMConst is the new constant used for bech32m checksum
	// verification.
	VersionMConst Bech32Const = 0x2bc830a3
)

// Bech32Version defines the current set of bech32 versions.
type Bech32Version uint8

const (
	// Version0 defines the original bech version.
	Version0 Bech32Version = iota

	// VersionM is the new bech32 version defined in BIP-350, also known as
	// bech32m.
	VersionM

	// VersionUnknown denotes an unknown bech version.
	VersionUnknown
)

// Bech32VersionToConsts maps bech32 versions to the checksum constant to be
// used when encoding, and asserting a particular version when decoding.
var Bech32VersionToConsts = map[Bech32Version]Bech32Const{
	Version0: Version0Const,
	VersionM: VersionMConst,
}

// Bech32ConstsToVersion maps a bech32 constant to the version it's associated
// with.
var Bech32ConstsToVersion = map[Bech32Const]Bech32Version{
	Version0Const: Version0,
	VersionMConst: VersionM,
}
