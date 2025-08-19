package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentsInfo  map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

// Метод Pay() принимает:
// 1. Описание проводимой оплаты
// 2. Сумма оплаты
// Возвращает:
// 1. ID проведенной операции
func (p PaymentModule) Pay(description string, usd int) int {
	// 1. проводить оплату
	// 2. получать id проведенной оплаты
	id := p.paymentMethod.Pay(usd)

	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}

	// 3. сохранять информацию о проведенной операции
	// - описание операции
	// - сколько было потрачено
	// - отмененная ли операция
	p.paymentsInfo[id] = info

	// 4. возвращать id проведенной оплаты
	return id
}

// Метод Cancel()
// Принимает:
// 1. ID операции
// Возвращает:
// - ничего не возвращает
func (p PaymentModule) Cancel(id int) {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}

	p.paymentMethod.Cancel(id)

	info.Cancelled = true

	p.paymentsInfo[id] = info
}

// Метод Info()
// Принимает:
// 1. ID операции
// Возвращает:
// 1. Информацию о проведенной операции
func (p PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}

	return info
}

// Метод AllInfo()
// Принимает:
// - ничего не принимает
// Возвращает:
// Возвращает информацию о всех проведенных операциях
func (p PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))
	for k, v := range p.paymentsInfo {
		tempMap[k] = v
	}

	return tempMap // временная мапа необходима для того, чтобы оригинальная мапа не подверглась изменению
}
