import { useState } from "react";
import { useShop } from "../hooks/ShopContext";

export default function Payments() {
  const { paymentStatus, sendPayment } = useShop();
  const [formData, setFormData] = useState({
    fullName: "",
    email: "",
    amount: "",
  });

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
      setFormData({ fullName: "", email: "", amount: "" });
    }
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
      {paymentStatus ? <p>{paymentStatus}</p> : null}
    </section>
  );
}
