import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";
import { apiClient, API_PATHS } from "../api/client";
import { parseMessage, parseProducts } from "../api/validators";
import { generateCartId } from "../utils/id";

const PRODUCTS_ERROR_MSG = "Nie udalo sie pobrac produktow";
const CART_ERROR_MSG = "Nie udalo sie wyslac koszyka";
const PAYMENT_ERROR_MSG = "Nie udalo sie wyslac platnosci";

const ShopContext = createContext(null);
ShopContext.displayName = "ShopContext";

function parseAmount(value) {
  const amount = Number.parseFloat(value);
  if (!Number.isFinite(amount) || amount <= 0) {
    return null;
  }
  return amount;
}

export function ShopProvider({ children }) {
  const [products, setProducts] = useState([]);
  const [productsError, setProductsError] = useState("");
  const [cartItems, setCartItems] = useState([]);
  const [cartStatus, setCartStatus] = useState("");
  const [paymentStatus, setPaymentStatus] = useState("");

  useEffect(() => {
    const controller = new AbortController();

    async function loadProducts() {
      try {
        const response = await apiClient.get(API_PATHS.products, {
          signal: controller.signal,
        });
        const parsed = parseProducts(response.data);
        if (parsed === null) {
          setProductsError(PRODUCTS_ERROR_MSG);
          return;
        }
        setProducts(parsed);
        setProductsError("");
      } catch {
        if (controller.signal.aborted) {
          return;
        }
        setProductsError(PRODUCTS_ERROR_MSG);
      }
    }

    loadProducts();
    return () => controller.abort();
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
      const response = await apiClient.post(API_PATHS.cart, { items });
      const message = parseMessage(response.data);
      if (message === null) {
        setCartStatus(CART_ERROR_MSG);
        return false;
      }
      setCartStatus(message);
      return true;
    } catch {
      setCartStatus(CART_ERROR_MSG);
      return false;
    }
  }, [cartItems]);

  const sendPayment = useCallback(async (formData) => {
    setPaymentStatus("");
    const amount = parseAmount(formData.amount);
    if (amount === null) {
      setPaymentStatus(PAYMENT_ERROR_MSG);
      return false;
    }

    try {
      const response = await apiClient.post(API_PATHS.payments, {
        fullName: formData.fullName.trim(),
        email: formData.email.trim(),
        amount,
      });
      const message = parseMessage(response.data);
      if (message === null) {
        setPaymentStatus(PAYMENT_ERROR_MSG);
        return false;
      }
      setPaymentStatus(message);
      return true;
    } catch {
      setPaymentStatus(PAYMENT_ERROR_MSG);
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
