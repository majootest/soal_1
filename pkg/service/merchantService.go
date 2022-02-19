package service

import (
	"fmt"

	"github.com/majoo_test/soal_1/internal/pkg"
	"github.com/majoo_test/soal_1/pkg/entity"
)

type MerchantService struct {
	repo   entity.MerchantRepository
	txRepo entity.TransactionRepository
}

func NewMerchantService(repo entity.MerchantRepository, txRepo entity.TransactionRepository) *MerchantService {
	return &MerchantService{repo, txRepo}
}

func (srv *MerchantService) FindOmzetReportNovember(userId int64, pagination *entity.Pagination) (results []entity.MerchantOmzet, err *pkg.Errors) {

	results = make([]entity.MerchantOmzet, 0)

	merchantList, e := srv.repo.FindByUserID(userId)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Error getting merchant list : %s", e.Error()),
			500,
		)
	}

	merchantIds := make([]int64, 0)

	for _, v := range merchantList {
		merchantIds = append(merchantIds, v.ID)
	}

	dateFrom := "2021-11-01 00:00:00"
	dateTo := "2021-11-30 23:59:59"
	txList, e := srv.txRepo.FindByMerchant(merchantIds, dateFrom, dateTo, pagination)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Error getting transaction report : %s", e.Error()),
			500,
		)
	}

	mapMerchant := make(map[int64]string)
	for _, v := range txList {
		merchantName, ok := mapMerchant[v.MerchantID]
		if !ok {
			merchant, e := srv.repo.FindByID(v.MerchantID)
			if e != nil {
				err = pkg.NewError(
					fmt.Sprintf("Error getting merchant data : %s", e.Error()),
					500,
				)
			}
			merchantName = merchant.MerchantName
			mapMerchant[v.MerchantID] = merchantName
		}

		results = append(
			results,
			entity.MerchantOmzet{
				MerchantName: merchantName,
				Omzet:        v.BillTotal,
			},
		)
	}
	return
}
