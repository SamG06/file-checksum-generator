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
		value TEXT,
		PRIMARY KEY ()
	)`

return []string{Link, Checksum}

}