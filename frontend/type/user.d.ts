export interface IResponseSearch {
  id_user: number;
  name: string;
  username: string;
  level: number;
  created_at: string;
  updated_at: string;
  last_login: string | null;
  is_active: boolean;
  account_number: string;
  balance: number;
  pockets: IResponseUserPocket[];
  term_deposits: IResponseUserTermDeposit[];
}

export interface IResponseUserPocket {
  id: number;
  bank_account_id: number;
  pocket_name: string;
  balance: number;
  created_at: string;
}

export interface IResponseUserTermDeposit {
  id: number;
  bank_account_id: number;
  amount: number;
  interest_rate: number;
  term_months: number;
  start_date: string;
  maturity_date: string;
  status: string;
}
