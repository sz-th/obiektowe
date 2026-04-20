import { createContext, useContext, useEffect, useMemo, useState } from "react";

const API_PRODUCTS_URL = "http://localhost:8080/api/products";
const API_CART_URL = "http://localhost:8080/api/cart";
const API_PAYMENTS_URL = "http://localhost:8080/api/payments";

const ShopContext = createContext(null);

export function ShopProvider({ children }) {
  const [products, setProducts] = useState([]);
  const [productsError, setProductsError] = useState("");
  const [cartItems, setCartItems] = useState([]);
  const [paymentStatus, setPaymentStatus] = useState("");
  const [cartStatus, setCartStatus] = useState("");
  const [status, setStatus] = useState("");

  useEffect(() => {
    async function loadProducts() {
      try {
        const response = await fetch(API_PRODUCTS_URL);
        if (!response.ok) {
          throw new Error("Nie udalo sie pobrac produktow");
        }
        const data = await response.json();
        setProducts(data);
        setProductsError("");
      } catch (error) {
        setProductsError("Nie udalo sie pobrac produktow");
      }
    }

    loadProducts();
  }, []);

  function addToCart(product) {
    setCartItems((prev) => [...prev, product]);
  }

  async function sendCart() {
    setCartStatus("");
    try {
      const response = await fetch(API_CART_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          items: cartItems,
          total: cartItems.reduce((sum, item) => sum + item.price, 0),
        }),
      });

      if (!response.ok) {
        setCartStatus("Nie udalo sie wyslac koszyka");
        setStatus("Nie udalo sie wyslac koszyka");
        return false;
      }

      const data = await response.json();
      setCartStatus(data.message);
      setStatus(data.message);
      return true;
    } catch (error) {
      setCartStatus("Nie udalo sie wyslac koszyka");
      setStatus("Nie udalo sie wyslac koszyka");
      return false;
    }
  }

  async function sendPayment(formData) {
    setPaymentStatus("");
    try {
      const response = await fetch(API_PAYMENTS_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ...formData,
          amount: Number(formData.amount),
        }),
      });

      if (!response.ok) {
        setPaymentStatus("Nie udalo sie wyslac platnosci");
        setStatus("Nie udalo sie wyslac platnosci");
        return false;
      }

      const data = await response.json();
      setPaymentStatus(data.message);
      setStatus(data.message);
      return true;
    } catch (error) {
      setPaymentStatus("Nie udalo sie wyslac platnosci");
      setStatus("Nie udalo sie wyslac platnosci");
      return false;
    }
  }

  const value = useMemo(
    () => ({
      products,
      productsError,
      cartItems,
      cartStatus,
      paymentStatus,
      status,
      addToCart,
      sendCart,
      sendPayment,
    }),
    [products, productsError, cartItems, cartStatus, paymentStatus, status],
  );

  return <ShopContext.Provider value={value}>{children}</ShopContext.Provider>;
}

export function useShop() {
  const context = useContext(ShopContext);
  if (!context) {
    throw new Error("useShop must be used within ShopProvider");
  }
  return context;
}
