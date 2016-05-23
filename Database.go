/*
TODO:
*/

// TODO: Implement SQLite Database

package bf

type Database struct {
	content []Card
	MaxSize int
}

func (db Database) GetFirstN(n int) []Card {
	return db.GetRange(0, n)
}

func (db Database) GetRange(start int, end int) []Card {
	return db.content[start:end]
}

func (db Database) putCard(c Card) bool {
	if len(db.content) == db.MaxSize {
		return false
	}

	if len(db.content) > db.MaxSize {
		db.content = db.content[0:db.MaxSize]
		return false
	}

	return true
}
