import { useState } from "react";
import { useShop } from "../hooks/ShopContext";

const INITIAL_FORM = { fullName: "", email: "", amount: "" };

export default function Payments() {
  const { paymentStatus, sendPayment } = useShop();
  const [formData, setFormData] = useState(INITIAL_FORM);

  function onChange(event) {
    setFormData((prev) => ({
      ...prev,
      [event.target.name]: event.target.value,
    }));
  }

  async function onSubmit(event) {
    event.preventDefault();
    const success = await sendPayment(formData);
    if (success) {
      setFormData(INITIAL_FORM);
    }
  }

  return (
    <section>
      <h2>Platnosci</h2>
      <form onSubmit={onSubmit}>
        <label htmlFor="fullName">Imie i nazwisko</label>
        <input
          id="fullName"
          name="fullName"
          placeholder="Imie i nazwisko"
          value={formData.fullName}
          onChange={onChange}
          required
          maxLength={200}
        />
        <label htmlFor="email">Email</label>
        <input
          id="email"
          name="email"
          type="email"
          placeholder="Email"
          value={formData.email}
          onChange={onChange}
          required
          maxLength={320}
        />
        <label htmlFor="amount">Kwota</label>
        <input
          id="amount"
          name="amount"
          type="number"
          step="0.01"
          min="0.01"
          placeholder="Kwota"
          value={formData.amount}
          onChange={onChange}
          required
        />
        <button type="submit">Wyslij</button>
      </form>
      {paymentStatus ? <p>{paymentStatus}</p> : null}
    </section>
  );
}
