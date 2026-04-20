import { useState } from "react";
import Products from "./components/Products";
import Payments from "./components/Payments";

export default function App() {
  const [status, setStatus] = useState("");

  return (
    <div className="container">
      <h1>Zadanie 5</h1>
      <Products />
      <Payments onPaymentSuccess={setStatus} />
      {status ? <p className="status">{status}</p> : null}
    </div>
  );
}
