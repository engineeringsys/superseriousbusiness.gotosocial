// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package migrations

import (
	"context"
	"strings"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect"
)

func init() {
	up := func(ctx context.Context, db *bun.DB) error {
		if db.Dialect().Name() == dialect.SQLite { // sqlite does not have an array type
			_, err := db.ExecContext(ctx, "ALTER TABLE ? ADD COLUMN ? VARCHAR", bun.Ident("reports"), bun.Ident("rules"))
			if err != nil && !(strings.Contains(err.Error(), "already exists") || strings.Contains(err.Error(), "duplicate column name") || strings.Contains(err.Error(), "SQLSTATE 42701")) {
				return err
			}
		} else {
			_, err := db.ExecContext(ctx, "ALTER TABLE ? ADD COLUMN ? VARCHAR[]", bun.Ident("reports"), bun.Ident("rules"))
			if err != nil && !(strings.Contains(err.Error(), "already exists") || strings.Contains(err.Error(), "duplicate column name") || strings.Contains(err.Error(), "SQLSTATE 42701")) {
				return err
			}
		}
		return nil
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			return nil
		})
	}

	if err := Migrations.Register(up, down); err != nil {
		panic(err)
	}
}
