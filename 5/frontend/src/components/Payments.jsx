import { useState } from "react";

const API_URL = "http://localhost:8080/api/payments";

export default function Payments({ onPaymentSuccess }) {
  const [formData, setFormData] = useState({
    fullName: "",
    email: "",
    amount: "",
  });
  const [status, setStatus] = useState("");

  function onChange(event) {
    setFormData((prev) => ({
      ...prev,
      [event.target.name]: event.target.value,
    }));
  }

  async function onSubmit(event) {
    event.preventDefault();
    setStatus("");

    const response = await fetch(API_URL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formData),
    });

    if (!response.ok) {
      setStatus("Nie udalo sie wyslac platnosci");
      return;
    }

    const data = await response.json();
    setStatus(data.message);
    onPaymentSuccess(data.message);
    setFormData({ fullName: "", email: "", amount: "" });
  }

  return (
    <section>
      <h2>Platnosci</h2>
      <form onSubmit={onSubmit}>
        <input
          name="fullName"
          placeholder="Imie i nazwisko"
          value={formData.fullName}
          onChange={onChange}
          required
        />
        <input
          name="email"
          type="email"
          placeholder="Email"
          value={formData.email}
          onChange={onChange}
          required
        />
        <input
          name="amount"
          type="number"
          step="0.01"
          placeholder="Kwota"
          value={formData.amount}
          onChange={onChange}
          required
        />
        <button type="submit">Wyslij</button>
      </form>
      {status ? <p>{status}</p> : null}
    </section>
  );
}
