# subscribe
--
    import "github.com/autom8ter/subscribe"


## Usage

#### type CardOpts

```go
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
```


#### func (*CardOpts) ToParams

```go
func (c *CardOpts) ToParams() *stripe.CardParams
```

#### type ChargeOpts

```go
type ChargeOpts struct {
	Amount      int64
	Currency    string
	CustomerId  string
	Description string
	Card        *CardOpts
}
```


#### type PlanOpts

```go
type PlanOpts struct {
	Active   bool
	Amount   int64
	Currency string
	Interval string
	APIID    string
}
```


#### type Subscriber

```go
type Subscriber struct {
	Secret string
}
```


#### func  New

```go
func New(apiKey string) *Subscriber
```

#### func (*Subscriber) CancelSubscription

```go
func (s *Subscriber) CancelSubscription(id string) (*stripe.Subscription, error)
```

#### func (*Subscriber) ChargeUser

```go
func (s *Subscriber) ChargeUser(c *ChargeOpts) (*stripe.Charge, error)
```

#### func (*Subscriber) DeletePlan

```go
func (s *Subscriber) DeletePlan(id string) (*stripe.Plan, error)
```

#### func (*Subscriber) DeleteUser

```go
func (s *Subscriber) DeleteUser(id string) (*stripe.Customer, error)
```

#### func (*Subscriber) GetAPI

```go
func (s *Subscriber) GetAPI(id string) (*stripe.Product, error)
```

#### func (*Subscriber) GetCharge

```go
func (s *Subscriber) GetCharge(id string) (*stripe.Charge, error)
```

#### func (*Subscriber) GetPlan

```go
func (s *Subscriber) GetPlan(id string) (*stripe.Plan, error)
```

#### func (*Subscriber) GetSubscription

```go
func (s *Subscriber) GetSubscription(id string) (*stripe.Subscription, error)
```

#### func (*Subscriber) GetUser

```go
func (s *Subscriber) GetUser(id string) (*stripe.Customer, error)
```

#### func (*Subscriber) ListAPIs

```go
func (s *Subscriber) ListAPIs() *product.Iter
```

#### func (*Subscriber) ListCharges

```go
func (s *Subscriber) ListCharges(limit int) *charge.Iter
```

#### func (*Subscriber) ListPlans

```go
func (s *Subscriber) ListPlans(id string) *plan.Iter
```

#### func (*Subscriber) ListSubscriptions

```go
func (s *Subscriber) ListSubscriptions(id string) *sub.Iter
```

#### func (*Subscriber) ListUsers

```go
func (s *Subscriber) ListUsers(id string) *customer.Iter
```

#### func (*Subscriber) NewAPI

```go
func (s *Subscriber) NewAPI(name, description string) (*stripe.Product, error)
```

#### func (*Subscriber) NewPlan

```go
func (s *Subscriber) NewPlan(planOpts *PlanOpts) (*stripe.Plan, error)
```

#### func (*Subscriber) NewSubscription

```go
func (s *Subscriber) NewSubscription(subscription *SubscriptionOpts) (*stripe.Subscription, error)
```

#### func (*Subscriber) NewUser

```go
func (s *Subscriber) NewUser(name, email, phone *string) (*stripe.Customer, error)
```

#### func (*Subscriber) UpdateAPIDescription

```go
func (s *Subscriber) UpdateAPIDescription(id, description string) (*stripe.Product, error)
```

#### func (*Subscriber) UpdateAPIName

```go
func (s *Subscriber) UpdateAPIName(id, name string) (*stripe.Product, error)
```

#### func (*Subscriber) UpdateChargeAmount

```go
func (s *Subscriber) UpdateChargeAmount(id string, amount int64) (*stripe.Charge, error)
```

#### func (*Subscriber) UpdateChargeCard

```go
func (s *Subscriber) UpdateChargeCard(id string, c *CardOpts) (*stripe.Charge, error)
```

#### func (*Subscriber) UpdateChargeCustomer

```go
func (s *Subscriber) UpdateChargeCustomer(id, customerId string) (*stripe.Charge, error)
```

#### func (*Subscriber) UpdateSubscription

```go
func (s *Subscriber) UpdateSubscription(id string) (*stripe.Subscription, error)
```

#### func (*Subscriber) UpdateUserBalance

```go
func (s *Subscriber) UpdateUserBalance(id string, balance int64) (*stripe.Customer, error)
```

#### func (*Subscriber) UpdateUserEmail

```go
func (s *Subscriber) UpdateUserEmail(id string, email string) (*stripe.Customer, error)
```

#### func (*Subscriber) UpdateUserName

```go
func (s *Subscriber) UpdateUserName(id string, name string) (*stripe.Customer, error)
```

#### func (*Subscriber) UpdateUserPhone

```go
func (s *Subscriber) UpdateUserPhone(id string, phone string) (*stripe.Customer, error)
```

#### func (*Subscriber) UpdateUserPlan

```go
func (s *Subscriber) UpdateUserPlan(id string, plan string) (*stripe.Customer, error)
```

#### type SubscriptionOpts

```go
type SubscriptionOpts struct {
	CustomerID string
	Plan       string
	Card       *CardOpts
}
```
