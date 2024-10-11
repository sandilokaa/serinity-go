package supplier

type SupplierFormatter struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Postal  string `json:"postal"`
}

func FormatSupplier(supplier Supplier) SupplierFormatter {
	supplierFormatter := SupplierFormatter{}
	supplierFormatter.ID = supplier.ID
	supplierFormatter.UserID = supplier.UserID
	supplierFormatter.Name = supplier.Name
	supplierFormatter.Address = supplier.Address
	supplierFormatter.Postal = supplier.Postal

	return supplierFormatter
}

func FormatSuppliers(suppliers []Supplier) []SupplierFormatter {
	suppliersFormatter := []SupplierFormatter{}

	for _, supplier := range suppliers {
		supplierFormatter := FormatSupplier(supplier)
		suppliersFormatter = append(suppliersFormatter, supplierFormatter)
	}

	return suppliersFormatter
}

type SupplierDetailFormatter struct {
	ID      int                   `json:"id"`
	UserID  int                   `json:"user_id"`
	Name    string                `json:"name"`
	Address string                `json:"address"`
	Postal  string                `json:"postal"`
	User    SupplierUserFormatter `json:"user"`
}

type SupplierUserFormatter struct {
	Name string `json:"name"`
}

func (s *Supplier) FormatSupplierDetail(supplier Supplier) SupplierDetailFormatter {

	supplierDetailFormatter := SupplierDetailFormatter{}
	supplierDetailFormatter.ID = supplier.ID
	supplierDetailFormatter.Name = supplier.Name
	supplierDetailFormatter.Address = supplier.Address
	supplierDetailFormatter.Postal = supplier.Postal

	user := supplier.User
	supplierUserFormatter := SupplierUserFormatter{}
	supplierUserFormatter.Name = user.Name

	supplierDetailFormatter.User = supplierUserFormatter

	return supplierDetailFormatter
}

func UpdatedFormatSupplier(updatedSupplier Supplier, oldSupplier Supplier) map[string]interface{} {
	updatedFields := make(map[string]interface{})

	if oldSupplier.ID != updatedSupplier.ID {
		updatedFields["id"] = updatedSupplier.ID
	}
	if oldSupplier.UserID != updatedSupplier.UserID {
		updatedFields["user_id"] = updatedSupplier.UserID
	}
	if oldSupplier.Name != updatedSupplier.Name {
		updatedFields["name"] = updatedSupplier.Name
	}
	if oldSupplier.Address != updatedSupplier.Address {
		updatedFields["address"] = updatedSupplier.Address
	}
	if oldSupplier.Postal != updatedSupplier.Postal {
		updatedFields["postal"] = updatedSupplier.Postal
	}

	return updatedFields
}
