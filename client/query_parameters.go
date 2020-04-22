package client

import (
	"net/url"
	"strconv"
	"time"
)

// QueryParameter is a function that sets a URl value
type QueryParameter func(u *url.Values)

func query(params []QueryParameter) (query string) {
	if len(params) > 0 {
		vals := url.Values{}
		for _, arg := range params {
			arg(&vals)
		}

		query = "?" + query
	}

	return query
}

// WithAll Can only be used by admins or managers to fetch all entities
func WithAll(val bool) QueryParameter {
	return func(u *url.Values) {
		u.Set("all", strconv.FormatBool(val))
	}
}

// WithID to fetch one or more devices. Multiple params can be passed
func WithID(val ...int) QueryParameter {
	return func(u *url.Values) {
		for _, v := range val {
			u.Add("id", strconv.FormatInt(int64(v), 10))
		}
	}
}

// WithUniqueID to fetch one or more devices. Multiple params can be passed
func WithUniqueID(val ...string) QueryParameter {
	return func(u *url.Values) {
		for _, v := range val {
			u.Add("uniqueId", v)
		}
	}
}

// WithUserID Standard users can use this only with their own userId
func WithUserID(val int) QueryParameter {
	return func(u *url.Values) {
		u.Set("userId", strconv.FormatInt(int64(val), 10))
	}
}

// WithDeviceID Standard users can use this only with _deviceId_s, they have access to
func WithDeviceID(val ...int) QueryParameter {
	return func(u *url.Values) {
		for _, v := range val {
			u.Add("deviceId", strconv.FormatInt(int64(v), 10))
		}
	}
}

// WithGroupID Standard users can use this only with _groupId_s, they have access to
func WithGroupID(val ...int) QueryParameter {
	return func(u *url.Values) {
		for _, v := range val {
			u.Add("groupId", strconv.FormatInt(int64(v), 10))
		}
	}
}

// WithRefresh ...
func WithRefresh(val bool) QueryParameter {
	return func(u *url.Values) {
		u.Set("refresh", strconv.FormatBool(val))
	}
}

// WithTextChannel ...
func WithTextChannel(val bool) QueryParameter {
	return func(u *url.Values) {
		u.Set("refresh", strconv.FormatBool(val))
	}
}

// WithFrom ...
func WithFrom(val time.Time) QueryParameter {
	return func(u *url.Values) {
		u.Set("from", val.Format(time.RFC3339))
	}
}

// WithTo ..
func WithTo(val time.Time) QueryParameter {
	return func(u *url.Values) {
		u.Set("to", val.Format(time.RFC3339))
	}
}
