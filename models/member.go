package models

type MemberData struct {
	Name  string `bson:"name" json:"name"`
	Image string `bson:"image" json:"image"`
}

type MostWatchedMember struct {
	Member MemberData `bson:"member" json:"member"`
	Watch  int        `bson:"watch" json:"watch"`
}
