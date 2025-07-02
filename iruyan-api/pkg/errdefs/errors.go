package errdefs

import "errors"

var (
	// --- User-related errors ---
	ErrDuplicateIruyanID    = errors.New("iruyan_id is already taken")
	ErrDuplicateEmail       = errors.New("email is already registered")
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrInvalidEmail         = errors.New("invalid email format")
	ErrPasswordTooShort     = errors.New("password must be at least 6 characters long")
	ErrPasswordMissingChars = errors.New("password must contain at least one letter and one number")

	// --- Room-related errors ---
	ErrRoomNameRequired      = errors.New("room name is required")
	ErrRoomNameInvalidFormat = errors.New("room name must contain only lowercase letters and numbers, with no special characters")
	ErrInvalidRoomName       = errors.New("invalid room name")
	ErrRoomAlreadyExists     = errors.New("room already exists")
	ErrRoomNotFound          = errors.New("room not found")
	ErrUserAlreadyInRoom     = errors.New("user is already in the room")
	ErrNotInRoom             = errors.New("user is already in the room")
	ErrAlreadyInRoom         = errors.New("user is not currently in the room")
	ErrInvalidDuration       = errors.New("invalid duration: shorter than time spent")
	ErrSeatNotFound          = errors.New("Seat not found in this room")
	ErrSeatAlreadyTaken      = errors.New("seat is already taken")
	ErrNotSeated             = errors.New("the user is not currently seated")
	ErrSeatMismatch          = errors.New("the seat number does not match the user's current seat")
	ErrDataRetrievalFailed   = errors.New("failed to retrieve seat status")
	ErrNoActiveSession       = errors.New("no active session found")

	// --- General errors ---
	ErrInternalServer = errors.New("internal server error")
)
