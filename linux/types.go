//go:build ignore
// +build ignore

// Input to cgo -godefs. See README.md

package libevdev

/*
#include <linux/input.h>
*/
import "C"

type InputEvent C.struct_input_event

type InputId C.struct_input_id

type InputAbsinfo C.struct_input_absinfo

type InputKeymapEntry C.struct_input_keymap_entry

type InputMask C.struct_input_mask

const (
	sizeofInputAbsinfo     = C.sizeof_struct_input_absinfo
	sizeofInputId          = C.sizeof_struct_input_id
	sizeofInputKeymapEntry = C.sizeof_struct_input_keymap_entry
)

const (
	EV_VERSION = C.EV_VERSION
)

const (
	EVIOCGVERSION    = C.EVIOCGVERSION    // get driver version
	EVIOCGID         = C.EVIOCGID         // get device ID
	EVIOCGREP        = C.EVIOCGREP        // get repeat settings
	EVIOCSREP        = C.EVIOCSREP        // set repeat settings
	EVIOCGKEYCODE    = C.EVIOCGKEYCODE    // get keycode
	EVIOCGKEYCODE_V2 = C.EVIOCGKEYCODE_V2 // get keycode
	EVIOCSKEYCODE    = C.EVIOCSKEYCODE    // set keycode
	EVIOCSKEYCODE_V2 = C.EVIOCSKEYCODE_V2 // set keycode
	EVIOCSFF         = C.EVIOCSFF         // send a force effect to a force feedback device
	EVIOCRMFF        = C.EVIOCRMFF        // erase a force effect
	EVIOCGEFFECTS    = C.EVIOCGEFFECTS    // report number of effects playable at the same time
	EVIOCGRAB        = C.EVIOCGRAB        // grab/release device
	EVIOCREVOKE      = C.EVIOCREVOKE      // revoke device access
	EVIOCGMASK       = C.EVIOCGMASK       // get event-masks
	EVIOCSMASK       = C.EVIOCSMASK       // set event-masks
	EVIOCSCLOCKID    = C.EVIOCSCLOCKID    // set clockid to be used for timestamps
)

// IDs
const (
	ID_BUS          = C.ID_BUS
	ID_VENDOR       = C.ID_VENDOR
	ID_PRODUCT      = C.ID_PRODUCT
	ID_VERSION      = C.ID_VERSION
	BUS_PCI         = C.BUS_PCI
	BUS_ISAPNP      = C.BUS_ISAPNP
	BUS_USB         = C.BUS_USB
	BUS_HIL         = C.BUS_HIL
	BUS_BLUETOOTH   = C.BUS_BLUETOOTH
	BUS_VIRTUAL     = C.BUS_VIRTUAL
	BUS_ISA         = C.BUS_ISA
	BUS_I8042       = C.BUS_I8042
	BUS_XTKBD       = C.BUS_XTKBD
	BUS_RS232       = C.BUS_RS232
	BUS_GAMEPORT    = C.BUS_GAMEPORT
	BUS_PARPORT     = C.BUS_PARPORT
	BUS_AMIGA       = C.BUS_AMIGA
	BUS_ADB         = C.BUS_ADB
	BUS_I2C         = C.BUS_I2C
	BUS_HOST        = C.BUS_HOST
	BUS_GSC         = C.BUS_GSC
	BUS_ATARI       = C.BUS_ATARI
	BUS_SPI         = C.BUS_SPI
	BUS_RMI         = C.BUS_RMI
	BUS_CEC         = C.BUS_CEC
	BUS_INTEL_ISHTP = C.BUS_INTEL_ISHTP
)

// MT_TOOL types
const (
	MT_TOOL_FINGER = C.MT_TOOL_FINGER
	MT_TOOL_PEN    = C.MT_TOOL_PEN
	MT_TOOL_PALM   = C.MT_TOOL_PALM
	MT_TOOL_DIAL   = C.MT_TOOL_DIAL
	MT_TOOL_MAX    = C.MT_TOOL_MAX
)

const (
	FF_STATUS_STOPPED = C.FF_STATUS_STOPPED
	FF_STATUS_PLAYING = C.FF_STATUS_PLAYING
	FF_STATUS_MAX     = C.FF_STATUS_MAX
)

type FFRelay C.struct_ff_replay

type FFTrigger C.struct_ff_trigger

type FFEnvelop C.struct_ff_envelope

type FFConstantEffect C.struct_ff_constant_effect

type FFRampEffect C.struct_ff_ramp_effect

type FFConditionEffect C.struct_ff_condition_effect

type FFPeriodicEffect C.struct_ff_periodic_effect

type FFRumbleEffect C.struct_ff_rumble_effect

type FFEffect C.struct_ff_effect

// Force feedback effect types
const (
	FF_RUMBLE       = C.FF_RUMBLE
	FF_PERIODIC     = C.FF_PERIODIC
	FF_CONSTANT     = C.FF_CONSTANT
	FF_SPRING       = C.FF_SPRING
	FF_FRICTION     = C.FF_FRICTION
	FF_DAMPER       = C.FF_DAMPER
	FF_INERTIA      = C.FF_INERTIA
	FF_RAMP         = C.FF_RAMP
	FF_EFFECT_MIN   = C.FF_EFFECT_MIN
	FF_EFFECT_MAX   = C.FF_EFFECT_MAX
	FF_SQUARE       = C.FF_SQUARE
	FF_TRIANGLE     = C.FF_TRIANGLE
	FF_SINE         = C.FF_SINE
	FF_SAW_UP       = C.FF_SAW_UP
	FF_SAW_DOWN     = C.FF_SAW_DOWN
	FF_CUSTOM       = C.FF_CUSTOM
	FF_WAVEFORM_MIN = C.FF_WAVEFORM_MIN
	FF_WAVEFORM_MAX = C.FF_WAVEFORM_MAX
	FF_GAIN         = C.FF_GAIN
	FF_AUTOCENTER   = C.FF_AUTOCENTER
	FF_MAX_EFFECTS  = C.FF_MAX_EFFECTS
	FF_MAX          = C.FF_MAX
)
