package product

import (
	"context"
	"gofr-curd/models"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

// func TestCoreLayer(t *testing.T) {
// 	app := gofr.New()
// 	// seeder := datastore.NewSeeder(&app.DataStore,"../db")
// 	// seeder.ResetCounter = true
// 	db, mock, _ := sqlmock.New()
// 	database, err := gorm.Open("mysql", db)
// 	if err != nil {
// 		log.Println("Can't Open the DataBase")
// 	}
// 	app.ORM = database

// 	rows := sqlmock.NewRows([]string{"Id", "Name", "Type"}).AddRow(1, "daikinn", "AC")

// 	tests := []struct {
// 		desc        string
// 		id          int
// 		expected    *models.Product
// 		expectedErr error
// 		mockQuery   *sqlmock.ExpectedQuery
// 	}{
// 		{desc: "Case1", id: 1, expectedErr: nil, expected: &models.Product{Id: 1, Name: "daikinn", Type: "AC"}, mockQuery: mock.ExpectQuery("select * from Product where id = ?").WithArgs(1).WillReturnRows(rows)},
// 		{desc: "Case2", id: 100, expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"}, expected: nil, mockQuery: mock.ExpectQuery("select * from Product where id = ?").WithArgs(100).WillReturnError(sql.ErrNoRows)},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.desc, func(t *testing.T) {
// 			ctx := gofr.NewContext(nil, nil, app)
// 			ctx.Context = context.Background()
// 			istore := New()
// 			res, err := istore.GetProductById(ctx, test.id)
// 			if !reflect.DeepEqual(err, test.expectedErr) {
// 				t.Error("expected: ", test.expectedErr, "obtained: ", err)
// 			}
// 			if err == nil && !reflect.DeepEqual(test.expected, res) {
// 				t.Error("expected: ", test.expected, "obtained: ", res)
// 			}
// 		})
// 	}

// }

func TestStoreLayer(t *testing.T) {
	app := gofr.New()
	seeder := datastore.NewSeeder(&app.DataStore, "../db")
	seeder.ResetCounter = true
	testGetProductById(t, app)
	testGetAllProducts(t, app)
	testCreateProduct(t, app)
	testDeleteProduct(t, app)
	testUpdateProductById(t, app)

}

func testGetProductById(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc        string
		id          int
		expected    *models.Product
		expectedErr error
		// mockQuery   *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", id: 1, expectedErr: nil, expected: &models.Product{Id: 1, Name: "daikin", Type: "AC"}},
		{desc: "Case2", id: 100, expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"}, expected: nil},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			istore := New()
			res, err := istore.GetProductById(ctx, test.id)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			if err == nil && !reflect.DeepEqual(test.expected, res) {
				t.Error("expected: ", test.expected, "obtained: ", res)
			}
		})
	}

}

func testGetAllProducts(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc string
		// id          int
		expected    []*models.Product
		expectedErr error
		// mockQuery   *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", expectedErr: nil,
			expected: []*models.Product{&models.Product{Id: 1, Name: "daikin", Type: "AC"},
				&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"},
				&models.Product{Id: 3, Name: "Kenstar", Type: "Microwave"},
				&models.Product{Id: 4, Name: "Ultra", Type: "RedGrinder"},
				&models.Product{Id: 5, Name: "Crompton", Type: "Fan"},
				&models.Product{Id: 6, Name: "Prestige", Type: "RiceCooker"},
				&models.Product{Id: 13, Name: "Nivvea", Type: "Moisturizzerr"},
				&models.Product{Id: 16, Name: "Kenstarr1", Type: "Microwavee1"},
				&models.Product{Id: 18, Name: "Kenstarr7", Type: "Microwavee7"},
				&models.Product{Id: 21, Name: "daikin", Type: "AC"},
			},
		},
		// {desc: "Case2", expectedErr: errors.EntityNotFound{Entity: "Product"}, expected: []*models.Product{}},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			istore := New()
			res, err := istore.GetAllProducts(ctx)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			if err == nil && !reflect.DeepEqual(test.expected, res) {
				t.Error("expected: ", test.expected, "obtained: ", res)
			}
		})
	}

}

func testCreateProduct(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc string
		// id          int
		input models.Product
		// expected    *models.Product
		expected    int
		expectedErr error
		// mockQuery   *sqlmock.ExpectedQuery
	}{
		{desc: "Case1" /*id: 1,*/, input: models.Product{Name: "daikin", Type: "AC"}, expectedErr: nil, expected: 31},
		// {desc: "Case2", /*id: 100,*/ input:,expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"}, expected: nil},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			istore := New()
			res, err := istore.CreateProduct(ctx, test.input)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			if err == nil && !reflect.DeepEqual(test.expected, res) {
				t.Error("expected: ", test.expected, "obtained: ", res)
			}

		})
	}

}

func testDeleteProduct(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc string
		id   int
		// expected    *models.Product
		expectedErr error
		// mockQuery   *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", id: 31, expectedErr: nil},
		{desc: "Case2", id: 100, expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"}},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			istore := New()
			err := istore.DeleteById(ctx, test.id)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			// if err == nil && !reflect.DeepEqual(test.expected, res) {
			// 	t.Error("expected: ", test.expected, "obtained: ", res)
			// }
		})
	}

}

func testUpdateProductById(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc     string
		id       int
		input    models.Product
		expected int
		// expected    int
		expectedErr error
		// mockQuery   *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", id: 21, input: models.Product{Name: "daikinn", Type: "ACC"}, expectedErr: nil, expected: 21},
		{desc: "Case2", id: 21, input: models.Product{}, expectedErr: errors.Error("Nothing to Update")},
		{desc: "Case3", id: 21, input: models.Product{Name: "daikin", Type: "AC"}, expectedErr: nil, expected: 21},
		{desc: "Case4", id: 21, input: models.Product{Name: "daikin", Type: "AC"}, expectedErr: errors.Error("SAME DATA GIVEN TO PREVIOUS DATA"), expected: 21},

		// {desc: "Case2", /*id: 100,*/ input:,expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"}, expected: nil},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			istore := New()
			res, err := istore.UpdateById(ctx, test.id, test.input)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			if err == nil && !reflect.DeepEqual(test.expected, res) {
				t.Error("expected: ", test.expected, "obtained: ", res)
			}

		})
	}

}
