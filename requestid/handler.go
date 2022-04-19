/*
 * Copyright (c) The go-kit Authors
 */

package requestid

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// NewRequestIDHttpHandler sets unique request id.
// If header `x-request-id` is already present in the request, that is considered the
// request id. Otherwise, generates a new unique ID.
func NewRequestIDHttpHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// find the request ID
		requestID := r.Header.Get(XRequestIDMetadataKey)
		// if not found generate a new one
		if requestID == "" {
			requestID = uuid.NewString()
			r.Header.Set(XRequestIDMetadataKey, requestID)
		}
		// set the request ID in the context
		ctx := context.WithValue(r.Context(), XRequestIDKey{}, requestID)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
