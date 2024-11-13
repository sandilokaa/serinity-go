package transaction

import "time"

type TransactionFormatter struct {
	ID         int
	UserID     int
	ClothID    int
	Quantity   int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.UserID = transaction.UserID
	formatter.ClothID = transaction.ClothID
	formatter.Quantity = transaction.Quantity
	formatter.Amount = transaction.Amount
	formatter.Code = transaction.Code
	formatter.Status = transaction.Status
	formatter.PaymentURL = transaction.PaymentURL

	return formatter
}

type UserTransactionFormatter struct {
	ID        int            `json:"id"`
	Quantity  int            `json:"quantity"`
	Amount    int            `json:"amount"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	Cloth     ClothFormatter `json:"cloth"`
}

type ClothFormatter struct {
	Name      string                    `json:"name"`
	Color     string                    `json:"color"`
	Size      string                    `json:"size"`
	ImageURL  string                    `json:"image_url"`
	Variation []ClothVariationFormatter `json:"variation"`
}

type ClothVariationFormatter struct {
	Size  string `json:"size"`
	Color string `json:"color"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Quantity = transaction.Quantity
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	clothFormatter := ClothFormatter{}
	clothFormatter.Name = transaction.Cloth.Name
	clothFormatter.ImageURL = ""

	if len(transaction.Cloth.ClothImages) > 0 {
		clothFormatter.ImageURL = transaction.Cloth.ClothImages[0].FileName
	}

	for _, variation := range transaction.Cloth.Variation {
		clothVariationFormatter := ClothVariationFormatter{
			Color: variation.Color,
			Size:  variation.Size,
		}

		formatter.Cloth.Variation = append(formatter.Cloth.Variation, clothVariationFormatter)
	}

	formatter.Cloth = clothFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type TransactionDetailFormatter struct {
	ID        int                          `json:"id"`
	Quantity  int                          `json:"quantity"`
	Amount    int                          `json:"amount"`
	Status    string                       `json:"status"`
	CreatedAt time.Time                    `json:"created_at"`
	Cloth     ClothDetailFormatter         `json:"cloth"`
	Material  ClothDetailMaterialFormatter `json:"material"`
	User      ClothUserProfileFormatter    `json:"user"`
}

type ClothDetailFormatter struct {
	Name      string                          `json:"name"`
	Color     string                          `json:"color"`
	Size      string                          `json:"size"`
	ImageURL  string                          `json:"image_url"`
	Variation []ClothVariationDetailFormatter `json:"variation"`
}

type ClothDetailMaterialFormatter struct {
	MaterialName string `json:"material_name"`
}

type ClothUserProfileFormatter struct {
	Name string `json:"name"`
}

type ClothVariationDetailFormatter struct {
	Size  string `json:"size"`
	Color string `json:"color"`
}

func FormatTransactionDetail(transaction Transaction) TransactionDetailFormatter {
	formatter := TransactionDetailFormatter{}
	formatter.ID = transaction.ID
	formatter.Quantity = transaction.Quantity
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	clothFormatter := ClothDetailFormatter{}
	clothFormatter.Name = transaction.Cloth.Name
	clothFormatter.ImageURL = ""

	materialFormatter := ClothDetailMaterialFormatter{}
	materialFormatter.MaterialName = transaction.Cloth.Material.MaterialName

	userFormatter := ClothUserProfileFormatter{}
	userFormatter.Name = transaction.User.Name

	if len(transaction.Cloth.ClothImages) > 0 {
		clothFormatter.ImageURL = transaction.Cloth.ClothImages[0].FileName
	}

	for _, variation := range transaction.Cloth.Variation {
		clothVariationDetailFormatter := ClothVariationDetailFormatter{
			Color: variation.Color,
			Size:  variation.Size,
		}

		formatter.Cloth.Variation = append(formatter.Cloth.Variation, clothVariationDetailFormatter)
	}

	formatter.Cloth = clothFormatter
	formatter.Material = materialFormatter
	formatter.User = userFormatter

	return formatter
}
