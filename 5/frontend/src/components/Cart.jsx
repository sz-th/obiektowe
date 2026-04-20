import { useShop } from "../hooks/ShopContext";

export default function Cart() {
  const { cartItems, cartStatus, sendCart } = useShop();
  const total = cartItems.reduce((sum, item) => sum + item.price, 0);

  async function onSubmit() {
    await sendCart();
  }

  return (
    <section>
      <h2>Koszyk</h2>
      {cartItems.length === 0 ? (
        <p>Koszyk jest pusty</p>
      ) : (
        <ul>
          {cartItems.map((item, index) => (
            <li key={`${item.id}-${index}`}>
              {item.name} - {item.price.toFixed(2)} PLN
            </li>
          ))}
        </ul>
      )}
      <p>Razem: {total.toFixed(2)} PLN</p>
      <button type="button" disabled={cartItems.length === 0} onClick={onSubmit}>
        Wyslij koszyk
      </button>
      {cartStatus ? <p>{cartStatus}</p> : null}
    </section>
  );
}
