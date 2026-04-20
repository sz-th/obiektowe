import { useEffect, useState } from "react";

const API_URL = "http://localhost:8080/api/products";

export default function Products({ onAddToCart }) {
  const [products, setProducts] = useState([]);
  const [error, setError] = useState("");

  useEffect(() => {
    async function loadProducts() {
      try {
        const response = await fetch(API_URL);
        if (!response.ok) {
          throw new Error("Nie udalo sie pobrac produktow");
        }
        const data = await response.json();
        setProducts(data);
      } catch (err) {
        setError(err.message);
      }
    }

    loadProducts();
  }, []);

  return (
    <section>
      <h2>Produkty</h2>
      {error ? <p>{error}</p> : null}
      <ul>
        {products.map((product) => (
          <li key={product.id}>
            {product.name} - {product.price.toFixed(2)} PLN
            <button type="button" onClick={() => onAddToCart(product)}>
              Dodaj do koszyka
            </button>
          </li>
        ))}
      </ul>
    </section>
  );
}
