package usecase

type Soal2Usecase interface {
	FindUsersWhoNoPhoneNumbers()
	FindUsersWhoHaveArticles()
	FindUsersWhoHaveAnnisInTheirName()
	FindUsersWhoHaveArticlesIn2020()
	FindUsersWhoBornOn1996()
	FindArticlesContainTips()
	FindArticlesPublishedBeforeAugust2019()
}
