package dto

type Base struct {
	Query  string `json:"query"`
	Result Result `json:"result"`
}

type Result struct {
	Description        string           `json:"description"`
	Business_line_data BusinessLineData `json:"business_line_data"`
	Coverage_data      []CoverageData   `json:"coverage_data"`
}

type BusinessLineData struct {
	Business_line                BusinessLine                 `json:"business_line"`
	Product_payment_method       []ProductPaymentMethod       `json:"product_payment_method"`
	Product_renewal_cycle        []ProductRenewalCycle        `json:"product_renewal_cycle"`
	Commercial_network_attribute []CommercialNetworkAttribute `json:"commercial_network_attribute"`
	Renewal_parameter            []RenewalParameter           `json:"renewal_parameter"`
}

type CoverageData struct {
	Coverage                    Coverage                    `json:"coverage"`
	Capital                     []Capital                   `json:"capital"`
	Coverage_concept            []CoverageConcept           `json:"coverage_concept"`
	Deductible                  []Deductible                `json:"deductible"`
	Subcoverage                 []Subcoverage               `json:"subcoverage"`
	Temporary_premium_reduction []TemporaryPremiumReduction `json:"temporary_premium_reduction"`
}

type BusinessLine struct {
	Ciaascod  string `json:"ciasscod"`
	Producto  string `json:"producto"`
	Ramopcod  string `json:"ramopcod"`
	Flujofor  string `json:"flujofor"`
	Antidias  int    `json:"antidias"`
	Cicloren  string `json:"cicloren"`
	Duracann  int    `json:"duracann"`
	Descridab string `json:"descridab"`
	Duractip  int    `json:"duractip"`
	Fechaefe  string `json:"fechaefe"`
	Fmprocol  string `json:"fmprocol"`
	Lsfljfor  string `json:"lsfljfor"`
	Ndiasdps  int    `json:"ndiasdps"`
	Prodbase  string `json:"prodbase"`
	Producer  string `json:"producer"`
	Productip string `json:"productip"`
	Swactivo  bool   `json:"swactivo"`
	Swctlcb   bool   `json:"swxtlcb"`
	Swreacep  bool   `json:"swreacep"`
	Swsusbpm  bool   `json:"swsusbpm"`
	Swtramed  bool   `json:"swtramed"`
	Swtranot  bool   `json:"swtranot"`
	Swunicue  bool   `json:"swunicue"`
	Tablabas  string `json:"tablabas"`
	Tablacom  string `json:"tablacom"`
	Tipocole  string `json:"tipocole"`
	Tipoprod  string `json:"tipoprod"`
}

type ProductPaymentMethod struct {
	Mediopag string `json:"mediopag"`
	Medireci string `json:"medireci"`
	Mediexto string `json:"mediexto"`
	Swexcole bool   `json:"swexcole"`
	Recimini int    `json:"recimini"`
}

type ProductRenewalCycle struct {
	Cicloren string `json:"cicloren"`
}

type CommercialNetworkAttribute struct {
	Fieldcod    string `json:"fieldcod"`
	Descripcion string `json:"descripcion"`
	Orden       int    `json:"orden"`
	Swobliga    bool   `json:"swobliga"`
	Lenght      int    `json:"lenght"`
	Precision   int    `json:"Precision"`
	Tipodato    string `json:"tipodato"`
	Classhelp   string `json:"classhelp"`
}

type RenewalParameter struct {
	Swrenova bool   `json:"swrenova"`
	Proceren string `json:"proceren"`
	Swpreren bool   `json:"swpreren"`
	Ndiasren int    `json:"ndiasren"`
	Swvigrie bool   `json:"swvigrie"`
	Ndiaspre int    `json:"ndiaspre"`
	Ndiascnr int    `json:"ndiascnr"`
	Ndiascre int    `json:"ndiascre"`
	NdiasVri int    `json:"ndiasvri"`
}

type Coverage struct {
	Garancod    string `json:"garancod"`
	Gramodgs    string `json:"gramodgs"`
	Cobernum    int    `json:"cobernum"`
	Agrupacion  string `json:"agrupacion"`
	Garappal    string `json:"garappal"`
	Fechaefe    string `json:"fechaefe"`
	Fechbaja    string `json:"fechbaja"`
	Edadmini    int    `json:"edadmini"`
	Edadmaxi    int    `json:"edadmaxi"`
	Swpriniv    bool   `json:"swpriniv"`
	Swtranot    bool   `json:"swtranot"`
	Swopcdto    bool   `json:"swopcdto"`
	Swtarman    bool   `json:"swtarman"`
	Swbonif     bool   `json:"swbonif"`
	Swobliga    bool   `json:"swobliga"`
	Swfran      bool   `json:"swfran"`
	Swcapita    bool   `json:"swcapita"`
	Swcalpro    bool   `json:"swcalpro"`
	Swsimpro    bool   `json:"swsimpro"`
	Cobecmin    int    `json:"cobecmin"`
	Cobecmax    int    `json:"cobecmax"`
	Primamin    int    `json:"primamin"`
	Dedumaxi    int    `json:"dedumaxi"`
	Recamaxi    int    `json:"recamaxi"`
	Cobepmin    int    `json:"cobepmin"`
	Cobepmax    int    `json:"cobepmax"`
	Porcreva    int    `json:"porcreva"`
	Basestec    string `json:"basestec"`
	Swips       bool   `json:"swips"`
	Swclea      bool   `json:"swclea"`
	Edadsali    int    `json:"edadsali"`
	Swcons      bool   `json:"swcons"`
	Tipocapcons string `json:"tipocapcons"`
	Tiporecasir string `json:"tiporecasir"`
}

type Capital struct {
	Capitdat    string `json:"capitdat"`
	Descripcion string `json:"descripcion"`
	Swforpor    bool   `json:"swforpor"`
	Swforcap    bool   `json:"swforcap"`
}

type CoverageConcept struct {
	Conceimp  string `json:"conceimp"`
	Tipoimpu  string `json:"tipoimpu"`
	Swliqimpu bool   `json:"swliqimpu"`
	Gastocob  int    `json:"gastocob"`
}

type Deductible struct {
	Codifran    string `json:"codifran"`
	Descripcion string `json:"descripcion"`
	Diasfran    string `json:"diasfran"`
	Impofran    int    `json:"impofran"`
}

type Subcoverage struct {
	Cobernum int    `json:"cobernum"`
	Capiaseg int    `json:"capiaseg"`
	Impofran int    `json:"impofran"`
	Codsubco string `json:"codsubco"`
	Porccapi int    `json:"porccapi"`
	Porcrepa int    `json:"porcrepa"`
}

type TemporaryPremiumReduction struct {
	Diasanti int `json:"diasanti"`
	Porcprim int `json:"porcprim"`
}
