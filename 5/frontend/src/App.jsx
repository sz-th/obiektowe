import { BrowserRouter, Link, Route, Routes } from "react-router-dom";
import Products from "./components/Products";
import Cart from "./components/Cart";
import Payments from "./components/Payments";
import { ShopProvider, useShop } from "./hooks/ShopContext";

function AppContent() {
  const { status } = useShop();
  return (
    <div className="container">
      <h1>Zadanie 5</h1>
      <nav className="nav">
        <Link to="/">Produkty</Link>
        <Link to="/cart">Koszyk</Link>
        <Link to="/payments">Platnosci</Link>
      </nav>
      <Routes>
        <Route path="/" element={<Products />} />
        <Route path="/cart" element={<Cart />} />
        <Route path="/payments" element={<Payments />} />
      </Routes>
      {status ? <p className="status">{status}</p> : null}
    </div>
  );
}

export default function App() {
  return (
    <BrowserRouter>
      <ShopProvider>
        <AppContent />
      </ShopProvider>
    </BrowserRouter>
  );
}
