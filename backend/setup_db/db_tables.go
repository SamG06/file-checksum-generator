package setupDB

func TablesSetup() []string {
	
	Link := `CREATE TABLE IF NOT EXISTS checksum_links(
		id uuid DEFAULT ,
		filename TEXT,
		PRIMARY KEY ()
	)`;

	Checksum := `CREATE TABLE IF NOT EXISTS checksum(
		id SERIAL PRIMARY KEY,
		type TEXT,
		CHECK (type in ('sha512', 'sha256', 'md5', 'sha384'))
		value TEXT,
		PRIMARY KEY ()
	)`

return []string{Link, Checksum}

}