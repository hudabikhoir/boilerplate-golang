package order

//Order is transaction order user
type Order struct {
	ID                  string
	OrderID             string
	NoPKS               string
	CustomerID          string
	CustomerName        string
	FunderID            string
	MemberID            string
	HubID               string
	AgenID              string
	OrderDate           string
	OrderType           string
	Tenor               int
	TenorUnit           string
	DueDate             string
	TotalAmountAsli     int
	TotalAmountMargin   int
	TotalAmountPetani   int
	TotalAmountKodeUnik int
	LinkBuktiPemesanan  string
	Notes               string
	Details             []OrderDetail
	SystemSuccess       bool
	CreatedBy           CreatedBy
	CreatedAt           string
	DirtyBy             DirtyBy
	DirtyAt             string
	PointID             int
	OrderStatus         string
	StatusReason        string
}

//OrderDetail is transaction detail order user
type OrderDetail struct {
	OrderPakanID        string
	OrderID             string
	KodeBarang          string
	JenisBarang         string
	SupplierID          string
	FunderID            string
	DetailPakanJumlahKG int
	Merek               string
	AreaHargaPakan      string
	Qty                 int
	Unit                string
	UnitPriceAsli       int
	UnitPricePetani     int
	SubTotalAsli        int
	SubTotalPetani      int
	UsedHistoryID       string
	ProductID           int
}

//CreatedBy user who created data
type CreatedBy struct {
	Name  string
	Email string
}

//DirtyBy user who delete data
type DirtyBy struct {
	Name  string
	Email string
}
