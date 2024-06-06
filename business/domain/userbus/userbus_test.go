package userbus

//func Test_User(t *testing.T) {
//	t.Parallel()
//
//	db := dbtest.NewDatabase(t, "Test_User")
//	defer func() {
//		if r := recover(); r != nil {
//			t.Log(r)
//			t.Error(string(debug.Stack()))
//		}
//		db.Teardown()
//	}()
//
//	sd, err := insertSeedData(db.BusDomain)
//	if err != nil {
//		t.Fatalf("Seeding error: %s", err)
//	}
//
//	// -------------------------------------------------------------------------
//
//	unitest.Run(t, query(db.BusDomain, sd), "query")
//	// unitest.Run(t, create(db.BusDomain), "create")
//	// unitest.Run(t, update(db.BusDomain, sd), "update")
//	// unitest.Run(t, delete(db.BusDomain, sd), "delete")
//}
