# Gobank

An digital bank for studying

## Entities

### Customer
Representa a entidade juridica ou física. Quem é o dono do dinheiro e dos acessos?
Guarda dados cadastrais e status de conformidade (KYC).

 - ID
 - Tipo (Individual, Business, System)
 - Nome
 - Documento
 - Email
 - Senha (auth-service)
 - Telefone
 - Endereço
 - Status (pendente_documentos, em analise, approved, rejeitada)


### Account
A conta bancária representa o registro financeiro. Onde o dinheiro oficial está depositado?
É usada para transações que transitam pelo sistema financeiro. Contas de classes asset são ativos e contas das classes liability e equity são passivos. Contas de ativo são bens do banco (debito aumenta e credito diminiu ativo) e contas passivas são dividas(credito aumenta e debito diminiu passivo) que o banco precisa honrar. Liabilities são passivos com os terceiros(clientes, bandeiras, fornecedores) e Equity são passivos com os donos ou investidores.

 - ID
 - CustomerID
 - Numero da Conta
 - Agencia
 - Tipo da Moeda (BRL, USD)
 - Classe (liability(passivo), equity(patrimonio/passivo), asset(ativo))
 - Tipo de Conta (Customer, vault, Clearing, Revenue)
 - Status (pendente, active, blocked, canceled)
 - Saldo

### Transaction 
A Transaction representa o fato que aconteceu. Registra a intenção e o contexto do movimentação do dinheiro. Quem mandou dinheiro para quem e quando, por qual motivo?

- ID
- Tipo (Pix, TED, DOC, Estorno, Cartão, Cartão de cédito)
- Valor total
- Status (Pendente, concluido, falhou)
- Descrição
- Timestamp

### AccountEntry
Cada Transaction gera duas ou mais AccountEntry. Ela representa o impacto real no saldo de uma conta específica. Partidas dobradas, Sai de uma Account (Debito) e entra em outra (Crédito). É o Ledger.

 - ID
 - TransactionID
 - AccountID
 - Type (credito ou débito)
 - Amount
 - timestamp


 # Onboarding Macro

 [Internet Banking] 
       │
       ▼ (POST /api/v1/onboarding)
[API Gateway] 
       │
       ▼ (Repassa a requisição inteira)
[OnboardingService] (O Orquestrador)
       │
       ├──► 1. Chama [UserService]    -> "Cria cliente UNDER_ANALYSIS"
       ├──► 2. Chama [AccountService] -> "Cria conta PENDING"
       ├──► 3. Chama [KYC Service]    -> "Valida os documentos"
       │
       ▼ (Após o resultado do KYC)
       └──► 4. Atualiza os estados no [UserService] e [AccountService]






