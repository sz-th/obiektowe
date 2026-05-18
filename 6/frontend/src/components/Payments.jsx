import { useState } from "react";
import { useShop } from "../hooks/ShopContext";

const INITIAL_FORM = { fullName: "", email: "", amount: "" };

const FORM_FIELDS = [
  {
    id: "fullName",
    name: "fullName",
    label: "Imie i nazwisko",
    placeholder: "Imie i nazwisko",
    type: "text",
    maxLength: 200,
  },
  {
    id: "email",
    name: "email",
    label: "Email",
    placeholder: "Email",
    type: "email",
    maxLength: 320,
  },
  {
    id: "amount",
    name: "amount",
    label: "Kwota",
    placeholder: "Kwota",
    type: "number",
    step: "0.01",
    min: "0.01",
  },
];

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
        {FORM_FIELDS.map((field) => (
          <div key={field.id}>
            <label htmlFor={field.id}>{field.label}</label>
            <input
              id={field.id}
              name={field.name}
              type={field.type}
              placeholder={field.placeholder}
              value={formData[field.name]}
              onChange={onChange}
              required
              maxLength={field.maxLength}
              step={field.step}
              min={field.min}
            />
          </div>
        ))}
        <button type="submit">Wyslij</button>
      </form>
      {paymentStatus ? <p>{paymentStatus}</p> : null}
    </section>
  );
}
