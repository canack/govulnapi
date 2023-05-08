package database

// CWE-340: Generation of Predictable Numbers or Identifiers
// Autoincrement is used for user id
const sqlSchema = `
CREATE TABLE IF NOT EXISTS "coin" (
	"id"	TEXT NOT NULL,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "user" (
	"id"	INTEGER NOT NULL,
	"email"	TEXT NOT NULL UNIQUE,
	"password"	TEXT NOT NULL,
	"usd_balance"	REAL NOT NULL DEFAULT 10000,
	"usd_starting_balance"	REAL NOT NULL DEFAULT 10000,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "coin_balance" (
	"user_id"	INTEGER,
	"coin_id"	TEXT,
	"address"	TEXT NOT NULL UNIQUE,
	"qty"	REAL NOT NULL,
	PRIMARY KEY("user_id","coin_id"),
	FOREIGN KEY("coin_id") REFERENCES "coin"("id"),
	FOREIGN KEY("user_id") REFERENCES "user"("id")
);
CREATE TABLE IF NOT EXISTS "order" (
	"id"	INTEGER,
	"user_id"	INTEGER NOT NULL,
	"coin_id"	TEXT NOT NULL,
	"price"	REAL NOT NULL,
	"is_buy"	INTEGER NOT NULL,
	"qty"	REAL NOT NULL,
	"date"	TEXT NOT NULL,
	FOREIGN KEY("user_id") REFERENCES "user"("id"),
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("coin_id") REFERENCES "coin"("id")
);
CREATE TABLE IF NOT EXISTS "transaction" (
	"id"	INTEGER,
	"sender_id"	INTEGER NOT NULL,
	"receiver_id"	INTEGER NOT NULL,
	"coin_id"	TEXT NOT NULL,
	"address"	TEXT NOT NULL,
	"qty"	REAL NOT NULL,
	"date"	TEXT NOT NULL,
	"note"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("sender_id") REFERENCES "user"("id"),
	FOREIGN KEY("receiver_id") REFERENCES "user"("id"),
	FOREIGN KEY("address") REFERENCES "coin_balance"("address"),
	FOREIGN KEY("coin_id") REFERENCES "coin"("id")
);
INSERT INTO "coin" ("id") VALUES ('bitcoin'),
 ('litecoin'),
 ('namecoin'),
 ('ripple'),
 ('dogecoin');
`
