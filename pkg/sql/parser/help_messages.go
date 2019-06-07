// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1078
	`ALTER`: {
		//line sql.y: 1079
		Category: hGroup,
		//line sql.y: 1080
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1094
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1095
		Category: hDDL,
		//line sql.y: 1096
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP STORED
  ALTER TABLE ... ALTER [COLUMN] <colname> [SET DATA] TYPE <type> [COLLATE <collation>]
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER TABLE ... UNSPLIT AT <selectclause>
  ALTER TABLE ... UNSPLIT ALL
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER TABLE ... INJECT STATISTICS ...  (experimental)
  ALTER TABLE ... PARTITION BY RANGE ( <name...> ) ( <rangespec> )
  ALTER TABLE ... PARTITION BY LIST ( <name...> ) ( <listspec> )
  ALTER TABLE ... PARTITION BY NOTHING
  ALTER TABLE ... CONFIGURE ZONE <zoneconfig>
  ALTER PARTITION ... OF TABLE ... CONFIGURE ZONE <zoneconfig>

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1134
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1149
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1150
		Category: hDDL,
		//line sql.y: 1151
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1153
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1160
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1161
		Category: hDDL,
		//line sql.y: 1162
		Text: `
ALTER SEQUENCE [IF EXISTS] <name>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]
ALTER SEQUENCE [IF EXISTS] <name> RENAME TO <newname>
`,
	},
	//line sql.y: 1185
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1186
		Category: hPriv,
		//line sql.y: 1187
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1189
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1194
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1195
		Category: hDDL,
		//line sql.y: 1196
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1198
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1206
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1207
		Category: hDDL,
		//line sql.y: 1208
		Text: `
ALTER RANGE <zonename> <command>

Commands:
  ALTER RANGE ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1220
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1225
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1226
		Category: hDDL,
		//line sql.y: 1227
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER INDEX ... UNSPLIT AT <selectclause>
  ALTER INDEX ... UNSPLIT ALL
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER PARTITION ... OF INDEX ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1244
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1709
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1710
		Category: hCCL,
		//line sql.y: 1711
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1728
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1736
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1737
		Category: hCCL,
		//line sql.y: 1738
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1754
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1772
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1773
		Category: hCCL,
		//line sql.y: 1774
		Text: `
-- Import both schema and table data:
IMPORT [ TABLE <tablename> FROM ]
       <format> <datafile>
       [ WITH <option> [= <value>] [, ...] ]

-- Import using specific schema, use only table data from external file:
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   MYSQLOUTFILE
   MYSQLDUMP
   PGCOPY
   PGDUMP

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   delimiter = '...'      [CSV, PGCOPY-specific]
   nullif = '...'         [CSV, PGCOPY-specific]
   comment = '...'        [CSV-specific]

`,
		//line sql.y: 1802
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1841
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1842
		Category: hCCL,
		//line sql.y: 1843
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1852
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 1945
	`CANCEL`: {
		//line sql.y: 1946
		Category: hGroup,
		//line sql.y: 1947
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 1954
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 1955
		Category: hMisc,
		//line sql.y: 1956
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 1959
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 1977
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 1978
		Category: hMisc,
		//line sql.y: 1979
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 1982
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 2013
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 2014
		Category: hMisc,
		//line sql.y: 2015
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 2018
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2090
	`CREATE`: {
		//line sql.y: 2091
		Category: hGroup,
		//line sql.y: 2092
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2173
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 2174
		Category: hMisc,
		//line sql.y: 2175
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2318
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2319
		Category: hDML,
		//line sql.y: 2320
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2324
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2339
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2340
		Category: hCfg,
		//line sql.y: 2341
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2353
	`DROP`: {
		//line sql.y: 2354
		Category: hGroup,
		//line sql.y: 2355
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2372
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2373
		Category: hDDL,
		//line sql.y: 2374
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2375
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2387
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2388
		Category: hDDL,
		//line sql.y: 2389
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2390
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2402
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2403
		Category: hDDL,
		//line sql.y: 2404
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2405
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2417
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2418
		Category: hDDL,
		//line sql.y: 2419
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2420
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2440
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2441
		Category: hDDL,
		//line sql.y: 2442
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2443
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2463
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2464
		Category: hPriv,
		//line sql.y: 2465
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2466
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2478
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2479
		Category: hPriv,
		//line sql.y: 2480
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2481
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2505
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2506
		Category: hMisc,
		//line sql.y: 2507
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>
EXPLAIN ANALYZE [(DISTSQL)] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN

Plan options:
    TYPES, VERBOSE, OPT

`,
		//line sql.y: 2520
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2594
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2595
		Category: hMisc,
		//line sql.y: 2596
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2597
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2628
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2629
		Category: hMisc,
		//line sql.y: 2630
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2631
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2661
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2662
		Category: hMisc,
		//line sql.y: 2663
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2664
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2684
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2685
		Category: hPriv,
		//line sql.y: 2686
		Text: `
Grant privileges:
  GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>
Grant role membership (CCL only):
  GRANT <roles...> TO <grantees...> [WITH ADMIN OPTION]

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2699
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2715
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2716
		Category: hPriv,
		//line sql.y: 2717
		Text: `
Revoke privileges:
  REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>
Revoke role membership (CCL only):
  REVOKE [ADMIN OPTION FOR] <roles...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2730
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2784
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2785
		Category: hCfg,
		//line sql.y: 2786
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2787
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2799
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2800
		Category: hCfg,
		//line sql.y: 2801
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2802
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2811
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2812
		Category: hCfg,
		//line sql.y: 2813
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2816
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2837
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2838
		Category: hExperimental,
		//line sql.y: 2839
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2847
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2853
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2854
		Category: hExperimental,
		//line sql.y: 2855
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2863
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2871
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2872
		Category: hExperimental,
		//line sql.y: 2873
		Text: `
SCRUB TABLE <tablename>
            [AS OF SYSTEM TIME <expr>]
            [WITH OPTIONS <option> [, ...]]

Options:
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT (<constraint>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS PHYSICAL
`,
		//line sql.y: 2884
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2939
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2940
		Category: hCfg,
		//line sql.y: 2941
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2942
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2963
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2964
		Category: hCfg,
		//line sql.y: 2965
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 2971
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2988
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2989
		Category: hTxn,
		//line sql.y: 2990
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2997
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3180
	`SHOW`: {
		//line sql.y: 3181
		Category: hGroup,
		//line sql.y: 3182
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
JOBS, SHOW QUERIES, SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW
SESSION, SHOW SESSIONS, SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES,
SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3216
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3217
		Category: hCfg,
		//line sql.y: 3218
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3219
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3240
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3241
		Category: hExperimental,
		//line sql.y: 3242
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3249
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3262
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3263
		Category: hExperimental,
		//line sql.y: 3264
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3268
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3281
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3282
		Category: hCCL,
		//line sql.y: 3283
		Text: `SHOW BACKUP [FILES|RANGES] <location>
`,
		//line sql.y: 3284
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3311
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3312
		Category: hCfg,
		//line sql.y: 3313
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 3316
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3333
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3334
		Category: hDDL,
		//line sql.y: 3335
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3336
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3344
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3345
		Category: hDDL,
		//line sql.y: 3346
		Text: `SHOW DATABASES
`,
		//line sql.y: 3347
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3355
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3356
		Category: hPriv,
		//line sql.y: 3357
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3363
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3376
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3377
		Category: hDDL,
		//line sql.y: 3378
		Text: `SHOW INDEXES FROM { <tablename> | DATABASE <database_name> }
`,
		//line sql.y: 3379
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3409
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3410
		Category: hDDL,
		//line sql.y: 3411
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3412
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3425
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3426
		Category: hMisc,
		//line sql.y: 3427
		Text: `SHOW [ALL] [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3428
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3449
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3450
		Category: hMisc,
		//line sql.y: 3451
		Text: `SHOW [AUTOMATIC] JOBS
`,
		//line sql.y: 3452
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3464
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3465
		Category: hMisc,
		//line sql.y: 3466
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3468
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3491
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3492
		Category: hMisc,
		//line sql.y: 3493
		Text: `SHOW [ALL] [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3494
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3507
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3508
		Category: hDDL,
		//line sql.y: 3509
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ] [WITH COMMENT]
`,
		//line sql.y: 3510
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3542
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3543
		Category: hDDL,
		//line sql.y: 3544
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3556
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 3557
		Category: hDDL,
		//line sql.y: 3558
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 3570
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3571
		Category: hMisc,
		//line sql.y: 3572
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3581
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3582
		Category: hCfg,
		//line sql.y: 3583
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3584
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3603
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3604
		Category: hDDL,
		//line sql.y: 3605
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3606
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3624
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3625
		Category: hPriv,
		//line sql.y: 3626
		Text: `SHOW USERS
`,
		//line sql.y: 3627
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3635
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3636
		Category: hPriv,
		//line sql.y: 3637
		Text: `SHOW ROLES
`,
		//line sql.y: 3638
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3694
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3695
		Category: hMisc,
		//line sql.y: 3696
		Text: `
SHOW EXPERIMENTAL_RANGES FROM TABLE <tablename>
SHOW EXPERIMENTAL_RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 3933
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 3934
		Category: hMisc,
		//line sql.y: 3935
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 3938
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3955
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3956
		Category: hDDL,
		//line sql.y: 3957
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3984
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4575
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4576
		Category: hDDL,
		//line sql.y: 4577
		Text: `
CREATE SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 4587
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 4634
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 4635
		Category: hDML,
		//line sql.y: 4636
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 4637
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 4645
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 4646
		Category: hPriv,
		//line sql.y: 4647
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 4648
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 4670
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 4671
		Category: hPriv,
		//line sql.y: 4672
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4673
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 4691
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 4692
		Category: hDDL,
		//line sql.y: 4693
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 4694
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 4728
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 4729
		Category: hDDL,
		//line sql.y: 4730
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4738
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 4956
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 4957
		Category: hTxn,
		//line sql.y: 4958
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 4959
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4967
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 4968
		Category: hMisc,
		//line sql.y: 4969
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 4972
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 4989
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 4990
		Category: hTxn,
		//line sql.y: 4991
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 4992
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5007
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5008
		Category: hTxn,
		//line sql.y: 5009
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5017
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5030
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5031
		Category: hTxn,
		//line sql.y: 5032
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5035
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5059
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 5060
		Category: hTxn,
		//line sql.y: 5061
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 5062
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5180
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5181
		Category: hDDL,
		//line sql.y: 5182
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5183
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5252
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5253
		Category: hDML,
		//line sql.y: 5254
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5259
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5278
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5279
		Category: hDML,
		//line sql.y: 5280
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5284
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5391
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5392
		Category: hDML,
		//line sql.y: 5393
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5400
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 5574
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 5575
		Category: hDML,
		//line sql.y: 5576
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 5587
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 5588
		Category: hDML,
		//line sql.y: 5589
		Text: `
SELECT [DISTINCT [ ON ( <expr> [ , ... ] ) ] ]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 5601
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 5676
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 5677
		Category: hDML,
		//line sql.y: 5678
		Text: `TABLE <tablename>
`,
		//line sql.y: 5679
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5959
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 5960
		Category: hDML,
		//line sql.y: 5961
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 5962
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6071
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6072
		Category: hDML,
		//line sql.y: 6073
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexflags> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> [ <jointype> ] JOIN <source> ON <expr>
  <source> [ <jointype> ] JOIN <source> USING ( <colnames...> )
  <source> NATURAL [ <jointype> ] JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'
  '{' IGNORE_FOREIGN_KEYS [, ...] '}'

Join types:
  { INNER | { LEFT | RIGHT | FULL } [OUTER] } [ { HASH | MERGE | LOOKUP } ]

`,
		//line sql.y: 6095
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
