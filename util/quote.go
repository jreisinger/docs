package util

import "math/rand"

type quote struct {
	what   string
	author string
}

var quotes = []quote{
	{
		`I'm here because computer science is wonderful, but it isn't
		everything.`, `Donald Knuth`,
	},
	{
		`Raising different kids requires different approaches, just
		like computer problems do.`, `Larry Wall`,
	},
	{
		`The only way to become smart is to be stupid first.`, `Larry
		Wall`,
	},
	{
		`It can get confusing. Hey, I'm confused most of the time.`,
		`Larry Wall`,
	},
	{
		`The best way to avoid burnout is to do something you truly
		enjoy in an environment that supports you.`, `Rob Pike`,
	},
	{
		`Know your place in the world and evaluate yourself fairly,
		not in terms of your naive ideals of your own youth, nor in
		terms of what you erroneously imagine your teacher's ideals
		are.`, `Richard Feynman`,
	},
}

func randQuote() []byte {
	q := quotes[rand.Intn(len(quotes))]
	md := "> " + q.what + " --- " + q.author
	return []byte(md)
}
