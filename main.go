import "fmt"

///

type GatewayPagamento interface {
	Pagar(account string, valor float64) error
	Depositar(account string, valor float64) error
}

type PagSeguro struct {
	Token string
}

func (p *PagSeguro) Depositar(account string, valor float64) error {
	/// request para https://pagseguro.uol.com.br/api/pagamento/barbara/joia/babi?account={account}
	return nil
}

func (p *PagSeguro) Pagar(account string, valor float64) error {
	/// request para https://pagseguro.uol.com.br/api/pagar/barbara/joia/babi?account={account}
	return nil
}

type Stone struct {
}

func (s *Stone) Depositar(account string, valor float64) error {
	/// request para https://stone.com.br/api/pagamento/barbara/joia/babi?account={account}
	return nil
}

func (s *Stone) Pagar(account string, valor float64) error {
	/// request para https://stone.com.br/api/pagar/barbara/joia/babi?account={account}
	return nil
}

func NewGatewayPagamento(name string) GatewayPagamento {
	if name == "pagseguro" {
		return &PagSeguro{}
	} else {
		return &Stone{}
	}
}

func APIBlah() {
	gp := NewGatewayPagamento("stone")
	err := gp.Pagar("account", 50.12)
	fmt.Println("blah", err)
}
