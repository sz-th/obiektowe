import { useState } from "react";
import { BrowserRouter, Link, Route, Routes } from "react-router-dom";
import Products from "./components/Products";
import Cart from "./components/Cart";
import Payments from "./components/Payments";

export default function App() {
  const [status, setStatus] = useState("");
  const [cartItems, setCartItems] = useState([]);

  function addToCart(product) {
    setCartItems((prev) => [...prev, product]);
  }

  return (
    <BrowserRouter>
      <div className="container">
        <h1>Zadanie 5</h1>
        <nav className="nav">
          <Link to="/">Produkty</Link>
          <Link to="/cart">Koszyk</Link>
          <Link to="/payments">Platnosci</Link>
        </nav>
        <Routes>
          <Route path="/" element={<Products onAddToCart={addToCart} />} />
          <Route path="/cart" element={<Cart items={cartItems} />} />
          <Route
            path="/payments"
            element={<Payments onPaymentSuccess={setStatus} />}
          />
        </Routes>
        {status ? <p className="status">{status}</p> : null}
      </div>
    </BrowserRouter>
  );
}
