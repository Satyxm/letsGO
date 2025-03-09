package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type UserStore struct {
	users  map[int]User
	mu     sync.Mutex
	nextID int
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (s *UserStore) Create(name string, age int) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := User{ID: s.nextID, Name: name, Age: age}
	s.users[s.nextID] = user
	s.nextID++

	return user
}

func (s *UserStore) Read(id int) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserStore) Update(id int, name string, age int) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}

	user.Name = name
	user.Age = age
	s.users[id] = user

	return user, nil
}

func (s *UserStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(s.users, id)
	return nil
}

func (s *UserStore) List() []User {
	s.mu.Lock()
	defer s.mu.Unlock()

	var userList []User
	for _, user := range s.users {
		userList = append(userList, user)
	}
	return userList
}

func main() {
	store := NewUserStore()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose an option: 1) Create 2) Read 3) Update 4) Delete 5) List 6) Exit")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Enter age: ")
			ageStr, _ := reader.ReadString('\n')
			age, _ := strconv.Atoi(strings.TrimSpace(ageStr))
			user := store.Create(name, age)
			fmt.Println("Created User:", user)

		case "2":
			fmt.Print("Enter user ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			user, err := store.Read(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("User:", user)
			}

		case "3":
			fmt.Print("Enter user ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			fmt.Print("Enter new name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Enter new age: ")
			ageStr, _ := reader.ReadString('\n')
			age, _ := strconv.Atoi(strings.TrimSpace(ageStr))
			user, err := store.Update(id, name, age)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Updated User:", user)
			}

		case "4":
			fmt.Print("Enter user ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			err := store.Delete(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("User deleted.")
			}

		case "5":
			fmt.Println("User List:", store.List())

		case "6":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
