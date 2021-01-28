package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/rizalhamdana/refactory_test/soal_2/model"
	"github.com/rizalhamdana/refactory_test/soal_2/repo"
)

type Soal2UsecaseImpl struct {
	repo repo.Soal2Repository
}

func NewUsecase(repo *repo.Soal2RepoImpl) *Soal2UsecaseImpl {
	return &Soal2UsecaseImpl{repo: repo}
}

func (u *Soal2UsecaseImpl) FindUsersWhoNoPhoneNumbers() {
	allUsers, err := u.repo.GetAllUsersFromJson()

	if err != nil {
		fmt.Println(err.Error())
	}
	count := 0
	fmt.Println("USER(S) WHO DO NOT HAVE PHONE NUMBER:")
	printFormat := "id: %d - username: %s - fullname: %s \n"
	for _, user := range *allUsers {
		if len(user.Profile.Phones) <= 0 {
			fmt.Printf(printFormat, user.Id, user.Username, user.Profile.FullName)
			count++
		}
	}
	fmt.Printf("\nCount: %d\n", count)
}

func (u *Soal2UsecaseImpl) FindUsersWhoHaveArticles() {
	allUsers, err := u.repo.GetAllUsersFromJson()

	if err != nil {
		fmt.Println(err.Error())
	}
	count := 0
	fmt.Println("USER(S) WHO HAVE ARTICLES:")
	printFormat := "id: %d - username: %s - fullname: %s \n"

	for _, user := range *allUsers {
		if len(user.Articles) > 0 {
			fmt.Printf(printFormat, user.Id, user.Username, user.Profile.FullName)
			count++
		}
	}
	fmt.Printf("\nCount: %d\n", count)

}

func (u *Soal2UsecaseImpl) FindUsersWhoHaveAnnisInTheirName() {
	allUsers, err := u.repo.GetAllUsersFromJson()

	if err != nil {
		fmt.Println(err.Error())
	}
	count := 0
	fmt.Println("USER(S) WHO HAVE \"annis\" in their name")
	printFormat := "id: %d - username: %s - fullname: %s \n"
	for _, user := range *allUsers {
		fullNameLowerCase := strings.ToLower(user.Profile.FullName)
		if strings.Contains(fullNameLowerCase, "annis") {
			fmt.Printf(printFormat, user.Id, user.Username, user.Profile.FullName)
			count++
		}
	}
	fmt.Printf("\nCount: %d\n", count)

}

func (u *Soal2UsecaseImpl) FindUsersWhoBornOn1996() {
	allUsers, err := u.repo.GetAllUsersFromJson()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("USER(S) WHO WERE BORN IN 1986")
	printFormat := "id: %d - username: %s - fullname: %s - birthday: %s \n"
	count := 0
	for _, user := range *allUsers {
		birthDayTime, err := time.Parse("2006-01-02", user.Profile.BirthDay)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if birthDayTime.Year() == 1986 {
			fmt.Printf(printFormat+"", user.Id, user.Username, user.Profile.FullName, user.Profile.BirthDay)
			count++
		}
	}
	fmt.Printf("\nCount: %d\n", count)
}

func (u *Soal2UsecaseImpl) FindUsersWhoHaveArticlesIn2020() {
	allUsers, err := u.repo.GetAllUsersFromJson()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("USER(S) WHO HAVE ARTICLES PUBLISHED IN 2020")
	printFormat := "id: %d - username: %s - fullname: %s \n"
	count := 0
	for _, user := range *allUsers {
		for _, article := range user.Articles {
			publishedTime, err := time.Parse("2006-01-02T15:04:05.000Z", article.PublishedAt+".000Z")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if publishedTime.Year() == 2020 {
				fmt.Printf(printFormat+"", user.Id, user.Username, user.Profile.FullName, user.Profile.BirthDay)
				count++
			}
		}

	}
	fmt.Printf("\nCount: %d\n", count)
}

func (u *Soal2UsecaseImpl) FindArticlesContainTips() {
	allUsers, err := u.repo.GetAllUsersFromJson()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("ARTICLE(S) WHO HAVE \"tips\" in their title")
	printFormat := "id: %d - title: %s - published at: %s \n"

	allArticles := u.GetAllArticlesFromAllUsers(allUsers)
	count := 0
	for _, article := range allArticles {
		titleLower := strings.ToLower(article.Title)
		if strings.Contains(titleLower, "tips") {
			fmt.Printf(printFormat+"", article.Id, article.Title, article.PublishedAt)
			count++
		}
	}
	fmt.Printf("\nCount: %d\n", count)
}

func (u *Soal2UsecaseImpl) FindArticlesPublishedBeforeAugust2019() {
	allUsers, err := u.repo.GetAllUsersFromJson()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("ARTICLE(S) WHICH PUBLISHED BEFORE AUGUST 2019")
	printFormat := "id: %d - title: %s - published at: %s \n"

	allArticles := u.GetAllArticlesFromAllUsers(allUsers)
	count := 0

	for _, article := range allArticles {
		firstAugust2019Time, err := time.Parse("2006-01-02", "2019-08-01")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		publishedTime, err := time.Parse("2006-01-02T15:04:05.000Z", article.PublishedAt+".000Z")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if publishedTime.Before(firstAugust2019Time) {
			fmt.Printf(printFormat+"", article.Id, article.Title, article.PublishedAt)
			count++
		}
	}
	fmt.Printf("\nCount: %d\n", count)

}

func (u *Soal2UsecaseImpl) GetAllArticlesFromAllUsers(users *[]model.User) (articles []model.Article) {

	for _, user := range *users {
		articles = append(articles, user.Articles...)
	}
	return
}
