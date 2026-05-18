import { useShop } from "../hooks/ShopContext";

function formatPrice(value) {
  return `${value.toFixed(2)} PLN`;
}

export default function Products() {
  const { products, productsError, addToCart } = useShop();

  return (
    <section>
      <h2>Produkty</h2>
      {productsError ? <p>{productsError}</p> : null}
      <ul>
        {products.map((product) => (
          <li key={product.id}>
            {product.name} - {formatPrice(product.price)}
            <button type="button" onClick={() => addToCart(product)}>
              Dodaj do koszyka
            </button>
          </li>
        ))}
      </ul>
    </section>
  );
}
