package main

import "fmt"

type UserWalletRegistryService interface {
	AddWallet(name string, wallet string)
	DeleteWallet(name string, wallet string)
	WalletsOfUser(name string) []string
	ViewAllUsers()
	UsersCount() int
}

type UserWalletRegistry struct {
	User map[string][]string
}

func (u *UserWalletRegistry) AddWallet(name string, wallet string) {
	if _, exists := u.User[name]; !exists {
		u.User[name] = []string{}
	}
	for _, existsWallet := range u.User[name] {
		if existsWallet == wallet {
			fmt.Println("Wallet already exist!", wallet)
		}
	}

	u.User[name] = append(u.User[name], wallet)
	fmt.Println("Wallet added:", name, wallet)
}

func (u *UserWalletRegistry) DeleteWallet(name string, wallet string) {
	wallets, exist := u.User[name]
	if !exist {
		fmt.Println("User not found", name)
		return
	}

	for i, currenWallet := range wallets {
		if currenWallet == wallet {
			wallets = append(wallets[:i], wallets[i+1:]...)
			u.User[name] = wallets
			fmt.Println("Wallet deleted!", wallet)
			return
		}
	}
	fmt.Println("Wallet not found!", wallet)
}

func (u *UserWalletRegistry) WalletsOfUser(name string) []string {
	return u.User[name]
}

func (u *UserWalletRegistry) ViewAllUsers() {
	for name, wallets := range u.User {
		fmt.Println("User info:", name, wallets)
	}
}

func (u *UserWalletRegistry) UsersCount() int {
	return len(u.User)
}

func PrintInfo(u UserWalletRegistryService, name string, wallet string) {
	u.ViewAllUsers()
	fmt.Println("Users count:", u.UsersCount())

	u.AddWallet(name, wallet)
	fmt.Println("Wallets of user:", name, u.WalletsOfUser(name))

	u.ViewAllUsers()

	u.DeleteWallet(name, wallet)
	fmt.Println("Wallets of user:", name, u.WalletsOfUser(name))

	u.ViewAllUsers()
	fmt.Println("Users count:", u.UsersCount())
}

func main() {
	userRegistry := UserWalletRegistry{
		User: map[string][]string{
			"Vladislav": {"asdjw02x", "09xajsdwj", "asdwsadkw"},
			"Viola":     {"gjwusk03x", "dkgjwux123", "sdjwufks0982"},
			"Afina":     {"sdkjgue982", "fsawfkxjw", "judyajwdk"},
		},
	}
	PrintInfo(&userRegistry, "Vladislav", "lokdsajw817")
}
