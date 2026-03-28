package formcontroller_test

// func TestDoValidation(t *testing.T) {
// 	testCases := []struct {
// 		Desc                       string
// 		Form                       entities.Form
// 		mockedValidationController func(*testing.T, *mock_validationcontroller.MockIValidationController)
// 	}{
// 		{
// 			Desc: "TEST001 - Test validation for happy path",
// 			Form: entities.Form{
// 				Fields: []entities.FormFields{
// 					{
// 						Name:      "txbUsername",
// 						InputType: "",
// 					},
// 				},
// 			},
// 			mockedValidationController: func(t *testing.T, mic *mock_validationcontroller.MockIValidationController) {
// 				mic.EXPECT().DoValidation().Return(true, []validationcontroller.FieldValidator{})
// 			},
// 		},
// 	}
// 	for _, tc := range testCases {
// 		tc := tc

// 		t.Run(tc.Desc, func(t *testing.T) {

// 			ctrl := gomock.NewController(t)
// 			mockedValidationController := mock_validationcontroller.NewMockIValidationController(ctrl)
// 			if fn := tc.mockedValidationController; fn != nil {
// 				fn(t, mockedValidationController)
// 			}

// 			formController := formcontroller.NewFormController(tc.Form, mockedValidationController)

// 			formController.DoValidation()

// 			// assert.Equal(t, tc.ExpectedValidationResult, validationResult)
// 		})
// 	}
// }
