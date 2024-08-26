// Code generated by mtgroup-generator.   
package api

import (
	"errors"
	"time"
	"demo/internal/def"
	"github.com/go-openapi/swag"

	"demo/internal/api/restapi/models"
	transaction "demo/internal/api/restapi/restapi/operations/transaction"
	"demo/internal/app"
		"demo/internal/types"

	extauthapi "github.com/mtgroupit/mt-mock-extauthapi"
		"github.com/go-openapi/strfmt"
	"github.com/go-openapi/runtime/middleware" 
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!
func (svc *service) GetTransaction(params transaction.GetTransactionParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
		c, err := svc.app.GetTransaction(toAppProfile(prof), params.Body.ID)
		switch {
		default:
			log.PrintErr("GetTransaction server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
			return transaction.NewGetTransactionDefault(codeInternal.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeInternal.extra),
				Message: swag.String("internal error"),
			})
		case errors.Is(err, app.ErrAccessDenied):
			log.Info("GetTransaction client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
			return transaction.NewGetTransactionDefault(codeForbidden.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeForbidden.extra),
				Message: swag.String(err.Error()),
			})
		case errors.Is(err, app.ErrNotFound):
			log.Info("GetTransaction client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
			return transaction.NewGetTransactionDefault(codeNotFound.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeNotFound.extra),
				Message: swag.String(err.Error()),
			})
		case err == nil:
			log.Info("GetTransaction ok", "id", params.Body.ID)
			return transaction.NewGetTransactionOK().WithPayload(apiTransaction(c))
		}
}
func (svc *service) AddTransaction(params transaction.AddTransactionParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
		c, err := svc.app.AddTransaction(toAppProfile(prof), appTransactionAdd(params.Body))
		switch {
		default:
			log.PrintErr("AddTransaction server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
			return transaction.NewAddTransactionDefault(codeInternal.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeInternal.extra),
				Message: swag.String("internal error"),
			})
		case errors.Is(err, app.ErrAccessDenied):
			log.Info("AddTransaction client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
			return transaction.NewAddTransactionDefault(codeForbidden.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeForbidden.extra),
				Message: swag.String(err.Error()),
			})
		case err == nil:
			log.Info("AddTransaction ok")
			return transaction.NewAddTransactionCreated().WithPayload(apiTransaction(c))
		}
}
func (svc *service) DeleteTransaction(params transaction.DeleteTransactionParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
		err := svc.app.DeleteTransaction(toAppProfile(prof), params.Body.ID)
		switch {
		default:
			log.PrintErr("DeleteTransaction server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
			return transaction.NewDeleteTransactionDefault(codeInternal.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeInternal.extra),
				Message: swag.String("internal error"),
			})
		case errors.Is(err, app.ErrAccessDenied):
			log.Info("DeleteTransaction client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
			return transaction.NewDeleteTransactionDefault(codeForbidden.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeForbidden.extra),
				Message: swag.String(err.Error()),
			})
		case errors.Is(err, app.ErrNotFound):
			log.Info("DeleteTransaction client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
			return transaction.NewDeleteTransactionDefault(codeNotFound.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeNotFound.extra),
				Message: swag.String(err.Error()),
			})
		case err == nil:
			log.Info("DeleteTransaction ok", "id", params.Body.ID)
			return transaction.NewDeleteTransactionNoContent()
		}
}
func (svc *service) EditTransaction(params transaction.EditTransactionParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
		c, err := svc.app.EditTransaction(toAppProfile(prof), params.Body.ID, appTransactionAdd(params.Body.Data))
		switch {
		default:
			log.PrintErr("EditTransaction server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
			return transaction.NewEditTransactionDefault(codeInternal.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeInternal.extra),
				Message: swag.String("internal error"),
			})
		case errors.Is(err, app.ErrAccessDenied):
			log.Info("EditTransaction client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
			return transaction.NewEditTransactionDefault(codeForbidden.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeForbidden.extra),
				Message: swag.String(err.Error()),
			})
		case errors.Is(err, app.ErrNotFound):
			log.Info("EditTransaction client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
			return transaction.NewEditTransactionDefault(codeNotFound.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeNotFound.extra),
				Message: swag.String(err.Error()),
			})
		case err == nil:
			log.Info("EditTransaction ok")
			return transaction.NewEditTransactionOK().WithPayload(apiTransaction(c))
		}
}
func (svc *service) ListTransaction(params transaction.ListTransactionParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
		c, count, warnings, err := svc.app.ListTransaction(toAppProfile(prof), appListParams(params.Body))
		switch {
		default:
			log.PrintErr("ListTransaction server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
			return transaction.NewListTransactionDefault(codeInternal.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeInternal.extra),
				Message: swag.String("internal error"),
			})
		case errors.Is(err, app.ErrAccessDenied):
			log.Info("ListTransaction client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
			return transaction.NewListTransactionDefault(codeForbidden.status).WithPayload(&models.Error{
				Code:    swag.Int32(codeForbidden.extra),
				Message: swag.String(err.Error()),
			})
		case err == nil:
			log.Info("ListTransaction ok")
			return transaction.NewListTransactionOK().WithPayload(&transaction.ListTransactionOKBody{
				Items: apiTransactions(c),
				Warnings: warnings,
				Count: int32(count),
			})
		}
}
      



func apiTransaction(a *app.Transaction) *models.Transaction {
	if a == nil {
		return nil
	}
	return &models.Transaction{
			ID: a.ID,
			Amount: a.Amount.Float64(),
			CreatedAt: (*strfmt.DateTime)(a.CreatedAt),
			Info: a.Info,
			Type: a.Type,
			User: apiUser(a.User),
	}
}

func apiTransactions(apps []*app.Transaction) []*models.Transaction {
	apis := []*models.Transaction{}
	for i := range apps {
		apis = append(apis, apiTransaction(apps[i]))
	}
	return apis
}

func appTransaction(a *models.Transaction, withStructs bool) *app.Transaction {
	if a == nil {
		return nil
	}
	transaction := &app.Transaction{}
		if withStructs {
					transaction.User = appUser(a.User)
		}
			transaction.ID = a.ID
			transaction.Amount = types.NewDecimal(a.Amount)
			transaction.CreatedAt = (*time.Time)(a.CreatedAt)
			transaction.Info = a.Info
			transaction.Type = a.Type
	
	return transaction
}

func appTransactions(apis []*models.Transaction, withStructs bool) []*app.Transaction {
	apps := []*app.Transaction{}
	for i := range apis {
		apps = append(apps, appTransaction(apis[i], withStructs))
	}
	return apps
}

func appTransactionAdd(a *models.TransactionAdd) *app.Transaction {
	if a == nil {
		return nil
	}
	transaction := &app.Transaction{}
			transaction.Amount = types.NewDecimal(a.Amount)
			transaction.Info = a.Info
			transaction.Type = a.Type
			if a.User != "" {
				transaction.User = &app.User{ID: a.User}
			}
	
	
	return transaction
}

func appTransactionsAdd(apis []*models.TransactionAdd) []*app.Transaction {
	apps := []*app.Transaction{}
	for i := range apis {
		apps = append(apps, appTransactionAdd(apis[i]))
	}
	return apps
}
