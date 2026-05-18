import { BrowserRouter, Link, Route, Routes } from "react-router-dom";
import Products from "./components/Products";
import Cart from "./components/Cart";
import Payments from "./components/Payments";
import { ShopProvider } from "./hooks/ShopContext";

export default function App() {
  return (
    <BrowserRouter>
      <ShopProvider>
        <div className="container">
          <h1>Zadanie 6</h1>
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
        </div>
      </ShopProvider>
    </BrowserRouter>
  );
}
