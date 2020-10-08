
/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package contract

import (
	"github.com/mysteriumnetwork/node/tequilapi/validation"
	"net/url"
)

// QueryParseContext is used for binding url key value
type QueryParseContext struct {
	query url.Values
	key  string
	value  string
}

// NewQueryParse create query parse context for later
func NewQueryParse(query url.Values, key string) *QueryParseContext {
	return &QueryParseContext{query, key, query.Get(key)}
}

// BindInt binds int value and errors if query key value is not empty
func (qp *QueryParseContext) BindInt(to *int, errs *validation.FieldErrorMap) {
	if qp.value != "" {
		v, e := parseInt(qp.value)

		if (e == nil) {
			*to = *v
		}

		errs.Set(validation.NewSingleErrorMap(qp.key, e))
	}
}