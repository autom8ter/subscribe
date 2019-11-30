package subscribe

import (
	"errors"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/product"
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

func (s *Subscriber) NewUser(name, email, phone *string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{}
	params.Name = name
	params.Email = email
	params.Phone = phone
	return customer.New(params)
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

func (s *Subscriber) NewSubscription(subscription *SubscriptionOpts) (*stripe.Subscription, error) {
	if subscription.Card != nil {
		params := &stripe.CardParams{}
		params.Number = stripe.String(subscription.Card.Number)
		params.CVC = stripe.String(subscription.Card.CVC)
		params.ExpMonth = stripe.String(subscription.Card.ExpMonth)
		params.ExpMonth = stripe.String(subscription.Card.ExpYear)
		params.Name = stripe.String(subscription.Card.Name)
		params.AddressCity = stripe.String(subscription.Card.AddressCity)
		params.AddressCountry = stripe.String(subscription.Card.AddressCountry)
		params.AddressState = stripe.String(subscription.Card.AddressState)
		params.AddressLine1 = stripe.String(subscription.Card.AddressLine1)
		params.AddressZip = stripe.String(subscription.Card.AddressZip)
		return sub.New(&stripe.SubscriptionParams{
			Card:             params,
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

func (s *Subscriber) ListSubscriptions(id string) *sub.Iter {
	return sub.List(nil)
}

func (s *Subscriber) UpdateSubscription(id string) (*stripe.Subscription, error) {
	return sub.Get(id, nil)
}

func (s *Subscriber) ListUsers(id string) *customer.Iter {
	return customer.List(nil)
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
