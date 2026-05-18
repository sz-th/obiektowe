import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";
import axios from "axios";

const API_BASE_URL =
  import.meta.env.VITE_API_BASE_URL || "http://localhost:8080";
const API_PRODUCTS_URL = `${API_BASE_URL}/api/products`;
const API_CART_URL = `${API_BASE_URL}/api/cart`;
const API_PAYMENTS_URL = `${API_BASE_URL}/api/payments`;

const PRODUCTS_ERROR_MSG = "Nie udalo sie pobrac produktow";
const CART_ERROR_MSG = "Nie udalo sie wyslac koszyka";
const PAYMENT_ERROR_MSG = "Nie udalo sie wyslac platnosci";

const ShopContext = createContext(null);

function generateCartId() {
  if (
    typeof crypto !== "undefined" &&
    typeof crypto.randomUUID === "function"
  ) {
    return crypto.randomUUID();
  }
  return `${Date.now()}-${Math.random().toString(36).slice(2)}`;
}

export function ShopProvider({ children }) {
  const [products, setProducts] = useState([]);
  const [productsError, setProductsError] = useState("");
  const [cartItems, setCartItems] = useState([]);
  const [paymentStatus, setPaymentStatus] = useState("");
  const [cartStatus, setCartStatus] = useState("");
  const [status, setStatus] = useState("");

  useEffect(() => {
    let cancelled = false;

    async function loadProducts() {
      try {
        const response = await axios.get(API_PRODUCTS_URL);
        if (cancelled) return;
        if (response.status == 200) {
          setProducts(response.data);
          setProductsError("");
        }
      } catch (error) {
        if (cancelled) return;
        console.error("loadProducts failed", error);
        setProductsError(PRODUCTS_ERROR_MSG);
      }
    }

    loadProducts();
    return () => {
      cancelled = true;
    };
  }, []);

  const addToCart = useCallback((product) => {
    setCartItems((prev) => [...prev, { ...product, cartId: generateCartId() }]);
  }, []);

  const sendCart = useCallback(async () => {
    setCartStatus("");
    try {
      const items = cartItems.map((item) => ({
        id: item.id,
        name: item.name,
        price: item.price,
      }));
      const response = await axios.post(API_CART_URL, { items });
      setCartStatus(response.data.message);
      setStatus(response.data.message);
      return true;
    } catch (error) {
      console.error("sendCart failed", error);
      setCartStatus(CART_ERROR_MSG);
      setStatus(CART_ERROR_MSG);
      return false;
    }
  }, [cartItems]);

  const sendPayment = useCallback(async (formData) => {
    setPaymentStatus("");
    try {
      const response = await axios.post(API_PAYMENTS_URL, {
        ...formData,
        amount: parseInt(formData.amount),
      });
      setPaymentStatus(response.data.message);
      setStatus(response.data.message);
      return true;
    } catch (error) {
      console.error("sendPayment failed", error);
      setPaymentStatus(PAYMENT_ERROR_MSG);
      setStatus(PAYMENT_ERROR_MSG);
      return false;
    }
  }, []);

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
    [
      products,
      productsError,
      cartItems,
      cartStatus,
      paymentStatus,
      status,
      addToCart,
      sendCart,
      sendPayment,
    ],
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
