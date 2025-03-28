"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { IResponseSearch } from "../../../type/user";

export default function Dashboard() {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState<IResponseSearch[]>([]);
  const [error, setError] = useState("");
  const router = useRouter();

  useEffect(() => {
    const token = localStorage.getItem("access_token");
    if (!token) {
      router.push("/");
    }
  }, [router]);

  const handleSearch = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    const token = localStorage.getItem("access_token");
    if (!token) {
      router.push("/");
      return;
    }

    try {
      const res = await fetch(
        `http://localhost:8000/v1/user/search?q=${query}`,
        {
          method: "GET",
          headers: {
            Authorization: `Bearer ${token}`, // Add Authentication header
          },
        }
      );

      if (!res.ok) {
        throw new Error("Failed to fetch search results");
      }

      const response = await res.json();
      setResults(response.data);
    } catch {
      setError("An error occurred while searching");
    }
  };

  const handleLogout = () => {
    localStorage.removeItem("access_token");
    router.push("/");
  };

  return (
    <div style={{ padding: "20px" }}>
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          marginBottom: "20px",
        }}
      >
        <h1>Customer Search</h1>
        <button
          onClick={handleLogout}
          style={{
            padding: "8px 16px",
            backgroundColor: "#ff4444",
            color: "white",
            border: "none",
            borderRadius: "4px",
            cursor: "pointer",
          }}
        >
          Logout
        </button>
      </div>

      <form onSubmit={handleSearch}>
        <input
          type="text"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          placeholder="Search by name or username"
          style={{ padding: "8px", marginRight: "10px" }}
        />
        <button type="submit" style={{ padding: "8px 16px" }}>
          Search
        </button>
      </form>

      {error && <p style={{ color: "red", marginTop: "10px" }}>{error}</p>}

      <div>
        {results.map((customer, index) => (
          <div
            key={index}
            style={{
              margin: "20px 0",
              border: "1px solid #ccc",
              padding: "10px",
            }}
          >
            <h2>
              {customer.name} ({customer.username})
            </h2>
            <h3>Bank Account:</h3>
            <p>
              Account Number: {customer.account_number} - Balance: $
              {customer.balance}
            </p>
            <h4>Pockets:</h4>
            <ul>
              {customer.pockets.map((pocket, index) => (
                <li key={index}>
                  {pocket.pocket_name}: ${pocket.balance}
                </li>
              ))}
            </ul>
            <h4>Term Deposits:</h4>
            <ul>
              {customer.term_deposits.map((deposit, index) => (
                <li key={index}>
                  Amount: ${deposit.amount}, Rate: {deposit.interest_rate}%,
                  Term: {deposit.term_months} months
                </li>
              ))}
            </ul>
          </div>
        ))}
      </div>
    </div>
  );
}
