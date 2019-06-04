// Copyright 2019 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package delegate

import "github.com/cockroachdb/cockroach/pkg/sql/sem/tree"

func (d *delegator) delegateShowDatabases(stmt *tree.ShowDatabases) (tree.Statement, error) {
	query := `SELECT
  DISTINCT
  catalog_name AS database_name`

	if stmt.WithComment {
		query += `,
  shobj_description(oid, 'pg_database') AS comment`
	}

	query += `
FROM
  "".information_schema.schemata`

	if stmt.WithComment {
		query += `
  JOIN pg_database ON
    schemata.catalog_name = pg_database.datname`
	}

	query += `
  ORDER BY 1`

	return parse(query)
}
