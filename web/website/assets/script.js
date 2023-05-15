document.addEventListener("alpine:init", () => {
  let API = "http://localhost:8081/api";
  Alpine.store("user", {
    isAuthenticated: false,
    id: Alpine.$persist(0),
    email: Alpine.$persist(""),
    // CWE-256: Plaintext Storage of a Password
    password: Alpine.$persist(""),
    usdBalances: [],
    coinBalances: [],
    orders: [],
    transactions: [],
    async init() {
      // TODO: improve cookie parsing
      this.isAuthenticated = document.cookie.split("=")[1] != undefined;
      if (this.isAuthenticated) {
        await this.refreshData();
      }
    },

    async register() {
      parameters = new URLSearchParams({
        email: this.email,
        password: this.password,
      });
      let r = await fetch(`${API}/register?${parameters}`, {
        credentials: "include",
      });
      response = await r.text();
      return response;
    },
    async login() {
      parameters = new URLSearchParams({
        email: this.email,
        password: this.password,
      });
      let r = await fetch(`${API}/login?${parameters}`, {
        credentials: "include",
      });
      response = await r.text();
      if (r.status == 200) {
        // Read user_id from jwt payload
        this.id = JSON.parse(atob(response.split(".")[1])).user_id;
        this.isAuthenticated = true;
        await this.refreshData();
      } else {
        return response;
      }
    },
    logout() {
      document.cookie = `jwt=;expires=Thu, 01 Jan 1970 00:00:01 GMT;`;
      this.isAuthenticated = false;
    },

    async get(endpoint) {
      let r = await fetch(`${API}/${endpoint}`, {
        credentials: "include",
      });
      return await r.json();
    },
    async getCoinBalances() {
      this.coinBalances = await this.get("balances/coin");
    },
    async getUsdBalances() {
      this.usdBalances = await this.get("balances/usd");
    },
    async getOrders() {
      this.orders = await this.get("orders");
    },
    async getTransactions() {
      this.transactions = await this.get("transactions");
    },
    async refreshData() {
      await this.getCoinBalances();
      await this.getUsdBalances();
      await this.getOrders();
      await this.getTransactions();
    },
    async addOrder(coinId, qty, isBuy) {
      let r = await fetch(`${API}/orders`, {
        credentials: "include",
        method: "POST",
        body: JSON.stringify({
          coinId: coinId,
          isBuy: isBuy,
          qty: parseFloat(qty),
        }),
      });
      await this.refreshData();
      return r.text();
    },
    async addTransaction(coinId, address, qty, note) {
      let r = await fetch(`${API}/transactions`, {
        credentials: "include",
        method: "POST",
        body: JSON.stringify({
          coinId: coinId,
          address: address,
          qty: parseFloat(qty),
          note: note,
        }),
      });
      await this.refreshData();
      return r.text();
    },
    async updateEmail(newEmail) {
      let formData = new FormData();
      formData.append("email", newEmail);
      let r = await fetch(`${API}/user/email`, {
        credentials: "include",
        method: "PUT",
        body: formData,
      });
      if (r.status == 200) {
        this.email = newEmail;
      }
      return r.text();
    },
    async updatePassword(newPassword) {
      let formData = new FormData();
      formData.append("password", newPassword);
      let r = await fetch(`${API}/user/password`, {
        credentials: "include",
        method: "PUT",
        body: formData,
      });
      if (r.status == 200) {
        this.password = newPassword;
      }
      return r.text();
    },
  });

  Alpine.store("coins", {
    list: [],
    async init() {
      await this.refresh();
      setInterval(async () => {
        await this.refresh();
      }, 5000);
    },
    async refresh() {
      let r = await fetch(`${API}/coins`);
      this.list = await r.json();
    },
  });
});
