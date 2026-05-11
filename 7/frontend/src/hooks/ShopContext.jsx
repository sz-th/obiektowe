import { createContext, useContext, useEffect, useMemo, useState } from "react";
import axios from "axios";

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
        const response = await axios.get(API_PRODUCTS_URL);
        setProducts(response.data);
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
      const response = await axios.post(API_CART_URL, {
        items: cartItems,
        total: cartItems.reduce((sum, item) => sum + item.price, 0),
      });

      setCartStatus(response.data.message);
      setStatus(response.data.message);
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
      const response = await axios.post(API_PAYMENTS_URL, {
        ...formData,
        amount: Number(formData.amount),
      });

      setPaymentStatus(response.data.message);
      setStatus(response.data.message);
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
