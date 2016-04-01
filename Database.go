/*
TODO:
*/

// TODO: Implement SQLite Database

package bf

type Database struct {
	content []Card
	MaxSize int
}

func GetFirstN(n int, db Database) []Card {
	return GetRange(0, n, db)
}

func GetRange(start int, end int, db Database) []Card {
	return db.content[start:end]
}

func putCard(c Card, db Database) bool {
	if len(db.content) == db.MaxSize {
		return false
	}

	if len(db.content) > db.MaxSize {
		db.content = db.content[0:db.MaxSize]
		return false
	}

	return true
}
