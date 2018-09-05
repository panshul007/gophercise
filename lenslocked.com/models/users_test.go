package models

//func testingUserService() (*UserService, error) {
//	const (
//		host   = "localhost"
//		port   = 5432
//		user   = "postgres"
//		dbname = "lenslocked_test"
//	)
//
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
//	us, err := NewUserService(psqlInfo)
//	if err != nil {
//		return nil, err
//	}
//	us.db.LogMode(false)
//	// Clear the users table between tests
//	us.DestructiveReset()
//	return us, nil
//}
//
//func TestUserService_Create(t *testing.T) {
//	us, err := testingUserService()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	us.DestructiveReset()
//	user := User{
//		Name:  "Michael Scott",
//		Email: "michael@email.io",
//	}
//
//	err = us.Create(&user)
//	if err != nil {
//		t.Fatal(err)
//	}
//	if user.ID == 0 {
//		t.Errorf("Expected ID > 0. Received %d", user.ID)
//	}
//
//	if time.Since(user.CreatedAt) > time.Duration(5*time.Second) {
//		t.Errorf("Expected CreatedAt to be recent. Received %s", user.CreatedAt)
//	}
//
//	if time.Since(user.UpdatedAt) > time.Duration(5*time.Second) {
//		t.Errorf("Expected CreatedAt to be recent. Received %s", user.UpdatedAt)
//	}
//}
