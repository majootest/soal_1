package service

import (
	"fmt"

	"github.com/majoo_test/soal_1/internal/pkg"
	"github.com/majoo_test/soal_1/pkg/entity"
)

type OutletService struct {
	repo         entity.OutletRepository
	merchantRepo entity.MerchantRepository
	txRepo       entity.TransactionRepository
}

func NewOutletService(repo entity.OutletRepository, merchantRepo entity.MerchantRepository, txRepo entity.TransactionRepository) *OutletService {
	return &OutletService{repo, merchantRepo, txRepo}
}

func (srv *OutletService) FindOmzetReportNovember(userId int64, pagination *entity.Pagination) (results []entity.OutletOmzet, err *pkg.Errors) {

	results = make([]entity.OutletOmzet, 0)

	merchantList, e := srv.merchantRepo.FindByUserID(userId)
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

	outletList, e := srv.repo.FindByMerchantID(merchantIds)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Error getting outlet list : %s", e.Error()),
			500,
		)
	}

	outletIds := make([]int64, 0)
	for _, v := range outletList {
		outletIds = append(outletIds, v.ID)
	}

	dateFrom := "2021-11-01 00:00:00"
	dateTo := "2021-11-30 23:59:59"
	txList, e := srv.txRepo.FindByOutlet(outletIds, dateFrom, dateTo, pagination)
	if e != nil {
		err = pkg.NewError(
			fmt.Sprintf("Error getting transaction report : %s", e.Error()),
			500,
		)
	}

	mapMerchant := make(map[int64]string)
	mapOutlet := make(map[int64]string)
	for _, v := range txList {

		merchantName, ok := mapMerchant[v.MerchantID]
		if !ok {
			merchant, e := srv.merchantRepo.FindByID(v.MerchantID)
			if e != nil {
				err = pkg.NewError(
					fmt.Sprintf("Error getting merchant data : %s", e.Error()),
					500,
				)
			}
			merchantName = merchant.MerchantName
			mapMerchant[v.MerchantID] = merchantName
		}

		outletName, ok := mapMerchant[v.OutletID]
		if !ok {
			outlet, e := srv.repo.FindByID(v.OutletID)
			if e != nil {
				err = pkg.NewError(
					fmt.Sprintf("Error getting outlet data : %s", e.Error()),
					500,
				)
			}
			outletName = outlet.OutletName
			mapOutlet[v.OutletID] = outletName
		}

		results = append(
			results,
			entity.OutletOmzet{
				MerchantName: merchantName,
				OutletName:   outletName,
				Omzet:        v.BillTotal,
			},
		)
	}
	return
}
