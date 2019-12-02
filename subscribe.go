//go:generate godocdown -o README.md

package subscribe

import (
	"errors"
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/product"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/sub"
)

type Subscriber struct {
	Secret string
}

func New(apiKey string) *Subscriber {
	stripe.Key = apiKey
	return &Subscriber{
		Secret: apiKey,
	}
}

type UserOpts struct {
	Name string
	Email string
	Phone string
	DefaultCard *CardOpts
}

func (u *UserOpts) ToCustomerParams() *stripe.CustomerParams{
	params := &stripe.CustomerParams{}
	if u.Name != "" {
		params.Name = stripe.String(u.Name)
	}
	if u.Email != "" {
		params.Email = stripe.String(u.Email)
	}
	if u.Phone != "" {
		params.Phone = stripe.String(u.Phone)
	}
	return params
}

func (s *Subscriber) NewUser(u *UserOpts) (*stripe.Customer, error) {
	return customer.New(u.ToCustomerParams())
}

func (s *Subscriber) DeleteUser(id string) (*stripe.Customer, error) {
	return customer.Del(id, nil)
}

func (s *Subscriber) GetUser(id string) (*stripe.Customer, error) {
	return customer.Get(id, nil)
}

func (s *Subscriber) UpdateUserName(id string, name string) (*stripe.Customer, error) {
	return customer.Update(id, &stripe.CustomerParams{
		Name: stripe.String(name),
	})
}

func (s *Subscriber) UpdateUserPhone(id string, phone string) (*stripe.Customer, error) {
	return customer.Update(id, &stripe.CustomerParams{
		Phone: stripe.String(phone),
	})
}

func (s *Subscriber) UpdateUserEmail(id string, email string) (*stripe.Customer, error) {
	return customer.Update(id, &stripe.CustomerParams{
		Email: stripe.String(email),
	})
}

func (s *Subscriber) UpdateUserPlan(id string, plan string) (*stripe.Customer, error) {
	return customer.Update(id, &stripe.CustomerParams{
		Plan: stripe.String(plan),
	})
}

func (s *Subscriber) UpdateUserBalance(id string, balance int64) (*stripe.Customer, error) {
	return customer.Update(id, &stripe.CustomerParams{
		Balance: stripe.Int64(balance),
	})
}

type SubscriptionOpts struct {
	CustomerID string
	Plan       string
	Card       *CardOpts
}

type CardOpts struct {
	Name           string
	Number         string
	CVC            string
	ExpMonth       string
	ExpYear        string
	AddressCity    string
	AddressCountry string
	AddressState   string
	AddressLine1   string
	AddressZip     string
}

func (c *CardOpts) ToParams() *stripe.CardParams {
	params := &stripe.CardParams{}
	params.Number = stripe.String(c.Number)
	params.CVC = stripe.String(c.CVC)
	params.ExpMonth = stripe.String(c.ExpMonth)
	params.ExpMonth = stripe.String(c.ExpYear)
	params.Name = stripe.String(c.Name)
	params.AddressCity = stripe.String(c.AddressCity)
	params.AddressCountry = stripe.String(c.AddressCountry)
	params.AddressState = stripe.String(c.AddressState)
	params.AddressLine1 = stripe.String(c.AddressLine1)
	params.AddressZip = stripe.String(c.AddressZip)
	return params
}

func (s *Subscriber) NewSubscription(subscription *SubscriptionOpts) (*stripe.Subscription, error) {
	if subscription.Card != nil {
		return sub.New(&stripe.SubscriptionParams{
			Card:             subscription.Card.ToParams(),
			CollectionMethod: stripe.String(string(stripe.SubscriptionCollectionMethodChargeAutomatically)),
			Customer:         stripe.String(subscription.CustomerID),
			Plan:             stripe.String(subscription.Plan),
		})
	}
	return nil, errors.New("empty subscription card info")
}

func (s *Subscriber) CancelSubscription(id string) (*stripe.Subscription, error) {
	return sub.Cancel(id, nil)
}

func (s *Subscriber) GetSubscription(id string) (*stripe.Subscription, error) {
	return sub.Get(id, nil)
}

func (s *Subscriber) ListSubscriptions(id string, limit int) *sub.Iter {
	params := &stripe.SubscriptionListParams{}
	limitList(limit, &params.Filters)
	return sub.List(params)
}

func (s *Subscriber) UpdateSubscription(id string) (*stripe.Subscription, error) {
	return sub.Get(id, nil)
}

func (s *Subscriber) ListUsers(id string, limit int) *customer.Iter {
	params := &stripe.CustomerListParams{}
	limitList(limit, &params.Filters)
	return customer.List(params)
}

type PlanOpts struct {
	Active   bool
	Amount   int64
	Currency string
	Interval string
	APIID    string
}

func (s *Subscriber) NewPlan(planOpts *PlanOpts) (*stripe.Plan, error) {
	return plan.New(&stripe.PlanParams{
		Active:    stripe.Bool(planOpts.Active),
		Amount:    stripe.Int64(planOpts.Amount),
		Currency:  stripe.String(planOpts.Currency),
		Interval:  stripe.String(planOpts.Interval),
		ProductID: stripe.String(planOpts.APIID),
	})
}

func (s *Subscriber) GetPlan(id string) (*stripe.Plan, error) {
	return plan.Get(id, nil)
}

func (s *Subscriber) ListPlans(id string) *plan.Iter {
	return plan.List(nil)
}

func (s *Subscriber) DeletePlan(id string) (*stripe.Plan, error) {
	return plan.Del(id, nil)
}

func (s *Subscriber) NewAPI(name, description string) (*stripe.Product, error) {
	return product.New(&stripe.ProductParams{
		Active:      stripe.Bool(true),
		Name:        stripe.String(name),
		Type:        stripe.String("service"),
		Description: stripe.String(description),
	})
}

func (s *Subscriber) ListAPIs() *product.Iter {
	return product.List(nil)
}

func (s *Subscriber) GetAPI(id string) (*stripe.Product, error) {
	return product.Get(id, nil)
}

func (s *Subscriber) UpdateAPIName(id, name string) (*stripe.Product, error) {
	return product.Update(id, &stripe.ProductParams{
		Active: stripe.Bool(true),
		Name:   stripe.String(name),
	})
}

func (s *Subscriber) UpdateAPIDescription(id, description string) (*stripe.Product, error) {
	return product.Update(id, &stripe.ProductParams{
		Active:      stripe.Bool(true),
		Description: stripe.String(description),
	})
}

type ChargeOpts struct {
	Amount int64
	Currency string
	CustomerId string
	Description string
	Card *CardOpts
}

func (s *Subscriber) ChargeUser(c *ChargeOpts) (*stripe.Charge, error) {
	return charge.New(&stripe.ChargeParams{
		Amount:                    stripe.Int64(c.Amount),
		Capture:                   nil,
		Currency:                  stripe.String(c.Currency),
		Customer:                  stripe.String(c.CustomerId),
		Description: stripe.String(c.Description),
		Source:                    &stripe.SourceParams{
			Card:  c.Card.ToParams(),
		},
	})
}

func (s *Subscriber) GetCharge(id string) (*stripe.Charge, error) {
	return charge.Get(id, nil)
}

func (s *Subscriber) ListCharges(limit int) (*charge.Iter) {
	params := &stripe.ChargeListParams{}
	limitList(limit, &params.Filters)
	return charge.List(&stripe.ChargeListParams{})
}

func (s *Subscriber) UpdateChargeCard(id string, c *CardOpts) (*stripe.Charge, error) {
	return charge.Update(id, &stripe.ChargeParams{
		Source: &stripe.SourceParams{
			Card: c.ToParams(),
		},
	})
}

func (s *Subscriber) UpdateChargeCustomer(id, customerId string) (*stripe.Charge, error) {
	return charge.Update(id, &stripe.ChargeParams{
		Customer: stripe.String(customerId),
	})
}

func (s *Subscriber) UpdateChargeAmount(id string, amount int64) (*stripe.Charge, error) {
	return charge.Update(id, &stripe.ChargeParams{
		Amount: stripe.Int64(amount),
	})
}


func limitList(limit int, filters *stripe.Filters) {
	if limit != 0 {
		filters.AddFilter("limit", "", fmt.Sprintf("%d", limit))
	}
}